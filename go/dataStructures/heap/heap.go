package heap

type Heap[T any] struct {
    data []T
    less func(T, T) bool
}
func (h *Heap[T]) Len() int           { return len(h.data) }
func (h *Heap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h *Heap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *Heap[T]) Push(x any)         { h.data = append(h.data, x.(T)) }
func (h *Heap[T]) Pop() any {
    old := h.data
    n := len(old)
    x := old[n-1]
    h.data = old[0 : n-1]
    return x
}
func NewHeap[T any](less func(T, T) bool) *Heap[T] {
    return &Heap[T]{less: less}
}
