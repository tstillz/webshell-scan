package timestamps

import (
	cm "../common"
)

func StatTimes(filePath string) (wts cm.FileTimes, err error) {
	wts.Modified = ""
	wts.Accessed =  ""
	wts.Created = ""
	wts.Birth = ""
	return
}
