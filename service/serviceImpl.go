package service

import (
	"context"
	"election_system/auth"
	"election_system/entity"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Server      string
	Database    string
	Collection1 string
}

var CollectionUser *mongo.Collection
var ctx = context.TODO()

func (e *Connection) Connect() {
	clientOptions := options.Client().ApplyURI(e.Server)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	CollectionUser = client.Database(e.Database).Collection(e.Collection1)
}

// ===================================userDetails============================================
func (e *Connection) SaveUserDetails(reqBody entity.User) ([]*entity.User, error) {
	var data []*entity.User

	bool, err := validateByNameAndEmail(reqBody)
	if err != nil {
		return data, err
	}
	if !bool {
		return data, errors.New("User already present")
	}

	finalData, err := CollectionUser.InsertOne(ctx, reqBody)
	if err != nil {
		log.Println(err)
		return data, errors.New("Unable to store data")
	}
	result, err := CollectionUser.Find(ctx, bson.D{primitive.E{Key: "_id", Value: finalData.InsertedID}})
	if err != nil {
		log.Println(err)
		return data, err
	}
	data, err = convertDbResultIntoUserStruct(result)
	if err != nil {
		log.Println(err)
		return data, err
	}
	return data, nil
}

// =====================================Verify-user========================================
func (e *Connection) UpdateUserDetails(reqData entity.User, idStr string) (bson.M, error) {
	var updatedDocument bson.M
	id, err := primitive.ObjectIDFromHex(idStr)
	fmt.Println(id)
	if err != nil {
		return updatedDocument, err
	}
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", id}},
			},
		},
	}
	UpdateQuery := bson.D{primitive.E{Key: "is_verified", Value: reqData.IsVerified}}
	/*	if reqData.Name != "" {
			UpdateQuery = append(UpdateQuery, primitive.E{Key: "name", Value: reqData.Name})
		}
		if reqData.Email != "" {
			UpdateQuery = append(UpdateQuery, primitive.E{Key: "email", Value: reqData.Email})
		}
		if reqData.Role != "" {
			UpdateQuery = append(UpdateQuery, primitive.E{Key: "role", Value: reqData.Role})
		}
		if reqData.IsVerified != true {
			UpdateQuery = append(UpdateQuery, primitive.E{Key: "is_verified", Value: reqData.IsVerified})
		}*/

	update := bson.D{{"$set", UpdateQuery}}

	r := CollectionUser.FindOneAndUpdate(ctx, filter, update).Decode(&updatedDocument)
	if r != nil {
		return updatedDocument, r
	}
	fmt.Println(updatedDocument)
	if updatedDocument == nil {
		return updatedDocument, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

// ===================================Update-user=================================
func (e *Connection) UpdateUser(reqData entity.User, email string) (bson.M, error) {
	var updatedDocument bson.M

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"email", email}},
			},
		},
	}
	UpdateQuery := bson.D{}
	if reqData.Role != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "role", Value: reqData.Role})
	}
	if reqData.Name != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "name", Value: reqData.Name})
	}
	if reqData.Email != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "email", Value: reqData.Email})
	}
	if reqData.Password != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "password", Value: reqData.Password})
	}
	if reqData.Phone_number != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "phone_no", Value: reqData.Phone_number})
	}
	if reqData.IsVerified != false {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "is_verified", Value: reqData.IsVerified})
	}
	/*if reqData.V != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "password", Value: reqData.Password})
	}
	if reqData.DOB != "" {
		dob, err := convertDate(reqData.DOB)
		if err != nil {
			log.Println(err)
			return updatedDocument, err
		}
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "dob", Value: dob})
	}*/
	update := bson.D{{"$set", UpdateQuery}}

	r := CollectionUser.FindOneAndUpdate(ctx, filter, update).Decode(&updatedDocument)
	if r != nil {
		return updatedDocument, r
	}
	fmt.Println(updatedDocument)
	if updatedDocument == nil {
		return updatedDocument, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

//==================================search-user==================================

func (e *Connection) UseSearch(reqData entity.UserSearch) (*entity.User, error) {
	var result *entity.User
	/*	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"id", reqData.ID}},
				bson.D{{"role", reqData.Role}},
				bson.D{{"name", reqData.Name}},
				bson.D{{"email", reqData.Email}},
			},
		},
	}*/
	filter := bson.D{primitive.E{Key: "_id", Value: reqData.ID}, primitive.E{Key: "role", Value: reqData.Role},
		primitive.E{Key: "name", Value: reqData.Name}, primitive.E{Key: "email", Value: reqData.Email}}
	err := CollectionUser.FindOne(ctx, filter).Decode(result)

	if err != nil {
		log.Println(err)
		return result, err
	}

	/*finalResult, err := convertDbResultIntoUserStruct(result)
	fmt.Println(finalResult)
	if err != nil {
		log.Println(err)
		return finalResult, err
	}*/
	return result, nil
}

//===================================multiple-user-search============================

func (e *Connection) MultiSearch(reqData entity.MultipleUserSearch) ([]*entity.User, error) {
	var result []*entity.User
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"role", reqData.Role}},
				bson.D{{"city", reqData.City}},
				bson.D{{"state", reqData.State}},
				bson.D{{"country", reqData.Country}},
				bson.D{{"zip_code", reqData.Zip_code}},
				bson.D{{"is_verified", reqData.IsVerified}},
			},
		},
	}

	//var query := {{"role":reqData.Role}, {"city": reqData.City}{"state": reqData.State},{"zip_code": reqData.Zip_code}, {"country": reqData.Country}}
	data, err := CollectionUser.Find(ctx, filter)

	if err != nil {
		log.Println(err)
		return result, err
	}

	result, err = convertDbResultIntoUserStruct(data)
	fmt.Println(result)
	if err != nil {
		log.Println(err)
		return result, err
	}
	return result, nil
}

