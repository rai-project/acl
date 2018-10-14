package acl

import "time"

// easyjson:json
type ACL struct {
	NetworkAccess bool     `json:"network_access,omitempty"`
	MountAccess   bool     `json:"mount_access,omitempty"`
	QueueAccess   []string `json:"queue_access,omitempty"`
	Limit         Limit    `json:"limit,omitempty"`
}

// easyjson:json
type Limit struct {
	Runtime         time.Duration `json:"runtime,omitempty"`
	Storage         uint64        `json:"storage,omitempty"`
	MountStorage    uint64        `json:"mount_storage,omitempty"`
	NumberOfGPUs    uint64        `json:"number_of_gpus,omitempty"`
	CPUArchitecture string        `json:"cpu_architecture,omitempty"`
}
