package inflect_test

import (
	"fmt"

	inflect "github.com/cv/go-inflect"
)

func ExampleAn() {
	fmt.Println(inflect.An("apple"))
	fmt.Println(inflect.An("banana"))
	fmt.Println(inflect.An("hour"))
	fmt.Println(inflect.An("university"))
	// Output:
	// an apple
	// a banana
	// an hour
	// a university
}

func ExampleA() {
	fmt.Println(inflect.A("cat"))
	fmt.Println(inflect.A("elephant"))
	fmt.Println(inflect.A("honor"))
	fmt.Println(inflect.A("European"))
	// Output:
	// a cat
	// an elephant
	// an honor
	// a European
}

func ExamplePlural() {
	fmt.Println(inflect.Plural("cat"))
	fmt.Println(inflect.Plural("box"))
	fmt.Println(inflect.Plural("child"))
	fmt.Println(inflect.Plural("sheep"))
	fmt.Println(inflect.Plural("cactus"))
	// Output:
	// cats
	// boxes
	// children
	// sheep
	// cacti
}

func ExampleSingular() {
	fmt.Println(inflect.Singular("cats"))
	fmt.Println(inflect.Singular("boxes"))
	fmt.Println(inflect.Singular("children"))
	fmt.Println(inflect.Singular("sheep"))
	fmt.Println(inflect.Singular("cacti"))
	// Output:
	// cat
	// box
	// child
	// sheep
	// cactus
}

func ExampleNumberToWords() {
	fmt.Println(inflect.NumberToWords(1))
	fmt.Println(inflect.NumberToWords(42))
	fmt.Println(inflect.NumberToWords(100))
	fmt.Println(inflect.NumberToWords(1000))
	fmt.Println(inflect.NumberToWords(-5))
	// Output:
	// one
	// forty-two
	// one hundred
	// one thousand
	// negative five
}

func ExampleOrdinal() {
	fmt.Println(inflect.Ordinal(1))
	fmt.Println(inflect.Ordinal(2))
	fmt.Println(inflect.Ordinal(3))
	fmt.Println(inflect.Ordinal(11))
	fmt.Println(inflect.Ordinal(21))
	// Output:
	// 1st
	// 2nd
	// 3rd
	// 11th
	// 21st
}

func ExampleOrdinalWord() {
	fmt.Println(inflect.OrdinalWord(1))
	fmt.Println(inflect.OrdinalWord(2))
	fmt.Println(inflect.OrdinalWord(11))
	fmt.Println(inflect.OrdinalWord(21))
	fmt.Println(inflect.OrdinalWord(100))
	// Output:
	// first
	// second
	// eleventh
	// twenty-first
	// one hundredth
}

func ExampleJoin() {
	fmt.Println(inflect.Join([]string{}))
	fmt.Println(inflect.Join([]string{"apple"}))
	fmt.Println(inflect.Join([]string{"apple", "banana"}))
	fmt.Println(inflect.Join([]string{"apple", "banana", "cherry"}))
	// Output:
	//
	// apple
	// apple and banana
	// apple, banana, and cherry
}

func ExamplePresentParticiple() {
	fmt.Println(inflect.PresentParticiple("run"))
	fmt.Println(inflect.PresentParticiple("make"))
	fmt.Println(inflect.PresentParticiple("play"))
	fmt.Println(inflect.PresentParticiple("die"))
	fmt.Println(inflect.PresentParticiple("panic"))
	// Output:
	// running
	// making
	// playing
	// dying
	// panicking
}

func ExamplePastParticiple() {
	fmt.Println(inflect.PastParticiple("walk"))
	fmt.Println(inflect.PastParticiple("stop"))
	fmt.Println(inflect.PastParticiple("try"))
	fmt.Println(inflect.PastParticiple("go"))
	fmt.Println(inflect.PastParticiple("take"))
	// Output:
	// walked
	// stopped
	// tried
	// gone
	// taken
}

func ExampleCompare() {
	fmt.Println(inflect.Compare("cat", "cat"))
	fmt.Println(inflect.Compare("cat", "cats"))
	fmt.Println(inflect.Compare("cats", "cat"))
	fmt.Println(inflect.Compare("indexes", "indices"))
	fmt.Println(inflect.Compare("cat", "dog"))
	// Output:
	// eq
	// s:p
	// p:s
	// p:p
	//
}
