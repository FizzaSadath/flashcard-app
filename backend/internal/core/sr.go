package core

func CalculateReview(stats CardStats, grade int) CardStats {
	newStats := stats
	newStats.Repetitions += 1

	if newStats.Repetitions == 1 {
		newStats.Interval = 1
	} else if newStats.Repetitions == 2 {
		newStats.Interval = 6
	}

	return newStats
}
