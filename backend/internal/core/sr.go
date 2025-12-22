package core

import "math"

func CalculateReview(stats CardStats, grade int) CardStats {
	newStats := stats

	if grade < 3 {
		newStats.Repetitions = 0
		newStats.Interval = 1
	} else {
		newStats.Repetitions += 1

		if newStats.Repetitions == 1 {
			newStats.Interval = 1
		} else if newStats.Repetitions == 2 {
			newStats.Interval = 6
		} else {
			newInterval := float64(newStats.Interval) * newStats.EaseFactor
			newStats.Interval = int(math.Ceil(newInterval))
		}
	}

	newStats.EaseFactor = stats.EaseFactor + (0.1 - (5.0-float64(grade))*(0.08+(5.0-float64(grade))*0.02))

	if newStats.EaseFactor < 1.3 {
		newStats.EaseFactor = 1.3
	}

	return newStats
}
