package logger

import (
	"errors"
	"log"
	"runtime"
	"strings"
)

func Logger(err error, args ...interface{}) {
	programCounter, sourceFileName, sourceFileLineNum, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(programCounter)
	fnames := strings.Split(fn.Name(), "/")
	fname := fnames[len(fnames)-1]
	log.Printf("%s(%d)[%s]: %s - %s", sourceFileName, sourceFileLineNum, fname, err, args)

}
func WriteErr(error string) error {
	log.Println(error)
	return errors.New(error)
}
