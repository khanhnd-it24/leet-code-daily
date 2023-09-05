package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	res := make([]int, 26)
	res2 := make([]int, 26)

	for _, v := range ransomNote {
		res[v-'a']++
	}
	for _, c := range magazine {
		res2[c-'a']++
	}
	for i := 0; i < len(res); i++ {
		if res[i] > res2[i] {
			return false
		}
	}
	return true
}
