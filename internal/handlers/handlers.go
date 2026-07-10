package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(res http.ResponseWriter, req *http.Request) {
	htmlFile, err := os.ReadFile("index.html")

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write(htmlFile)

}

func UploadHandler(res http.ResponseWriter, req *http.Request) {

	err := req.ParseMultipartForm(10 << 20)

	if err != nil {
		http.Error(res, "cannot parse form", http.StatusInternalServerError)
		return
	}

	file, header, err := req.FormFile("myFile")

	if err != nil {
		http.Error(res, "cannot get file from form", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		http.Error(res, "cannot read file", http.StatusInternalServerError)
		return
	}

	result, err := service.DetectAndConvert(string(data))

	if err != nil {
		http.Error(res, "cannot convert data", http.StatusInternalServerError)
		return
	}

	name := time.Now().UTC().Format("2026-01-02_15-04-05") + filepath.Ext(header.Filename)
	out, err := os.Create(name)

	if err != nil {
		http.Error(res, "cannot create file", http.StatusInternalServerError)
		return
	}

	defer out.Close()

	if _, err = io.WriteString(out, result); err != nil {
		http.Error(res, "cannot write file", http.StatusInternalServerError)
		return
	}

	if _, err := res.Write([]byte(result)); err != nil {
		http.Error(res, "cannot write file", http.StatusInternalServerError)
		return
	}

}
