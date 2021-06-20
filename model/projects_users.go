package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/yerassyldanay/invest/utils/constants"
)

func (pu *ProjectsUsers) Chech_whether_user_is_assigned_to_project(tx *gorm.DB) error {
	var count int
	err := tx.Table("projects_users").Where("project_id = ? and user_id = ?", pu.ProjectId, pu.UserId).Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return errors.New("there is no such user & project relation")
	}

	return nil
}

func (pu *ProjectsUsers) OnlyCreate(tx *gorm.DB) (err error) {
	err = tx.Create(pu).Error
	return err
}

func (pu *ProjectsUsers) OnlyDelete(tx *gorm.DB) (err error) {
	err = tx.Delete(pu, "project_id = ? and user_id = ?", pu.ProjectId, pu.UserId).Error
	return err
}

func (pu *ProjectsUsers) OnlyDeleteRelation(tx *gorm.DB) (err error) {
	err = tx.Delete(pu, "project_id = ? and user_id = ?", pu.ProjectId, pu.UserId).Error
	return err
}

// count number of users assigned to the project
func (pu *ProjectsUsers) OnlyCountByRoleAndProjectId(role string, tx *gorm.DB) (int, error) {
	var count int
	if err := tx.Table("projects_users pu").
		Joins("join users u on pu.user_id = u.id").
		Joins("join roles r on u.role_id = r.id").
		Where("project_id = ? and r.name = ?", pu.ProjectId, role).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

//
func (pu *ProjectsUsers) OnlyAssignExpertsToProject(project_id uint64, tx *gorm.DB) error {
	main_query := `insert into projects_users select ? as project_id, u.id as id from users u ` +
		` join roles r on r.id = u.role_id where r.name = '` + constants.RoleExpert + `' ;`
	err := tx.Exec(main_query, project_id).Error

	return err
}

// assign expert to all projects
func (pu *ProjectsUsers) OnlyAssignExpertToAllProjects(tx *gorm.DB) error {
	main_query := `insert into projects_users (project_id, user_id) select id, ? from projects;`
	err := tx.Raw(main_query, pu.UserId).Error
	return err
}

func (pu *ProjectsUsers) OnlyDeleteByProjectId(project_id uint64, tx *gorm.DB) error {
	//fmt.Println(project_id)
	err := tx.Delete(pu, "project_id = ?", project_id).Error
	return err
}
