package helpers

import (
	"log"
	"runtime"
	"time"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func LogError(err error) error {
	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, filename, line, _ := runtime.Caller(1)

		log.Printf("Error in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
	}
	return err
}
