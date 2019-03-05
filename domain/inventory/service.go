package inventory

type Service interface {
	FindAllBooks() []Book
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
