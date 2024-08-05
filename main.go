/*
1  - Boas vindas, este é um exemplo de CRUD usando Golang e Postgres
2 - Selecione uma opcao abaixo:
	1 - Criar um novo registro
	2 - Verificar registros existentes
	3 - Selecionar um determinado registro
	4 - Deletar um registro existente
	0 - Sair da aplicacao
Desenvolvido por Fernando Fraga para fins de portfólio.
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	escolha := showmenu()

	switch escolha {
	default:
		os.Exit(1)
	case 1:
		createRecord()
	case 2:
		selectAll()
	case 3:
		selectOne()
	case 4:
		deleteRecord()
	case 5:
		os.Exit(1)
	}

}

func createRecord() {}
func selectAll()    {}
func selectOne()    {}
func deleteRecord() {}

func showmenu() int {
	fmt.Println("Boas vindas, este é um exemplo de CRUD usando Golang e Postgres!")
	fmt.Println(`Selecione uma opção abaixo:
	1 - Criar um novo registro
	2 - Verificar registros existentes
	3 - Selecionar um determinado registro
	4 - Deletar um registro existente
	0 - Sair da aplicacao`)

	var escolha int

	fmt.Scan(&escolha)

	return escolha

}
