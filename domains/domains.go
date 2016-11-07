/*
keep struc in one place
*/
package domains

import (
	"time"
)

//Log record simple struct
type Log struct {
	Date time.Time
	Log  string
}
