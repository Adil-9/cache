package main

import (
	"cache/funcs"
	"fmt"
	"strconv"
	"time"
)

// using jsonplaceholder api and REDIS caching instances will be shown
func main() {

	funcs.RedisClientInit()

	for {
		fmt.Println(`Chose id to retrieve user data or enter "all"`)
		var scanned string
		fmt.Scanln(&scanned)
		if scanned == "all" {
			counter := time.Now()
			data, err := funcs.RetrieveAll()
			if err != nil {
				continue
			}
			fmt.Println(data)
			elapsedTime := time.Since(counter)
			fmt.Printf("Time spent on request: %v", elapsedTime)
		} else if scanned == "quit" || scanned == "exit" {
			fmt.Println("Bye!!!")
			break
		} else {
			num, err := strconv.Atoi(scanned)
			if err != nil {
				fmt.Println("The input is incorrect!")
				fmt.Println("Enter valid argument!")
				fmt.Println(`To quit print "quit" or "exit"`)
				continue
			}
			counter := time.Now()
			data, err := funcs.RetrieveWithId(num)
			if err != nil {
				continue
			}
			fmt.Println(data)
			elapsedTime := time.Since(counter)
			fmt.Printf("Time spent on request: %v", elapsedTime)
		}
	}
}
