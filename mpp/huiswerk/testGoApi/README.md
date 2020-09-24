# How?

## Stappenplan voor het uitvoeren van de verschillen requests

Voor het starten van de server, moet er eerst een commando gedraaid worden om een library te importeren voor het gebruik van sqlite
- Doe dit door `go get github.com/mattn/go-sqlite3` te draaien in de terminal in de root van dit project

Wanneer alles is geinstalleerd kan het go bestand gerunt of gecompiled worden:
- Draai `go run main.go` of `go build main.go`

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

- Het verkrijgen van allen films die ingeladen zijn zullen worden gedaan met een get request. Deze zal naar de volgende url gestuurd moeten worden:
```
localhost:8080/movies
```

