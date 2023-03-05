package entities

import (
	"time"
	// uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"github.com/google/uuid"
)

type Nasabah struct {
	ID              uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Nama            string     `gorm:"type:varchar(100)" json:"nama"`
	NomorKTP        string     `gorm:"type:varchar(20)" json:"nomor_ktp"`
	TempatLahir     string     `gorm:"type:varchar(100)" json:"tempat_lahir"`
	TanggalLahir    string     `gorm:"type:varchar(100)" json:"tanggal_lahir"`
	AlamatAsal      string     `gorm:"type:text" json:"alamat_asal"`
	NomorHP         string     `gorm:"type:varchar(30)" json:"no_hp"`
	Email           string     `gorm:"type:varchar(50)" json:"email"`
	JenisKelamin    string     `gorm:"type:varchar(20)" json:"jenis_kelamin"`
	Pekerjaan       string     `gorm:"type:varchar(100)" json:"pekerjaan"`
	AlamatPekerjaan string     `gorm:"type:text" json:"alamat_pekerjaaan" `
	CreatedAt       time.Time  `gorm:"type:timestamp with time zone" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"type:timestamp with time zone" json:"updated_at"`
	Rekenings       []Rekening `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"rekening_numbers,omitempty"`
}
// unique, required

// 