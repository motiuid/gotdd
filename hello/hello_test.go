package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func (t *testing.T, got, want string)  {
		t.Helper() // This function helps to tell the exact error location instead of throwing inside the helper fucntpion.
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Saying hello to people", func (t *testing.T)  {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("Say 'Hello, World' when an empty string is passed", func (t *testing.T)  {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func (t *testing.T)  {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func (t *testing.T)  {
		got := Hello("Vladimir", "French")
		want := "Bonjour, Vladimir"
		assertCorrectMessage(t, got, want)
	})
}