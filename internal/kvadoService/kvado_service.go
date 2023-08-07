package kvadoService

import (
	"context"
	"database/sql"
	pb "github.com/poggerr/kvado/api/kvado"
)

// Реализация GRPC-сервера
type Server struct {
	pb.YourServiceServer
	db *sql.DB // Объект базы данных
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		db: db,
	}
}

// Метод для поиска книг по автору
func (s *Server) FindBooksByAuthor(ctx context.Context, req *pb.AuthorRequest) (*pb.BookResponse, error) {
	// Выполнение SQL-запроса к базе данных
	rows, err := s.db.Query("SELECT Id, Name FROM Book WHERE AuthorId = ?", req.AuthorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Создание ответа
	resp := &pb.BookResponse{}
	for rows.Next() {
		var bookId int32
		var bookName string
		if err := rows.Scan(&bookId, &bookName); err != nil {
			return nil, err
		}
		book := &pb.Book{
			BookId:   bookId,
			BookName: bookName,
		}
		resp.Books = append(resp.Books, book)
	}

	return resp, nil
}

// Метод для поиска авторов по книге
func (s *Server) FindAuthorsByBook(ctx context.Context, req *pb.BookRequest) (*pb.AuthorResponse, error) {
	// Выполнение SQL-запроса к базе данных
	rows, err := s.db.Query("SELECT Author.Id, Author.Name FROM Author JOIN Book ON Author.Id = Book.AuthorId WHERE Book.Name = ?", req.BookName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Создание ответа
	resp := &pb.AuthorResponse{}
	for rows.Next() {
		var authorId int32
		var authorName string
		if err = rows.Scan(&authorId, &authorName); err != nil {
			return nil, err
		}
		author := &pb.Author{
			AuthorId:   authorId,
			AuthorName: authorName,
		}
		resp.Authors = append(resp.Authors, author)
	}

	return resp, nil
}
