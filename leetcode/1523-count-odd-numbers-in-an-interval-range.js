// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range

/**
 * @param {number} low
 * @param {number} high
 * @return {number}
 */
var countOdds = function(low, high) {
  if (low % 2 === 0) {
      low += 1;
  }

  if (high % 2 === 0) {
      high -= 1;
  }

  if (low > high) {
      return 0;
  }
  return Math.floor((high - low) / 2) + 1;
};

console.log(
  countOdds(2, 2), // 0
  countOdds(3, 7), // 3
  countOdds(0, 10), // 5
  countOdds(8, 10), // 1
  countOdds(16, 17), // 1
)

