package main

import (
	"fmt"
)

func goRoutineStop() {
	stopCh := make(chan bool)
	go func() {
		<-stopCh
		fmt.Println("stopped by chan")
		return
	}()

}
