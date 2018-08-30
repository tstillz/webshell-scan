# About Webshell Scan
Webshell scan is a cross platform standalone binary that recursively scans through a specified directory with either user defined or default regex. 

This tool is related to the following write up:
 
https://blog.stillztech.com/2018/08/analyzing-and-detecting-web-shells.html

To test effectiveness of the scanner, it was tested against the tennc webshell repo: 

https://github.com/tennc/webshell 

### Requirements
None. Simply download the binary for your OS, supply the directory you wish to scan and let it rip.

### Running the binary
Running `webscan.exe` with no arguments shows the support flags:

	C:\Users\beastmode> webscan.exe
	
	Options:
	  -dir string
        	Directory to scan for webshells 
      -regex string
        	Override default regex with your own
      -output string
            json or csv (pipe separated). Default is csv
            
The only required argument is `dir`, but you can override system defaults if you wish. 
	
The output of the scan we be written to console. Example below:

	C:\Users\beastmode> webscan.exe -dir C:\tennc

	C:\tennc\webshell-master\138shell\C\Crystal.txt|map[exec(:7 eval(:1 cmd:15]
    C:\tennc\webshell-master\138shell\C\CyberSpy5.Asp.txt|map[]
    C:\tennc\webshell-master\138shell\C\c100.txt|map[exec(:41 cmd:38 eval(:3]
    C:\tennc\webshell-master\138shell\C\c2007.php.txt|map[cmd:32 exec(:32 eval(:3]
    C:\tennc\webshell-master\138shell\C\c99(1).php.txt|map[cmd:21 eval(:3 exec(:7]
    C:\tennc\webshell-master\138shell\C\c99.txt|map[cmd:30 exec(:11 eval(:2]

### Custom regex
You can also supply your own regex if you have some other specific regex term you're looking for:

    C:\Users\beastmode\webscan.exe -dir C:\tennc -regex "eval\\(|cmd|exec\\("
    
### Building the project
If you decide to modify the source code, you can build it using the following commands:

    cd <path-to-project>
    ## Windows
    GOOS=windows GOARCH=386 go build -o webscan.exe main.go
    ## Linux
    GOOS=linux GOARCH=386 go build -o webscan main.go
    ## Darwin
    GOOS=darwin GOARCH=386 go build -o webscan main.go
