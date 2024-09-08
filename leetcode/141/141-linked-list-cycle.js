// https://leetcode.com/problems/linked-list-cycle/

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {boolean}
 */
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