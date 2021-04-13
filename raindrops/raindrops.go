// Package raindrops will handle raindrops sounds
package raindrops

import "fmt"

// Convert receive num integer and return sounds string based on number's factor
func Convert(num int) string {
	var result string
	if num%3 == 0 {
		result += "Pling"
	}
	if num%5 == 0 {
		result += "Plang"
	}
	if num%7 == 0 {
		result += "Plong"
	}
	if (num%3 != 0) && (num%5 != 0) && (num%7 != 0) {
		result += fmt.Sprint(num)
	}
	return result
}
