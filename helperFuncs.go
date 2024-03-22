package main

import (
	"fmt"
)

func WriteLog(text ...string) {
	fmt.Println(fmt.Sprint(text), "log/monitor")
}
