# https://leetcode.com/problems/robot-return-to-origin/

class Solution:
    def judgeCircle(self, moves: str) -> bool:
        position = [0, 0]
        initialPosition = [0, 0]

        movement = list(moves)
        
        movementValue = {
            'U': [0, 1],
            'D': [0, -1],
            'R': [1, 0],
            'L': [-1, 0]
        }

        for m in movement:
            value = movementValue.get(m, [0, 0])
            position[0] += value[0]
            position[1] += value[1]

        return position == initialPosition
    

case1 = "UD"
case2 = "LL"
want1 = True 
want2 = False

s = Solution()

got1 = s.judgeCircle(case1)
got2 = s.judgeCircle(case2)

if got1 != want1:
    print("case1 - expected = {0}, but got = {1}".format(want1, got1))
  
if got2 != want2:
    print("case2 - expected = {0}, but got = {1}".format(want2, got2))
  