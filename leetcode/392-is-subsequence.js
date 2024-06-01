// https://leetcode.com/problems/is-subsequence

/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isSubsequence = function(s, t) {
  target = s.split('')

  t.split('').forEach(str => {
    if (target[0] == str) {
      target.shift()
    }
  })

  return target.length == 0
};

console.log(
  isSubsequence("abc", "ahbgdc") == true,
  isSubsequence("axc", "ahbgdc") == false,
)