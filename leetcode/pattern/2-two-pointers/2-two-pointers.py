from typing import List 

class Solution:
  def two_pointers(self, nums: List[int], target: int) -> List[int]:
    left = 0
    right = len(nums)-1

    try:
      while True:
        if target == nums[left] + nums[right]:
          return [left, right]

        if target > nums[left] + nums[right]:
          left += 1
          continue 

        if target < nums[left] + nums[right]:
          right -= 1
          continue 
    except:
      return []


s = Solution()
print(
  [1, 3] == s.two_pointers([1, 2, 3, 4, 6], 6)
)