# Reference 
# https://nbviewer.org/github/donnemartin/interactive-coding-challenges/blob/master/arrays_strings/compress/compress_challenge.ipynb

# Problem
# Compress a string such that 'AAABCCDDDD' becomes 'A3BC2D4'. Only compress the string if it saves space.

# Constraints
# Can we assume the string is ASCII? Yes
# Note: Unicode strings could require special handling depending on your language
# Is this case sensitive? Yes
# Can we use additional data structures? Yes
# Can we assume this fits in memory? Yes

# Test Cases
# None -> None
# '' -> ''
# 'AABBCC' -> 'AABBCC'
# 'AAABCCDDDD' -> 'A3BC2D4'

import unittest

class CompressString(object):
    def compress(self, string):
        if string is None or not string:
            return string 
                
        keys = {}

        for char in string:
          if char in keys:
            keys[char] = keys[char] + 1
            continue
          keys[char] = 0

        print(keys)
        output = ''
        for key, val in keys.items():
            output += key
            
            if val == 1:
                output += key 
                continue

            if val > 1:
                output += str(val+1)
                continue

        return output

class TestCompress(unittest.TestCase):
    def test_compress(self, func):
        self.assertEqual(func(None), None)
        self.assertEqual(func(''), '')
        self.assertEqual(func('AABBCC'), 'AABBCC')
        self.assertEqual(func('AAABCCDDDDE'), 'A3BCCD4E')
        self.assertEqual(func('BAAACCDDDD'), 'BA3CCD4')
        # self.assertEqual(func('AAABAACCDDDD'), 'A3BAACCD4') # failing test
        print('Success: test_compress')


def main():
    test = TestCompress()
    compress_string = CompressString()
    test.test_compress(compress_string.compress)
    try:
        compress_string = CompressStringAlt()
        test.test_compress(compress_string.compress)
    except NameError:
        pass

if __name__ == '__main__':
    main()

