# https://leetcode.com/problems/two-sum-ii-input-array-is-sorted

from typing import List 

class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        pair = {}
        for i in range(len(numbers)):
            if pair.get(numbers[i], 0) != 0:
              return [pair[numbers[i]], i+1]

            pair[target-numbers[i]] = i+1

        return pair
    



s = Solution()
print(
  [2, 4] == s.twoSum([1, 2, 3, 4, 6], 6)
)