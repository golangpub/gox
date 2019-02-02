package types

type IDList []ID

func (ids IDList) Len() int {
	return len(ids)
}

func (ids IDList) Less(i, j int) bool {
	return ids[i] < ids[j]
}

func (ids IDList) Swap(i, j int) {
	tmp := ids[i]
	ids[i] = ids[j]
	ids[j] = tmp
}

func (p *IDList) Append(i ID) {
	*p = append(*p, i)
}

func (p *IDList) Remove(i ID) bool {
	a := *p
	for idx, v := range a {
		if v == i {
			if idx < len(a)-1 {
				*p = append(a[0:idx], a[idx+1:len(a)]...)
			} else {
				*p = a[0:idx]
			}
			return true
		}
	}

	return false
}

func (a IDList) IndexOf(i ID) int {
	for idx, v := range a {
		if v == i {
			return idx
		}
	}
	return -1
}

func Int64sToIDs(a []int64) []ID {
	if a == nil {
		return nil
	}

	ids := make([]ID, len(a))
	for i, v := range a {
		ids[i] = ID(v)
	}

	return ids
}

func IDsToInt64s(ids []ID) []int64 {
	if ids == nil {
		return nil
	}

	a := make([]int64, len(ids))
	for i, v := range ids {
		a[i] = int64(v)
	}

	return a
}

func IDSliceToSet(ids []ID) map[ID]bool {
	m := make(map[ID]bool, len(ids))
	for _, v := range ids {
		m[v] = true
	}

	return m
}

func IDSetToSlice(ids map[ID]bool) []ID {
	a := make([]ID, 0, len(ids))
	for id, _ := range ids {
		a = append(a, id)
	}
	return a
}
