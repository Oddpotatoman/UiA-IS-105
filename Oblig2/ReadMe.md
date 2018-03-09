Oppgave 1:

I oppgave 1 har vi laget et program som skal returnere definert informasjon om en fil "fileinfo.go". I tillegg har vi
bygget en build- program som kan utføres via kommandolinja.

Oppgave 2:

I oppgave 2 har vi laget en go fil som har filnavnet "filecount.go". Denne filen skal lese en tekst fil og i tillegg
skrive ut maks antall linjer samt de fem runene som forekommer mest i fila i medfølgende fil text.txt. Vi har også bygget
filen om til et kjørbart program sånn at den han bli utført i en terminal. Det blir vist en resultat i formatform med info om
filen, antal linjer i filen, og hvilken runer som er fremkalt mest.

Oppgave 3:

Oppgave 3 er delt inn i flere deloppgaver fra a til e.

a.
I Oppgave a har vi laget en fil med navn "addup.go" som inneholder 2 funksjoner som kommuniserer med hverandre gjennom channels.
I funksjon A har vi lest inn to tall fra ternimalen og deretter sendt disse to verdiene til funksjon B. Funksjonen B har i oppgave å legge
sammen de to tallene og sende resultatet tilbake til funksjon A og skrive ut resusltatet i terminalen.

b.
I oppgave b har vi laget 2 programmer med navnene "addtofile.go" og "sumfromfile.go". Opggaven her er å få dem til å kommunisere med hverandre
gjennom en fil. "Addtofile" leser to tall og skriver disse til en fil. "Sumfromfile" derimot leser tall fra en fil som er laget av program A
og adderer de to tallene. Summen skrives tilmake til samme fil. Program A leser resultatet og skriver ut summen til stdout.

c.

I oppgave c har vi beskrivet feilhåndteringer i både a og b.

d.

I oppgave d har vi implementert en håndtering av SIGINT i a og b. Dette skriver ut et avslutningsmelding
dersom de mottar SIGINT før de fullføres naturlig.

e.

I oppgave e har vi bygget om filene i både 1 og 2 og gjort dem til kjørbare  filer på operativsystemet (.exe)
og lagt dem til i en egen bin-mappe.