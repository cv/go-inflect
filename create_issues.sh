#!/bin/bash

# Create epics for main feature areas
echo "Creating epics..."
EPIC_PLURAL=$(bd create "Pluralization" -t epic --silent)
EPIC_SINGULAR=$(bd create "Singularization" -t epic --silent)
EPIC_ARTICLES=$(bd create "Indefinite Articles (a/an)" -t epic --silent)
EPIC_NUMBERS=$(bd create "Number Conversion" -t epic --silent)
EPIC_ORDINALS=$(bd create "Ordinals" -t epic --silent)
EPIC_JOIN=$(bd create "Word Joining" -t epic --silent)
EPIC_CLASSICAL=$(bd create "Classical Mode" -t epic --silent)
EPIC_PARTICIPLE=$(bd create "Present Participle" -t epic --silent)
EPIC_COMPARE=$(bd create "Comparison Functions" -t epic --silent)
EPIC_CUSTOM=$(bd create "Custom Definitions" -t epic --silent)
EPIC_UTIL=$(bd create "Utility Functions" -t epic --silent)
EPIC_SETUP=$(bd create "Project Setup" -t epic --silent)
EPIC_DOCS=$(bd create "Documentation" -t epic --silent)

echo "Created epics. Now creating tasks..."

# Pluralization tasks
bd create "Implement plural() - unconditional pluralization" --parent $EPIC_PLURAL --silent
bd create "Implement plural_noun() - pluralize nouns" --parent $EPIC_PLURAL --silent
bd create "Implement plural_verb() - pluralize verbs" --parent $EPIC_PLURAL --silent
bd create "Implement plural_adj() - pluralize adjectives" --parent $EPIC_PLURAL --silent
bd create "Test 860+ inflection test cases from inflections.txt" --parent $EPIC_PLURAL --silent
bd create "Implement compound word pluralization (hyphenated and spaced)" --parent $EPIC_PLURAL --silent
bd create "Implement unit handling (degree celsius, per square inch, etc)" --parent $EPIC_PLURAL --silent
bd create "Implement open compound nouns (high school, prima donna)" --parent $EPIC_PLURAL --silent

# Singularization tasks
bd create "Implement singular_noun() - convert plural to singular" --parent $EPIC_SINGULAR --silent
bd create "Test singular_noun() with all plural forms" --parent $EPIC_SINGULAR --silent

# Indefinite Articles tasks
bd create "Implement a() and an() - indefinite article selection" --parent $EPIC_ARTICLES --silent
bd create "Test a/an with common words (cat, ant, honest, etc)" --parent $EPIC_ARTICLES --silent
bd create "Test a/an with abbreviations (YAML, JSON, Core ML, etc)" --parent $EPIC_ARTICLES --silent
bd create "Test a/an with silent letters and special cases" --parent $EPIC_ARTICLES --silent

# Number Conversion tasks
bd create "Implement number_to_words() - convert numbers to words" --parent $EPIC_NUMBERS --silent
bd create "Support basic numbers (0-20, 21-99, 100s, 1000s, millions, etc)" --parent $EPIC_NUMBERS --silent
bd create "Support decimal numbers" --parent $EPIC_NUMBERS --silent
bd create "Support threshold parameter (don't convert above threshold)" --parent $EPIC_NUMBERS --silent
bd create "Support comma parameter (format numbers with commas)" --parent $EPIC_NUMBERS --silent
bd create "Support group parameter (group digits by N)" --parent $EPIC_NUMBERS --silent
bd create "Support decimal parameter (decimal point character)" --parent $EPIC_NUMBERS --silent
bd create "Test 100+ number_to_words test cases" --parent $EPIC_NUMBERS --silent

# Ordinals tasks
bd create "Implement ordinal() - convert number to ordinal (1st, 2nd, 3rd)" --parent $EPIC_ORDINALS --silent
bd create "Support numeric ordinals (1st, 2nd, 3rd, 21st, etc)" --parent $EPIC_ORDINALS --silent
bd create "Support word ordinals (first, second, third, etc)" --parent $EPIC_ORDINALS --silent
bd create "Test ordinal() with number_to_words() roundtrip" --parent $EPIC_ORDINALS --silent

