package repository

import "github.com/dirgadm/fithub-api/internal/common"

// ROption anything any repo object needed
type ROption struct {
	Common common.Options
}

// Repository all repo object injected here
type Repository struct {
	User IUser
}
