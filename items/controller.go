package items

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cassiosantos/DungeonAndDevs/models"
	"github.com/julienschmidt/httprouter"
)

//ItemController |
type ItemController struct {
	service *ItemService
}

// NewItemController |
func NewItemController(service *ItemService) *ItemController {
	return &ItemController{service: service}
}

//Create |
func (c *ItemController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newItem models.Item
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&newItem); err != nil {
		writeJSON(w, http.StatusBadRequest, models.NewHTTPErrorMessage("Malformed request"))
		return
	}

	if valid, msg := c.service.IsValid(newItem); !valid {
		writeJSON(w, http.StatusConflict, models.NewHTTPErrorMessage(msg))
		return
	}
	err := c.service.AddItem(newItem)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusCreated, nil)
}

//Update |
func (c *ItemController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newItem models.Item
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&newItem); err != nil {
		writeJSON(w, http.StatusBadRequest, models.NewHTTPErrorMessage("Malformed request"))
		return
	}
	if exist := c.service.ItemExists(newItem.Name); !exist {
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Item named "+newItem.Name+" was found"))
		return
	}
	err := c.service.UpdateItem(newItem)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusOK, nil)
}

//Find |
func (c *ItemController) Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	var item models.Item
	item, err := c.service.GetItemByName(name)
	if err != nil {
		log.Panicf("%s\n", err)
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Item named "+name+" was found"))
		return
	}
	writeJSON(w, http.StatusOK, item)
}

//List |
func (c *ItemController) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var itemList []models.Item
	itemList, err := c.service.GetAllItems()
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusOK, itemList)
}

//Delete |
func (c *ItemController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if exist := c.service.ItemExists(name); !exist {
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Item with the name "+name+" was found"))
		return
	}
	err := c.service.DeleteItemByName(name)
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
