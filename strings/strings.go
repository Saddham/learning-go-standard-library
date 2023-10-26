package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"

	// ** Basic string operations ***

	// length of string
	fmt.Println(len(s))

	// iterate over each char
	for _, ch := range s {
		fmt.Print(string(ch), ",")
	}

	fmt.Println()

	// using operators < > == !=
	fmt.Println("dog" < "cat")
	fmt.Println("dog" > "cat")
	fmt.Println("dog" == "cat")
	fmt.Println("dog" != "cat")

	// comparing strings
	result := strings.Compare("dog", "cat")
	fmt.Println(result)
	result = strings.Compare("dog", "dog")
	fmt.Println(result)

	// equalFold tests using unicode case-folding
	fmt.Println(strings.EqualFold("dog", "cat"))
	fmt.Println(strings.EqualFold("dog", "Dog"))
	fmt.Println(strings.EqualFold("dog", "dog"))

	// toupper, tolower, title
	s1 := strings.ToUpper(s)
	s2 := strings.ToLower(s)
	s3 := strings.Title(s)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// *** String Searching ***

	// use contains to see if a substrin gis in a string
	fmt.Println(strings.Contains(s, "jump"))    // See if string contains given string
	fmt.Println(strings.ContainsAny(s, "abcd")) // See if string contains any of the individual chars in the substring

	// find the offset of the first instance of a substring
	fmt.Println(strings.Index(s, "fox"))
	fmt.Println(strings.Index(s, "cat"))

	vowels := "aeiouAEIOU"
	fmt.Println(strings.IndexAny("grzbl", vowels))
	fmt.Println(strings.IndexAny("Golang!", vowels))

	// Search for suffix/prefix
	fname := "filename.txt"
	fname2 := "temp_picfile.jpeg"

	fmt.Println(strings.HasSuffix(fname, "txt"))
	fmt.Println(strings.HasPrefix(fname2, "temp"))
	fmt.Println(strings.Count(s, "the")) // no. of occurrences of "the"
	fmt.Println(strings.Count(s, "he"))

	// *** String Manipulation ***

	sub1 := strings.Split(s, " ")
	fmt.Printf("%q\n", sub1)

	sub2 := strings.Split(s, "the")
	fmt.Printf("%q\n", sub2)

	s4 := []string{"one", "two", "three"}
	fmt.Println(strings.Join(s4, " - "))

	fmt.Printf("%q\n", strings.Fields(s)) // Split the string around whitespace

	// Split the string with callback
	s5 := "This is a string. With some punctionation, for a demo! Yep."
	fmt.Printf("%q\n", strings.FieldsFunc(s5, func(c rune) bool {
		return unicode.IsPunct(c)
	}))

	// String replacements
	rep := strings.NewReplacer(".", "|", ",", "|", "!", "|")
	fmt.Println(rep.Replace(s5))

	// String transformations
	shift := 2
	transform := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			value := (int(r) + shift)
			if value > 91 {
				value -= 26
			} else if value < 65 {
				value += 26
			}

			return rune(value)
		case r >= 'a' && r <= 'z':
			value := (int(r) + shift)
			if value > 122 {
				value -= 26
			} else if value < 97 {
				value += 26
			}

			return rune(value)
		}

		return r
	}

	encode := strings.Map(transform, s)
	fmt.Println(encode)

	shift = -shift
	decode := strings.Map(transform, encode)
	fmt.Println(decode)

	// String builder
	var sb strings.Builder

	sb.WriteString("This is a string 1\n")
	sb.WriteString("This is a string 2\n")
	sb.WriteString("This is a string 3\n")

	fmt.Println(sb.String())

	fmt.Println("Capacity:", sb.Cap())
	sb.Grow(1024)
	fmt.Println("Capacity:", sb.Cap())

	for i := 0; i <= 10; i++ {
		fmt.Fprint(&sb, "String ", i, " -- ")
	}

	fmt.Println(sb.String())
	fmt.Printf("Builder len - %d\n", sb.Len())

	sb.Reset()
	fmt.Println("After reset:")
	fmt.Println("Capacity: ", sb.Cap())
	fmt.Println("Length: ", sb.Len())

	// String conversions
	num := 100

	// Integer to string
	s7 := strconv.Itoa(num)
	fmt.Printf("%T, %v\n", s7, s7)

	// String to integer
	num2, err := strconv.Atoi(s7)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%T, %v\n", num2, num2)

	// String to primitive
	b, _ := strconv.ParseBool("true")
	fmt.Println(b)

	f, _ := strconv.ParseFloat("3.14159", 64) // 64 bit float
	fmt.Println(f)

	i, _ := strconv.ParseInt("-45", 10, 64)
	fmt.Println(i)

	ui, _ := strconv.ParseUint("45", 10, 64)
	fmt.Println(ui)

	// Primitive to string
	s8 := strconv.FormatBool(true)
	fmt.Println(s8)

	s9 := strconv.FormatFloat(3.14159, 'E', -1, 64)
	fmt.Println(s9)

	s10 := strconv.FormatInt(-42, 10)
	fmt.Println(s10)

	s11 := strconv.FormatUint(42, 10)
	fmt.Println(s11)

	const s12 = "The 'quick' brown fox jumps over the *LAZY* dog"

	punctCount := 0
	lowerCount, upperCount := 0, 0
	spaceCount := 0
	hexdigitCount := 0

	for _, ch := range s12 {
		if unicode.IsPunct(ch) {
			punctCount++
		}

		if unicode.IsLower(ch) {
			lowerCount++
		}

		if unicode.IsUpper(ch) {
			upperCount++
		}

		if unicode.IsSpace(ch) {
			spaceCount++
		}

		if unicode.Is(unicode.Hex_Digit, ch) {
			hexdigitCount++
		}
	}

	fmt.Println("Punctuation:", punctCount)
	fmt.Println("Lowercase:", lowerCount)
	fmt.Println("Uppercase:", upperCount)
	fmt.Println("Space:", spaceCount)
	fmt.Println("Hexdigit:", hexdigitCount)
}
