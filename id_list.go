package gox

type KeyIDList map[ID][]ID

func (kl KeyIDList) Append(key ID, val ID) {
	l, ok := kl[key]
	if !ok {
		l = make([]ID, 0, 1)
	}
	kl[key] = append(l, val)
}

type IDList []ID

func (l IDList) Len() int {
	return len(l)
}

func (l IDList) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l IDList) Swap(i, j int) {
	tmp := l[i]
	l[i] = l[j]
	l[j] = tmp
}

func (l *IDList) Append(i ID) {
	*l = append(*l, i)
}

func (l *IDList) Remove(i ID) bool {
	a := *l
	for idx, v := range a {
		if v == i {
			*l = append(a[0:idx], a[idx+1:]...)
			return true
		}
	}
	return false
}

func (l IDList) IndexOf(i ID) int {
	for idx, v := range l {
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
	for id := range ids {
		a = append(a, id)
	}
	return a
}
