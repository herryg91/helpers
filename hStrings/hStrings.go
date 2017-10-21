package hStrings

import "bytes"

//Concatenate faster concatenate string.
func Concatenate(str ...string) string {
	var buffer bytes.Buffer
	for _, v := range str {
		buffer.WriteString(v)
	}
	return buffer.String()
}

//Reverse make string reverse. Ex: Hello World => dlroW olleH
func Reverse(s string) (result string) {
	for _, v := range s {
		result = Concatenate(string(v), result)
	}
	return
}

// //NumberFormat give thousand dot separator
// func NumberFormat(number string) string {
// 	number = Reverse(number)

// 	var numberStr string
// 	for i, val := range number {
// 		if i%3 == 0 && i != 0 {
// 			numberStr = numberStr + "."
// 		}
// 		numberStr = numberStr + string(val)
// 	}

// 	return Reverse(numberStr)
// }

// //RupiahFormat format string of number rupiah format including thousan dot separator
// func RupiahFormat(priceStr string) string {
// 	var priceRupiah = Reverse(NumberFormat(priceStr))

// 	priceRupiah = priceRupiah + " pR"
// 	return Reverse(priceRupiah)

// }
