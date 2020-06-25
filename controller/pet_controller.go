package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"

	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is pet controlller
type PetController struct {
	Petservice service.PetService
}

// List action: GET /pets
func (pc PetController) List(c *gin.Context) {
	log.Debug().Caller().Msg("pets list")

	p, err := pc.Petservice.List()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// Create action: POST /pets
func (pc PetController) Create(c *gin.Context) {
	log.Debug().Caller().Msg("pets create")

	p, err := pc.Petservice.Create(c)

	if err != nil {
		if strings.Contains(err.Error(), "InvalidAddress") {
			c.AbortWithStatus(http.StatusNotFound)
			log.Error().Caller().Err(err)
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Error().Caller().Err(err)
		}
	} else {
		c.JSON(http.StatusCreated, p)
	}
}

// Get action: GET /pets/:id
func (pc PetController) Get(c *gin.Context) {
	log.Debug().Caller().Msg("pets get")

	id := c.Params.ByName("id")

	p, err := pc.Petservice.Get(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// Update aciton: PATCH /pets/:id
func (pc PetController) Update(c *gin.Context) {
	log.Debug().Caller().Msg("pets update")

	id := c.Params.ByName("id")
	p, err := pc.Petservice.Update(id, c)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatus(http.StatusNotFound)
			log.Error().Caller().Err(err)
		} else {
			c.AbortWithStatus(http.StatusBadRequest)
			log.Error().Caller().Err(err)
		}
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// Delete action: DELETE /pets/:id
func (pc PetController) Delete(c *gin.Context) {
	log.Debug().Caller().Msg("pets delete")

	id := c.Params.ByName("id")
	p, err := pc.Petservice.Delete(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(http.StatusNoContent, p)
	}
}
