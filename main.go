package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ValidatePin(pin string, rulers ...Ruler) bool {
	for _, ruler := range rulers {
		if !ruler.Check(pin) {
			return false
		}
	}
	return true
}

type Ruler interface {
	Check(pin string) bool
}

type MinDigits struct {
	min int
}

func (c *MinDigits) Check(pin string) bool {
	re := regexp.MustCompile(fmt.Sprintf(`^\d{%v,}$`, c.min))
	isValid := re.MatchString(pin)
	return isValid
}

type CheckNotRepeatedAdjacentNumber struct{}

func (c *CheckNotRepeatedAdjacentNumber) Check(pin string) bool {
	repeatRegex := regexp.MustCompile(`0{2,}|1{2,}|2{2,}|3{2,}|4{2,}|5{2,}|6{2,}|7{2,}|8{2,}|9{2,}`)
	isValid := !repeatRegex.MatchString(pin)
	return isValid
}

type CheckNotRepeatedAllNumber struct{}

func (c *CheckNotRepeatedAllNumber) Check(pin string) bool {
	repeatRegex := regexp.MustCompile(`^1+$|^2+$|^3+$|^4+$|^5+$|^6+$|^7+$|^8+$|^9+$`)
	isValid := !repeatRegex.MatchString(pin)
	return isValid
}

type CheckNotSequence struct{}

func (c *CheckNotSequence) Check(pin string) bool {
	isValid := !strings.Contains("0123456789", pin) &&
		!strings.Contains("9876543210", pin)
	return isValid
}

type CheckMinDifferentCharacter struct {
	min int
}

func (c *CheckMinDifferentCharacter) Check(pin string) bool {
	counter := make(map[rune]bool)
	for _, c := range pin {
		counter[c] = true
	}

	return len(counter) >= c.min
}

func main() {
	for true {
		fmt.Print("Enter pin: ")
		var input string
		fmt.Scanln(&input)
		pin := strings.Trim(input, "\n")
		fmt.Printf("Pin la: %s\n", pin)
		isValid := ValidatePin(
			pin,
			&MinDigits{min: 6},
			&CheckNotSequence{},
			&CheckNotRepeatedAllNumber{},
		)
		fmt.Printf("%v\n", isValid)
	}
}
