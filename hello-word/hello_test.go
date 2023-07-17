package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		testName := "ExampleName"
		got := Hello(testName, "")
		want := "Hello, " + testName
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, Word' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hola, Elodie' if the language 'Spanish' and name 'Elodie' is specified", func(t *testing.T) {
		got := Hello("Elodie", "ES")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Bonjour' if the language 'French' and name 'Ela' is specified", func(t *testing.T) {
		got := Hello("Ela", "FR")
		want := "Bonjour, Ela"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// conscript this method as a helper; with this the reporting for failing tests
	// includes the line number of the function call, not the one from the helper
	t.Helper()

	if got != want {
		t.Errorf("\nreceived:\t%q\nwanted:\t%q", got, want)
	}
}
