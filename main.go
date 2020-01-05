package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	//open the file
	csvfile, err := os.Open("top50.csv")

	if err != nil {
		log.Fatalln("can't open", err)
	}

	//parse the file

	r := csv.NewReader(csvfile)

	//iterate through the records

	for {

		//read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		info := map[string]string{

			"id":           record[0],
			"name":         record[1],
			"artist":       record[2],
			"genre":        record[3],
			"tempo":        record[4],
			"energy":       record[5],
			"danceability": record[6],
			"loudness":     record[7],
			"liveness":     record[8],
			"valence":      record[9],
			"length":       record[10],
			"acousticness": record[11],
			"Speechiness":  record[12],
			"popularity":   record[13],
		}

		fmt.Printf(
			"%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s,\n",
			info["id"],
			info["name"],
			info["artist"],
			info["genre"],
			info["tempo"],
			info["energy"],
			info["danceability"],
			info["loudness"],
			info["liveness"],
			info["valence"],
			info["length"],
			info["acousticness"],
			info["Speechiness"],
			info["popularity"],
		)

	}

}
