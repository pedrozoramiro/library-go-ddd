package inventory

type Repository interface {
	FindAllBooks() []Book
}
