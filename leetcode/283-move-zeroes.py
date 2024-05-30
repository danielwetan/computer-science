# https://leetcode.com/problems/move-zeroes/

from typing import List

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """

        for i in range(len(nums)):
          if nums[i] == 0:
            for j in range(len(nums)-i):
              return
              #  temp = nums[i]
              #  nums[i] = nums[j]
              #  nums[j] = temp

        return
    
input = [0,1,0,3,12]
want = [1,3,12,0,0]
Solution().moveZeroes(input)

if input != want:
    print("expected = {0}, but got = {1}".format(want, input))
else:
   print("success")