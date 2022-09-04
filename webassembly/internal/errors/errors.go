package errors

import "log"

func Check(e error, msg string) {
	if e != nil {
		log.Println(e)
		log.Panic(msg)
	}
}
