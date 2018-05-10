package main

import (
	"html/template"
	"net/http"
	"io/ioutil"
	"log"
	"path"
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
type Symbol struct {
	XMLName xml.Name `xml:"symbol"`
	Name string `xml:"name,attr"`
}
type Times struct {
	XMLName     xml.Name `xml:"time"`
	From        string   `xml:"from,attr"`
	To          string   `xml:"to,attr"`
	Period      string   `xml:"period,attr"`
	Temperature Temps    `xml:"temperature"`
	WindSpeed   Ws       `xml:"windSpeed"`
	Symbol      Symbol   `xml:"symbol"`
}
type Forecasts struct {
	XMLName xml.Name `xml:"forecast"`
	Tabular []Times  `xml:"tabular>time"`
}
type weatherdata struct {
	XMLName  xml.Name  `xml:"weatherdata"`
	Forecast Forecasts `xml:"forecast"`
}

func main() {
	//Make a quick index
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fp := path.Join("Oblig4/templates/index.html")

		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, "index"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})


	http.HandleFunc("/sandnes", func(w http.ResponseWriter, r *http.Request) {
		weatherSuggestion(w, r, "Rogaland/Sandnes/Sandnes")
	})
	http.HandleFunc("/oslo", func(w http.ResponseWriter, r *http.Request) {
		weatherSuggestion(w, r, "Oslo/Oslo/Oslo")
	})
	http.HandleFunc("/kristiansand", func(w http.ResponseWriter, r *http.Request) {
		weatherSuggestion(w, r, "Vest-Agder/Kristiansand/Kristiansand")
	})
	http.ListenAndServe(":8080", nil)
}

func weatherSuggestion(w http.ResponseWriter, r *http.Request, extension string) {

	url := "http://www.yr.no/sted/Norge/" + extension + "/varsel.xml"

	var XMLweather = ""
	if getWdFromUrl, err := getXML(url); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		log.Println("Received XML:")
		log.Println(getWdFromUrl)
		XMLweather = getWdFromUrl
	}

	fp := path.Join("Oblig4/templates/suggestion.html")
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
	var weatherType = weatherData.Forecast.Tabular[0].Symbol.Name
	var totalTemp = convertStringToInt(weatherData.Forecast.Tabular[0].Temperature.Value)
	var totalWs = convertStringToFloat(weatherData.Forecast.Tabular[0].WindSpeed.Mps)

	println(weatherType)
	println(totalWs)
	println(totalTemp)
	if err := tmpl.Execute(w, getMessage(totalTemp, totalWs, weatherType)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func getMessage(temp int, ws float64, wt string) (string) {
	//Disse værtypene var de jeg fant ifra yr. Siden vi ikke har all verdens tid så gjør vi det ikke altfor detaljert.
	var message string
	if wt == "Regn" || wt == "Kraftig regn" || wt == "Kraftige regnbyger" {
		message = "Det regner idag, bruk regnjakke."
		if temp > 15 {
			message += " Det er også ganske varmt, så ha ett tynt lag med klær under."
		} else if temp > 10 {
			message += " Det er en grei varme ute, så ikke bombarder deg med klær."
		} else if temp > 0 {
			message += " Det er også ganske kjølig ute, vi foreslår at du kler deg godt."
		} else {
			message += " Det er også SYKT kaldt ute så vi foreslår du kler deg utroligt godt #ullErGull"
		}
	} else if wt == "Skyet" || wt == "Lettskyet" || wt == "Delvis skyet" {
		message = "Det er skyet ute."
		if temp > 20 {
			message += " Det er også varmt da, så T-skjorte og bukse høres greit ut."
		} else if temp > 10 {
			message += " Det er en grei varme ute, så ikke bombarder deg med klær men ikke gå lettkledd heller."
		} else if temp > 0 {
			message += " Det er også ganske kjølig ute, vi foreslår at du kler deg godt."
		} else {
			message += " Det er også SYKT kaldt ute så vi foreslår du kler deg utroligt godt #ullErGull"
		}
	} else if wt == "Klarvær" || wt == "Sol" {
		message = "Det er sol ute, husk solkrem."
		if temp > 20 {
			message += " Det er også varmt da, så hiv på en t-skjorte og shorts."
		} else if temp > 10 {
			message += " Det er en grei varme ute, så en t-skjorte, genser og bukse vil funke fett."
		} else if temp > 0 {
			message += " Det er fortsatt ganske kjølig ute da så hiv på en jakke i tilegg til genser og bukse."
		} else {
			message += " Det er også SYKT kaldt ute så vi foreslår du kler deg utroligt godt #ullErGull"
		}
	} else {
		message = ""
		if temp > 20 {
			message += "Det er varmt ute, hiv på deg sommeklær og hopp i vannet."
		} else if temp > 10 {
			message += "Det er litt smålig kjølig ute, hiv på deg en jakke."
		} else if temp > 0 {
			message += "Her var det litt kaldt altså, litt tykkere klær en vanelig kan være greit."
		} else {
			message += " Det er SYKT kaldt ute så vi foreslår du kler deg utroligt godt #ullErGull"
		}
		if ws > 10 {
			message += " Det blåser en del så vær litt obs på det."
		}
	}
	return message
}
func convertStringToFloat(toFloat string) (float64) {
	f, err := strconv.ParseFloat(toFloat, 64)
	if err == nil {
		/** displayg the type of the b variable */
		fmt.Printf("Type: %T \n", f)

		/** displaying the string variable into the console */
		fmt.Println("Value:", f)
	}
	return f
}
func convertStringToInt(toInt string) (int) {
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
