package role

import (
	"encoding/json"
	"errors"

	"fms/internal/interactor/pkg/util"

	roleModel "fms/internal/interactor/models/roles"
	roleService "fms/internal/interactor/service/role"

	"gorm.io/gorm"

	"fms/internal/interactor/pkg/util/code"
	"fms/internal/interactor/pkg/util/log"
)

type Manager interface {
	Create(trx *gorm.DB, input *roleModel.Create) (int, any)
	GetByList(input *roleModel.Fields) (int, any)
	GetBySingle(input *roleModel.Field) (int, any)
	Delete(input *roleModel.Field) (int, any)
	Update(input *roleModel.Update) (int, any)
}

type manager struct {
	RoleService roleService.Service
}

func Init(db *gorm.DB) Manager {
	return &manager{
		RoleService: roleService.Init(db),
	}
}

func (m *manager) Create(trx *gorm.DB, input *roleModel.Create) (int, any) {
	defer trx.Rollback()

	// 判斷角色是否重複
	quantity, _ := m.RoleService.GetByQuantity(&roleModel.Field{
		Name: util.PointerString(input.Name),
	})

	if quantity > 0 {
		log.Info("Role already exists. Name: ", input.Name)
		return code.BadRequest, code.GetCodeMessage(code.BadRequest, "Role already exists.")
	}

	roleBase, err := m.RoleService.WithTrx(trx).Create(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	trx.Commit()
	return code.Successful, code.GetCodeMessage(code.Successful, roleBase.ID)
}

func (m *manager) GetByList(input *roleModel.Fields) (int, any) {
	output := &roleModel.List{}
	output.Limit = input.Limit
	output.Page = input.Page
	quantity, roleBase, err := m.RoleService.GetByList(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Total.Total = quantity
	roleByte, err := json.Marshal(roleBase)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}
	output.Pages = util.Pagination(quantity, output.Limit)
	err = json.Unmarshal(roleByte, &output.Roles)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	for i, role := range output.Roles {
		role.CreatedBy = *roleBase[i].CreatedByUsers.Name
		role.UpdatedBy = *roleBase[i].UpdatedByUsers.Name
	}

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) GetBySingle(input *roleModel.Field) (int, any) {
	roleBase, err := m.RoleService.GetBySingle(input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output := &roleModel.Single{}
	roleByte, _ := json.Marshal(roleBase)
	err = json.Unmarshal(roleByte, &output)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	output.CreatedBy = *roleBase.CreatedByUsers.Name
	output.UpdatedBy = *roleBase.UpdatedByUsers.Name

	return code.Successful, code.GetCodeMessage(code.Successful, output)
}

func (m *manager) Delete(input *roleModel.Field) (int, any) {
	_, err := m.RoleService.GetBySingle(&roleModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.RoleService.Delete(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, "Delete ok!")
}

func (m *manager) Update(input *roleModel.Update) (int, any) {
	roleBase, err := m.RoleService.GetBySingle(&roleModel.Field{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.DoesNotExist, code.GetCodeMessage(code.DoesNotExist, err.Error())
		}

		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	err = m.RoleService.Update(input)
	if err != nil {
		log.Error(err)
		return code.InternalServerError, code.GetCodeMessage(code.InternalServerError, err.Error())
	}

	return code.Successful, code.GetCodeMessage(code.Successful, roleBase.ID)
}
