package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"

)

func main() {
	listenForInput()

}

func listenForInput() {
	printCommands()
	fmt.Println("O que deseja fazer?")

	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	errHandle(err)
	switch string(line) {
	case "quit":
		os.Exit(1)
	case "create":
		startGenerating(reader)
	case "help":
		printCommands()
	}
}

func errHandle(err error) {
	if err != nil {
		fmt.Errorf("Some error %s", err)
	}
}

func printCommands() {
	fmt.Println("Está é a lista de comandos:")
	fmt.Println("help")
	fmt.Println("create")
	fmt.Println("quit")
}

func startGenerating(reader *bufio.Reader) {
	fmt.Println("Qual tabela deseja a tabela?")
	tableName,_, err := reader.ReadLine()
	errHandle(err)

	insert := strings.Join([]string{"CREATE TABLE ", string(tableName)}, " ( ")
	inserts := []byte(insert)
	fmt.Println(insert)
	
	fmt.Println("Quais colunas?")
	fmt.Println("id, name, lastname, ...")
	attributes,_, err2 := reader.ReadLine()
	attributes = append(attributes, ')')
	errHandle(err2)

	


	numberOfAttributes := 1
	for i := 0; i < len(attributes); i++ {
		if attributes[i] == ',' {
			numberOfAttributes++
		}
	}

	values := []string {}
	for j := 0; j < numberOfAttributes; j++ {
		fmt.Println("Escolha o tipo da coluna", j+1, "")
		value, _, err := reader.ReadLine()
		errHandle(err)
		value = append(value, ',')
		values = append(values, string(value))
	}
	lastValue := values[len(values)-1]
	s := strings.Replace(lastValue, ",", "", -1)
	values[len(values)-1] = s


	writeToFile(inserts, values)

	fmt.Println(inserts)

}
func writeToFile(insert []byte, values []string) {
	file, err := os.Create("sql/deila.sql")
	errHandle(err)

	file.Write(insert)
	numberOfEntries :=1
	val := strings.Join(values, "")
	for i := 0; i < numberOfEntries; i++ {
		if i == numberOfEntries-1 {
			file.Write([]byte("("))
			file.Write([]byte(val))
			file.Write([]byte(");"))
		} else {
			file.Write([]byte("("))
			file.Write([]byte(val))
			file.Write([]byte("), "))
		}
	}

}