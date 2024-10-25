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

	assertUserEqual(t, got, want)
}

func createRandomUser() CreateUserParams {
	return CreateUserParams{
		Email:     th.RandomString(5) + "@" + th.RandomString(5) + ".com",
		Hash:      th.RandomString(64),
		FirstName: th.RandomString(7),
		LastName:  th.RandomString(7),
	}
}

func TestGetUser(t *testing.T) {
	want := createRandomUser()
	_, err := testQueries.CreateUser(context.Background(), want)
	require.NoError(t, err)

	got, err := testQueries.GetUser(context.Background(), want.Email)
	require.NoError(t, err)

	assertUserEqual(t, got, want)
}

func assertUserEqual(t *testing.T, got UserCore, want CreateUserParams) {
	t.Helper()

	require.Equal(t, got.Email, want.Email)
	require.Equal(t, got.Hash, want.Hash)
	require.Equal(t, got.FirstName, want.FirstName)
	require.Equal(t, got.LastName, want.LastName)
}
