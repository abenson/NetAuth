package client

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/grpc"

	pb "github.com/NetAuth/NetAuth/proto"
)

func NewClient(server string, port int) (pb.NetAuthClient, error) {
	// Setup the connection and defer the close.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", server, port), grpc.WithInsecure())

	// Create a client to use later on.
	return pb.NewNetAuthClient(conn), err
}

func Ping(server string, port int, clientID string) (string, error) {
	request := new(pb.PingRequest)
	request.ClientID = ensureClientID(clientID)

	client, err := NewClient(server, port)
	if err != nil {
		return "", err
	}
	pingResult, err := client.Ping(context.Background(), request)
	return pingResult.GetMsg(), nil
}

func Authenticate(server string, port int, clientID string, serviceID string, entity string, secret string) (string, error) {
	e := new(pb.Entity)
	e.ID = &entity
	e.Secret = &secret

	request := new(pb.NetAuthRequest)
	request.Entity = e
	request.ClientID = ensureClientID(clientID)
	request.ServiceID = ensureServiceID(serviceID)

	c, err := NewClient(server, port)
	if err != nil {
		return "", err
	}
	authResult, err := c.AuthEntity(context.Background(), request)
	if err != nil {
		return "", err
	}

	return authResult.GetMsg(), nil
}

// NewEntity makes a request to the server to add a new entity.  It
// requires an existing entity for authentication and authorization to
// add the new one, as well as parameters to populate the core fields
// on the new entity.  This function returns a string message from the
// server and an error describing whether or not the server was able
// to add the requested entity.
func NewEntity(server string, port int, clientID string, serviceID string, entity string, secret string, newEntity string, newUIDNumber int32, newSecret string) (string, error) {
	// e is the entity that is requesting this change.  This
	// entity must have the correct capabilities to actually add a
	// new entity to the system.
	e := new(pb.Entity)
	e.ID = &entity
	e.Secret = &secret

	// ne is the new entity.  These fields are the ones that must
	// be set at the time of creation for a new entity.
	ne := new(pb.Entity)
	ne.ID = &newEntity
	ne.UidNumber = &newUIDNumber
	ne.Secret = &newSecret

	request := new(pb.ModEntityRequest)
	request.Entity = e
	request.ModEntity = ne

	c, err := NewClient(server, port)
	if err != nil {
		return "", err
	}

	result, err := c.NewEntity(context.Background(), request)
	if err != nil {
		return "", err
	}

	return result.GetMsg(), nil
}

// RemoveEntity makes a request to the server to remove the named
// entity.  This must be authorized by an entity which has the
// appropriate capabilities to fulfill the remove request
func RemoveEntity(server string, port int, clientID string, serviceID string, entity string, secret string, delEntity string) (string, error) {
	// e is the entity requesting this change, it must have the
	// correct permissions to run the remove.
	e := new(pb.Entity)
	e.ID = &entity
	e.Secret = &secret

	// re is the entity to be removed
	re := new(pb.Entity)
	re.ID = &delEntity

	request := new(pb.ModEntityRequest)
	request.Entity = e
	request.ModEntity = re

	c, err := NewClient(server, port)
	if err != nil {
		return "", err
	}

	result, err := c.RemoveEntity(context.Background(), request)
	if err != nil {
		return "", err
	}

	return result.GetMsg(), nil
}

func ensureClientID(clientID string) *string {
	if clientID == "" {
		hostname, err := os.Hostname()
		if err != nil {
			clientID = "BOGUS_CLIENT"
			return &clientID
		}
		clientID = hostname
	}
	return &clientID
}

func ensureServiceID(serviceID string) *string {
	if serviceID == "" {
		serviceID = "BOGUS_SERVICE"
	}
	return &serviceID
}
