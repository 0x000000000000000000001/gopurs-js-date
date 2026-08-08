package Data_JSDate

import (
	"gopurs/output/gopurs_runtime"
)

var Now = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var IsValid = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var ToInstantImpl = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(date gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Apply(just, gopurs_runtime.Box(0.0))
		})
	})
})

var FromInstant = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var Jsdate = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})

var JsdateLocal = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{}
	})
})

var DateMethod = gopurs_runtime.Func(func(method gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(date gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{}
	})
})

var DateMethodEff = gopurs_runtime.Func(func(method gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(date gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Value{}
		})
	})
})

var Parse = gopurs_runtime.Func(func(dateString gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Value{}
	})
})

var FromTime = gopurs_runtime.Func(func(time gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Value{}
})
