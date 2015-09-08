package main

import (
	"fmt"
	"math"
)

// Cosine converts texts to vectors
// associatting each chracter with its frequncy
// and caculates cosine similarities.
// (https://en.wikipedia.org/wiki/Cosine_similarity)
func Cosine(txt1, txt2 []byte) float64 {
	vect1 := make(map[byte]int)
	for _, t := range txt1 {
		vect1[t]++
	}
	vect2 := make(map[byte]int)
	for _, t := range txt2 {
		vect2[t]++
	}
	//
	// dot-product two vectors
	// map[byte]int return 0 for non-existing key
	// and if two texts are equal, product will be highest
	// and if two texts are totally different, it will be 0
	//
	// to calculate A·B
	dotProduct := 0.0
	for k, v := range vect1 {
		dotProduct += float64(v) * float64(vect2[k])
	}
	// to calculate |A|*|B|
	sum1 := 0.0
	for _, v := range vect1 {
		sum1 += math.Pow(float64(v), 2)
	}
	sum2 := 0.0
	for _, v := range vect2 {
		sum2 += math.Pow(float64(v), 2)
	}
	magnitude := math.Sqrt(sum1) * math.Sqrt(sum2)
	if magnitude == 0 {
		return 0.0
	}
	return float64(dotProduct) / float64(magnitude)
}

func main() {
	fmt.Println(Cosine([]byte("Hello"), []byte("Hello")))
	// 0.9999999999999999

	fmt.Println(Cosine([]byte(""), []byte("Hello")))
	// 0

	fmt.Println(Cosine([]byte("hello"), []byte("Hello")))
	// 0.857142857142857

	fmt.Println(Cosine([]byte("abc"), []byte("bcd")))
	// 0.6666666666666667

	fmt.Println(Cosine([]byte("Hello"), []byte("Hel lo")))
	// 0.9354143466934852
}
