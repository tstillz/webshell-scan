package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type JsonOut struct {
	FilePath    string         `json:"filePath"`
	Size        int64          `json:"size"`
	MD5         string         `json:"md5"`
	Matches     map[string]int `json:"matches"`
	RawContents string         `json:"rawContents,omitempty"`
}
type Metrics struct {
	Scanned    int     `json:"scanned"`
	Matched    int     `json:"matches"`
	Clear      int     `json:"noMatches"`
	ScannedDir string  `json:"directory"`
	ScanTime   float64 `json:"scanDuration"`
}

var matched = 0
var cleared = 0

func md5HashFile(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}
func compressEncode(filePath string, fileSize int64) string {

	fileItem, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer fileItem.Close()

	buf := make([]byte, fileSize)
	fReader := bufio.NewReader(fileItem)
	fReader.Read(buf)

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(buf); err != nil {
		fmt.Println(err)
		return ""
	}
	if err := gz.Flush(); err != nil {
		fmt.Println(err)
		return ""
	}
	if err := gz.Close(); err != nil {
		fmt.Println(err)
		return ""
	}

	readBuf, _ := ioutil.ReadAll(&b)
	imgBase64Str := base64.StdEncoding.EncodeToString(readBuf)

	return imgBase64Str

}
func Scan_worker(id int, filesToScan <-chan string, r regexp.Regexp, wg *sync.WaitGroup, rawContents bool) {
	for j := range filesToScan {
		//fmt.Println("Worker:", id, "File:", j)
		//fmt.Println(len(filesToScan))

		fileHandle, err := os.Open(j)
		if err != nil {
			log.Fatal(err)
		}
		defer fileHandle.Close()

		fileScanner := bufio.NewScanner(fileHandle)
		fileMatches := make(map[string]int)

		for fileScanner.Scan() {
			matches := r.FindStringSubmatch(strings.ToLower(fileScanner.Text()))
			if len(matches) > 0 {
				for _, it := range matches {
					fileMatches[it] += 1
				}
			}
		}

		if len(fileMatches) != 0 {
			matched = matched + 1
			Jdata := JsonOut{}
			Jdata.FilePath = j
			Jdata.Matches = fileMatches
			fi, err := os.Stat(j)
			if err != nil {
				log.Println(err)
			}
			Jdata.Size = fi.Size()
			fHash, err := md5HashFile(j)
			Jdata.MD5 = fHash

			if rawContents {
				Jdata.RawContents = compressEncode(j, Jdata.Size)
			}

			data, err := json.Marshal(Jdata)

			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", data)
		} else if len(fileMatches) == 0 {
			cleared = cleared + 1
		}
	}
	wg.Done()
}

