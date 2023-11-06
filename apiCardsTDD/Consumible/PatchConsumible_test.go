package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"testing"
	
)
func simulatePatchConsumible() (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
    //ID (necesario para PATCH)
	writer.WriteField("id", "650f390b7aaeb67f7dfc7140")
    //Dato a actualizar
	writer.WriteField("descripcion", "Actualizar descripcion")
	// Cerrar el escritor del formulario formData
	writer.Close()

    // Crear una solicitud POST con el cuerpo del formulario
	req, err := http.NewRequest("PATCH", "https://cards.thenexusbattles2.cloud/api/consumible/", body)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Content-Type", writer.FormDataContentType())

    // Hacer la solicitud POST al servidor de prueba
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

//Se prueba la petición anteriormente definida
func TestSimulatePatchConsumible(t *testing.T) {
    t.Log("Prueba de la petición...")

    resp, err := simulatePatchConsumible()
    if err != nil {
        t.Fatalf("Error al simular la solicitud POST: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        var rtx interface{}

        decoder := json.NewDecoder(resp.Body)

        decoder.Decode(&rtx)

        fmt.Println(rtx)

        t.Logf("Solicitud exitosa: Se obtuvo un código de estado 200")
    } else {
        t.Fatalf("Se esperaba un código de estado 200, pero se obtuvo: %d", resp.StatusCode)
    }
}

//Petición para simular un badRequest
func simulatePatchBadRequest() (*http.Response, error) {
	
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("id", "dato Invalido")
	writer.WriteField("descripcion", "Actualizar descripcion")
	writer.Close()
	req, err := http.NewRequest("PATCH", "https://cards.thenexusbattles2.cloud/api/consumible/", body)
    if err != nil {
        return nil, err
    }
    req.Header.Add("Content-Type", writer.FormDataContentType())
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    return resp, nil
}
//Probar peticion badRequest
func TestSimulatePatchBadRequest(t *testing.T) {
    t.Log("Prueba de una mala petición...")
    resp, err := simulatePatchBadRequest()
    if err != nil {
        t.Fatalf("Error al simular la solicitud POST: %v", err)
    }
    defer resp.Body.Close()
    if resp.StatusCode == http.StatusBadRequest {
        var rtx interface{}
        decoder := json.NewDecoder(resp.Body)
        decoder.Decode(&rtx)
        fmt.Println(rtx)
        t.Logf("Prueba exitosa: Se obtuvo un código de estado 400")
    } else {
        t.Fatalf("Se esperaba un código de estado 400, pero se obtuvo: %d", resp.StatusCode)
    }
}