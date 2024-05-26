# https://leetcode.com/problems/baseball-game/

from typing import List

class Solution:
    def calPoints(self, operations: List[str]) -> int:
        result = []
        sum = 0

        for i in range(len(operations)):
          if operations[i] == '+':
            if i < 2:
               continue
            two_previous_score = int(result[len(result)-2]) + int(result[len(result)-1])
            result.append(str(two_previous_score))
          elif operations[i] == 'D':
            if i < 1:
               continue
            previous_score = int(result[len(result)-1])
            result.append(str(previous_score * 2))
          elif operations[i] == 'C':
             if i < 1:
                continue
             result.pop()
          else:
             result.append(operations[i])

        for x in result:
          sum += int(x)   
        
        return sum
    
ops = ["5","-2","4","C","D","9","+","+"]
got = Solution().calPoints(ops)
want = 27

if got != want:
    print("expected = {0}, but got = {1}".format(want, got))
else:
   print("success")