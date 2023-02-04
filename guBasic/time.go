package guBasic

import (
	"fmt"
	"strings"
	"time"
)

// time format
func FindTimeFormat(datetime string) (format string) {
	if len(datetime) == 8 && !strings.Contains(datetime, ":") {
		format = "20060102"
	} else if len(datetime) == 8 && strings.Contains(datetime, ":") {
		format = "03:04:05"
	} else if len(datetime) == 17 {
		format = "20060102 03:04:05"
	} else if len(datetime) == 14 {
		format = "20060102 03:04"
	} else if len(datetime) == 5 {
		format = "03:04"
	} else if len(datetime) == 4 {
		format = "2006"
	} else if len(datetime) == 6 {
		format = "200601"
	}
	return
}

// format: "20060102"
func TimeFromStr(datetime, format string) (t time.Time) {
	t, err := time.Parse(format, datetime)
	if err != nil {
		fmt.Printf("\n%v\n", err)
	}
	return t
}

func StrFromTime(datetime time.Time, format string) string {
	return datetime.Format(format)
}

// time(str) -> timestamp
// TimeStampFromTimeStr("20230202", "20060102") -> 1675296000
func TimeStampFromTimeStr(datetime, format string) int64 {
	t, err := time.Parse(format, datetime)
	if err != nil {
		fmt.Printf("\n%v\n", err)
	}
	return t.Unix()
}

// timestamp -> time(str)
// TimeStrFromTimeStamp(1534809600, "20060102") -> 20180821
func TimeStrFromTimeStamp(timestamp int64, format string) string {
	return time.Unix(timestamp, 0).Format(format)
}

// now -> format string
func Now(format string) string {
	return time.Now().Format(format)
}

// format: "20060102" / "20060102 03:04:05" / "03:04:05"
func AddTime(date string, add int, unit byte, format string) string {
	t := TimeFromStr(date, format)
	t_ := time.Time{}
	switch unit {
	case 'S':
		t_ = t.Add(time.Second * time.Duration(add))
	case 'M':
		t_ = t.Add(time.Minute * time.Duration(add))
	case 'H':
		t_ = t.Add(time.Hour * time.Duration(add))
	}
	return t_.Format(format)
}

// Candle Datetime 연산
func addTimeInCandle(datetime string, add int, unit byte) string {
	format := FindTimeFormat(datetime)

	t := TimeFromStr(datetime, format)
	t_ := time.Time{}

	switch unit {
	case 'S':
		t_ = t.Add(time.Second * time.Duration(add))
	case 'M':
		t_ = t.Add(time.Minute * time.Duration(add))
	case 'H':
		t_ = t.Add(time.Hour * time.Duration(add))
	case 'd':
		t_ = t.AddDate(0, 0, add)
	case 'm':
		t_ = t.AddDate(0, add, 0)
	case 'y':
		t_ = t.AddDate(add, 0, 0)
	}
	return t_.Format(format)
}

// Candle Datetime 연산
// interval: 30_S, 3_M, ...
func AddTimeInCandle(datetime, interval string) string {
	intvs := strings.Split(interval, "_")
	return addTimeInCandle(datetime, IntFromStr(intvs[0]), intvs[1][0])
}

// Duration between t1 - t0
func SubTime(t1, t0 string) time.Duration {
	format := FindTimeFormat(t1)
	return TimeFromStr(t1, format).Sub(TimeFromStr(t0, format))
}

func intervalFromDuration(duration time.Duration) (interval string) {
	ds := fmt.Sprintf("%v", duration)

	hour := 0
	min := 0
	sec := 0
	if strings.Contains(ds, "h") {
		hms := strings.Split(ds, "h")
		hour = IntFromStr(hms[0])
		ds = hms[1]
	}
	if strings.Contains(ds, "m") {
		hms := strings.Split(ds, "m")
		min = IntFromStr(hms[0])
		ds = hms[1]
	}
	if strings.Contains(ds, "s") {
		sec = IntFromStr(strings.Trim(ds, "s"))
	}

	if sec != 0 { // sec 단위로
		interval = fmt.Sprintf("%v_S", sec+60*min+3600*hour)
	} else if min != 0 { // min 단위로
		interval = fmt.Sprintf("%v_M", min+60*hour)
	} else if hour < 23 { // hour 단위로
		interval = fmt.Sprintf("%v_H", hour)
	} else if hour%24 == 0 { // day 단위로
		interval = fmt.Sprintf("%v_d", hour/24)
	} else {
		interval = fmt.Sprintf("%v_H", hour)
	}
	return
}

// interval string from time diff(t1 - t0)
func IntervalFromTimes(t1, t0 string) (interval string) {
	return intervalFromDuration(SubTime(t1, t0))
}

// 이후시간 여부 t1 > t0이면 true
func IsAfter(t1, t0 string) bool {
	format := FindTimeFormat(t1)
	return TimeFromStr(t1, format).After(TimeFromStr(t0, format))
}
