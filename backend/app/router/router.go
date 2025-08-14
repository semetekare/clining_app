package router

import (
	controllers "clining_app/app/controllers"
	"clining_app/app/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterHTTPEndpoints(router fiber.Router) {
	// OK запрос
	router.Get("api/Ok", controllers.Ok)

	// Services
	servicesGroup := router.Group("api/services")
	serviceCtrl := controllers.NewServiceController(service.NewServiceService())
	servicesGroup.Post("/", serviceCtrl.CreateService)
	servicesGroup.Get("/:id", serviceCtrl.GetService)
	servicesGroup.Get("/", serviceCtrl.GetAllServices)
	servicesGroup.Put("/:id", serviceCtrl.UpdateService)
	servicesGroup.Delete("/:id", serviceCtrl.DeleteService)

	// RoomTypes
	roomTypesGroup := router.Group("api/roomtypes")
	roomTypeCtrl := controllers.NewRoomTypeController(service.NewRoomTypeService())
	roomTypesGroup.Post("/", roomTypeCtrl.CreateRoomType)
	roomTypesGroup.Get("/:id", roomTypeCtrl.GetRoomType)
	roomTypesGroup.Get("/", roomTypeCtrl.GetAllRoomTypes)
	roomTypesGroup.Put("/:id", roomTypeCtrl.UpdateRoomType)
	roomTypesGroup.Delete("/:id", roomTypeCtrl.DeleteRoomType)

	// Surcharges
	surchargesGroup := router.Group("api/surcharges")
	surchargeCtrl := controllers.NewSurchargeController(service.NewSurchargeService())
	surchargesGroup.Post("/", surchargeCtrl.CreateSurcharge)
	surchargesGroup.Get("/:id", surchargeCtrl.GetSurcharge)
	surchargesGroup.Get("/", surchargeCtrl.GetAllSurcharges)
	surchargesGroup.Put("/:id", surchargeCtrl.UpdateSurcharge)
	surchargesGroup.Delete("/:id", surchargeCtrl.DeleteSurcharge)

	// ServicePrices
	servicePricesGroup := router.Group("api/serviceprices")
	servicePriceCtrl := controllers.NewServicePriceController(service.NewServicePriceService())
	servicePricesGroup.Post("/", servicePriceCtrl.CreateServicePrice)
	servicePricesGroup.Get("/:id", servicePriceCtrl.GetServicePrice)
	servicePricesGroup.Get("/", servicePriceCtrl.GetAllServicePrices)
	servicePricesGroup.Put("/:id", servicePriceCtrl.UpdateServicePrice)
	servicePricesGroup.Delete("/:id", servicePriceCtrl.DeleteServicePrice)

	// ExtraServices
	extraServicesGroup := router.Group("api/extraservices")
	extraServiceCtrl := controllers.NewExtraServiceController(service.NewExtraServiceService())
	extraServicesGroup.Post("/", extraServiceCtrl.CreateExtraService)
	extraServicesGroup.Get("/:id", extraServiceCtrl.GetExtraService)
	extraServicesGroup.Get("/", extraServiceCtrl.GetAllExtraServices)
	extraServicesGroup.Put("/:id", extraServiceCtrl.UpdateExtraService)
	extraServicesGroup.Delete("/:id", extraServiceCtrl.DeleteExtraService)

	// Clients
	clientsGroup := router.Group("api/clients")
	clientCtrl := controllers.NewClientController(service.NewClientService())
	clientsGroup.Post("/", clientCtrl.CreateClient)
	clientsGroup.Get("/:id", clientCtrl.GetClient)
	clientsGroup.Get("/", clientCtrl.GetAllClients)
	clientsGroup.Put("/:id", clientCtrl.UpdateClient)
	clientsGroup.Delete("/:id", clientCtrl.DeleteClient)

	// Orders
	ordersGroup := router.Group("api/orders")
	orderCtrl := controllers.NewOrderController(service.NewOrderService())
	ordersGroup.Post("/", orderCtrl.CreateOrder)
	ordersGroup.Get("/:id", orderCtrl.GetOrder)
	ordersGroup.Get("/", orderCtrl.GetAllOrders)
	ordersGroup.Put("/:id", orderCtrl.UpdateOrder)
	ordersGroup.Delete("/:id", orderCtrl.DeleteOrder)
}
