package model

import (
	"fmt"
	"invest/utils"

	"sync"
)

//func (u *User) Add_stats_of_projects_to_this_investor(wg *sync.WaitGroup) {
//
//}

func (u *User) Add_statistics_to_this_user_on_project_statuses(wg *sync.WaitGroup) {
	defer wg.Done()

	if u.Role.Name == "" {
		if u.RoleId == 0 {
			_ = GetDB().Table(u.TableName()).Where("id = ?", u.Id).First(&u).Error
			u.Password = ""
		}

		err := GetDB().Table(Role{}.TableName()).Where("id = ?", u.RoleId).First(&u.Role).Error
		if err != nil {
		fmt.Println("could not load role. user stats on project statuses")
		return
		}
	}

	var err error
	var stats UserStatsRawList
	if u.Role.Name == utils.RoleInvestor {
		stats, err = u.get_stats_on_project_statuses_of_investor()
	} else {
		stats, err = u.get_stats_on_project_statuses_of_user_return_raw_stats()
	}

	if err != nil {
		return
	}

	var user_stat = UserStats{}
	for _, _ = range stats {

	}

	u.Statistics = user_stat
}

/*
	returns:
		{
			"newone": 2,
			...
		}
 */
func (u *User) get_stats_on_project_statuses_of_user_return_raw_stats() (UserStatsRawList, error) {
	var main_query = `select status, count(*) as number from projects_users pu join projects p on pu.project_id = p.id
    where pu.user_id = ? group by p.status;`

	var stats = UserStatsRawList{}
	if err := GetDB().Raw(main_query, u.Id).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return stats, nil
}

/*
	this is an another function, this is because of the database architecture
		Projects possess the column called 'offered_by_id', which is the id of the investor
 */
func (u *User) get_stats_on_project_statuses_of_investor() (UserStatsRawList, error) {
	var main_query = `select p.status as status, count(*) as number from projects p
    join users u on p.offered_by_id = u.id
    join roles r on u.role_id = r.id
	where u.id = ?
    group by p.status;`

	var stats = UserStatsRawList{}
	err := GetDB().Raw(main_query, u.Id).Scan(&stats).Error

	return stats, err
}
