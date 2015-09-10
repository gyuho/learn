/*
go run -race 31_no_race_surbl_with_mutex.go
*/
package main

import (
	"log"
	"net"
	"net/url"
	"strings"
	"sync"
)

// Data is a set of data in map data structure.
// Every element is unique, and it is unordered.
// It maps its value to frequency.
type Data struct {
	// m maps an element to its frequency
	m map[interface{}]int

	// RWMutex is more expensive
	// https://blogs.oracle.com/roch/entry/beware_of_the_performance_of
	// sync.RWMutex
	//
	// to synchronize access to shared state across multiple goroutines.
	//
	sync.Mutex
}

// NewData returns a new Data.
// Map supports the built-in function "make"
// so we do not have to use "new" and
// "make" does not return pointer.
func NewData() *Data {
	nmap := make(map[interface{}]int)
	return &Data{
		m: nmap,
	}
	// return make(Data)
}

// Init initializes the Data.
func (d *Data) Init() {
	// (X) d = NewData()
	// This only updates its pointer
	// , not the Data itself
	//
	*d = *NewData()
}

// GetSize returns the size of set.
func (d Data) GetSize() int {
	return len(d.m)
}

// IsEmpty returns true if the set is empty.
func (d Data) IsEmpty() bool {
	return d.GetSize() == 0
}

// Insert insert values to the set.
func (d *Data) Insert(items ...interface{}) {
	for _, value := range items {
		d.Lock()
		v, ok := d.m[value]
		d.Unlock()
		if ok {
			d.Lock()
			d.m[value] = v + 1
			d.Unlock()
			continue
		}
		d.Lock()
		d.m[value] = 1
		d.Unlock()
	}
}

func main() {
	d := NewData()
	d.Insert(1, 2, -.9, "A", 0, 2, 2, 2)
	if d.IsEmpty() {
		log.Fatalf("IsEmpty() should return false: %#v", d)
	}
	if d.GetSize() != 5 {
		log.Fatalf("GetSize() should return 5: %#v", d)
	}

	rmap2 := Check(goodSlice...)
	for k, v := range rmap2 {
		if v.IsSpam {
			log.Fatalf("Check | Unexpected %+v %+v but it's ok", k, v)
		}
	}
}

var goodSlice = []string{
	"google.com",
}

// DomainInfo contains domain information from Surbl.org.
type DomainInfo struct {
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
