package utils

func ScoreToPoints(score int) float64 {
	switch {
	case score >= 90:
		return 4.0
	case score >= 80:
		return 3.0
	case score >= 70:
		return 2.0
	case score >= 60:
		return 1.0
	default:
		return 0.0
	}
}

func CalculateGPA(scores []int) float64 {
	if len(scores) == 0 {
		return 0.0
	}

	var total float64
	for _, s := range scores {
		total += ScoreToPoints(s)
	}
	return total / float64(len(scores))
}
