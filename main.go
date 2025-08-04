package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const str = "0.beevik-ntp.pool.ntp.org"

func GetNTPTime(str string) (time.Time, error) {
	return ntp.Time(str)
}

func main() {
	currTime, err := GetNTPTime(str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting ntp time: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Current Time: %v\n", currTime.Format(time.RFC3339))
}
