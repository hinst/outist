package outist

import "time"

func FormatFullTime(time time.Time) string {
	var toS = func(i int) string {
		return IntToStr(i)
	}
	var toS2 = func(i int) string {
		return Add0ToStr(IntToStr(i), 2)
	}
	var toS3 = func(i int) string {
		return Add0ToStr(IntToStr(i), 3)
	}
	return "" + toS(time.Year()) +
		"." + toS2(int(time.Month())) +
		"." + toS2(time.Day()) +
		"-" + toS2(time.Hour()) +
		":" + toS2(time.Minute()) +
		":" + toS2(time.Second()) +
		"." + toS3(time.Nanosecond()/1000/1000)
}

func FormatFileTime(time time.Time) string {
	var toS = func(i int) string {
		return IntToStr(i)
	}
	var toS2 = func(i int) string {
		return Add0ToStr(IntToStr(i), 2)
	}
	return "" + toS(time.Year()) +
		"-" + toS2(int(time.Month())) +
		"-" + toS2(time.Day()) +
		"_" + toS2(time.Hour()) +
		"-" + toS2(time.Minute()) +
		"-" + toS2(time.Second())
}

func FormateDateForFileName(time time.Time) string {
	var yearString = IntToStr(time.Year())
	var monthString = Add0ToStr(IntToStr(int(time.Month())), 2)
	var dayString = Add0ToStr(IntToStr(time.Day()), 2)
	return yearString + "-" + monthString + "-" + dayString
}