//==============================================delete-user=======================

func (e *Connection) UserDelete(reqData entity.DeleteUser) (string, error) {

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", reqData.ID}},
				bson.D{{"email", reqData.Email}},
			},
		},
	}

	err := CollectionUser.FindOneAndDelete(ctx, filter)
	if err != nil {
		log.Println(err)
		return "data not found", nil
	}

	return "deleted successfully", nil
}

//====================================deactivate-user============================

func (e *Connection) UserDeactivate(idStr string) (bson.M, error) {
	var updatedDocument bson.M
	id, err := primitive.ObjectIDFromHex(idStr)
	fmt.Println(id)
	if err != nil {
		return updatedDocument, err
	}
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"_id", id}},
			},
		},
	}
	UpdateQuery := bson.D{primitive.E{Key: "is_verified", Value: false}}
	/*	if reqData.Name != "" {
			UpdateQuery = append(UpdateQuery, primitive.E{Key: "name", Value: reqData.Name})
		}
		if reqData.Email != "" {
			UpdateQuery = append(UpdateQuery, primitive.E{Key: "email", Value: reqData.Email})
		}
		if reqData.Role != "" {
			UpdateQuery = append(UpdateQuery, primitive.E{Key: "role", Value: reqData.Role})
		}
		if reqData.IsVerified != true {
			UpdateQuery = append(UpdateQuery, primitive.E{Key: "is_verified", Value: reqData.IsVerified})
		}*/

	update := bson.D{{"$set", UpdateQuery}}

	r := CollectionUser.FindOneAndUpdate(ctx, filter, update).Decode(&updatedDocument)
	if r != nil {
		return updatedDocument, r
	}
	fmt.Println(updatedDocument)
	if updatedDocument == nil {
		return updatedDocument, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

//========================================generate-token=========================

func (e *Connection) GenerateToken(request entity.LoginDetails) (string, error) {

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"email", request.Email}},
				bson.D{{"role", request.Role}},
				bson.D{{"password", request.Password}},
			},
		},
	}

	// check if email exists and password is correct
	record, err := CollectionUser.Find(ctx, filter)
	if err != nil {
		fmt.Println("not found user")
		return "", err
	}

	convertData, err := convertDbResultIntoUserStruct(record)
	if err != nil {
		return "", err
	}

	if len(convertData) != 0 {
		tokenString, err := auth.GenerateJWT(request.Email, request.Password)
		if err != nil {
			return "", err
		}
		return tokenString, err
	} else {
		return "", errors.New("Invalid Credentials")
	}
}

// ===============================================================================
func SetValueInUserModel(req entity.User) (entity.User, error) {
	var data entity.User
	/*dob, err := convertDate(req.PersonalInfo.Dob)
	if err != nil {
		log.Println(err)
		return data, err
	}*/
	data.PersonalInfo.Dob = req.PersonalInfo.Dob
	data.Role = req.Role
	data.Name = req.Name
	data.Email = req.Email
	data.Password = req.Password
	data.Phone_number = req.Phone_number
	data.PersonalInfo.FatherName = req.PersonalInfo.FatherName
	data.PersonalInfo.Age = req.PersonalInfo.Age
	data.PersonalInfo.VoterId = req.PersonalInfo.VoterId
	data.PersonalInfo.DocumentType = req.PersonalInfo.DocumentType
	data.PersonalInfo.Address.Street = req.PersonalInfo.Address.Street
	data.PersonalInfo.Address.City = req.PersonalInfo.Address.City
	data.PersonalInfo.Address.State = req.PersonalInfo.Address.State
	data.PersonalInfo.Address.ZipCode = req.PersonalInfo.Address.ZipCode
	data.PersonalInfo.Address.Country = req.PersonalInfo.Address.Country
	data.IsVerified = req.IsVerified
	return data, nil
}

func convertDate(dateStr string) (time.Time, error) {

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Println(err)
		return date, err
	}
	return date, nil
}

func convertDbResultIntoUserStruct(fetchDataCursor *mongo.Cursor) ([]*entity.User, error) {
	var finalData []*entity.User
	for fetchDataCursor.Next(ctx) {
		var data entity.User
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finalData, err
		}
		finalData = append(finalData, &data)
	}
	return finalData, nil
}

func validateByNameAndEmail(reqbody entity.User) (bool, error) {
	email := reqbody.Email
	var result []*entity.User
	data, err := CollectionUser.Find(ctx, bson.D{{"name", reqbody.Name}, {"email", email}})
	if err != nil {
		return false, err
	}
	result, err = convertDbResultIntoUserStruct(data)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return true, err
	}
	return false, err
}
