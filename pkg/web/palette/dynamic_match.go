package palette

func dynamicMatch(
	pattern []rune,
	text []rune,
	bonuses []int,
) (int, []int) {
	rows := len(pattern)
	columns := len(text)
	scores := make([][]int, rows)
	consecutive := make([][]int, rows)

	for i := range rows {
		scores[i] = make([]int, columns)
		consecutive[i] = make([]int, columns)

		for j := range columns {
			scores[i][j] = minScore
		}
	}

	for j := range columns {
		if text[j] != pattern[0] {
			continue
		}

		bonus := bonuses[j] * bonusFirstCharacter
		scores[0][j] = scoreMatch + bonus
		consecutive[0][j] = 1
	}

	for i := 1; i < rows; i++ {
		for j := i; j < columns; j++ {
			if text[j] != pattern[i] {
				continue
			}

			bonus := bonuses[j]
			diagonalScore := minScore

			if j > 0 && scores[i-1][j-1] > minScore {
				past := consecutive[i-1][j-1]
				consecutiveBonus := max(bonus, bonusConsecutiveMin)

				if past > 0 {
					bonus = consecutiveBonus
				}

				diagonalScore = scores[i-1][j-1] + scoreMatch + bonus
			} else if i == 0 {
				diagonalScore = scoreMatch + bonus*bonusFirstCharacter
			}

			gapScore := minScore

			for k := i - 1; k < j; k++ {
				if scores[i-1][k] <= minScore {
					continue
				}

				gap := j - k - 1
				penalty := scoreGapStart + scoreGapExtension*gap
				candidate := scores[i-1][k] + penalty + scoreMatch + bonus

				if candidate > gapScore {
					gapScore = candidate
				}
			}

			best := max(diagonalScore, gapScore)

			if best > minScore {
				scores[i][j] = best

				if diagonalScore >= gapScore && j > 0 {
					consecutive[i][j] = consecutive[i-1][j-1] + 1
				} else {
					consecutive[i][j] = 1
				}
			}
		}
	}

	bestScore := minScore
	bestEnd := -1

	for j := rows - 1; j < columns; j++ {
		if scores[rows-1][j] > bestScore {
			bestScore = scores[rows-1][j]
			bestEnd = j
		}
	}

	if bestEnd < 0 {
		return -1, nil
	}

	positions := backtrack(scores, consecutive, pattern, text, bestEnd)

	return bestScore, positions
}
