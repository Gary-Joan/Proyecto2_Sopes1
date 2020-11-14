package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	//"html/template"
	"io/ioutil"
	"net/http"
	//"os"
	"strconv"
	//"strings"
	"log"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)
//struct de casos
type caso struct {
    Name string `json:"name"`
	Location string `json:"location"`
	Age int `json:"age"`
	Infectedtype string `json:"infectedtype"`
	State string `json:"state"`
}

type test_struct struct {
    Test string
}
func CrearCaso(w http.ResponseWriter, r *http.Request) {
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    var c caso
    err = json.Unmarshal(body, &c)
    if err != nil {
        panic(err)
	}
	// Do something with the Person struct...
	var jsonStr = []byte(`{"name":"` + c.Name + `","location":"` + c.Location +`","age":`+strconv.Itoa(c.Age)+`,"infectedtype":"`+c.Infectedtype+`","state":"`+c.State+ `"}`)
	





	//Enviando a Rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@mu-rabbit-rabbitmq.project.svc.cluster.local:5672/")
	failOnError(err, "Fallo al conectar con RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Fallo para abrir un canal a rabbitmq.")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Case", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Error fallo crear encolador.")

	//body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "aplication/json",
			Body:        jsonStr,
		})
	failOnError(err, "Fallo enviar mensage!!")
	


}
//funcion para ver si hay error en la conexion
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}


func main() {
	//rutas del servidor

	router := mux.NewRouter()
    router.HandleFunc("/NewCaso", CrearCaso).Methods("POST")
	router.HandleFunc("/", DoHealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8082", router))

	
}

func DoHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor GOLAND con Rabbitmq!!!")
}