package main

import (
	"fmt"
	"github.com/mgutz/ansi"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

func main() {

	var filePath string
	pflag.StringVarP(&filePath, "file", "f", "", "Path to a single file, you want to exposed")

	var dirPath string
	pflag.StringVarP(&dirPath, "dir", "d", "", "Path to the directory, you want to exposed")

	var outPath string
	pflag.StringVarP(&outPath, "expose", "e", "", "Url path, where the file/directory will be exposed")

	var port string
	pflag.StringVarP(&port, "port", "p", "8080", "Port to run the server on. Default is 8080")

	pflag.Parse()

	if filePath != "" && dirPath != "" {
		fmt.Printf("%sError parsing flags: Only one of --file and --dir can be set!\n", ansi.LightRed)
		os.Exit(2)
	}

	isFileMode := filePath != ""

	if isFileMode {
		reg, _ := regexp.Compile("\\.[a-zA-Z0-9]+$")
		if !reg.MatchString(filePath) {
			fmt.Printf("%sError parsing filepath: Flag --file must be a valid filepath!\n", ansi.LightRed)
			os.Exit(3)
		}
		if !reg.MatchString(outPath) {
			fmt.Printf("%sError parsing exposepath: Flag --expose must be a valid filepath!\n", ansi.LightRed)
			os.Exit(4)
		}
		http.HandleFunc(outPath, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Expires", time.Unix(0, 0).Format(time.RFC1123))
			w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("X-Accel-Expires", "0")
			http.ServeFile(w, r, filePath)
		})
	} else {
		if !strings.HasSuffix(dirPath, "/") {
			dirPath += "/"
		}
		if !strings.HasSuffix(outPath, "/") {
			outPath += "/"
		}
		http.Handle(outPath, http.StripPrefix(outPath, http.FileServer(http.Dir(dirPath))))
	}

	msg := ansi.Reset + "Exposing "
	if isFileMode {
		msg += fmt.Sprintf("file '%s%s%s' ", ansi.LightCyan, filePath, ansi.Reset)
	} else {
		msg += fmt.Sprintf("'%s%s%s' ", ansi.LightCyan, dirPath, ansi.Reset)
	}
	msg += fmt.Sprintf("on '%shttp://localhost:%s%s%s'\n", ansi.LightGreen, port, outPath, ansi.Reset)

	fmt.Println(msg)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(ansi.LightRed+"Error starting server:", err)
	}
}
