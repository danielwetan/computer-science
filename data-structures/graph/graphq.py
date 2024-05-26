# A graph data structure is a collection of nodes that have data and are connected to other nodes.
# 
# Let's try to understand this through an example. 
# On facebook, everything is a node. 
# That includes User, Photo, Album, Event, Group, Page, Comment, Story, Video, Link, Note...anything that has data is a node.
# 
# Every relationship is an edge from one node to another. 
# Whether you post a photo, join a group, like a page, etc., a new edge is created for that relationship.
# 
# All of facebook is then a collection of these nodes and edges. This is because facebook uses a graph data structure to store its data.
# 
# More precisely, a graph is a data structure (V, E) that consists of
# A collection of vertices V
# A collection of edges E, represented as ordered pairs of vertices (u,v)
# 
# For example:
# V = {0, 1, 2, 3}
# E = {(0,1), (0,2), (0,3), (1,2)}
# G = {V, E}
#
# Reference: https://www.programiz.com/dsa/graph

# Graph terminology
# Adjacency: A vertex is said to be adjacent to another vertex if there is an edge connecting them. Vertices 2 and 3 are not adjacent because there is no edge between them.
# Path: A sequence of edges that allows you to go from vertex A to vertex B is called a path. 0-1, 1-2 and 0-2 are paths from vertex 0 to vertex 2.
# Directed Graph: A graph in which an edge (u,v) doesn't necessarily mean that there is an edge (v, u) as well. The edges in such a graph are represented by arrows to show the direction of the edge.

# Graph representation
# Graphs are commonly represented in two ways:
# 1. Adjacency Matrix
# An adjacency matrix is a 2D array of V x V vertices. Each row and column represent a vertex.
# If the value of any element a[i][j] is 1, it represents that there is an edge connecting vertex i and vertex j.
# The adjacency matrix for the graph we created above is
adjacency_matrix = [
  [0, 1, 1, 1],
  [1, 0, 1, 0],
  [1, 1, 0, 0],
  [1, 0, 0, 0], 
]
# Since it is an undirected graph, for edge (0,2), we also need to mark edge (2,0); 
# making the adjacency matrix symmetric about the diagonal.
# Edge lookup(checking if an edge exists between vertex A and vertex B) is extremely fast in adjacency matrix representation 
# but we have to reserve space for every possible link between all vertices(V x V), so it requires more space.
# 
# 2. Adjacency List
# An adjacency list represents a graph as an array of linked lists.
# The index of the array represents a vertex 
# and each element in its linked list represents the other vertices that form an edge with the vertex.
# The adjacency list for the graph we made in the first example is as follows:
# 
# An adjacency list is efficient in terms of storage because we only need to store the values for the edges. 
# For a graph with millions of vertices, this can mean a lot of saved space.
class Node:
  def __init__(self, item):
    self.item = item 
    self.next = None 

class LinkedList:
  def __init__(self):
    self.head = None 

zero = LinkedList()
zero.head = Node(0)
zero.head.next = Node(1)
zero.head.next.next = Node(2)
zero.head.next.next.next = Node(3)
# 0 -> 1 -> 2 -> 3

one = LinkedList()
one.head = Node(1)
one.head.next = Node(0)
one.head.next.next = Node(2)
# 1 -> 0 -> 2

two = LinkedList()
two.head = Node(2)
two.head.next = Node(0)
two.head.next.next = Node(1)
# 2 -> 0 -> 1

three = LinkedList()
three.head = Node(3)
three.head.next = Node(0)
# 3 -> 0

# Graph Operations
# The most common graph operations are:
# - Check if the element is present in the graph
# - Graph Traversal
# - Add elements(vertex, edges) to graph
# - Finding the path from one vertex to another