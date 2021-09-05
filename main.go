package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("csv/bairro.csv")
	if err != nil {
		fmt.Println(err)
	}
	readerFile1 := csv.NewReader(file)
	records, _ := readerFile1.ReadAll()
	WriteCsv(records)

	file2, err := os.Open("csv/cidade.csv")

	readerFile2 := csv.NewReader(file2)
	records2, _ := readerFile2.ReadAll()

	AppendFile(records2)

}

func WriteCsv(records [][]string) {
	f, e := os.Create("csv/output.csv")
	if e != nil {
		fmt.Println(e)
	}

	var formattedStrings = FormatFileNeighborhood(records)

	var data = [][]string{
		{"Id_bairro", "Bairro", "id_cidade", "cidade", "uf", "cod_ibge", "area"},
	}

	data = append(data, formattedStrings...)

	writer := csv.NewWriter(f)

	e = writer.WriteAll(data)
	if e != nil {
		fmt.Println(e)
	}
}

func AppendFile(records [][]string) {
	file, err := os.OpenFile("csv/output.csv", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var formattedStrings = FormatFileCity(records)

	w := csv.NewWriter(file)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteAll(formattedStrings)

}

func FormatFileCity(records [][]string) [][]string {
	var formattedStrings [][]string
	for i, data := range records {
		if i != 0 {
			explode := strings.FieldsFunc(data[0], func(r rune) bool {
				return r == '|'
			})

			formattedStrings = append(formattedStrings, []string{"", "", explode[0], explode[1], explode[2], explode[3], explode[4]})

		}
	}
	return formattedStrings
}

func FormatFileNeighborhood(records [][]string) [][]string {
	var formattedStrings [][]string
	for i, data := range records {
		if i != 0 {
			explode := strings.FieldsFunc(data[0], func(r rune) bool {
				return r == '|'
			})

			formattedStrings = append(formattedStrings, []string{explode[0], explode[1], explode[2], "", "", "", ""})
		}
	}
	return formattedStrings
}
