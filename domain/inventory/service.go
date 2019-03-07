package inventory

type Service interface {
	FindAllBooks() []Book
	CreateBook(book Book)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) FindAllBooks() []Book {
	return s.repository.FindAllBooks()
}

func (s *service) CreateBook(book Book) {
	s.repository.CreateBook(book)
}
