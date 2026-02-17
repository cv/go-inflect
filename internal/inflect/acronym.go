package inflect

import (
	"slices"
	"sort"
	"strings"
	"unicode"
)

// defaultAcronyms contains sensible default acronyms that should preserve their case.
var defaultAcronyms = []string{
	"ACL", "API", "AWS", "CPU", "CRUD", "CSS", "DNS", "GCP", "GPU",
	"gRPC", "HTML", "HTTP", "HTTPS", "ID", "JSON", "LDAP", "OAuth",
	"REST", "SQL", "SSH", "SSL", "TLS", "URI", "URL", "UUID", "VM", "XML",
}

// AddAcronym registers an acronym that should preserve its case in humanization.
//
// Acronyms are matched case-insensitively. For example, AddAcronym("GPU") will
// preserve the case for "GPU", "gpu", or "Gpu" as "GPU".
//
// Examples:
//
//	AddAcronym("GPU")
//	Humanize("GPUConfig")  // returns "GPU config"
//	Humanize("myGPU")      // returns "My GPU"
func AddAcronym(acronym string) {
	defaultEngine.AddAcronym(acronym)
}

// AddAcronym registers an acronym that should preserve its case in humanization.
//
// Acronyms are matched case-insensitively. For example, AddAcronym("GPU") will
// preserve the case for "GPU", "gpu", or "Gpu" as "GPU".
//
// Examples:
//
//	e := NewEngine()
//	e.AddAcronym("GPU")
//	e.Humanize("GPUConfig")  // returns "GPU config"
//	e.Humanize("myGPU")      // returns "My GPU"
func (e *Engine) AddAcronym(acronym string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.acronyms == nil {
		e.acronyms = make(map[string]string)
		// Initialize with defaults
		for _, a := range defaultAcronyms {
			e.acronyms[strings.ToUpper(a)] = a
		}
	}
	e.acronyms[strings.ToUpper(acronym)] = acronym
}

// RemoveAcronym removes an acronym from the registry.
//
// Returns true if the acronym was removed, false if it wasn't registered.
//
// Examples:
//
//	RemoveAcronym("GPU")
//	Humanize("GPUConfig")  // returns "Gpu config"
func RemoveAcronym(acronym string) bool {
	return defaultEngine.RemoveAcronym(acronym)
}

// RemoveAcronym removes an acronym from the registry.
//
// Returns true if the acronym was removed, false if it wasn't registered.
//
// Examples:
//
//	e := NewEngine()
//	e.RemoveAcronym("GPU")
//	e.Humanize("GPUConfig")  // returns "Gpu config"
func (e *Engine) RemoveAcronym(acronym string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.acronyms == nil {
		return false
	}
	upper := strings.ToUpper(acronym)
	if _, exists := e.acronyms[upper]; !exists {
		return false
	}
	delete(e.acronyms, upper)
	return true
}

// ClearAcronyms removes all registered acronyms, including defaults.
//
// Example:
//
//	ClearAcronyms()
//	Humanize("GPUConfig")  // returns "Gpu config"
func ClearAcronyms() {
	defaultEngine.ClearAcronyms()
}

// ClearAcronyms removes all registered acronyms, including defaults.
//
// Example:
//
//	e := NewEngine()
//	e.ClearAcronyms()
//	e.Humanize("GPUConfig")  // returns "Gpu config"
func (e *Engine) ClearAcronyms() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.acronyms = make(map[string]string)
}

// ResetAcronyms restores the acronym registry to its default state.
//
// Example:
//
//	ClearAcronyms()
//	ResetAcronyms()
//	Humanize("GPUConfig")  // returns "GPU config"
func ResetAcronyms() {
	defaultEngine.ResetAcronyms()
}

// ResetAcronyms restores the acronym registry to its default state.
//
// Example:
//
//	e := NewEngine()
//	e.ClearAcronyms()
//	e.ResetAcronyms()
//	e.Humanize("GPUConfig")  // returns "GPU config"
func (e *Engine) ResetAcronyms() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.acronyms = make(map[string]string)
	for _, a := range defaultAcronyms {
		e.acronyms[strings.ToUpper(a)] = a
	}
}

