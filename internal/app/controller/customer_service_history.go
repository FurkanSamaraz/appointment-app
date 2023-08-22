package controller

import (
	"math"
	api_service "meeting_app/internal/app/service"
	api_structure "meeting_app/internal/app/structures"
	runtime_tools "meeting_app/runtime-tools"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CustomerServiceHistoryController struct {
	Svc api_service.CustomerServiceHistoryService
}
type PaginationCustomerServiceHistory struct {
	Items      []api_structure.CustomerServiceHistoryFilter `json:"items"`
	TotalCount int64                                        `json:"totalCount"`
	Page       int                                          `json:"page"`
	PageSize   int                                          `json:"pageSize"`
	TotalPages int                                          `json:"totalPages"`
	MaxPage    int64                                        `json:"maxPage"`
	Total      int64                                        `json:"total"`
	Last       bool                                         `json:"last"`
	First      bool                                         `json:"first"`
	Visible    int64                                        `json:"visible"`
	Order      string                                       `json:"order"`
}

// ShowAccount godoc
// @Summary      Pagination CustomerServiceHistory
// @Description  Pagination GetCustomerServiceHistoryWithPagination
// @Tags        	meeting CustomerServiceHistory
// @Id					 ApiV1MeetingAppCustomerServiceHistoryPagination
// @Success      200  {array}  api_structure.CustomerServiceHistory
// @Failure      400  {object}  error
// @Router       /meeting/CustomerServiceHistory/pagination [get]
func (controller *CustomerServiceHistoryController) GetCustomerServiceHistoryWithPagination(c *fiber.Ctx) error {
	var CustomerServiceHistorys []api_structure.CustomerServiceHistoryFilter
	db := controller.Svc.DB.Table("public.customer_service_history")

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
		paginationResponse := PaginationCustomerServiceHistory{
			Items:      []api_structure.CustomerServiceHistoryFilter{},
			TotalCount: totalCount,
			Page:       page,
			PageSize:   pageSize,
			TotalPages: totalPages,
		}
		return c.JSON(paginationResponse)
	} else if page <= 0 {
		// Sayfa numarası belirtilmemişse veya geçersizse, tüm verileri çek
		db.Find(&CustomerServiceHistorys)
	}
	offset := (page - 1) * pageSize
	db.Offset(offset).Limit(pageSize).Find(&CustomerServiceHistorys)

	// Eksik kısımları doldururken aynı zamanda ilişkileri olan verileri kullanın

	// Sayfalama bilgilerini oluşturun
	paginationResponse := PaginationCustomerServiceHistory{
		Items:      CustomerServiceHistorys,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		MaxPage:    int64(totalPages - 1),
		Total:      totalCount,
		Last:       page >= totalPages,
		First:      page <= 1,
		Visible:    int64(len(CustomerServiceHistorys)),
		Order:      orderParam,
	}

	// Diziyi JSON olarak cevaplayın
	return c.JSON(paginationResponse)
}

// ShowAccount godoc
// @Summary      Show CustomerServiceHistory
// @Description  Get GetCustomerServiceHistory
// @Tags        	meeting CustomerServiceHistory
// @Id					 ApiV1MeetingAppCustomerServiceHistorySearch
// @Success      200  {array}  api_structure.CustomerServiceHistory
// @Failure      400  {object}  error
// @Router       /meeting/CustomerServiceHistory/search [get]
func (controller *CustomerServiceHistoryController) GetCustomerServiceHistory(c *fiber.Ctx) error {
	var CustomerServiceHistory_table_name api_structure.CustomerServiceHistoryFilter

	response, err := controller.Svc.GetCustomerServiceHistory(CustomerServiceHistory_table_name)
	if err != nil {
		return c.Status(400).JSON(c.Status(400).JSON(c.JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})))
	}
	return c.JSON(response)
}

