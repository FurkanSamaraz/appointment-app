package app

import (
	"meeting_app/internal/app/controller"
	"meeting_app/internal/app/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Run(app fiber.Router, db *gorm.DB) {

	{
		{
			ServiceGroup := app.Group("/service")
			Servicesservices := controller.ServiceController{Svc: service.ServiceService{DB: db}}
			ServiceGroup.Get("/pagination", Servicesservices.GetServiceWithPagination).Name("GetServiceWithPagination")
			ServiceGroup.Get("/search", Servicesservices.GetService).Name("GetService")
			ServiceGroup.Post("/single", Servicesservices.CreateService).Name("CreateService")
			ServiceMultipleGroup := ServiceGroup.Group("/multiple")
			{
				ServiceMultipleGroup.Post("/", Servicesservices.CreateServiceMultiple).Name("CreateServiceMultiple")
				ServiceMultipleGroup.Put("/", Servicesservices.UpdateServiceMultiple).Name("UpdateServiceMultiple")
			}
			subServiceGroup := ServiceGroup.Group("/with-id/:Id")
			{
				subServiceGroup.Get("/", Servicesservices.GetServiceById).Name("GetServiceById")
				subServiceGroup.Put("/", Servicesservices.UpdateService).Name("UpdateService")
				subServiceGroup.Delete("/", Servicesservices.DeleteService).Name("DeleteService")
			}
		}

		{
			CustomerGroup := app.Group("/customer")
			CustomerServices := controller.CustomerController{Svc: service.CustomerService{DB: db}}
			CustomerGroup.Get("/pagination", CustomerServices.GetCustomerWithPagination).Name("GetCustomerWithPagination")
			CustomerGroup.Get("/search", CustomerServices.GetCustomer).Name("GetCustomer")
			CustomerGroup.Post("/single", CustomerServices.CreateCustomer).Name("CreateCustomer")
			CustomerMultipleGroup := CustomerGroup.Group("/multiple")
			{
				CustomerMultipleGroup.Post("/", CustomerServices.CreateCustomerMultiple).Name("CreateCustomerMultiple")
				CustomerMultipleGroup.Put("/", CustomerServices.UpdateCustomerMultiple).Name("UpdateCustomerMultiple")
			}
			subCustomerGroup := CustomerGroup.Group("/with-id/:Id")
			{
				subCustomerGroup.Get("/", CustomerServices.GetCustomerById).Name("GetCustomerById")
				subCustomerGroup.Put("/", CustomerServices.UpdateCustomer).Name("UpdateCustomer")
				subCustomerGroup.Delete("/", CustomerServices.DeleteCustomer).Name("DeleteCustomer")
			}
		}

		{
			AppointmentGroup := app.Group("/appointment")
			AppointmentServices := controller.AppointmentController{Svc: service.AppointmentService{DB: db}}
			AppointmentGroup.Get("/pagination", AppointmentServices.GetAppointmentWithPagination).Name("GetAppointmentWithPagination")
			AppointmentGroup.Get("/search", AppointmentServices.GetAppointment).Name("GetAppointment")
			AppointmentGroup.Post("/single", AppointmentServices.CreateAppointment).Name("CreateAppointment")
			AppointmentMultipleGroup := AppointmentGroup.Group("/multiple")
			{
				AppointmentMultipleGroup.Post("/", AppointmentServices.CreateAppointmentMultiple).Name("CreateAppointmentMultiple")
				AppointmentMultipleGroup.Put("/", AppointmentServices.UpdateAppointmentMultiple).Name("UpdateAppointmentMultiple")
			}
			subAppointmentGroup := AppointmentGroup.Group("/with-id/:Id")
			{
				subAppointmentGroup.Get("/", AppointmentServices.GetAppointmentById).Name("GetAppointmentById")
				subAppointmentGroup.Put("/", AppointmentServices.UpdateAppointment).Name("UpdateAppointment")
				subAppointmentGroup.Delete("/", AppointmentServices.DeleteAppointment).Name("DeleteAppointment")
			}
		}

		{
			CustomerServiceHistoryGroup := app.Group("/customer-service-history")
			CustomerServiceHistoryServices := controller.CustomerServiceHistoryController{Svc: service.CustomerServiceHistoryService{DB: db}}
			CustomerServiceHistoryGroup.Get("/pagination", CustomerServiceHistoryServices.GetCustomerServiceHistoryWithPagination).Name("GetCustomerServiceHistoryWithPagination")
			CustomerServiceHistoryGroup.Get("/search", CustomerServiceHistoryServices.GetCustomerServiceHistory).Name("GetCustomerServiceHistory")
			CustomerServiceHistoryGroup.Post("/single", CustomerServiceHistoryServices.CreateCustomerServiceHistory).Name("CreateCustomerServiceHistory")
			CustomerServiceHistoryMultipleGroup := CustomerServiceHistoryGroup.Group("/multiple")
			{
				CustomerServiceHistoryMultipleGroup.Post("/", CustomerServiceHistoryServices.CreateCustomerServiceHistoryMultiple).Name("CreateCustomerServiceHistoryMultiple")
				CustomerServiceHistoryMultipleGroup.Put("/", CustomerServiceHistoryServices.UpdateCustomerServiceHistoryMultiple).Name("UpdateCustomerServiceHistoryMultiple")
			}
			subCustomerServiceHistoryGroup := CustomerServiceHistoryGroup.Group("/with-id/:Id")
			{
				subCustomerServiceHistoryGroup.Get("/", CustomerServiceHistoryServices.GetCustomerServiceHistoryById).Name("GetCustomerServiceHistoryById")
				subCustomerServiceHistoryGroup.Put("/", CustomerServiceHistoryServices.UpdateCustomerServiceHistory).Name("UpdateCustomerServiceHistory")
				subCustomerServiceHistoryGroup.Delete("/", CustomerServiceHistoryServices.DeleteCustomerServiceHistory).Name("DeleteCustomerServiceHistory")
			}
		}
	}

}
