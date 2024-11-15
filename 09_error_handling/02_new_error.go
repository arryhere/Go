package error_handling

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func SafeFloatInput() (float64, float64, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Dividend: ")
	dividend_read_value, dividend_read_error := reader.ReadString('\n')
	if dividend_read_error != nil {
		return 0.0, 0.0, dividend_read_error
	}
	dividend_trim := strings.TrimSpace(dividend_read_value)
	dividend, dividend_parse_float_error := strconv.ParseFloat(dividend_trim, 64)
	if dividend_parse_float_error != nil {
		return 0.0, 0.0, errors.New("[Error]: Enter a valid Dividend [float64]")
	}

	fmt.Print("Enter Divisor: ")
	divisor_read_value, divisor_read_error := reader.ReadString('\n')
	if divisor_read_error != nil {
		return 0.0, 0.0, divisor_read_error
	}
	divisor_trim := strings.TrimSpace(divisor_read_value)
	divisor, divisor_parse_float_error := strconv.ParseFloat(divisor_trim, 64)
	if divisor_parse_float_error != nil {
		return 0.0, 0.0, errors.New("[Error]: Enter a valid Divisor [float64]")
	}

	return dividend, divisor, nil

}

func SafeDivide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.0, errors.New("[Error]: Divisor cannot be 0")
	}
	return dividend / divisor, nil
}

func New_Error_2() {
	dividend, divisor, safe_float_input_error := SafeFloatInput()

	if safe_float_input_error != nil {
		fmt.Println(safe_float_input_error.Error())
		return
	}

	safe_divide_res, safe_divide_error := SafeDivide(dividend, divisor)

	if safe_divide_error != nil {
		fmt.Println(safe_divide_error.Error())
		return
	}

	fmt.Printf("res: [%v], type: [%v]\n", safe_divide_res, reflect.TypeOf(safe_divide_error))
}
