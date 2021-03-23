// Package leap will handle function regarding leap year
package leap

// IsLeapYear receive year in integer and return boolean type whether it is leap year
func IsLeapYear(year int) bool {
	if ((year % 4) == 0) && ((year % 100) != 0) {
		return true
	} else if (year%4 == 0) && ((year % 100) == 0) {
		if (year % 400) == 0 {
			return true
		}
	}
	return false
}
