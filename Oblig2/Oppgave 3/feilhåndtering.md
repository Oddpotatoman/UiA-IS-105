Feilhåndteringer:



	file, err := os.Create("3b.txt")
	if err != nil {
		log.Fatal("Kan ikke opprette fil", err)
	}
	defer file.Close()

	f, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(f, "%d\n%d", n1, n2); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

Over ser vi en feilhåndtering i filen "addtofile.go". Her ser vi at det blir lagt en fil
kalt (result.txt) over og over igjen hver gang programmet kjører. Dette gir os et blank ark hver gang. Hvis
programmet ikke kan kjøre en fil, vil vi få en error som sier "Kan ikke opprette fil".
Hvis dette skjer så åpner programmet et nytt fil og dersom det oppstår en feil så vil det komme opp et error som sier
"Kan ikke åpne fil". Derfor bruker vi os.O_WRONLY sånn at vi er sikre på at den nye filen
vi åpner kan skrives i. Videre ser vi at programmet har brukt verdier som vi har satt inn i filen.
Hvis det kommer opp error så vil det komme opp "Kan ikke skrive fil"
Til slutt lukkes filen og deretter blir skrevet i sumfromfile.go.

    func checkErr(e error) {
	if e != nil {
		panic(e)

Koden over har også en feilhåndteringsfunksjon som har i oppgave å lese sluttresusltatet fra "3b.txt". Dersom den finner
feil vil den utføre kommandoen panic(e). Panic er en innebygd funksjon i go språket som gjør at den ordinære flyten i
koden stopper opp. Deretter vil den få panikk. Dette vil si at når en funksjon kaller på panic vil den stoppe
utførelsen av denne funksjonen. Den vil fortsette å utføre defer funksjoner i koden.


    file, err := os.OpenFile("result.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Kan ikke åpne fil", err)
	}
	if _, err := fmt.Fprintf(file,"\n%d\n", result); err != nil {
		log.Fatal("Kan ikke skrive fil", err)
	}

	if err := file.Close(); err != nil {
		log.Fatal(err)
	}


Over ser vi på feilhåndteringen i filen sumfromfile.go. Vi kan se at den såpass er lik som den andre, men her blir filen åpnet med os.O_APPEND i tillegg til O_WRONLY.
Denne filen er ikke tom og vi vil beholde informasjonen i den for å printe imput som en del av resultatet.
Vi har derfor brukt APPEND for å  legge til den nye informasjonen etter den eksisterende.