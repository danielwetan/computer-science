# Traversing a tree means visiting every node in the tree. You might, for instance, want to add all the values in the tree or find the largest one. 
# For all these operations, you will need to visit each node of the tree.
#
# Linear data structures like arrays, stacks, queues, and linked list have only one way to read the data. 
# But a hierarchical data structure like a tree can be traversed in different ways.
# 
# Reference: https://www.programiz.com/dsa/tree-traversal

class TreeNode:
  def __init__(self, value):
    self.value = value 
    self.left = None
    self.right = None

def inorder_traversal(root):
  if root:
    inorder_traversal(root.left)
    print(root.value, end=" ")
    inorder_traversal(root.right)

def preorder_traversal(root):
    if root:
        print(root.value, end=" ")
        preorder_traversal(root.left)
        preorder_traversal(root.right)

def postorder_traversal(root):
    if root:
        postorder_traversal(root.left)
        postorder_traversal(root.right)
        print(root.value, end=" ")

# inorder_traversal performs an inorder traversal (left-root-right) of the tree.
# preorder_traversal performs a preorder traversal (root-left-right) of the tree.
# postorder_traversal performs a postorder traversal (left-right-root) of the tree.

# Example tree
#        1
#      /   \
#     2     3
#    / \   / \
#   4   5 6   7

root = TreeNode(1)
root.left = TreeNode(2)
root.right = TreeNode(3)
root.left.left = TreeNode(4)
root.left.right = TreeNode(5)
root.right.left = TreeNode(6)
root.right.right = TreeNode(7)

print("Inorder Traversal:")
inorder_traversal(root)
print("\nPreorder Traversal:")
preorder_traversal(root)
print("\nPostorder Traversal:")
postorder_traversal(root)

