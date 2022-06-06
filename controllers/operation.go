package controllers

import (
    "os"
    "fmt"
    "io"
    "time"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/sandlayth/abyss/models"
)

var SupportedFileTypes = [1]string{"csv"}

func isFileTypeSupported(f string) (bool) {
    for _, v := range SupportedFileTypes {
        if f == v {
            return true
        }
    }
    return false
}

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

// POST /operations/:filetype
func ImportOperation(c *gin.Context) {
    if !isFileTypeSupported(c.Param("filetype")) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "File type not supported!"})
        return
    }
    fileHeader, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "File couldn't be imported!"})
        return
    }
    file, err := fileHeader.Open()
    if err != nil {
         c.JSON(http.StatusBadRequest, gin.H{"error": "File couldn't be read!"})
         return
    }
    defer file.Close()

    var operations []models.Operation
    buf := make([]byte, 1)
    line := ""
    for {
        n, err := file.Read(buf)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println(os.Stderr, err)
            continue
        }
        char := string(buf[:n])
        if char != "\n" {
            line += char
        } else {
            operation, err := models.ParseOperation(line)
            if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }
            operations = append(operations, operation)
            line = ""
        }
    }
    models.DB.Create(&operations)
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
