class LazySegmentTree {
  tree;
  lazy;
  n = 0;
  updater;

  constructor(array, updater = (a, b) => a + b) {
    this.n = 1 << this.getNextPowerOf2(array.length);
    this.build(array);
    this.updater = updater;
  }

  updateRange(left, right, value, currLeft = 0, currRight = this.n - 1, currIndex = 1) {
    if (currLeft > currRight || left > currRight || right < currLeft) return;
    if (this.lazy[currIndex]) {
      this.tree[currIndex] += this.lazy[currIndex];
      this.updateChildren(currIndex, this.lazy[currIndex]);
      this.lazy[currIndex] = 0;
    }
    if (currLeft >= left && currRight <= right) {
      this.tree[currIndex] += value;
      this.updateChildren(currIndex, value);
      return;
    }
    const mid = currLeft + currRight >> 1;
    this.updateRange(left, right, value, currLeft, mid, currIndex << 1);
    this.updateRange(left, right, value, mid + 1, currRight, currIndex << 1 | 1);
  }

  queryRange(left, right, currLeft, currRight, currIndex) {
    if (this.lazy[currIndex]) {
      this.tree[currIndex] += this.lazy[currIndex];
      this.updateChildren(currIndex, this.lazy[currIndex]);
      this.lazy[currIndex] = 0;
    }
    if (currLeft > currRight || left > currRight || right < currLeft) return 0;
    if (currLeft >= left && currRight <= right) {
      return this.tree[currIndex];
    }
    const mid = currLeft + currRight >> 1;
    return this.updater(
      this.queryRange(left, right, value, currLeft, mid, currIndex << 1),
      this.queryRange(left, right, value, mid + 1, currRight, currIndex << 1 | 1)
    );
  }

  updateChildren(currIndex, value) {
    if (currIndex << 1 | 1 < this.n) {
      this.lazy[currIndex << 1] += value;
      this.lazy[currIndex << 1 | 1] += value;
    }
  }

  getNextPowerOf2(length) {
    let n = 1;
    let i = 0;
    while (n << i < length) {
      i++;
    }
    return i;
  }

  build(array) {
    this.tree = new Array(this.n * 2).fill(0);
    this.lazy = new Array(this.n * 2).fill(0);
    for(let i = 0; i < array.length; i++) {
      this.tree[this.n + i] = array[i];
    }
    for(let i = this.n - 1; i > 0; i--) {
      this.tree[i] = this.updater(this.tree[i << 1], this.tree[i << 1 | 1]);
    }
  }
}