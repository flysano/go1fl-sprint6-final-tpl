package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go1fl-sprint6-final/internal/service"
)

func IndexHendler(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "Не удалось открыть файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}

func UploaderHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка разбора формы", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Не удалось получить файл", http.StatusInternalServerError)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	converted, err := service.Converter(string(data))
	if err != nil {
		http.Error(w, "Ошибка  конвертации", http.StatusInternalServerError)
	}

	ext := filepath.Ext(header.Filename)
	t := time.Now().UTC().Format("01022006_150405")
	filename := fmt.Sprintf("%s%s", t, ext)

	if err := os.WriteFile(filename, []byte(converted), 0644); err != nil {
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(converted))
}
