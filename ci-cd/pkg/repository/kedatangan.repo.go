package repository

import (
	"fmt"
	"mini_project/pkg/database"
	"mini_project/pkg/models"
)

type KedatanganRepository interface {
	GetAllKedatanganKapal() ([]models.KedatanganKapal, error)
	GetKedatanganKapalById(id string) (models.KedatanganKapal, error)
	CreateNewKedatanganKapal(kedatangan models.KedatanganKapal) error
	UpdateKedatanganKapalById(id string, kedatangan models.KedatanganKapal) error
	GroupingDaerah() ([]models.KedatanganKapal, error)
}

func GetAllKedatanganKapal() ([]models.KedatanganKapal, error) {
	var listkapal []models.KedatanganKapal

	err := database.DB.Find(&listkapal).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return listkapal, err
}

func GetKedatanganKapalById(id string) (models.KedatanganKapal, error) {
	var kedatangan models.KedatanganKapal
	err := database.DB.First(&kedatangan, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return kedatangan, err
}

func CreateNewKedatanganKapal(newdatang models.KedatanganKapal) error {
	err := database.DB.Create(&newdatang).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateKedatanganKapalById(id string, newinfo models.KedatanganKapal) error {
	err := database.DB.Model(&newinfo).Where("id = ?", id).Updates(&newinfo).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func GroupingDaerah() ([]models.KedatanganKapal, error) {
	var Data []models.KedatanganKapal
	err := database.DB.Select("daerah_penangkapan", "id_kapal").Order("daerah_penangkapan, id_kapal asc").
		Find(&Data).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return Data, err
}
