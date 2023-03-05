package handlers

import (
	"errors"
	"net/http"

	"github.com/Caknoooo/nasabah-bank/common"
	"github.com/Caknoooo/nasabah-bank/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RekeningHandler struct {
	DB *gorm.DB
}

func (h *RekeningHandler) HandleInsertRekening(ctx *gin.Context) {
	var rekening entities.Rekening

	if err := ctx.ShouldBind(&rekening); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if err := h.DB.Create(&rekening).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, common.Response{
			Status:  false,
			Message: "Permintaan dari body tidak valid",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Response{
		Status:  true,
		Message: "Berhasil menambahkan rekening",
		Data:    rekening,
	})
}

func (h *RekeningHandler) HandleUpdateRekening(ctx *gin.Context){
	rekeningID := ctx.Param("id")

	var rekening entities.Rekening
	if err := h.DB.First(&rekening, "id = ?", rekeningID).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			ctx.JSON(http.StatusBadRequest, common.Response{
				Status: false,
				Message: "ID Rekening tidak ditemukan",
				Data: nil,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, common.Response{
			Status: false,
			Message: "Gagal mendapatkan rekening",
			Data: nil,
		})
		return
	}

	// Parsing data yang akan diupdate dari body
	if err := ctx.ShouldBind(&rekening); err != nil{
		ctx.JSON(http.StatusBadRequest, common.Response{
			Status: false,
			Message: "Permintaan dari body tidak valid",
			Data: nil,
		})
		return
	}

	if err := h.DB.Save(&rekening).Error; err != nil{
		ctx.JSON(http.StatusBadRequest, common.Response{
			Status: false,
			Message: "Gagal melakukan update rekening",
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Response{
		Status: true,
		Message: "Berhasil update rekening",
		Data: rekening,
	})
}

func (h *RekeningHandler) HandleDeleteRekening(ctx *gin.Context) {
	rekeningID := ctx.Param("id")

	var rekening entities.Rekening
	if err := h.DB.First(&rekening, "id = ?", rekeningID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusBadRequest, common.Response{
				Status:  false,
				Message: "ID Rekening tidak ditemukan",
				Data:    nil,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, common.Response{
			Status:  false,
			Message: "Gagal mendapatkan rekening",
			Data:    nil,
		})
		return
	}

	if err := h.DB.Delete(&rekening).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, common.Response{
			Status:  false,
			Message: "Gagal menghapus rekening",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Response{
		Status:  true,
		Message: "Berhasil menghapus rekening",
		Data:    nil,
	})
}
