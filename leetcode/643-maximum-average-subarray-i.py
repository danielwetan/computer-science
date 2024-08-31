from typing import List

class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        window_sum = sum(nums[:k])
        max_sum = window_sum

        for i in range(1, len(nums) - k + 1):
            window_sum += nums[i + k - 1] - nums[i - 1]
            max_sum = max(max_sum, window_sum)

        return max_sum / k  # Return the maximum average

# Test cases
s = Solution()

print(
    12.75000 == s.findMaxAverage([1, 12, -5, -6, 50, 3], 4),
    -1.0 == s.findMaxAverage([-1], 1),
)
