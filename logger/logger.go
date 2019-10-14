package logger

import (
	"fmt"
	"runtime"
	"strings"
)

func Logger(err error, args ...interface{}) string {
	programCounter, sourceFileName, sourceFileLineNum, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(programCounter)
	fnames := strings.Split(fn.Name(), "/")
	fname := fnames[len(fnames)-1]
	fmtMsg := fmt.Sprintf("%s(%d)[%s]: %s - %s", sourceFileName, sourceFileLineNum, fname, err, args)
	return fmtMsg
}
