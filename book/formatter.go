package book

type BookFormatter struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Publisher string `json:"publisher"`
	UserID    int    `json:"user_id"`
}

func FormatBook(book Book) BookFormatter {
	bookFormatter := BookFormatter{
		ID:        int(book.ID),
		Title:     book.Title,
		Year:      book.Year,
		Publisher: book.Publisher,
		UserID:    book.UserID,
	}
	return bookFormatter
}

func FormatBooks(books []Book) []BookFormatter {

	booksFormatter := []BookFormatter{}

	for _, book := range books {
		bookFormatter := FormatBook(book)
		booksFormatter = append(booksFormatter, bookFormatter)
	}
	return booksFormatter
}
