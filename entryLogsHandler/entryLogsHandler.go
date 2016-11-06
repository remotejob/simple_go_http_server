package entryLogsHandler

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/remotejob/simple_go_http_server/domains"
)

type EntryLog struct {
	EntryLog []domains.Log
}

func NewEntryLog() *EntryLog {
	return &EntryLog{[]domains.Log{}}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
func (logs *EntryLog) AddLastRecords(file string, deltaTime time.Duration, init bool) {

	csvfile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	if init {
		log.Println("Start initfunc")
		reader.FieldsPerRecord = -1
		rawCSVdata, err := reader.ReadAll()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, each := range rawCSVdata {

			layout := time.RFC1123Z
			timetoinsert, _ := time.Parse(layout, each[0])

			delta := time.Since(timetoinsert)

			if deltaTime.Seconds() > delta.Seconds() {
				logrecord := domains.Log{timetoinsert, each[1]}
				logs.EntryLog = append(logs.EntryLog, logrecord)

			}

		}

	} else {
		log.Println("Periodic CleanUP")
	}
	csvfile.Close()

	os.Truncate(csvfile.Name(), 0)

	csvfile, err = os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer csvfile.Close()

	var tmpArrStr [][]string
	for _, value := range logs.EntryLog {
		layout := time.RFC1123Z
		logarr := []string{value.Date.Format(layout), value.Log}

		tmpArrStr = append(tmpArrStr, logarr)

	}

	wr := csv.NewWriter(csvfile)

	log.Println("tmpArrStr", len(tmpArrStr))

	wr.WriteAll(tmpArrStr)

	// for _, record := range tmpArrStr {

	// 	log.Println(record)

	// 	wr.Write(record)
	// 	wr.Flush()

	// }

	// defer w.Flush()

}
func (logs *EntryLog) AddNewHit(logstr string) {

	logrecord := domains.Log{time.Now(), logstr}

	logs.EntryLog = append(logs.EntryLog, logrecord)

}

func (logs *EntryLog) DeleteExtraRecords(index int) {

	var tmparr []domains.Log

	tmparr = append(logs.EntryLog[:index], logs.EntryLog[index+1:]...)

	logs.EntryLog = tmparr

}

func (logs *EntryLog) CleanExtraRecords(deltaTimeMin int) {

	// log.Println(len(logs.EntryLog))

}
