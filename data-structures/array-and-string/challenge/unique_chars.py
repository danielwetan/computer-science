# Reference
# https://nbviewer.org/github/donnemartin/interactive-coding-challenges/blob/master/arrays_strings/unique_chars/unique_chars_challenge.ipynb

# Problem: Implement an algorithm to determine if a string has all unique characters.

# Constraints
# Can we assume the string is ASCII?
# Yes
# Note: Unicode strings could require special handling depending on your language
# Can we assume this is case sensitive?
# Yes
# Can we use additional data structures?
# Yes
# Can we assume this fits in memory?
# Yes

# Test Cases
# None -> False
# '' -> True
# 'foo' -> False
# 'bar' -> True

"""
Python loop through map
capitals = {
  "Indonesia": "Jakarta",
  "Malaysia": "Kuala Lumpur",
  "Thailand": "Bangkok"
}

for key in capitals.keys():
  print(key) # Indonesia, Malaysia, Thailand

for val in capitals.values():
  print(val) # Jakarta, Kuala Lumpur, Bangkok

for key, value in capitals.items():
  print(key, value) # Indonesia, Jakarta, ...
"""


import unittest

# Solution 1
class UniqueChars(object):
  def has_unique_chars(self, string):
    if string == None:
       return False
    
    if string == '':
       return True

    keys = {}
    for char in string:
      if char in keys:
        keys[char] = keys[char] + 1
        continue
      keys[char] = 0

    for value in keys.values():
       if value != 0:
          return False    

    return True

# Solution 2
class UniqueCharsSet(object):
   def has_unique_chars(self, string):
      if string is None:
         return False 
      
      return len(set(string)) == len(string)

# Solution 3
class UniqueCharsInPlace(object):
   def has_unique_chars(self, string):
      if string is None:
         return False 
      
      for char in string:
         if string.count(char) > 1:
            return False 
         
      return True

class TestUniqueChars(unittest.TestCase):
    def test_unique_chars(self, func):
        self.assertEqual(func(None), False)
        self.assertEqual(func(''), True)
        self.assertEqual(func('foo'), False)
        self.assertEqual(func('bar'), True)
        print('Success: test_unique_chars')


def main():
    test = TestUniqueChars()
    unique_chars = UniqueChars()
    test.test_unique_chars(unique_chars.has_unique_chars)
    try:
        unique_chars_set = UniqueCharsSet()
        test.test_unique_chars(unique_chars_set.has_unique_chars)
        unique_chars_in_place = UniqueCharsInPlace()
        test.test_unique_chars(unique_chars_in_place.has_unique_chars)
    except NameError:
        pass

if __name__ == '__main__':
    main()