func main() {
	filesToScan := make(chan string, 100000)

	start := time.Now()
	var dir = flag.String("dir", "", "Directory to scan for webshells")
	var customRegex = flag.String("regex", "", "Override default regex with your own")
	var size = flag.Int64("size", 10, "Specify max file size to scan (default is 10 MB)")
	var exts = flag.String("exts", "", "Specify extensions to target. Multiple extensions should be passed with pipe separator (asp|aspx|php|cfm). Default is all extensions")
	var rawContents = flag.Bool("raw_contents", false, "If a match is found, grab the raw contents and base64 + gzip compress the file into the JSON object.")
	flag.Parse()

	if *dir == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	regexString := ""

	if *customRegex == "" {
		regexString = `Filesman|eval\(|Load\(Request\.BinaryRead\(int\.Parse\(Request\.Cookies|Html \= Replace\(Html\, \"\%26raquo\;\"\, \"?\"\)|pinkok|internal class reDuh|c0derz shell|md5 cracker|umer rock|Function CP\(S\,D\)\{sf\=CreateObject\(\"java\"\,\"java\.io\.File|Arguments\=xcmd\.text|asp cmd shell|Maceo|TEXTAREA id\=TEXTAREA1 name\=SqlQuery|CMD Bilgileri|sbusqlmod|php assert\(\$\_POST\[|oWshShellNet\.UserName|PHP C0nsole|rhtools|WinX Shell|system\(\$\_GET\[\'cmd\'|Successfully uploadet|\'Are you sure delete|sbusqlcmd|CFSWITCH EXPRESSION\=\#Form\.chopper|php\\HFile|\"ws\"\+\"cr\"\+\"ipt\.s\"\+\"hell\"|eval\(request\(|string rootkey|uZE Shell|Copyed success\!|InStr\(\"\$rar\$mdb\$zip\$exe\$com\$ico\$\"|Folder dosen\'t exists|Buradan Dosya Upload|echo passthru\(\$\_GET\[\'cmd\'|javascript:Bin\_PostBack|The file you want Downloadable|arguments\=\"/c \#cmd\#\"|cmdshell|AvFBP8k9CDlSP79lDl|AK-74 Security Team Web Shell|cfexecute name \= \"\#Form\.cmd\#\"|execute any shell commandn|Gamma Web Shell|System\.Reflection\.Assembly\.Load\(Request\.BinaryRead\(int\.Parse\(Request\.Cookies|fcreateshell|bash to execute a stack overflow|Safe Mode Shell|ASPX Shell|dingen\.php|azrailphp|\$\_POST\[\'sa\']\(\$\_POST\[\'sb\']\)|AspSpy|ntdaddy|\.HitU\. team|National Cracker Crew|eval\(base64\_decode\(\$\_REQUEST\[\'comment\'|Rootshell|geshi\\tsql\.php|tuifei\.asp|GRP WebShell|No Permission :\(|powered by zehir|will be delete all|WebFileManager Browsing|Dive Shell|diez\=server\.urlencode|@eval\(\$\_POST\[\'|ifupload\=\"ItsOk\"|eval\(request\.item|\(eval request\(|wsshn\.username|connect to reDuh|eval\(gzinflate\(base64\_decode|Ru24PostWebShell|ASPXTOOL\"|aspshell|File upload successfully you can download here|eval request\(|if\(is\_uploaded\_file\(\$HTTP|Sub RunSQLCMD|STNC WebShell|doosib|WinExec\(Target\_copy\_of\_cmd|php passthru\(getenv|win\.com cmd\.exe /c cacls\.exe|TUM HAKLARI SAKLIDIR|Created by PowerDream|Then Request\.Files\(0\)\.SaveAs\(Server\.MapPath\(Request|cfmshell|\{ Request\.Files\[0]\.SaveAs\(Server\.MapPath\(Request|\%execute\(request\(\"|php eval\(\$\_POST\[|lama\'s\'hell|RHTOOLS|data\=request\(\"dama\"|digitalapocalypse|hackingway\.tk|\.htaccess stealth web shell|strDat\.IndexOf\(\"EXEC \"|ExecuteGlobal request\(|Deleted file have finished|bin\_filern|CurrentVersionRunBackdoor|Chr\(124\)\.O\.Chr\(124\)|does not have permission to execute CMD\.EXE|G-Security Webshell|system\( \"\./findsock|configwizard|textarea style\=\"width:600\;height:200\" name\=\"cmd\"|ASPShell|repair/sam|BypasS Command eXecute|\%execute\(request\(|arguments\=\"/c \#hotmail|Coded by Loader|Call oS\.Run\(\"win\.com cmd\.exe|DESERTSUN SERVER CRASHER|ASPXSpy|cfparam name\=\"form\.shellpath\"|IIS Spy Using ADSI|p4ssw0rD|WARNING: Failed to daemonise|C0mmand line|phpinfo\(\) function has non-permissible|letaksekarang|Execute Shell Command|DXGLOBALSHIT|IISSpy|execute request\(|Chmod Ok\!|Upload Gagal|awen asp\.net|execute\(request\(\"|oSNet\.ComputerName"`
	} else {
		regexString = *customRegex
	}

	r := regexp.MustCompile(regexString)

	totalScanned := 0

	filepath.Walk(*dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if f.Size() < (*size * 1024 * 1024) {
			//fmt.Println(f.Size(), *size * 1024 * 1024)
			//fmt.Println(path, f.Size())

			/// Scan all files with all extensions
			if *exts == "" {
				filesToScan <- path
				totalScanned = totalScanned + 1

				/// Scan files with specific extensions
			} else {
				items := strings.SplitAfter(*exts, "|")
				for _, e := range items {
					if strings.HasSuffix(path, e) {
						filesToScan <- path
						totalScanned = totalScanned + 1
					}
				}
			}
		}
		return nil
	})

	close(filesToScan)

	var wg sync.WaitGroup
	for w := 1; w <= 10; w++ {
		wg.Add(1)
		go Scan_worker(w, filesToScan, *r, &wg, *rawContents)
	}
	wg.Wait()

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
