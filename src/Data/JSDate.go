package Data_JSDate

import (
	"gopurs/output/gopurs_runtime"
	"math"
	"time"
	"strings"
)

type Date struct {
	Ms float64
}

func getMs(date gopurs_runtime.Value) float64 {
	if date.Type == gopurs_runtime.TypeAny && date.UnsafePtr != nil {
		val := *(*any)(date.UnsafePtr)
		if d, ok := val.(Date); ok {
			return d.Ms
		}
	}
	return date.FloatVal()
}

func boxDate(ms float64) gopurs_runtime.Value {
	return gopurs_runtime.Box(Date{Ms: ms})
}

func isValidMs(ms float64) bool {
	return !math.IsNaN(ms)
}

var Now = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return boxDate(float64(time.Now().UnixMilli()))
	})
})

var IsValid = gopurs_runtime.Func(func(date gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Box(isValidMs(getMs(date)))
})

var ToInstantImpl = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(date gopurs_runtime.Value) gopurs_runtime.Value {
			ms := getMs(date)
			if !isValidMs(ms) {
				return nothing
			}
			return gopurs_runtime.Apply(just, gopurs_runtime.Box(ms))
		})
	})
})

var FromInstant = gopurs_runtime.Func(func(instant gopurs_runtime.Value) gopurs_runtime.Value {
	return boxDate(instant.FloatVal())
})

func msFromParts(year, month, day, hour, minute, second, millisecond float64, local bool) float64 {
	if math.IsNaN(year) || math.IsNaN(month) || math.IsNaN(day) || math.IsNaN(hour) || math.IsNaN(minute) || math.IsNaN(second) || math.IsNaN(millisecond) {
		return math.NaN()
	}
	loc := time.UTC
	if local {
		loc = time.Local
	}
	t := time.Date(int(year), time.Month(int(month)+1), int(day), int(hour), int(minute), int(second), int(millisecond)*1e6, loc)
	return float64(t.UnixMilli())
}

var Jsdate = gopurs_runtime.Func(func(rec gopurs_runtime.Value) gopurs_runtime.Value {
	return boxDate(msFromParts(
		gopurs_runtime.RecordGet(rec, "year").FloatVal(),
		gopurs_runtime.RecordGet(rec, "month").FloatVal(),
		gopurs_runtime.RecordGet(rec, "day").FloatVal(),
		gopurs_runtime.RecordGet(rec, "hour").FloatVal(),
		gopurs_runtime.RecordGet(rec, "minute").FloatVal(),
		gopurs_runtime.RecordGet(rec, "second").FloatVal(),
		gopurs_runtime.RecordGet(rec, "millisecond").FloatVal(),
		false,
	))
})

var JsdateLocal = gopurs_runtime.Func(func(rec gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return boxDate(msFromParts(
			gopurs_runtime.RecordGet(rec, "year").FloatVal(),
			gopurs_runtime.RecordGet(rec, "month").FloatVal(),
			gopurs_runtime.RecordGet(rec, "day").FloatVal(),
			gopurs_runtime.RecordGet(rec, "hour").FloatVal(),
			gopurs_runtime.RecordGet(rec, "minute").FloatVal(),
			gopurs_runtime.RecordGet(rec, "second").FloatVal(),
			gopurs_runtime.RecordGet(rec, "millisecond").FloatVal(),
			true,
		))
	})
})

func handleMethod(method string, date gopurs_runtime.Value) gopurs_runtime.Value {
	ms := getMs(date)
	if !isValidMs(ms) {
		if method == "toString" || method == "toDateString" || method == "toTimeString" || method == "toUTCString" || method == "toISOString" {
			return gopurs_runtime.Str("Invalid Date")
		}
		return gopurs_runtime.Float(math.NaN())
	}

	sec := int64(ms / 1000)
	nsec := int64(math.Mod(ms, 1000)) * 1e6
	t := time.Unix(sec, nsec)

	switch method {
	case "getTime": return gopurs_runtime.Float(ms)
	case "getUTCDate": return gopurs_runtime.Float(float64(t.UTC().Day()))
	case "getUTCDay": return gopurs_runtime.Float(float64(t.UTC().Weekday()))
	case "getUTCFullYear": return gopurs_runtime.Float(float64(t.UTC().Year()))
	case "getUTCHours": return gopurs_runtime.Float(float64(t.UTC().Hour()))
	case "getUTCMilliseconds": return gopurs_runtime.Float(float64(t.UTC().Nanosecond() / 1e6))
	case "getUTCMinutes": return gopurs_runtime.Float(float64(t.UTC().Minute()))
	case "getUTCMonth": return gopurs_runtime.Float(float64(t.UTC().Month() - 1))
	case "getUTCSeconds": return gopurs_runtime.Float(float64(t.UTC().Second()))
	
	case "getDate": return gopurs_runtime.Float(float64(t.Local().Day()))
	case "getDay": return gopurs_runtime.Float(float64(t.Local().Weekday()))
	case "getFullYear": return gopurs_runtime.Float(float64(t.Local().Year()))
	case "getHours": return gopurs_runtime.Float(float64(t.Local().Hour()))
	case "getMilliseconds": return gopurs_runtime.Float(float64(t.Local().Nanosecond() / 1e6))
	case "getMinutes": return gopurs_runtime.Float(float64(t.Local().Minute()))
	case "getMonth": return gopurs_runtime.Float(float64(t.Local().Month() - 1))
	case "getSeconds": return gopurs_runtime.Float(float64(t.Local().Second()))
	case "getTimezoneOffset":
		_, offset := t.Local().Zone()
		return gopurs_runtime.Float(float64(-offset / 60))
	
	case "toDateString": return gopurs_runtime.Str(t.Local().Format("Mon Jan 02 2006"))
	case "toString": return gopurs_runtime.Str(t.Local().Format("Mon Jan 02 2006 15:04:05 GMT-0700 (MST)"))
	case "toTimeString": return gopurs_runtime.Str(t.Local().Format("15:04:05 GMT-0700 (MST)"))
	case "toUTCString": return gopurs_runtime.Str(t.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
	case "toISOString": return gopurs_runtime.Str(t.UTC().Format("2006-01-02T15:04:05.000Z"))
	}
	return gopurs_runtime.Value{}
}

var DateMethod = gopurs_runtime.Func2(func(method, date gopurs_runtime.Value) gopurs_runtime.Value {
	return handleMethod(method.StrVal(), date)
})

var DateMethodEff = gopurs_runtime.Func2(func(method, date gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return handleMethod(method.StrVal(), date)
	})
})

var Parse = gopurs_runtime.Func(func(dateString gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		str := dateString.StrVal()
		t, err := time.Parse(time.RFC3339, str)
		if err == nil {
			return boxDate(float64(t.UnixMilli()))
		}

		t, err = time.Parse("2006-01-02T15:04:05.000Z", str)
		if err == nil {
			return boxDate(float64(t.UnixMilli()))
		}
		
		t, err = time.Parse("2006-01-02T15:04:05", str)
		if err == nil {
			return boxDate(float64(t.UnixMilli()))
		}
		
		str = strings.Replace(str, "GMT", "", 1)
		t, err = time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", str)
		if err == nil {
			return boxDate(float64(t.UnixMilli()))
		}
		
		return boxDate(math.NaN())
	})
})

var FromTime = gopurs_runtime.Func(func(timeVal gopurs_runtime.Value) gopurs_runtime.Value {
	return boxDate(timeVal.FloatVal())
})
