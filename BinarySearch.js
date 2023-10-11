/** Returns index that's coming AFTER the last element that's less or equal to target.
 *  That index may be out of bounds of the array.
 */
function rightBinSearch(arr, target) {
  let l = 0;
  let r = arr.length - 1;
  let mid;
  while (l <= r) {
    mid = l + r + 1 >> 1;
    if (target < arr[mid]) {
      r = mid - 1;
    } else {
      l = mid + 1;
    }
  }
  return l;
}

/** Returns index of the first element that's greater or equal to target. */
function leftBinSearch(arr, target) {
  let l = 0;
  let r = arr.length - 1;
  let mid;
  while (l <= r) {  
    mid = l + r >> 1;
    if (target <= arr[mid]) {
      r = mid - 1;
    } else {
      l = mid + 1;
    }
  }
  return l;
}
