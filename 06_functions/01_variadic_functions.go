package functions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func safeArgsInput() ([]int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter space separated arguments: ")
	args_read_value, args_read_error := reader.ReadString('\n')
	if args_read_error != nil {
		return nil, args_read_error
	}

	args_list := strings.Split(strings.TrimSpace(args_read_value), " ")

	int_list := []int{}

	for i := 0; i < len(args_list); i++ {
		int_value, err := strconv.ParseInt(args_list[i], 10, 0)

		if err != nil {
			return nil, fmt.Errorf("[Error]: Invalid argument [%v], cannot parse to integer", args_list[i])
		}

		int_list = append(int_list, int(int_value))
	}

	return int_list, nil
}

func multiply(args ...int) (int, error) {
	if len(args) == 0 {
		return 0, errors.New("[Error]: No arguments provided")
	}
	res := 1

	for i := 0; i < len(args); i++ {
		res *= args[i]
	}

	return res, nil
}

func Variadic_Functions() {
	safeArgsInputValue, safeArgsInputError := safeArgsInput()

	if safeArgsInputError != nil {
		fmt.Println(safeArgsInputError)
		return
	}

	multiplyValue, multiplyError := multiply(safeArgsInputValue...)

	if multiplyError != nil {
		fmt.Println(multiplyError)
		return
	}

	fmt.Println("res:", multiplyValue) // res: 1239504
}
