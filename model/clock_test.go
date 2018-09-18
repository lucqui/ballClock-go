package model

import (
	"testing"
	"fmt"
)

func TestClock(t *testing.T){
	c,err := NewClock(30)
	if err != nil {
		t.Fatalf("NewClock failes %s",err)
	}
	counter := 1
	for  {
		b,err := c.MoveArm()
		if err != nil {
			t.Fatalf("Error: %s",err)
		}
		if b {
			break
		}
		counter ++
		if counter%(24*60) == 0{
			fmt.Println((counter/(24*60)))
		}
		if (counter/(24*60)) > 30 {
			t.Fatalf("Counter passed 15 days")
		}
	}
	fmt.Println(counter)
	fmt.Println("Days: ", (counter/(24*60)))
	if (counter/(24*60)) != 15 {
		t.Fatalf("Error not expect 15, got: %d",(counter/(24*60)))
	}
}

