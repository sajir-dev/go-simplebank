package db

import (
	"context"
	"testing"

	"github.com/sajir-dev/go-simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(
	t *testing.T,
	from_account int64,
	to_account int64,
	amount int64) Transfer {

	args := CreateTransferParams{
		FromAccountID: from_account,
		ToAccountID:   to_account,
		Amount:        amount,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.NotZero(t, transfer.ID)
	require.Equal(t, transfer.FromAccountID, from_account)
	require.Equal(t, transfer.ToAccountID, to_account)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)
	amount := utils.RandomMoney()

	createRandomTransfer(t, from_account.ID, to_account.ID, amount)
}

func TestGetTransfer(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)
	amount := utils.RandomMoney()

	transfer1 := createRandomTransfer(t, from_account.ID, to_account.ID, amount)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
}

func TestListTransfers(t *testing.T) {
	from_account := createRandomAccount(t)
	to_account := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		amount := utils.RandomMoney()
		createRandomTransfer(t, from_account.ID, to_account.ID, amount)
	}

	args := ListTransfersParams{
		FromAccountID: from_account.ID,
		ToAccountID:   to_account.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotZero(t, transfer)
	}
}
