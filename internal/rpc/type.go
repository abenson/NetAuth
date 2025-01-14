package rpc

import (
	"github.com/NetAuth/NetAuth/internal/db"
	"github.com/NetAuth/NetAuth/internal/token"

	pb "github.com/NetAuth/Protocol"
)

// An EntityTree is a mechanism for storing entities and information
// about them.
type EntityTree interface {
	Bootstrap(string, string)
	DisableBootstrap()

	CreateEntity(string, int32, string) error
	FetchEntity(string) (*pb.Entity, error)
	SearchEntities(db.SearchRequest) ([]*pb.Entity, error)
	ValidateSecret(string, string) error
	SetSecret(string, string) error
	LockEntity(string) error
	UnlockEntity(string) error
	UpdateEntityMeta(string, *pb.EntityMeta) error
	UpdateEntityKeys(string, string, string, string) ([]string, error)
	ManageUntypedEntityMeta(string, string, string, string) ([]string, error)
	DestroyEntity(string) error

	CreateGroup(string, string, string, int32) error
	FetchGroup(string) (*pb.Group, error)
	SearchGroups(db.SearchRequest) ([]*pb.Group, error)
	UpdateGroupMeta(string, *pb.Group) error
	ManageUntypedGroupMeta(string, string, string, string) ([]string, error)
	DestroyGroup(string) error

	AddEntityToGroup(string, string) error
	RemoveEntityFromGroup(string, string) error
	ListMembers(string) ([]*pb.Entity, error)
	GetMemberships(*pb.Entity, bool) []string
	ModifyGroupExpansions(string, string, pb.ExpansionMode) error

	SetEntityCapability(string, string) error
	DropEntityCapability(string, string) error
	SetGroupCapability(string, string) error
	DropGroupCapability(string, string) error
}

// A NetAuthServer is a collection of methods that satisfy the
// requirements of the NetAuthServer protocol buffer.
type NetAuthServer struct {
	Tree  EntityTree
	Token token.Service
}
