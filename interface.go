package Hosting_System_Plugin_Library

type Listener struct {
	Name    string
	Handler func() error
}
