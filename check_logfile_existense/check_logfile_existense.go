/*
simple check if logfile exist
*/
package check_logfile_existense

import "os"

//Check simple check if logfile exist
func Check(logfile string) bool {

	if _, err := os.Stat(logfile); os.IsNotExist(err) {

		return false

	}

	return true

}
