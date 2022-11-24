package debugx

import (
	"fmt"
	"runtime"
	"strings"
)

//GetStackTrace get stack trace
func GetStackTrace(skip int) string {
	const skipOffset = 2 //默认跳过两层
	var stackTraces []string
	for i := skipOffset + skip; ; i++ {
		pc, f, l, got := runtime.Caller(i)
		if !got {
			break
		}
		pcName := runtime.FuncForPC(pc).Name()
		stackTraces = append(stackTraces, fmt.Sprintf("%s\n\t%s:%d", pcName, f, l))
	}
	return strings.Join(stackTraces, "\n")
}

//LocatePanic 定位发生panic的代码行，需直接写在defer_recover方法中
//不能用于包装的公共方法中，会定位错误
func LocatePanic() string {
	var name, file string
	var line int
	var pc [16]uintptr

	n := runtime.Callers(3, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line = fn.FileLine(pc)
		name = fn.Name()
		if !strings.HasPrefix(name, "runtime.") {
			break
		}
	}

	if name != "" && file != "" && line != 0 {
		return fmt.Sprintf("%s:%d\t%s", file, line, name)

	}

	return fmt.Sprintf("pc:%x", pc)
}
