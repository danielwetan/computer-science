# Reference
# https://nbviewer.org/github/donnemartin/interactive-coding-challenges/blob/master/arrays_strings/rotation/rotation_challenge.ipynb

# Problem
# Determine if a string s1 is a rotation of another string s2, by calling (only once) a function is_substring.

# Constraints
# Can we assume the string is ASCII? Yes
# Note: Unicode strings could require special handling depending on your language
# Is this case sensitive? Yes
# Can we use additional data structures? Yes
# Can we assume this fits in memory? Yes

# Test Cases
# Any strings that differ in size -> False
# None, 'foo' -> False (any None results in False)
# ' ', 'foo' -> False
# ' ', ' ' -> True
# 'foobarbaz', 'barbazfoo' -> True

import unittest

class Rotation(object):
    def is_substring(self, s1, s2):
        return s1 in s2 

    def is_rotation(self, s1, s2):
        if s1 is None or s2 is None:
            return False
        if len(s1) != len(s2):
            return False

        return self.is_substring(s1, s2 + s2)        

class TestRotation(unittest.TestCase):
    def test_rotation(self):
        rotation = Rotation()
        self.assertEqual(rotation.is_rotation('o', 'oo'), False)
        self.assertEqual(rotation.is_rotation(None, 'foo'), False)
        self.assertEqual(rotation.is_rotation('', 'foo'), False)
        self.assertEqual(rotation.is_rotation('', ''), True)
        self.assertEqual(rotation.is_rotation('foobarbaz', 'barbazfoo'), True)
        print('Success: test_rotation')

def main():
    test = TestRotation()
    test.test_rotation()


if __name__ == '__main__':
    main()


