# https://leetcode.com/problems/two-sum-ii-input-array-is-sorted

from typing import List 

class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
      left = 0
      right = len(numbers)-1

      try:
        while True:
          if target == numbers[left] + numbers[right]:
            return [left+1, right+1]

          if target > numbers[left] + numbers[right]:
            left += 1
            continue 

          if target < numbers[left] + numbers[right]:
            right -= 1
            continue 
      except:
        return []


s = Solution()
print(
  [2, 4] == s.twoSum([1, 2, 3, 4, 6], 6)
)