# Exit Code Checker
This tool returns a designated exit code.
## Getting Started
### Dependencies
* Mac
* Linux
* Windows
### Installing
* go build -o exitcode main.go
### Executing program
Mac, linux:
```
$ ./exitcode 123
$ echo $?
123
```
Windows:
```
C:\ exitcode.ext 123
C:\> echo %ERRORLEVEL%
123
```
## Author
Shiro Konaka
## Version History
* 0.1
    * Initial Release