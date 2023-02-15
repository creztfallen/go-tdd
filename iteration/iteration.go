package iteration

import "strings"

func Repeat(stringToBeRepeated string, numberToRepeat int) string {
	var repetitions string
	// for i := 0; i < numberToRepeat; i++ {
	// 	repetitions += stringToBeRepeated
	// }
	repetitions = strings.Repeat(stringToBeRepeated, numberToRepeat)
	return repetitions
}
