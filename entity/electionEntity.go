package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Role         string             `bson:"role,omitempty" json:"role,omitempty"`
	Name         string             `bson:"name,omitempty" json:"name,omitempty"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Password     string             `bson:"password,omitempty" json:"password,omitempty"`
	Phone_number string             `bson:"phone_no,omitempty" json:"phone_no,omitempty"`
	IsVerified   bool               `bson:"is_verified,omitempty" json:"is_verified,omitempty"`
	PersonalInfo PersonalInfo       `bson:"personal_info,omitempty" json:"personal_info,omitempty"`
	VerifiedBy   VerifiedBy         `bson:"verified_by,omitempty" json:"verified_by,omitempty"`
	UploadedDocs UploadedDocs       `bson:"uploaded_docs,omitempty" json:"uploaded_docs,omitempty"`
}

type PersonalInfo struct {
	FatherName   string    `bson:"father_name,omitempty" json:"father_name,omitempty"`
	Dob          time.Time `bson:"dob,omitempty" json:"dob,omitempty"`
	Age          int       `bson:"age,omitempty" json:"age,omitempty"`
	VoterId      string    `bson:"voter_id,omitempty" json:"voter_id,omitempty"`
	DocumentType string    `bson:"document_type,omitempty" json:"document_type,omitempty"`
	Address      Address   `bson:"address,omitempty" json:"address,omitempty"`
}

type Address struct {
	Street  string `bson:"street,omitempty" json:"street,omitempty"`
	City    string `bson:"city,omitempty" json:"city,omitempty"`
	State   string `bson:"state,omitempty" json:"state,omitempty"`
	ZipCode string `bson:"zip_code,omitempty" json:"zip_code,omitempty"`
	Country string `bson:"country,omitempty" json:"country,omitempty"`
}

type VerifiedBy struct {
	UserID   string `bson:"user_id,omitempty" json:"user_id,omitempty"`
	UserName string `bson:"user_name,omitempty" json:"user_name,omitempty"`
}

type UploadedDocs struct {
	DocumentType             string `bson:"document_type,omitempty" json:"document_type,omitempty"`
	DocumentIdentificationNo string `bson:"document_identification_no,omitempty" json:"document_identification_no,omitempty"`
	DocumentPath             string `bson:"document_path,omitempty" json:"document_path,omitempty"`
}

type Response struct {
	Success    string      `json:"success,omitempty"`
	SucessCode string      `json:"successCode,omitempty"`
	Response   interface{} `json:"response,omitempty"`
}

type VerifyRequest struct {
	Role       string `bson:"role,omitempty" json:"role,omitempty"`
	Name       string `bson:"name,omitempty" json:"name,omitempty"`
	Email      string `bson:"email,omitempty" json:"email,omitempty"`
	IsVerified bool   `bson:"is_verified,omitempty" json:"is_verified,omitempty"`
}

type UserSearch struct {
	ID    string `bson:"id,omitempty" json:"id,omitempty"`
	Role  string `bson:"role,omitempty" json:"role,omitempty"`
	Name  string `bson:"name,omitempty" json:"name,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
}

type MultipleUserSearch struct {
	Role       string `bson:"role,omitempty" json:"role,omitempty"`
	City       string `bson:"city,omitempty" json:"city,omitempty"`
	State      string `bson:"state,omitempty" json:"state,omitempty"`
	Zip_code   string `bson:"zip_code,omitempty" json:"zip_code,omitempty"`
	Country    string `bson:"country,omitempty" json:"country,omitempty"`
	IsVerified bool   `bson:"is_verified,omitempty" json:"is_verified,omitempty"`
}

type DeleteUser struct {
	ID    string `bson:"id,omitempty" json:"id,omitempty"`
	Email string `bson:"email,omitempty" json:"email,omitempty"`
}

type LoginDetails struct {
	Role     string `bson:"role,omitempty" json:"role,omitempty"`
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
	Name     string `bson:"name,omitempty" json:"name,omitempty"`
}
