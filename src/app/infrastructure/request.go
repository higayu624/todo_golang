package infrastructure

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Connect() {
	res, err := http.Get("https://youtube.googleapis.com/youtube/v3/playlists?part=snippet&id=RDMMRgKAFK5djSk&key=AIzaSyBL_3aORl9f_AenKMBhEyEuRC83jm7sWYw")
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return "Response failed"
	}
	if err != nil {
		return err
	}
	return body
}
