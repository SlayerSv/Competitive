package bitmask

type bitmask int64

func (bm bitmask) Set(i, val int) {
	if val == 1 {
		bm |= 1 << i
	} else if val == 0 {
		bm &= ^(1 << i)
	}
}

func (bm bitmask) IsBit(i int) bool {
	return bm&(1<<i) != 0
}

func (bm bitmask) Invert(i int) {
	bm ^= 1 << i
}

func (bm bitmask) Count() int {
	c := bm
	count := 0
	for c != 0 {
		c &= c - 1
		count++
	}
	return count
}
