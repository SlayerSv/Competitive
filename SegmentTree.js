class SegmentTree {
  tree;
  updater;
  n = 0;
  constructor(array, updater = (a, b) => Math.max(a, b)) {
    this.updater = updater;
    this.n = 1 << this.getNextPowerOf2(array.length);
    this.build(array);
  }

  build(array) {
    this.tree = new Array(this.n * 2).fill(0);
    for(let i = 0; i < array.length; i++) {
      this.tree[this.n + i] = array[i];
    }
    for(let i = this.n - 1; i > 0; i--) {
      this.tree[i] = this.updater(this.tree[i << 1], this.tree[i << 1 | 1]);
    }
  }

  update(index, value) {
    index = index + this.n;
    this.tree[index] = value;
    index = index >> 1;
    while (index > 0) {
      this.tree[index] = this.updater(this.tree[index << 1], this.tree[index << 1 | 1]);
      index = index >> 1;
    }
  }

  query(left, right) {
    if (left > right) return;
    left = left < 0 ? this.n : left + this.n;
    right = right >= this.n ? this.tree.length - 1 : right + this.n;
    let result = 0;
    while (left <= right) {
      if ((left & 1) === 1) {
        result = this.updater(result, this.tree[left++]);
      }
      if ((right & 1) === 0) {
        result = this.updater(result, this.tree[right--]);
      }
      left = left >> 1;
      right = right >> 1;
    }
    return result;
  }

  getNextPowerOf2(length) {
    let n = 1;
    let i = 0;
    while (n << i < length) {
      i++;
    }
    return i;
  }
}