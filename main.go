package main

import (
	"ExemploControllers/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/soma", controllers.SomaController)
	http.HandleFunc("/subtracao", controllers.SubtracaoController)
	http.HandleFunc("/multiplicacao", controllers.MultiplicacaoController)
	http.HandleFunc("/divisao", controllers.DivisaoController)
	http.ListenAndServe(":8080", nil)
}