// GetAcronyms returns a sorted list of all registered acronyms.
//
// Example:
//
//	acronyms := GetAcronyms()
//	// returns ["ACL", "API", "AWS", "CPU", ...] (sorted)
func GetAcronyms() []string {
	return defaultEngine.GetAcronyms()
}

// GetAcronyms returns a sorted list of all registered acronyms.
//
// Example:
//
//	e := NewEngine()
//	acronyms := e.GetAcronyms()
//	// returns ["ACL", "API", "AWS", "CPU", ...] (sorted)
func (e *Engine) GetAcronyms() []string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if e.acronyms == nil {
		result := make([]string, len(defaultAcronyms))
		copy(result, defaultAcronyms)
		sort.Strings(result)
		return result
	}
	result := make([]string, 0, len(e.acronyms))
	for _, v := range e.acronyms {
		result = append(result, v)
	}
	sort.Strings(result)
	return result
}

// IsAcronym checks if a word is a registered acronym.
//
// The check is case-insensitive.
//
// Examples:
//
//	IsAcronym("GPU")  // returns true (if GPU is registered)
//	IsAcronym("gpu")  // returns true (case-insensitive)
//	IsAcronym("Cat")  // returns false
func IsAcronym(word string) bool {
	return defaultEngine.IsAcronym(word)
}

// IsAcronym checks if a word is a registered acronym.
//
// The check is case-insensitive.
//
// Examples:
//
//	e := NewEngine()
//	e.IsAcronym("GPU")  // returns true (if GPU is registered)
//	e.IsAcronym("gpu")  // returns true (case-insensitive)
//	e.IsAcronym("Cat")  // returns false
func (e *Engine) IsAcronym(word string) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if e.acronyms == nil {
		// Check against defaults
		for _, a := range defaultAcronyms {
			if strings.EqualFold(a, word) {
				return true
			}
		}
		return false
	}
	_, exists := e.acronyms[strings.ToUpper(word)]
	return exists
}

// getAcronymCase returns the preferred case for a word if it's a registered acronym.
// Returns the word unchanged if it's not an acronym.
func (e *Engine) getAcronymCase(word string) string {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if e.acronyms == nil {
		// Check against defaults
		for _, a := range defaultAcronyms {
			if strings.EqualFold(a, word) {
				return a
			}
		}
		return word
	}
	if preferred, exists := e.acronyms[strings.ToUpper(word)]; exists {
		return preferred
	}
	return word
}

// SplitPascalCase splits a PascalCase identifier into words.
//
// Consecutive uppercase letters are treated as acronyms. The function handles:
//   - PascalCase: "HelloWorld" -> ["Hello", "World"]
//   - Acronyms: "GPUConfig" -> ["GPU", "Config"]
//   - Mixed: "CoreweaveGPUStatus" -> ["Coreweave", "GPU", "Status"]
//   - Numbers: "OAuth2Client" -> ["O", "Auth", "2", "Client"]
//   - Leading lowercase: "myXMLParser" -> ["my", "XML", "Parser"]
//
// Examples:
//
//	SplitPascalCase("GPUConfig")          // ["GPU", "Config"]
//	SplitPascalCase("DNSRecord")          // ["DNS", "Record"]
//	SplitPascalCase("LDAPUser")           // ["LDAP", "User"]
//	SplitPascalCase("CoreweaveGPUStatus") // ["Coreweave", "GPU", "Status"]
//	SplitPascalCase("myXMLParser")        // ["my", "XML", "Parser"]
func SplitPascalCase(s string) []string {
	return splitIntoWords(s)
}

// isAllUppercase checks if a string contains only uppercase letters.
func isAllUppercase(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

// DefaultAcronyms returns a copy of the default acronyms list.
func DefaultAcronyms() []string {
	result := make([]string, len(defaultAcronyms))
	copy(result, defaultAcronyms)
	slices.Sort(result)
	return result
}
