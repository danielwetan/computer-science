// https://leetcode.com/problems/merge-two-sorted-lists

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} list1
 * @param {ListNode} list2
 * @return {ListNode}
 */
var mergeTwoLists = function(list1, list2) {
  let result = new ListNode(0)
  let tempNode = result

  while (list1 && list2) {
    if (list1.val < list2.val) {
      tempNode.next = list1 
      list1 = list1.next
    } else {
      tempNode.next = list2
      list2 = list2.next
    }
    tempNode = tempNode.next
  }
  tempNode.next = list1 || list2

  return result.next
};

function ListNode(val, next = null) {
  this.val = val;
  this.next = next;
}

let list1 = new ListNode(1, new ListNode(2, new ListNode(4)))
let list2 = new ListNode(1, new ListNode(3, new ListNode(4)))

console.log(
  mergeTwoLists(list1, list2)
)
