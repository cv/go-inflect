package inflect

import (
	"math"
	"strings"
)

// currencyInfo holds the singular and plural forms for major and minor currency units.
type currencyInfo struct {
	majorSingular string
	majorPlural   string
	minorSingular string
	minorPlural   string
	hasMinorUnit  bool
}

// currencies maps currency codes to their unit names.
// Covers the top 50 most traded currencies by volume.
var currencies = map[string]currencyInfo{
	// Major world currencies
	"USD": {"dollar", "dollars", "cent", "cents", true},
	"EUR": {"euro", "euros", "cent", "cents", true},
	"JPY": {"yen", "yen", "", "", false},
	"GBP": {"pound", "pounds", "penny", "pence", true},
	"CHF": {"franc", "francs", "centime", "centimes", true},
	"CNY": {"yuan", "yuan", "fen", "fen", true},
	// Dollar variants
	"AUD": {"Australian dollar", "Australian dollars", "cent", "cents", true},
	"CAD": {"Canadian dollar", "Canadian dollars", "cent", "cents", true},
	"HKD": {"Hong Kong dollar", "Hong Kong dollars", "cent", "cents", true},
	"SGD": {"Singapore dollar", "Singapore dollars", "cent", "cents", true},
	"NZD": {"New Zealand dollar", "New Zealand dollars", "cent", "cents", true},
	"TWD": {"New Taiwan dollar", "New Taiwan dollars", "cent", "cents", true},
	// Scandinavian currencies
	"SEK": {"krona", "kronor", "öre", "öre", true},
	"NOK": {"krone", "kroner", "øre", "øre", true},
	"DKK": {"krone", "kroner", "øre", "øre", true},
	// Asian currencies
	"KRW": {"won", "won", "", "", false},
	"INR": {"rupee", "rupees", "paisa", "paise", true},
	"THB": {"baht", "baht", "satang", "satang", true},
	"IDR": {"rupiah", "rupiah", "sen", "sen", true},
	"PHP": {"peso", "pesos", "centavo", "centavos", true},
	"MYR": {"ringgit", "ringgit", "sen", "sen", true},
	"VND": {"dong", "dong", "", "", false},
	"PKR": {"rupee", "rupees", "paisa", "paise", true},
	"BDT": {"taka", "taka", "poisha", "poisha", true},
	"LKR": {"rupee", "rupees", "cent", "cents", true},
	// Latin American currencies
	"MXN": {"peso", "pesos", "centavo", "centavos", true},
	"BRL": {"real", "reais", "centavo", "centavos", true},
	"CLP": {"peso", "pesos", "centavo", "centavos", true},
	"COP": {"peso", "pesos", "centavo", "centavos", true},
	"PEN": {"sol", "soles", "céntimo", "céntimos", true},
	"ARS": {"peso", "pesos", "centavo", "centavos", true},
	"CRC": {"colón", "colones", "céntimo", "céntimos", true},
	// European currencies (non-EUR)
	"PLN": {"zloty", "zlotys", "grosz", "groszy", true},
	"CZK": {"koruna", "koruny", "haléř", "haléřů", true},
	"HUF": {"forint", "forints", "fillér", "fillérs", true},
	"RON": {"leu", "lei", "ban", "bani", true},
	"UAH": {"hryvnia", "hryvnias", "kopiyka", "kopiykas", true},
	"RUB": {"ruble", "rubles", "kopeck", "kopecks", true},
	// Middle Eastern currencies
	"ILS": {"shekel", "shekels", "agora", "agorot", true},
	"AED": {"dirham", "dirhams", "fil", "fils", true},
	"SAR": {"riyal", "riyals", "halala", "halalas", true},
	"TRY": {"lira", "liras", "kurus", "kurus", true},
	"QAR": {"riyal", "riyals", "dirham", "dirhams", true},
	"KWD": {"dinar", "dinars", "fil", "fils", true},
	"EGP": {"pound", "pounds", "piastre", "piastres", true},
	"MAD": {"dirham", "dirhams", "centime", "centimes", true},
	// African currencies
	"ZAR": {"rand", "rand", "cent", "cents", true},
	"NGN": {"naira", "naira", "kobo", "kobo", true},
	"KES": {"shilling", "shillings", "cent", "cents", true},
	"GHS": {"cedi", "cedis", "pesewa", "pesewas", true},
}

// CurrencyToWords converts a currency amount to its English word representation.
//
// The function handles various cases:
//   - Proper singular/plural forms based on amount
//   - Zero minor units are omitted (e.g., "one hundred dollars" not "one hundred dollars and zero cents")
//   - Minor units only when major is zero (e.g., "fifty cents")
//   - Negative amounts are prefixed with "negative"
//   - Amounts are rounded to 2 decimal places
//
// Supported currency codes: USD, GBP, EUR, CAD, AUD, JPY, CHF, CNY, INR.
// Returns empty string for unknown currency codes.
//
// Examples:
//   - CurrencyToWords(123.45, "USD") returns "one hundred twenty-three dollars and forty-five cents"
//   - CurrencyToWords(1.00, "USD") returns "one dollar"
//   - CurrencyToWords(0.50, "USD") returns "fifty cents"
//   - CurrencyToWords(1000000.00, "USD") returns "one million dollars"
//   - CurrencyToWords(123.45, "GBP") returns "one hundred twenty-three pounds and forty-five pence"
//   - CurrencyToWords(-5.00, "USD") returns "negative five dollars"
func CurrencyToWords(amount float64, currency string) string {
	// Handle special float values
	if math.IsNaN(amount) || math.IsInf(amount, 0) {
		return ""
	}

	// Normalize currency code to uppercase for lookup
	info, ok := currencies[strings.ToUpper(currency)]
	if !ok {
		return ""
	}

	// Handle negative amounts
	negative := false
	if amount < 0 {
		negative = true
		amount = -amount
	}

	// Round to 2 decimal places
	amount = math.Round(amount*100) / 100

	// If amount rounds to zero, it's not negative
	if amount == 0 {
		negative = false
	}

	// Split into major and minor units
	major := int(amount)
	minor := int(math.Round((amount - float64(major)) * 100))

	// Build the result
	var result string

	switch {
	case major == 0 && minor == 0:
		// Zero amount
		result = "zero " + info.majorPlural
	case major == 0:
		// Minor units only
		result = formatMinorUnit(minor, info)
	case minor == 0 || !info.hasMinorUnit:
		// Major units only (or currency has no minor unit)
		result = formatMajorUnit(major, info)
	default:
		// Both major and minor units
		result = formatMajorUnit(major, info) + " and " + formatMinorUnit(minor, info)
	}

	if negative {
		return "negative " + result
	}
	return result
}

// formatMajorUnit formats the major currency unit with proper singular/plural.
func formatMajorUnit(amount int, info currencyInfo) string {
	word := NumberToWords(amount)
	if amount == 1 {
		return word + " " + info.majorSingular
	}
	return word + " " + info.majorPlural
}

// formatMinorUnit formats the minor currency unit with proper singular/plural.
func formatMinorUnit(amount int, info currencyInfo) string {
	word := NumberToWords(amount)
	if amount == 1 {
		return word + " " + info.minorSingular
	}
	return word + " " + info.minorPlural
}
