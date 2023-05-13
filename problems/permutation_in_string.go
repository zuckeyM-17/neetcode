package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s1, s2 string
	fmt.Fscan(in, &s1)
	fmt.Fscan(in, &s2)

	fmt.Fprintln(out, checkInclusion(s1, s2))
}
func checkInclusion(s1, s2 string) bool {
	freq1, freq2 := make([]int, 26), make([]int, 26)
	for i := range s1 {
		freq1[s1[i]-'a']++
		freq2[s2[i]-'a']++
	}

	matches := 0
	for i := 0; i < 26; i++ {
		if freq1[i] == freq2[i] {
			matches++
		}
	}

	l := 0

	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}
		idx := s2[r] - 'a'
		freq2[idx]++
		if freq2[idx] == freq1[idx] {
			matches++
		} else if freq2[idx] == freq1[idx]+1 {
			matches--
		}

		idx = s2[l] - 'a'
		freq2[idx]--
		if freq2[idx] == freq1[idx] {
			matches++
		} else if freq2[idx] == freq1[idx]-1 {
			matches--
		}
		l++
	}
	return matches == 26
}
