package skills

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lucasmdrs/DungeonAndDevs/models"
)

//SkillController |
type SkillController struct {
	service *SkillService
}

// NewSkillController |
func NewSkillController(service *SkillService) *SkillController {
	return &SkillController{service: service}
}

//Create |
func (c *SkillController) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newSkill models.Skill
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&newSkill); err != nil {
		writeJSON(w, http.StatusBadRequest, models.NewHTTPErrorMessage("Malformed request"))
		return
	}

	if valid, msg := c.service.IsValid(newSkill); !valid {
		writeJSON(w, http.StatusConflict, models.NewHTTPErrorMessage(msg))
		return
	}
	err := c.service.AddSkill(newSkill)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusCreated, nil)
}

//Update |
func (c *SkillController) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var newSkill models.Skill
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&newSkill); err != nil {
		writeJSON(w, http.StatusBadRequest, models.NewHTTPErrorMessage("Malformed request"))
		return
	}
	if exist := c.service.SkillExists(newSkill.Name); !exist {
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Skill named "+newSkill.Name+" was found"))
		return
	}
	err := c.service.UpdateSkill(newSkill)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusOK, nil)
}

//Find |
func (c *SkillController) Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	var skill models.Skill
	skill, err := c.service.GetSkillByName(name)
	if err != nil {
		log.Panicf("%s\n", err)
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Skill named "+name+" was found"))
		return
	}
	writeJSON(w, http.StatusOK, skill)
}

//List |
func (c *SkillController) List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var skillList []models.Skill
	skillList, err := c.service.GetAllSkills()
	if err != nil {
		log.Panicf("%s\n", err)
	}
	writeJSON(w, http.StatusOK, skillList)
}

//Delete |
func (c *SkillController) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	if exist := c.service.SkillExists(name); !exist {
		writeJSON(w, http.StatusPreconditionFailed, models.NewHTTPErrorMessage("No Skill with the name "+name+" was found"))
		return
	}
	err := c.service.DeleteSkillByName(name)
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
