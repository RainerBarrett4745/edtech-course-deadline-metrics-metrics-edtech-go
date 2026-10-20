package main

import (
	"fmt"
	"time"
)

func main() {
	c, err := NewClient()
	if err != nil {
		panic(err)
	}
	now := time.Now().UTC()
	learners := []Learner{{ID: "learner-17", Deadline: now.Add(48 * time.Hour)}, {ID: "learner-42", Deadline: now.Add(-2 * time.Hour)}}
	counts, err := reportLearners(c, learners, now)
	if err != nil {
		panic(err)
	}
	usage, err := c.UsageSeries()
	if err != nil {
		panic(err)
	}
	fmt.Printf("deadline metrics: %+v\naccount usage series: %s\n", counts, usage)
}
