class Solution:
  def prefix_sum(arr, start, end):
    new_arr = arr[start:end+1]
    
    sum = 0
    for x in new_arr:
      sum += x

    return sum
  

s = Solution
print(
  9 == s.prefix_sum([1, 2, 3, 4, 5, 6], 1, 3),
  21 == s.prefix_sum([1, 2, 3, 4, 5, 6], 0, 5)
)

