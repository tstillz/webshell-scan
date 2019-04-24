package timestamps

import (
	cm "../common"
	"os"
)
func StatTimes(filePath string) (wts cm.FileTimes, err error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return
	}
	wts.Modified = cm.FormatTimestamp(fi.ModTime())
	wts.Accessed =  ""
	wts.Created = ""
	wts.Birth =  ""
	return
}
