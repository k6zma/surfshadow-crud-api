package utils

import (
	"fmt"
	"time"
)

const (
	interval = 3 * time.Millisecond
)

func LogoPrint(text string) {
	for _, v := range text {
		fmt.Printf("%c", v)
		time.Sleep(interval)
	}
}
