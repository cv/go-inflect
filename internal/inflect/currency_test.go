package inflect_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

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

		// ===== Special float values =====
		{name: "positive infinity", amount: math.Inf(1), currency: "USD", want: ""},
		{name: "negative infinity", amount: math.Inf(-1), currency: "USD", want: ""},
		{name: "NaN", amount: math.NaN(), currency: "USD", want: ""},

		// ===== Case insensitive currency codes =====
		{name: "lowercase usd", amount: 100.00, currency: "usd", want: "one hundred dollars"},
		{name: "mixed case Usd", amount: 100.00, currency: "Usd", want: "one hundred dollars"},
		{name: "lowercase gbp", amount: 1.01, currency: "gbp", want: "one pound and one penny"},

		// ===== Additional currencies =====
		{name: "NZD basic", amount: 123.45, currency: "NZD", want: "one hundred twenty-three New Zealand dollars and forty-five cents"},
		{name: "NZD one dollar", amount: 1.00, currency: "NZD", want: "one New Zealand dollar"},
		{name: "MXN basic", amount: 100.50, currency: "MXN", want: "one hundred pesos and fifty centavos"},
		{name: "MXN one peso", amount: 1.00, currency: "MXN", want: "one peso"},
		{name: "MXN one centavo", amount: 0.01, currency: "MXN", want: "one centavo"},
		{name: "BRL basic", amount: 100.50, currency: "BRL", want: "one hundred reais and fifty centavos"},
		{name: "BRL one real", amount: 1.00, currency: "BRL", want: "one real"},
		{name: "SGD basic", amount: 100.50, currency: "SGD", want: "one hundred Singapore dollars and fifty cents"},
		{name: "HKD basic", amount: 100.50, currency: "HKD", want: "one hundred Hong Kong dollars and fifty cents"},
		{name: "KRW basic", amount: 1000.00, currency: "KRW", want: "one thousand won"},
		{name: "KRW one won", amount: 1.00, currency: "KRW", want: "one won"},
		{name: "SEK basic", amount: 100.50, currency: "SEK", want: "one hundred kronor and fifty öre"},
		{name: "SEK one krona", amount: 1.00, currency: "SEK", want: "one krona"},
		{name: "NOK basic", amount: 100.50, currency: "NOK", want: "one hundred kroner and fifty øre"},
		{name: "NOK one krone", amount: 1.00, currency: "NOK", want: "one krone"},
		{name: "DKK basic", amount: 100.50, currency: "DKK", want: "one hundred kroner and fifty øre"},
		{name: "DKK one krone", amount: 1.00, currency: "DKK", want: "one krone"},
		{name: "ZAR basic", amount: 100.50, currency: "ZAR", want: "one hundred rand and fifty cents"},
		{name: "ZAR one rand", amount: 1.00, currency: "ZAR", want: "one rand"},
		{name: "PLN basic", amount: 100.50, currency: "PLN", want: "one hundred zlotys and fifty groszy"},
		{name: "PLN one zloty", amount: 1.00, currency: "PLN", want: "one zloty"},
		{name: "PLN one grosz", amount: 0.01, currency: "PLN", want: "one grosz"},
		{name: "THB basic", amount: 100.50, currency: "THB", want: "one hundred baht and fifty satang"},
		{name: "THB one baht", amount: 1.00, currency: "THB", want: "one baht"},
		{name: "ILS basic", amount: 100.50, currency: "ILS", want: "one hundred shekels and fifty agorot"},
		{name: "ILS one shekel", amount: 1.00, currency: "ILS", want: "one shekel"},
		{name: "ILS one agora", amount: 0.01, currency: "ILS", want: "one agora"},
		{name: "AED basic", amount: 100.50, currency: "AED", want: "one hundred dirhams and fifty fils"},
		{name: "AED one dirham", amount: 1.00, currency: "AED", want: "one dirham"},
		{name: "SAR basic", amount: 100.50, currency: "SAR", want: "one hundred riyals and fifty halalas"},
		{name: "SAR one riyal", amount: 1.00, currency: "SAR", want: "one riyal"},
		{name: "TRY basic", amount: 100.50, currency: "TRY", want: "one hundred liras and fifty kurus"},
		{name: "TRY one lira", amount: 1.00, currency: "TRY", want: "one lira"},
		{name: "RUB basic", amount: 100.50, currency: "RUB", want: "one hundred rubles and fifty kopecks"},
		{name: "RUB one ruble", amount: 1.00, currency: "RUB", want: "one ruble"},
		{name: "RUB one kopeck", amount: 0.01, currency: "RUB", want: "one kopeck"},

		// ===== Additional top 50 currencies =====
		{name: "TWD basic", amount: 100.50, currency: "TWD", want: "one hundred New Taiwan dollars and fifty cents"},
		{name: "TWD one dollar", amount: 1.00, currency: "TWD", want: "one New Taiwan dollar"},
		{name: "IDR basic", amount: 10000.00, currency: "IDR", want: "ten thousand rupiah"},
		{name: "IDR one rupiah", amount: 1.00, currency: "IDR", want: "one rupiah"},
		{name: "CZK basic", amount: 100.50, currency: "CZK", want: "one hundred koruny and fifty haléřů"},
		{name: "CZK one koruna", amount: 1.00, currency: "CZK", want: "one koruna"},
		{name: "HUF basic", amount: 1000.00, currency: "HUF", want: "one thousand forints"},
		{name: "HUF one forint", amount: 1.00, currency: "HUF", want: "one forint"},
		{name: "CLP basic", amount: 1000.00, currency: "CLP", want: "one thousand pesos"},
		{name: "CLP one peso", amount: 1.00, currency: "CLP", want: "one peso"},
		{name: "PHP basic", amount: 100.50, currency: "PHP", want: "one hundred pesos and fifty centavos"},
		{name: "PHP one peso", amount: 1.00, currency: "PHP", want: "one peso"},
		{name: "MYR basic", amount: 100.50, currency: "MYR", want: "one hundred ringgit and fifty sen"},
		{name: "MYR one ringgit", amount: 1.00, currency: "MYR", want: "one ringgit"},
		{name: "COP basic", amount: 10000.00, currency: "COP", want: "ten thousand pesos"},
		{name: "COP one peso", amount: 1.00, currency: "COP", want: "one peso"},
		{name: "RON basic", amount: 100.50, currency: "RON", want: "one hundred lei and fifty bani"},
		{name: "RON one leu", amount: 1.00, currency: "RON", want: "one leu"},
		{name: "PEN basic", amount: 100.50, currency: "PEN", want: "one hundred soles and fifty céntimos"},
		{name: "PEN one sol", amount: 1.00, currency: "PEN", want: "one sol"},
		{name: "EGP basic", amount: 100.50, currency: "EGP", want: "one hundred pounds and fifty piastres"},
		{name: "EGP one pound", amount: 1.00, currency: "EGP", want: "one pound"},
		{name: "VND basic", amount: 10000.00, currency: "VND", want: "ten thousand dong"},
		{name: "VND one dong", amount: 1.00, currency: "VND", want: "one dong"},
		{name: "PKR basic", amount: 100.50, currency: "PKR", want: "one hundred rupees and fifty paise"},
		{name: "PKR one rupee", amount: 1.00, currency: "PKR", want: "one rupee"},
		{name: "NGN basic", amount: 100.50, currency: "NGN", want: "one hundred naira and fifty kobo"},
		{name: "NGN one naira", amount: 1.00, currency: "NGN", want: "one naira"},
		{name: "BDT basic", amount: 100.50, currency: "BDT", want: "one hundred taka and fifty poisha"},
		{name: "BDT one taka", amount: 1.00, currency: "BDT", want: "one taka"},
		{name: "ARS basic", amount: 100.50, currency: "ARS", want: "one hundred pesos and fifty centavos"},
		{name: "ARS one peso", amount: 1.00, currency: "ARS", want: "one peso"},
		{name: "QAR basic", amount: 100.50, currency: "QAR", want: "one hundred riyals and fifty dirhams"},
		{name: "QAR one riyal", amount: 1.00, currency: "QAR", want: "one riyal"},
		{name: "KWD basic", amount: 100.50, currency: "KWD", want: "one hundred dinars and fifty fils"},
		{name: "KWD one dinar", amount: 1.00, currency: "KWD", want: "one dinar"},
		{name: "MAD basic", amount: 100.50, currency: "MAD", want: "one hundred dirhams and fifty centimes"},
		{name: "MAD one dirham", amount: 1.00, currency: "MAD", want: "one dirham"},
		{name: "UAH basic", amount: 100.50, currency: "UAH", want: "one hundred hryvnias and fifty kopiykas"},
		{name: "UAH one hryvnia", amount: 1.00, currency: "UAH", want: "one hryvnia"},
		{name: "KES basic", amount: 100.50, currency: "KES", want: "one hundred shillings and fifty cents"},
		{name: "KES one shilling", amount: 1.00, currency: "KES", want: "one shilling"},
		{name: "GHS basic", amount: 100.50, currency: "GHS", want: "one hundred cedis and fifty pesewas"},
		{name: "GHS one cedi", amount: 1.00, currency: "GHS", want: "one cedi"},
		{name: "CRC basic", amount: 1000.50, currency: "CRC", want: "one thousand colones and fifty céntimos"},
		{name: "CRC one colon", amount: 1.00, currency: "CRC", want: "one colón"},
		{name: "LKR basic", amount: 100.50, currency: "LKR", want: "one hundred rupees and fifty cents"},
		{name: "LKR one rupee", amount: 1.00, currency: "LKR", want: "one rupee"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inflect.CurrencyToWords(tt.amount, tt.currency)
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, ex.want, got)
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
