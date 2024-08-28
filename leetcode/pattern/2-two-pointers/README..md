The Two Pointers pattern involves using two pointers to iterate through an array or list, often used to find pairs or elements that meet specific criteria.

Use this pattern when dealing with sorted arrays or lists where you need to find pairs that satisfy a specific condition.

Sample Problem:
Find two numbers in a sorted array that add up to a target value.

Example:
- Input: nums = [1, 2, 3, 4, 6], target = 6
- Output: [1, 3]

Explanation:
- Initialize two pointers, one at the start (left) and one at the end (right) of the array.
- Check the sum of the elements at the two pointers.
- If the sum equals the target, return the indices.
- If the sum is less than the target, move the left pointer to the right.
- If the sum is greater than the target, move the right pointer to the left.