<?php 

class Test 
{
  static public function getNew() 
  {
    return new static;
  }
}

class Child extends Test {}

$obj1 = new Test();
$obj2 = new $obj1;
$obj3 =& $obj1;
var_dump($obj1 === $obj2);
var_dump($obj1 === $obj3);

$obj4 = Test::getNew();
var_dump($obj4 instanceof Test);

$obj5 = Child::getNew();
var_dump($obj5 instanceof Child);

