package timerounding_test

import (
	"fmt"
	"time"

	"github.com/andreyvit/timerounding"
)

func ExampleFormatSet_Format() {
	fmt.Println(timerounding.Concise.Format(time.Date(2017, 1, 7, 9, 37, 12, 0, time.UTC), timerounding.Hours))
	// Output: 20170107-09
}
