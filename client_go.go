package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
	"io"
)

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide host:port.")
                return
        }

        port_connexion := arguments[1]
        connexion_serveur, err := net.Dial("tcp", port_connexion)//réception des données du serveur
	msg_affiche,_ := bufio.NewReader(connexion_serveur).ReadString('\n')//Convertit les données reçues en string
	fmt.Print(msg_affiche)
	if err != nil {
                fmt.Println(err)
                return
        }

        for {
		requete, _ := bufio.NewReader(os.Stdin).ReadString('\n')//on lis des entrées au clavier
		io.WriteString(connexion_serveur, requete)
		reponse_serveur,_:=bufio.NewReader(connexion_serveur).ReadString('\n')
		fmt.Print(reponse_serveur)
		if strings.TrimSpace(string(requete)) == "STOP" {
                        fmt.Println("TCP client exiting...")
                        return
                }
        }
}
    
