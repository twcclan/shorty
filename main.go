package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var urls = make(map[string]string)

func redirect(w http.ResponseWriter, r *http.Request) {
	target := r.URL.RequestURI()[1:]

	if url, ok := urls[target]; ok {
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Short url was not found, please check for correct spelling.")
	}
}

func main() {
	file, err := os.Open("shorty.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		log.Println(strings.Join(record, " => "))
		urls[record[0]] = record[1]
	}

	http.HandleFunc("/", redirect)

	log.Println(http.ListenAndServe(":6060", nil))
}
