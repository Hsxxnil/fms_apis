package jwx

import (
	"time"

	"fms/config"
	model "fms/internal/interactor/models/jwx"
	"fms/internal/interactor/pkg/jwx"
	"fms/internal/interactor/pkg/util"
	"fms/internal/interactor/pkg/util/log"
)

type Service interface {
	CreateAccessToken(input *model.JWX) (output *model.Token, err error)
	CreateRefreshToken(input *model.JWX) (output *model.Token, err error)
}

type service struct {
}

func Init() Service {
	return &service{}
}

func (s service) CreateAccessToken(input *model.JWX) (output *model.Token, err error) {
	other := map[string]any{
		"user_id":  input.UserID,
		"name":     input.Name,
		"role":     input.Role,
		"fleet_id": input.FleetID,
	}

	accessExpiration := util.NowToUTC().Add(time.Minute * 5).Unix()
	j := &jwx.JWE{
		PublicKey:     config.AccessPublicKey,
		Other:         other,
		ExpirationKey: accessExpiration,
	}

	j, err = j.Create()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	accessToken := j.Token

	output = &model.Token{
		AccessToken: accessToken,
	}

	return output, nil
}

func (s service) CreateRefreshToken(input *model.JWX) (output *model.Token, err error) {
	other := map[string]any{
		"user_id": input.UserID,
	}

	refreshExpiration := util.NowToUTC().Add(time.Hour * 8).Unix()
	j := &jwx.JWT{
		PrivateKey:    config.RefreshPrivateKey,
		Other:         other,
		ExpirationKey: refreshExpiration,
	}

	j, err = j.Create()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	refreshToken := j.Token
	output = &model.Token{
		RefreshToken: refreshToken,
	}

	return output, nil
}
