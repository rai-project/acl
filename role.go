package acl

import (
	"strings"

	"github.com/pkg/errors"
)

// roles are essentially permission groups
type Role string

const (
	All = "all"
)

var (
	// A list of all the roles. Currently, the roles
	// are defined statically within the program
	Roles = []Role{}
)

func getRoles() {
	ACLs.Range(func(key, value interface{}) bool {
		name, ok := key.(Role)
		if !ok {
			return true
		}
		Roles = append(Roles, Role(strings.ToLower(string(name))))
		return true
	})
}

func (r Role) ToACL() (ACL, error) {
	s := strings.ToLower(string(r))
	val, ok := ACLs.Load(s)
	if !ok {
		return ACL{}, errors.Errorf("the role %v is not valid. valid roles are %v", r, Roles)
	}
	acl, ok := val.(ACL)
	if !ok {
		return ACL{}, errors.Errorf("the role %v is not valid type. valid roles are %v", r, Roles)
	}
	return acl, nil
}

func (r0 Role) Validate() bool {
	r := strings.ToLower(string(r0))
	for _, e := range Roles {
		if string(e) == r {
			return true
		}
	}
	return false
}
