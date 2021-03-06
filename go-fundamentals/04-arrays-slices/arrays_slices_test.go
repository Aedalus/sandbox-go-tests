package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){
	t.Run("collection of 5 numbers", func(t *testing.T){
		numbers := [5]int{1,2,3,4,5}

		got := Sum(numbers[:])
		expected := 15

		if got != expected {
			t.Errorf("got %d expected %d", got, expected)
		}
	})
}

func TestSumAll(t *testing.T){

	got := SumAll([]int{1,2},[]int{0,9})
	want := []int{3,9}

	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v want %v", got, want)
	}
}
