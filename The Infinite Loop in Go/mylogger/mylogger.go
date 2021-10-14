package mylogger

import "log"

func ListenForLog(ch chan string) {
	for {
		msg := <-ch

		log.Println(msg)
	}
}
