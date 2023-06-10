package Hosting_System_Plugin_Library

import "github.com/Cerberus-Labs-Technologies/Hosting-System-Plugin-Library/events"

type HostingListener struct {
	Name    events.ServerEvent
	Handler func() error
}

type HostingListenerInterface interface {
	GetName() string
	GetHandler() func() error
	Execute() error
}

func (hl *HostingListener) GetName() events.ServerEvent {
	return hl.Name
}

func (hl *HostingListener) GetHandler() func() error {
	return hl.Handler
}

func (hl *HostingListener) Execute() error {
	return hl.Handler()
}
