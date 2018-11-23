package facade

import "fmt"

type (
	BookList    struct{}
	LendingList struct{}
	Librarian   struct {
		BookList    *BookList
		LendingList *LendingList
	}
)

func (booklist *BookList) SearchBook(bookName string) string {
	// search book and get location
	location := "15-F"
	return location
}

func (list *LendingList) Check(bookName string) bool {
	// check lending list
	return true
}

func NewLibrarian(booklist *BookList, lendinglist *LendingList) *Librarian {
	return &Librarian{
		BookList:    booklist,
		LendingList: lendinglist,
	}
}

func (librarian *Librarian) SearchBook(bookName string) string {
	location := librarian.BookList.SearchBook(bookName)

	if location != "" {
		isAvailable := librarian.LendingList.Check(bookName)
		if isAvailable {
			return "Available at " + location
		} else {
			return "Not available at " + location
		}
	} else {
		return "Not found"
	}
}

func Run() {
	bookName := "Zero to One"
	booklist := &BookList{}
	lendinglist := &LendingList{}
	librarian := NewLibrarian(booklist, lendinglist)
	status := librarian.SearchBook(bookName)
	fmt.Println(status)
}
