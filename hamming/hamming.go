//package hamming will handle DNA strands function related with hamming counting
package hamming

import "errors"

//Distance will accept 2 string of 2 strands DNA
//it will return int and error:
// int: number of hamming
//error: error condition if strands lenght is not equal or nil of no error happen
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("error: sequences length is not equal")
	}
	var counter int = 0
	for i, _ := range a {
		if a[i] != b[i] {
			counter++
		}
	}
	return counter, nil
}
