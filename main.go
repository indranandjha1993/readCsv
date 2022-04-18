package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type ChangeHistory struct {
	GEOGCD    string `json:"geogcd"`
	GEOGNM    string `json:"geognm"`
	GEOGNMW   string `json:"geognmw"`
	SI_ID     string `json:"si_id"`
	SI_TITLE  string `json:"si_title"`
	OPER_DATE string `json:"oper_date"`
	TERM_DATE string `json:"term_date"`
	PARENTCD  string `json:"parentcd"`
	ENTITYCD  string `json:"entitycd"`
	OWNER     string `json:"owner"`
	STATUS    string `json:"status"`
	AREAEHECT string `json:"areaehect"`
	AREACHECT string `json:"areachect"`
	AREAIHECT string `json:"areaihect"`
	AREALHECT string `json:"arealhect"`
}

func main() {
	changeHistory, err := LoadCSV("/home/indra/Downloads/Code_History_Database_(December_2021)_UK/ChangeHistory.csv")
	if err != nil {
		log.Fatal(err)
	}
	changeHistoryJson, _ := json.Marshal(changeHistory)

	// Find by cell value with gjson
	//res := gjson.GetMany(string(changeHistoryJson), `#(si_id=="1237/2016").geogcd`)

	fmt.Println(len(changeHistoryJson))
}

func LoadCSV(filePath string) ([]ChangeHistory, error) {
	csvFile, _ := os.Open(filePath)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var changeHistory []ChangeHistory
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		changeHistory = append(changeHistory, ChangeHistory{
			GEOGCD:    line[0],
			GEOGNM:    line[1],
			GEOGNMW:   line[2],
			SI_ID:     line[3],
			SI_TITLE:  line[4],
			OPER_DATE: line[5],
			TERM_DATE: line[6],
			PARENTCD:  line[7],
			ENTITYCD:  line[8],
			OWNER:     line[9],
			STATUS:    line[10],
			AREAEHECT: line[11],
			AREACHECT: line[12],
			AREAIHECT: line[13],
			AREALHECT: line[14],
		})
	}
	return changeHistory, nil
}
