


package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    listener, err := net.Listen("tcp", ":80")
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    fmt.Println("Servidor escuchando en el puerto 80")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error al aceptar conexión:", err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    currentTime := time.Now().Format(time.RFC3339)
    conn.Write([]byte(currentTime))
}
