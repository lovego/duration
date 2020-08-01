package duration

import (
	"fmt"
	"time"
)

func ExampleParseDuation() {
	fmt.Println(Parse(`1h3m5`))
	fmt.Println(Parse(`1h3m5`))

	// Output:
	// 3785 <nil>
	// 3785 <nil>
}

func ExampleParseDuation_noUnit() {
	fmt.Println(Parse(``))
	fmt.Println(Parse(`0`))
	fmt.Println(Parse(`1`))
	fmt.Println(Parse(`12`))
	fmt.Println(Parse(`123`))

	// Output:
	// 0 <nil>
	// 0 <nil>
	// 1 <nil>
	// 12 <nil>
	// 123 <nil>
}

func ExampleParseDuation_seconds() {
	fmt.Println(Parse(`0s`))
	fmt.Println(Parse(`1s`))
	fmt.Println(Parse(`12s`))
	fmt.Println(Parse(`123s`))

	// Output:
	// 0 <nil>
	// 1 <nil>
	// 12 <nil>
	// 123 <nil>
}

func ExampleParseDuation_minutes() {
	fmt.Println(Parse(`0m`))
	fmt.Println(Parse(`1m`))
	fmt.Println(Parse(`12m`))
	fmt.Println(Parse(`123m`))

	// Output:
	// 0 <nil>
	// 60 <nil>
	// 720 <nil>
	// 7380 <nil>
}

func ExampleTime_ParseDuation() {
	fmt.Println(time.ParseDuration(`1h3m5`))

	// Output:
	// 0s time: missing unit in duration 1h3m5
}
