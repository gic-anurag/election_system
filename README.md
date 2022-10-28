# election_system

**User Table crud api’s**

    Description -> this api is used to add new user in the collection api
    -> [localhost:8080/add-user](http://localhost:8080/user/add-user)

input -> {
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": ISODate("2016-04-08T15:06:21.595Z"),
        "age": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    }
}

output -> { “successCode” : true,
                    “successMsg”: “User details saved successfully!”, 
                    “response”: {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": ISODate("2016-04-08T15:06:21.595Z"),
        "age": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "is_verified":false,
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    }
}

                  }

    > Description -> this api is used to verify voter in the collection
    > using the documents uploaded by voter, can be accessed after login
    > only by ‘admin’ type user api -> [localhost:8080/verify-user](http://localhost:8080/user/verifyUser/)

input -> {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "is_verified": true
}
output -> { “successCode” : true,
                    “successMsg”: “Voter verified successfully!”, 
                    “response”: {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": ISODate("2016-04-08T15:06:21.595Z"),
        "age": "Lorem",
        "voter_id": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "is_verified": true,
    "verified_by": {
        "_id": ObjectId("6c6c2d84de0f09d5007bdef8"),
        "name": "Lorem"
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    }
}
}


    Description -> this api is used to update user in the collection, can be accessed after login only by ‘admin’ type user
    api -> [localhost:8080/update-user](http://localhost:8080/user/update-User)

input -> {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": ISODate("2016-04-08T15:06:21.595Z"),
        "age": "Lorem",
        "voter_id": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "is_verified": true,
    "verified_by": {
        "_id": ObjectId("6c6c2d84de0f09d5007bdef8"),
        "name": "Lorem"
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    }
    "voted": [
        ObjectId("ccebdc4611f903debe1b15ac")
    ]

}

output -> { “successCode” : true,
                    “successMsg”: “User details updated successfully!”, 
                    “response”: {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": ISODate("2016-04-08T15:06:21.595Z"),
        "age": "Lorem",
        "voter_id": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "is_verified": true,
    "verified_by": {
        "_id": ObjectId("6c6c2d84de0f09d5007bdef8"),
        "name": "Lorem"
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    },
    "voted": [
        ObjectId("ccebdc4611f903debe1b15ac")
    ]
}

                  }

    Description -> this api is used to search user in the collection
    api -> [localhost:8080/search-one-user](http://localhost:8080/user/searchUser)

input -> {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
}
output -> { “successCode” : true,
                    “successMsg”: “User details searched successfully!”, 
                    “response”: {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": ISODate("2016-04-08T15:06:21.595Z"),
        "age": "Lorem",
        "voter_id": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "is_verified": true,
    "verified_by": {
        "_id": ObjectId("6c6c2d84de0f09d5007bdef8"),
        "name": "Lorem"
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    },
    "voted": [
        ObjectId("ccebdc4611f903debe1b15ac")
    ]
}

                  }

    Description -> this api is used to search list of users in the collection based on certain criteria, any one or all keys can be used in input json for searching
    api -> [localhost:8080/search-multiple-user](http://localhost:8080/user/search-multiple-user)

input -> {
"role": "Lorem",
"city": "Lorem",
"state": "Lorem",
"zip_code": "Lorem",
"country": "Lorem",
"is_verified": true,
"voted": [
        ObjectId("ccebdc4611f903debe1b15ac")
    ]

}
output -> { “successCode” : true,
                    “successMsg”: “User details searched successfully!”, 
                    “response”: [{
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": ISODate("2016-04-08T15:06:21.595Z"),
        "age": "Lorem",
        "voter_id": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "is_verified": true,
    "verified_by": {
        "_id": ObjectId("6c6c2d84de0f09d5007bdef8"),
        "name": "Lorem"
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    },
    "voted": [
        ObjectId("ccebdc4611f903debe1b15ac")
    ]
}]
 }

    Description -> this api is used to delete user in the collection
    api -> [localhost:8080/delete-user](http://localhost:8080/delete-user)

input -> {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
}

output -> { “successCode” : true,
                    “successMsg”: “User details deleted successfully!”, 
                    “response”: {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": ISODate("2016-04-08T15:06:21.595Z"),
        "age": "Lorem",
        "voter_id": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "is_verified": true,
    "verified_by": {
        "_id": ObjectId("6c6c2d84de0f09d5007bdef8"),
        "name": "Lorem"
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    },
    "voted": [
        ObjectId("ccebdc4611f903debe1b15ac")
    ]
}

                  }

    Description -> this api is used to deactivate user in the collection, mark is_verified as false
    api -> [localhost:8080/deactivate-user](http://localhost:8080/user/deactivate-user/)

input -> {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
“is_verified”:false
}

output -> { “successCode” : true,
                    “successMsg”: “User details deactivate successfully!”, 
                    “response”: {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "role": "Lorem",
    "name": "Lorem",
    "email": "Lorem",
    "password": "Lorem",
    "phone_number": "Lorem",
    "personal_info": {
        "name": "Lorem",
        "father_name": "Lorem",
        "dob": ISODate("2016-04-08T15:06:21.595Z"),
        "age": "Lorem",
        "voter_id": "Lorem",
        "document_type": "Lorem",
        "address": {
            "street": "Lorem",
            "city": "Lorem",
            "state": "Lorem",
            "zip_code": "Lorem",
            "country": "Lorem"
        }
    },
    "is_verified":false,
    "verified_by": {
        "_id": ObjectId("6c6c2d84de0f09d5007bdef8"),
        "name": "Lorem"
    },
    "uploaded_docs": {
        "document_type": "Lorem",
        "document_identification_no": "Lorem",
        "document_path": "Lorem"
    },
    "voted": [
        ObjectId("ccebdc4611f903debe1b15ac")
    ]
}

                  }
================================================================================================

**Election Details**

    Description -> this api is used to add new election
    API:- http://localhost:8080/election

Method POST
InputJson:-
{
"election_date": " ",
"result_date": " ",
"election_staus": " ",
"result": " ",
"location":" "
}

output -> { “successCode” : true,
                    “successMsg”: “Election details saved successfully!”,
                    “response”: {
 "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
"election_date": " ",
"result_date": " ",
"election_staus": " ",
"result": " ",
"location":" "
}
}


    Description -> this api is used to add new candidate
    API:- http://localhost:8080/api/candidate

Method POST
InputJson:-

"candidate":[{
"id":"",
"name":" ",
"vote_count":" ",
"vote_sign":" ",
"nomination_status": " "
"nomination_verified_by": " "
}]
}


output -> { “successCode” : true,
                    “successMsg”: “Candidate details saved successfully!”,
                    “response”: {
 "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
"candidate":[{
"id":"",
"name":" ",
"vote_count":" ",
"is_nomination_verified": " "
"nomination_verified_by": " "
"vote_sign":" ",
}]
}

    Description -> this api is used to verify candidate only by admin
    API:- http://localhost:8080/verify-candidate(only admin)
    Method POST

InputJson:-

{"nomination_status": " ",

"id":" ",

}

output -> { “successCode” : true,
                    “successMsg”: “Candidate details saved successfully!”,
                    “response”: {
 "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
"candidate":[{
"id":"",
"name":" ",
"vote_count":" ",
"election_id": " ",
"nomination_status": " ",
"nomination_verified_by": " ",
"vote_sign":" ",
}]
}


    Description -> this api is used to update election only by admin
    API:- http://localhost:8080/update_election{id}
    Method PUT

InputJson:-
{
"election_date": " ",
"result_date": " ",
"election_staus": " ",
"location":" "
}

output -> { “successCode” : true,
                    “successMsg”: “Election details update successfully!”,
                    “response”: {
 "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
"election_date": " ",
"result_date": " ",
"election_staus": " ",
"result": " ",
"location":" "
}
}


    Description -> this api is used to search election in the collection
    API:- http://localhost:8080/election_search
    Method POST

InputJson:- {
    "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
    "election_date": " ",
    "result_date": " ",
    "election_staus": " ",
    "candidate_name": " ",
    "location":" "
}

output -> { “successCode” : true,
                    “successMsg”: “Election details search successfully!”,
                    “response”: {
 "_id": ObjectId("0e1efc01f9961f8af5cee4f6"),
"election_date": " ",
"result_date": " ",
"election_staus": " ",
"result": " ",
"location":" "
}
}


    Description -> this api is used to deactivate election in the collection only admin
    API:- http://localhost:8080/deactivate_election/{id}
    Method DELETE

Output-
record deleted successfully

    Description -> this api is used to find election in the collection
    API:- http://localhost:8080/find_election/{id}
    Method GET

Output-
{}
==============================================================================================

    Description:- Caste vote
    API:- http://localhost:8080/cast-vote

Input:-{
"election_id": "",
"candidate_id": ""
}
Output:-{Candidates profile}


    Description:- Election result search
    API:- http://localhost:8080/election-result

Input:- {
"election_id": "",
"location": "",
"election_date": "",
"result_date": ""
"election_status": ""
}
Output:-{Candidates profile}


    Description:-Search Candidates profile by election ID, Location, election dates and election status
    API:- http://localhost:8080/candidates-profile

Input:- {
"election_id": "",
"location": "",
"election_date": "",
"election_status": ""
}
Output:- {Each candidate profile for this election}

===================================================
**Cast, Vote and candidates**

    Description:- Caste vote
    API:- http://localhost:8080/cast-vote

Input:-{
"election_id": "",
"candidate_id": ""
}
Output:-{Candidates profile}


    Description:- Election result search
    API:- http://localhost:8080/election-result

Input:- {
"election_id": "",
"location": "",
"election_date": "",
"result_date": ""
"election_status": ""
}
Output:-{Candidates profile}


    Description:-Search Candidates profile by election ID, Location, election dates and election status
    API:- http://localhost:8080/candidates-profile

Input:- {
"election_id": "",
"location": "",
"election_date": "",
"election_status": ""
}
Output:- {Each candidate profile for this election}




