package brankas

import db "Brankas/base/db/postgres"

type ImageDetails struct {
	TableName   struct{} `sql:"all_images" json:"-"`
	Id          string   `param:"id" json:"id" sql:"id"`
	FileName    string   `param:"fileName" json:"fileName" sql:"file_name"`
	ContentType string   `param:"contentType" json:"contentType" sql:"content_type"`
	Size        int64    `param:"size" json:"size" sql:"size"`
}

// create business user
func CreateImageDetailsRow(data ImageDetails) (ImageDetails, error) {
	err := db.PG.Create(&data)
	return data, err
}
