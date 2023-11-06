package main

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

//fetch al endpoint
func getCartas(baseURL string, queryParams map[string]string) ([]byte, int, error) {
	url := baseURL + "?"
	for key, value := range queryParams {
		url += key + "=" + value + "&"
	}
	url = url[:len(url)-1]
	resp, err := http.Get(url)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return body, resp.StatusCode, nil
}


func TestFetchData(t *testing.T) {
	t.Log("Prueba de la petición...")
	baseURL := "https://cards.thenexusbattles2.cloud/api/cartas/"
	queryParams := map[string]string{
		"size":        "5",
		"page":        "1",
		"coleccion":   "Heroes",
		"onlyActives": "true",
	}
	responseBody, statusCode, err := getCartas(baseURL, queryParams)
	//Probar  la solicitud HTTP
	if err != nil {
		t.Fatalf("Error al realizar la solicitud: %v", err)
	}
	//Probar solicitud exitosa (status code 200)		
	if statusCode != http.StatusOK {
		t.Errorf("Código de estado esperado %d, pero obtuvo %d", http.StatusOK, statusCode)
	}
	// Probar tipo dato retornado (matriz/slice de objetos JSON)
	var jsonData []map[string]interface{} 
	if err := json.Unmarshal(responseBody, &jsonData); err != nil {
		t.Fatalf("Error al analizar la respuesta JSON: %v", err)
	} else if len(jsonData) == 0 {
		t.Fatal("La respuesta es una matriz de objetos JSON vacía.")
	} else {
		t.Log("La respuesta es una matriz de objetos JSON.")
	}
	//Exito
	if err == nil && statusCode ==http.StatusOK {
		t.Logf("Prueba exitosa estado %d", statusCode)
	}
}

//Funcion la cual sirve para probar distintos casos de error en la petición
func TestCasesFetchCartas(t *testing.T) {
	t.Log("Prueba de manejo de  parametros en la petición...")
	testCases := []struct {
		queryParams map[string]string 
		Expected    int             
	}{
		
		// Parámetros con valor de "page" inválido
		{
			queryParams: map[string]string{
				"size":        "5",
				"page":        "invalido",
				"coleccion":   "Heroes",
				"onlyActives": "true",
			},
			Expected: http.StatusBadRequest,
		},
	}

	baseURL := "https://cards.thenexusbattles2.cloud/api/cartas/"
	for _, tc := range testCases {
		_, statusCode, err := getCartas(baseURL, tc.queryParams)

		if err == nil && statusCode == tc.Expected {
			t.Logf("Prueba exitosa, código de estado %d", statusCode)
		} else if err != nil && statusCode == tc.Expected {
			t.Errorf("Error inesperado: %v", err)
		} else if err == nil && statusCode != tc.Expected {
			t.Errorf("Código de estado esperado %d, pero obtuvo %d", tc.Expected, statusCode)
		} else {
			t.Logf("Solicitud con error %d, como se esperaba", tc.Expected)
		}
	}
}

