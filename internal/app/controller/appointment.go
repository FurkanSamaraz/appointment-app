package controller

import (
	"math"
	api_service "meeting_app/internal/app/service"
	api_structure "meeting_app/internal/app/structures"
	runtime_tools "meeting_app/runtime-tools"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AppointmentController struct {
	Svc api_service.AppointmentService
}
type PaginationAppointment struct {
	Items      []api_structure.AppointmentFilter `json:"items"`
	TotalCount int64                             `json:"totalCount"`
	Page       int                               `json:"page"`
	PageSize   int                               `json:"pageSize"`
	TotalPages int                               `json:"totalPages"`
	MaxPage    int64                             `json:"maxPage"`
	Total      int64                             `json:"total"`
	Last       bool                              `json:"last"`
	First      bool                              `json:"first"`
	Visible    int64                             `json:"visible"`
	Order      string                            `json:"order"`
}

// ShowAccount godoc
// @Summary      Pagination Appointment
// @Description  Pagination GetAppointmentWithPagination
// @Tags        	meeting Appointment
// @Id					 ApiV1MeetingAppAppointmentPagination
// @Success      200  {array}  api_structure.Appointment
// @Failure      400  {object}  error
// @Router       /api/v1/Appointment/pagination [get]
func (controller *AppointmentController) GetAppointmentWithPagination(c *fiber.Ctx) error {
	var Appointments []api_structure.AppointmentFilter
	db := controller.Svc.DB.Table("public.appointment")

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
		paginationResponse := PaginationAppointment{
			Items:      []api_structure.AppointmentFilter{},
			TotalCount: totalCount,
			Page:       page,
			PageSize:   pageSize,
			TotalPages: totalPages,
		}
		return c.JSON(paginationResponse)
	} else if page <= 0 {
		// Sayfa numarası belirtilmemişse veya geçersizse, tüm verileri çek
		db.Find(&Appointments)
	}
	offset := (page - 1) * pageSize
	db.Offset(offset).Limit(pageSize).Find(&Appointments)

	// Eksik kısımları doldururken aynı zamanda ilişkileri olan verileri kullanın

	// Sayfalama bilgilerini oluşturun
	paginationResponse := PaginationAppointment{
		Items:      Appointments,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
		MaxPage:    int64(totalPages - 1),
		Total:      totalCount,
		Last:       page >= totalPages,
		First:      page <= 1,
		Visible:    int64(len(Appointments)),
		Order:      orderParam,
	}

	// Diziyi JSON olarak cevaplayın
	return c.JSON(paginationResponse)
}

// ShowAccount godoc
// @Summary      Show Appointment
// @Description  Get GetAppointment
// @Tags        	meeting Appointment
// @Id					 ApiV1MeetingAppAppointmentSearch
// @Success      200  {array}  api_structure.Appointment
// @Failure      400  {object}  error
// @Router       /api/v1/Appointment/search [get]
func (controller *AppointmentController) GetAppointment(c *fiber.Ctx) error {
	var Appointment_table_name api_structure.AppointmentFilter

	response, err := controller.Svc.GetAppointment(Appointment_table_name)
	if err != nil {
		return c.Status(400).JSON(c.Status(400).JSON(c.JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})))
	}
	return c.JSON(response)
}

// ShowAccount godoc
// @Summary      Get Appointment by Unique IDs
// @Description  Get GetAppointmentById
// @Tags        	meeting Appointment
// @Id					 ApiV1MeetingAppAppointmentWithID
// @Success      200  {array}  api_structure.Appointment
// @Failure      400  {object}  error
// @Router       /api/v1/Appointment/with-id/:id [get]
func (controller *AppointmentController) GetAppointmentById(c *fiber.Ctx) error {

	id := c.Params("id")
	idTrue, err := runtime_tools.Decrypt(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	result, rerr := controller.Svc.GetAppointmentById(idTrue)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Update with id in Appointment
// @Description  Update with id in UpdateAppointment
// @Tags         meeting Appointment
// @Id					 ApiV1MeetingAppAppointmentUpdateWithID
// @Param id path string false "id uuid"
// @Param request body api_structure.AppointmentEdit true "update params"
// @Success      200  {object}  api_structure.Appointment
// @Failure      400  {object}  error
// @Router       /api/v1/Appointment/with-id/:id [put]
func (controller *AppointmentController) UpdateAppointment(c *fiber.Ctx) error {

	id := c.Params("id")

	editData := api_structure.AppointmentEdit{}
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
	uerr := controller.Svc.UpdateAppointment(idTrue, editData)
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
// @Summary      Update Multiple in Appointment
// @Description  Update Multiple in UpdateAppointment
// @Tags         meeting Appointment
// @Id					 ApiV1MeetingAppAppointmentsUpdateMultiple
// @Param id path string false "id uuid"
// @Param request body api_structure.AppointmentEdit true "update params"
// @Success      200  {array}  api_structure.Appointment
// @Failure      400  {object}  error
// @Router       /api/v1/Appointment/multiple [put]
func (controller *AppointmentController) UpdateAppointmentMultiple(c *fiber.Ctx) error {
	editData := []api_structure.AppointmentEdit{}
	if err := c.BodyParser(&editData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	uerr := controller.Svc.UpdateAppointmentMultiple(editData)
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
// @Summary      Create in Appointment
// @Description  Create in CreateAppointment
// @Tags         meeting Appointment
// @Id					 ApiV1MeetingAppAppointmentCreate
// @Param id path string false "id uuid"
// @Param request body api_structure.AppointmentFilter true "update params"
// @Success      200  {object}  api_structure.Appointment
// @Failure      400  {object}  error
// @Router       /api/v1/Appointment/single [post]
func (controller *AppointmentController) CreateAppointment(c *fiber.Ctx) error {
	data := api_structure.AppointmentForm{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	result, rerr := controller.Svc.CreateAppointment(data)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Create Multiple in Appointment
// @Description  Create Multiple in CreateAppointment
// @Tags         meeting Appointment
// @Id					 ApiV1MeetingAppAppointmentCreateMultiple
// @Param   Authorization  header     string     true  "Current Session Token" 	default([[authToken]])
// @Param   x-company  		 header     string     true  "Current CompanyID" 			default([[companyID]])
// @Param id path string false "id uuid"
// @Param request body api_structure.AppointmentFilter true "update params"
// @Success      200  {array}  api_structure.Appointment
// @Failure      400  {object}  error
// @Router       /api/v1/Appointment/multiple [post]
func (controller *AppointmentController) CreateAppointmentMultiple(c *fiber.Ctx) error {

	bulkData := []api_structure.AppointmentForm{}
	if err := c.BodyParser(&bulkData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Invalid Data",
			"message": err.Error(),
		})
	}

	result, rerr := controller.Svc.CreateAppointmentMultiple(bulkData)
	if rerr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Create Bulk Data",
			"message": rerr.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

// ShowAccount godoc
// @Summary      Delete with path params in Appointment
// @Description  Delete with path params in DeleteAppointment
// @Tags         meeting Appointment
// @Id					 ApiV1MeetingAppAppointmentDelete
// @Accept       json
// @Produce      json
// @Param id path string false "id uuid"
// @Success      200  {string}  string
// @Failure      400  {object}  error
// @Router       /api/v1/Appointment/with-id/:id [delete]
func (controller *AppointmentController) DeleteAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	idTrue, err := runtime_tools.Decrypt(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	deleteErr := controller.Svc.DeleteAppointment(idTrue)
	// err := controller.Svc.DeleteAppointment(id)
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
