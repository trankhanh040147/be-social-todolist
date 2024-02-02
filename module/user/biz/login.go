package biz

import (
	"context"
	"go-200lab-g09/common"
	"go-200lab-g09/component/tokenprovider"
	"go-200lab-g09/module/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*model.User, error)
}

type loginBusiness struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider,
	hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry}
}

// 1. Find User, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1 Access token and refresh token
// 4. Return token to client
func (biz *loginBusiness) Login(ctx context.Context, data *model.UserLogin) (tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, model.ErrEmailOrPasswordInvalid
	}

	payload := &common.TokenPayLoad{
		UId:   user.Id,
		URole: user.Role.String(),
	}

	// fmt.Println("payload: ", payload)
	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	// fmt.Println("accessToken: ", accessToken)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	//refreshToken, err := biz.tokenProvider.GenerateRefreshToken(payload, biz.tkCfg.GetRtExp())
	// if err != nil {
	// 	return nil, common.ErrInternal(err)
	// }

	// account := &model.NewAccount(user, refreshToken)

	return accessToken, nil
}
