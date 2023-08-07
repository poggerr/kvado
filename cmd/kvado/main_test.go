package main

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/poggerr/kvado/api/kvado"
	"github.com/poggerr/kvado/internal/kvadoService"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindBooksByAuthor(t *testing.T) {
	// Создание mock базы данных
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание экземпляра сервера с mock базой данных
	server := kvadoService.NewServer(db)

	// Ожидаемый запрос к базе данных
	mock.ExpectQuery("SELECT Id, Name FROM Book WHERE AuthorId = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"Id", "Name"}).
			AddRow(1, "Book 1").
			AddRow(2, "Book 2"))

	// Выполнение метода FindBooksByAuthor
	req := &pb.AuthorRequest{AuthorId: 1}
	resp, err := server.FindBooksByAuthor(context.Background(), req)

	// Проверка результатов
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Books, 2)
	assert.Equal(t, int32(1), resp.Books[0].BookId)
	assert.Equal(t, "Book 1", resp.Books[0].BookName)
	assert.Equal(t, int32(2), resp.Books[1].BookId)
	assert.Equal(t, "Book 2", resp.Books[1].BookName)

	// Проверка mock базы данных
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestFindAuthorsByBook(t *testing.T) {
	// Создание mock базы данных
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Создание экземпляра сервера с mock базой данных
	server := kvadoService.NewServer(db)

	// Ожидаемый запрос к базе данных
	mock.ExpectQuery("SELECT Author.Id, Author.Name FROM Author JOIN Book ON Author.Id = Book.AuthorId WHERE Book.Name = ?").
		WithArgs("Book 1").
		WillReturnRows(sqlmock.NewRows([]string{"Id", "Name"}).
			AddRow(1, "Author 1").
			AddRow(2, "Author 2"))

	// Выполнение метода FindAuthorsByBook
	req := &pb.BookRequest{BookName: "Book 1"}
	resp, err := server.FindAuthorsByBook(context.Background(), req)

	// Проверка результатов
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Authors, 2)
	assert.Equal(t, int32(1), resp.Authors[0].AuthorId)
	assert.Equal(t, "Author 1", resp.Authors[0].AuthorName)
	assert.Equal(t, int32(2), resp.Authors[1].AuthorId)
	assert.Equal(t, "Author 2", resp.Authors[1].AuthorName)

	// Проверка mock базы данных
	assert.NoError(t, mock.ExpectationsWereMet())
}