// ShowAccount godoc
// @Summary      Get CustomerServiceHistory by Unique IDs
// @Description  Get GetCustomerServiceHistoryById
// @Tags        	meeting CustomerServiceHistory
// @Id					 ApiV1MeetingAppCustomerServiceHistoryWithID
// @Success      200  {array}  api_structure.CustomerServiceHistory
// @Failure      400  {object}  error
// @Router       /meeting/CustomerServiceHistory/with-id/:id [get]
func (controller *CustomerServiceHistoryController) GetCustomerServiceHistoryById(c *fiber.Ctx) error {

	id := c.Params("id")
	idTrue, err := runtime_tools.Decrypt(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	result, rerr := controller.Svc.GetCustomerServiceHistoryById(idTrue)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Update with id in CustomerServiceHistory
// @Description  Update with id in UpdateCustomerServiceHistory
// @Tags         meeting CustomerServiceHistory
// @Id					 ApiV1MeetingAppCustomerServiceHistoryUpdateWithID
// @Param id path string false "id uuid"
// @Param request body api_structure.CustomerServiceHistoryEdit true "update params"
// @Success      200  {object}  api_structure.CustomerServiceHistory
// @Failure      400  {object}  error
// @Router       /meeting/CustomerServiceHistory/with-id/:id [put]
func (controller *CustomerServiceHistoryController) UpdateCustomerServiceHistory(c *fiber.Ctx) error {

	id := c.Params("id")

	editData := api_structure.CustomerServiceHistoryEdit{}
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
	uerr := controller.Svc.UpdateCustomerServiceHistory(idTrue, editData)
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
// @Summary      Update Multiple in CustomerServiceHistory
// @Description  Update Multiple in UpdateCustomerServiceHistory
// @Tags         meeting CustomerServiceHistory
// @Id					 ApiV1MeetingAppCustomerServiceHistorysUpdateMultiple
// @Param id path string false "id uuid"
// @Param request body api_structure.CustomerServiceHistoryEdit true "update params"
// @Success      200  {array}  api_structure.CustomerServiceHistory
// @Failure      400  {object}  error
// @Router       /meeting/CustomerServiceHistory/multiple [put]
func (controller *CustomerServiceHistoryController) UpdateCustomerServiceHistoryMultiple(c *fiber.Ctx) error {
	editData := []api_structure.CustomerServiceHistoryEdit{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateCustomerServiceHistoryMultiple(editData)
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
// @Summary      Create in CustomerServiceHistory
// @Description  Create in CreateCustomerServiceHistory
// @Tags         meeting CustomerServiceHistory
// @Id					 ApiV1MeetingAppCustomerServiceHistoryCreate
// @Param id path string false "id uuid"
// @Param request body api_structure.CustomerServiceHistoryFilter true "update params"
// @Success      200  {object}  api_structure.CustomerServiceHistory
// @Failure      400  {object}  error
// @Router       /meeting/CustomerServiceHistory/single [post]
func (controller *CustomerServiceHistoryController) CreateCustomerServiceHistory(c *fiber.Ctx) error {
	data := api_structure.CustomerServiceHistoryForm{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	result, rerr := controller.Svc.CreateCustomerServiceHistory(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Create Multiple in CustomerServiceHistory
// @Description  Create Multiple in CreateCustomerServiceHistory
// @Tags         meeting CustomerServiceHistory
// @Id					 ApiV1MeetingAppCustomerServiceHistoryCreateMultiple
// @Param   Authorization  header     string     true  "Current Session Token" 	default([[authToken]])
// @Param   x-company  		 header     string     true  "Current CompanyID" 			default([[companyID]])
// @Param id path string false "id uuid"
// @Param request body api_structure.CustomerServiceHistoryFilter true "update params"
// @Success      200  {array}  api_structure.CustomerServiceHistory
// @Failure      400  {object}  error
// @Router       /meeting/CustomerServiceHistory/multiple [post]
func (controller *CustomerServiceHistoryController) CreateCustomerServiceHistoryMultiple(c *fiber.Ctx) error {

	bulkData := []api_structure.CustomerServiceHistoryForm{}
	if err := c.BodyParser(&bulkData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	result, rerr := controller.Svc.CreateCustomerServiceHistoryMultiple(bulkData)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Bulk Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Delete with path params in CustomerServiceHistory
// @Description  Delete with path params in DeleteCustomerServiceHistory
// @Tags         meeting CustomerServiceHistory
// @Id					 ApiV1MeetingAppCustomerServiceHistoryDelete
// @Accept       json
// @Produce      json
// @Param id path string false "id uuid"
// @Success      200  {string}  string
// @Failure      400  {object}  error
// @Router       /meeting/CustomerServiceHistory/with-id/:id [delete]
func (controller *CustomerServiceHistoryController) DeleteCustomerServiceHistory(c *fiber.Ctx) error {
	id := c.Params("id")
	idTrue, err := runtime_tools.Decrypt(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	deleteErr := controller.Svc.DeleteCustomerServiceHistory(idTrue)
	// err := controller.Svc.DeleteCustomerServiceHistory(id)
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
