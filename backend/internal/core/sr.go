package core

func CalculateReview(stats CardStats, grade int) CardStats {
	newStats := stats

	newStats.Repetitions += 1
	if newStats.Repetitions == 1 {
		newStats.Interval = 1
	}

	return newStats
}
