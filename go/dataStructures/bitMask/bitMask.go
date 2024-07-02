package bitmask

type bitmask struct {
	mask int64
}

func (bm bitmask) Set(i, val int) {
	if val == 1 {
		bm.mask |= 1 << i
	} else if val == 0 {
		bm.mask &= ^(1 << i)
	}
}

func (bm bitmask) IsBit(i int) bool {
	return bm.mask&(1<<i) != 0
}

func (bm bitmask) Invert(i int) {
	bm.mask ^= 1 << i
}

func (bm bitmask) Count() int {
	c := bm.mask
	count := 0
	for c != 0 {
		c &= c - 1
		count++
	}
	return count
}
