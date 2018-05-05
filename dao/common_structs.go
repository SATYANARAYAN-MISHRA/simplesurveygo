package dao

import (
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Alias    string `json:"alias"`
}

type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
}

func AllMoviesEndPoint(w http.ResponseWriter, r * http.Request) {
    movies, err: = dao.FindAll()
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJson(w, http.StatusOK, movies)
}

type Session struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func AuthenticateUser(cred UserCredentials) Session {
	session := MgoSession.Clone()
	defer session.Close()

	var response interface{}
	clctn := session.DB("simplesurveys").C("user")
	query := clctn.Find(bson.M{"username": cred.Username, "password": cred.Password})
	err := query.One(&response)

	uuidStr := uuid.Must(uuid.NewV4()).String()
	sessionStruct := Session{cred.Username, uuidStr}
	if err != nil {
		return Session{}
	}

	sessionClctn := session.DB("simplesurveys").C("session")
	sessionClctn.Insert(sessionStruct)
	return sessionStruct
}

func GetSessionDetails(token string) UserCredentials {
	session := MgoSession.Clone()
	defer session.Close()

	var response Session
	sessionClctn := session.DB("simplesurveys").C("session")
	query := sessionClctn.Find(bson.M{"token": token})
	err := query.One(&response)
	if err != nil {
		return UserCredentials{}
	}

	var cred UserCredentials
	clctn := session.DB("simplesurveys").C("user")
	query = clctn.Find(bson.M{"username": response.Username})
	err = query.One(&cred)
	return cred
}

func SignupUser(cred UserCredentials) string {
	session := MgoSession.Clone()
	defer session.Close()

	clctn := session.DB("simplesurveys").C("user")

	var signup UserCredentials
	query := clctn.Find(bson.M{"username": cred.Username})
	err := query.One(&signup)
	if signup.Username != "" {
		return "User Already Present"
	}

	err = clctn.Insert(&cred)
	if err != nil {
		panic(err)
	}
	return "User Registered"

}
