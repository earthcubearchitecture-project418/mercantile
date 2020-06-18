package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fils/ocdGraphQL/graph"
	"github.com/fils/ocdGraphQL/graph/generated"
	"github.com/spf13/viper"
)

const defaultPort = "8080"

var viperVal string

func main() {
	var v1 *viper.Viper

	v1, err := readConfig(viperVal, map[string]interface{}{
		"address": "localhost",
		"port":    "8080",
	})
	if err != nil {
		panic(fmt.Errorf("error when reading config: %v", err))
	}

	mcfg := v1.GetStringMapString("server")

	port := os.Getenv(mcfg["port"])
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://0.0.0.0:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}
