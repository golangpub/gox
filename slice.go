package gox

func IndexOfInt(a []int, i int) int {
	for idx, v := range a {
		if v == i {
			return idx
		}
	}

	return -1
}

func IndexOfInt64(a []int64, i int64) int {
	for idx, v := range a {
		if v == i {
			return idx
		}
	}

	return -1
}

func IndexOfString(a []string, s string) int {
	for i, str := range a {
		if str == s {
			return i
		}
	}

	return -1
}

func ReverseIntSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ReverseInt64Slice(s []int64) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
