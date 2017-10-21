package hNumbers

import (
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/herryg91/helpers/hStrings"
)

//Round round number
func Round(number float64, decimal int) float64 {
	if decimal < 0 {
		decimal = 0
	}
	rounder := math.Pow10(decimal)
	adder := 0.5
	if number < 0 {
		adder = -0.5
	}
	return float64(int((number*rounder)+adder)) / rounder
}

//IsMinus check if number minus
func IsMinus(number interface{}) (result bool) {
	switch reflect.TypeOf(number).Kind() {
	case reflect.Int:
		result = number.(int) < 0
	case reflect.Int8:
		result = number.(int8) < 0
	case reflect.Int16:
		result = number.(int16) < 0
	case reflect.Int32:
		result = number.(int32) < 0
	case reflect.Int64:
		result = number.(int64) < 0
	case reflect.Float32:
		result = number.(float32) < 0
	case reflect.Float64:
		result = number.(float64) < 0
	}
	return
}

//NumberFormat make number format
func NumberFormat(number float64, decimal int, decimalPoint string, thousandSeparator string) (result string) {
	if decimalPoint == "" {
		decimalPoint = ","
	}
	if thousandSeparator == "" {
		thousandSeparator = "."
	}

	isMinus := IsMinus(number)
	if isMinus {
		number *= -1
	}

	numberInString := strconv.FormatFloat(number, 'f', decimal, 64)
	splitNumber := strings.Split(numberInString, ".")

	/*Get Fraction*/
	fraction := ""
	if len(splitNumber) > 1 {
		fraction = splitNumber[1]
	}

	/*Append Thousand Separator*/
	pos := 0
	if len(splitNumber[0])%3 != 0 {
		pos += len(splitNumber[0]) % 3
		result = hStrings.Concatenate(result, splitNumber[0][:pos])
	}
	for ; pos < len(splitNumber[0]); pos += 3 {
		if result != "" {
			result = hStrings.Concatenate(result, thousandSeparator)
		}
		result = hStrings.Concatenate(result, splitNumber[0][pos:pos+3])
	}
	if fraction != "" {
		result = hStrings.Concatenate(result, decimalPoint, fraction)
	}
	if isMinus {
		result = hStrings.Concatenate("-", result)
	}
	return result
}
