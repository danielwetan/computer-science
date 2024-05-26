# A queue is a data structure that similar to the ticket queue outside a cinema hall, 
# where the first person entering the queue is the first person who gets the ticket.
# 
# Queue follows the First In First Out (FIFO) rule 
# the item that goes in first is the item that comes out first.
#
# In programming terms, putting items in the queue is called enqueue, 
# and removing items from the queue is called dequeue.
#
# Basic Operations of Queue
# A queue is an object (an abstract data structure - ADT) that allows the following operations:
# - Enqueue: Add an element to the end of the queue
# - Dequeue: Remove an element from the front of the queue
# - IsEmpty: Check if the queue is empty
# - IsFull: Check if the queue is full
# - Peek: Get the value of the front of the queue without removing it
#
# Reference: https://www.programiz.com/dsa/queue

class Queue:
  def __init__(self):
    self.queue = []

  def enqueue(self, item):
    self.queue.append(item)

  def dequeue(self):
    if len(self.queue) < 1:
      return None 
    return self.queue.pop(0)

  def display(self):
    print(self.queue)

  def size(self):
    return len(self.queue)

q = Queue()
q.enqueue(1)
q.enqueue(2)
q.enqueue(3)
q.enqueue(4)
q.enqueue(5)

q.display()

q.dequeue()

print("After removing an element")
q.display()