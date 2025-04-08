package reflection

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
		Input         any
		ExpectedCalls []string
	}{
		{
			Name: "struct with one field",
			Input: struct {
				Name string
			}{"Chrys"},
			ExpectedCalls: []string{"Chrys"},
		},
		{
			Name: "struct with two fields",
			Input: struct {
				Name string
				City string
			}{"Chrys", "London"},
			ExpectedCalls: []string{"Chrys", "London"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chrys", 30},
			ExpectedCalls: []string{"Chrys"},
		},
		{
			Name: "nested fields",
			Input: Person{
				"Chrys",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chrys", "London"},
		},
		{
			Name: "Pointers to things",
			Input: &Person{
				"Chrys",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chrys", "London"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			ExpectedCalls: []string{"London", "Reykjavik"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			var got []string
			walk(tc.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tc.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, tc.ExpectedCalls)
			}
		})
	}
}
