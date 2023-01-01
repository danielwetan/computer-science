package challenges

func personalLearningScore(rule, log []int) int {

	// L: menit terakhir
	// M: batas dapat thropy
	// N: total menit
	L, M := rule[0], rule[1]

	var score, trophies int

	for i := 0; i < L; i++ {
		score += log[i]

		if score >= M {
			trophies++
		}
	}

	return trophies
}
