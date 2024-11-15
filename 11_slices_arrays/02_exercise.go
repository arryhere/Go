package slices_arrays

import "fmt"

type cost struct {
	day   int
	value float64
}

func formatByDays(costList []cost) []string {
	result := []string{}

	valueToInsert := 0.0

	for i := 0; i < len(costList); i++ {
		if valueToInsert == 0 {
			valueToInsert = costList[i].value
		}

		if i == len(costList)-1 {
			result = append(result, fmt.Sprintf("Day [%v], value: [%v]", costList[i].day, valueToInsert))
			break
		}

		if costList[i].day == costList[i+1].day {
			valueToInsert += costList[i+1].value
			continue
		}
		result = append(result, fmt.Sprintf("Day [%v], value: [%v]", costList[i].day, valueToInsert))

		valueToInsert = 0
	}

	return result
}

func formatByDaysIncludingMissing(costList []cost) []string {
	// Find the maximum day in the costList
	maxDay := 0
	for i := range costList {
		if costList[i].day > maxDay {
			maxDay = costList[i].day
		}
	}

	// Create a slice to store cumulative values for each day
	dayValues := make([]float64, maxDay+1) // maxDay+1 as day is starting from 0

	// Populate the slice with the values from costList
	for i := range costList {
		dayValues[costList[i].day] += costList[i].value
	}

	// Generate the result slice
	result := []string{}
	for i := range dayValues {
		result = append(result, fmt.Sprintf("Day [%v], value: [%v]", i, dayValues[i]))
	}

	return result
}

func Exercise() {
	costList := []cost{
		{day: 0, value: 4.0},
		{day: 1, value: 2.1},
		{day: 2, value: 1.1},
		{day: 2, value: 9.1},
		{day: 5, value: 5.2},
	}

	res := formatByDays(costList)
	res2 := formatByDaysIncludingMissing(costList)

	for i := range len(res) {
		fmt.Println(res[i])
	}

	fmt.Println("-------------------------")

	for i := range len(res2) {
		fmt.Println(res2[i])
	}
}
