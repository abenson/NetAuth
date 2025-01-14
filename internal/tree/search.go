package tree

import (
	"github.com/NetAuth/NetAuth/internal/db"

	pb "github.com/NetAuth/Protocol"
)

// SearchGroups returns a list of groups filtered by the search
// criteria.
func (m *Manager) SearchGroups(r db.SearchRequest) ([]*pb.Group, error) {
	return m.db.SearchGroups(r)
}

// SearchEntities returns a list of entities filtered by the search
// criteria.
func (m *Manager) SearchEntities(r db.SearchRequest) ([]*pb.Entity, error) {
	entities, err := m.db.SearchEntities(r)
	if err != nil {
		return nil, err
	}

	out := []*pb.Entity{}
	for i := range entities {
		out = append(out, entities[i])
	}
	return out, nil
}
