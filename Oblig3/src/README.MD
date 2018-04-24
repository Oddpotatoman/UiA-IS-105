Obligatorisk innlevering 3

Madina Abdul Sattar,Anders Hodne,Anders Ekse og Odd Andre Vatne.

Oppgave 1
I oppgave 1 have vi satt opp en webserver som lytter på port 8080.
Når webserveren kjører, kan man søke "localhost:8080" i nettleseren, og da får man opp "Hello, client".

Oppgave 2
I oppgave 2 så henter vi inn json data fra 5 forskjellige kilder. De to første ble brukt for å vende seg til måten man strukturerer og henter inn dataen. De 3 siste er mer eksperimentelle og er blandt annet hentet ifra Amerika sine offentlige databaser. Om mer tid hadde blitt brukt på oppgaven så kunne man ha omfaktorisert litt på funksjonene for å gjøre type routing og utsendelse/pakking av data litt enklere i tilleg til å ungå en del code duplication.

Oppgave 3
I oppgave 3 har vi implementert et serverprogram i henhold til RFC 865, serveren svarer på både UDP og TCP.
Om man kjører "go run client.go" så kommer det opp en melding "Do you wanna connect to the server through
TCP or UDP?". Man kan da enten skrive inn TCP eller UDP og da får man opp en "quote of the day". 
