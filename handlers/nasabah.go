package handlers

import (
	"errors"
	"net/http"

	"github.com/Caknoooo/nasabah-bank/common"
	"github.com/Caknoooo/nasabah-bank/entities"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NasabahHandler struct {
	DB *gorm.DB
}

func (h *NasabahHandler) HandleGetUserNasabah(ctx *gin.Context) {
	var userNasabah []entities.Nasabah
	tx := h.DB.Preload("Rekenings").Find(&userNasabah) 
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest, common.Response{
			Status:  false,
			Message: tx.Error.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Response{
		Status:  true,
		Message: "Users fetched successfully",
		Data:    userNasabah,
	})
}

func (h *NasabahHandler) HandleGetUserNasabahByID(ctx *gin.Context) {
	nasabahID := ctx.Param("id") 

	var userNasabah entities.Nasabah
	if err := h.DB.Preload("Rekenings").First(&userNasabah, "id = ?", nasabahID).Error; err != nil { 
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, common.Response{
				Status:  false,
				Message: "ID Nasabah tidak ditemukan",
				Data:    nil,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, common.Response{
			Status:  false,
			Message: "Gagal mendapatkan User Nasabah",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Response{
		Status:  true,
		Message: "User fetched successfully",
		Data:    userNasabah,
	})
}

func (h *NasabahHandler) HandleInsertUserNasabah(ctx *gin.Context) {
	var userNasabah entities.Nasabah 

	err := ctx.ShouldBind(&userNasabah) 
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	tx := h.DB.Create(&userNasabah)
	if tx.Error != nil {
		ctx.JSON(http.StatusBadRequest, common.Response{
			Status:  false,
			Message: tx.Error.Error(),
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, common.Response{
		Status:  true,
		Message: "Nasabah berhasil dibuat",
		Data:    userNasabah,
	})
}

func (h *NasabahHandler) HandleUpdateUserNasabah(ctx *gin.Context) {
	nasabahId := ctx.Param("id")

	var userNasabah entities.Nasabah
	if err := h.DB.First(&userNasabah, "id = ?", nasabahId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusBadRequest, common.Response{
				Status:  false,
				Message: "ID Nasabah tidak ditemukan",
				Data:    nil,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, common.Response{
			Status:  false,
			Message: "Gagal mendapatkan User Nasabah",
			Data:    nil,
		})
		return
	}

	if err := ctx.ShouldBindJSON(&userNasabah); err != nil {
		ctx.JSON(http.StatusBadRequest, common.Response{
			Status:  false,
			Message: "Permintaan dari body tidak valid",
			Data:    nil,
		})
		return
	}

	if err := h.DB.Omit("rekening_numbers").Save(&userNasabah).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, common.Response{
			Status:  false,
			Message: "Gagal memperbarui User Nasabah",
			Data:    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Response{
		Status:  true,
		Message: "Berhasil memperbarui data",
		Data:    userNasabah,
	})
}

func (h *NasabahHandler) HandleDeleteUserNasabah(ctx *gin.Context) {
    nasabahID := ctx.Param("id")

    var userNasabah entities.Nasabah
    if err := h.DB.First(&userNasabah, "id = ?", nasabahID).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            ctx.JSON(http.StatusBadRequest, common.Response{
                Status:  false,
                Message: "ID Nasabah tidak ditemukan",
                Data:    nil,
            })
            return
        }

        ctx.JSON(http.StatusBadRequest, common.Response{
            Status:  false,
            Message: "Gagal mendapatkan User Nasabah",
            Data:    nil,
        })
        return
    }

    if err := h.DB.Delete(&userNasabah).Error; err != nil {
        ctx.JSON(http.StatusBadRequest, common.Response{
            Status:  false,
            Message: "Gagal menghapus User Nasabah",
            Data:    nil,
        })
        return
    }

    ctx.JSON(http.StatusOK, common.Response{
        Status:  true,
        Message: "Berhasil menghapus User Nasabah",
        Data:    nil,
    })
}

