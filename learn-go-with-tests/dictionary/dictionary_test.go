package dictionary

import "testing"

func TestSearch(t *testing.T) {
    dictionary := Dictionary{"test": "This is just a test"}

    t.Run("should return value when key exists", func(t *testing.T) {
        got, _ := dictionary.Search("test")
        want := "This is just a test"
        assertStrings(t, got, want)
    })

    t.Run("should return empty string when key doesn't exist", func(t *testing.T) {
        _, err := dictionary.Search("tester")

        if err == nil {
            t.Fatal("Expected an error to be raised but got none")
        }

        assertError(t, err, KeyNotFoundError)
    })
}

func TestAdd(t *testing.T) {
    t.Run("should add item given new key", func(t *testing.T) {
        dictionary := Dictionary{}
        dictionary.Add("test", "this is just a test")

        want := "this is just a test"
        got, err := dictionary.Search("test")

        assertStrings(t, got, want)
        assertNoError(t, err)
    })

    t.Run("should not add item given existing key", func(t *testing.T) {
        dictionary := Dictionary{"test": "This is just a test"}
        err := dictionary.Add("test", "something else")

        assertError(t, err, DuplicateKeyError)
    })
}

func TestUpdate(t *testing.T) {
    t.Run("should update item given existing key", func(t *testing.T) {
        dictionary := Dictionary{"test": "This is just a test"}
        want := "Something else"
        dictionary.Update("test", want)

        got, err := dictionary.Search("test")

        assertStrings(t, got, want)
        assertNoError(t, err)
    })

    t.Run("should raise error given key not found", func(t *testing.T) {
        dictionary := Dictionary{"test": "This is just a test"}
        want := "Something else"
        err := dictionary.Update("tester", want)

        assertError(t, err, KeyNotFoundError)
    })
}

func assertStrings(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func assertError(t *testing.T, got, want error) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func assertNoError(t *testing.T, got error)  {
    t.Helper()
    if got != nil {
        t.Errorf("Unexpected error %s", got)
    }
}

