package controller

import (
	"math"
	api_service "meeting_app/internal/app/service"
	api_structure "meeting_app/internal/app/structures"
	runtime_tools "meeting_app/runtime-tools"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ServiceController struct{ Svc api_service.ServiceService }
type PaginationResponse struct {
	Items      []api_structure.ServiceFilter `json:"items"`
	TotalCount int64                         `json:"totalCount"`
	Page       int                           `json:"page"`
	PageSize   int                           `json:"pageSize"`
	TotalPages int                           `json:"totalPages"`
	MaxPage    int64                         `json:"maxPage"`
	Total      int64                         `json:"total"`
	Last       bool                          `json:"last"`
	First      bool                          `json:"first"`
	Visible    int64                         `json:"visible"`
	Order      string                        `json:"order"`
}

// ShowAccount godoc
// @Summary      Pagination Service
// @Description  Pagination GetServiceWithPagination
// @Tags        	meeting Service
// @Id					 ApiV1MettingAppServicePagination
// @Success      200  {array}  api_structure.Service
// @Failure      400  {object}  error
// @Router       /meeting/Service/pagination [get]
func (controller *ServiceController) GetServiceWithPagination(c *fiber.Ctx) error {
	var Services []api_structure.ServiceFilter
	db := controller.Svc.DB.Table("public.service")

	statusParam := c.Query("status")
	pageParam := c.Query("page")
	page, _ := strconv.Atoi(pageParam)

	pageSize := 10
	orderParam := c.Query("order")
	if orderParam != "asc" && orderParam != "desc" {
		orderParam = "desc"
	}
	if statusParam != "" {
		db = db.Where("status = ?", statusParam)
	}
	db = db.Order("created_at " + orderParam)

	var totalCount int64
	db.Count(&totalCount)
	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))

	if page > totalPages {
		// Sayfa numarası toplam sayfa sayısını aştığında boş bir dizi döndür
		paginationResponse := PaginationResponse{
			Items:      []api_structure.ServiceFilter{},
			TotalCount: totalCount,
			Page:       page,
			PageSize:   pageSize,
			TotalPages: totalPages,
		}
		return c.JSON(paginationResponse)
	} else if page <= 0 {
		// Sayfa numarası belirtilmemişse veya geçersizse, tüm verileri çek
		db.Find(&Services)
	}
	offset := (page - 1) * pageSize
	db.Offset(offset).Limit(pageSize).Find(&Services)

	// Eksik kısımları doldururken aynı zamanda ilişkileri olan verileri kullanın

	// Sayfalama bilgilerini oluşturun
	paginationResponse := PaginationResponse{
		Items:      Services,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		MaxPage:    int64(totalPages - 1),
		Total:      totalCount,
		Last:       page >= totalPages,
		First:      page <= 1,
		Visible:    int64(len(Services)),
		Order:      orderParam,
	}

	// Diziyi JSON olarak cevaplayın
	return c.JSON(paginationResponse)
}

// ShowAccount godoc
// @Summary      Show Service
// @Description  Get GetService
// @Tags        	meeting Service
// @Id					 ApiV1MettingAppServiceSearch
// @Param page query int false "1" default(1)
// @Param size query int false "10" default(10)
// @Param sort query string false "-id,name" default(integrationId)
// @Param filters query string false "[[col,eq,1],[col,eq,2],(missing and rule)[[col,eq,1],and|or,[col,eq,2]]]" default([])
// @Param preload query string false "Service,Service" default(Service,Service)
// @Param joins query string false "Service,Service" default(Service,Service)
// @Param fields query string false "integrationId,name,createdBy,createdAt,updatedAt,deletedAt" default(integrationId,name,createdBy,createdAt,updatedAt,deletedAt)
// @Success      200  {array}  api_structure.Service
// @Failure      400  {object}  error
// @Router       /meeting/Service/search [get]
func (controller *ServiceController) GetService(c *fiber.Ctx) error {
	var Service_table_name api_structure.ServiceFilter

	response, err := controller.Svc.GetService(Service_table_name)
	if err != nil {
		return c.Status(400).JSON(c.Status(400).JSON(c.JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})))
	}
	return c.JSON(response)
}

