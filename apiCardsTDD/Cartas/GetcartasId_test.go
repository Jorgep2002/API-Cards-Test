package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"fmt"
)

//fetch al endpoint
func getCardById(cardID string) ([]byte, int, error) {
	url := "http://cards.thenexusbattles2.cloud/api/cartas/" + cardID
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	// Leer la respuesta en bytes
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return body, resp.StatusCode, nil
}

func TestGetCardById(t *testing.T) {
	t.Log("Prueba de la petición para traer una carta por ID")
	//ID de una carta existente en el sistema
	cardID := "650f390b7aaeb67f7dfc7140"
	responseBody, statusCode, err := getCardById(cardID)	
	//Probar  la solicitud HTTP
	if err != nil {
		t.Fatalf("Error al realizar la solicitud: %v", err)
	}
	//Probar solicitud exitosa (status code 200)
	if statusCode != http.StatusOK {
		t.Errorf("Código de estado esperado %d, pero obtuvo %d", http.StatusOK, statusCode)
	}
	//Probar tipo dato retornado
	var jsonData map[string]interface{}
	if err := json.Unmarshal(responseBody, &jsonData); err == nil {
		t.Log("La respuesta es un objeto JSON válidoS.")
	} else {
		t.Fatalf("Error al analizar la respuesta JSON: %v", err)
	}
	//Exito
	if err == nil && statusCode ==http.StatusOK {
		t.Logf("Prueba exitosa estado %d", statusCode)
	}
}

//Probar distintos tipos de error en la API
func TestCaseGetCardById(t *testing.T) {
    testCases := []struct {
        idCard   interface{} 
        Expected int        
    }{
		//El valor es de tipo numerico cuando se recibe un string
        {idCard: 12351231, Expected: http.StatusBadRequest},     
		// El id de la carta no es un id valido(no existe una carta con ese id)
        {idCard: "aeo12k3wawe", Expected: http.StatusBadRequest},
		//El id no fue dado 
		{idCard: "", Expected: http.StatusBadRequest},
    }
    for _, tc := range testCases {
        idCard := fmt.Sprintf("%v", tc.idCard) 
        _, statusCode, err := getCardById(idCard)

        if err == nil && statusCode == tc.Expected {
            t.Logf("Prueba de caso de error exitosa se manejo con el código de estado %d", statusCode)
        } else if err != nil && statusCode == tc.Expected {
            t.Errorf("Error inesperado: %v", err)
        } else if err == nil && statusCode != tc.Expected {
            t.Errorf("Código de estado esperado %d, pero obtuvo %d", tc.Expected, statusCode)
        } else {
            t.Logf("Solicitud con error %d, como se esperaba", tc.Expected)
        }
    }
}