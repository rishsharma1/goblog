package blog

import (
	"log"
)

// LogError logs an error
func LogError(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}
