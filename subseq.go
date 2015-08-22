// public domain.

package subseq

import "sort"

func LongestIncreasing(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	var m []int
	p := make([]int, len(a))
	for i, xi := range a {
		j := sort.Search(len(m), func(j int) bool { return a[m[j]] > xi })
		if j > 0 {
			p[i] = m[j-1]
		}
		if j == len(m) {
			m = append(m, i)
		} else if xi < a[m[j]] {
			m[j] = i
		}
	}
	s := make([]int, len(m))
	last := len(m) - 1
	k := m[last]
	for i := last; i >= 0; i-- {
		s[i] = a[k]
		k = p[k]
	}
	return s
}

func LongestDecreasing(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	var m []int
	p := make([]int, len(a))
	for i, xi := range a {
		j := sort.Search(len(m), func(j int) bool { return a[m[j]] < xi })
		if j > 0 {
			p[i] = m[j-1]
		}
		if j == len(m) {
			m = append(m, i)
		} else if xi > a[m[j]] {
			m[j] = i
		}
	}
	s := make([]int, len(m))
	last := len(m) - 1
	k := m[last]
	for i := last; i >= 0; i-- {
		s[i] = a[k]
		k = p[k]
	}
	return s
}

func LongestCommon(a, b []int) []int {
	m := make([][]int, len(a)+1)
	for i := range m {
		m[i] = make([]int, len(b)+1)
	}
	for i, ai := range a {
		for j, bj := range b {
			switch {
			case ai == bj:
				m[i+1][j+1] = m[i][j] + 1
			case m[i+1][j] > m[i][j+1]:
				m[i+1][j+1] = m[i+1][j]
			default:
				m[i+1][j+1] = m[i][j+1]
			}
		}
	}
	c := make([]int, m[len(a)][len(b)])
	cx := len(c)
	for x, y := len(a), len(b); x != 0 && y != 0; {
		switch {
		case m[x][y] == m[x-1][y]:
			x--
		case m[x][y] == m[x][y-1]:
			y--
		default:
			x--
			y--
			cx--
			c[cx] = a[x]
		}
	}
	return c
}
