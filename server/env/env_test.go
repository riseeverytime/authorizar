package env

import (
	"testing"

	"github.com/authorizerdev/authorizer/server/constants"
	"github.com/authorizerdev/authorizer/server/enum"
	"github.com/stretchr/testify/assert"
)

func TestEnvs(t *testing.T) {
	constants.ENV_PATH = "../../.env.sample"
	InitEnv()

	assert.Equal(t, constants.ADMIN_SECRET, "admin")
	assert.Equal(t, constants.ENV, "production")
	assert.Equal(t, constants.DATABASE_URL, "data.db")
	assert.Equal(t, constants.DATABASE_TYPE, enum.Sqlite.String())
	assert.True(t, constants.DISABLE_EMAIL_VERIFICATION)
	assert.True(t, constants.DISABLE_MAGIC_LOGIN)
	assert.False(t, constants.DISABLE_BASIC_AUTHENTICATION)
	assert.Equal(t, constants.JWT_TYPE, "HS256")
	assert.Equal(t, constants.JWT_SECRET, "random_string")
	assert.Equal(t, constants.JWT_ROLE_CLAIM, "role")
	assert.EqualValues(t, constants.ROLES, []string{"user"})
	assert.EqualValues(t, constants.DEFAULT_ROLES, []string{"user"})
	assert.EqualValues(t, constants.PROTECTED_ROLES, []string{"admin"})
	assert.EqualValues(t, constants.ALLOWED_ORIGINS, []string{"*"})
}
