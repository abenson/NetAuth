package entity_tree

import "errors"

var (
	E_DUPLICATE_ID        = errors.New("An entity with that ID already exists!")
	E_DUPLICATE_UIDNUMBER = errors.New("An entity with that uidNumber already exists!")
)