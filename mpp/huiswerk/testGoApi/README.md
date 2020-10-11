# How?


## Stappenplan voor het uitvoeren van de verschillen requests

Voor het starten van de server, moet er eerst een commando gedraaid worden om een library te importeren voor het gebruik van sqlite
- Doe dit door `go get github.com/mattn/go-sqlite3` te draaien in de terminal in de root van dit project

Wanneer alles is geinstalleerd kan het go bestand gerunt of gecompileerd worden:
- Draai `go run main.go` of `go build main.go`
- Als de build gedraaid is kan deze worden uitgevoerd met `./main` vanaf de locatie waar het build commando is gedraaid



### On start

Tijdens de start van de server zal de server een bestand proberen te importeren genaamd `watchlist.csv`. De server verwacht dit bestand te vinden in dezelfde locatie waar de server is opgestart



### Post

- Het posten van een film kan gedaan worden door het sturen van een postrequest met in body in de vorm van:
    ```JSON
    {
        "IMDBid": "[ID| string]",
        "Name": "[name| string]",
        "Year": [Year| int],
        "Score": [score| float]
    }
    ```

- Deze body moet verstuurd worden als "Content-Type: application/json"




### Get

Voordat er een get request wordt uitgevoerd moet er iets in de database staan. Mocht dit niet het geval zijn wordt er niets returned.

- Het verkrijgen van allen films die ingeladen zijn zullen worden gedaan met een get request. Deze zal naar de volgende url gestuurd moeten worden:
```
localhost:8080/movies
```

- De url die gebruikt moet worden om een enkele entry op te halen is:
```
localhost:8080/movies/{id}
```


### Add descriptions to movies

Voordat dit beschreven wordt is het handig om te weten dat het bijgeleverde csv bestand 260 films bevat.
Deze worden automatisch ingeladen mits het bestand te vinden is in de DIR waar het main.go bestand gedraaid wordt.

**Er** zijn twee manieren van het toevoegen van descriptions aan de films, 
1. Dit kan je doen door een get request te doen naar localhost:8080/descriptions/{ID}
2. Of je haalt ze allemaal op door een GET request te doen naar localhost:8080/descriptions.
Bij deze tweede optie zorg je ervoor dat alle films concurrent worden toegevoegd. Dit zie je voorbij komen in je terminal
