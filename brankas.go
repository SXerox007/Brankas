package main

import (
	"Brankas/base/utils"
	"Brankas/models/brankas"
	"Brankas/models/common"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// page struct
type Page struct {
	Title      string
	StaticHost string
	Auth       string
}

func GetFileUploadPage() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		page := &Page{
			Title:      "BRANKAS",
			StaticHost: "http://localhost:50051",
			Auth:       os.Getenv("AUTH"),
		}
		utils.OutputHTML(w, "./templates/brankas.html", page)
		return
	}
}

// upload file
func UploadFile() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(10 << 20)
		file, handler, err := r.FormFile("myImage")
		if err != nil {
			log.Println("Error:", err)
			utils.RespondWithJSON(w, http.StatusForbidden, common.CommonResponse{Success: false, Message: "Upload File Error.", Code: http.StatusForbidden, Data: nil})
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)

		if handler.Size > 8000000 {
			utils.RespondWithJSON(w, http.StatusForbidden, common.CommonResponse{Success: false, Message: "Error: Image size will not be larger than 8MB.", Code: http.StatusForbidden, Data: nil})
			return
		}

		attachmentType := strings.Split(handler.Header.Get("Content-Type"), ";")
		if attachmentType[0] == "image/jpeg" || attachmentType[0] == "image/png" || attachmentType[0] == "image/jpg" {
			// read all of the contents of our uploaded file into a
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				utils.RespondWithJSON(w, http.StatusForbidden, common.CommonResponse{Success: false, Message: "Something went wrong.", Code: http.StatusForbidden, Data: err})
				return
			}
			// write this byte array to our temporary file
			err = ioutil.WriteFile(handler.Filename, fileBytes, 0644)
			if err != nil {
				utils.RespondWithJSON(w, http.StatusForbidden, common.CommonResponse{Success: false, Message: "Something went wrong.", Code: http.StatusForbidden, Data: err})
			} else {
				// save image details in db
				if _, err := brankas.CreateImageDetailsRow(brankas.ImageDetails{
					FileName:    handler.Filename,
					ContentType: attachmentType[0],
					Size:        handler.Size,
				}); err != nil {
					utils.RespondWithJSON(w, http.StatusConflict, common.CommonResponse{Success: false, Message: "DB Error.", Code: http.StatusConflict, Data: err})
				} else {
					utils.RespondWithJSON(w, http.StatusOK, common.CommonResponse{Success: true, Message: "Success.", Code: http.StatusOK, Data: nil})
				}
			}
		} else {
			utils.RespondWithJSON(w, http.StatusForbidden, common.CommonResponse{Success: false, Message: "Invalid file.", Code: http.StatusForbidden, Data: nil})
		}

		return

	}
}

// test handler
func TestHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.OutputHTML(w, "./templates/static/test.html", nil)
		return
	}
}
