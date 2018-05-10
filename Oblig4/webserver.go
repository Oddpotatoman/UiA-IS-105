package main


import (
	"html/template"
	"net/http"
	"time"
	"io/ioutil"
	"log"
	"path"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"os"
)
type Temps struct {
	XMLName xml.Name `xml:"temperature"`
	Value   string   `xml:"value,attr"`
	Unit    string   `xml:"unit,attr"`
}
type Ws struct {
	//XMLName xml.Name `xml:"windSpeed"`
	Mps  string `xml:"mps,attr"`
	Name string `xml:"name,attr"`
}
type Times struct {
	XMLName     xml.Name `xml:"time"`
	From        string   `xml:"from,attr"`
	To          string   `xml:"to,attr"`
	Period      string   `xml:"period,attr"`
	Temperature Temps  `xml:"temperature"`
	WindSpeed   Ws     `xml:"windSpeed"`
}
type Time struct {
	XMLName xml.Name `xml:"time"`
	From        string   `xml:"from,attr"`
	To          string   `xml:"to,attr"`
	EventType string `xml:"eventType,attr"`
}
type Location struct {
	XMLName xml.Name `xml:"location"`
	Time Time `xml:"time"`
}
type Text struct {
	XMLName xml.Name `xml:"text"`
	Location Location `xml:"location"`
}

type Forecasts struct {
	XMLName xml.Name `xml:"forecast"`
	Tabular []Times  `xml:"tabular>time"`
	Text Text `xml:"text"`
}
type weatherdata struct {
	XMLName  xml.Name  `xml:"weatherdata"`
	Forecast Forecasts `xml:"forecast"`
}

func main () {
	http.HandleFunc("/1", api1)
	http.HandleFunc("/5", weatherSuggestion)
	http.ListenAndServe(":8080", nil)
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

func weatherSuggestion(w http.ResponseWriter, r * http.Request){

	extension := "Rogaland/Sandnes/Sandnes"

	url := "http://www.yr.no/sted/Norge/"+extension+"/varsel.xml"

	var XMLweather = ""
	if cryptoCurrency, err := getXML(url); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		log.Println("Received XML:")
		log.Println(cryptoCurrency)
		XMLweather = cryptoCurrency
	}

	fp := path.Join("Oblig4/templates/api5.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var weatherData weatherdata
	xmlErr := xml.Unmarshal([]byte(XMLweather), &weatherData)
	if xmlErr != nil {
		log.Fatal(xmlErr)
	}
	var message string
	var weatherType = weatherData.Forecast.Text.Location.Time.EventType
	var totalTemp = convertStringToInt(weatherData.Forecast.Tabular[0].Temperature.Value)
	var totalWs = convertStringToFloat(weatherData.Forecast.Tabular[0].WindSpeed.Mps)

	if weatherType == "Rain" {
		message = "Det regner idag, bruk regnjakke."
		if totalTemp > 15 {
			message += " Det er også ganske varmt, så ha ett tynt lag med klær under."
		} else if totalTemp > 10 {
			message += " Det er en grei varme ute, så ikke bombarder deg med klær."
		} else if totalTemp > 0 {
			message += " Det er også ganske kjølig ute, vi foreslår at du kler deg godt."
			} else {
				message += " Det er også SYKT kaldt ute så vi foreslår du kler deg utroligt godt #ullErGull"
		}
	}

	println(weatherType)
	println(totalWs)
	println(totalTemp)
	if err := tmpl.Execute(w, XMLweather); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func convertStringToFloat(toFloat string)(float64){
	f, err := strconv.ParseFloat(toFloat, 64)
	if err == nil {
		/** displayg the type of the b variable */
		fmt.Printf("Type: %T \n", f)

		/** displaying the string variable into the console */
		fmt.Println("Value:", f)
	}
	return f
}
func convertStringToInt(toInt string)(int){
	convertedInt, err := strconv.Atoi(toInt)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return convertedInt
}
func getXML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil
}
