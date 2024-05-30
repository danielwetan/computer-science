# https://leetcode.com/problems/lemonade-change

from typing import List

class Solution:
    def lemonadeChange(self, bills: List[int]) -> bool:
        cash = {
            '5': 0,
            '10': 0,
            '20': 0,
        }
        
        for x in bills:
            cash[str(x)] = cash.get(str(x), 0) + 1

            if x == 10:
                if cash.get('5', 0) >= 1:
                    cash['5'] = cash.get('5')-1
                else:
                    return False
                
            if x == 20:
                if cash.get('10', 0) >= 1 and cash.get('5', 0) >= 1:
                    cash['10'] = cash.get('10')-1
                    cash['5'] = cash.get('5')-1
                    continue

                if cash.get('5',0) >= 3:
                    cash['5'] = cash.get('5')-3
                else:
                    return False
                        
        return True
    

input = [5,5,5,5,10,5,10,10,10,20]
want = True
got = Solution().lemonadeChange(input)

if got != want:
    print("expected = {0}, but got = {1}".format(want, got))
else:
   print("success")