class BSTNode {
  priority;
  value;
  count = 1;
  left = null;
  right = null;
  constructor(value = 0) {
    this.priority = Math.random();
    this.value = value;
  }
}

class Treap {
  size = 0;
  root = null;

  constructor() {}

  split(node, key) {
    if (!node) return [null, null];
    if (node.value < key) {
      const [lesser, greater] = this.split(node.right, key);
      node.right = lesser;
      return [node, greater]
    } else {
      const [lesser, greater] = this.split(node.left, key);
      node.left = greater;
      return [lesser, node];
    }
  }

  merge(node1, node2) {
    if (!node1  || !node2) return node1 || node2;
    if (node1.priority > node2.priority) {
      node1.right = this.merge(node1.right, node2);
      return node1;
    } else {
      node2.left = this.merge(node1, node2.left);
      return node2;
    }
  }

  add(value) {
    let newNode;
    let [lesser, greater] = this.split(this.root, value);
    let equal;
    [equal, greater] = this.split(greater, value + 1);
    if (equal) {
      equal.count++;
      newNode = equal;
    } else {
      newNode = new BSTNode(value);
    }
    lesser = this.merge(lesser, newNode);
    this.root = this.merge(lesser, greater);
    this.size++;
  }

  delete(value) {
    let [lesser, greater] = this.split(this.root, value);
    let equal = null;
    [equal, greater] = this.split(greater, value + 1);
    if (equal) {
      if (equal.count > 1) {
        equal.count--;
        lesser = this.merge(lesser, equal);
      }
      this.size--;
    }
    this.root = this.merge(lesser, greater);
  }

  has(key) {
    let tmp = this.root;
    while (tmp) {
      if (key < tmp.value) {
        tmp = tmp.left;
      } else if (tmp.value === key) {
        return true;
      } else {
        tmp = tmp.right;
      }
    }
    return false;
  }

  prevValue(key) {
    const [lesser, greater] = this.split(this.root, key);
    let tmp = lesser;
    while (tmp && tmp.right) {
      tmp = tmp.right;
    }
    this.root = this.merge(lesser, greater)
    return tmp ? tmp.value : null;
  }

  nextValue(key) {
    const [lesser, greater] = this.split(this.root, key + 1);
    let tmp = greater;
    while (tmp && tmp.left) {
      tmp = tmp.left;
    }
    this.root = this.merge(lesser, greater)
    return tmp ? tmp.value : null;
  }

  min() {
    let tmp = this.root;
    while (tmp && tmp.left) {
      tmp = tmp.left;
    }
    return tmp ? tmp.value : null;
  }

  max() {
    let tmp = this.root;
    while (tmp && tmp.right) {
      tmp = tmp.right;
    }
    return tmp ? tmp.value : null;
  }

  print() {
    let tmp = this.root;
    const stack = [];
    while (stack.length || tmp) {
      while (tmp) {
        stack.push(tmp);
        tmp = tmp.left;
      }
      tmp = stack.pop();
      console.log(tmp.value, tmp.count);
      tmp = tmp.right;
    }
  }

  indexOf(key) {
    let tmp = this.root;
    let index = 0;
    const stack = [];
    while (stack.length || tmp) {
      while (tmp) {
        stack.push(tmp);
        tmp = tmp.left;
      }
      tmp = stack.pop();
      if (tmp.value === key) return index;
      index++;
      tmp = tmp.right;
    }
    return -1;
  }

  getByIndex(i) {
    let tmp = this.root;
    let index = 0;
    const stack = [];
    while (stack.length || tmp) {
      while (tmp) {
        stack.push(tmp);
        tmp = tmp.left;
      }
      tmp = stack.pop();
      if (index === i) return tmp.value;
      index++;
      tmp = tmp.right;
    }
    return null;
  }
}