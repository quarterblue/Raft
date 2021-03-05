package raft

import "sync"

type Raft struct {
	mu sync.Mutex
}
