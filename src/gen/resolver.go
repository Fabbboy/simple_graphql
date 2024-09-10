package gen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Users     []*User
	Messages  []*Message
	Listeners []chan *Message
}
