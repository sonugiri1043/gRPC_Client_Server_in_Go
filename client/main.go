package main

import (
  "log"

  "../api"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)

func main() {
  var conn *grpc.ClientConn

  // Create the client TLS credentials
  creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")
  if err != nil {
    log.Fatalf("could not load tls cert: %s", err)
  }

  // Initiate a connection with the server
  conn, err = grpc.Dial("localhost:7777", grpc.WithTransportCredentials(creds))
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()

  c := api.NewPingClient(conn)

  response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Response from server: %s", response.Greeting)
}