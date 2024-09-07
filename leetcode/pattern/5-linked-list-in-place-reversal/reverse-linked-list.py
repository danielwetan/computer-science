class ListNode:
  def __init__(self, val, next=None):
    self.val = val 
    self.next = next 

class Solution:
  def reverseList(self, head: ListNode) -> ListNode:
    prev, curr = None, head 

    while curr:
      nxt = curr.next
      curr.next = prev 
      prev = curr 
      curr = nxt 
    
    return prev
  
  def reverseListRecursino(self, head: ListNode) -> ListNode:
    if not head:
      return None 
    
    newHead = head 
    if head.next:
      newHead = self.reverseList(head.next)
      head.next.next = head 
    head.next = None 
    return newHead
  

def printLinkedList(head: ListNode):
  if head == None:
    return 
  
  while head:
    print(head.val) 
    head = head.next

a = ListNode('a')
b = ListNode('b')
c = ListNode('c')
d = ListNode('d')
a.next = b 
b.next = c 
c.next = d 

s = Solution()
printLinkedList(a)
print("reversed")
printLinkedList(s.reverseList(a))
# print("reversed recursion")
# printLinkedList(s.reverseListRecursino(a))
