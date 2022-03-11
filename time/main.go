package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	reqTime := "1746838781"

	testTime, _ := strconv.ParseInt(reqTime, 10, 64)
	myTime := time.Unix(testTime, 0)

	fmt.Println("My = ", myTime.Unix())

	diff := time.Since(myTime).Seconds()

	fmt.Println("Diff = ", diff)
}
