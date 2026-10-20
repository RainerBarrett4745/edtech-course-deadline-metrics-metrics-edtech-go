package main

import (
	"testing"
	"time"
)

func TestCompletionStatusDrivesReportCounts(t *testing.T) {
	now := time.Date(2026, 9, 13, 12, 0, 0, 0, time.UTC)
	submitted := now.Add(-time.Hour)
	learners := []Learner{{ID: "a", Deadline: now.Add(time.Hour)}, {ID: "b", Deadline: now.Add(-time.Hour)}, {ID: "c", Deadline: now.Add(time.Hour), SubmittedAt: &submitted}}
	want := map[string]int{"open": 1, "overdue": 1, "submitted": 1}
	got := map[string]int{"submitted": 0, "overdue": 0, "open": 0}
	for _, l := range learners {
		got[completionStatus(l, now)]++
	}
	for k, v := range want {
		if got[k] != v {
			t.Fatalf("%s: got %d want %d", k, got[k], v)
		}
	}
}
