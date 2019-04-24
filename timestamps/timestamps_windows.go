package timestamps

func StatTimes(filePath string) (wts cm.FileTimes, err error) {
	wts.Modified = ""
	wts.Accessed =  ""
	wts.Created = ""
	wts.Birth = ""
	return
}
