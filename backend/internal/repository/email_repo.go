package repository

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"vendorsage/backend/internal/models"
)

var ErrDuplicate = errors.New("duplicate email")

type EmailRepository struct {
	db *gorm.DB
}

func NewEmailRepository(db *gorm.DB) *EmailRepository {
	return &EmailRepository{db: db}
}

func (r *EmailRepository) Create(email *models.Email) error {
	result := r.db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "content_hash"}},
			DoNothing: true,
		}).
		Create(email)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrDuplicate
	}
	return nil
}

func (r *EmailRepository) UpdateType(id uint, emailType models.EmailType) {
	r.db.Model(&models.Email{}).Where("id = ?", id).Update("type", emailType)
}

func (r *EmailRepository) BulkCreate(emails []models.Email) error {
	return r.db.CreateInBatches(emails, 100).Error
}

func (r *EmailRepository) FindAll() ([]models.Email, error) {
	var emails []models.Email
	err := r.db.Find(&emails).Error
	return emails, err
}
