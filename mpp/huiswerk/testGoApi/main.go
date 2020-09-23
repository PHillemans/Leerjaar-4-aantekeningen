package main

import (
	"database/sql"
	"encoding/json"
	"errors"
    _ "github.com/mattn/go-sqlite3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type movie struct {
    IMDBId int          `json:"IMDBId"`
    Name string         `json:"Name"`
    Year int            `json:"Year"`
    Score float64       `json:"Score"`
}

type allMovies []movie

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

func getMovies(request string, w http.ResponseWriter) {
    db := getDatabase()
    defer db.Close()
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
            imdbid int
            name string
            year int
            score float64
        )
        results.Scan(&id, &imdbid, &name, &year, &score)
        dbMovie := movie{
            imdbid,
            name,
            year,
            score,
        }
        dbMovies = append(dbMovies, dbMovie)
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(dbMovies)
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
        log.Fatalln(err)
    }

    err = addMovie(newMovie)
    if err != nil {
        http.Error(w, err.Error(), 400)
    }

    json.NewEncoder(w).Encode(newMovie)
}

func movieGetHandler(w http.ResponseWriter, req *http.Request) {
    if req.Method != "GET" {
        return
    }

    //@TODO get imdbId from parameters in request query
    var imdbId int

    db := getDatabase()
    defer db.Close()

    query := "SELECT * FROM movies WHERE imdbid is ?;"

    prep, _ := db.Prepare(query)
    result := prep.QueryRow(imdbId)
    var (
        id int
        imdbid int
        movieName string
        year int
        score float64
    )
    result.Scan(&id, imdbid, &movieName, &year, &score)
    dbMovie := movie{
        id, movieName, year, score,
    }

    json.NewEncoder(w).Encode(dbMovie)
}

func main() {
    createTable()
    http.HandleFunc("/movies", movieHandler)
    http.HandleFunc("/movies/", movieGetHandler)
    http.ListenAndServe(":8080", nil)
}

func addMovie(newMovie movie) (error) {
    var faultMessage string
    switch 0 {
    case newMovie.IMDBId:
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
        log.Fatal(err.Error())
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
                imdbid int,
                name TEXT,
                year int,
                score float
            )`

    prep, err := db.Prepare(query)
    if err != nil {
        log.Fatalln("db error:", err.Error())
    }
    prep.Exec()
}
