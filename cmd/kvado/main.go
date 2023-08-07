package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	pb "github.com/poggerr/kvado/api/kvado"
	"github.com/poggerr/kvado/internal/kvadoService"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051" // Порт, на котором будет запущен сервер GRPC
)

func main() {
	db, err := sql.Open("mysql", "kvadoProto:password@tcp(127.0.0.1:3306)/kvadoProto")
	if err != nil {
		log.Fatalln(err)
	}

	serv := kvadoService.NewServer(db)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Регистрация сервиса GRPC
	pb.RegisterYourServiceServer(s, serv)

	// Запуск сервера GRPC
	log.Printf("GRPC server listening on %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
