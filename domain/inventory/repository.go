package inventory

type Repository interface {
	FindAllBooks() []Book
	CreateBook(book Book)
}
