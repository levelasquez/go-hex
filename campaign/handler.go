package campaign

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler - Actua como adaptador
// Técnicamente no es parte del dominio
type Handler interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	Get(c *gin.Context)
}

type handler struct {
	service Service
}

// NewHandler - init
func NewHandler(service Service) Handler {
	return &handler{
		service,
	}
}

func (h *handler) Create(c *gin.Context) {
	var campaign Campaign
	c.BindJSON(&campaign)

	err := h.service.CreateCampaign(campaign)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.Status(http.StatusCreated)
}

func (h *handler) GetByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.Status(http.StatusBadRequest)
	}

	campaign, err := h.service.FindCampaignByID(id)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, campaign)
}

func (h *handler) Get(c *gin.Context) {
	campaigns, err := h.service.FindAllCampaigns()

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, campaigns)
}
