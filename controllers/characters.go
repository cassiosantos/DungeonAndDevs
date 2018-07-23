package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucasmdrs/DungeonAndDevs/db"
	"github.com/lucasmdrs/DungeonAndDevs/models"
	mgo "gopkg.in/mgo.v2"
)

//CharacterController |
type characterController struct {
	DB *mgo.Session
}

// NewCharController creates a new controller
func NewCharController(db *mgo.Session) *characterController {
	return &characterController{
		DB: db,
	}
}

//Create |
func (c *characterController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newChar models.Character
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&newChar); err != nil {
		writeJSON(w, http.StatusNotAcceptable, err)
	}
	err := db.AddChar(c.DB, newChar)
	if err != nil {
		panic(err)
	}
	writeJSON(w, http.StatusCreated, newChar)
}

//Update |
func (c *characterController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newChar models.Character
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&newChar); err != nil {
		writeJSON(w, http.StatusNotAcceptable, err)
	}
	updatedChar, err := db.UpdateChar(c.DB, newChar)
	if err != nil {
		panic(err)
	}
	writeJSON(w, http.StatusOK, updatedChar)
}

//List |
func (c *characterController) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var charList []models.Character
	charList, err := db.GetAllChars(c.DB)
	if err != nil {
		panic(err)
	}
	writeJSON(w, http.StatusOK, charList)
}

//Find |
func (c *characterController) Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	var char models.Character
	char, err := db.GetCharByName(c.DB, name)
	if err != nil {
		panic(err)
	}
	writeJSON(w, http.StatusOK, char)
}

//Delete |
func (c *characterController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	var char models.Character
	err := db.DeleteChar(c.DB, name)
	if err != nil {
		panic(err)
	}
	writeJSON(w, http.StatusOK, char)
}

func writeJSON(response http.ResponseWriter, statusCode int, content interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	sb, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	_, err = response.Write(sb)
	if err != nil {
		panic(err)
	}
}