// ShowAccount godoc
// @Summary      Get Service by Unique IDs
// @Description  Get GetServiceById
// @Tags        	meeting Service
// @Id					 ApiV1MettingAppServiceWithID
// @Param preload query string false "Service,Service" default(Service,Service)
// @Param joins query string false "Service,Service" default(Service,Service)
// @Param fields query string false "integrationId,name,createdBy,createdAt,updatedAt,deletedAt" default(integrationId,name,createdBy,createdAt,updatedAt,deletedAt)
// @Success      200  {array}  api_structure.Service
// @Failure      400  {object}  error
// @Router       /meeting/Service/with-id/:id [get]
func (controller *ServiceController) GetServiceById(c *fiber.Ctx) error {

	id := c.Params("id")
	idTrue, err := runtime_tools.Decrypt(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	result, rerr := controller.Svc.GetServiceById(idTrue)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Update with id in Service
// @Description  Update with id in UpdateService
// @Tags         meeting Service
// @Id					 ApiV1MettingAppServiceUpdateWithID
// @Param request body api_structure.ServiceEdit true "update params"
// @Success      200  {object}  api_structure.Service
// @Failure      400  {object}  error
// @Router       /meeting/Service/with-id/:id [put]
func (controller *ServiceController) UpdateService(c *fiber.Ctx) error {

	id := c.Params("id")

	editData := api_structure.ServiceEdit{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}
	idTrue, err := runtime_tools.Decrypt(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	uerr := controller.Svc.UpdateService(idTrue, editData)
	if uerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Update Data",
			"message": uerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"succes":  true,
		"message": "Updated Successfully",
		"type":    "Update Data",
	})
}

// ShowAccount godoc
// @Summary      Update Multiple in Service
// @Description  Update Multiple in UpdateService
// @Tags         meeting Service
// @Id					 ApiV1MettingAppServicesUpdateMultiple
// @Param request body api_structure.ServiceEdit true "update params"
// @Success      200  {array}  api_structure.Service
// @Failure      400  {object}  error
// @Router       /meeting/Service/multiple [put]
func (controller *ServiceController) UpdateServiceMultiple(c *fiber.Ctx) error {
	editData := []api_structure.ServiceEdit{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateServiceMultiple(editData)
	if uerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Update Data",
			"message": uerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"succes":  true,
		"message": "Updated Successfully",
		"type":    "Update Data",
	})
}

// ShowAccount godoc
// @Summary      Create in Service
// @Description  Create in CreateService
// @Tags         meeting Service
// @Id					 ApiV1MettingAppServiceCreate
// @Param request body api_structure.ServiceFilter true "update params"
// @Success      200  {object}  api_structure.Service
// @Failure      400  {object}  error
// @Router       /meeting/Service/single [post]
func (controller *ServiceController) CreateService(c *fiber.Ctx) error {
	data := api_structure.ServiceForm{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	result, rerr := controller.Svc.CreateService(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Create Multiple in Service
// @Description  Create Multiple in CreateService
// @Tags         meeting Service
// @Id					 ApiV1MettingAppServiceCreateMultiple
// @Param id path string false "id uuid"
// @Param request body api_structure.ServiceFilter true "update params"
// @Success      200  {array}  api_structure.Service
// @Failure      400  {object}  error
// @Router       /meeting/Service/multiple [post]
func (controller *ServiceController) CreateServiceMultiple(c *fiber.Ctx) error {

	bulkData := []api_structure.ServiceForm{}
	if err := c.BodyParser(&bulkData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	result, rerr := controller.Svc.CreateServiceMultiple(bulkData)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Bulk Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Delete with path params in Service
// @Description  Delete with path params in DeleteService
// @Tags         meeting Service
// @Id					 ApiV1MettingAppServiceDelete
// @Accept       json
// @Produce      json
// @Param id path string false "id uuid"
// @Success      200  {string}  string
// @Failure      400  {object}  error
// @Router       /meeting/Service/with-id/:id [delete]
func (controller *ServiceController) DeleteService(c *fiber.Ctx) error {
	id := c.Params("id")

	idTrue, err := runtime_tools.Decrypt(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}

	deleteErr := controller.Svc.DeleteService(idTrue)

	// err := controller.Svc.DeleteService(id)
	if deleteErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Delete Data",
			"message": deleteErr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Deleted Successfully",
		"type":    "Delete Data",
		"success": true,
	})
}
