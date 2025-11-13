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

