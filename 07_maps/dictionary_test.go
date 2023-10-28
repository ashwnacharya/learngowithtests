package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want:= "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}


func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition:= "this is just a test"


	dictionary.Add(word, definition)

	want := "this is just a test"
	got, _ := dictionary.Search("test")

	assertString(t, got, want)
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition:= "this is an updated test"
	
		dictionary := Dictionary{word: definition}
		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		newDefinition:= "this is an updated test"
	
		dictionary := Dictionary{}
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err:= dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}


func assertDefinition(t testing.TB, d Dictionary, word, definition string) {

	t.Helper()

	got, err := d.Search(word)

	if err != nil {
		t.Fatal("should find added word", err)
	}

	assertString(t, got, definition)
}


func assertString(t testing.TB, got string, want string) {

	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
