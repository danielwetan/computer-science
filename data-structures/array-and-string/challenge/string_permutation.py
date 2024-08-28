# Reference
# https://nbviewer.org/github/donnemartin/interactive-coding-challenges/blob/master/arrays_strings/permutation/permutation_challenge.ipynb

# Constraints
# Can we assume the string is ASCII? Yes
# Note: Unicode strings could require special handling depending on your language
# Is whitespace important? Yes
# Is this case sensitive? 'Nib', 'bin' is not a match? Yes
# Can we use additional data structures? Yes
# Can we assume this fits in memory? Yes

# Test Cases
# One or more None inputs -> False
# One or more empty strings -> False
# 'Nib', 'bin' -> False
# 'act', 'cat' -> True
# 'a ct', 'ca t' -> True

"""
# Python default dict
# If the key is not found it will return default value based on data type

from collection import defaultdict
words = defaultdict(int)
words['a'] = 1
words['b'] = 2

print(
  words['a'], # 1
  words['b'], # 2
  words['c'], # 0
)
"""

import unittest
from collections import defaultdict

class Permutations(object):
    def is_permutation(self, str1, str2):
        if str1 == None or str2 == None:
          return False

        set1 = set(str1)
        set2 = set(str2)
        return set1 == set2 and len(str1) == len(str2)
    
class PermutationsAlt(object):
    def is_permutation(self, str1, str2):
        if str1 is None or str2 is None:
          return False
        if len(str1) != len(str2):
            return False
        
        keys1 = defaultdict(int)
        keys2 = defaultdict(int)
        for char in str1:
          keys1[char] += 1
        for char in str2:
          keys2[char] += 1

        return keys1 == keys2
    
class TestPermutation(unittest.TestCase):
    def test_permutation(self, func):
        self.assertEqual(func(None, 'foo'), False)
        self.assertEqual(func('', 'foo'), False)
        self.assertEqual(func('Nib', 'bin'), False)
        self.assertEqual(func('act', 'cat'), True)
        self.assertEqual(func('a ct', 'ca t'), True)
        self.assertEqual(func('dog', 'doggo'), False)
        print('Success: test_permutation')

def main():
    test = TestPermutation()
    permutations = Permutations()
    test.test_permutation(permutations.is_permutation)
    try:
        permutations_alt = PermutationsAlt()
        test.test_permutation(permutations_alt.is_permutation)
    except NameError:
        pass

if __name__ == '__main__':
    main()

