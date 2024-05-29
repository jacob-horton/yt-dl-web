package main

import (
	_ "embed"
	"io"
	"net/http"
	"os/exec"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.HandleFunc("/download", downloadMP3)

	err := http.ListenAndServe("0.0.0.0:3333", mux)
	if err != nil {
		panic(err)
	}
}

func downloadMP3(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, "method not allowed")
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	url := string(body)
	data := getMP3(url)

	w.Header().Add("Content-Disposition", "attachment")
	w.Write(data)
}

func getMP3(url string) []byte {
	out, err := exec.Command("yt-dlp", "-o", "-", "-x", "-f", "m4a", url).Output()
	if err != nil {
		panic(err)
	}

	return out
}
