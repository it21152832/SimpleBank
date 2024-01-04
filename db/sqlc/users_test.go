package db

import (
	"context"
	"new/learning/user/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:        util.RandomString(6),
		FirstName:       util.RandomString(10),
		LastName:        util.RandomString(10),
		Email:           util.RandomString(10),
		Password:        util.RandomString(8),
		ConfirmPassword: util.RandomString(8),
	}
	User, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, User)

	require.Equal(t, arg.Username, User.Username)
	require.Equal(t, arg.FirstName, User.FirstName)
	require.Equal(t, arg.LastName, User.LastName)
	require.Equal(t, arg.Email, User.Email)
	require.Equal(t, arg.Password, User.Password)
	require.Equal(t, arg.ConfirmPassword, User.ConfirmPassword)

	require.NotZero(t, User.ID)
	require.NotZero(t, User.CreatedAt)
}
func TestGetUser(t *testing.T) {
	arg := CreateUserParams{
		Username:        util.RandomString(6),
		FirstName:       util.RandomString(10),
		LastName:        util.RandomString(10),
		Email:           util.RandomString(10),
		Password:        util.RandomString(8),
		ConfirmPassword: util.RandomString(8),
	}
	User, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, User)

	require.Equal(t, arg.Username, User.Username)
	require.Equal(t, arg.FirstName, User.FirstName)
	require.Equal(t, arg.LastName, User.LastName)
	require.Equal(t, arg.Email, User.Email)
	require.Equal(t, arg.Password, User.Password)
	require.Equal(t, arg.ConfirmPassword, User.ConfirmPassword)

	require.NotZero(t, User.ID)
	require.NotZero(t, User.CreatedAt)
}
