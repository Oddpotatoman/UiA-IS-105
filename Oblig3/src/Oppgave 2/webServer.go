package main


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
	http.HandleFunc("/3", api3)
	http.HandleFunc("/4", api4)
	http.HandleFunc("/5", api5)
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
	//Very varied when this works or not
	fp := path.Join("templates","api1.html")

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
	fp := path.Join("templates","api2.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, people1); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
//Severe weather warning systems
func api3(w http.ResponseWriter, r * http.Request){

	type Sites struct{
		County string `json:"county"`
		Notifications string `json:"notifications"`
		User_fees string `json:"user_fees"`
		Web_site string `json:"web_site"`
	}

	url := "https://data.mo.gov/resource/b69t-kpq2.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var warningSites []Sites
	jsonErr := json.Unmarshal(body, &warningSites)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fp := path.Join("templates","api3.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, warningSites); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func api4(w http.ResponseWriter, r * http.Request){

	type Hospitals struct{
		Average_covered_charges string `json:"average_covered_charges"`
		Average_medicare_payments string `json:"average_medicare_payments"`
		Provider_city string `json:"provider_city"`
		Provider_name string `json:"provider_name"`
		Provider_street_address string `json:"provider_street_address"`
		Provider_zip_code string `json:"provider_zip_code"`
		Total_discharges string `json:"total_discharges"`
	}

	url := "https://data.cms.gov/resource/ehrv-m9r6.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var hospitalAverages []Hospitals
	jsonErr := json.Unmarshal(body, &hospitalAverages)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fp := path.Join("templates","api4.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, hospitalAverages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func api5(w http.ResponseWriter, r * http.Request){

	type Tick struct {
		Name string `json:"name"`
		Symbol string `json:"symbol"`
		Rank string `json:"rank"`
		Price_usd string `json:"price_usd"`
		Price_btc string `json:"price_btc"`
	}

	url := "https://api.coinmarketcap.com/v1/ticker/?start=0&limit=100"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var cryptoCurrency []Tick
	jsonErr := json.Unmarshal(body, &cryptoCurrency)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fp := path.Join("templates","api5.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, cryptoCurrency); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

