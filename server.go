package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	"net/smtp"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Henoch-kargo/kargo-trucks/graph"
	"github.com/Henoch-kargo/kargo-trucks/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	from := "from@yahoo.com"
	password := "<Password>"

	to := []string{
		"sender@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("This is test email")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, from, to, message)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent successfully!")

	resolver := &graph.Resolver{}
	resolver.Init()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
