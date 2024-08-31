class ListNode {
  constructor(value) {
    this.value = value;
    this.next = null;
  }
}

function reverseBetween(head, m, n) {
  if (!head || m === n) {
    return head; // No need to reverse if the list is empty or m equals n
  }

  let current = head;
  let previous = null;

  // Move the two pointers until they reach the proper starting point in the list.
  for (let i = 1; current !== null && i < m; i++) {
    previous = current;
    current = current.next;
  }

  // We are interested in three parts of the linked list:
  // - The part before index 'm'
  // - The part between 'm' and 'n'
  // - The part after index 'n'
  const lastNodeOfFirstPart = previous; // Points to the node at index 'm-1'
  const lastNodeOfSubList = current; // After reversal, 'current' will become the last node of the sublist
  let next = null; // Will be used to temporarily store the next node

  // Reverse nodes between 'm' and 'n'
  for (let i = 0; current !== null && i < n - m + 1; i++) {
    next = current.next;
    current.next = previous;
    previous = current;
    current = next;
  }

  // Connect with the first part
  if (lastNodeOfFirstPart !== null) {
    lastNodeOfFirstPart.next = previous;
  } else {
    head = previous; // If m = 1, then we are changing the first node (head) of the LinkedList
  }

  // Connect with the last part
  lastNodeOfSubList.next = current;

  return head;
}

// for testing purpose
// Helper function to create a linked list from an array
function createLinkedList(arr) {
  let head = new ListNode(arr[0]);
  let current = head;
  for (let i = 1; i < arr.length; i++) {
    current.next = new ListNode(arr[i]);
    current = current.next;
  }
  return head;
}

// Helper function to print the linked list
function printLinkedList(head) {
  let result = [];
  while (head !== null) {
    result.push(head.value);
    head = head.next;
  }
  console.log(result);
}

// Example: Reverse sublist from position 2 to 4
let head = createLinkedList([1, 2, 3, 4, 5]);
printLinkedList(head); // Output: [1, 2, 3, 4, 5]

head = reverseBetween(head, 2, 4);
printLinkedList(head); // Output: [1, 4, 3, 2, 5]
