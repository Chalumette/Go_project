//Nécessite un argument : le numéro du port TCP
package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
	"io"
)
func handle_connection (connexion net.Conn){	
	for {
		io.WriteString(connexion, "Chargez un fichier ? O/N\n")
		netData, err := bufio.NewReader(connexion).ReadString('\n')
		if err != nil {
                      fmt.Println(err)
                      return
	        }
		var reponse string 
		reponse = strings.TrimSpace(string(netData))
		fmt.Println(reponse)
		switch reponse{
		case "o","O":
			fmt.Println("En attente de fichier...\n")
			io.WriteString(connexion, "Chargemenet de fichier à implémenter\n")
		case "STOP":
			fmt.Println ("Fermeture de la connexion")
			return
		case "N", "n":
			io.WriteString(connexion,"Aurevoir\n")
			return
		default:
			io.WriteString(connexion,"Réponse invalide\n")
		}

	}
	connexion.Close()
}
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
	for {
		connexion, err := auditeur_reseau.Accept()
        	if err != nil {
                	fmt.Println(err)
                	return
        	}
		fmt.Println("Connexion établie")
        	go handle_connection (connexion)
		
	}
}
    
