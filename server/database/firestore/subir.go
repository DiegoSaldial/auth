package firestore

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/vincent-petithory/dataurl"
	"google.golang.org/api/option"
	"google.golang.org/api/storage/v1"
)

func getJson() []byte {
	key := os.Getenv("FIREBASE_STORAGE_CRED")
	return []byte(key)
}

func SubirImagen(base64, filename string, resize bool) (string, error) {
	ctx := context.Background()
	if !strings.HasPrefix(base64, "data:") {
		return "", errors.New("base64 no valido")
	}

	opt := option.WithCredentialsJSON(getJson())
	srv, err := storage.NewService(ctx, opt)
	if err != nil {
		return "", err
	}

	if resize {
		newb64, err := Resize(base64, 400)
		if err == nil {
			base64 = newb64
		}
	}

	mydataurl, err := dataurl.DecodeString(base64)
	if err != nil {
		return "", err
	}

	buck := os.Getenv("FIREBASE_STORAGE_BUCKET")
	carpeta := os.Getenv("FIREBASE_STORAGE_FOLDER")
	objeto := storage.Object{
		Name: carpeta + "/" + filename,
	}

	upload := srv.Objects.Insert(buck, &objeto).Media(strings.NewReader(string(mydataurl.Data)))
	_, err = upload.Do()
	if err != nil {
		return "", err
	}
	base := "https://firebasestorage.googleapis.com/v0/b/imagenesstore-sladia.appspot.com/o/taxisapp%2F"
	url := fmt.Sprintf("%s%s?alt=media", base, filename)
	// fmt.Println("url: ", url)
	return url, nil
}
