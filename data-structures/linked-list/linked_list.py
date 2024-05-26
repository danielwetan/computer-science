
# Linked List 
# A linked list is a linear data structure that includes a series of connected nodes. 
# Here, each node stores the data and the address of the next node. For example,
#
# Time Complexity
#          |  Worst Case | Average Case 
# Search   | O(n)        | O(n)
# Insert   | O(n)        | O(n)
# Deletion | O(n)        | O(n)
#
# Space Complexity
# O(n)
#
# Linked List Applications
# - Dynamic memory allocation
# - Implemented in stack and queue
# - In undo functionality of softwares
# - Hash tables, Graphs
#
# Reference: https://www.programiz.com/dsa/linked-list

class Node:
  def __init__(self, item):
    self.item = item 
    self.next = None 

class LinkedList:
  def __init__(self):
    self.head = None 

linked_list = LinkedList()
linked_list.head = Node(1)
second = Node(2)
third = Node(3)

linked_list.head.next = second
second.next = third

# print the linked list 
while linked_list.head != None:
    print(linked_list.head.item, end=" ")
    linked_list.head = linked_list.head.next