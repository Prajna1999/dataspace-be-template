package repository

import (
	"github.com/Prajna1999/dataspace-be/internal/models"
	"gorm.io/gorm"
)

type OrganizationRepository struct {
	DB *gorm.DB
}

// inititalize the repo and write the queries
// takes a pointer to the database and spit out the repository

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{DB: db}
}

// create an organization
func (r *OrganizationRepository) Create(organization *models.Organization) error {
	return r.DB.Create(organization).Error
}

// get an organization by ID
func (r *OrganizationRepository) getByID(id uint) (*models.Organization, error) {
	var organization models.Organization
	err := r.DB.First(&organization, id).Error
	return &organization, err
}
