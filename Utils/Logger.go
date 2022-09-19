package Utils

import (
	"log"
)

type Log4g struct {
}

func (l *Log4g) Info(prefix string, msg interface{}) {
	log.SetPrefix(INFO_PRX)
	log.Printf("%s:%#v\n", prefix, msg)
}

func (l *Log4g) Error(prefix string, msg interface{}) {
	log.SetPrefix(ERROR_PRX)
	log.Printf("%s:%#v\n", prefix, msg)
	log.SetPrefix(INFO_PRX)
}

func (l *Log4g) Warn(prefix string, msg interface{}) {
	log.SetPrefix(WARN_PRX)
	log.Printf("%s:%#v\n", prefix, msg)
	log.SetPrefix(INFO_PRX)
}
