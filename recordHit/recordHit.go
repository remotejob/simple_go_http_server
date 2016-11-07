/*Package recordHit
used for goroutine, write in background
last hit record
*/
package recordHit

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/remotejob/simple_go_http_server/domains"
)

//Record  Write last Log record in CVS file
//it used in goroutine
func Record(file string, log domains.Log) {

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	w := csv.NewWriter(f)
	layout := time.RFC1123Z
	logarr := []string{log.Date.Format(layout), log.Log}

	w.Write(logarr)
	w.Flush()

}
