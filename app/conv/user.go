package conv

import (
	"github.com/ProgrammingLab/toshokan/app/dao"
	"github.com/ProgrammingLab/toshokan/models"
)

func UserDAOToUserModel(u *dao.User, includeEmail bool) *models.User {
	res := &models.User{
		UserID:      int32(u.UserID),
		DisplayName: u.DisplayName,
		Name:        u.Name,
		IsAdmin:     u.IsAdmin,
	}
	if includeEmail {
		res.Email = u.Email
	}
	return res
}
