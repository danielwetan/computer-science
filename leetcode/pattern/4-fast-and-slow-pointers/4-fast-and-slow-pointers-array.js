function hasCycleInArray(arr) {
  let slow = 0; // Slow pointer (tortoise)
  let fast = 0; // Fast pointer (hare)

  while (true) {
    // Move slow pointer by 1 step
    slow = arr[slow];

    // Move fast pointer by 2 steps
    fast = arr[arr[fast]];

    // If either pointer reaches an invalid index, there is no cycle
    if (slow === undefined || fast === undefined || arr[fast] === undefined) {
      return false;
    }

    // Cycle detected if slow and fast pointers meet
    if (slow === fast) {
      return true;
    }
  }
}

// Example 1: Array with No Cycle
const arr1 = [1, 2, 3, 4, 5, -1]; // The last element points to an invalid index
console.log(hasCycleInArray(arr1)); // Output: false

// Example 2: Array with a Cycle
const arr2 = [1, 2, 3, 4, 2]; // The last element points back to index 2, creating a cycle
console.log(hasCycleInArray(arr2)); // Output: true


