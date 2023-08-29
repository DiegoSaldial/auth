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
	key := `
	{
		"type": "service_account",
		"project_id": "imagenesstore-sladia",
		"private_key_id": "9aa71601ad8ea878b85a101e82592a1299854a50",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEugIBADANBgkqhkiG9w0BAQEFAASCBKQwggSgAgEAAoIBAQC+tGi5RQ/AVMmL\nUBpT6+jLB6PSA5Lt9SCWuxXWr9m+D4hzzn9vtU4rLbIXCR5YIAqnCAcar4xrAzcF\n3bvz/FTxcX+BlibXE609lfK20GOHLpskWNk7y9vvmXnR6WBwj/egSZBT7/1g3Jvp\nxfdGcWNsXpolhPMzHhnchM7tgjjghYt0R7/IKBHDEQLgfSMDVTIYklp6MPgE2Esl\nm9OQKw3qEmTvh0nIrXz10hJdpCwWVbwkF8KpuOJMrhxG2I7SfgrN2kH3BYk8gOfI\npk1bS0VMEwm13AiQA2ruia0Gg2tJrdjma01tHlbzko+Kr34KAI1UvsJI2Nj7Av+Q\n6M7oVHsRAgMBAAECgf8g6TTDFmBe9KlYrkwuyRHt9+3Np6dILD3fDRh1LLrxXrzg\n7YnTKPbpAdX8vVOOpF9bNqYsLR+Oi3VAQR2P7l1eQ1vr8dnrEPpAT5EJTZtkqZcI\nGk06HiyywUnO9TUQx0zQ+2YvkQ0Z2FTSeQrciNOvDimWRcb6duqZPYVuTr4brwpv\nYleVYHISf+YLgy+C+HvU511tw8zfTvcAG3XfDoEFmaFxGoinAD6xQBxTSPPOQqM2\nu6N9vIFotcwb+xkBp6R9ZQWgLcfxsxS/CBYsjmrPpTgMzq02CyWX4Uv00GH+WDBH\neagNshgkOsPvA7jUjc1qHWce6CgLUOqg4a6pHiMCgYEA4p6j6oJ1khogILnDHktL\nv8NPzD582v7HwIkEiMYjLVATk1Iwn+7OpuzAieRwlLFXG6K3L0ZvECrWz5Sw3lxY\n1aFFPvhiWHoutWRw5o9FUV4k5zxvhwR8ZRl05w51wU0QME9HBaSV2UbQCHUowXtO\nT+e372zIVkxEPl0sCzsXjR8CgYEA123GRLayNtRA5mmlj4f27m/8sZ9te/Qqv8Do\n/URtTtItkNI3oWoRQ5bxwwWvX+/9Gf3lxtGk7eofk/igZ99M59lGI3Ehs+3GwHBN\nx6Xg6zHTiaR/luC0lEdsj0pF4C3xUYk2gA4cyMmYEnsLzZCwqOwsYyrt4EHaN2Ia\nsE97wc8CgYBBV/7psr5V52SWKxog1RM9cwLCYM49kzNjx95f1cn02d8bHprYstIx\nfZiy6gSwS0ZTuKJbZlLF9aRE7JnnM/eFed/unU42Ntza/uAzCuKw6JV5e4qpAtkf\niayUpy27FA5z75gJ/4AZy3pWxfl/eJ+HGMiHp/VINOuEk/cMMfjDfwKBgG3x2Vli\nQEvZQZIWYueAxZZk0vs32WGEeHjKlF4FmR+8BoM9tiSMBGmRVKtqcFmvGmY0fRte\nDSVa8mjLX9oDTbWZ3sDh0QiKFoBKEHgPAV4nXMWHjZJL9f0jApWSm1zRNbKeK5Fp\nImma8SClaZ2s9WwQtIqb8zEbiqnvSJ2owwljAoGAElm1PqrB24NV/xEyPRoB/6PM\natFzcCLlCXhTMfRWJOF+Fw5IR9tjVwMmX3D6yYZUODcNgnqcEPtodigS3HaTHagp\nRQ/UtDvAgqFY6tJec8Mcfv87MA9e1G5oBhmrFI8IvhV0/hQDuM1W9k6WrXNDqBaO\nQQB6JilrWcVYr4R5+lA=\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-f4xrh@imagenesstore-sladia.iam.gserviceaccount.com",
		"client_id": "104554994921687289504",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-f4xrh%40imagenesstore-sladia.iam.gserviceaccount.com",
		"universe_domain": "googleapis.com"
	  }
	`
	return []byte(key)
}

func SubirImagen(base64, filename string) (string, error) {
	ctx := context.Background()
	if !strings.HasPrefix(base64, "data:") {
		return "", errors.New("base64 no valido")
	}

	opt := option.WithCredentialsJSON(getJson())
	srv, err := storage.NewService(ctx, opt)
	if err != nil {
		return "", err
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
	fmt.Println("url: ", url)
	return url, nil
}
