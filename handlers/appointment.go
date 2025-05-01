package handlers

import (
	"net/http"
	"strconv"

	"healthcare-api/database"
	"healthcare-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateAppointment godoc
// @Summary Create a new appointment
// @Tags appointments
// @Accept json
// @Produce json
// @Param appointment body models.Appointment true "Appointment data"
// @Success 201 {object} models.Appointment
// @Failure 400 {object} map[string]string "Appointment time must be in the future"
// @Router /appointments [post]
func CreateAppointment(c *gin.Context) {
	var appointment models.Appointment

	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment data"})
		return
	}

	// No appointment time in the past!
	if appointment.DateTime.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Appointment time must be in the future"})
		return
	}

	// Ensure patient exists!
	var patient models.Patient
	if err := database.DB.First(&patient, appointment.PatientID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Patient not found"})
		return
	}

	database.DB.Create(&appointment)
	c.JSON(http.StatusCreated, appointment)
}

// GetAppointments godoc
// @Summary Get all appointments with pagination
// @Tags appointments
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} models.Appointment
// @Router /appointments [get]
func GetAppointments(c *gin.Context) {
	var appointments []models.Appointment
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	database.DB.Limit(limit).Offset(offset).Find(&appointments)
	c.JSON(http.StatusOK, appointments)
}

// GetAppointmentByID godoc
// @Summary Get an appointment by ID
// @Tags appointments
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} models.Appointment
// @Failure 404 {object} map[string]string
// @Router /appointments/{id} [get]
func GetAppointmentByID(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment

	if err := database.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

// UpdateAppointment godoc
// @Summary Update an appointment by ID
// @Tags appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Param appointment body models.Appointment true "Updated appointment data"
// @Success 200 {object} models.Appointment
// @Failure 404 {object} map[string]string
// @Router /appointments/{id} [put]
func UpdateAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment

	if err := database.DB.First(&appointment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	var input models.Appointment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&appointment).Updates(input)
	c.JSON(http.StatusOK, appointment)
}

// DeleteAppointment godoc
// @Summary Delete an appointment by ID
// @Tags appointments
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /appointments/{id} [delete]
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	var appointment models.Appointment

	if err := database.DB.Delete(&appointment, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete appointment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted"})
}
