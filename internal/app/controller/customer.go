package controller

import (
	"math"
	api_service "meeting_app/internal/app/service"
	api_structure "meeting_app/internal/app/structures"
	runtime_tools "meeting_app/runtime-tools"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CustomerController struct{ Svc api_service.CustomerService }
type PaginationCustomer struct {
	Items      []api_structure.CustomerFilter `json:"items"`
	TotalCount int64                          `json:"totalCount"`
	Page       int                            `json:"page"`
	PageSize   int                            `json:"pageSize"`
	TotalPages int                            `json:"totalPages"`
	MaxPage    int64                          `json:"maxPage"`
	Total      int64                          `json:"total"`
	Last       bool                           `json:"last"`
	First      bool                           `json:"first"`
	Visible    int64                          `json:"visible"`
	Order      string                         `json:"order"`
}

// ShowAccount godoc
// @Summary      Pagination Customer
// @Description  Pagination GetCustomerWithPagination
// @Tags        	meeting Customer
// @Id					 ApiV1MeetingAppCustomerPagination
// @Success      200  {array}  api_structure.Customer
// @Failure      400  {object}  error
// @Router       /api/v1/Customer/pagination [get]
func (controller *CustomerController) GetCustomerWithPagination(c *fiber.Ctx) error {
	var customers []api_structure.CustomerFilter
	db := controller.Svc.DB.Table("public.customer")

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
		paginationResponse := PaginationCustomer{
			Items:      []api_structure.CustomerFilter{},
			TotalCount: totalCount,
			Page:       page,
			PageSize:   pageSize,
			TotalPages: totalPages,
		}
		return c.JSON(paginationResponse)
	} else if page <= 0 {
		// Sayfa numarası belirtilmemişse veya geçersizse, tüm verileri çek
		db.Find(&customers)
	}
	offset := (page - 1) * pageSize
	db.Offset(offset).Limit(pageSize).Find(&customers)

	// Eksik kısımları doldururken aynı zamanda ilişkileri olan verileri kullanın

	// Sayfalama bilgilerini oluşturun
	paginationResponse := PaginationCustomer{
		Items:      customers,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		MaxPage:    int64(totalPages - 1),
		Total:      totalCount,
		Last:       page >= totalPages,
		First:      page <= 1,
		Visible:    int64(len(customers)),
		Order:      orderParam,
	}

	// Diziyi JSON olarak cevaplayın
	return c.JSON(paginationResponse)
}

// ShowAccount godoc
// @Summary      Show Customer
// @Description  Get GetCustomer
// @Tags        	meeting Customer
// @Id					 ApiV1MeetingAppCustomerSearch
// @Success      200  {array}  api_structure.Customer
// @Failure      400  {object}  error
// @Router       /api/v1/Customer/search [get]
func (controller *CustomerController) GetCustomer(c *fiber.Ctx) error {
	var Customer_table_name api_structure.CustomerFilter

	response, err := controller.Svc.GetCustomer(Customer_table_name)
	if err != nil {
		return c.Status(400).JSON(c.Status(400).JSON(c.JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})))
	}
	return c.JSON(response)
}

// ShowAccount godoc
// @Summary      Get Customer by Unique IDs
// @Description  Get GetCustomerById
// @Tags        	meeting Customer
// @Id					 ApiV1MeetingAppCustomerWithID
// @Success      200  {array}  api_structure.Customer
// @Failure      400  {object}  error
// @Router       /api/v1/Customer/with-id/:id [get]
func (controller *CustomerController) GetCustomerById(c *fiber.Ctx) error {

	id := c.Params("id")
	idTrue, err := runtime_tools.Decrypt(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	result, rerr := controller.Svc.GetCustomerById(idTrue)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Update with id in Customer
// @Description  Update with id in UpdateCustomer
// @Tags         meeting Customer
// @Id					 ApiV1MeetingAppCustomerUpdateWithID
// @Param id path string false "id uuid"
// @Param request body api_structure.CustomerEdit true "update params"
// @Success      200  {object}  api_structure.Customer
// @Failure      400  {object}  error
// @Router       /api/v1/Customer/with-id/:id [put]
func (controller *CustomerController) UpdateCustomer(c *fiber.Ctx) error {

	id := c.Params("id")

	editData := api_structure.CustomerEdit{}
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
	uerr := controller.Svc.UpdateCustomer(idTrue, editData)
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
// @Summary      Update Multiple in Customer
// @Description  Update Multiple in UpdateCustomer
// @Tags         meeting Customer
// @Id					 ApiV1MeetingAppCustomersUpdateMultiple
// @Param id path string false "id uuid"
// @Param request body api_structure.CustomerEdit true "update params"
// @Success      200  {array}  api_structure.Customer
// @Failure      400  {object}  error
// @Router       /api/v1/Customer/multiple [put]
func (controller *CustomerController) UpdateCustomerMultiple(c *fiber.Ctx) error {
	editData := []api_structure.CustomerEdit{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateCustomerMultiple(editData)
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
// @Summary      Create in Customer
// @Description  Create in CreateCustomer
// @Tags         meeting Customer
// @Id					 ApiV1MeetingAppCustomerCreate
// @Param id path string false "id uuid"
// @Param request body api_structure.CustomerFilter true "update params"
// @Success      200  {object}  api_structure.Customer
// @Failure      400  {object}  error
// @Router       /api/v1/Customer/single [post]
func (controller *CustomerController) CreateCustomer(c *fiber.Ctx) error {
	data := api_structure.CustomerForm{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	result, rerr := controller.Svc.CreateCustomer(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Create Multiple in Customer
// @Description  Create Multiple in CreateCustomer
// @Tags         meeting Customer
// @Id					 ApiV1MeetingAppCustomerCreateMultiple
// @Param   Authorization  header     string     true  "Current Session Token" 	default([[authToken]])
// @Param   x-company  		 header     string     true  "Current CompanyID" 			default([[companyID]])
// @Param id path string false "id uuid"
// @Param request body api_structure.CustomerFilter true "update params"
// @Success      200  {array}  api_structure.Customer
// @Failure      400  {object}  error
// @Router       /api/v1/Customer/multiple [post]
func (controller *CustomerController) CreateCustomerMultiple(c *fiber.Ctx) error {

	bulkData := []api_structure.CustomerForm{}
	if err := c.BodyParser(&bulkData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	result, rerr := controller.Svc.CreateCustomerMultiple(bulkData)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Bulk Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Delete with path params in Customer
// @Description  Delete with path params in DeleteCustomer
// @Tags         meeting Customer
// @Id					 ApiV1MeetingAppCustomerDelete
// @Accept       json
// @Produce      json
// @Param id path string false "id uuid"
// @Success      200  {string}  string
// @Failure      400  {object}  error
// @Router       /api/v1/Customer/with-id/:id [delete]
func (controller *CustomerController) DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	idTrue, err := runtime_tools.Decrypt(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	deleteErr := controller.Svc.DeleteCustomer(idTrue)
	// err := controller.Svc.DeleteCustomer(id)
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
