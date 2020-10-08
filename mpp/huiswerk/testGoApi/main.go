package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type movie struct {
    IMDBId string       `json:"IMDBId"`
    Name string         `json:"Name"`
    Year int            `json:"Year"`
    Score float64       `json:"Score"`
    Description string  `json:"Description"`
}

type rating struct{
    Source string
    Value string
}

type omdbMovie struct {
    Title string          `json:"Title"`
    Year string           `json:"Year"`
    Rated string          `json:"Rated"`
    Released string       `json:"Released"`
    Runtime string        `json:"Runtime"`
    Genre string          `json:"Genre"`
    Director string       `json:"Director"`
    Writer string         `json:"Writer"`
    Actors string         `json:"Actors"`
    Plot string           `json:"Plot"`
    Language string       `json:"Language"`
    Country string        `json:"Country"`
    Awards string         `json:"Awards"`
    Poster string         `json:"Poster"`
    Ratings []rating      `json:"Ratings"`
    Metascore string      `json:"Metascore"`
    imdbRating string
    imdbVotes string
    imdbID string
    Type string           `json:"Type"`
    DVD string            `json:"DVD"`
    BoxOffice string      `json:"BoxOffice"`
    Production string     `json:"Production"`
    Website string        `json:"Website"`
    Response string       `json:"Response"`
}


func movieHandler(w http.ResponseWriter, req *http.Request) {
    switch req.Method {
        case "POST":
            postMovie("movie", req.Body, w)
        case "GET":
            getMovies("movie", w)
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
            w.Write([]byte("method not allowed"))
    }
}

func movieDescriptionHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method != "GET" { return }

    ApiKey := "e1843a60"
    requestURL := "http://omdbapi.com/?apikey=" + ApiKey + "&plot=full&i="

    reqDescId := strings.Split(req.RequestURI, "/descriptions/")[1];
    if reqDescId != "" {
        resp, err := http.Get(requestURL + reqDescId)
        if err != nil { http.Error(w, err.Error(), 500) }
        b, err := ioutil.ReadAll(resp.Body)
        defer resp.Body.Close()
        w.Write(b)
    } else {
        db := getDatabase()
        query := "SELECT id, imdbid FROM movies"
        prep, err := db.Prepare(query)
        if err != nil {log.Fatalln("openin")}
        results,_ := prep.Query()
        db.Close()
        for results.Next() {
            var (
                id int
                imdbid string
            )
            results.Scan(&id, &imdbid)
            go importMovieDesc(requestURL, imdbid)
        }
        w.Write([]byte("Descriptions has been added"))
    }
}

func getMovies(request string, w http.ResponseWriter) {
    dbMovies := getMoviesFromDB()

    if dbMovies == nil {
        w.Write([]byte("There are no movies yet"))
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(dbMovies)
}

func getMoviesFromDB() ([]movie){
    db := getDatabase()
    query := "SELECT * FROM movies"
    prep, err := db.Prepare(query)
    if err != nil {
        log.Fatalln("opening db gives: ", err.Error())
    }

    var dbMovies []movie
    results,_ := prep.Query()
    for results.Next() {
        var (
            id int
            imdbid string
            name string
            year int
            score float64
            description string
        )
        results.Scan(&id, &imdbid, &name, &year, &score, &description)
        dbMovie := movie{
            imdbid,
            name,
            year,
            score,
            description,
        }
        dbMovies = append(dbMovies, dbMovie)
    }
    return dbMovies
}

func postMovie(request string, body io.ReadCloser, w http.ResponseWriter) {
    // reading the body and closing it after its done
    b, err := ioutil.ReadAll(body)
    defer body.Close()
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    // make new movie and make it a movie(type)
    var newMovie movie
    err = json.Unmarshal(b, &newMovie)
    if err != nil {
        w.Write([]byte("There maybe something wrong with the types the provided json.\n"))
        return
    }

    err = addMovie(newMovie)
    if err != nil {
        log.Println(err)
    }

    json.NewEncoder(w).Encode(newMovie)
}

func movieGetHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method != "GET" {
        return
    }

    var imdbId string
    imdbId = strings.Split(req.RequestURI, "movies/")[1]

    db := getDatabase()
    defer db.Close()

    query := "SELECT * FROM movies WHERE imdbid is ?;"

    prep, _ := db.Prepare(query)
    result := prep.QueryRow(imdbId)
    var (
        id int
        imdbid string
        movieName string
        year int
        score float64
        description string
    )
    result.Scan(&id, &imdbid, &movieName, &year, &score, &description)
    if imdbid == "" {
        w.Write([]byte("No item has been found with the id: "+ imdbId))
        return
    }
    dbMovie := movie{
        imdbid, movieName, year, score, description,
    }

    json.NewEncoder(w).Encode(dbMovie)
}

