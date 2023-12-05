package main

import (
  "fmt"
  "net"
  "os"
)

func main() {
  
  var port string 
  
  //Check for correct number of cli arguments
  if len(os.Args) == 2 {
   port = os.Args[1]
  } else {
    //default port is 1234
    port = "1234"
  }

  //set up the listener
  listener, err := net.Listen("tcp", ":"+port)
  if err != nil {
    fmt.Println("Error setting up listener: ", err)
    os.Exit(1)
  }

  defer listener.Close()
  fmt.Println("Server listening on port: ", port)

  for {
    
    conn, err := listener.Accept()
    if err != nil {
      fmt.Println("Error accepting connection: ", err)
      continue
    }

    go handleConnection(conn)

  }

}

func handleConnection(conn net.Conn) {
  defer conn.Close()

  buffer := make([]byte, 1024)

  for {
    n, err := conn.Read(buffer)
    if err != nil {
      if err == io.EOF {
        fmt.Println("Connection closed by the sender.")
        break
      }
      fmt.Println("Error reading message: ", err)
      break
    }
    fmt.Printf("Received message: %s\n", buffer[:n])
  }
}
