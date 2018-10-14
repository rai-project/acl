package acl

import (
	"math"
	"strings"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/pkg/errors"
)

type Role string

const (
	All = "all"
)

var Roles = []Role{
	"ece408_student",
	"student",
	"guest",
	"power",
	"admin",
}

func (r Role) ToACL() (ACL, error) {
	switch r {
	case "ece408_student":
		return ACL{
			MountAccess:   false,
			NetworkAccess: false,
			QueueAccess:   []string{"rai-raiders"},
			Limit: Limit{
				Runtime:         5 * time.Minute,
				Storage:         10 * humanize.MiByte,
				MountStorage:    0,
				NumberOfGPUs:    1,
				CPUArchitecture: "amd64",
			},
		}, nil
	case "student":
		return ACL{
			MountAccess:   false,
			NetworkAccess: false,
			QueueAccess:   []string{All},
			Limit: Limit{
				Runtime:         time.Hour,
				Storage:         humanize.GiByte,
				MountStorage:    humanize.GiByte,
				NumberOfGPUs:    math.MaxUint64,
				CPUArchitecture: All,
			},
		}, nil
	case "guest":
		return ACL{
			MountAccess:   true,
			NetworkAccess: false,
			QueueAccess:   []string{All},
			Limit: Limit{
				Runtime:         time.Hour,
				Storage:         humanize.GiByte,
				MountStorage:    humanize.GiByte,
				NumberOfGPUs:    math.MaxUint64,
				CPUArchitecture: All,
			},
		}, nil
	case "power":
		return ACL{
			MountAccess:   true,
			NetworkAccess: true,
			QueueAccess:   []string{All},
			Limit: Limit{
				Runtime:         10 * time.Hour,
				Storage:         humanize.GiByte,
				MountStorage:    10 * humanize.GiByte,
				NumberOfGPUs:    math.MaxUint64,
				CPUArchitecture: All,
			},
		}, nil
	case "admin":
		return ACL{
			MountAccess:   true,
			NetworkAccess: true,
			QueueAccess:   []string{All},
			Limit: Limit{
				Runtime:         math.MaxInt64,
				Storage:         math.MaxUint64,
				MountStorage:    math.MaxUint64,
				NumberOfGPUs:    math.MaxUint64,
				CPUArchitecture: All,
			},
		}, nil
	}
	return ACL{}, errors.Errorf("the role %v is not valid. valid roles are %v", r, Roles)
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
