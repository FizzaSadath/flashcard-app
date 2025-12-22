package core

import "testing"

func TestCalculateReview_InitialCorrectResponse(t *testing.T) {
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

func TestCalculateReview_SecondCorrectResponse(t *testing.T) {
	currentStats := CardStats{Repetitions: 1, EaseFactor: 2.5, Interval: 1}
	grade := 4

	newStats := CalculateReview(currentStats, grade)

	if newStats.Repetitions != 2 {
		t.Errorf("Expected Repetitions 2,but got %d", newStats.Repetitions)
	}
	if newStats.Interval != 6 {
		t.Errorf("Expected Interval 6,but got %d", newStats.Interval)
	}
}
func TestCalculateReview_ThirdCorrectResponse(t *testing.T) {
	currentStats := CardStats{
		Repetitions: 2,
		EaseFactor:  2.5,
		Interval:    6,
	}
	grade := 5

	newStats := CalculateReview(currentStats, grade)

	if newStats.Repetitions != 3 {
		t.Errorf("Expected Repetitions 3,but got %d", newStats.Repetitions)
	}
	if newStats.Interval != 15 {
		t.Errorf("Expected Interval 15,but got %d", newStats.Interval)
	}
	if newStats.EaseFactor == 2.5 {
		t.Errorf("EaseFactor should have updated")
	}
}
func TestCalculateReview_IncorrectResponse(t *testing.T) {
	currentStats := CardStats{Repetitions: 10, EaseFactor: 2.9, Interval: 100}
	grade := 0 // Forgot completely

	newStats := CalculateReview(currentStats, grade)

	if newStats.Repetitions != 0 {
		t.Errorf("Expected Repetitions reset to 0,but got %d", newStats.Repetitions)
	}
	if newStats.Interval != 1 {
		t.Errorf("Expected Interval reset to 1,but got %d", newStats.Interval)
	}
}
