package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"testing"
	"os"
	"io"
)
func simulatePostHeroe() (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// Datos formData
	writer.WriteField("icono", "http://ejemplo.com")
	writer.WriteField("nombre", "test")
	writer.WriteField("clase", "Guerrero")
	writer.WriteField("tipo", "Armas")
	writer.WriteField("poder", "6")
	writer.WriteField("vida", "6")
	writer.WriteField("defensa", "6")
	writer.WriteField("ataqueBase", "6")
	writer.WriteField("ataqueRnd", "6")
	writer.WriteField("daño", "6")
	writer.WriteField("descripcion", "Alguna descripcion")
	writer.WriteField("precio", "1000")
	writer.WriteField("descuento", "20")
	writer.WriteField("stock", "1")
	writer.WriteField("estado", "true")
	file, err := os.Open("image.webp")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	part, err := writer.CreateFormFile("imagen", "image.webp")
	if err != nil{
		fmt.Println((err))
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
	}
	// Cerrar el escritor del formulario formData
	writer.Close()
    // Crear una solicitud POST con el cuerpo del formulario
req, err := http.NewRequest("POST", "https://cards.thenexusbattles2.cloud/api/heroes/", body)
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
//Prueba petición anteriormente definida
func TestPostHeroe(t *testing.T) {
    t.Log("Prueba de la petición...")
    resp, err := simulatePostHeroe()
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
//Petición de flujo alternativo
func PostHeroe() (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// Datos formData
	writer.WriteField("nombre", "test")
	// Cerrar el escritor del formulario formData
	writer.Close()
    // Crear una solicitud POST con el cuerpo del formulario
	req, err := http.NewRequest("POST", "https://cards.thenexusbattles2.cloud/api/heroes/", body)
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
//Prueba petición de flujo alternativo
func TestSimulatePostBadRequest(t *testing.T) {
    t.Log("Prueba de una mala petición...")
    resp, err := PostHeroe()
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