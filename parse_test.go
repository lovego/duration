package duration

import (
	"fmt"
	"time"
)

func ExampleParse_empty() {
	fmt.Println(Parse(``))
	fmt.Println(Parse(`0`))
	fmt.Println(Parse(`+0`))
	fmt.Println(Parse(`-0`))

	// Output:
	// 0 <nil>
	// 0 <nil>
	// 0 <nil>
	// 0 <nil>
}

func ExampleParse_EN() {
	fmt.Println(Parse(`123s`))
	fmt.Println(Parse(`year`))
	fmt.Println(Parse(`1Y2M3D4h306s`))
	fmt.Println(Parse(`1.1m`))

	// Output:
	// 2m3s <nil>
	// 1Y <nil>
	// 1Y2M3D4h5m6s <nil>
	// 1m6s <nil>
}

func ExampleParse_ZH() {
	fmt.Println(Parse(`123秒`))
	fmt.Println(Parse(`年`))
	fmt.Println(Parse(`1年2月3日4时306秒`))
	fmt.Println(Parse(`1.1分`))

	// Output:
	// 2分3秒 <nil>
	// 1年 <nil>
	// 1年2个月3天4小时5分6秒 <nil>
	// 1分6秒 <nil>
}

func ExampleTime_Parse() {
	fmt.Println(time.ParseDuration(`1h3m5`))
	// Output:
	// 0s time: missing unit in duration 1h3m5
}
