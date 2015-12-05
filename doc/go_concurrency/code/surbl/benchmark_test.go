package surbl

import (
	"log"
	"net"
	"runtime"
	"strings"
	"sync"
	"testing"
)

func checkNoCurrency(domains ...string) map[string]DomainInfo {
	final := make(map[string]DomainInfo)
	for _, domain := range domains {
		dom := hosten(domain)
		dmToLook := dom + ".multi.surbl.org"
		ads, err := net.LookupHost(dmToLook)
		if err != nil {
			switch err.(type) {
			case net.Error:
				if err.(*net.DNSError).Err == "no such host" {
					final[dom] = nonSpam
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
			final[dom] = info
		}
	}
	return final
}

func BenchmarkCheckNoCuncurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkNoCurrency(goodSlice...)
		checkNoCurrency(spamSlice...)
	}
}

// checkWithLockBatch checks SURLB list distributing by the number of CPU.
// This is slower. Just use goroutine per each domain.
func checkWithLockBatch(domains ...string) map[string]DomainInfo {
	cpus := runtime.NumCPU()
	length := len(domains)
	if cpus > length {
		cpus = 1
	}
	final := make(map[string]DomainInfo)
	if cpus > 1 {
		finals := []map[string]DomainInfo{}
		var wg sync.WaitGroup
		wg.Add(cpus)
		var mutex sync.Mutex
		bucketCap := length / cpus
		for i := 0; i < cpus; i++ {
			var newIdx1, newIdx2 int
			if i == cpus-1 {
				newIdx1 = bucketCap * i
				newIdx2 = len(domains)
			} else {
				newIdx1 = bucketCap * i
				newIdx2 = bucketCap * (i + 1)
			}
			mutex.Lock()
			newSlice := domains[newIdx1:newIdx2]
			mutex.Unlock()
			go func(slice []string, mutex *sync.Mutex) {
				defer wg.Done()
				// maps are not thread-safe
				tempmap := make(map[string]DomainInfo)
				for idx, dm := range slice {
					if idx != 0 && idx%500 == 0 {
						log.Println("Processing:", idx, "/", len(slice))
					}
					dom := hosten(dm)
					dmToLook := dom + ".multi.surbl.org"
					ads, err := net.LookupHost(dmToLook)
					if err != nil {
						switch err.(type) {
						case net.Error:
							if err.(*net.DNSError).Err == "no such host" {
								mutex.Lock()
								final[dom] = nonSpam
								mutex.Unlock()
								continue
								// next lines will be skipped
							}
						default:
							log.Fatal(err)
						}
					}
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
					tempmap[dom] = info
					mutex.Unlock()
				}
				mutex.Lock()
				finals = append(finals, tempmap)
				mutex.Unlock()
			}(newSlice, &mutex)
		}
		wg.Wait()
		for _, elem := range finals {
			for key, val := range elem {
				mutex.Lock()
				final[key] = val
				mutex.Unlock()
				runtime.Gosched()
			}
		}
	} else {
		// Do without concurrency
		for idx, dm := range domains {
			if idx != 0 && idx%500 == 0 {
				log.Println("Processing:", idx, "/", len(domains))
			}
			dom := hosten(dm)
			dmToLook := dom + ".multi.surbl.org"
			ads, err := net.LookupHost(dmToLook)
			if err != nil {
				switch err.(type) {
				case net.Error:
					if err.(*net.DNSError).Err == "no such host" {
						final[dom] = nonSpam
						continue
						// next lines will be skipped
					}
				default:
					log.Fatal(err)
				}
			}
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
			final[dom] = info
		}
	}
	return final
}

func BenchmarkCheckWithLockBatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkWithLockBatch(goodSlice...)
		checkWithLockBatch(spamSlice...)
	}
}

func BenchmarkCheckWithLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckWithLock(goodSlice...)
		CheckWithLock(spamSlice...)
	}
}

func BenchmarkCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Check(goodSlice...)
		Check(spamSlice...)
	}
}
