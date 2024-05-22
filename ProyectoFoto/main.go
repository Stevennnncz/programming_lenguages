package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/tiff"
	"github.com/fogleman/gg"
)

var globalimg = "" // Variable global para almacenar la imagen codificada en base64

// Request representa la estructura de la solicitud para procesar la imagen
type Request struct {
	BgImgPath string  // Ruta de la imagen de fondo
	FontPath  string  // Ruta de la fuente
	FontSize  float64 // Tamaño de la fuente
	Text1     string  // Primer texto a dibujar
	Text2     string  // Segundo texto a dibujar
}

// EmailRequest representa la estructura de la solicitud para enviar el correo electrónico
type EmailRequest struct {
	Email string // Dirección de correo electrónico del destinatario
}

// TextOnImg procesa la imagen y dibuja el texto sobre ella
func TextOnImg(request Request) (image.Image, error) {
	// Abrir la imagen de fondo
	bgFile, err := os.Open(request.BgImgPath)
	if err != nil {
		return nil, err
	}
	defer bgFile.Close()

	var bgImage image.Image
	// Decodificar la imagen según su extensión
	switch extension := filepath.Ext(request.BgImgPath); extension {
	case ".png":
		bgImage, err = png.Decode(bgFile)
	case ".jpg", ".jpeg":
		bgImage, err = jpeg.Decode(bgFile)
	case ".tiff", ".tif":
		bgImage, err = tiff.Decode(bgFile)
	default:
		return nil, fmt.Errorf("formato de imagen no compatible")
	}
	if err != nil {
		return nil, err
	}

	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	// Crear un contexto de dibujo con las dimensiones de la imagen
	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	// Cargar la fuente con el tamaño especificado
	if err := dc.LoadFontFace(request.FontPath, request.FontSize); err != nil {
		return nil, err
	}

	x := float64(imgWidth / 2)

	// Validar el texto
	if !validateText(request.Text1, request.Text2) {
		return nil, errors.New("el texto debe tener un mínimo de 5 letras y una palabra, y un máximo de dos palabras y 12 caracteres")
	}

	// Calcular el tamaño óptimo del texto para abarcar el ancho completo de la imagen
	maxWidth := float64(imgWidth) - 60.0
	estimatedTextWidth1 := estimateTextWidth(request.Text1, request.FontSize)
	estimatedTextWidth2 := estimateTextWidth(request.Text2, request.FontSize)
	scale1 := maxWidth / estimatedTextWidth1
	scale2 := maxWidth / estimatedTextWidth2

	newFontSize1 := request.FontSize * scale1
	newFontSize2 := request.FontSize * scale2

	// Ajustar el tamaño de la fuente según sea necesario
	if request.Text1 != " " && request.Text2 == " " {
		if err := dc.LoadFontFace(request.FontPath, newFontSize1); err != nil {
			return nil, err
		}
	} else if request.Text1 == " " && request.Text2 != " " {
		if err := dc.LoadFontFace(request.FontPath, newFontSize2); err != nil {
			return nil, err
		}
	} else {
		if err := dc.LoadFontFace(request.FontPath, newFontSize1); err != nil {
			return nil, err
		}
		if err := dc.LoadFontFace(request.FontPath, newFontSize2); err != nil {
			return nil, err
		}
	}

	y1 := 40.0                                                 // Coordenada Y para el primer texto
	y2 := float64(imgHeight) - float64(dc.FontHeight()) - 20.0 // Coordenada Y para el segundo texto

	dc.SetColor(color.White)
	dc.DrawStringWrapped(request.Text1, x, y1, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	dc.DrawStringWrapped(request.Text2, x, y2, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}

// validateText verifica si el texto cumple con los requisitos mínimos y máximos
func validateText(text1, text2 string) bool {
	if (text1 != " ") && (text2 == " ") {
		if len(text1) < 5 {
			return false
		}
		words := strings.Fields(text1)
		if len(words) < 1 || len(words) > 2 {
			return false
		}
		if len(text1) > 12 {
			return false
		}

	} else if (text1 == " ") && (text2 != " ") {
		if len(text2) < 5 {
			return false
		}
		words := strings.Fields(text2)
		if len(words) < 1 || len(words) > 2 {
			return false
		}
		if len(text2) > 12 {
			return false
		}

	} else if (text1 != " ") && (text2 != " ") {
		if len(text1) < 5 && len(text2) < 5 {
			return false
		}
		words1 := strings.Fields(text1)
		words2 := strings.Fields(text2)
		if (len(words1) < 1 || len(words1) > 2) || (len(words2) < 1 || len(words2) > 2) {
			return false
		}
		if len(text1) > 12 || len(text2) > 12 {
			return false
		}
	}
	return true
}

// sendEmail envía un correo electrónico con la imagen adjunta
func sendEmail(imgBase64Str string, mail string) {

	// Datos de autenticación
	from := "stevenlenguajes@gmail.com"
	password := "igza rzzt lfcu uukw"

	// Configuración del servidor SMTP
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Destinatario
	to := mail

	// Decodificar la cadena Base64
	imageBytes, err := base64.StdEncoding.DecodeString(imgBase64Str)
	if err != nil {
		fmt.Println("Error al decodificar la cadena Base64:", err)
		return
	}

	// Crear el mensaje con el archivo adjunto
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	writer.WriteField("Subject", "Hola")
	writer.WriteField("Text", "Este es un correo de prueba enviado desde Go con una imagen adjunta.")

	// Adjuntar el archivo
	part, _ := writer.CreateFormFile("attachment", "image.jpg")
	part.Write(imageBytes)
	writer.Close()

	// Crear el mensaje MIME
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\nMIME-version: 1.0;\r\nContent-Type: multipart/mixed; boundary=%s\r\n\r\n--%s\r\nContent-Type: text/plain; charset=\"UTF-8\"\r\n\r\n%s\r\n--%s\r\nContent-Type: application/octet-stream; name=\"%s\"\r\nContent-Transfer-Encoding: base64\r\nContent-Disposition: attachment; filename=\"%s\"\r\n\r\n%s\r\n--%s--",
		to, "Hola", writer.Boundary(), writer.Boundary(), "Este es un correo de prueba enviado desde Go con una imagen adjunta.", writer.Boundary(), "image.jpg", "image.jpg", imgBase64Str, writer.Boundary())

	// Autenticación
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Envío del correo
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		fmt.Println("Error al enviar el correo:", err)
		return
	}
	fmt.Println("Correo enviado correctamente.")
}

