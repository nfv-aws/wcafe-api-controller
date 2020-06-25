package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"

	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is store controlller
type StoreController struct {
	Storeservice service.StoreService
}

// List action: GET /stores
func (sc StoreController) List(c *gin.Context) {
	log.Debug().Caller().Msg("stores list")

	p, err := sc.Storeservice.List()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// Create action: POST /stores
func (sc StoreController) Create(c *gin.Context) {
	log.Debug().Caller().Msg("stores create")
	p, err := sc.Storeservice.Create(c)

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

// Get action: GET /stores/:id
func (sc StoreController) Get(c *gin.Context) {
	log.Debug().Caller().Msg("stores get")
	id := c.Params.ByName("id")
	p, err := sc.Storeservice.Get(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

// Update action: PATCH /stores/:id
func (sc StoreController) Update(c *gin.Context) {
	log.Debug().Caller().Msg("stores update")
	id := c.Params.ByName("id")
	p, err := sc.Storeservice.Update(id, c)

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

// Delete action: DELETE /stores/:id
func (sc StoreController) Delete(c *gin.Context) {
	log.Debug().Caller().Msg("stores delete")
	id := c.Params.ByName("id")
	p, err := sc.Storeservice.Delete(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatus(http.StatusNotFound)
			log.Error().Caller().Err(err)
		} else {
			c.AbortWithStatus(http.StatusConflict)
			log.Error().Caller().Err(err)
		}
	} else {
		c.JSON(http.StatusNoContent, p)
	}
}

// List action: GET /stores/:id/pets
func (sc StoreController) PetsList(c *gin.Context) {
	log.Debug().Caller().Msg("stores petslist")
	id := c.Params.ByName("id")
	p, err := sc.Storeservice.PetsList(id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Error().Caller().Err(err)
	} else {
		c.JSON(http.StatusOK, p)
	}
}
