package main

func characterReplacement(s string, k int) int {
	start := 0
	m := make(map[byte]int)
	var pop byte
	answer := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		m[c]++
		if pop == 0 || m[pop] < m[c] {
			pop = c
		}

		if i-start+1 > m[pop]+k {
			m[s[start]]--
			start++
		}

		answer = max(i-start+1, answer)
	}

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
