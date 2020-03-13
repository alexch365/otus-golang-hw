package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	fmt.Println("current time:", time.Now())
	exactTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatalln("exact time:", err)
	} else {
		fmt.Println("exact time:", exactTime)
	}
}
