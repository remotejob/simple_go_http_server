package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"time"

	"github.com/remotejob/simple_go_http_server/check_logfile_existense"
	"github.com/remotejob/simple_go_http_server/domains"
	"github.com/remotejob/simple_go_http_server/entryLogsHandler"
	"github.com/remotejob/simple_go_http_server/recordHit"
)

const logfile = "logs.csv"
const timeSeconds int = 60

var counterToClenUpLogfile int
var newLogsEntry *entryLogsHandler.EntryLog

var deltaTime time.Duration

func init() {

	deltaTime = time.Duration(time.Second * time.Duration(timeSeconds))

	newLogsEntry = entryLogsHandler.NewEntryLog()

	exist := check_logfile_existense.Check(logfile)

	if !exist {

		log.Println("log file not exist")
		file, err := os.Create(logfile)
		if err != nil {

			log.Fatalln(err.Error())
		}
		file.Close()

	} else {

		newLogsEntry.AddLastRecords(logfile, deltaTime, true)
	}

}

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		now := time.Now()
		log := domains.Log{now, req.UserAgent()}
		go recordHit.Record(logfile, log)

		h.ServeHTTP(rw, req)
	})
}

func loghandler(rw http.ResponseWriter, req *http.Request) {

	newLogsEntry.AddNewHit(req.UserAgent())

	log.Println("Click ", len(newLogsEntry.EntryLog))

	var htmlLogs [][]string

	for i, logentry := range newLogsEntry.EntryLog {

		delta := time.Since(logentry.Date)

		if deltaTime.Seconds() > delta.Seconds() {
			htmlLogs = append(htmlLogs, []string{logentry.Date.Format(time.RFC3339), logentry.Log})

		} else {

			newLogsEntry.DeleteExtraRecords(i)

			counterToClenUpLogfile++

			if counterToClenUpLogfile > 10 {
				go newLogsEntry.AddLastRecords(logfile, deltaTime, false)
				counterToClenUpLogfile = 0
			}

		}

	}

	//Show All information
	fmt.Fprintln(rw, "Count", len(htmlLogs), "hits last 60 seconds")
	fmt.Fprintln(rw, "-------------------------------------------------")
	for i, htmlLog := range htmlLogs {
		fmt.Fprintln(rw, i, htmlLog)

	}

}

func main() {

	h := middleware(http.HandlerFunc(loghandler))
	log.Fatal(http.ListenAndServe(":8080", h))

}
