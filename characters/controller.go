package characters

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucasmdrs/DungeonAndDevs/models"
)

//CharacterController |
type CharacterController struct {
	service *CharacterService
}

// NewCharacterController |
func NewCharacterController(service *CharacterService) *CharacterController {
	return &CharacterController{service: service}
}

//Create |
func (c *CharacterController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newChar models.Character
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&newChar); err != nil {
		writeJSON(w, http.StatusBadRequest, models.NewHTTPErrorMessage("Malformed request"))
		return
	}

	if valid, msg := c.service.IsValid(newChar); !valid {
		writeJSON(w, http.StatusConflict, models.NewHTTPErrorMessage(msg))
		return
	}
	err := c.service.AddChar(newChar)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusCreated, nil)
}

//Update |
func (c *CharacterController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newChar models.Character
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&newChar); err != nil {
		writeJSON(w, http.StatusBadRequest, models.NewHTTPErrorMessage("Malformed request"))
		return
	}
	if exist := c.service.CharExists(newChar.Name); !exist {
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Character named "+newChar.Name+" was found"))
		return
	}
	err := c.service.UpdateChar(newChar)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusOK, nil)
}

//Find |
func (c *CharacterController) Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	var char models.Character
	char, err := c.service.GetCharByName(name)
	if err != nil {
		log.Panicf("%s\n", err)
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Character named "+name+" was found"))
		return
	}
	writeJSON(w, http.StatusOK, char)
}

//List |
func (c *CharacterController) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var charList []models.Character
	charList, err := c.service.GetAllChars()
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusOK, charList)
}

//Delete |
func (c *CharacterController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if exist := c.service.CharExists(name); !exist {
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Character with the name "+name+" was found"))
		return
	}
	err := c.service.DeleteCharByName(name)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusOK, nil)
}

func writeJSON(response http.ResponseWriter, statusCode int, content interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	if content != nil {
		sb, err := json.Marshal(content)
		if err != nil {
			log.Panicf("%s\n", err)
		}
		_, err = response.Write(sb)
		if err != nil {
			log.Panicf("%s\n", err)
		}
		return
	}
	response.Write([]byte{})
}
