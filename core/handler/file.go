package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/labstack/echo/v4"
	"io"
	"moechat/core/database"
	"net/http"
	"time"
)

func File(c echo.Context) error {
	switch c.Request().Method {
	case http.MethodPost:

		// Parse the form data, limit the size to 10MB
		err := c.Request().ParseMultipartForm(10 << 20)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Unable to parse form")
		}
		// Get the uploaded files
		form := c.Request().MultipartForm
		files := form.File["files"] // 'files' is the name attribute of the input field

		if len(files) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "No files uploaded")
		}
		uploadedFilesData := make([]*database.File, 0, len(files))
		for _, fileHeader := range files {
			// Open the uploaded file
			file, err := fileHeader.Open()
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Unable to open file")
			}
			defer file.Close()

			// Read the file content
			fileData, err := io.ReadAll(file)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Unable to read file")
			}

			// Generate a hash for the file (SHA256)
			hash := generateHash(fileData)

			// Create a new File instance
			newFile := &database.File{
				Hash:      hash,
				Email:     `c.Get("email").(string)`, //todo
				Filename:  fileHeader.Filename,
				MimeType:  fileHeader.Header.Get("Content-Type"),
				Data:      fileData,
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}

			// Save the file record to the database
			if _, err := database.DB.NamedQuery(`INSERT OR IGNORE INTO file 
    	(hash, email, filename, mime_type, data, created_at, updated_at)
		VALUES (:hash, :email, :filename, :mime_type, :data, :created_at, :updated_at)`, newFile); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Unable to save file to database")
			}
			newFile.Data = nil
			uploadedFilesData = append(uploadedFilesData, newFile)
		}

		return c.JSON(http.StatusOK, uploadedFilesData)
	case http.MethodGet:
		var file database.File
		err := database.DB.Get(&file, "SELECT * FROM file WHERE hash = ?", c.QueryParam("hash"))
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, file.MimeType, file.Data)
	}
	return echo.ErrMethodNotAllowed
}

func generateHash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