// estimateTextWidth estima el ancho del texto basándose en el número de caracteres y el tamaño de la fuente
func estimateTextWidth(text string, fontSize float64) float64 {
	// Asumimos un ancho de aproximadamente 10 unidades por carácter
	return float64(len(text)) * fontSize * 0.5
}

// handleUpload maneja la subida de archivos y procesamiento de la imagen
func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("upload.html"))
		tmpl.Execute(w, nil)
		return
	}

	// Obtener el archivo subido
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error al obtener el archivo:", err)
		renderErrorPage(w, "Error al obtener el archivo")
		return
	}
	defer file.Close()

	// Guardar el archivo subido
	filePath := "./" + handler.Filename
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		renderErrorPage(w, "Error al abrir el archivo")
		return
	}
	defer f.Close()

	io.Copy(f, file)

	// Obtener la fuente seleccionada desde el formulario
	font := r.FormValue("font")

	// Crear la solicitud para procesar la imagen
	request := Request{
		BgImgPath: filePath,
		FontPath:  font, // Asignar la fuente seleccionada al campo FontPath
		FontSize:  24,
		Text1:     r.FormValue("text1"),
		Text2:     r.FormValue("text2"),
	}

	if len(request.Text1) == 0 {
		request.Text1 = " "
	} else if len(request.Text2) == 0 {
		request.Text2 = " "
	}

	// Procesar la imagen
	img, err := TextOnImg(request)
	if err != nil {
		fmt.Println("Error:", err)
		renderErrorPage(w, err.Error())
		return
	}

	// Codificar la imagen en base64
	var imgBase64 bytes.Buffer
	if err := jpeg.Encode(&imgBase64, img, nil); err != nil {
		fmt.Println("Error al codificar la imagen:", err)
		renderErrorPage(w, "Error al codificar la imagen")
		return
	}

	imgBase64Str := base64.StdEncoding.EncodeToString(imgBase64.Bytes())
	globalimg = imgBase64Str

	// Renderizar la página con la imagen procesada
	tmpl := template.Must(template.New("result").Parse(resultPage))
	tmpl.Execute(w, imgBase64Str)
}

// renderErrorPage renderiza una página de error con el mensaje proporcionado
func renderErrorPage(w http.ResponseWriter, errorMessage string) {
	type ErrorPageData struct {
		Message string
	}

	data := ErrorPageData{
		Message: errorMessage,
	}

	tmpl := template.Must(template.ParseFiles("error.html"))
	tmpl.Execute(w, data)
}

// HTML para la página de resultado
const resultPage = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Resultado</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f0f0f0;
            margin: 0;
            padding: 0;
        }
        .container {
            width: 80%;
            margin: 50px auto;
            background-color: #fff;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
        }
        h2 {
            color: #333;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>Imagen procesada</h2>
        <img src="data:image/jpeg;base64,{{ . }}" alt="Processed Image">
    </div>
</body>
</html>`

// emailHandler maneja la solicitud para enviar el correo electrónico
func emailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var request EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Solicitud no válida", http.StatusBadRequest)
		return
	}

	sendEmail(globalimg, request.Email)
	fmt.Fprintln(w, "Correo enviado correctamente.")
}

// Función principal para iniciar el servidor HTTP
func main() {
	http.HandleFunc("/", handleUpload)
	http.HandleFunc("/send-email", emailHandler)
	http.ListenAndServe(":8080", nil)
}
