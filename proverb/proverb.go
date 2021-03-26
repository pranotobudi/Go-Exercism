// Package proverb will handle list of words and turn it into proverb.
// https://golang.org/doc/effective_go.html#commentary
package proverb

// Proverb function will accept words as slice of string and return proverbs in a slice of string.
func Proverb(rhyme []string) []string {
	var result []string
	var text string
	for i, _ := range rhyme {
		if i == len(rhyme)-1 {
			text = "And all for the want of a " + rhyme[0] + "."
			result = append(result, text)
		} else {
			text = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
			result = append(result, text)
		}
	}
	return result
}
