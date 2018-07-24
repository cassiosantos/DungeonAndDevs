package attributes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucasmdrs/DungeonAndDevs/models"
)

//AttrController |
type AttrController struct {
	service *AttrService
}

// NewAttrController |
func NewAttrController(service *AttrService) *AttrController {
	return &AttrController{service: service}
}

//Create |
func (c *AttrController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newAttr models.Attribute
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&newAttr); err != nil {
		writeJSON(w, http.StatusBadRequest, models.NewHTTPErrorMessage("Malformed request"))
		return
	}

	if valid, msg := c.service.IsValid(newAttr); !valid {
		writeJSON(w, http.StatusConflict, models.NewHTTPErrorMessage(msg))
		return
	}
	err := c.service.AddAttr(newAttr)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusCreated, nil)
}

//Update |
func (c *AttrController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newAttr models.Attribute
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&newAttr); err != nil {
		writeJSON(w, http.StatusBadRequest, models.NewHTTPErrorMessage("Malformed request"))
		return
	}
	if exist := c.service.AttrExists(newAttr.Name); !exist {
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Attribute named "+newAttr.Name+" was found"))
		return
	}
	err := c.service.UpdateAttr(newAttr)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusOK, nil)
}

//Find |
func (c *AttrController) Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	var attr models.Attribute
	attr, err := c.service.GetAttrByName(name)
	if err != nil {
		log.Panicf("%s\n", err)
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Attribute named "+name+" was found"))
		return
	}
	writeJSON(w, http.StatusOK, attr)
}

//List |
func (c *AttrController) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var attrList []models.Attribute
	attrList, err := c.service.GetAllAttrs()
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusOK, attrList)
}

//Delete |
func (c *AttrController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if exist := c.service.AttrExists(name); !exist {
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Attribute with the name "+name+" was found"))
		return
	}
	err := c.service.DeleteAttrByName(name)
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
