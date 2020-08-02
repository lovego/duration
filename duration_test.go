package duration

import (
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

func ExampleDuration_String() {
	d, _ := Parse(`2m3s`)
	fmt.Println(d)

	d, _ = Parse(`-1Y2M9D13h5m`)
	d.N = 2
	fmt.Println(d)

	// Output:
	// 2m3s
	// -1Y2M
}

func ExampleDuration_UnmarshalYAML() {
	var d Duration
	if err := yaml.Unmarshal([]byte(`123秒`), &d); err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)

	if err := yaml.Unmarshal([]byte(`"123秒"`), &d); err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
	// Output:
	// 2分3秒
	// 2分3秒
}

func ExampleDuration_UnmarshalJSON() {
	var d Duration
	if err := json.Unmarshal([]byte(`"123秒"`), &d); err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
	// Output: 2分3秒
}

func ExampleDuration_Truncate() {
	d, _ := Parse(`2m3s`)
	fmt.Println(d.Truncate(Minute))

	d, _ = Parse(`-1Y2M9D13h5m`)
	fmt.Println(d.Truncate(Day))

	// Output:
	// 2m
	// -1Y2M9D
}

func ExampleDuration_Round() {
	d, _ := Parse(`2m29s`)
	fmt.Println(d.Round(Minute))

	d, _ = Parse(`2m30s`)
	fmt.Println(d.Round(Minute))

	d, _ = Parse(`-2m29s`)
	fmt.Println(d.Round(Minute))

	d, _ = Parse(`-2m30s`)
	fmt.Println(d.Round(Minute))

	// Output:
	// 2m
	// 3m
	// -2m
	// -3m
}
