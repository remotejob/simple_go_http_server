/*Package entryLogsHandler
used to keep records slice of struc in pointer,
add new records and delete not used records in pointer
keep clean Database (.csv file but it can be other backend for example redis etc..)
*/
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

//AddLastRecords used mostly for control database (in our case simple .csv file),
//it keep only needed records.
//1. It used at spart up (init func) init paramaters must be set TRUE
//2. Used as goroutines periodicly cleapUp Database init paramaters must be set FALSE
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

	wr.WriteAll(tmpArrStr)

}

// AddNewHit add last click click hits in pointer
func (logs *EntryLog) AddNewHit(logstr string) {

	logrecord := domains.Log{time.Now(), logstr}

	logs.EntryLog = append(logs.EntryLog, logrecord)

}

//DeleteExtraRecords so it delete not used records from slice
//as well used in goroutines
func (logs *EntryLog) DeleteExtraRecords(index int) {

	var tmparr []domains.Log

	tmparr = append(logs.EntryLog[:index], logs.EntryLog[index+1:]...)

	logs.EntryLog = tmparr

}
