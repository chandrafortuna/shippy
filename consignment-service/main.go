package main

import (
	"context"
	"log"
	"sync"

	pb "github.com/chandrafortuna/shippy/consignment-service/proto/consignment"
	micro "github.com/micro/go-micro"
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

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()

	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("consignment"),
	)

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
