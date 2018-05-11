

 
I denne systemarkitekturen er det en klient som kjører applikasjonen vår. Applikasjonen starter en webserver som lytter på port “:8080”. Webserver sender ut en API- request til   yr.no sine servere for å hente værdata. Vi har begrenset værmeldingene til Sandnes, Oslo og Kristiansand. Når vi bruker applikasjonen vil brukeren se resultatet av dataen som webleseren henter fra de ulike API’ene slik at det blir lesbart. Dette hierarkiet viser hvordan den fungerer:






![sandnes](https://user-images.githubusercontent.com/35718955/39917798-ec226f0c-550e-11e8-8484-9739b5a86443.png)





Nodene består av Klient, webserver og databaseserveren. Klienten oppretter en tilkobling med server hos yr.no. Deretter blir det sendt en forespørsel fra webserveren og vi får svar fra yr.no. Forespørselen om API  blir prossesert hos yr.no, og vi får svar tilbake i form av våre fårspurte data (XML fil). 
