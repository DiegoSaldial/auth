package firestore

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

func Resize(base64DataURI string, nuevoAncho int) (string, error) {
	coI := strings.Index(string(base64DataURI), ",")
	rawImage := string(base64DataURI)[coI+1:]
	unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))
	res := bytes.NewReader(unbased)

	var img image.Image

	switch strings.TrimSuffix(base64DataURI[5:coI], ";base64") {
	case "image/png":
		pngI, err := png.Decode(res)
		if err != nil {
			return "", err
		}
		img = pngI
	case "image/jpeg":
		jpgI, err := jpeg.Decode(res)
		if err != nil {
			return "", err
		}
		img = jpgI
	default:
		imge, _, err := image.Decode(res)
		if err != nil {
			return "", err
		}
		img = imge
	}

	// nuevoAncho := 200

	// Calcular el nuevo alto manteniendo la relación de aspecto
	anchoOriginal := img.Bounds().Dx()
	altoOriginal := img.Bounds().Dy()
	nuevoAlto := (altoOriginal * nuevoAncho) / anchoOriginal

	if anchoOriginal <= nuevoAncho {
		return base64DataURI, nil
	}

	// Crear una nueva imagen con las dimensiones deseadas
	imgRedimensionada := image.NewRGBA(image.Rect(0, 0, nuevoAncho, nuevoAlto))

	// Redimensionar la imagen mediante interpolación simple
	for y := 0; y < nuevoAlto; y++ {
		for x := 0; x < nuevoAncho; x++ {
			srcX := x * anchoOriginal / nuevoAncho
			srcY := y * altoOriginal / nuevoAlto
			imgRedimensionada.Set(x, y, img.At(srcX, srcY))
		}
	}

	// Guardar la imagen redimensionada
	/* outputFile, err := os.Create("imagen_redimensionada.jpg") // Reemplaza con la ruta deseada
	if err != nil {
		return "", err
	}
	defer outputFile.Close()
	jpeg.Encode(outputFile, imgRedimensionada, nil)
	fmt.Println("Imagen redimensionada guardada correctamente.") */

	var buffer bytes.Buffer
	err := jpeg.Encode(&buffer, imgRedimensionada, nil)
	if err != nil {
		return "", err
	}
	imgBase64 := base64.StdEncoding.EncodeToString(buffer.Bytes())

	// Construir la cadena base64 data URI
	// format := base64DataURI[5:coI]
	newBase64DataURI := fmt.Sprintf("data:image/jpeg;base64,%s", imgBase64)
	// fmt.Println(format)
	return newBase64DataURI, nil

}
