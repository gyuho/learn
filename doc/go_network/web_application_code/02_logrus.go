package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
)

func openToAppend(fpath string) (*os.File, error) {
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		f, err = os.Create(fpath)
		if err != nil {
			return f, err
		}
	}
	return f, nil
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(new(log.JSONFormatter))

	// https://godoc.org/github.com/Sirupsen/logrus#Level
	// log.SetLevel(log.PanicLevel)
	// log.SetLevel(log.FatalLevel)
	// log.SetLevel(log.ErrorLevel)
	// log.SetLevel(log.WarnLevel)
	// log.SetLevel(log.InfoLevel)
	log.SetLevel(log.DebugLevel)
}

func main() {
	lf, err := openToAppend("my.log")
	if err != nil {
		panic(err)
	}
	defer lf.Close()
	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(lf)

	log.Println("Hello World!") // Println belongs to Infoln
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.Panic("panic")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Panic("The ice breaks!")
}
