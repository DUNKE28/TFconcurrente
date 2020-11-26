package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	//"strconv"
	"bufio"
	//"io"
	//"net/http"
	"net"
	"strings"
	//"time"
)

//estructura
type Data struct {
	Altura float64 `json:"altura"`
	Horas  float64 `json:"horas"`
}

var datos []Data

//var arrResult []float64
var altura string
var hora string

//var direccion_nodo string

func main() {

	datos = []Data{
		{120.3, 3000.2},
		{150.0, 5000},
		{130.5, 4000}}

	//conexion

	handleResquest()

	//direccion_nodo = "localhost:9000"
	conn, _ := net.Dial("tcp", "localhost:9000")
	defer conn.Close()

	allDate := altura + "," + hora //concatenamos los datos recibidos
	fmt.Fprintf(conn, allDate)     //envia el string al nodo01
	// for {
	// 	con, _ := ln1.Accept()
	// 	go obtenerData(con)
	// }
	obtenerData(conn)

}

func handleResquest() {
	http.HandleFunc("/results", getAll)
	http.HandleFunc("/pushresult", pushResult)
	http.HandleFunc("/getdata", getResult)

	log.Fatal(http.ListenAndServe(":9001", nil))
}

func getAll(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	//serializacion
	jsonBytes, _ := json.MarshalIndent(datos, "", " ")
	io.WriteString(res, string(jsonBytes))
}

func pushResult(res http.ResponseWriter, req *http.Request) {
	var consulta Data

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(res, "Inserte datos validos")
	}

	json.Unmarshal(reqBody, &consulta)
	datos = append(datos, consulta)

	res.Header().Set("Content-Type", "application/json")

}

func getResult(res http.ResponseWriter, req *http.Request) {
	//http://localhost:3000/getdata?alt=155&hour=8555

	altura = req.FormValue("alt") //se obtiene los valores que se ingresa
	hora = req.FormValue("hour")

	fmt.Println(altura)
	fmt.Println(hora)

	//sendResult(con, alturaData, horaData)

	res.Header().Set("Content-Type", "application/json")
}

func obtenerData(con net.Conn) {
	//defer con.Close()
	//sendResult(con)
	r := bufio.NewReader(con)
	respuesta, _ := r.ReadString('\n')
	resp := strings.Split(respuesta, ",")
	// val, err := strconv.ParseFloat(strings.ReplaceAll(resp,"\r\n",""), 64)
	// if err != nil{
	// 	fmt.Println("Error: ", err)
	// }

	//arrResult = append(arrResult, val)
	fmt.Println(resp)
}

func sendResult(conn net.Conn) {

	allDate := altura + "," + hora

	fmt.Fprintf(conn, allDate)
}
