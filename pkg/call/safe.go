package call

import (
	"runtime/debug"
)

func RunSafe(fn func()) {
	defer Recover()
	fn()
}

// RunSafeWrapper 捕获记录panic
func RunSafeWrapper(fn func() error) func() error {
	return func() (err error) {
		defer Recover()
		err = fn()
		return
	}
}

//Recover 捕获打印panic
//用法：defer Recover()
func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}
	if p := recover(); p != nil {
		//err := errorx.AsError(p)
		debug.PrintStack() //控制台输出
	}
}
