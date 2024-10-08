package login

import (
	"encoding/json"
	"errors"
	fleetModel "fms/internal/interactor/models/fleets"
	userRoleModel "fms/internal/interactor/models/user_roles"
	fleetService "fms/internal/interactor/service/fleet"
	userRoleService "fms/internal/interactor/service/user_role"

	"fms/config"

	jwxModel "fms/internal/interactor/models/jwx"
	loginsModel "fms/internal/interactor/models/logins"
	usersModel "fms/internal/interactor/models/users"
	"fms/internal/interactor/pkg/jwx"
	"fms/internal/interactor/pkg/util"
	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
	jwxService "fms/internal/interactor/service/jwx"
	userService "fms/internal/interactor/service/user"

	"gorm.io/gorm"
)

type Manager interface {
	Login(input *loginsModel.Login) (int, any)
	Refresh(input *jwxModel.Refresh) (int, any)
}

type manager struct {
	UserService     userService.Service
	JwxService      jwxService.Service
	FleetService    fleetService.Service
	UserRoleService userRoleService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		UserService:     userService.Init(db),
		JwxService:      jwxService.Init(),
		FleetService:    fleetService.Init(db),
		UserRoleService: userRoleService.Init(db),
	}
}

func (m *manager) Login(input *loginsModel.Login) (int, any) {
	// 取得車隊ID
	fleetBase, err := m.FleetService.GetBySingle(&fleetModel.Field{
		FleetCode: util.PointerString(input.FleetCode),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.PermissionDenied, code.GetCodeMessage(code.PermissionDenied, "Incorrect fleet_code.")
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 驗證帳密
	acknowledge, userBase, err := m.UserService.AcknowledgeUser(&usersModel.Field{
		UserName: util.PointerString(input.UserName),
		Password: util.PointerString(input.Password),
	})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error(err)
			return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
		}
	}

	if acknowledge == false {
		return code.PermissionDenied, code.GetCodeMessage(code.PermissionDenied, "Incorrect user_name or password.")
	}

	_, err = m.UserRoleService.GetBySingle(&userRoleModel.Field{
		FleetID: *fleetBase.ID,
		UserID:  userBase.ID,
	})
	if err != nil {
		log.Error(err)
		return code.PermissionDenied, code.GetCodeMessage(code.PermissionDenied, "Incorrect user_name or fleet_code.")
	}

	// 取得使用者fleet & role
	userRoleBase, _ := m.UserRoleService.GetBySingle(&userRoleModel.Field{
		UserID: userBase.ID,
	})

	// 產生accessToken
	output := &jwxModel.Token{}
	accessToken, err := m.JwxService.CreateAccessToken(&jwxModel.JWX{
		UserID:  userBase.ID,
		FleetID: userRoleBase.FleetID,
		Name:    userBase.Name,
		Role:    userRoleBase.Roles.Name,
	})

	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	accessTokenByte, _ := json.Marshal(accessToken)
	err = json.Unmarshal(accessTokenByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 產生refreshToken
	refreshToken, err := m.JwxService.CreateRefreshToken(&jwxModel.JWX{
		UserID: userBase.ID,
	})

	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	refreshTokenByte, _ := json.Marshal(refreshToken)
	err = json.Unmarshal(refreshTokenByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.Role = *userRoleBase.Roles.Name
	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Refresh(input *jwxModel.Refresh) (int, any) {
	// 驗證refreshToken
	j := &jwx.JWT{
		PublicKey: config.RefreshPublicKey,
		Token:     input.RefreshToken,
	}

	if len(input.RefreshToken) == 0 {
		return code.JWTRejected, code.GetCodeMessage(code.JWTRejected, "RefreshToken is null.")
	}

	j, err := j.Verify()
	if err != nil {
		log.Error(err)
		return code.JWTRejected, code.GetCodeMessage(code.JWTRejected, "RefreshToken is error.")
	}

	field, err := m.UserService.GetBySingle(&usersModel.Field{
		ID: j.Other["user_id"].(string),
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.JWTRejected, code.GetCodeMessage(code.JWTRejected, "RefreshToken is error.")
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	// 取得使用者fleet & role
	userRoleBase, _ := m.UserRoleService.GetBySingle(&userRoleModel.Field{
		UserID: field.ID,
	})

	// 產生accessToken
	token, err := m.JwxService.CreateAccessToken(&jwxModel.JWX{
		UserID:  field.ID,
		FleetID: userRoleBase.FleetID,
		Name:    field.Name,
		Role:    userRoleBase.Roles.Name,
	})
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	token.RefreshToken = input.RefreshToken
	token.Role = *userRoleBase.Roles.Name
	return code.Successful, code.GetCodeMessage(code.Successful, token)
}
