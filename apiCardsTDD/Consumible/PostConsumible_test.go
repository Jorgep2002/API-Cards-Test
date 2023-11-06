package main

import (
	"testing"
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"io"
)
//Modelo de EfectoModel (estructura usada en un parametro solicitado en la petición)
type EfectoModel struct {
    Estadistica   string `json:"estadistica" validate:"required,oneof=Poder Vida Defensa Ataque Daño"`
    ValorAfectado int    `json:"valorAfectado" validate:"required"`
    TurnosValidos int    `json:"turnosValidos" validate:"required,gte=-1"`
    Id_Estrategia int    `json:"id_Estrategia" validate:"required,gt=0"`
}
func simulatePostConsumible() (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	//Datos
	efecto := EfectoModel{
		Estadistica:   "Poder",
		ValorAfectado: 10,
		TurnosValidos: 5,
		Id_Estrategia: 1,
	}
	efectoHeroe := EfectoModel{
		Estadistica:   "Poder",
		ValorAfectado: 10,
		TurnosValidos: 5,
		Id_Estrategia: 1,
	}
	writer.WriteField("icono", "http://ejemplo.com")
	writer.WriteField("nombre", "test")
	writer.WriteField("clase", "Guerrero")
	writer.WriteField("tipo", "Armas")
	writer.WriteField("coleccion", "Armas")
	efectoJSON, _ := json.Marshal(efecto)
	writer.WriteField("efecto", string(efectoJSON))
	efectoHeroeJSON, _ := json.Marshal(efectoHeroe)
	writer.WriteField("efectoHeroe", string(efectoHeroeJSON))
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
	req, err := http.NewRequest("POST", "https://cards.thenexusbattles2.cloud/api/consumible/", body)
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
//Prueba de  solicitud anteriormente definida
func TestPostConsumible(t *testing.T) {
    t.Log("Prueba de la petición...")

    resp, err := simulatePostConsumible()
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
func PostConsumible() (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.WriteField("nombre", "test")
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
	writer.Close()
	req, err := http.NewRequest("POST", "https://cards.thenexusbattles2.cloud/api/consumible/", body)
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
//Prueba de petición con flujo alternativo
func TestSimulatePostBadRequest(t *testing.T) {
    t.Log("Prueba de una mala petición...")
    resp, err := PostConsumible()
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