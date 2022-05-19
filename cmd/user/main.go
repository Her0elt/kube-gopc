package main

import (
    pb "kube-gopc/internal/pb/user"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "kube-gopc/pkg/user/models"
)


const ds = "host=postgres user=root password=password dbname=kube_gopc port=5432"
var db *gorm.DB

func init() {
	var e error
	db, e = gorm.Open(postgres.Open(ds), &gorm.Config{})

	if e != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.User{})  //this will create a table in your database if it isn't there already
}

type server struct {
    pb.UnimplementedUserServiceServer
}

func (s *server) GetUserList(ctx context.Context, in *pb.GetUserListRequest) (*pb.GetUserListResponse, error){
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
    return &pb.GetUserListResponse {
        Users: GetUsers(),
    }, nil
}

func GetUsers() []*pb.User {
    var users []*pb.User
    db.Model(&models.User{}).Find(&users)
    return users
}


func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
    var user models.User;
    user.FromCreateRequest(in)
    db.Create(&user)
    return &pb.CreateUserResponse{
        User: user.GetGRPCModel(),
    }, nil
}





func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
