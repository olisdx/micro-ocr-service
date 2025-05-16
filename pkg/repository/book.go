package repository

import (
	"context"
	"micro-ocr-service/api/presenter"
	"micro-ocr-service/pkg/entities"
	"micro-ocr-service/settings"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionName = "books"

// Repository interface allows us to access the CRUD Operations in mongo here.
type Repository interface {
	CreateBook(book *entities.Book) (*entities.Book, error)
	ReadBook(page int64) (*[]presenter.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	DeleteBook(ID string) error
}
type repository struct {
	Collection *mongo.Collection
}

// NewRepo is the single instance repo that is being created.
func NewRepo() Repository {

	bookCollection := settings.ConnectMongoDb(collectionName)

	return &repository{
		Collection: bookCollection,
	}
}

// CreateBook is a mongo repository that helps to create books
func (r *repository) CreateBook(book *entities.Book) (*entities.Book, error) {
	book.ID = primitive.NewObjectID()
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

// ReadBook is a mongo repository that helps to fetch books
func (r *repository) ReadBook(page int64) (*[]presenter.Book, error) {
	var books []presenter.Book

	limit, _ := strconv.ParseInt(os.Getenv("DATA_PER_PAGE"), 10, 64) // convert string to int

	filter := bson.D{{}} // selects all documents
	options := new(options.FindOptions)
	if limit != 0 {
		if page == 0 {
			page = 1
		}
		options.SetSkip(int64((page - 1) * limit))
		options.SetLimit(int64(limit))
	}

	cursor, err := r.Collection.Find(context.TODO(), filter, options)

	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var book presenter.Book
		_ = cursor.Decode(&book)
		books = append(books, book)
	}
	return &books, nil
}

// UpdateBook is a mongo repository that helps to update books
func (r *repository) UpdateBook(book *entities.Book) (*entities.Book, error) {
	book.UpdatedAt = time.Now()
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	if err != nil {
		return nil, err
	}
	return book, nil
}

// DeleteBook is a mongo repository that helps to delete books
func (r *repository) DeleteBook(ID string) error {
	bookID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}
	_, err = r.Collection.DeleteOne(context.Background(), bson.M{"_id": bookID})
	if err != nil {
		return err
	}
	return nil
}
