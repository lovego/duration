package duration

import (
	"encoding/json"
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

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
