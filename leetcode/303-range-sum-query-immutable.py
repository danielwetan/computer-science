# https://leetcode.com/problems/range-sum-query-immutable

from typing import List 

class NumArray:
    def __init__(self, nums: List[int]):
      self.nums = nums 

    def sumRange(self, left: int, right: int) -> int:
        new_arr = self.nums[left:right+1]
        sum = 0
        for x in new_arr:
          sum += x

        return sum

# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# param_1 = obj.sumRange(left,right)
    
# ["NumArray","sumRange","sumRange","sumRange"]
# [[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]
    
s = NumArray([-2,0,3,-5,2,-1])
print(s.sumRange(0, 2) == 1)
print(s.sumRange(2, 5) == -1)
print(s.sumRange(0, 5) == -3)
