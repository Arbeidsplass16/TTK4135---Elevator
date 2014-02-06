package main

import(
    "net"
    "fmt"
)

func main(){
    out_socket, in_socket := connect()
    out_socket.Write([]byte("Hello socket\x00"))
}
