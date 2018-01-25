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
