package handlers

import (
	"goHtmx/web/templates/documents"
	"io"
	"os"

	"github.com/labstack/echo/v4"
)

func DocumentPage(c echo.Context) error {
	return render(c, documents.DocumentPage())
}

func UploadDocument(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create("docs/temp/" + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return DocumentPage(c)
}

func CreateDocument() error {
	return nil
}
