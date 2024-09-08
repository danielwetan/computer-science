# https://leetcode.com/problems/matrix-diagonal-sum/

from typing import List 

class Solution:
    def diagonalSum(self, mat: List[List[int]]) -> int:
      sum = 0
      y = 0
      x = 0

      for _ in range(len(mat[0])):
        left = mat[x][y]
        right = mat[len(mat)-1-x][y]

        if x == len(mat)-1-x and y == len(mat)-1-y:
          sum += left  
        else:
          sum += left
          sum += right

        x += 1
        y += 1

      return sum
    
input = [
  [7,3,1,9],
  [3,4,6,9],
  [6,9,6,6],
  [9,5,8,5]
]
want = 55

s = Solution()
got = s.diagonalSum(input)

if got != want:
  print("expected = {0}, but got = {1}".format(want, got))
else:
  print("passed")