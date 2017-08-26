type service struct{}

func (s *service) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil
}

func main() {
	s := grpc.NewServer()
	pb.RegisterContractServer(s, &service{})

	l, _ := net.Listen("tcp", ADDR)
	s.Serve(l)
}
