package facade

import "testing"

func TestFacade(t *testing.T) {
	booklist := &BookList{}
	lendinglist := &LendingList{}
	librarian := NewLibrarian(booklist, lendinglist)

	result := librarian.SearchBook("Zero to One")
	expect := "Available at 15-F"

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.", expect, result)
	}
}
