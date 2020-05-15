package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nfv-aws/wcafe-api-controller/service"
)

// Controller is store controlller
type StoreController struct {
	Service service.StoreService
}

// List action: GET /stores
func (sc StoreController) List(c *gin.Context) {
	p, err := sc.Service.List()

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Create action: POST /stores
func (sc StoreController) Create(c *gin.Context) {
	p, err := sc.Service.Create(c)

	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
	} else {
		c.JSON(201, p)
	}
}

// Get action: GET /stores/:id
func (sc StoreController) Get(c *gin.Context) {
	id := c.Params.ByName("id")
	p, err := sc.Service.Get(id)

	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Update action: PATCH /stores/:id
func (sc StoreController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	p, err := sc.Service.Update(id, c)

	if err != nil {
		c.AbortWithStatus(400)
		log.Println(err)
	} else {
		c.JSON(200, p)
	}
}

// Delete action: DELETE /stores/:id
func (sc StoreController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	p, err := sc.Service.Delete(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.AbortWithStatus(404)
			log.Println(err)
		} else {
			c.AbortWithStatus(409)
			log.Println(err)
		}
	} else {
		c.JSON(204, p)
	}
}
