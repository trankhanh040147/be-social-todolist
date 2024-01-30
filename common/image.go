package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url"`
	Width     int    `json:"width" gorm:"column:width"`
	Height    int    `json:"height" gorm:"column:height"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string { return "images" }

func (j *Image) FullFill(domain string) {
	j.Url = fmt.Sprintf("%s/%s", domain, j.Url)
}

// Scan scans the JSONB value and unmarshals it into the Image object.
// It takes a value of type interface{} and returns an error if the unmarshaling fails.
// The value should be a byte slice ([]byte) representing the JSONB value.
// If the unmarshaling is successful, the Image object is updated with the unmarshaled data.
func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json image, implement driver.Valuer interface
// Value returns the value of the Image as a driver.Value.
// If the Image is nil, it returns nil, nil.
// Otherwise, it marshals the Image into JSON format and returns the result.
func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
