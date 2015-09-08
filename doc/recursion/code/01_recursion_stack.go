package main

import "fmt"

var keys = []string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
}

func recursion(index int, rmap *map[string]string) {
	if index == len(keys) {
		fmt.Println()
		fmt.Println("recursion is done")
		fmt.Println()
		return
	}

	fmt.Printf("beginning recursion with index %d / key %s / map %v\n", index, keys[index], (*rmap)[keys[index]])

	recursion(index+1, rmap)

	(*rmap)[keys[index]] = "done"
	fmt.Printf("after     recursion with index %d / key %s / map %v\n", index, keys[index], (*rmap)[keys[index]])
}

func main() {
	executed := make(map[string]string)
	for _, k := range keys {
		executed[k] = "not yet"
	}
	recursion(0, &executed)
}

/*
beginning recursion with index 0 / key A / map not yet
beginning recursion with index 1 / key B / map not yet
beginning recursion with index 2 / key C / map not yet
beginning recursion with index 3 / key D / map not yet
beginning recursion with index 4 / key E / map not yet
beginning recursion with index 5 / key F / map not yet
beginning recursion with index 6 / key G / map not yet
beginning recursion with index 7 / key H / map not yet
beginning recursion with index 8 / key I / map not yet

recursion is done

after     recursion with index 8 / key I / map done
after     recursion with index 7 / key H / map done
after     recursion with index 6 / key G / map done
after     recursion with index 5 / key F / map done
after     recursion with index 4 / key E / map done
after     recursion with index 3 / key D / map done
after     recursion with index 2 / key C / map done
after     recursion with index 1 / key B / map done
after     recursion with index 0 / key A / map done
*/
