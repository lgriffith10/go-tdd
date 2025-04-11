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
		{
			Name: "arrays",
			Input: [2]Profile{
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baaa",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baaa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "London"}
			aChannel <- Profile{34, "Reykjavik"}
			close(aChannel)
		}()

		var got []string
		want := []string{"London", "Reykjavik"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()

	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)

	}
}
