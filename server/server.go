package main

import (
	"io"
	"log"
	"net"

	pb "github.com/gocs/ymmr/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (s server) Move(srv pb.PedalService_MoveServer) error {

	log.Println("start new server")
	ctx := srv.Context()
	pedal := &pb.Pedal{}

	resp := pb.MovePedal{Pedal: pedal}
	if err := srv.Send(&resp); err != nil {
		log.Printf("send error %v", err)
	}

	for {

		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		// update max and send it to stream
		pedal = req.Pedal
		resp := pb.MovePedal{Pedal: pedal}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Println("pedal:", req.GetPedal())
	}
}

func main() {
	// create listiner
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create grpc server
	s := grpc.NewServer()
	pb.RegisterPedalServiceServer(s, server{})

	log.Println("start...")
	// and start...
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
