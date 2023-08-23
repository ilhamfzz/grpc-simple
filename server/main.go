package main

import (
	"context"
	"encoding/json"
	pb "go-grpc-simple/student"
	"log"
	"net"
	"os"
	"sync"

	"google.golang.org/grpc"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*pb.Student
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}
	return nil, nil
}

func (d *dataStudentServer) loadData() {
	pwd := os.Getenv("PWD")
	data, err := os.ReadFile(pwd + "server/data/data.json")
	if err != nil {
		log.Fatalln("error in read file", err.Error())
	}

	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("error in unmarshal data json", err.Error())
	}
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	log.Println("server is running on port 8080")
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("error in listen", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error in serve grpc", err.Error())
	}
}
