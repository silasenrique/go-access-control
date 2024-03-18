package main

import "go-access-control/src/internal/api"

func main() {
	err := api.NewServer(
		":9959",
		"postgres://user:password@localhost:5432/db?sslmode=disable",
	).LoadRoutes().Run()

	if err != nil {
		panic(err)
	}
}
