package reflection

import (
	"slices"
	"testing"
)

func TestWalk(t *testing.T) {

	t.Run("structs,slices,arrays,pointers", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"Andy"},
				[]string{"Andy"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{"Andy", "Bangor"},
				[]string{"Andy", "Bangor"},
			},
			{
				"struct with non string field",
				struct {
					Name string
					Age  int
				}{"Andy", 51},
				[]string{"Andy"},
			},
			{
				"struct with nested fields",
				Person{"Andy", Profile{51, "Bangor"}},
				[]string{"Andy", "Bangor"},
			},
			{
				"pointers to things",
				&Person{
					"Andy",
					Profile{51, "Bangor"},
				},
				[]string{"Andy", "Bangor"},
			},
			{
				"slices",
				[]Profile{
					{51, "Bangor"},
					{33, "Belfast"},
				},
				[]string{"Bangor", "Belfast"},
			},
			{
				"arrays",
				[2]Profile{
					{51, "Bangor"},
					{33, "Belfast"},
				},
				[]string{"Bangor", "Belfast"},
			},
			{
				"maps",
				map[string]string{
					"Cow":   "Moo",
					"Sheep": "Baa",
				},
				[]string{"Moo", "Baa"},
			},
		}

		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				var got []string
				walk(c.Input, func(input string) {
					got = append(got, input)
				})

				if !slices.Equal(got, c.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, c.ExpectedCalls)
				}
			})
		}
	})

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{51, "Bangor"}
			aChannel <- Profile{33, "Belfast"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Bangor", "Belfast"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{51, "Bangor"}, Profile{33, "Belfast"}
		}

		var got []string
		want := []string{"Bangor", "Belfast"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
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
		t.Errorf("expected %v to contain %s but it didn't", haystack, needle)
	}
}
