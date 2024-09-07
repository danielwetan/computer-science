from typing import List

class ListNode:
  def __init__(self, val, next=None):
    self.val = val 
    self.next = next 


def reverse_linked_list(head: List[int]) -> List[int]:
  prev = None 
  current = head 

  while current is not None:
    next = current.next 
    current.next = prev 
    prev = current 
    current = next 

  return prev


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

printLinkedList(reverse_linked_list(a))