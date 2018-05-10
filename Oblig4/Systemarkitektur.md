

Nodene består av Klient, webserver og databaseserveren. I denne systemarkitekturen er det en klient som kjører applikasjonen vår. Applikasjonen starter en webserver som lytter på port “:8080”. Webserver sender ut en API- request til   yr.no sine servere for å hente værdata. Vi har begrenset værmeldingene til Bergen, Stavanger, Oslo og Kristiansand. Når vi bruker applikasjonen, vil brukeren se resultatet av dataen som webleseren henter fra de ulike API’ene slik at det blir lesbart. Dette hierarkiet viser hvordan den fungerer:






![org](https://user-images.githubusercontent.com/35718955/39878636-1bcfab52-5479-11e8-9dd2-a0e7a936f914.png)




