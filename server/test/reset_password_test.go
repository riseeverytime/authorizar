package test

import (
	"testing"

	"github.com/authorizerdev/authorizer/server/db"
	"github.com/authorizerdev/authorizer/server/enum"
	"github.com/authorizerdev/authorizer/server/graph/model"
	"github.com/authorizerdev/authorizer/server/resolvers"
	"github.com/stretchr/testify/assert"
)

func resetPasswordTest(s TestSetup, t *testing.T) {
	email := "reset_password." + s.TestInfo.Email
	_, err := resolvers.Signup(s.Ctx, model.SignUpInput{
		Email:           email,
		Password:        s.TestInfo.Password,
		ConfirmPassword: s.TestInfo.Password,
	})

	_, err = resolvers.ForgotPassword(s.Ctx, model.ForgotPasswordInput{
		Email: email,
	})
	assert.Nil(t, err, "no errors for forgot password")

	verificationRequest, err := db.Mgr.GetVerificationByEmail(email, enum.ForgotPassword.String())
	assert.Nil(t, err, "should get forgot password request")

	_, err = resolvers.ResetPassword(s.Ctx, model.ResetPasswordInput{
		Token:           verificationRequest.Token,
		Password:        "test1",
		ConfirmPassword: "test",
	})

	assert.NotNil(t, err, "passowrds don't match")

	_, err = resolvers.ResetPassword(s.Ctx, model.ResetPasswordInput{
		Token:           verificationRequest.Token,
		Password:        "test1",
		ConfirmPassword: "test1",
	})

	assert.Nil(t, err, "password changed successfully")

	cleanData(email)
}
