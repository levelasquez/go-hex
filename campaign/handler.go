package campaign

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CampaignHandler - Acts as adapter
// Technically not part of the domain
// It could be part of a handler package
type CampaignHandler interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	Get(c *gin.Context)
}

type campaignHandler struct {
	campaignService CampaignService
}

// NewCampaignHandler - initialize the handler
func NewCampaignHandler(campaignService CampaignService) CampaignHandler {
	return &campaignHandler{
		campaignService,
	}
}

func (h *campaignHandler) Create(c *gin.Context) {
	var campaign Campaign
	c.BindJSON(&campaign)

	err := h.campaignService.CreateCampaign(&campaign)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.Status(http.StatusCreated)
}

func (h *campaignHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.Status(http.StatusBadRequest)
	}

	campaign, err := h.campaignService.FindCampaignByID(id)

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, campaign)
}

func (h *campaignHandler) Get(c *gin.Context) {
	campaigns, err := h.campaignService.FindAllCampaigns()

	if err != nil {
		c.Status(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, campaigns)
}
