# basic data structures
# Two very common data structures that you will use as a Ruby programmer are arrays and hashes. 

# arrays
# An array is used to organize information into an ordered list. 
# The list can be made up of strings, integers, floats, booleans, or any other data type. 
# In Ruby, an array literal is denoted by square brackets [ ]. Inside the brackets you can create a list of elements separated by commas. Let's make one in irb.

arr = [1, 2, 3, 4, 5]
puts arr 
puts arr[0]

# hashes
# A hash, sometimes referred to as a dictionary, is a set of key-value pairs. 
# Hash literals are represented with curly braces { }. 
# A key-value pair is an association where a key is assigned a specific value. 

location = {
  :country => "Indonesia",
  "city" => "Jakarta",
  "code": "JKT"
}
puts location
puts location[:country]
puts location["country"] # can't access the key because string and symbol are different
puts location["city"]