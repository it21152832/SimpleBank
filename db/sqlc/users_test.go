package db

import (
	"context"
	"new/learning/user/util"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestCreateUser(t *testing.T) {
// 	arg := CreateUserParams{
// 		Username: util.RandomString(6),
// 		HashedPassword: util.RandomString(20),
// 		FullName: util.RandomString(20),
// 		Email: util.RandomString(20),

// 	}
// 	User, err := testQueries.CreateUser(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, User)

// 	require.Equal(t, arg.Username, User.Username)
// 	require.Equal(t, arg.HashedPassword, User.HashedPassword)
// 	require.Equal(t, arg.FullName, User.FullName)
// 	require.Equal(t, arg.Email, User.Email)

// 	// require.NotZero(t, User.ID)
// 	require.NotZero(t, User.CreatedAt)
// }

func createRandomUser(t *testing.T) User {
	hashedPassword,err := util.HashPassword(util.RandomString(6))
	require.NoError(t,err)
	
	arg := CreateUserParams{
		Username: util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName: util.RandomOwner(),
		Email: util.RandomString(20),

	}
	User, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, User)

	require.Equal(t, arg.Username, User.Username)
	require.Equal(t, arg.HashedPassword, User.HashedPassword)
	require.Equal(t, arg.FullName, User.FullName)
	require.Equal(t, arg.Email, User.Email)

	// require.NotZero(t, User.ID)
	require.True(t, User.PasswordChangedAt.IsZero())
	require.NotZero(t, User.CreatedAt)
	
	return User
}

func TestCreateUser(t *testing.T){
	createRandomUser(t)
}
// func TestGetUser(t *testing.T) {
// 	arg := CreateUserParams{
// 		Username:        util.RandomString(6),
// 		FirstName:       util.RandomString(10),
// 		LastName:        util.RandomString(10),
// 		Email:           util.RandomString(10),
// 	}
// 	User, err := testQueries.CreateUser(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, User)

// 	require.Equal(t, arg.Username, User.Username)
// 	require.Equal(t, arg.FirstName, User.FirstName)
// 	require.Equal(t, arg.LastName, User.LastName)
// 	require.Equal(t, arg.Email, User.Email)
// 	require.Equal(t, arg.Password, User.Password)
// 	require.Equal(t, arg.ConfirmPassword, User.ConfirmPassword)

// 	require.NotZero(t, User.ID)
// 	require.NotZero(t, User.CreatedAt)
// }
