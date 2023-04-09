package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindByUserID(userID int) ([]Book, error)
	FindByID(ID int) (Book, error)
	Save(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(ID int) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *repository) FindByUserID(userID int) ([]Book, error) {
	var books []Book
	err := r.db.Where("user_id = ?", userID).Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *repository) FindByID(ID int) (Book, error) {
	var book Book
	err := r.db.Where("id = ?", ID).Find(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) Delete(ID int) (Book, error) {
	var book Book

	err := r.db.First(&book, ID).Error
	if err != nil {
		return book, err
	}
	deletedBook := book

	err = r.db.Where("id = ?", ID).Delete(&book).Error
	if err != nil {
		return deletedBook, err
	}
	return deletedBook, nil
}

func (r *repository) Save(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error
	if err != nil {
		return book, err
	}
	return book, nil
}
