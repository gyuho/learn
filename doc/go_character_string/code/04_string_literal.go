package main

import "fmt"

func main() {
	txt1 := "\xE2\x88\x83y \xE2\x88\x80x \xC2\xAC(x \xE2\x89\xBA y)"
	txt2 := "∃y ∀x ¬(x ≺ y)"
	fmt.Println(txt1)         // ∃y ∀x ¬(x ≺ y)
	fmt.Println(txt2)         // ∃y ∀x ¬(x ≺ y)
	fmt.Printf("%x\n", txt2)  // e288837920e288807820c2ac287820e289ba207929
	fmt.Println(txt1 == txt2) // true
}
