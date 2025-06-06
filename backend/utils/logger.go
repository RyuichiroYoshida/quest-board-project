package utils

import (
	"fmt"
	"log"
)

func LogInfo(message string, opt ...any) {
	log.Printf("[INFO]: %s", message+" "+fmt.Sprint(opt...))
}

func LogError(message string, opt ...any) {
	log.Printf("[ERROR]: %s", message+" "+fmt.Sprint(opt...))
}

func LogWarning(message string, opt ...any) {
	log.Printf("[WARNING]: %s", message+" "+fmt.Sprint(opt...))
}

func LogDebug(message string, opt ...any) {
	log.Printf("[DEBUG]: %s", message+" "+fmt.Sprint(opt...))
}
