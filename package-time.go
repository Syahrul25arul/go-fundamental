package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Day())
	fmt.Println(now.Month())
	fmt.Println(now.Year())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	utc := time.Date(1996, 2, 25, 5, 30, 25, 22, time.UTC)
	fmt.Println(utc)

	layout := "02-01-2006"
	parse, _ := time.Parse(layout, "25-02-1996")
	fmt.Println(parse.Date())
	times := time.Now()
	fmt.Println(times)
}
