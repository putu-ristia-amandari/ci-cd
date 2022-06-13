package controllers

import (
	"mini_project/pkg/models"
	"mini_project/pkg/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllMuatanController(c echo.Context) error {
	data, err := repository.GetAllMuatanKapal()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                       "Kamu Berhasil Melihat Semua Data Jenis Muatan",
		"Daftar Muatan Kapal Perikanan": data,
	})
}

func GetMuatanController(c echo.Context) error {
	MuatanKapal, err := repository.GetMuatanKapalById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, MuatanKapal)

}

func DeleteMuatanController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeleteMuatanKapalById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menghapus Data Pelabuhan dengan id '" + id + "'",
	})

}

func CreateMuatanController(c echo.Context) error {
	JenisMuatan := models.Muatan{}
	c.Bind(&JenisMuatan)
	err := repository.CreateNewMuatanKapal(JenisMuatan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Data Jenis Muatan",
	})

}

func UpdateMuatanController(c echo.Context) error {
	MuatanKapal := models.Muatan{}
	c.Bind(&MuatanKapal)
	err := repository.UpdateMuatanKapalById(c.Param("id"), MuatanKapal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Memperbaharui Data",
	})

}
