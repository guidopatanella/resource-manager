package rmanager

import "time"

// Resource - general resource interface
type Resource interface {
	SetID(string)
	GetID() string
	SetProgress(float32) 
	GetProgress() float32
	IsAvailable() bool
	Start(time.Duration) error
}