func importMovieDesc(req string , imdbid string) {
    log.Println("getting results for imdbid: ", imdbid)

    // get request for description
    resp, err := http.Get(req + imdbid)
    if err != nil {log.Println("error in req to omdb")}

    // read body
    b,_ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    // translate to omdbMovie struct
    var res omdbMovie
    json.Unmarshal([]byte(b), &res)

    // get the database and exec the query
    db := getDatabase()
    query := "UPDATE movies SET desc = ? WHERE imdbid is ?;"

    prep,_ := db.Prepare(query)
    prep.Exec(res.Plot, imdbid)
    db.Close()
    log.Println("Got desc for: " + imdbid + ", which is: " + res.Plot)
}

func importData() {
    db := getDatabase()
    defer db.Close()

    csvFile, err := os.Open("watchlist.csv")
    if err != nil {
        log.Println(err)
        return
    }

    r := csv.NewReader(csvFile)
    if _, err := r.Read(); err != nil {
        log.Println("ERRRROR terror")
    }

    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }

        imdbId := record[1]
        name := record[5]
        year,_ := strconv.ParseFloat(record[10], 64)
        score,_ := strconv.Atoi(record[8])
        description := ""

        newMovie := movie{
            imdbId,
            name,
            score,
            year,
            description,
        }
        insertMovieInToDatabase(newMovie)
    }

}

func main() {
    createTable()
    go importData()

    http.HandleFunc("/movies", movieHandler)
    http.HandleFunc("/movies/", movieGetHandler)
    http.HandleFunc("/descriptions", movieDescriptionHandler)
    http.HandleFunc("/descriptions/", movieDescriptionHandler)
    http.ListenAndServe(":8080", nil)
}

func addMovie(newMovie movie) (error) {
    var faultMessage string
    switch 0 {
    case len(newMovie.IMDBId):
        faultMessage = faultMessage + "Missing IMDBid; "
        fallthrough
    case len(newMovie.Name):
        faultMessage = faultMessage + "Missing name; "
        fallthrough
    case newMovie.Year:
        faultMessage = faultMessage + "Missing Year; "
        fallthrough
    case int(newMovie.Score):
        faultMessage = faultMessage + "Missing Score; "
    default:
        faultMessage = "0"
    }

    if faultMessage == "0" {
        insertMovieInToDatabase(newMovie)
        return nil
    }

    return errors.New(faultMessage)
}

func insertMovieInToDatabase(newMovie movie) {
    db := getDatabase()
    defer db.Close()

    query := "INSERT INTO MOVIES (imdbid, name, year, score) VALUES (?, ?, ?, ?)"
    prep, err := db.Prepare(query)
    if err != nil {
        log.Println(err.Error())
        return
    }

    prep.Exec(newMovie.IMDBId, newMovie.Name, newMovie.Year, newMovie.Score)
}

func getDatabase() *sql.DB {
    db, err := sql.Open("sqlite3", "watchlist.db")

    if err != nil {
        log.Fatalln("Couldn't open database", err)
    }
    return db
}

func createTable() {
    db := getDatabase()
    defer db.Close()
    query := `
        Create TABLE IF NOT EXISTS movies
            (
                id INTEGER PRIMARY KEY,
                imdbid TEXT,
                name TEXT,
                year int,
                score float,
                desc TEXT
            )`

    prep, err := db.Prepare(query)
    if err != nil {
        log.Fatalln("db error:", err.Error())
    }
    prep.Exec()
}
