# Tree terminologies
# Node
# A node is an entity that contains a key or value and pointers to its child nodes.
# The last nodes of each path are called leaf nodes or external nodes that do not contain a link/pointer to child nodes.
# The node having at least a child node is called an internal node.
#
# Edge
# It is the link between any two nodes.
#
# Root
# It is the topmost node of a tree.
#
# Height of a Node
# The height of a node is the number of edges from the node to the deepest leaf (ie. the longest path from the node to a leaf node).
#
# Depth of a Node
# The depth of a node is the number of edges from the root to the node.
# 
# Height of a Tree
# The height of a Tree is the height of the root node or the depth of the deepest node.
# 
# Degree of a Node
# The degree of a node is the total number of branches of that node.
#
# Forest
# A collection of disjoint trees is called a forest.
# 
# Reference: https://www.programiz.com/dsa/trees

class TreeNode:
  def __init__(self, value):
    self.value = value 
    self.left = None
    self.right = None

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

