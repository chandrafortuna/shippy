package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/chandrafortuna/shippy/consignment-service/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5050"
)

//Repository consigment interface
type repository interface {
	Create(pb *pb.Consignment) (*pb.Consignment, error)
}

//Repository consigment data store
type Repository struct {
	mu          sync.Mutex
	consigments []*pb.Consignment
}

// Create new consigment
func (repo *Repository) Create(consigment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consigments, consigment)
	repo.consigments = updated
	repo.mu.Unlock()
	return consigment, nil
}

//GetAll consignment items
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consigments
}

type service struct {
	repo *Repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consigment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consigment}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()

	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterShippingServiceServer(s, &service{repo})

	reflection.Register(s)

	log.Println("Running on port", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
