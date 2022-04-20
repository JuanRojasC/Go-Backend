package gotesting

func SortSlice(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if s[j] > s[i] {
				s[j], s[i] = s[i], s[j]
			}
		}
	}
}
