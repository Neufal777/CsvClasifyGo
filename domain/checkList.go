package domain

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Song struct {
	Id           string
	Name         string
	Artist       string
	Genere       string
	Tempo        string
	Energy       string
	Danceability string
	Loudness     string
	Liveness     string
	Valance      string
	Length       string
	Acousticness string
	Speechness   string
	Popularity   string
}

func (song *Song) readList(fileName []string) {

	//read all files passed
	for _, file := range fileName {

		//open the file
		csvfile, err := os.Open(file)

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

			song := &Song{

				Id:           record[0],
				Name:         record[1],
				Artist:       record[2],
				Genere:       record[3],
				Tempo:        record[4],
				Energy:       record[5],
				Danceability: record[6],
				Loudness:     record[7],
				Liveness:     record[8],
				Valance:      record[9],
				Length:       record[10],
				Acousticness: record[11],
				Speechness:   record[12],
				Popularity:   record[13],
			}

			//Print struct with the information of each song in the list
			fmt.Println(song)
		}
	}

}
