package main

import (
	"fmt"
	"time"
)

func MySleepFunc(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Спим 5 секунд")

	sleepDuration := 5
	MySleepFunc(sleepDuration)

	fmt.Println("Проспались")
}
