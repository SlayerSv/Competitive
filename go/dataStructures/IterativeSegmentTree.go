package dataStructures

type SegmentTree struct {
	tree []int
	n    int
}

func (st *SegmentTree) update(a, b int) int {
	return a + b
}

func (st *SegmentTree) Build(arr []int) {
	n := len(arr)
	l := 1
	for l < n {
		l *= 2
	}
	st.tree = make([]int, l*2)
	st.n = l
	for i := 0; i < n; i++ {
		st.tree[st.n+i] = arr[i]
	}
	for i := st.n - 1; i > 0; i-- {
		st.tree[i] = st.update(st.tree[i<<1], st.tree[i<<1|1])
	}
}

func (st *SegmentTree) Set(i, val int) {
	i += st.n
	st.tree[i] += val
	i = i >> 1
	for i > 0 {
		st.tree[i] = st.update(st.tree[i<<1], st.tree[i<<1|1])
		i = i >> 1
	}
}

func (st *SegmentTree) Get(left, right int) int {
	if left > right {
		return 0
	}
	if left < 0 {
		left = st.n
	} else {
		left += st.n
	}
	if right >= st.n {
		right = len(st.tree) - 1
	} else {
		right += st.n
	}
	result := 0
	for left <= right {
		if left&1 == 1 {
			result = st.update(result, st.tree[left])
			left++
		}
		if right&1 == 0 {
			result = st.update(result, st.tree[right])
			right--
		}
		left = left >> 1
		right = right >> 1
	}
	return result
}
