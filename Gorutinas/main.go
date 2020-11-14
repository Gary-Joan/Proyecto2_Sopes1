package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Users struct {
	Users []User `json:"Casos"`
}


type User struct {
	Name   string `json:"name"`
	Location   string `json:"location"`
	Age    int    `json:"age"`
	InfectedType    string    `json:"infectedtype"`
	State    string    `json:"state"`
}

func envio(users Users, cantidadFinal int, cantidadInicial int, direccion string){
	for i := cantidadInicial; i < cantidadFinal; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
		url := direccion
		fmt.Println("URL:>", url)

		var jsonStr = []byte(`{"name":"`+users.Users[i].Name+`",
			"location": "`+users.Users[i].Location+`",
			"age": `+strconv.Itoa(users.Users[i].Age)+`,
			"infectedtype": "`+users.Users[i].InfectedType+`",
			"state": "`+users.Users[i].InfectedType+`"}`)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))

		}
}
func goRutinas(cantidad int, solicitudes int, users Users, direccion string){
	var valor1 int
	var valor2 int
	var aux int
	var anterior int
	valor1 = solicitudes/cantidad
	valor2 = solicitudes%cantidad
	aux = 0
	anterior = 0

	for true{

		if aux >= solicitudes - (valor1+valor2) {
			go envio(users,solicitudes, aux,direccion)
			break
		}

		anterior = aux
		aux = aux + valor1
		go envio(users,aux,anterior,direccion)
	}

}

func main ()  {
	reader := bufio.NewReader(os.Stdin)
	var url string
	var gorutinas int
	var solicitudes int
	var ruta string
	var users Users

	for true {
		fmt.Println("1. URL del balanceador de carga")
		fmt.Println("2. Cantidad de gorutinas")
		fmt.Println("3. Cantidad de solicitudes")
		fmt.Println("4. Ruta del Archivo")
		fmt.Println("5. Cerrar programa")
		fmt.Println("6. Mostrar datos")
		fmt.Println("7. Enviar datos")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\n")
		if text == "5" {
			break
		} else if text == "1" {
			fmt.Println("Ingrese la URL:")
			text, _ := reader.ReadString('\n')
			url = strings.TrimRight(text, "\n")

		}else if text == "2" {
			fmt.Println("Ingrese la cantidad de gorutinas:")
			text, _ := reader.ReadString('\n')
			text = strings.TrimRight(text, "\n")
			cantidad, _ := strconv.Atoi(text)
			gorutinas = cantidad

		}else if text == "3" {
			fmt.Println("Ingrese la cantidad de solicitudes:")
			text, _ := reader.ReadString('\n')
			text = strings.TrimRight(text, "\n")
			cantidad, _ := strconv.Atoi(text)

			solicitudes = cantidad

			if(solicitudes>len(users.Users)){
				fmt.Println("Cantidad de solicitudes mayor a la cantidad de casos que se desea enviar")
				solicitudes = len(users.Users)
			}

		}else if text == "4" {
			fmt.Println("Ingrese la ruta del archivo:")
			text, _ := reader.ReadString('\n')
			ruta = strings.TrimRight(text, "\n")
			jsonFile, err := os.Open(ruta)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Datos cargados con exito")
			defer jsonFile.Close()
			byteValue, _ := ioutil.ReadAll(jsonFile)


			json.Unmarshal(byteValue, &users)

		}else if text == "6" {
			fmt.Println("Url: "+url)
			fmt.Print("Gorutinas: ")
			fmt.Println(gorutinas)
			fmt.Print("Solicitudes: ")
			fmt.Println(solicitudes)
			fmt.Println("Datos cargados: ")
			envio(users, solicitudes,0,url)
		}else if text == "7" {
			goRutinas(gorutinas,solicitudes,users,url)
		}
	}

}