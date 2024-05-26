# A stack is a linear data structure that follows the principle of Last In First Out (LIFO). 
# This means the last element inserted inside the stack is removed first.
# You can think of the stack data structure as the pile of plates on top of another.
# 
# Here, you can:
# Put a new plate on top
# Remove the top plate
# And, if you want the plate at the bottom, you must first remove all the plates on top. This is exactly how the stack data structure works.
#
# LIFO Principle of Stack
# In programming terms, putting an item on top of the stack is called push and removing an item is called pop.
#
# Basic Operations of Stack
# There are some basic operations that allow us to perform different actions on a stack.
# - Push: Add an element to the top of a stack
# - Pop: Remove an element from the top of a stack
# - IsEmpty: Check if the stack is empty
# - IsFull: Check if the stack is full
# - Peek: Get the value of the top element without removing it
#
# Reference: https://www.programiz.com/dsa/stack

def create_stack():
    stack = []
    return stack

# Creating an empty stack
def check_empty(stack):
    return len(stack) == 0


# Adding items into the stack
def push(stack, item):
    stack.append(item)
    print("pushed item: " + item)


# Removing an element from the stack
def pop(stack):
    if (check_empty(stack)):
        return "stack is empty"

    return stack.pop()

stack = create_stack()
push(stack, str(1))
push(stack, str(2))
push(stack, str(3))
push(stack, str(4))
print("popped item: " + pop(stack))
print("stack after popping an element: " + str(stack))
