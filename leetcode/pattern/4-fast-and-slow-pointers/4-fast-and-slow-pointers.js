const assert = require('assert')

class ListNode {
  constructor(value) {
    this.value = value;
    this.next = null;
  }
}

function hasCycle(head) {
  if (!head || !head.next) {
    return false;
  }

  let slow = head; // slow pointers (tortoise)
  let fast = head; // fast pointers (hare)

  while (fast && fast.next) {
    slow = slow.next; // move slow pointers by 1 step
    fast = fast.next.next; // move fast pointers by 2 step

    if (slow === fast) {
      return true; // cycle detected
    }
  }

  return false;
}

// step over the time
// X -> Y -> Z -> Y
// SF        F
//      S    F
//           SF

// // Example 1: Linked List with No Cycle
const nodeA = new ListNode("A");
const nodeB = new ListNode("B");
const nodeC = new ListNode("C");

nodeA.next = nodeB;
nodeB.next = nodeC;

assert.equal(hasCycle(nodeA), false)

// Example 2: Linked List with a Cycle
const nodeX = new ListNode("X");
const nodeY = new ListNode("Y");
const nodeZ = new ListNode("Z");

nodeX.next = nodeY;
nodeY.next = nodeZ;
nodeZ.next = nodeY; // Creates a cycle

assert.equal(hasCycle(nodeX), true)
