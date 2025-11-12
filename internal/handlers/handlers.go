package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(filepath.Join(".", "index.html"))
	if err != nil {
		http.Error(w, "page loading error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "file download error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "file reading error", http.StatusInternalServerError)
		return
	}
	result, err := service.ConvertAutomatically(string(data))
	if err != nil {
		http.Error(w, "conversion error", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(handler.Filename)

	timeStr := time.Now().UTC().String()
	timeStr = strings.ReplaceAll(timeStr, ":", "-")
	timeStr = strings.ReplaceAll(timeStr, " ", "_")
	newFileName := fmt.Sprintf("result_%s%s", timeStr, ext)

	newFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "file creation error", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.Write([]byte(result))
	if err != nil {
		http.Error(w, "file write error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Файл успешно конвертирован!\n\nИсходный текст:\n%s\n\nРезультат:\n%s", string(data), result)

}
