package dictionary

import (
	"github.com/catarinabombaca/learngowithtests/testhelpers"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		testhelpers.AssertCorrectString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrNotFound.Error()

		testhelpers.AssertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		definition := "this is just a test"
		dictionary.Add("test", definition)

		got, err := dictionary.Search("test")

		testhelpers.AssertNoError(t, err)
		testhelpers.AssertCorrectString(t, got, definition)
	})

	t.Run("word already exists", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		want := ErrWordExists
		testhelpers.AssertError(t, err, want.Error())
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"

	err := dictionary.Update(word, newDefinition)
	testhelpers.AssertNoError(t, err)

	got, _ := dictionary.Search(word)
	testhelpers.AssertCorrectString(t, got, newDefinition)
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	testhelpers.AssertError(t, err, ErrNotFound.Error())
}
