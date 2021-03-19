package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/sajir-dev/go-simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	amount := utils.RandomMoney()

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    amount,
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NotEmpty(t, entry)
	require.NoError(t, err)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NotEmpty(t, entry1)
	require.NotEmpty(t, entry2)
	require.NoError(t, err)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	args := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: utils.RandomMoney(),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), args)

	require.NotEmpty(t, entry1)
	require.NotEmpty(t, entry2)
	require.NoError(t, err)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, args.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	err := testQueries.DeleteEntries(context.Background(), entry1.ID)

	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Empty(t, entry2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListEntries(t *testing.T) {

	account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		amount := utils.RandomMoney()

		args := CreateEntryParams{
			AccountID: account.ID,
			Amount:    amount,
		}

		testQueries.CreateEntry(context.Background(), args)
		// require.NoError(t, err)
	}

	listEntryArgs := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), listEntryArgs)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
