package main

import (
	"./model"
	"fmt"
	"os"
	"strconv"
)

//Two modes time and json
func main() {
	args := os.Args

	if args[1] != "time" && args[1] != "json" {
		fmt.Println("Only time and json modes allowed")
		return
	}

	if args[1] == "time" {

		size, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error converting size to int", err)
		}

		timeMode(size)
	} else {
		size, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error converting size to int", err)
		}
		minutes, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Println("Error converting minutes to int", err)
		}
		jsonMode(size, minutes)
	}

}
func jsonMode(size int, minutes int) {
	c, err := model.NewClock(size)
	if err != nil {
		fmt.Println("NewClock fails %s", err)
		return
	}
	i := 0
	for ; i < minutes; i++ {
		_, err := c.MoveArm()
		if err != nil {
			fmt.Println("Error: %s", err)
			return
		}
		if i%(24*60) == 0 {
			fmt.Print(".")
		}
	}
	//TODO write using json.marshall
	fmt.Println()
	fmt.Println(c.String())
}

func timeMode(size int) {
	fmt.Println("Running in time mode with", size, "balls.")
	c, err := model.NewClock(size)
	if err != nil {
		fmt.Println("NewClock fails: %s", err)
		return
	}
	counter := 1
	for {
		b, err := c.MoveArm()
		if err != nil {
			fmt.Println("Error: %s", err)
			return
		}
		if b {
			break
		}
		counter++
		if counter%(24*60) == 0 {
			fmt.Print(".")
		}
	}
	fmt.Println()
	fmt.Println("Days until back in order: ", (counter / (24 * 60)))
}
