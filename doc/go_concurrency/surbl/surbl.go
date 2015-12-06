package surbl

import (
	"log"
	"net"
	"net/url"
	"strings"
	"sync"
)

// DomainInfo contains domain information from Surbl.org.
type DomainInfo struct {
	Domain string
	IsSpam bool
	Types  []string
}

var nonSpam = DomainInfo{
	IsSpam: false,
	Types:  []string{"none"},
}

var addressMap = map[string]string{
	"2":  "SC: SpamCop web sites",
	"4":  "WS: sa-blacklist web sited",
	"8":  "AB: AbuseButler web sites",
	"16": "PH: Phishing sites",
	"32": "MW: Malware sites",
	"64": "JP: jwSpamSpy + Prolocation sites",
	"68": "WS JP: sa-blacklist web sited jwSpamSpy + Prolocation sites",
}

// Check concurrently checks SURBL spam list.
// http://www.surbl.org/guidelines
// http://www.surbl.org/surbl-analysis
func Check(domains ...string) map[string]DomainInfo {
	ch := make(chan DomainInfo)
	for _, domain := range domains {
		go func(domain string) {
			dom := hosten(domain)
			dmToLook := dom + ".multi.surbl.org"
			ads, err := net.LookupHost(dmToLook)
			if err != nil {
				switch err.(type) {
				case net.Error:
					if err.(*net.DNSError).Err == "no such host" {
						nonSpam := DomainInfo{
							Domain: domain,
							IsSpam: false,
							Types:  []string{"none"},
						}
						ch <- nonSpam
					}
				default:
					log.Fatal(err)
				}
			} else {
				stypes := []string{}
				for _, add := range ads {
					tempSlice := strings.Split(add, ".")
					flag := tempSlice[len(tempSlice)-1]
					if val, ok := addressMap[flag]; !ok {
						stypes = append(stypes, "unknown_source")
					} else {
						stypes = append(stypes, val)
					}
				}
				info := DomainInfo{
					Domain: domain,
					IsSpam: true,
					Types:  stypes,
				}
				ch <- info
			}
		}(domain)
	}
	final := make(map[string]DomainInfo)
	checkSize := len(domains)
	cn := 0
	for info := range ch {
		final[info.Domain] = info
		cn++
		if cn == checkSize {
			close(ch)
		}
	}
	return final
}

// CheckWithLock concurrently checks SURBL spam list with Mutex.
func CheckWithLock(domains ...string) map[string]DomainInfo {
	final := make(map[string]DomainInfo)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, domain := range domains {
		dom := hosten(domain)
		dmToLook := dom + ".multi.surbl.org"
		wg.Add(1)
		go func() {
			defer wg.Done()
			ads, err := net.LookupHost(dmToLook)
			if err != nil {
				switch err.(type) {
				case net.Error:
					if err.(*net.DNSError).Err == "no such host" {
						mutex.Lock()
						final[dom] = nonSpam
						mutex.Unlock()
					}
				default:
					log.Fatal(err)
				}
			} else {
				stypes := []string{}
				for _, add := range ads {
					tempSlice := strings.Split(add, ".")
					flag := tempSlice[len(tempSlice)-1]
					if val, ok := addressMap[flag]; !ok {
						stypes = append(stypes, "unknown_source")
					} else {
						stypes = append(stypes, val)
					}
				}
				info := DomainInfo{
					IsSpam: true,
					Types:  stypes,
				}
				mutex.Lock()
				final[dom] = info
				mutex.Unlock()
			}
		}()
	}
	wg.Wait()
	return final
}

// hosten returns the host of url.
func hosten(dom string) string {
	dom = strings.TrimSpace(dom)
	var domain string
	if strings.HasPrefix(dom, "http:") ||
		strings.HasPrefix(dom, "https:") {
		dmt, err := url.Parse(dom)
		if err != nil {
			log.Fatal(err)
		}
		domain = dmt.Host
	} else {
		domain = dom
	}
	return domain
}
