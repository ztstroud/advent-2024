package main

import (
	"testing"
)

func TestSimple(t *testing.T) {
	input := "mul(2,3)"
	expected := 6
	actual := sumMuls(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestMulti(t *testing.T) {
	input := "mul(2,3)mul(1,4)mul(2,2)"
	expected := 14
	actual := sumMuls(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestGarbageExternal(t *testing.T) {
	input := "^)|mmul(1,1)x8("
	expected := 1
	actual := sumMuls(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestGarbageName(t *testing.T) {
	input := "mXul(1,1)"
	expected := 0
	actual := sumMuls(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestGarbageArgument(t *testing.T) {
	input := "mul(1X,1)"
	expected := 0
	actual := sumMuls(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

/*
The conditional tests use mul that produce powers of two. This means the binary
of the expected value tells you exactly which muls were included.
*/

func TestConditionalNoConditional(t *testing.T) {
	input := "mul(1,1)mul(1,2)"
	expected := 3
	actual := sumMulsConditional(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestConditionalOff(t *testing.T) {
	input := "don't()mul(1,1)"
	expected := 0
	actual := sumMulsConditional(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestConditionalOffOn(t *testing.T) {
	input := "don't()mul(1,1)do()mul(1,2)"
	expected := 2
	actual := sumMulsConditional(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

func TestConditionalOffOnOffOn(t *testing.T) {
	input := "mul(1,1)don't()mul(1,2)do()mul(1,4)don't()mul(1,8)do()mul(1,16)"
	expected := 21
	actual := sumMulsConditional(input)

	if expected != actual {
		t.Errorf("Expected %v to be %v", actual, expected)
	}
}

