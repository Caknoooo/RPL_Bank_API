package router

import (
	"github.com/Caknoooo/nasabah-bank/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) *gin.Engine{
	server := gin.Default()

	nasabahHandler := handlers.NasabahHandler{DB: db}
	rekeningHandler := handlers.RekeningHandler{DB: db}

	// Nasabah 
	server.GET("/api/nasabah", nasabahHandler.HandleGetUserNasabah) // GetAll
	server.GET("/api/nasabah/:id", nasabahHandler.HandleGetUserNasabahByID) // GetID
	server.POST("/api/nasabah", nasabahHandler.HandleInsertUserNasabah) // Insert Nasabah
	server.PUT("/api/nasabah/:id", nasabahHandler.HandleUpdateUserNasabah) // Update Nasabah
	server.DELETE("/api/nasabah/:id", nasabahHandler.HandleDeleteUserNasabah) // Delete Nasabah

	// Rekening
	server.POST("/api/rekening", rekeningHandler.HandleInsertRekening) // Add Rekening yang mereferensi ke nasabah_id
	server.PUT("/api/rekening/:id", rekeningHandler.HandleUpdateRekening)
	server.DELETE("/api/rekening/:id", rekeningHandler.HandleDeleteRekening)
	return server
}