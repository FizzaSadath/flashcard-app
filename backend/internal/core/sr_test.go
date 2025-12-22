package core

import "testing"

func TestCalculateReview(t *testing.T) {
	currentStats := CardStats{
		Repetitions: 0,
		EaseFactor:  2.5,
		Interval:    0,
	}
	grade := 5

	newStats := CalculateReview(currentStats, grade)

	if newStats.Repetitions != 1 {
		t.Errorf("Expected Repetitions 1, but got %d", newStats.Repetitions)
	}
	if newStats.Interval != 1 {
		t.Errorf("Expected Interval 1, but got %d", newStats.Interval)
	}
}
