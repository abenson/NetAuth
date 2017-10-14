package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/NetAuth/NetAuth/proto"
)

var (
	port     = flag.Int("port", 8080, "Serving port, defaults to 8080")
	bindAddr = flag.String("bind", "localhost", "Bind address, defaults to localhost")
	useTLS   = flag.Bool("tls", false, "Enable TLS, off by default")
	certFile = flag.String("cert_file", "", "Path to certificate file")
	keyFile  = flag.String("key_file", "", "Path to key file")
)

type netAuthServer struct{}

func (s *netAuthServer) AuthEntity(ctx context.Context, entity *pb.Entity) (*pb.AuthResult, error) {
	log.Printf("Authenticating %s", entity.GetID())
	return &pb.AuthResult{}, nil
}

func (s *netAuthServer) EntityInfo(ctx context.Context, entity *pb.Entity) (*pb.EntityMeta, error) {
	return &pb.EntityMeta{}, nil
}

func (s *netAuthServer) Ping(ctx context.Context, PingRequest *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{}, nil
}

func newServer() *netAuthServer {
	return new(netAuthServer)
}

func main() {
	flag.Parse()

	log.Println("NetAuth server is starting!")

	// Bind early so that if this fails we can just bail out.
	sock, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *bindAddr, *port))
	if err != nil {
		log.Fatalf("could not bind! %v", err)
	}
	log.Printf("server bound on %s:%d", *bindAddr, *port)

	// Setup the TLS parameters if necessary.
	var opts []grpc.ServerOption
	if *useTLS {
		log.Printf("this server will use TLS with the certificate %s and key %s", *certFile, *keyFile)
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("TLS credentials could not be generated! %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	if ! *useTLS {
		// Not using TLS in an auth server?  For shame...
		log.Println("launching without TLS! Your passwords will be shipped in the clear!")
		log.Println("You should really start the server with -tls -key_file <keyfile> -cert_file <certfile>")
	}

	// Instantiate and launch.  This will block and the server
	// will server forever.
	log.Println("Server is launching...")
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSystemAuthServer(grpcServer, newServer())
	grpcServer.Serve(sock)
}
