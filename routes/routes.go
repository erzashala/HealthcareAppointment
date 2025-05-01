package routes

import (
	"healthcare-api/handlers"
	"healthcare-api/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures API routes and middleware
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger())

	// Patients CRUD
	patients := r.Group("/patients")
	{
		patients.POST("", handlers.CreatePatient)
		patients.GET("", handlers.GetPatients)
		patients.GET("/:id", handlers.GetPatientByID)
		patients.PUT("/:id", handlers.UpdatePatient)
		patients.DELETE("/:id", handlers.DeletePatient)
		patients.GET("/:id/appointments", handlers.GetAppointmentByID)
	}

	// Appointments CRUD
	appointments := r.Group("/appointments")
	{
		appointments.POST("", handlers.CreateAppointment)
		appointments.GET("", handlers.GetAppointments)
		appointments.GET("/:id", handlers.GetAppointmentByID)
		appointments.PUT("/:id", handlers.UpdateAppointment)
		appointments.DELETE("/:id", handlers.DeleteAppointment)
	}

	return r
}
