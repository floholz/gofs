package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/mgutz/ansi"
	"github.com/spf13/pflag"
)

func main() {

	var urlIn string
	pflag.StringVarP(&urlIn, "url", "u", "localhost:8080",
		"Host and Port to run the server on. Default is 'localhost:8080'.")

	pflag.Parse()

	inPath := pflag.Arg(0)

	urlIn = strings.TrimPrefix(urlIn, "https://")
	urlIn = strings.TrimPrefix(urlIn, "http://")

	urlParsed, urlErr := url.Parse("http://" + urlIn)
	if urlErr != nil {
		fmt.Printf("%sError parsing host: Flag --host must be a valid 'address:port/path' combination!\n", ansi.LightRed)
		os.Exit(1)
	}

	if urlParsed.Port() == "" {
		urlParsed.Host = urlParsed.Host + ":8080"
	}

	if urlParsed.Hostname() == "" {
		urlParsed.Host = "localhost" + urlParsed.Host
	}

	var isFileMode bool
	if inPath == "" {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("%sError retreiving working directory!\n", ansi.LightRed)
			os.Exit(2)
		}
		inPath = wd
		isFileMode = false
	} else {
		_isFileMode, fodErr := fileOrDir(inPath)
		if fodErr != nil {
			fmt.Printf("%sError parsing the in-path, --in must be a valid path to a file or directory!\n", ansi.LightRed)
			os.Exit(2)
		}
		isFileMode = _isFileMode
	}

	if isFileMode {
		_inPath := strings.Split(inPath, "/\\")
		fileName := _inPath[len(_inPath)-1]
		if urlParsed.Path == "" {
			urlParsed.Path = "/" + fileName
		}

		http.HandleFunc(urlParsed.Path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Expires", time.Unix(0, 0).Format(time.RFC1123))
			w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("X-Accel-Expires", "0")
			http.ServeFile(w, r, inPath)
		})
	} else {
		if !strings.HasSuffix(inPath, "/") {
			inPath += "/"
		}
		if !strings.HasSuffix(urlParsed.Path, "/") {
			urlParsed.Path += "/"
		}
		http.Handle(urlParsed.Path, http.StripPrefix(urlParsed.Path, http.FileServer(http.Dir(inPath))))
	}

	msg := ansi.Reset + "Exposing "
	if isFileMode {
		msg += "file "
	} else {
		msg += "directory "
	}
	msg += fmt.Sprintf("'%s%s%s' ", ansi.LightCyan, inPath, ansi.Reset)
	msg += fmt.Sprintf("on '%s%s%s'\n", ansi.LightGreen, urlParsed.String(), ansi.Reset)

	fmt.Println(msg)

	err := http.ListenAndServe(urlParsed.Host, nil)
	if err != nil {
		fmt.Println(ansi.LightRed+"Error starting server:", err)
	}
}

func fileOrDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return !fileInfo.IsDir(), err
}
