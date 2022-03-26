package controllers

import (
    "time"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/sandlayth/abyss/models"
)

type CreateOperationInput struct {
    Date          time.Time `json: "date" binding:"required"`
    EffectiveDate time.Time `json: "effectiveDate" binding:"required"`
    Label         string    `json: "id" binding:"required"`
    Amount        float64   `json: "amount" binding:"required"`
}

type UpdateOperationInput struct {
    Date          time.Time `json: "date"`
    EffectiveDate time.Time `json: "effectiveDate"`
    Label         string    `json: "id"`
    Amount        float64   `json: "amount"`
}

// GET /operations
func FindOperations(c *gin.Context) {
    var operations []models.Operation
    models.DB.Find(&operations)
    c.JSON(http.StatusOK, gin.H{"data": operations})
}

// GET /operations/:id
func FindOperation(c *gin.Context) {
    var operation models.Operation
    if err := models.DB.Where("id = ?", c.Param("id")).First(&operation).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": operation})
}

// POST /operations
func CreateOperation(c *gin.Context) {
    var input CreateOperationInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    operation := models.Operation{Date: input.Date, EffectiveDate: input.EffectiveDate, Label: input.Label, Amount: input.Amount}
    models.DB.Create(&operation)
    c.JSON(http.StatusOK, gin.H{"data": operation})
}

// PATCH /operations/:id
func UpdateOperation(c *gin.Context) {
    var operation models.Operation
    if err := models.DB.Where("id = ?", c.Param("id")).First(&operation).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }
    var input UpdateOperationInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    models.DB.Model(&operation).Updates(input)
    c.JSON(http.StatusOK, gin.H{"data": operation})
}


// DELETE /operations/:id
func DeleteOperation(c *gin.Context) {
    var operation models.Operation
    if err := models.DB.Where("id = ?", c.Param("id")).First(&operation).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }
    models.DB.Delete(&operation)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
