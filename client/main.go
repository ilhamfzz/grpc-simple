package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "go-grpc-simple/student"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) /*(*pb.Student, error)*/ {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	student, err := client.FindStudentByEmail(ctx, &pb.Student{Email: email})
	if err != nil {
		log.Fatalln("error in call FindStudentByEmail", err.Error())
	}

	fmt.Println(student)
	// return student, nil
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatalln("error in dial", err.Error())
	}

	defer conn.Close()
	client := pb.NewDataStudentClient(conn)

	getDataStudentByEmail(client, "ilhamfz117@gmail.com")
	getDataStudentByEmail(client, "rizqizami0@gmail.com")
}
