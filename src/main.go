package main

import (
  "fmt"
  "./config_reader"
  "./udp_server"
  )

func main() {
  config := config_reader.ReadConfigFile("./settings.json")

  response := make(chan string)
  if config["address"] == nil {
    fmt.Println("Need an address in settings.json")
  }

  udp_server.CreateServer(config["address"].(string), response)

  for {
    message := <-response

    switch(message) {
      case "exit":
        break
      default:
        fmt.Println(message)
    }
  }

}

