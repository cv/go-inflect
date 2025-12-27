package inflect_test

import (
	"testing"

	inflect "github.com/cv/go-inflect"
)

func TestCurrencyToWords(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		currency string
		want     string
	}{
		// ===== USD - Basic cases =====
		{name: "USD basic", amount: 123.45, currency: "USD", want: "one hundred twenty-three dollars and forty-five cents"},
		{name: "USD one dollar", amount: 1.00, currency: "USD", want: "one dollar"},
		{name: "USD cents only", amount: 0.50, currency: "USD", want: "fifty cents"},
		{name: "USD one million", amount: 1000000.00, currency: "USD", want: "one million dollars"},
		{name: "USD zero", amount: 0.00, currency: "USD", want: "zero dollars"},
		{name: "USD one cent", amount: 0.01, currency: "USD", want: "one cent"},
		{name: "USD ninety-nine cents", amount: 0.99, currency: "USD", want: "ninety-nine cents"},

		// ===== USD - Singular/plural =====
		{name: "USD singular dollar", amount: 1.00, currency: "USD", want: "one dollar"},
		{name: "USD plural dollars", amount: 2.00, currency: "USD", want: "two dollars"},
		{name: "USD singular cent", amount: 0.01, currency: "USD", want: "one cent"},
		{name: "USD plural cents", amount: 0.02, currency: "USD", want: "two cents"},
		{name: "USD one dollar one cent", amount: 1.01, currency: "USD", want: "one dollar and one cent"},
		{name: "USD one dollar two cents", amount: 1.02, currency: "USD", want: "one dollar and two cents"},
		{name: "USD two dollars one cent", amount: 2.01, currency: "USD", want: "two dollars and one cent"},

		// ===== USD - Negative amounts =====
		{name: "USD negative", amount: -5.00, currency: "USD", want: "negative five dollars"},
		{name: "USD negative with cents", amount: -123.45, currency: "USD", want: "negative one hundred twenty-three dollars and forty-five cents"},
		{name: "USD negative one dollar", amount: -1.00, currency: "USD", want: "negative one dollar"},
		{name: "USD negative cents only", amount: -0.50, currency: "USD", want: "negative fifty cents"},

		// ===== USD - Large amounts =====
		{name: "USD thousand", amount: 1000.00, currency: "USD", want: "one thousand dollars"},
		{name: "USD ten thousand", amount: 10000.00, currency: "USD", want: "ten thousand dollars"},
		{name: "USD hundred thousand", amount: 100000.00, currency: "USD", want: "one hundred thousand dollars"},
		{name: "USD billion", amount: 1000000000.00, currency: "USD", want: "one billion dollars"},
		{name: "USD complex", amount: 1234567.89, currency: "USD", want: "one million two hundred thirty-four thousand five hundred sixty-seven dollars and eighty-nine cents"},

		// ===== USD - Rounding =====
		{name: "USD round down", amount: 1.234, currency: "USD", want: "one dollar and twenty-three cents"},
		{name: "USD round up", amount: 1.235, currency: "USD", want: "one dollar and twenty-four cents"},
		{name: "USD round half up", amount: 1.245, currency: "USD", want: "one dollar and twenty-five cents"},

		// ===== GBP =====
		{name: "GBP basic", amount: 123.45, currency: "GBP", want: "one hundred twenty-three pounds and forty-five pence"},
		{name: "GBP one pound", amount: 1.00, currency: "GBP", want: "one pound"},
		{name: "GBP plural pounds", amount: 5.00, currency: "GBP", want: "five pounds"},
		{name: "GBP one penny", amount: 0.01, currency: "GBP", want: "one penny"},
		{name: "GBP plural pence", amount: 0.50, currency: "GBP", want: "fifty pence"},
		{name: "GBP one pound one penny", amount: 1.01, currency: "GBP", want: "one pound and one penny"},
		{name: "GBP zero", amount: 0.00, currency: "GBP", want: "zero pounds"},

		// ===== EUR =====
		{name: "EUR basic", amount: 123.45, currency: "EUR", want: "one hundred twenty-three euros and forty-five cents"},
		{name: "EUR one euro", amount: 1.00, currency: "EUR", want: "one euro"},
		{name: "EUR plural euros", amount: 5.00, currency: "EUR", want: "five euros"},
		{name: "EUR cents only", amount: 0.50, currency: "EUR", want: "fifty cents"},
		{name: "EUR one cent", amount: 0.01, currency: "EUR", want: "one cent"},
		{name: "EUR zero", amount: 0.00, currency: "EUR", want: "zero euros"},

		// ===== CAD =====
		{name: "CAD basic", amount: 123.45, currency: "CAD", want: "one hundred twenty-three Canadian dollars and forty-five cents"},
		{name: "CAD one dollar", amount: 1.00, currency: "CAD", want: "one Canadian dollar"},
		{name: "CAD cents only", amount: 0.50, currency: "CAD", want: "fifty cents"},
		{name: "CAD zero", amount: 0.00, currency: "CAD", want: "zero Canadian dollars"},

		// ===== AUD =====
		{name: "AUD basic", amount: 123.45, currency: "AUD", want: "one hundred twenty-three Australian dollars and forty-five cents"},
		{name: "AUD one dollar", amount: 1.00, currency: "AUD", want: "one Australian dollar"},
		{name: "AUD cents only", amount: 0.50, currency: "AUD", want: "fifty cents"},
		{name: "AUD zero", amount: 0.00, currency: "AUD", want: "zero Australian dollars"},

		// ===== JPY (no minor unit) =====
		{name: "JPY basic", amount: 12345.00, currency: "JPY", want: "twelve thousand three hundred forty-five yen"},
		{name: "JPY one yen", amount: 1.00, currency: "JPY", want: "one yen"},
		{name: "JPY plural yen", amount: 5.00, currency: "JPY", want: "five yen"},
		{name: "JPY ignores decimals", amount: 123.45, currency: "JPY", want: "one hundred twenty-three yen"},
		{name: "JPY million", amount: 1000000.00, currency: "JPY", want: "one million yen"},
		{name: "JPY zero", amount: 0.00, currency: "JPY", want: "zero yen"},
		{name: "JPY negative", amount: -100.00, currency: "JPY", want: "negative one hundred yen"},

		// ===== CHF =====
		{name: "CHF basic", amount: 123.45, currency: "CHF", want: "one hundred twenty-three francs and forty-five centimes"},
		{name: "CHF one franc", amount: 1.00, currency: "CHF", want: "one franc"},
		{name: "CHF plural francs", amount: 5.00, currency: "CHF", want: "five francs"},
		{name: "CHF one centime", amount: 0.01, currency: "CHF", want: "one centime"},
		{name: "CHF plural centimes", amount: 0.50, currency: "CHF", want: "fifty centimes"},
		{name: "CHF zero", amount: 0.00, currency: "CHF", want: "zero francs"},

		// ===== CNY =====
		{name: "CNY basic", amount: 123.45, currency: "CNY", want: "one hundred twenty-three yuan and forty-five fen"},
		{name: "CNY one yuan", amount: 1.00, currency: "CNY", want: "one yuan"},
		{name: "CNY plural yuan", amount: 5.00, currency: "CNY", want: "five yuan"},
		{name: "CNY one fen", amount: 0.01, currency: "CNY", want: "one fen"},
		{name: "CNY plural fen", amount: 0.50, currency: "CNY", want: "fifty fen"},
		{name: "CNY zero", amount: 0.00, currency: "CNY", want: "zero yuan"},

		// ===== INR =====
		{name: "INR basic", amount: 123.45, currency: "INR", want: "one hundred twenty-three rupees and forty-five paise"},
		{name: "INR one rupee", amount: 1.00, currency: "INR", want: "one rupee"},
		{name: "INR plural rupees", amount: 5.00, currency: "INR", want: "five rupees"},
		{name: "INR one paisa", amount: 0.01, currency: "INR", want: "one paisa"},
		{name: "INR plural paise", amount: 0.50, currency: "INR", want: "fifty paise"},
		{name: "INR zero", amount: 0.00, currency: "INR", want: "zero rupees"},

		// ===== Unknown currency =====
		{name: "unknown currency", amount: 100.00, currency: "XYZ", want: ""},
		{name: "empty currency", amount: 100.00, currency: "", want: ""},

		// ===== Edge cases =====
		{name: "very small positive", amount: 0.001, currency: "USD", want: "zero dollars"},
		{name: "very small negative", amount: -0.001, currency: "USD", want: "zero dollars"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.CurrencyToWords(tt.amount, tt.currency)
			if got != tt.want {
				t.Errorf("CurrencyToWords(%v, %q) = %q, want %q", tt.amount, tt.currency, got, tt.want)
			}
		})
	}
}

