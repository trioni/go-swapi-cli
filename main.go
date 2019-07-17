package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
)

var apiURL = "https://swapi.co/api/"

// Person A Starwars character
type Person struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Weight    string `json:"mass"`
	HairColor string `json:"hair_color"`
}

func (p Person) String() string {
	return fmt.Sprintf("Name<%s>, Gender<%s>, Weight<%s>, Hair<%s>", p.Name, p.Gender, p.Weight, p.HairColor)
}

// Planet A Starwars planet
type Planet struct {
	Name    string `json:"name"`
	Climate string `json:"climate"`
}

// Starship A Starwars vessel
type Starship struct {
	Name  string `json:"name"`
	Model string `json:"model"`
	Type  string `json:"starship_class"`
}

func parseAsType(entityType reflect.Type, body []byte) interface{} {
	var entity interface{}
	err := json.Unmarshal(body, &entity)
	if err != nil {
		log.Fatalln("JSON Error: ", err)
	}
	return entity
}

func parseAsPerson(body []byte) Person {
	var entity Person
	err := json.Unmarshal(body, &entity)
	if err != nil {
		log.Fatalln("JSON Error: ", err)
	}
	return entity
}

func parseAsPlanet(body []byte) Planet {
	var entity Planet
	err := json.Unmarshal(body, &entity)
	if err != nil {
		log.Fatalln("JSON Error: ", err)
	}
	return entity
}

func parseAsStarship(body []byte) Starship {
	var entity Starship
	err := json.Unmarshal(body, &entity)
	if err != nil {
		log.Fatalln("JSON Error: ", err)
	}
	return entity
}

func parseBody(entityType string, body []byte) interface{} {
	switch entityType {
	case "people":
		return parseAsType(reflect.TypeOf(Person{}), body)
		// return parseAsPerson(body)
	case "planets":
		return parseAsPlanet(body)
	case "starships":
		return parseAsStarship(body)
	default:
		return Person{"N/A", "N/A", "N/A", "N/A"}
	}
}

func callAPI() {
	entityType := flag.String("entity", "people", "people, planets or startships")
	entityID := flag.String("id", "1", "Id of the entity")
	flag.Parse()

	callURL := apiURL + *entityType + "/" + *entityID + "/"
	fmt.Printf("Calling: %s\n", callURL)
	resp, err := http.Get(callURL)
	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode > 299 {
			log.Fatalf("Unsuccessful Request: %d\n", resp.StatusCode)
			os.Exit(1)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		entity := parseBody(*entityType, body)
		fmt.Printf("%v\n", entity)
	}
}

func main() {
	callAPI()
}
