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

	//TODO: Get all file timestamps using syscall
	wts.Modified = cm.FormatTimestamp(fi.ModTime())
	wts.Accessed =  ""
	wts.Created = ""
	wts.Birth =  ""
	return
}
