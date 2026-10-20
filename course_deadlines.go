package main

import "time"

type Learner struct {
	ID          string
	Deadline    time.Time
	SubmittedAt *time.Time
}

func completionStatus(l Learner, now time.Time) string {
	if l.SubmittedAt != nil {
		return "submitted"
	}
	if now.After(l.Deadline) {
		return "overdue"
	}
	return "open"
}

func reportLearners(c *Client, learners []Learner, now time.Time) (map[string]int, error) {
	counts := map[string]int{"submitted": 0, "overdue": 0, "open": 0}
	for _, learner := range learners {
		status := completionStatus(learner, now)
		counts[status]++
	}
	for status, count := range counts {
		if err := c.Report("course.deadlines."+status, float64(count), "gauge", map[string]string{"course": "go-foundations"}); err != nil {
			return nil, err
		}
	}
	return counts, nil
}
