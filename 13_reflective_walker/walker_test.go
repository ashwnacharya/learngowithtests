package reflectivewalker

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

func TestWalk(t *testing.T) {
	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"chris"},
			[]string{"chris"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age int
			}{"chris", 33},
			[]string{"chris"},
		},
		{
			"struct with nested fields",
			struct {
				Name string
				Profile struct {
					Age int
					City string
				}
			}{
				"Chris",
				struct{Age int; City string}{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointer to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string {"Chris", "London"},
		},
		{
			"slices",
			[]Profile {
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string { "London", "Reykjavik"},
		},
		{
			"arrays",
			[2]Profile {
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string { "London", "Reykjavik"},
		},
	}

	for _, test := range cases {
		
		t.Run(test.Name, func(t *testing.T) {
			got := []string{}

			WalkFunction := func(name string) {
				got = append(got, name)
			}

			Walker(test.Input, WalkFunction)

			want := test.ExpectedCalls

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %q want %q", got, want)
			}
		})

		t.Run("With maps", func(t *testing.T) {
			aMap := map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			}
		
			var got []string
			Walker(aMap, func(input string) {
				got = append(got, input)
			})
		
			assertContains(t, got, "Moo")
			assertContains(t, got, "Baa")
		
		})

		t.Run("with channels", func(t *testing.T) {
			aChannel := make(chan Profile)
		
			go func() {
				aChannel <- Profile{33, "Berlin"}
				aChannel <- Profile{34, "Katowice"}
				close(aChannel)
			}()
		
			var got []string
			want := []string{"Berlin", "Katowice"}
		
			Walker(aChannel, func(input string) {
				got = append(got, input)
			})
		
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})

		t.Run("with function", func(t *testing.T) {
			aFunction := func() (Profile, Profile) {
				return Profile{33, "Berlin"}, Profile{34, "Katowice"}
			}
		
			var got []string
			want := []string{"Berlin", "Katowice"}
		
			Walker(aFunction, func(input string) {
				got = append(got, input)
			})
		
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}