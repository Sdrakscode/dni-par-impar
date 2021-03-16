package main

import "fmt"

func main() {
	var dni int //iniamos al variable 
	fmt.Println("decime el ultimo digito de tu dni: ")
	fmt.Scanf("%d", &dni) //pide el ultimo numero del dni, y busca si su resto es 0 es par
	if dni%2 == 0 { //si su ultimo numero al dividirlo por dos su resto da 0 es par si no impoar
		fmt.Println("tu dni es par")
	} else {
		fmt.Println("tu dni es impar")
	}
} //fin del comando
