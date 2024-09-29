package app

import (
	"backend/model"
	"context"
	"errors"
	"testing"
)

func TestRegister(t *testing.T) {
	th := SetupTestHelper(t)

	var (
		email    = "danielwetan.io@gmail.com"
		password = "12345678"

		ctx            = context.Background()
		invalidPayload = &model.CreateUserRequest{}
		payload        = &model.CreateUserRequest{
			Email:    email,
			Password: password,
		}

		user = &model.User{
			Email: email,
		}
	)

	testCases := []struct {
		name     string
		request  *model.CreateUserRequest
		wantErr  bool
		mockFunc func()
	}{
		{
			name:     "#1. Register with invalid payload",
			request:  invalidPayload,
			wantErr:  true,
			mockFunc: func() {},
		},
		{
			name:    "#2. Register with already exists user",
			request: payload,
			wantErr: true,
			mockFunc: func() {
				th.Store.EXPECT().GetUserByEmail(ctx, payload.Email).Return(user, errors.New("user already exists"))
			},
		},
		{
			name:    "#3. Register success",
			request: payload,
			wantErr: false,
			mockFunc: func() {
				th.Store.EXPECT().GetUserByEmail(ctx, payload.Email).Return(nil, nil)
				th.Store.EXPECT().CreateUser(ctx, payload).Return(nil)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			err := th.App.Register(ctx, test.request)
			if (err != nil) != test.wantErr {
				t.Errorf("expected error = %v, but got = %v", test.wantErr, err)
			}
		})
	}
}
