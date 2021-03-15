package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "Struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name:          "Nested fields",
			Input:         Person{Name: "Chris", Profile: Profile{33, "London"}},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	//
	//expected := "Chris"
	//var got []string
	//
	//x := struct {
	//	Name string
	//}{expected}
	//
	//walk(x, func(input string) {
	//	got = append(got, input)
	//})
	//
	//if len(got) != 1 {
	//	t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	//}
	//
	//if got[0] != expected {
	//	t.Errorf("got %q, want %q", got[0], expected)
	//}
}
