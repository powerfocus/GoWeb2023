package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	lg *logrus.Logger
)

type Log struct {
	Log *logrus.Logger
}

func init() {
	lg = logrus.New()
	lg.Formatter = new(logrus.JSONFormatter)
	lg.Formatter = new(logrus.TextFormatter)                     //default
	lg.Formatter.(*logrus.TextFormatter).DisableColors = true    // remove colors
	lg.Formatter.(*logrus.TextFormatter).DisableTimestamp = true // remove timestamp from test output
	lg.Level = logrus.TraceLevel
	lg.Out = os.Stdout

	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	// 	log.Out = file
	// } else {
	// 	log.Info("Failed to log to file, using default stderr")
	// }

	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logrus.Entry)
			lg.WithFields(logrus.Fields{
				"omg":         true,
				"err_animal":  entry.Data["animal"],
				"err_size":    entry.Data["size"],
				"err_level":   entry.Level,
				"err_message": entry.Message,
				"number":      100,
			}).Error("The ice breaks!") // or use Fatal() to force the process to exit with a nonzero code
		}
	}()
}
func NewLog() *Log {
	log := new(Log)
	log.Log = lg
	return log
}
func Println(v ...any) {
	lg.Println(v)
}
