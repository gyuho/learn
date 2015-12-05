package count

import "testing"

func TestNaiveCounter(t *testing.T) {
	counter := new(NaiveCounter)
	if counter.Get() != 0.0 {
		t.Fatalf("Expected 0.0 but %v", counter.Get())
	}
	counter.Add(1.7)
	if counter.Get() != 1.7 {
		t.Fatalf("Expected 1.7 but %v", counter.Get())
	}
	counter.Add(-5.5)
	if counter.Get() != -3.8 {
		t.Fatalf("Expected -3.8 but %v", counter.Get())
	}
	counter.Add(100)
	if counter.Get() != 96.2 {
		t.Fatalf("Expected 96.2 but %v", counter.Get())
	}
}

func TestNaiveCounterRace(t *testing.T) {
	//
	// this will cause race conditions for:
	// go test -race
	//

	// counter := new(NaiveCounter)
	// var wg sync.WaitGroup
	// wg.Add(10)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		counter.Add(1.7)
	// 		counter.Get()
	// 		counter.Add(-5.5)
	// 		counter.Get()
	// 		counter.Add(100)
	// 		counter.Get()
	// 	}()
	// }
	// wg.Wait()
}

func TestMutexCounter(t *testing.T) {
	counter := new(MutexCounter)
	if counter.Get() != 0.0 {
		t.Fatalf("Expected 0.0 but %v", counter.Get())
	}
	counter.Add(1.7)
	if counter.Get() != 1.7 {
		t.Fatalf("Expected 1.7 but %v", counter.Get())
	}
	counter.Add(-5.5)
	if counter.Get() != -3.8 {
		t.Fatalf("Expected -3.8 but %v", counter.Get())
	}
	counter.Add(100)
	if counter.Get() != 96.2 {
		t.Fatalf("Expected 96.2 but %v", counter.Get())
	}
}

func TestRWMutexCounter(t *testing.T) {
	counter := new(RWMutexCounter)
	if counter.Get() != 0.0 {
		t.Fatalf("Expected 0.0 but %v", counter.Get())
	}
	counter.Add(1.7)
	if counter.Get() != 1.7 {
		t.Fatalf("Expected 1.7 but %v", counter.Get())
	}
	counter.Add(-5.5)
	if counter.Get() != -3.8 {
		t.Fatalf("Expected -3.8 but %v", counter.Get())
	}
	counter.Add(100)
	if counter.Get() != 96.2 {
		t.Fatalf("Expected 96.2 but %v", counter.Get())
	}
}

func TestAtomicIntCounter(t *testing.T) {
	counter := new(AtomicIntCounter)
	if counter.Get() != 0.0 {
		t.Fatalf("Expected 0.0 but %v", counter.Get())
	}
	counter.Add(1.7)
	if counter.Get() != 1.0 {
		t.Fatalf("Expected 1.0 but %v", counter.Get())
	}
	counter.Add(-5.5)
	if counter.Get() != -4.0 {
		t.Fatalf("Expected -4.0 but %v", counter.Get())
	}
	counter.Add(100)
	if counter.Get() != 96 {
		t.Fatalf("Expected 96 but %v", counter.Get())
	}
}

func TestAtomicCounter(t *testing.T) {
	counter := new(AtomicCounter)
	if counter.Get() != 0.0 {
		t.Fatalf("Expected 0.0 but %v", counter.Get())
	}
	counter.Add(1.7)
	if counter.Get() != 1.7 {
		t.Fatalf("Expected 1.7 but %v", counter.Get())
	}
	counter.Add(-5.5)
	if counter.Get() != -3.8 {
		t.Fatalf("Expected -3.8 but %v", counter.Get())
	}
	counter.Add(100)
	if counter.Get() != 96.2 {
		t.Fatalf("Expected 96.2 but %v", counter.Get())
	}
}

func TestChannelCounter(t *testing.T) {
	counter := NewChannelCounter(0)
	defer counter.Done()
	defer counter.Close()
	if counter.Get() != 0.0 {
		t.Fatalf("Expected 0.0 but %v", counter.Get())
	}
	counter.Add(1.7)
	if counter.Get() != 1.7 {
		t.Fatalf("Expected 1.7 but %v", counter.Get())
	}
	counter.Add(-5.5)
	if counter.Get() != -3.8 {
		t.Fatalf("Expected -3.8 but %v", counter.Get())
	}
	counter.Add(100)
	if counter.Get() != 96.2 {
		t.Fatalf("Expected 96.2 but %v", counter.Get())
	}
}
