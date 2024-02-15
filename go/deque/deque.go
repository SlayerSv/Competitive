package deque

type deque struct {
	arr      []int
	capacity int
	length   int
	front    int
	back     int
}

func NewDeque(n int) *deque {
	if n <= 0 {
		n = 10
	}
	return &deque{
		arr:      make([]int, n),
		capacity: n,
		length:   0,
		front:    n - 1,
		back:     0,
	}
}

func (d *deque) PushBack(val int) {
	if d.length == d.capacity {
		d.grow()
	}
	d.arr[d.back] = val
	d.back = (d.back + 1) % d.capacity
	d.length++
}

func (d *deque) PushFront(val int) {
	if d.length == d.capacity {
		d.grow()
	}
	d.arr[d.front] = val
	d.front = (d.front + d.capacity - 1) % d.capacity
	d.length++
}

func (d *deque) PopBack() int {
	d.back = (d.back + d.capacity - 1) % d.capacity
	d.length--
	return d.arr[d.back]
}

func (d *deque) PopFront() int {
	d.front = (d.front + 1) % d.capacity
	d.length--
	return d.arr[d.front]
}

func (d *deque) grow() {
	newArr := make([]int, d.capacity*2)
	d.front = (d.front + 1) % d.capacity
	for i := 0; i < d.length; i++ {
		newArr[i] = d.arr[d.front]
		d.front = (d.front + 1) % d.capacity
	}
	d.front = d.capacity*2 - 1
	d.back = d.length
	d.capacity *= 2
	d.arr = newArr
}

func (d *deque) Back() int {
	if d.length == 0 {
		return 0
	}
	return d.arr[(d.back+d.capacity-1)%d.capacity]
}

func (d *deque) Front() int {
	if d.length == 0 {
		return 0
	}
	return d.arr[(d.front+1)%d.capacity]
}

func (d *deque) Get(i int) int {
	if d.length == 0 || i >= d.length || i < -d.length {
		return 0
	}
	if i < 0 {
		i = d.length + i
	}
	return d.arr[(d.front+i+1)%d.capacity]
}