# Word Joining tasks
bd create "Implement join() - join words with proper grammar" --parent $EPIC_JOIN --silent
bd create "Support customizable conjunction (and/or)" --parent $EPIC_JOIN --silent
bd create "Support customizable final separator" --parent $EPIC_JOIN --silent
bd create "Support comma vs semicolon detection" --parent $EPIC_JOIN --silent
bd create "Test join with 1, 2, 3+ words" --parent $EPIC_JOIN --silent

# Classical Mode tasks
bd create "Implement classical() mode toggle" --parent $EPIC_CLASSICAL --silent
bd create "Support classical(all=True/False)" --parent $EPIC_CLASSICAL --silent
bd create "Support classical(zero=True) - '0 errors' vs '0 error'" --parent $EPIC_CLASSICAL --silent
bd create "Support classical(herd=True) - 'wildebeest' vs 'wildebeests'" --parent $EPIC_CLASSICAL --silent
bd create "Support classical(names=True) - proper name pluralization" --parent $EPIC_CLASSICAL --silent
bd create "Support classical(ancient=True) - Latin/Greek plurals (formulae)" --parent $EPIC_CLASSICAL --silent
bd create "Support classical(persons=True) - 'persons' vs 'people'" --parent $EPIC_CLASSICAL --silent
bd create "Test all classical mode combinations" --parent $EPIC_CLASSICAL --silent

# Present Participle task
bd create "Implement present_participle() - verb to -ing form" --parent $EPIC_PARTICIPLE --silent

# Comparison Functions tasks
bd create "Implement compare() - compare singular/plural forms" --parent $EPIC_COMPARE --silent
bd create "Implement compare_nouns() - compare noun forms" --parent $EPIC_COMPARE --silent
bd create "Implement compare_verbs() - compare verb forms" --parent $EPIC_COMPARE --silent
bd create "Implement compare_adjs() - compare adjective forms" --parent $EPIC_COMPARE --silent

# Custom Definitions tasks
bd create "Implement defnoun() - define custom noun pluralization" --parent $EPIC_CUSTOM --silent
bd create "Implement defverb() - define custom verb conjugation" --parent $EPIC_CUSTOM --silent
bd create "Implement defadj() - define custom adjective pluralization" --parent $EPIC_CUSTOM --silent
bd create "Implement defa() - define custom a/an pattern" --parent $EPIC_CUSTOM --silent
bd create "Implement defan() - define custom a/an pattern" --parent $EPIC_CUSTOM --silent
bd create "Support regex patterns in custom definitions" --parent $EPIC_CUSTOM --silent

# Utility Functions tasks
bd create "Implement num() - store/retrieve default count" --parent $EPIC_UTIL --silent
bd create "Implement no() - 'no N' vs 'N Xs'" --parent $EPIC_UTIL --silent
bd create "Implement inflect() - parse and inflect inline text" --parent $EPIC_UTIL --silent
bd create "Implement gender() - set pronoun gender" --parent $EPIC_UTIL --silent

# Project Setup tasks
bd create "Initialize Go module and project structure" --parent $EPIC_SETUP --silent
bd create "Set up testing framework" --parent $EPIC_SETUP --silent
bd create "Create README with examples" --parent $EPIC_SETUP --silent
bd create "Set up CI/CD" --parent $EPIC_SETUP --silent
bd create "Write contribution guidelines" --parent $EPIC_SETUP --silent

# Documentation tasks
bd create "Document all public APIs" --parent $EPIC_DOCS --silent
bd create "Create comprehensive examples" --parent $EPIC_DOCS --silent
bd create "Port README examples from Python" --parent $EPIC_DOCS --silent
bd create "Create performance benchmarks" --parent $EPIC_DOCS --silent

echo "Done! Created all issues."
