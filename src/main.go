package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fabbboy/graphql-test/src/gen"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	r := gin.Default()

	// Create a new GraphQL server
	srv := handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{Resolvers: &gen.Resolver{}}))

	// Add WebSocket support
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	// Serve the GraphQL Playground UI at the root path
	r.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	// Serve GraphQL queries and mutations over HTTP POST
	r.POST("/query", gin.WrapH(srv))

	// Serve GraphQL subscriptions over WebSockets
	r.GET("/query", gin.WrapH(srv))

	// Run the server
	r.Run()
}
