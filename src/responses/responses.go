package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma reposta em json para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// Erro retorna um erro em formato json
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
