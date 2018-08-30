package main

import (
	"flag"
	"os"
	"path/filepath"
	"fmt"
	"log"
	"bufio"
	"regexp"
	"strings"
	"encoding/json"
	"time"
)

type JsonOut struct {
	FilePath string `json:"filePath"`
	Matches map[string]int `json:"matches"`
}
type Metrics struct {
	Scanned int `json:"scanned"`
	Matched int `json:"matchs"`
	Clear int `json:"nomatchs"`
	ScannedDir string `json:"scannedDir"`
	ScanTime float64 `json:"scanTime"`
}

func main()  {
	start := time.Now()
	var dir = flag.String("dir", "", "Directory to scan for webshells")
	var customRegex = flag.String("regex", "", "Override default regex with your own")
	var output = flag.String("output", "", "json or csv (pipe separated). Default is csv")
	flag.Parse()

	if *dir == ""{
		flag.PrintDefaults()
		os.Exit(1)
	}

	regexString := ""

	if *customRegex == ""{
		regexString = `Load\(Request\.BinaryRead\(int\.Parse\(Request\.Cookies|Html \= Replace\(Html\, \"\%26raquo\;\"\, \"?\"\)|pinkok|internal class reDuh|c0derz shell|md5 cracker|umer rock|Function CP\(S\,D\)\{sf\=CreateObject\(\"java\"\,\"java\.io\.File|Arguments\=xcmd\.text|asp cmd shell|Maceo|TEXTAREA id\=TEXTAREA1 name\=SqlQuery|CMD Bilgileri|sbusqlmod|php assert\(\$\_POST\[|oWshShellNet\.UserName|PHP C0nsole|rhtools|WinX Shell|system\(\$\_GET\[\'cmd\'|Successfully uploadet|\'Are you sure delete|sbusqlcmd|CFSWITCH EXPRESSION\=\#Form\.chopper|php\\HFile|\"ws\"\+\"cr\"\+\"ipt\.s\"\+\"hell\"|eval\(request\(|string rootkey|uZE Shell|Copyed success\!|InStr\(\"\$rar\$mdb\$zip\$exe\$com\$ico\$\"|Folder dosen\'t exists|Buradan Dosya Upload|echo passthru\(\$\_GET\[\'cmd\'|javascript:Bin\_PostBack|The file you want Downloadable|arguments\=\"/c \#cmd\#\"|cmdshell|AvFBP8k9CDlSP79lDl|AK-74 Security Team Web Shell|cfexecute name \= \"\#Form\.cmd\#\"|execute any shell commandn|Gamma Web Shell|System\.Reflection\.Assembly\.Load\(Request\.BinaryRead\(int\.Parse\(Request\.Cookies|fcreateshell|bash to execute a stack overflow|Safe Mode Shell|ASPX Shell|dingen\.php|azrailphp|\$\_POST\[\'sa\']\(\$\_POST\[\'sb\']\)|AspSpy|ntdaddy|\.HitU\. team|National Cracker Crew|eval\(base64\_decode\(\$\_REQUEST\[\'comment\'|Rootshell|geshi\\tsql\.php|tuifei\.asp|GRP WebShell|No Permission :\(|powered by zehir|will be delete all|WebFileManager Browsing|Dive Shell|diez\=server\.urlencode|@eval\(\$\_POST\[\'|ifupload\=\"ItsOk\"|eval\(request\.item|\(eval request\(|wsshn\.username|connect to reDuh|eval\(gzinflate\(base64\_decode|Ru24PostWebShell|ASPXTOOL\"|aspshell|File upload successfully you can download here|eval request\(|if\(is\_uploaded\_file\(\$HTTP|Sub RunSQLCMD|STNC WebShell|doosib|WinExec\(Target\_copy\_of\_cmd|php passthru\(getenv|win\.com cmd\.exe /c cacls\.exe|TUM HAKLARI SAKLIDIR|Created by PowerDream|Then Request\.Files\(0\)\.SaveAs\(Server\.MapPath\(Request|cfmshell|\{ Request\.Files\[0]\.SaveAs\(Server\.MapPath\(Request|\%execute\(request\(\"|php eval\(\$\_POST\[|lama\'s\'hell|RHTOOLS|data\=request\(\"dama\"|digitalapocalypse|hackingway\.tk|\.htaccess stealth web shell|strDat\.IndexOf\(\"EXEC \"|ExecuteGlobal request\(|Deleted file have finished|bin\_filern|CurrentVersionRunBackdoor|Chr\(124\)\.O\.Chr\(124\)|does not have permission to execute CMD\.EXE|G-Security Webshell|system\( \"\./findsock|configwizard|textarea style\=\"width:600\;height:200\" name\=\"cmd\"|ASPShell|repair/sam|BypasS Command eXecute|\%execute\(request\(|arguments\=\"/c \#hotmail|Coded by Loader|Call oS\.Run\(\"win\.com cmd\.exe|DESERTSUN SERVER CRASHER|ASPXSpy|cfparam name\=\"form\.shellpath\"|IIS Spy Using ADSI|p4ssw0rD|WARNING: Failed to daemonise|C0mmand line|phpinfo\(\) function has non-permissible|letaksekarang|Execute Shell Command|DXGLOBALSHIT|IISSpy|execute request\(|Chmod Ok\!|Upload Gagal|awen asp\.net|execute\(request\(\"|oSNet\.ComputerName"`
	}else{
		regexString = *customRegex
	}

	r := regexp.MustCompile(regexString)

	totalScanned := 0
	matched := 0
	cleared := 0

	filepath.Walk(*dir, func(path string, f os.FileInfo, err error) error {
		fileHandle, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer fileHandle.Close()

		fileScanner := bufio.NewScanner(fileHandle)

		fileMatches := make(map[string]int)
		totalScanned +=1

		for fileScanner.Scan() {
			matches := r.FindStringSubmatch(strings.ToLower(fileScanner.Text()))
			if len(matches) > 0{
				for _, it := range matches{
					fileMatches[it] += 1
				}
			}
		}

		if len(fileMatches) == 0 {
			cleared +=1
		}else if len(fileMatches) != 0 {
			matched +=1
		}

		if *output == "" || *output == "csv"{
			fmt.Println(fmt.Sprintf("%s|%v", path, fileMatches))

		}else if *output == "json"{
			Jdata := JsonOut{}
			Jdata.FilePath = path
			Jdata.Matches = fileMatches
			data, err := json.Marshal(Jdata)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", data)

		}else{
			fmt.Println(fmt.Sprintf("%s|%v", path, fileMatches))
		}

		return nil
	})

	metrics := Metrics{}
	metrics.Scanned = totalScanned
	metrics.Clear = cleared
	metrics.Matched = matched
	metrics.ScannedDir = *dir
	metrics.ScanTime = time.Since(start).Minutes()

	data, err := json.Marshal(metrics)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
