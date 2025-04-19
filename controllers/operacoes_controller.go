package controllers

import (
	"ExemploControllers/models"
	"fmt"
	"net/http"
	"strconv"
)

func SomaController(w http.ResponseWriter, r *http.Request) {
	a, b, err := getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado := models.Somar(float64(a), float64(b))
	fmt.Fprintf(w, "Resultado da soma: %2f", resultado)
}

func SubtracaoController(w http.ResponseWriter, r *http.Request) {
	a, b, err := getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado := models.Subtracao(float64(a), float64(b))
	fmt.Fprintf(w, "Resultado da subtração: %2f", resultado)
}

func MultiplicacaoController(w http.ResponseWriter, r *http.Request) {
	a, b, err := getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado := models.Multiplicacao(float64(a), float64(b))
	fmt.Fprintf(w, "Resultado da multiplicação: %2f", resultado)
}

func DivisaoController(w http.ResponseWriter, r *http.Request) {
	a, b, err := getParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultado, err := models.Divisão(float64(a), float64(b))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Resultado da divisão: %2f", resultado)
}

// Função auxiliar para evitar repetição de código
func getParams(r *http.Request) (int, int, error) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err1 := strconv.Atoi(aStr)
	b, err2 := strconv.Atoi(bStr)

	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("parâmetros inválidos")
	}

	return a, b, nil
}
