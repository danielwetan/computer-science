from typing import List

class Solution:
  def findMaxAverage(self, nums: List[int], k: int) -> float:
    window_sum = sum(nums[:k])
    max_sum = window_sum

    for i in range(1, len(nums)-k + 1):
      window_sum += nums[i + k - 1] - nums[i - 1]
      max_sum = max(max_sum, window_sum)

    return max_sum

  def sliding_window(self, nums: List[int], k: int) -> int:
    start = 0 
    end = k

    greatest = 0
    while True:      
      if end > len(nums):
        break

      sum = self.calculate_sub_array(nums[start:end])

      if greatest < sum:
        greatest = sum

      start += 1
      end += 1
    return greatest
  
  def calculate_sub_array(self, arr: List[int]) -> int:
      sum = 0
      for x in arr:
        sum += x
      return sum
  
s = Solution()

print(
 9 == s.sliding_window([2, 1, 5, 1, 3, 2], 3)
)