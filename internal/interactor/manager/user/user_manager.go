package user

import (
	"encoding/json"
	"errors"
	fleetModel "fms/internal/interactor/models/fleets"
	fleetService "fms/internal/interactor/service/fleet"

	"fms/internal/interactor/pkg/util"

	userRoleModel "fms/internal/interactor/models/user_roles"
	userModel "fms/internal/interactor/models/users"
	userService "fms/internal/interactor/service/user"
	userRoleService "fms/internal/interactor/service/user_role"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *userModel.Create) (int, any)
	GetByList(input *userModel.Fields) (int, any)
	GetBySingle(input *userModel.Field) (int, any)
	Delete(trx *gorm.DB, input *userModel.Field) (int, any)
	Update(trx *gorm.DB, input *userModel.Update) (int, any)
}

type manager struct {
	UserService     userService.Service
	FleetService    fleetService.Service
	UserRoleService userRoleService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		UserService:     userService.Init(db),
		FleetService:    fleetService.Init(db),
		UserRoleService: userRoleService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *userModel.Create) (int, any) {
	defer trx.Rollback()

	// 判斷使用者ID是否重複
	quantity, _ := m.UserService.GetByQuantity(&userModel.Field{
		UserName: util.PointerString(input.UserName),
	})

	if quantity > 0 {
		log.Info("UserName already exists. UserName: ", input.UserName)
		return code.BadRequest, code.GetCodeMessage(code.BadRequest, "User already exists.")
	}

	// 取得車隊ID
	fleetBase, err := m.FleetService.GetBySingle(&fleetModel.Field{
		FleetCode: util.PointerString(input.FleetCode),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, "The fleet ID does not exist.")
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	userBase, err := m.UserService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步新增user_roles
	_, err = m.UserRoleService.WithTrx(trx).Create(&userRoleModel.Create{
		UserID:    *userBase.ID,
		RoleID:    input.RoleID,
		FleetID:   *fleetBase.ID,
		CreatedBy: input.CreatedBy,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, userBase.ID)
}

func (m *manager) GetByList(input *userModel.Fields) (int, any) {
	output := &userModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, userBase, err := m.UserService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	userByte, err := json.Marshal(userBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(userByte, &output.Users)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, user := range output.Users {
		if userBase[i].CreatedByUsers != nil && userBase[i].UpdatedByUsers != nil {
			user.CreatedBy = *userBase[i].CreatedByUsers.Name
			user.UpdatedBy = *userBase[i].UpdatedByUsers.Name
		}

		// 取得該使用者的fleet & role
		userRoleBase, _ := m.UserRoleService.GetBySingle(&userRoleModel.Field{
			UserID: userBase[i].ID,
		})
		if userRoleBase != nil {
			user.RoleID = *userRoleBase.RoleID
			user.RoleName = *userRoleBase.Roles.Name
			user.FleetID = *userRoleBase.FleetID
			user.FleetCode = *userRoleBase.Fleets.FleetCode
		}
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *userModel.Field) (int, any) {
	userBase, err := m.UserService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &userModel.Single{}
	userByte, _ := json.Marshal(userBase)
	err = json.Unmarshal(userByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if userBase.CreatedByUsers != nil && userBase.UpdatedByUsers != nil {
		output.CreatedBy = *userBase.CreatedByUsers.Name
		output.UpdatedBy = *userBase.UpdatedByUsers.Name
	}

	// 取得該使用者的fleet & role
	userRoleBase, _ := m.UserRoleService.GetBySingle(&userRoleModel.Field{
		UserID: userBase.ID,
	})
	if userRoleBase != nil {
		output.RoleID = *userRoleBase.RoleID
		output.RoleName = *userRoleBase.Roles.Name
		output.FleetID = *userRoleBase.FleetID
		output.FleetCode = *userRoleBase.Fleets.FleetCode
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(trx *gorm.DB, input *userModel.Field) (int, any) {
	defer trx.Rollback()

	_, err := m.UserService.GetBySingle(&userModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.UserService.WithTrx(trx).Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 同步刪除user_roles
	err = m.UserRoleService.WithTrx(trx).Delete(&userRoleModel.Field{
		UserID: util.PointerString(input.ID),
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(trx *gorm.DB, input *userModel.Update) (int, any) {
	defer trx.Rollback()

	// 驗證舊帳密
	if input.Password != nil {
		acknowledge, _, err := m.UserService.AcknowledgeUser(&userModel.Field{
			ID:       input.ID,
			Password: input.OldPassword,
		})
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				log.Error(err)
				return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
			}
		}

		if acknowledge == false {
			return code.PermissionDenied, code.GetCodeMessage(code.PermissionDenied, "Incorrect password.")
		}
	}

	// 判斷使用者ID是否重複
	if input.UserName != nil {
		quantity, _ := m.UserService.GetByQuantity(&userModel.Field{
			UserName: input.UserName,
		})

		if quantity > 0 {
			log.Info("UserName already exists. UserName: ", input.UserName)
			return code.BadRequest, code.GetCodeMessage(code.BadRequest, "User already exists.")
		}
	}

	userBase, err := m.UserService.GetBySingle(&userModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	if input.FleetCode != nil {
		// 取得車隊ID
		fleetBase, err := m.FleetService.GetBySingle(&fleetModel.Field{
			FleetCode: input.FleetCode,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, "The fleet ID does not exist.")
			}

			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}

		// 同步更新user_roles
		err = m.UserRoleService.WithTrx(trx).Update(&userRoleModel.Update{
			UserID:    util.PointerString(input.ID),
			RoleID:    input.RoleID,
			FleetID:   fleetBase.ID,
			UpdatedBy: input.UpdatedBy,
		})
		if err != nil {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	err = m.UserService.WithTrx(trx).Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, userBase.ID)
}
