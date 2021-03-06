package diagnostic_type

import (
    "github.com/google/uuid"
    
    // Import GORM-related packages.
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type DiagnosticType struct {
    Id      uuid.UUID     `json:"id" gorm:"column:id; type:uuid; primary_key;"`
    Name    string        `json:"name" gorm:"column:name; type:string; not_null"` 
}

func (DiagnosticType) TableName() string {
    return "diagnostic_types"
}

func Create(db *gorm.DB, name string) (*DiagnosticType, error) {
    var toInsert = &DiagnosticType{
        Name: name,
    }

    err := db.Create(toInsert).Error;

    if err != nil {
        toInsert = nil
    }
    return toInsert, err
}

func Update(db *gorm.DB, toUpdate *DiagnosticType) (*DiagnosticType, error) {
    err := db.Save(toUpdate).Error;

    if err != nil {
        toUpdate = nil
    }
    return toUpdate, err
}

func FetchById(db *gorm.DB, id uuid.UUID) (*DiagnosticType, error) {
    result :=  &DiagnosticType{}

    err := db.Where("id = ?", id).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchByName(db *gorm.DB, name string) (*DiagnosticType, error) {
    result :=  &DiagnosticType{}

    err := db.Where("name = ?", name).First(&result).Error;

    if err != nil {
        result = nil
    }

    return result, err
}

func FetchList(db *gorm.DB) ([]DiagnosticType, error) {
    var results []DiagnosticType =  nil

    err := db.Find(&results).Error;

    if err != nil {
        results = nil
    }

    return results, err
}