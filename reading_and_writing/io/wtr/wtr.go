/*
Assignment is to see if people are "cooking" accounting books by using Benfords Law which checks the distrubtion of leading digits in numbers
- Law states that leading digits with smaller values occur more frequently than larger values
- States that about 30% of numbers start with a 1 while less than 5% start with a 9
*/

package wtr

import "unicode"

// DigitsFreq calculates leading digit frequency
type DigitsFreq struct {
	Freqs map[rune]int // leading digit frequency
	inNum bool         // local state
}

// implementing io.Writer. You get a []byte and update the leading digit distribution
// Write implements io.Writer
func (d *DigitsFreq) Write(data []byte) (int, error) {
	if d.Freqs == nil {
		d.Freqs = make(map[rune]int)
	}

	for _, b := range data {
		if r := rune(b); unicode.IsDigit(r) {
			if !d.inNum {
				d.Freqs[r]++
				d.inNum = true
			}
			continue
		}
		// here its not a digit
		if d.inNum {
			d.inNum = false
		}
	}

	return len(data), nil
}