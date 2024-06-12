package longest_special_substring

// Description: https://leetcode.com/problems/find-longest-special-substring-that-occurs-thrice-i/description/

func maximumLength(s string) int {
	type st struct {
		ch  byte
		len int
	}

	mp := make(map[st]int)
	n := len(s)
	var count int
	for i := 0; i < n; i++ {
		count = 1
		mp[st{ch: s[i], len: count}]++
		j := i
		for j < n-1 && s[j] == s[j+1] {
			count++
			mp[st{ch: s[i], len: count}]++
			j++
		}
	}
	maxLen := -1
	for k, v := range mp {
		if v >= 3 && k.len > maxLen {
			maxLen = k.len
		}
	}
	return maxLen
}

// Solution 2:
func maximumLength2(s string) int {
	num := [26][3]int{}
	start := -1
	for i := range s {
		if start == -1 {
			start = i
		} else if s[i] != s[i-1] {
			prev := s[i-1] - 97
			count := i - start
			start = i
			a := num[prev][0]
			b := num[prev][1]
			c := num[prev][2]
			if count >= a {
				num[prev][2] = b
				num[prev][1] = a
				num[prev][0] = count
			} else if count >= b {
				num[prev][2] = b
				num[prev][1] = count
				num[prev][0] = a
			} else if count >= c {
				num[prev][2] = count
			}
		}
		if i == len(s)-1 {
			count := 1
			prev := s[i] - 97
			if s[i] == s[i-1] {
				count = i - start + 1
			}
			a := num[prev][0]
			b := num[prev][1]
			c := num[prev][2]
			if count >= a {
				num[prev][2] = b
				num[prev][1] = a
				num[prev][0] = count
			} else if count >= b {
				num[prev][2] = b
				num[prev][1] = count
				num[prev][0] = a
			} else if count >= c {
				num[prev][2] = count
			}

		}
	}
	res := -1
	for j := range num {
		sum := num[j][0] + num[j][1] + num[j][2]
		if sum < 3 {
			continue
		}
		max := num[j][2]
		if num[j][1] > 0 {
			if num[j][0] > num[j][1] {
				max = num[j][1]
			} else {
				// num[j][1] == num[j][0]
				if num[j][1]-1 > max {
					max = num[j][1] - 1
				}
			}
		}
		if num[j][0] >= 3 && num[j][0]-2 > max {
			max = num[j][0] - 2
		}
		if max > res {
			res = max
		}
	}
	return res
}
