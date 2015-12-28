package main

import "fmt"

type keyIndex struct {
	generations []generation
}

type generation struct {
	slice []int
}

func main() {
	gn := generation{slice: []int{0, 1}}
	ki := keyIndex{generations: []generation{gn}}

	g := ki.generations[0]
	g.slice = append(g.slice, 5)
	fmt.Println(ki) // {[{[0 1]}]}

	gp := &ki.generations[0]
	gp.slice = append(gp.slice, 5)
	fmt.Println(ki) // {[{[0 1 5]}]}
}
