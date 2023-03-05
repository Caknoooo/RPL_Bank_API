package entities

import (
	"time"
	// uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"github.com/google/uuid"
)

type Rekening struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	JenisAkun string    `gorm:"type:varchar(30)" json:"jenis_akun"`
	MataUang  string    `gorm:"type:varchar(50)" json:"mata_uang"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp with time zone" json:"updated_at"`
	NasabahID uuid.UUID `gorm:"foreignKey" json:"nasabah_id" binding:"required"`
	Nasabah		*Nasabah  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transaksi,omitempty"`
}