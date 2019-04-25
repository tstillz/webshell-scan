# About web shell scan
Web shell scan is a cross platform standalone binary that recursively scans through a specified directory with either user defined or default regex. 
Web shell scan utilizes a pool of go routines (10 total) to read from a channel and speed up the scanner. Note, the regex supplied with the scanner isn't 100% and does not guarantee it will find every web shell on disk.  

This tool is related to the following write up:
 
https://blog.stillztech.com/2018/08/analyzing-and-detecting-web-shells.html

To test effectiveness of the scanner, it was tested against the tennc web shell repo: 

https://github.com/tennc/webshell

### Requirements
None! Simply download the binary for your OS, supply the directory you wish to scan and let it rip.

### Running the binary
Running `webscan` with no arguments shows the following arguments:

	/Users/beastmode$ ./webscan
	Options:
	    -dir string
          	Directory to scan for web shells
        -exts string
          	Specify extensions to target. Multiple extensions should be passed with pipe separator (asp|aspx|php|cfm). Default is all extensions
        -raw_contents
          	If a match is found, grab the raw contents and base64 + gzip compress the file into the JSON object.
        -regex string
          	Override default regex with your own
        -size int
          	Specify max file size to scan (default is 10 MB) (default 10)
            
The only required argument is `dir`, but you can override the program defaults if you wish. 
	
The output of the scan will be written to console. Example below (For best results, send stdout to a json file and review/post process offline):

	/Users/beastmode$ ./webscan -dir /Users/beastmode/webshell-master

	{"filePath":"/Users/beastmode/webshell-master/xakep-shells/PHP/wacking.php.php.txt","size":142739,"md5":"9c5bb5e3a46ec28039e8986324e42792","timestamps":{"birth":"2019-02-03 02:02:22","created":"2019-03-17 13:18:52","modified":"2019-02-03 02:02:22","accessed":"2019-04-25 01:19:47"},"matches":{"eval(":2}}
	
    ### With STDOUT:
    
    /Users/beastmode$ ./webscan -dir /Users/beastmode/webshell-master -raw_contents=true > scan_results.json

Once the scanner finishes, it will output the overall scan metrics to STDOUT, as shown in the example below:

    {"scanned":1575,"matches":501,"noMatches":1074,"directory":"/Users/beastmode$/Downloads/webshell-master","scanDuration":1.2528534718166666,
    "systemInfo":{"hostname":"beastmode","envVars":["...."],"username":"beastmode","userID":"12345","realName":"","userHomeDir":"/Users/beastmode"}}

### Custom regex
You can also supply your own regex if you have some specific regex pattern you're looking for:

    ./webscan -dir=/opt/https -regex="eval\\(|cmd|exec\\(" -size=5 -raw_contents=true -exts=php|jsp
    
### Building the project
If you decide to modify the source code, you can build the project using the following commands:

    cd <path-to-project>
    ## Windows
    GOOS=windows GOARCH=386 go build -o webscan_windows.exe main.go
    ## Linux
    GOOS=linux GOARCH=386 go build -o webscan_linux main.go
    ## Darwin
    GOOS=darwin GOARCH=386 go build -o webscan_darwin main.go
