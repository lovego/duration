package duration

import (
	"log"
)

func init() {
	if err := RegisterUnitSet(EN); err != nil {
		log.Panic(err)
	}
	if err := RegisterUnitSet(ZH); err != nil {
		log.Panic(err)
	}
}

var EN = []Unit{
	{Value: Year, Name: "Y", OtherNames: []string{"year", "years"}},
	{Value: Month, Name: "M", OtherNames: []string{"month", "months"}},
	{Value: Week, Name: "W", OtherNames: []string{"week", "weeks"}},
	{Value: Day, Name: "D", OtherNames: []string{"day", "days"}},

	{Value: Hour, Name: "h", OtherNames: []string{"hour", "hours"}},
	{Value: Minute, Name: "m", OtherNames: []string{"minute", "minutes"}},
	{Value: Second, Name: "s", OtherNames: []string{"second", "seconds"}},

	{Value: Millisecond, Name: "ms", OtherNames: []string{"millisecond", "milliseconds"}},
	{Value: Microsecond, Name: "us", OtherNames: []string{
		"µs", // U+00B5 = micro symbol
		"μs", // U+03BC = Greek letter mu
		"microsecond", "microseconds",
	}},
	{Value: Nanosecond, Name: "ns", OtherNames: []string{"nanosecond", "nanoseconds"}},
}

var ZH = []Unit{
	{Value: Year, Name: "年", OtherNames: []string{}},
	{Value: Month, Name: "个月", OtherNames: []string{"月"}},
	{Value: Week, Name: "周", OtherNames: []string{"星期", "个星期"}},
	{Value: Day, Name: "天", OtherNames: []string{"日"}},

	{Value: Hour, Name: "小时", OtherNames: []string{"时"}},
	{Value: Minute, Name: "分", OtherNames: []string{"分钟"}},
	{Value: Second, Name: "秒", OtherNames: []string{"秒钟"}},

	{Value: Millisecond, Name: "毫秒", OtherNames: []string{}},
	{Value: Microsecond, Name: "微秒", OtherNames: []string{}},
	{Value: Nanosecond, Name: "纳秒", OtherNames: []string{}},
}
