package main

import (
	//"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"
	//"html/template"
	"io/ioutil"
	"net/http"
	//"os"
	"strconv"
	//"strings"
	"github.com/gorilla/mux"
	"log"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)
//struct de casos
type caso struct {
    Name string `json:"name"`
	Location string `json:"location"`
	Age int `json:"age"`
	Infectedtype string `json:"infectedtype"`
	State string `json:"state"`
}

const (
	address = "pythongrpc:50051"
)
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
	var jsonStr = string(`{"name":"` + c.Name + `","location":"` + c.Location +`","age":`+strconv.Itoa(c.Age)+`,"infectedtype":"`+c.Infectedtype+`","state":"`+c.State+ `"}`)
	//grpc
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
        log.Fatal("no se conecto: ",err)
	}
	defer conn.Close()
	ca := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	//enviando 
	ra, error := ca.SayHello(ctx, &pb.HelloRequest{Name: jsonStr})
	if error != nil{
		log.Fatal("error al enviar el mensaje: ",error)
	}
	log.Printf("respuesta: %s", ra.GetMessage())

}


func main() {
	//rutas del servidor

	router := mux.NewRouter()
    router.HandleFunc("/NewCaso", CrearCaso).Methods("POST")
	router.HandleFunc("/", Inicio).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))

	
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Servidor GOLAND ENCENDIDO para GRPC!!!")
}