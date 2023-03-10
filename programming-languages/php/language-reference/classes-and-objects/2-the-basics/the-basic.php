<?php 

/**
 * Basic class definitions begin with the keyword class, 
 * followed by a class name, followed by a pair of curly braces which enclose the definitions of the properties and methods belonging to the class.
 * 
 * The class name can be any valid label, provided it is not a PHP reserved word. 
 * A valid class name starts with a letter or underscore, followed by any number of letters, numbers, or underscores. 
 * As a regular expression, it would be expressed thus: ^[a-zA-Z_\x80-\xff][a-zA-Z0-9_\x80-\xff]*$.
 * 
 * A class may contain its own constants, 
 * variables (called "properties"),
 * and functions (called "methods").
 */

// Simple class definition

class SimpleClass 
{
  // property declaration
  public $var = "a default value";

  // method declaration
  public function displayVar() {
    echo $this->var;
  }
}

$instance = new SimpleClass();

// This can also be done with a variable
$className = "SimpleClass";
$instance = new $className(); // new SimpleClass()

$instance->displayVar();

