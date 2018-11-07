package acl

import (
	"math"
	"sync"
	"time"

	humanize "github.com/dustin/go-humanize"
)

var ACLs sync.Map

func init() {
	defer getRoles()
	ACLs.Store(Role("ECE408"),
		ACL{
			MountAccess:   false,
			NetworkAccess: false,
			DockerPush:    false,
			QueueAccess:   []string{"rai-raiders"},
			Limit: Limit{
				Runtime:         5 * time.Minute,
				Storage:         10 * humanize.MiByte,
				MountStorage:    0,
				NumberOfGPUs:    1,
				CPUArchitecture: "amd64",
			},
		},
	)
	ACLs.Store(Role("student"),
		ACL{
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
		},
	)
	ACLs.Store(Role("guest"),
		ACL{
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
		},
	)
	ACLs.Store(Role("power"),
		ACL{
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
		},
	)
	ACLs.Store(Role("admin"),
		ACL{
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
		},
	)
}
