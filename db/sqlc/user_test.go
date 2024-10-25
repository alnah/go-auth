package db

import (
	"context"
	"testing"

	th "github.com/alnah/go-auth/testhelper"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	want := createRandomUser()
	got, err := testQueries.CreateUser(context.Background(), want)
	require.NoError(t, err)

	require.Equal(t, got.Email, want.Email)
	require.Equal(t, got.Hash, want.Hash)
	require.Equal(t, got.FirstName, want.FirstName)
	require.Equal(t, got.LastName, want.LastName)
}

func createRandomUser() CreateUserParams {
	return CreateUserParams{
		Email:     th.RandomString(5) + "@" + th.RandomString(5) + ".com",
		Hash:      th.RandomString(64),
		FirstName: th.RandomString(7),
		LastName:  th.RandomString(7),
	}
}
