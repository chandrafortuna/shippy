package main

import(
	"log"
	"context"
	"errors"

	pb "github.com/chandrafortuna/shippy/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

//VesselRepository repo
type VesselRepository struct {
	vessels []*pb.Vessel
}

//FindAvailable check specfication
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("No vessel found on that spec")
}

type service struct {
	repo repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}

	repo := &VesselRepository{vessels}

	srv := micro.NewService(micro.Name("vessel"))

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		log.Println(err)
	}

}
