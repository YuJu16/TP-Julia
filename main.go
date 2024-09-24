package main

import "fmt"

func Ft_missing(nums []int) int {
	n := len(nums)
	m := 67
	for i := 0; i < n; i++ {
		for range nums {
			if i != nums[i] {
				m = i
			}
		}
	}
	if m == 67 {
		m = n
	}
	return m
}

func Ft_profit(prices []int) int {
	n := len(prices)
	minP := n + 1
	maxP := 0
	day := 0
	for i := 0; i < n; i++ {
		for range prices {
			if prices[i] <= minP {
				minP = prices[i]
				day = i
			}
		}
	}
	for j := day; j < n; j++ {
		for range prices {
			if prices[j] >= maxP {
				maxP = prices[j]
			}
		}
	}
	return maxP - minP
}

func Ft_max_substring(s string) int {
	var u []string
	var c string
	d := false
	for i := 0; i < len(s); i++ {
		c = string(s[i])
		for range s {
			for j := 0; j < len(u); j++ {
				for range u {
					if c == u[j] {
						d = true
					}
				}
			}
			if d == false {
				u = append(u, c)
			}

		}
		d = false
	}
	return len(u)
}

func Ft_coin(coins []int, amt int) int {
	var dp []int
	for j := 0; j < amt+1; j++ {
		dp = append(dp, j)
	}
	for i := 1; i <= amt; i++ {
		dp[i] = amt + 1
	}
	for i := 1; i <= amt; i++ {
		for _, c := range coins {
			if i >= c {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amt] == amt+1 {
		return -1
	}
	return dp[amt]
}

func Ft_non_overlap(intervals [][]int) int {
	ov := 0
	e := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < e {
			ov++
		} else {
			e = intervals[i][1]
		}
	}

	return ov
}

func Ft_min_window(s string, t string) string {
	n := len(s)
	minL := n
	res := ""

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sub := s[i : j+1]
			if ca(sub, t) {
				if len(sub) < minL {
					minL = len(sub)
					res = sub
				}
			}
		}
	}

	return res
}

func verif(s string, t string) bool {
	for i := 0; i < len(t); i++ {
		f := false
		for j := 0; j < len(s); j++ {
			if t[i] == s[j] {
				f = true
				break
			}
		}
		if !f {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Ft_missing([]int{3, 1, 2}))                   // resultat : 0
	fmt.Println(Ft_missing([]int{0, 1}))                      // resultat : 2
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // resultat : 8
	fmt.Print("\n")

	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4})) // resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1})) // resultat : 0
	fmt.Print("\n")

	fmt.Println(Ft_max_substring("abcabcbb")) // resultat : 3
	// "abc" est la plus grande sous chaine composé de caractère diffèrent
	fmt.Println(Ft_max_substring("bbbbb")) // resultat : 1
	// "b" est la plus grande sous chaine
	fmt.Print("\n")

	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3 car (11 == 5 + 5 + 1)
	fmt.Println(Ft_coin([]int{2}, 3))        // resultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // resultat : 0
	fmt.Print("\n")

	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // resultat : 1
	// pour que les intervalles soient tous des intervalles qui ne se superpose pas,
	// il suffit de d'enlever qu'un seul intervalle, [1,3]
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))         // resultat : 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}})) // resultat : 2
	fmt.Print("\n")

	fmt.Println(Ft_min_window("ADOBECODEBANC", "ABC")) // resultat : "BANC"
	fmt.Println(Ft_min_window("a", "aa"))              // resultat : ""

}
