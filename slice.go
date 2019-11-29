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

func IndexOfString(strs []string, s string) int {
	for i, str := range strs {
		if str == s {
			return i
		}
	}

	return -1
}
