package book

type Service interface {
	GetBooks(userID int) ([]Book, error)
	GetBookByID(input GetBookDetailInput) (Book, error)
	CreateBook(input CreateBookInput) (Book, error)
	UpdateBook(inputID int, inputData CreateBookInput) (Book, error)
	DeleteBook(inputID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetBooks(userID int) ([]Book, error) {
	if userID != 0 {
		books, err := s.repository.FindByUserID(userID)
		if err != nil {
			return books, err
		}
		return books, nil
	}

	books, err := s.repository.FindAll()
	if err != nil {
		return books, err
	}
	return books, nil
}

func (s *service) CreateBook(input CreateBookInput) (Book, error) {
	book := Book{}
	book.Title = input.Title
	book.Year = input.Year
	book.Publisher = input.Publisher
	book.UserID = int(input.User.ID)

	newBook, err := s.repository.Save(book)
	if err != nil {
		return newBook, err
	}
	return newBook, nil
}

func (s *service) GetBookByID(input GetBookDetailInput) (Book, error) {
	book, err := s.repository.FindByID(input.ID)

	if err != nil {
		return book, err
	}
	return book, nil
}

func (s *service) UpdateBook(inputID int, inputData CreateBookInput) (Book, error) {
	book, err := s.repository.FindByID(inputID)
	if err != nil {
		return book, err
	}
	book.Title = inputData.Title
	book.Year = inputData.Year
	book.Publisher = inputData.Publisher

	updatedBook, err := s.repository.Update(book)
	if err != nil {
		return updatedBook, err
	}
	return updatedBook, nil
}

func (s *service) DeleteBook(inputID int) (Book, error) {

	deletedBook, err := s.repository.Delete(inputID)
	if err != nil {
		return deletedBook, err
	}
	return deletedBook, nil
}
