package Oppgave_2

//People in Space JSON http://api.open-notify.org/astros.json


import (
	"html/template"
	"net/http"
	"fmt"
	"encoding/json"
	"time"
	"io/ioutil"
	"log"
	"path"
)

func main () {
	http.HandleFunc("/", client)
	http.HandleFunc("/1", api1)
	http.HandleFunc("/2", api2)
	http.ListenAndServe(":8080", nil)
}
func client (w http.ResponseWriter, r * http.Request) {
	fmt.Fprint(w, "<b>Hello client</b>")
}
func api1 (w http.ResponseWriter, r * http.Request){

	url := "https://hotell.difi.no/api/json/stavanger/miljostasjoner"
	apiClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	type Entry struct {
		Latitude float32 `json:"latitude,string"`
		Navn string `json:"navn"`
		Plast string `json:"plast"`
		Glass_metall string `json:"glass_metall"`
		Tekstil_sko string `json:"tekstil_sko"`
		Longitude float32 `json:"longitude,string"`
	}
	type Entries struct {
		Entries []Entry `json:"entries"`
		Page int `json:"page"`
		Pages int `json:"pages"`
		Posts int `json:"posts"`

	}


	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := apiClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

 	var entires Entries
 	jsonErr := json.Unmarshal(body, &entires)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fp := path.Join("Oblig3/Oppgave 1/templates/api1.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, entires); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func api2(w http.ResponseWriter, r * http.Request){

	type people struct{
		Name string `json:"name"`
		Craft string `json:"craft"`
	}
	type space struct {
		People []people `json:"people"`
	}

	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	people1 := space{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fp := path.Join("Oblig3/Oppgave 1/templates/api2.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, people1); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

