package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
)

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide host:port.")
                return
        }

        port_connexion := arguments[1]
        demande, err := net.Dial("tcp", port_connexion)
	if err != nil {
                fmt.Println(err)
                return
        }

        for {
		reponse, _ := bufio.NewReader(os.Stdin).ReadString('\n')
                fmt.Fprintf(demande, reponse+"\n")

                message, _ := bufio.NewReader(demande).ReadString('\n')
                fmt.Print("->: " + message)
                if strings.TrimSpace(string(reponse)) == "STOP" {
                        fmt.Println("TCP client exiting...")
                        return
                }
        }
}
    
