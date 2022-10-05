package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	file, err := os.Open("File.txt") // ouvre fichier.txt
	if err != nil {
		log.Fatalf("failed opening file: %s", err) // si le fichier n'existe pas, affiche une erreur
	}

	scanner := bufio.NewScanner(file) // scan le fichier
	scanner.Split(bufio.ScanLines)    // sépare les lignes du fichier
	var lines []string                // stocke les lignes du fichier la variable lines qui est un tableau de string

	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // ajoute les lignes du fichier dans la variable lines
	}

	// affiche la première ligne du fichier, à l'index 0, donc "generate"
	fmt.Println(lines[0])
	// affiche la dernière ligne du fichier,
	//en parcourant la longueur du tableau lines et en soustrayant 1 pour avoir l'index de la dernière ligne
	fmt.Println(lines[(len(lines) - 1)])

	// une boucle for range qui parcourt le tableau lines,
	//et qui stocke l'index et le mot dans la variable words
	for index, words := range lines {
		// si le mot est "before", alors le programme dans la condition
		// dans la condition, je déclare dynamiquement une variable value qui est égale à la ligne après "before"
		// la variable a est égale à la valeur de la variable value, mais en int avec la fonction strconv.Atoi
		// puis le programme affiche la ligne à l'index a-1
		if words == "before" {
			value := lines[index+1]
			a, _ := strconv.Atoi(value)
			fmt.Println(lines[a-1])
		}

		// si le mot est "now ", alors le programme entre dans la condition
		// dans la condition, je déclare dynamiquement une variable b qui est égale à la ligne avant "now"
		// la variable c est égale à la valeur de la variable b, mais converti en int()
		// divisé par la longueur du tableau lines
		// puis le programme affiche la ligne à l'index c-1
		if words == "now " {
			b := lines[index-1]
			c := int(b[1]) / len(lines)
			fmt.Println(lines[c-1])
		}

	}
	rand.Seed(time.Now().UnixNano()) // génére un nombre aléatoire, avec la fonction rand.Seed
	fmt.Println(rand.Intn(777))      // affiche un nombre aléatoire entre 0 et 777

	file.Close() // ferme le fichier
}
