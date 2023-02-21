package services

import (
	"context"
	"testing"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"github.com/stretchr/testify/require"
)

func createEntry(t *testing.T) models.Entry {

	account := CreateAccounts(t)
	entry_data := models.Entry{Account_id: account.Id, Amount: 10}
	entry, err := CreateEntry(context.Background(), &entry_data, TEST_DB)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry_data.Account_id, entry.Account_id)
	require.Equal(t, entry.Amount, entry.Amount)

	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createEntry(t)
}

func TestGetEntry(t *testing.T) {

	new_entry := createEntry(t)
	entry, err := GetEntryById(context.Background(), new_entry.Id, TEST_DB)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, new_entry.Account_id, entry.Account_id)
	require.Equal(t, new_entry.Amount, entry.Amount)
	require.Equal(t, new_entry.Id, entry.Id)

	require.NotZero(t, entry.CreatedAt)

}

func TestUpdateEntry(t *testing.T) {
	account := CreateAccounts(t)
	update_Entry := models.Entry{Amount: 2, Account_id: account.Id}
	entry := createEntry(t)
	updated_entry, err := UpdateEntry(context.Background(), &entry, &update_Entry, TEST_DB)

	require.NoError(t, err)
	require.NotEmpty(t, updated_entry)

	require.Equal(t, updated_entry.Account_id, update_Entry.Account_id)
	require.Equal(t, updated_entry.Amount, update_Entry.Amount)

	require.NotZero(t, updated_entry.UpdatedAt)
}

func TestDeleteEntry(t *testing.T) {

	entry := createEntry(t)
	deleted_row, err := DeleteEntry(context.Background(), entry.Id, TEST_DB)

	require.NoError(t, err)

	require.Equal(t, deleted_row, int64(1))
}

func TestListAllEntrys(t *testing.T) {

	for i := 0; i < 10; i++ {
		createEntry(t)
	}

	entries, err := ListAllEntries(context.Background(), 5, 5, TEST_DB)

	require.NoError(t, err)
	require.Equal(t, len(entries), 5)

	for _, i := range entries {
		require.NotEmpty(t, i)
	}
}
