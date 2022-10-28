package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveUserDetails(t *testing.T) {

	var jsonStr = []byte(`{
		"role": "arsh",
		"name": "arsh",
		"email": "arsh@gm",
		"password": "arsh",
		"phone_no": "995524",
		"is_verified": false,
		"personal_info":{
		  "father_name": "arshk",
		  "age": 36,
		  "voter_id": "",
		  "document_type": "aadharcard",
		  "address": {
			 "street": "2",
			 "city": "patna",
			 "state": "Bihar",
			 "zip_code": "800026",
			 "country": "India"
			}
		},
			"verified_by": {
				"user_id":"",
				"user_name": ""
			},
			"uploaded_docs": {
				"document_type": "aadharcard",
				"document_identification_no": "82802665243",
				"document_path": "E/project"
			}
		}`)

	req, err := http.NewRequest("POST", "/user/save-user", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(saveUserDetails)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{
		"success": "true",
		"successCode": "202",
		"response": [
			{
				"id": "635ba2ed558430306f32b287",
				"role": "admin",
				"name": "arsh",
				"email": "arsh@gm",
				"password": "arsh",
				"phone_no": "995524",
				"is_verified": false,
				"personal_info": {
					"father_name": "arsh",
					"dob": "0001-01-01T00:00:00Z",
					"age": 36,
					"document_type": "aadharcard",
					"address": {
						"street": "2",
						"city": "patna",
						"state": "Bihar",
						"zip_code": "800026",
						"country": "India"
					}
				},
				"verified_by": {},
				"uploaded_docs": {
					"document_type": "aadharcard",
					"document_identification_no": "82802665243",
					"document_path": "E/project"
				}
			}
		]
	}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
