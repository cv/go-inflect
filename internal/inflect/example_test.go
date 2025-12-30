package inflect_test

import (
	"bytes"
	"fmt"
	"text/template"

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

// --- Adjective examples ---

func ExampleComparative() {
	fmt.Println(inflect.Comparative("big"))
	fmt.Println(inflect.Comparative("happy"))
	fmt.Println(inflect.Comparative("beautiful"))
	fmt.Println(inflect.Comparative("good"))
	// Output:
	// bigger
	// happier
	// more beautiful
	// better
}

func ExampleSuperlative() {
	fmt.Println(inflect.Superlative("big"))
	fmt.Println(inflect.Superlative("happy"))
	fmt.Println(inflect.Superlative("beautiful"))
	fmt.Println(inflect.Superlative("good"))
	// Output:
	// biggest
	// happiest
	// most beautiful
	// best
}

// --- Adverb examples ---

func ExampleAdverb() {
	fmt.Println(inflect.Adverb("quick"))
	fmt.Println(inflect.Adverb("happy"))
	fmt.Println(inflect.Adverb("gentle"))
	fmt.Println(inflect.Adverb("good"))
	fmt.Println(inflect.Adverb("fast"))
	// Output:
	// quickly
	// happily
	// gently
	// well
	// fast
}

// --- Article customization examples ---

func ExampleDefA() {
	inflect.DefAReset() // Reset to defaults first
	inflect.DefA("herb")
	fmt.Println(inflect.An("herb"))
	inflect.DefAReset() // Clean up
	// Output:
	// a herb
}

func ExampleDefAn() {
	inflect.DefAReset() // Reset to defaults first
	inflect.DefAn("hotel")
	fmt.Println(inflect.An("hotel"))
	inflect.DefAReset() // Clean up
	// Output:
	// an hotel
}

func ExampleDefAPattern() {
	inflect.DefAReset() // Reset to defaults first
	_ = inflect.DefAPattern("herb.*")
	fmt.Println(inflect.An("herbal"))
	inflect.DefAReset() // Clean up
	// Output:
	// a herbal
}

func ExampleDefAnPattern() {
	inflect.DefAReset() // Reset to defaults first
	_ = inflect.DefAnPattern("histor.*")
	fmt.Println(inflect.An("historical"))
	inflect.DefAReset() // Clean up
	// Output:
	// an historical
}

func ExampleDefAReset() {
	inflect.DefA("ape")
	inflect.DefAReset()
	fmt.Println(inflect.An("ape"))
	// Output:
	// an ape
}

func ExampleUndefA() {
	inflect.DefAReset()
	inflect.DefA("apple")
	fmt.Println(inflect.An("apple"))
	inflect.UndefA("apple")
	fmt.Println(inflect.An("apple"))
	inflect.DefAReset()
	// Output:
	// a apple
	// an apple
}

func ExampleUndefAn() {
	inflect.DefAReset()
	inflect.DefAn("hotel")
	fmt.Println(inflect.An("hotel"))
	inflect.UndefAn("hotel")
	fmt.Println(inflect.An("hotel"))
	inflect.DefAReset()
	// Output:
	// an hotel
	// a hotel
}

func ExampleUndefAPattern() {
	inflect.DefAReset()
	_ = inflect.DefAPattern("app.*")
	fmt.Println(inflect.An("apple"))
	inflect.UndefAPattern("app.*")
	fmt.Println(inflect.An("apple"))
	inflect.DefAReset()
	// Output:
	// a apple
	// an apple
}

func ExampleUndefAnPattern() {
	inflect.DefAReset()
	_ = inflect.DefAnPattern("histor.*")
	fmt.Println(inflect.An("historical"))
	inflect.UndefAnPattern("histor.*")
	fmt.Println(inflect.An("historical"))
	inflect.DefAReset()
	// Output:
	// an historical
	// a historical
}

// --- Classical pluralization examples ---

func ExampleClassical() {
	inflect.Classical(true)
	fmt.Println(inflect.Plural("formula"))
	inflect.Classical(false)
	fmt.Println(inflect.Plural("formula"))
	// Output:
	// formulae
	// formulas
}

func ExampleClassicalAll() {
	inflect.ClassicalAll(true)
	fmt.Println(inflect.Plural("formula"))
	inflect.ClassicalAll(false)
	fmt.Println(inflect.Plural("formula"))
	// Output:
	// formulae
	// formulas
}

func ExampleClassicalAncient() {
	inflect.ClassicalAncient(true)
	fmt.Println(inflect.Plural("antenna"))
	inflect.ClassicalAncient(false)
	fmt.Println(inflect.Plural("antenna"))
	// Output:
	// antennae
	// antennas
}

func ExampleClassicalHerd() {
	inflect.ClassicalHerd(true)
	fmt.Println(inflect.Plural("buffalo"))
	inflect.ClassicalHerd(false)
	fmt.Println(inflect.Plural("buffalo"))
	// Output:
	// buffalo
	// buffaloes
}

func ExampleClassicalNames() {
	inflect.ClassicalNames(true)
	fmt.Println(inflect.Plural("Jones"))
	inflect.ClassicalNames(false)
	fmt.Println(inflect.Plural("Jones"))
	// Output:
	// Jones
	// Joneses
}

func ExampleClassicalPersons() {
	inflect.ClassicalPersons(true)
	fmt.Println(inflect.Plural("person"))
	inflect.ClassicalPersons(false)
	fmt.Println(inflect.Plural("person"))
	// Output:
	// persons
	// people
}

func ExampleClassicalZero() {
	inflect.ClassicalZero(true)
	fmt.Println(inflect.No("cat", 0))
	inflect.ClassicalZero(false)
	fmt.Println(inflect.No("cat", 0))
	// Output:
	// no cat
	// no cats
}

func ExampleIsClassical() {
	inflect.Classical(false)
	fmt.Println(inflect.IsClassical())
	inflect.Classical(true)
	fmt.Println(inflect.IsClassical())
	inflect.Classical(false)
	// Output:
	// false
	// true
}

func ExampleIsClassicalAll() {
	inflect.ClassicalAll(false)
	fmt.Println(inflect.IsClassicalAll())
	inflect.ClassicalAll(true)
	fmt.Println(inflect.IsClassicalAll())
	inflect.ClassicalAll(false)
	// Output:
	// false
	// true
}

func ExampleIsClassicalAncient() {
	inflect.ClassicalAncient(false)
	fmt.Println(inflect.IsClassicalAncient())
	inflect.ClassicalAncient(true)
	fmt.Println(inflect.IsClassicalAncient())
	inflect.ClassicalAncient(false)
	// Output:
	// false
	// true
}

func ExampleIsClassicalHerd() {
	inflect.ClassicalHerd(false)
	fmt.Println(inflect.IsClassicalHerd())
	inflect.ClassicalHerd(true)
	fmt.Println(inflect.IsClassicalHerd())
	inflect.ClassicalHerd(false)
	// Output:
	// false
	// true
}

func ExampleIsClassicalNames() {
	inflect.ClassicalNames(false)
	fmt.Println(inflect.IsClassicalNames())
	inflect.ClassicalNames(true)
	fmt.Println(inflect.IsClassicalNames())
	inflect.ClassicalNames(false)
	// Output:
	// false
	// true
}

func ExampleIsClassicalPersons() {
	inflect.ClassicalPersons(false)
	fmt.Println(inflect.IsClassicalPersons())
	inflect.ClassicalPersons(true)
	fmt.Println(inflect.IsClassicalPersons())
	inflect.ClassicalPersons(false)
	// Output:
	// false
	// true
}

func ExampleIsClassicalZero() {
	inflect.ClassicalZero(false)
	fmt.Println(inflect.IsClassicalZero())
	inflect.ClassicalZero(true)
	fmt.Println(inflect.IsClassicalZero())
	inflect.ClassicalZero(false)
	// Output:
	// false
	// true
}

// --- Compare examples ---

func ExampleCompareAdjs() {
	fmt.Println(inflect.CompareAdjs("this", "these"))
	fmt.Println(inflect.CompareAdjs("these", "this"))
	fmt.Println(inflect.CompareAdjs("this", "this"))
	// Output:
	// s:p
	// p:s
	// eq
}

func ExampleCompareNouns() {
	fmt.Println(inflect.CompareNouns("cat", "cats"))
	fmt.Println(inflect.CompareNouns("mice", "mouse"))
	fmt.Println(inflect.CompareNouns("cat", "cat"))
	// Output:
	// s:p
	// p:s
	// eq
}

func ExampleCompareVerbs() {
	fmt.Println(inflect.CompareVerbs("runs", "run"))
	fmt.Println(inflect.CompareVerbs("run", "runs"))
	fmt.Println(inflect.CompareVerbs("is", "are"))
	// Output:
	// s:p
	// p:s
	// s:p
}

// --- Counting examples ---

func ExampleCountingWord() {
	fmt.Println(inflect.CountingWord(1))
	fmt.Println(inflect.CountingWord(2))
	fmt.Println(inflect.CountingWord(3))
	fmt.Println(inflect.CountingWord(4))
	fmt.Println(inflect.CountingWord(10))
	// Output:
	// once
	// twice
	// thrice
	// four times
	// ten times
}

func ExampleCountingWordThreshold() {
	fmt.Println(inflect.CountingWordThreshold(5, 10))
	fmt.Println(inflect.CountingWordThreshold(15, 10))
	fmt.Println(inflect.CountingWordThreshold(1, 10))
	// Output:
	// five times
	// 15 times
	// once
}

func ExampleCountingWordWithOptions() {
	fmt.Println(inflect.CountingWordWithOptions(3, true))
	fmt.Println(inflect.CountingWordWithOptions(3, false))
	// Output:
	// thrice
	// three times
}

// --- Currency examples ---

func ExampleCurrencyToWords() {
	fmt.Println(inflect.CurrencyToWords(1.00, "USD"))
	fmt.Println(inflect.CurrencyToWords(0.50, "USD"))
	fmt.Println(inflect.CurrencyToWords(123.45, "USD"))
	fmt.Println(inflect.CurrencyToWords(1.01, "GBP"))
	// Output:
	// one dollar
	// fifty cents
	// one hundred twenty-three dollars and forty-five cents
	// one pound and one penny
}

// --- Custom noun/verb/adj examples ---

func ExampleDefNoun() {
	inflect.DefNoun("foo", "foos")
	fmt.Println(inflect.Plural("foo"))
	fmt.Println(inflect.Singular("foos"))
	inflect.DefNounReset()
	// Output:
	// foos
	// foo
}

func ExampleDefNounReset() {
	inflect.DefNoun("child", "childs")
	fmt.Println(inflect.Plural("child"))
	inflect.DefNounReset()
	fmt.Println(inflect.Plural("child"))
	// Output:
	// childs
	// children
}

func ExampleUndefNoun() {
	inflect.DefNoun("widget", "widgetz")
	fmt.Println(inflect.Plural("widget"))
	inflect.UndefNoun("widget")
	fmt.Println(inflect.Plural("widget"))
	inflect.DefNounReset()
	// Output:
	// widgetz
	// widgets
}

func ExampleDefVerb() {
	inflect.DefVerb("florp", "florps")
	inflect.DefVerbReset()
	fmt.Println("verb defined and reset")
	// Output:
	// verb defined and reset
}

func ExampleDefVerbReset() {
	inflect.DefVerb("bloop", "bloops")
	inflect.DefVerbReset()
	fmt.Println("verbs reset")
	// Output:
	// verbs reset
}

func ExampleUndefVerb() {
	inflect.DefVerb("zorp", "zorps")
	fmt.Println(inflect.UndefVerb("zorp"))
	fmt.Println(inflect.UndefVerb("nonexistent"))
	inflect.DefVerbReset()
	// Output:
	// true
	// false
}

func ExampleDefAdj() {
	inflect.DefAdj("snazzy", "snazzies")
	inflect.DefAdjReset()
	fmt.Println("adj defined and reset")
	// Output:
	// adj defined and reset
}

func ExampleDefAdjReset() {
	inflect.DefAdj("groovy", "groovies")
	inflect.DefAdjReset()
	fmt.Println("adjs reset")
	// Output:
	// adjs reset
}

func ExampleUndefAdj() {
	inflect.DefAdj("nifty", "nifties")
	fmt.Println(inflect.UndefAdj("nifty"))
	fmt.Println(inflect.UndefAdj("nonexistent"))
	inflect.DefAdjReset()
	// Output:
	// true
	// false
}

// --- Fraction examples ---

func ExampleFractionToWords() {
	fmt.Println(inflect.FractionToWords(1, 2))
	fmt.Println(inflect.FractionToWords(3, 4))
	fmt.Println(inflect.FractionToWords(2, 3))
	fmt.Println(inflect.FractionToWords(5, 8))
	fmt.Println(inflect.FractionToWords(-1, 2))
	// Output:
	// one half
	// three quarters
	// two thirds
	// five eighths
	// negative one half
}

func ExampleFractionToWordsWithFourths() {
	fmt.Println(inflect.FractionToWordsWithFourths(1, 4))
	fmt.Println(inflect.FractionToWordsWithFourths(3, 4))
	// Output:
	// one fourth
	// three fourths
}

// --- Gender examples ---

func ExampleGender() {
	inflect.Gender("m")
	fmt.Println(inflect.GetGender())
	inflect.Gender("f")
	fmt.Println(inflect.GetGender())
	inflect.Gender("t") // Reset to default
	// Output:
	// m
	// f
}

func ExampleGetGender() {
	inflect.Gender("t")
	fmt.Println(inflect.GetGender())
	inflect.Gender("n")
	fmt.Println(inflect.GetGender())
	inflect.Gender("t") // Reset to default
	// Output:
	// t
	// n
}

func ExamplePluralAdj() {
	fmt.Println(inflect.PluralAdj("this"))
	fmt.Println(inflect.PluralAdj("that"))
	fmt.Println(inflect.PluralAdj("a"))
	fmt.Println(inflect.PluralAdj("my"))
	// Output:
	// these
	// those
	// some
	// our
}

func ExamplePluralNoun() {
	fmt.Println(inflect.PluralNoun("I"))
	fmt.Println(inflect.PluralNoun("me"))
	fmt.Println(inflect.PluralNoun("cat"))
	fmt.Println(inflect.PluralNoun("cat", 1))
	// Output:
	// We
	// us
	// cats
	// cat
}

func ExamplePluralVerb() {
	fmt.Println(inflect.PluralVerb("is"))
	fmt.Println(inflect.PluralVerb("was"))
	fmt.Println(inflect.PluralVerb("has"))
	fmt.Println(inflect.PluralVerb("runs"))
	// Output:
	// are
	// were
	// have
	// run
}

func ExampleSingularNoun() {
	inflect.Gender("t")
	fmt.Println(inflect.SingularNoun("we"))
	fmt.Println(inflect.SingularNoun("us"))
	fmt.Println(inflect.SingularNoun("cats"))
	fmt.Println(inflect.SingularNoun("cats", 2))
	// Output:
	// I
	// me
	// cat
	// cats
}

// --- Join examples ---

func ExampleJoinNoOxford() {
	fmt.Println(inflect.JoinNoOxford([]string{"a", "b"}))
	fmt.Println(inflect.JoinNoOxford([]string{"a", "b", "c"}))
	// Output:
	// a and b
	// a, b and c
}

func ExampleJoinNoOxfordWithConj() {
	fmt.Println(inflect.JoinNoOxfordWithConj([]string{"a", "b", "c"}, "or"))
	// Output:
	// a, b or c
}

func ExampleJoinWithAutoSep() {
	fmt.Println(inflect.JoinWithAutoSep([]string{"a", "b", "c"}, "and"))
	fmt.Println(inflect.JoinWithAutoSep([]string{"Jan 1, 2020", "Feb 2, 2021"}, "and"))
	// Output:
	// a, b, and c
	// Jan 1, 2020; and Feb 2, 2021
}

func ExampleJoinWithConj() {
	fmt.Println(inflect.JoinWithConj([]string{"a", "b"}, "or"))
	fmt.Println(inflect.JoinWithConj([]string{"a", "b", "c"}, "or"))
	// Output:
	// a or b
	// a, b, or c
}

func ExampleJoinWithFinalSep() {
	fmt.Println(inflect.JoinWithFinalSep([]string{"a", "b", "c"}, "and", ", ", "; "))
	fmt.Println(inflect.JoinWithFinalSep([]string{"a", "b", "c"}, "and", ", ", " "))
	// Output:
	// a, b; and c
	// a, b and c
}

func ExampleJoinWithSep() {
	fmt.Println(inflect.JoinWithSep([]string{"a", "b", "c"}, "and", "; "))
	// Output:
	// a; b; and c
}

// --- Number examples ---

func ExampleFormatNumber() {
	fmt.Println(inflect.FormatNumber(1000))
	fmt.Println(inflect.FormatNumber(1000000))
	fmt.Println(inflect.FormatNumber(123456789))
	fmt.Println(inflect.FormatNumber(-1234))
	// Output:
	// 1,000
	// 1,000,000
	// 123,456,789
	// -1,234
}

func ExampleGetNum() {
	inflect.Num(5)
	fmt.Println(inflect.GetNum())
	inflect.Num(0)
	fmt.Println(inflect.GetNum())
	// Output:
	// 5
	// 0
}

func ExampleNo() {
	inflect.ClassicalZero(false)
	fmt.Println(inflect.No("error", 0))
	fmt.Println(inflect.No("error", 1))
	fmt.Println(inflect.No("error", 5))
	// Output:
	// no errors
	// 1 error
	// 5 errors
}

func ExampleNum() {
	fmt.Println(inflect.Num(42))
	fmt.Println(inflect.Num())
	// Output:
	// 42
	// 0
}

func ExampleNumberToWordsFloat() {
	fmt.Println(inflect.NumberToWordsFloat(3.14))
	fmt.Println(inflect.NumberToWordsFloat(0.5))
	fmt.Println(inflect.NumberToWordsFloat(-2.7))
	// Output:
	// three point one four
	// zero point five
	// negative two point seven
}

func ExampleNumberToWordsFloatWithDecimal() {
	fmt.Println(inflect.NumberToWordsFloatWithDecimal(3.14, "point"))
	fmt.Println(inflect.NumberToWordsFloatWithDecimal(3.14, "dot"))
	// Output:
	// three point one four
	// three dot one four
}

func ExampleNumberToWordsGrouped() {
	fmt.Println(inflect.NumberToWordsGrouped(1234, 2))
	fmt.Println(inflect.NumberToWordsGrouped(123456, 2))
	// Output:
	// twelve thirty-four
	// twelve thirty-four fifty-six
}

func ExampleNumberToWordsThreshold() {
	fmt.Println(inflect.NumberToWordsThreshold(5, 10))
	fmt.Println(inflect.NumberToWordsThreshold(15, 10))
	// Output:
	// five
	// 15
}

func ExampleNumberToWordsWithAnd() {
	fmt.Println(inflect.NumberToWordsWithAnd(101))
	fmt.Println(inflect.NumberToWordsWithAnd(1001))
	fmt.Println(inflect.NumberToWordsWithAnd(121))
	// Output:
	// one hundred and one
	// one thousand and one
	// one hundred and twenty-one
}

// --- Ordinal examples ---

func ExampleIsOrdinal() {
	fmt.Println(inflect.IsOrdinal("1st"))
	fmt.Println(inflect.IsOrdinal("first"))
	fmt.Println(inflect.IsOrdinal("twenty-first"))
	fmt.Println(inflect.IsOrdinal("one"))
	fmt.Println(inflect.IsOrdinal("cat"))
	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleOrdinalSuffix() {
	fmt.Println(inflect.OrdinalSuffix(1))
	fmt.Println(inflect.OrdinalSuffix(2))
	fmt.Println(inflect.OrdinalSuffix(3))
	fmt.Println(inflect.OrdinalSuffix(4))
	fmt.Println(inflect.OrdinalSuffix(11))
	// Output:
	// st
	// nd
	// rd
	// th
	// th
}

func ExampleOrdinalToCardinal() {
	fmt.Println(inflect.OrdinalToCardinal("1st"))
	fmt.Println(inflect.OrdinalToCardinal("first"))
	fmt.Println(inflect.OrdinalToCardinal("twenty-first"))
	fmt.Println(inflect.OrdinalToCardinal("one"))
	// Output:
	// 1
	// one
	// twenty-one
	// one
}

func ExampleWordToOrdinal() {
	fmt.Println(inflect.WordToOrdinal("1"))
	fmt.Println(inflect.WordToOrdinal("one"))
	fmt.Println(inflect.WordToOrdinal("twenty-one"))
	fmt.Println(inflect.WordToOrdinal("One"))
	// Output:
	// 1st
	// first
	// twenty-first
	// First
}

// --- Participle examples ---

func ExampleIsParticiple() {
	fmt.Println(inflect.IsParticiple("running"))
	fmt.Println(inflect.IsParticiple("walked"))
	fmt.Println(inflect.IsParticiple("taken"))
	fmt.Println(inflect.IsParticiple("walk"))
	// Output:
	// true
	// true
	// true
	// false
}

// --- Past tense examples ---

func ExamplePastTense() {
	fmt.Println(inflect.PastTense("walk"))
	fmt.Println(inflect.PastTense("go"))
	fmt.Println(inflect.PastTense("try"))
	fmt.Println(inflect.PastTense("stop"))
	// Output:
	// walked
	// went
	// tried
	// stopped
}

func ExampleFutureTense() {
	fmt.Println(inflect.FutureTense("walk"))
	fmt.Println(inflect.FutureTense("go"))
	fmt.Println(inflect.FutureTense("be"))
	fmt.Println(inflect.FutureTense("Walk"))
	// Output:
	// will walk
	// will go
	// will be
	// Will Walk
}

// --- Possessive examples ---

func ExampleGetPossessiveStyle() {
	inflect.PossessiveStyle(inflect.PossessiveModern)
	fmt.Println(inflect.GetPossessiveStyle() == inflect.PossessiveModern)
	inflect.PossessiveStyle(inflect.PossessiveTraditional)
	fmt.Println(inflect.GetPossessiveStyle() == inflect.PossessiveTraditional)
	inflect.PossessiveStyle(inflect.PossessiveModern)
	// Output:
	// true
	// true
}

func ExamplePossessive() {
	inflect.PossessiveStyle(inflect.PossessiveModern)
	fmt.Println(inflect.Possessive("cat"))
	fmt.Println(inflect.Possessive("cats"))
	fmt.Println(inflect.Possessive("children"))
	fmt.Println(inflect.Possessive("James"))
	// Output:
	// cat's
	// cats'
	// children's
	// James's
}

func ExamplePossessiveStyle() {
	inflect.PossessiveStyle(inflect.PossessiveModern)
	fmt.Println(inflect.Possessive("James"))
	inflect.PossessiveStyle(inflect.PossessiveTraditional)
	fmt.Println(inflect.Possessive("James"))
	inflect.PossessiveStyle(inflect.PossessiveModern)
	// Output:
	// James's
	// James'
}

// --- Utility examples ---

func ExampleCapitalize() {
	fmt.Println(inflect.Capitalize("hello"))
	fmt.Println(inflect.Capitalize("HELLO"))
	fmt.Println(inflect.Capitalize("hello world"))
	// Output:
	// Hello
	// HELLO
	// Hello world
}

func ExampleIsPlural() {
	fmt.Println(inflect.IsPlural("cats"))
	fmt.Println(inflect.IsPlural("cat"))
	fmt.Println(inflect.IsPlural("children"))
	fmt.Println(inflect.IsPlural("child"))
	// Output:
	// true
	// false
	// true
	// false
}

func ExampleIsSingular() {
	fmt.Println(inflect.IsSingular("cat"))
	fmt.Println(inflect.IsSingular("cats"))
	fmt.Println(inflect.IsSingular("child"))
	fmt.Println(inflect.IsSingular("children"))
	// Output:
	// true
	// false
	// true
	// false
}

func ExampleTitleize() {
	fmt.Println(inflect.Titleize("hello world"))
	fmt.Println(inflect.Titleize("HELLO WORLD"))
	fmt.Println(inflect.Titleize("hello-world"))
	// Output:
	// Hello World
	// Hello World
	// Hello-World
}

func ExampleWordCount() {
	fmt.Println(inflect.WordCount("hello world"))
	fmt.Println(inflect.WordCount("  one   two   three  "))
	fmt.Println(inflect.WordCount(""))
	fmt.Println(inflect.WordCount("single"))
	// Output:
	// 2
	// 3
	// 0
	// 1
}

// --- Case conversion examples ---

func ExampleDasherize() {
	fmt.Println(inflect.Dasherize("HelloWorld"))
	fmt.Println(inflect.Dasherize("hello_world"))
	fmt.Println(inflect.Dasherize("HTTPServer"))
	fmt.Println(inflect.Dasherize("getHTTPResponse"))
	// Output:
	// hello-world
	// hello-world
	// http-server
	// get-http-response
}

func ExampleKebabCase() {
	fmt.Println(inflect.KebabCase("HelloWorld"))
	fmt.Println(inflect.KebabCase("hello_world"))
	fmt.Println(inflect.KebabCase("XMLParser"))
	// Output:
	// hello-world
	// hello-world
	// xml-parser
}

func ExampleUnderscore() {
	fmt.Println(inflect.Underscore("HelloWorld"))
	fmt.Println(inflect.Underscore("hello-world"))
	fmt.Println(inflect.Underscore("HTTPServer"))
	fmt.Println(inflect.Underscore("getHTTPResponse"))
	// Output:
	// hello_world
	// hello_world
	// http_server
	// get_http_response
}

func ExampleSnakeCase() {
	fmt.Println(inflect.SnakeCase("HelloWorld"))
	fmt.Println(inflect.SnakeCase("hello-world"))
	fmt.Println(inflect.SnakeCase("XMLParser"))
	// Output:
	// hello_world
	// hello_world
	// xml_parser
}

func ExamplePascalCase() {
	fmt.Println(inflect.PascalCase("hello_world"))
	fmt.Println(inflect.PascalCase("hello-world"))
	fmt.Println(inflect.PascalCase("HTTP_SERVER"))
	fmt.Println(inflect.PascalCase("helloWorld"))
	// Output:
	// HelloWorld
	// HelloWorld
	// HttpServer
	// HelloWorld
}

func ExampleTitleCase() {
	fmt.Println(inflect.TitleCase("hello_world"))
	fmt.Println(inflect.TitleCase("hello-world"))
	fmt.Println(inflect.TitleCase("get_http_response"))
	// Output:
	// HelloWorld
	// HelloWorld
	// GetHttpResponse
}

func ExampleCamelCase() {
	fmt.Println(inflect.CamelCase("hello_world"))
	fmt.Println(inflect.CamelCase("hello-world"))
	fmt.Println(inflect.CamelCase("HTTP_SERVER"))
	fmt.Println(inflect.CamelCase("HelloWorld"))
	// Output:
	// helloWorld
	// helloWorld
	// httpServer
	// helloWorld
}

// Compatibility aliases for github.com/jinzhu/inflection

func ExampleAddIrregular() {
	inflect.AddIrregular("octopus", "octopodes")
	fmt.Println(inflect.Plural("octopus"))
	fmt.Println(inflect.Singular("octopodes"))
	// Output:
	// octopodes
	// octopus
}

func ExampleAddUncountable() {
	inflect.AddUncountable("rice", "wheat")
	fmt.Println(inflect.Plural("rice"))
	fmt.Println(inflect.Plural("wheat"))
	// Output:
	// rice
	// wheat
}

// Compatibility aliases for github.com/go-openapi/inflect

func ExamplePluralize() {
	fmt.Println(inflect.Pluralize("cat"))
	fmt.Println(inflect.Pluralize("child"))
	fmt.Println(inflect.Pluralize("Person"))
	// Output:
	// cats
	// children
	// People
}

func ExampleSingularize() {
	fmt.Println(inflect.Singularize("cats"))
	fmt.Println(inflect.Singularize("children"))
	fmt.Println(inflect.Singularize("People"))
	// Output:
	// cat
	// child
	// Person
}

func ExampleCamelize() {
	fmt.Println(inflect.Camelize("hello_world"))
	fmt.Println(inflect.Camelize("foo-bar"))
	fmt.Println(inflect.Camelize("get_http_response"))
	// Output:
	// HelloWorld
	// FooBar
	// GetHttpResponse
}

func ExampleCamelizeDownFirst() {
	fmt.Println(inflect.CamelizeDownFirst("hello_world"))
	fmt.Println(inflect.CamelizeDownFirst("foo-bar"))
	fmt.Println(inflect.CamelizeDownFirst("get_http_response"))
	// Output:
	// helloWorld
	// fooBar
	// getHttpResponse
}

// Rails-style helper functions

func ExampleHumanize() {
	fmt.Println(inflect.Humanize("employee_salary"))
	fmt.Println(inflect.Humanize("author_id"))
	fmt.Println(inflect.Humanize("XMLParser"))
	// Output:
	// Employee salary
	// Author
	// Xml parser
}

func ExampleForeignKey() {
	fmt.Println(inflect.ForeignKey("Person"))
	fmt.Println(inflect.ForeignKey("AdminUser"))
	fmt.Println(inflect.ForeignKey("XMLParser"))
	// Output:
	// person_id
	// admin_user_id
	// xml_parser_id
}

func ExampleForeignKeyCondensed() {
	fmt.Println(inflect.ForeignKeyCondensed("Person"))
	fmt.Println(inflect.ForeignKeyCondensed("AdminUser"))
	// Output:
	// personid
	// admin_userid
}

func ExampleTableize() {
	fmt.Println(inflect.Tableize("Person"))
	fmt.Println(inflect.Tableize("RawScaledScorer"))
	fmt.Println(inflect.Tableize("Child"))
	// Output:
	// people
	// raw_scaled_scorers
	// children
}

func ExampleParameterize() {
	fmt.Println(inflect.Parameterize("Hello World!"))
	fmt.Println(inflect.Parameterize("café au lait"))
	fmt.Println(inflect.Parameterize("  Multiple   Spaces  "))
	// Output:
	// hello-world
	// cafe-au-lait
	// multiple-spaces
}

func ExampleParameterizeJoin() {
	fmt.Println(inflect.ParameterizeJoin("Hello World!", "_"))
	fmt.Println(inflect.ParameterizeJoin("Hello World!", "-"))
	fmt.Println(inflect.ParameterizeJoin("Hello World!", ""))
	// Output:
	// hello_world
	// hello-world
	// helloworld
}

func ExampleTypeify() {
	fmt.Println(inflect.Typeify("users"))
	fmt.Println(inflect.Typeify("raw_scaled_scorers"))
	fmt.Println(inflect.Typeify("people"))
	// Output:
	// User
	// RawScaledScorer
	// Person
}

func ExampleAsciify() {
	fmt.Println(inflect.Asciify("café"))
	fmt.Println(inflect.Asciify("naïve"))
	fmt.Println(inflect.Asciify("Crème brûlée"))
	// Output:
	// cafe
	// naive
	// Creme brulee
}

// --- Engine examples ---

// ExampleNewEngine demonstrates creating isolated engine instances.
// Each engine has independent configuration, so changes to one
// don't affect others or the package-level default.
func ExampleNewEngine() {
	// Create an engine with classical pluralization
	classical := inflect.NewEngine()
	classical.Classical(true)

	// Create an engine with default (modern) settings
	modern := inflect.NewEngine()

	// Each engine operates independently
	fmt.Println("classical:", classical.Plural("formula"))
	fmt.Println("modern:", modern.Plural("formula"))

	// Output:
	// classical: formulae
	// modern: formulas
}

// ExampleEngine_Clone demonstrates creating a copy of an engine's configuration.
func ExampleEngine_Clone() {
	// Create and configure an engine
	original := inflect.NewEngine()
	original.DefNoun("gizmo", "gizmoz")

	// Clone it
	cloned := original.Clone()

	// Both have the custom noun
	fmt.Println("original:", original.Plural("gizmo"))
	fmt.Println("cloned:", cloned.Plural("gizmo"))

	// Modify the clone - original is unaffected
	cloned.DefNoun("widget", "widgetz")
	fmt.Println("cloned widget:", cloned.Plural("widget"))
	fmt.Println("original widget:", original.Plural("widget"))

	// Output:
	// original: gizmoz
	// cloned: gizmoz
	// cloned widget: widgetz
	// original widget: widgets
}

func ExampleFuncMap() {
	tmpl := template.New("example").Funcs(inflect.FuncMap())
	_, _ = tmpl.Parse(`I have {{plural "cat" .Count}} and {{an "apple"}}`)

	var buf bytes.Buffer
	_ = tmpl.Execute(&buf, map[string]int{"Count": 3})
	fmt.Println(buf.String())
	// Output:
	// I have cats and an apple
}

func ExampleEngine_FuncMap() {
	e := inflect.NewEngine()
	e.DefNoun("gremlin", "gremloz")

	tmpl := template.New("example").Funcs(e.FuncMap())
	_, _ = tmpl.Parse(`The plural of gremlin is {{plural "gremlin"}}`)

	var buf bytes.Buffer
	_ = tmpl.Execute(&buf, nil)
	fmt.Println(buf.String())
	// Output:
	// The plural of gremlin is gremloz
}
