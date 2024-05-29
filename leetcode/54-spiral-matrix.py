# https://leetcode.com/problems/spiral-matrix/

from typing import List

class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        # get total indices
        xLength = len(matrix[0])
        yLength = len(matrix)
        totalIndices = xLength * yLength

        # solution result
        result = []

        # set direction
        directionValue = {
           # direcitionName: [y, x]
           'N': [-1, 0],
           'S': [1, 0],
           'E': [0, 1],
           'W': [0, -1],
        }

        # merged xy for example {'00': True}
        traversedPath = {}

        # movement position
        # pos = f"{position[0]}{position[1]}" [y, x]
        position = [0, 0]

        # default N
        currentDirection = 'E'
        movementCounter = 0
        for _ in range(9):
          # record traveled path
          traversedPath[f"{position[0]}{position[1]}"] = True
          result.append(matrix[position[0]][position[1]])

          # move                  
          movement = directionValue.get(currentDirection)
          position[0] += movement[0]
          position[1] += movement[1]


          # check if already traversed path
          # if traversedPath.get(f"{position[0]}{position[1]}") == True:
          #   currentDirection = self.switchDirection(currentDirection)
          #   movementCounter = 0
          #   continue
          # else:
          #   movementCounter += 1   

          # switch direction
          if (currentDirection == 'E' or currentDirection == 'W') and movementCounter >= xLength-2:
            currentDirection = self.switchDirection(currentDirection)
            movementCounter = 0
            continue
          else:
            movementCounter += 1   

          # print(movementCounter)
          if (currentDirection == 'N' or currentDirection == 'S') and movementCounter >= yLength-1:
            currentDirection = self.switchDirection(currentDirection)
            movementCounter = 0
            continue
          else:
            movementCounter += 1

        return result

    def switchDirection(self, currentDirection) -> str:
      # create a spiral direction
      switchDirection = {
          'E': 'S',
          'S': 'W',
          'W': 'N',
          'N': 'E'
      }

      return switchDirection.get(currentDirection)


input = [ 
  [1,2,3],
  [4,5,6],
  [7,8,9]
]

want = [1,2,3,6,9,8,7,4,5]

s = Solution()
got = s.spiralOrder(input)

if got != want:
  print("expected = {0}, but got = {1}".format(want, got))
else:
  print("passed")


