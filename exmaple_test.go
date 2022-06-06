package holiday_test

import (
	"fmt"
	"github.com/Oppodelldog/holiday"
	"time"
)

func ExampleQuery_IsHoliday() {
	var (
		date   = time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local)
		query  = holiday.New()
		result = query.IsHoliday(date, holiday.Any)
	)

	if result {
		fmt.Println("it is a holiday")
	} else {
		fmt.Println("sorry, no holiday")
	}

	// Output: it is a holiday
}

func ExampleIsHoliday() {
	var (
		date   = time.Date(2024, 12, 25, 0, 0, 0, 0, time.Local)
		result = holiday.IsHoliday(date, holiday.Any)
	)

	if result {
		fmt.Println("it is a holiday")
	} else {
		fmt.Println("sorry, no holiday")
	}

	// Output: it is a holiday
}
