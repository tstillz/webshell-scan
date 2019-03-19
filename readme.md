# About Web shell scan
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
            
The only required argument is `dir`, but you can override system defaults if you wish. 
	
The output of the scan we be written to console. Example below (send stdout to a file and review offline is best):

	/Users/beastmode$ ./webscan -dir /Users/beastmode/webshell-master

	{"filePath":"/Users/beastmode/webshell-master/138shell/C/c99.txt","size":152950,"md5":"7a3cc460306cbf50b4f230884624acb0","matches":{"eval(":2}}
    {"filePath":"/Users/beastmode/webshell-master/138shell/F/Fatalshell.php.txt","size":16375,"md5":"b15583f4eaad10a25ef53ab451a4a26d","matches":{"eval(":1}}
    {"filePath":"/Users/beastmode/webshell-master/138shell/D/Dx.txt","size":111606,"md5":"9cfe372d49fe8bf2fac8e1c534153d9b","matches":{"eval(":4}}
    {"filePath":"/Users/beastmode/webshell-master/138shell/D/DxShell_hk.php.txt","size":111841,"md5":"8705c4495a9fd1811f31e2507f93e63e","matches":{"eval(":4}}

    ### With STDOUT:
    
    /Users/beastmode$ ./webscan -dir /Users/beastmode/webshell-master -raw_contents=true > scan_results.json

### Custom regex
You can also supply your own regex if you have some other specific regex term you're looking for:

    ./webscan -dir /opt/https -regex "eval\\(|cmd|exec\\(" -size=5 -raw_contents=true -exts=php|jsp
    
### Building the project
If you decide to modify the source code, you can build it using the following commands:

    cd <path-to-project>
    ## Windows
    GOOS=windows GOARCH=386 go build -o webscan_windows.exe main.go
    ## Linux
    GOOS=linux GOARCH=386 go build -o webscan_linux main.go
    ## Darwin
    GOOS=darwin GOARCH=386 go build -o webscan_darwin main.go
