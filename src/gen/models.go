// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gen

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	User    *User  `json:"user"`
}

type Mutation struct {
}

type Query struct {
}

type Subscription struct {
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
