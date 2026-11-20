package main

import (
	"fmt"
	"time"
	"net/http"
	
	
)


type Character struct {
Id  int `json:"id"`
Name  string `json:"name"`
Origin Origin `json:"origin"`
Location Location `json:"location"`
Image  string `json:"image"`
}

type Origin struct {
Name  string `json:"name"`
}

type Location struct {
Name  string `json:"name"`
}
func main() {
client := http.Client{
	Timeout : time.Second * 10,
 }
  id := 1
   endpoint := fmt.Sprint
}
func main() {

}
