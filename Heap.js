class Heap {
  heap;
  needSwap;

  constructor(array = [], compareFunc = (a, b) => a < b) {
    this.heap = array;
    this.needSwap = compareFunc;
    this.heapify();
  }

  down(i) {
    let childIndex = i * 2 + 1;
    while (childIndex < this.heap.length) {
      if (childIndex + 1 < this.heap.length
          && this.needSwap(this.heap[childIndex + 1], this.heap[childIndex])) {
        childIndex++;
      }
      if (this.needSwap(this.heap[childIndex], this.heap[i])) {
        const tmp = this.heap[i]
        this.heap[i] = this.heap[childIndex];
        this.heap[childIndex] = tmp;
        i = childIndex;
        childIndex = i * 2 + 1;
      } else {
        break;
      }
    }
  }

  up(i) {
    let parentIndex = (i - 1) >> 1;
    while (i !== 0) {
      if (this.needSwap(this.heap[i], this.heap[parentIndex])) {
        const tmp = this.heap[parentIndex]
        this.heap[parentIndex] = this.heap[i];
        this.heap[i] = tmp;
        i = parentIndex;
        parentIndex = (i - 1) >> 1;
      } else {
        break;
      }
    }
  }

  push(value) {
    this.heap.push(value);
    this.up(this.heap.length - 1);
  }

  pop() {
    const value = this.heap[0];
    this.heap[0] = this.heap[this.heap.length - 1];
    this.heap.pop();
    this.down(0);
    return value;
  }
  
  heapify() {
    if (this.heap.length < 2) {
      return;
    }
    for (let i = (this.heap.length >> 1) - 1; i >= 0; i--) {
      this.down(i);
    }
  }
}

// TESTING

const heap = new Heap();

for (let i = 0; i < 10; i++) {
  heap.push(10 - i);
}

for (let i = 0; i < 10; i++) {
  console.log(heap.heap);
  console.log(heap.pop());
}

const heap1 = new Heap([5, 3, 27, 99, 45, 1, 0, 20])

for (let i = 0; i < 10; i++) {
  console.log(heap1.heap);
  console.log(heap1.pop());
}