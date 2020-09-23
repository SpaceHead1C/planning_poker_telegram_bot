package logger

import (
	"log"
	"os"

	"../betypes"
)

var (
	outfile, _ = os.OpenFile(betypes.LOGS_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0775)
	LogFile    = log.New(outfile, "", 0)
)

func ForError(er error) {
	if er != nil {
		LogFile.Fatalln(er)
	}
}
