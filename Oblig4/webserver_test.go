package main

import (
	"testing"
	"log"
)

func TestGetMessage(t *testing.T){
	expected := "Det er sol ute, husk solkrem. Det er en grei varme ute, så en t-skjorte, genser og bukse vil funke fett."
	actual := getMessage(14,1.5,"Klarvær")
	if actual != expected {
		t.Error("Test failed")
	}

	}
func TestConvertStringToInt(t *testing.T){
	expected := 1.7
	actual := convertStringToFloat("1.7")
	if actual != expected {
		t.Error("Test failed")
	}
}
func TestConvertStringToFloat(t *testing.T){
	expected := 20
	actual := convertStringToInt("20")
	if actual != expected {
		t.Error("Test failed")
	}
}
func TestGetXML(t *testing.T){
	expected := `<?xml version="1.0" encoding="utf-8"?>`
	url := "http://www.yr.no/sted/Norge/Rogaland/Sandnes/Sandnes/varsel.xml"
	value := ""
	if getWdFromUrl, err := getXML(url); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		log.Println("Received XML")
		value = getWdFromUrl
	}
	runes := []rune(value)
	safeSubstring := string(runes[0:38])
	log.Println(safeSubstring)
	actual := safeSubstring
	if actual != expected {
		t.Error("Test failed")
	}

}