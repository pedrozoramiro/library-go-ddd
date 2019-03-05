package hardcode

import (
	"library-go-ddd/domain/inventory"
)

type BookRepository struct{}

func (BookRepository) FindAllBooks() []inventory.Book {
	books := []inventory.Book{
		inventory.Book{Name: "Domain Driven Design", Author: "Eric J. Evans"},
		inventory.Book{Name: "implementing Domain Driven design", Author: "Vaughn Vernon"},
		inventory.Book{Name: "Domain-Driven Design Using Naked Objects", Author: "Dan Haywood"},
		inventory.Book{Name: "Applying Domain-Driven Design and Patterns", Author: "Jimmy Nilsson"},
		inventory.Book{Name: "Patterns, Principles and Practices of Domain-Driven Design", Author: "Scott Millett"},
	}
	return books
}
