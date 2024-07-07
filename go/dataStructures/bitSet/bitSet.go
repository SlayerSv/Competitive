package bitSet

type bitMask int64

func (bm bitMask) Set(i int) {
	bm |= 1 << i
}

func (bm bitMask) Clear(i int) {
	bm &= ^(1 << i)
}

func (bm bitMask) IsEmpty() bool {
	return bm & ^(0) == 0
}

func (bm bitMask) IsBit(i int) bool {
	return bm&(1<<i) != 0
}

func (bm bitMask) Invert(i int) {
	bm ^= 1 << i
}

func (bm bitMask) Count() int {
	c := bm
	count := 0
	for c != 0 {
		c &= c - 1
		count++
	}
	return count
}

type bitSet []bitMask

func (bs *bitSet) Build(n int) {
	count := n / 64
	for i := 0; i < count; i++ {
		*bs = append(*bs, bitMask(0))
	}
}

func (bs bitSet) Set(n int) {
	i := n / 64
	bs[i].Set(n % 64)
}

func (bs bitSet) Clear(n int) {
	i := n / 64
	bs[i].Clear(n % 64)
}

func (bs bitSet) Clone() bitSet {
	n := len(bs)
	newBS := make([]bitMask, n)
	for i := 0; i < n; i++ {
		newBS[i] = bs[i]
	}
	return newBS
}

func (bs bitSet) Merge(bs2 bitSet) {
	n := len(bs)
	for i := 0; i < n || i < len(bs2); i++ {
		bs[i] |= bs2[i]
	}
}

func (bs bitSet) Has(n int) bool {
	i := n / 64
	return bs[i].IsBit(n % 64)
}
