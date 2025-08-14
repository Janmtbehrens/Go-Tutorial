package maps

import "testing"

func TestSearch(t *testing.T){
	message := "Just testing bruh"
	needle := "test"
	dictionary := Dictionary{needle:message}

	t.Run("Known word", func(t *testing.T){
		got, _ := dictionary.Search(needle)
		want := message

		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T){
		_, err := dictionary.Search("null")
		want := ErrNotInDictionary

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"Existing Key":"Existing Value"}

	t.Run("Adding New" , func(t *testing.T){
		err := dictionary.Add("New Key","New Value")
		got := dictionary["New Key"]
		want := "New Value"

		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("Adding existing", func(t *testing.T){
		err := dictionary.Add("Existing Key","Existing Value")

		assertError(t, err, ErrAlreadyInDictionary)
	})
}

func TestUpdate(t *testing.T){
	dictionary := Dictionary{"Existing Key":"Existing Value"}

	t.Run("Updating not existing key" , func(t *testing.T){
		err := dictionary.Update("New Key","New Value")

		assertError(t, err, ErrUpdatedKeyDoesntExist)
	})

	t.Run("Updating existing key", func(t *testing.T){
		err := dictionary.Update("Existing Key","New Value")
		got := dictionary["Existing Key"]
		want := "New Value"

		assertNoError(t, err)
		assertStrings(t, got, want)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got DictionaryErr) {
	t.Helper()
	if got != "" {
		t.Fatal("got an error but didn't want one")
	}
}
func assertError(t testing.TB, got, want DictionaryErr) {
	t.Helper()
	if got == "" {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}