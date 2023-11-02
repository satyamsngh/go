package repository

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"job-portal-api/internal/models"
)

func (r *Repo) ViewJobDetailsBy(ctx context.Context, jid uint64) (models.Job, error) {
	var job models.Job
	result := r.DB.First(&job, jid)

	if result.Error != nil {
		return models.Job{}, result.Error
	}
	return job, nil
}

func (r *Repo) ViewJobByCompanyId(ctx context.Context, id uint) ([]models.Job, error) {
	var jobs []models.Job
	result := r.DB.Where("company_id = ?", id).Find(&jobs)

	if result.Error != nil {
		return nil, result.Error
	}

	return jobs, nil
}

func (r *Repo) CreateJob(ctx context.Context, jobData models.Job) (models.Job, error) {
	result := r.DB.Create(&jobData)

	if result.Error != nil {
		return models.Job{}, result.Error
	}
	return jobData, nil
}

func (r *Repo) FindAllJobs(ctx context.Context) ([]models.Job, error) {
	var jobs []models.Job
	result := r.DB.Find(&jobs)
	if result.Error != nil {
		return nil, result.Error
	}
	return jobs, nil

}

func (r *Repo) FindJob(ctx context.Context, cid uint64) ([]models.Job, error) {
	var jobData []models.Job
	result := r.DB.Where("cid = ?", cid).Find(&jobData)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find the company")
	}
	return jobData, nil
}
func (r *Repo) CreateCompany(ctx context.Context, companyData models.Companies) (models.Companies, error) {
	tx := r.DB.WithContext(ctx).Create(&companyData)
	// If there's an error with the database transaction.
	if tx.Error != nil {
		// Return an empty 'Inventory' struct and the error.
		return models.Companies{}, tx.Error
	}
	return companyData, nil
}

func (r *Repo) ViewCompanies(ctx context.Context) ([]models.Companies, error) {
	var comp = make([]models.Companies, 0, 10)
	var companies = make([]models.Companies, 0, 10)
	r.DB.Find(&comp)
	for _, company := range comp {
		companies = append(companies, company)

	}
	return companies, nil
}

func (r *Repo) ViewCompanyById(ctx context.Context, cid uint) ([]models.Companies, error) {
	var company []models.Companies
	result := r.DB.Where("id = ?", cid).First(&company)

	if result.Error != nil {
		return []models.Companies{}, result.Error
	}
	return company, nil
}
func (r *Repo) AutoMigrate() error {
	//if s.db.Migrator().HasTable(&User{}se
	//	return services
	//}

	err := r.DB.Migrator().AutoMigrate(&models.User{}, &models.Companies{}, &models.Job{})
	if err != nil {
		return err
	}

	// AutoMigrate function will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns
	err = r.DB.Migrator().AutoMigrate(&models.User{}, &models.Companies{})
	if err != nil {
		// If there is an error while migrating, log the error message and stop the program
		return err
	}
	return nil
}
