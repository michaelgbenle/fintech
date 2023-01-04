package repository

import "github.com/michaelgbenle/fintech/internal/models"

func (p *Postgres) FindStudentByEmail(email string) (*models.User, error) {
	//var student *model.Student
	student := &models.User{}
	if err := p.DB.Where("email = ?", email).First(&student).Error; err != nil {
		return nil, err
	}
	return student, nil
}