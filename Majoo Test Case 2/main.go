package main

import (
	"errors"
	"log"

	"gorm.io/gorm"
	"majoo.id/dua/Model"
)

type AreaRepository struct {
	DB *gorm.DB
}

type Repository struct {
	repository *AreaRepository
}

// func (_r *AreaRepository) InsertArea(param1 int32, param2 int64, type []string, ar *Model.Area) (err error) { SALAH
//
//   - variabel type; `type` adalah kata built in dari golang, tidak bisa digunakan sebagai variabel
//   - type dari variabel type; seharusnya string bukan []string
func (_r *AreaRepository) InsertArea(param1 int32, param2 int64, tipe string, ar *Model.Area) (err error) { // BENAR

	// inst := _r.DB.Model(ar)
	//
	// Hapus variable inst, karena tidak digunakan

	// Var area int // SALAH
	//
	// - `Var` huruf kecil
	// - area menggunakan tipe data int64
	var area int64 = 0 // BENAR

	//switch type {
	switch tipe {
	case "persegi_panjang":
		// var area := param1 * param2 // SALAH
		//
		// - variabel area sudah di inisiasi, maka tidak dapat menggunakan :=,
		// - tipe data dari param1 dan param2 berbeda, harus disamakan
		area = int64(param1) * param2 // BENAR
		ar.AreaValue = area
		ar.AreaType = "persegi panjang"

		// err = _r.DB.create(&ar).Error // SALAH
		//
		// - Public method menggunakan kapital di awal, Create
		err = _r.DB.Create(&ar).Error // BENAR
		if err != nil {
			return err
		}
	case "persegi":
		// var area = param1 * param2
		area = int64(param1) * param2
		ar.AreaValue = area
		ar.AreaType = "persegi"

		// err = _r.DB.create(&ar).Error
		//
		// - Public method menggunakan kapital di awal, Create
		// - Optimalkan kode
		if err = _r.DB.Create(&ar).Error; err != nil {
			return err
		}

	// case segitiga: SALAH
	//
	// - Kurang tanda yang menyatakan bahwa segitiga merupakan string: `"`, karena tipe data yang di switch berupa string
	case "segitiga": // BENAR

		// area = 0.5 * (param1 * param2) // SALAH
		//
		// - Selesaikan perhitungan matematika dalam bentuk float, karena jika float diubah ke int64, akan menjadi 0
		area = int64(0.5 * (float64(param1) * float64(param2))) // BENAR

		ar.AreaValue = area
		ar.AreaType = "segitiga"
		err = _r.DB.Create(&ar).Error
		if err != nil {
			return err
		}
	default:
		ar.AreaValue = 0
		ar.AreaType = "undefiend data"
		err = _r.DB.Create(&ar).Error
		if err != nil {
			return err
		}
	}

	// - Add return, karena function return error
	return
}

func Testmain() (err error) {
	_u := Repository{}
	//err = _u.repository.InsertArea(10, 10, ‘persegi’) // SALAH
	//
	// - Kurang argument ke-empat: *Model.Area
	err = _u.repository.InsertArea(10, 10, "persegi", &Model.Area{}) // BENAR

	if err != nil {
		// log.Error().Msg(err.Error())
		log.Println(err.Error())
		err = errors.New("en.ERROR_DATABASE")
		return err
	}
	return
}
