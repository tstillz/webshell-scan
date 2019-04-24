package timestamps

func StatTimes(filePath string) (wts cm.FileTimes, err error) {
	fi, err := os.Stat(filePath)
	if err != nil {
		return
	}
	wts.Modified = ""
	wts.Accessed =  ""
	wts.Created = ""
	wts.Birth = ""
	return
}