// TestCurrencyToWordsExamples tests the specific examples from the issue.
func TestCurrencyToWordsExamples(t *testing.T) {
	examples := []struct {
		amount   float64
		currency string
		want     string
	}{
		{123.45, "USD", "one hundred twenty-three dollars and forty-five cents"},
		{1.00, "USD", "one dollar"},
		{0.50, "USD", "fifty cents"},
		{1000000.00, "USD", "one million dollars"},
		{123.45, "GBP", "one hundred twenty-three pounds and forty-five pence"},
	}

	for _, ex := range examples {
		t.Run("", func(t *testing.T) {
			got := inflect.CurrencyToWords(ex.amount, ex.currency)
			if got != ex.want {
				t.Errorf("CurrencyToWords(%v, %q) = %q, want %q", ex.amount, ex.currency, got, ex.want)
			}
		})
	}
}

func BenchmarkCurrencyToWords(b *testing.B) {
	benchmarks := []struct {
		name     string
		amount   float64
		currency string
	}{
		{"USD simple", 1.00, "USD"},
		{"USD with cents", 123.45, "USD"},
		{"USD large", 1234567.89, "USD"},
		{"GBP", 123.45, "GBP"},
		{"JPY no minor", 12345.00, "JPY"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for range b.N {
				inflect.CurrencyToWords(bm.amount, bm.currency)
			}
		})
	}
}
