package db

import (
	"context"
	"new/learning/user/util"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestCreateFile(t *testing.T) File {
// 	arg := CreateFileParams{
// 		FileName:   "file",
// 		Owner:      "trini",
// 		ChunkCount: 1,
// 	}

// 	File, err := testQueries.CreateFile(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, File)

// 	require.Equal(t, arg.FileName, File.FileName)
// 	require.Equal(t, arg.Owner, File.Owner)
// 	require.Equal(t, arg.ChunkCount, File.ChunkCount)

// 	require.NotZero(t, File.Hash)

//		return File
//	}
func createRandomFile(t *testing.T) File {

	arg := CreateFileParams{
		FileName:   util.RandomString(20),
		Owner:      util.RandomOwner(),
		ChunkCount: util.RandomInt(64, 64),
	}

	File, err := testQueries.CreateFile(context.Background(), arg)
	require.NoError(t, err)

	return File
}

func TestCreateFile(t *testing.T) {
	createRandomFile(t)
}
