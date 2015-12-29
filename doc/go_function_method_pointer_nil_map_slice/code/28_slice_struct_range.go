package main

import "fmt"

type event struct {
	val int
}

func main() {
	evs := []event{
		event{1},
		event{2},
		event{3},
	}
	for _, ev := range evs {
		ev.val = 10
	}
	fmt.Println(evs) // [{1} {2} {3}]

	for i := range evs {
		evs[i].val = 100
	}
	fmt.Println(evs) // [{100} {100} {100}]
}
