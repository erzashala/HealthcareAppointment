package handlers

import (
	"net/http"
	"strconv"

	"healthcare-api/database"
	"healthcare-api/models"

	"github.com/gin-gonic/gin"
)

// CreatePatient godoc
// @Summary Create a new patient
// @Tags patients
// @Accept json
// @Produce json
// @Param patient body models.Patient true "Patient data"
// @Success 201 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Router /patients [post]
func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create patient"})
		return
	}

	c.JSON(http.StatusCreated, patient)
}

// GetPatients godoc
// @Summary List all patients with pagination
// @Tags patients
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} models.Patient
// @Router /patients [get]
func GetPatients(c *gin.Context) {
	var patients []models.Patient
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	database.DB.Preload("Appointments").Limit(limit).Offset(offset).Find(&patients)
	c.JSON(http.StatusOK, patients)
}

// GetPatientByID godoc
// @Summary Get a patient by ID
// @Tags patients
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} models.Patient
// @Failure 404 {object} map[string]string
// @Router /patients/{id} [get]
func GetPatientByID(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if err := database.DB.Preload("Appointments").First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, patient)
}

// UpdatePatient godoc
// @Summary Update a patient by ID
// @Tags patients
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Param patient body models.Patient true "Updated patient data"
// @Success 200 {object} models.Patient
// @Failure 404 {object} map[string]string
// @Router /patients/{id} [put]
func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if err := database.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var input models.Patient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&patient).Updates(input)
	c.JSON(http.StatusOK, patient)
}

// DeletePatient godoc
// @Summary Delete a patient by ID
// @Tags patients
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /patients/{id} [delete]
func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient

	if err := database.DB.Delete(&patient, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
}
