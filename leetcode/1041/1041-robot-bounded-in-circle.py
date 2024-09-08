# https://leetcode.com/problems/robot-bounded-in-circle/

class Solution:
    def isRobotBounded(self, instructions: str) -> bool:
      if len(instructions) > 100:
         return False

      initialPosition = [0, 0]
      position = [0, 0]
      movement = list(instructions)

      direction = ''
      directionValue = {
        'N': [0, 1],
        'S': [0, -1],
        'E': [1, 0],
        'W': [-1, 0]
      }

      # limit loop to 4 times
      for _ in range(4):   
        for m in movement:     
          direction = self.getDirection(direction, m)
          value = directionValue.get(direction)
          if m == 'G':
            position[0] += value[0]
            position[1] += value[1]
        if position == initialPosition:
           return True 

      return False
    
    def getDirection(self, currentDirection: str, instruction: str) -> str:
      if currentDirection == '':
        return 'N'

      switchDirection = {
         'N': {
            'G': 'N',
            'L': 'W',
            'R': 'E',
         },
         'S': {
            'G': 'S',
            'L': 'E',
            'R': 'W',
         },
         'E': {
            'G': 'E',
            'L': 'N',
            'R': 'S'
         },
         'W': {
            'G': 'W',
            'L': 'S',
            'R': 'N'
         }
      }

      return switchDirection.get(currentDirection).get(instruction)

test_cases = {
    "case1": {
        "value": "GGLLGG",
        "want": True
    },
    "case2": {
        "value": "GG",
        "want": False
    },
    "case3": {
        "value": "GL",
        "want": True
    }
}


s = Solution()

for case_name, case_data in test_cases.items():
    value = case_data["value"]
    want = case_data["want"]
    got = s.isRobotBounded(value)
    if got != want:
      print(f"{case_name} - got = {got}, want = {want}")
    else:
      print(f"{case_name} - passed")
       
