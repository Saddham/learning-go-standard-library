package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {

	// Basic math

	fmt.Println(math.Pi)
	fmt.Println(math.Ceil(math.Pi))
	fmt.Println(math.Floor(math.Pi))
	fmt.Println(math.Trunc(math.Pi))
	fmt.Println(math.Max(math.Pi, math.Ln2))
	fmt.Println(math.Min(math.Pi, math.Ln2))
	fmt.Println(math.Mod(17.25, 5.0)) // Mod with floats

	fmt.Printf("%.1f\n", math.Round(10.5))
	fmt.Printf("%.1f\n", math.Round(-10.5))
	fmt.Printf("%.1f\n", math.RoundToEven(10.5))
	fmt.Printf("%.1f\n", math.RoundToEven(-10.5))

	x := 10.0

	fmt.Println(math.Abs(x), math.Abs(-x))
	fmt.Println(math.Pow(x, 2))
	fmt.Println(math.E, math.Exp(x))

	// Trignometry

	fmt.Println(math.Cos(math.Pi))
	fmt.Println(math.Sin(math.Pi / 2))
	fmt.Println(math.Tan(math.Pi / 4))

	fmt.Println(math.Log(10))       // natural log of 10
	fmt.Println(math.Hypot(30, 40)) // Hypotenuse of a triangle

	// Square and cube roots
	fmt.Println(math.Sqrt(25))
	fmt.Println(math.Cbrt(125))

	// Random numbers
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())

	// Create permutations
	s := []string{"apple", "pear", "grape", "orange", "kiwi", "melon"}
	indexes := rand.Perm(len(s))
	fmt.Println(indexes)

	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " ")
	}
	fmt.Println()

	const a, b = 10, 50
	for i := 0; i < 10; i++ {
		n := a + rand.Intn(b-a+1)
		fmt.Print(n, " ")
	}
	fmt.Println()

	for i := 0; i < 10; i++ {
		c := string('A' + rune(rand.Intn('Z'-'A'+1)))
		fmt.Printf("%s ", c)
	}
	fmt.Println()

	const numstring = "one two three four five six"
	words := strings.Fields(numstring)
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}
