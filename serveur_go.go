//Serveur en go qui affiche l'heure 
//Nécessite un argument : le numéro du port TCP
package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
        "time"
)

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide port number")
                return
        }

        port_tcp := ":" + arguments[1]
        auditeur_reseau, err := net.Listen("tcp", port_tcp)
        if err != nil {
                fmt.Println(err)
                return
        }
	fmt.Println("En attente de connexion...")
	defer auditeur_reseau.Close() //ferme le serveur à la fin du programme ie arrêt de l'écoute

	connexion, err := auditeur_reseau.Accept()
        if err != nil {
                fmt.Println(err)
                return
        }
	fmt.Println("Connexion établie")

        for {
                fmt.Println("Chargez un fichier ? O/N")
		netData, err := bufio.NewReader(connexion).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }
                if strings.TrimSpace(string(netData)) == "STOP" {
                        fmt.Println("Exiting TCP server!")
                        return
                }

                fmt.Print("-> ", string(netData))
                t := time.Now()
                myTime := t.Format(time.RFC3339) + "\n"
                connexion.Write([]byte(myTime))
        }
}
    
