package services

import (
	"context"
	"testing"

	"github.com/Blaiseniyo/go-simple-bank-backend/models"
	"github.com/stretchr/testify/require"
)

func createEntry( t *testing.T) models.Entry{

	account := CreateAccounts(t)
	Entry_data := models.Entry{ Account_id: account.Id, Amount: 10 }
	Entry,err := CreateEntry(context.Background(),&Entry_data,TEST_DB)

	require.NoError(t,err)
	require.NotEmpty(t,Entry)

	require.Equal(t, Entry_data.Account_id,Entry.Account_id)
	require.Equal(t, Entry.Amount,Entry.Amount)

	require.NotZero(t,Entry.CreatedAt)

	return *Entry
}

func TestCreateEntry(t *testing.T){
	createEntry(t)
}

func TestGetEntry(t *testing.T){

	Entry :=models.Entry{}
	new_Entry:= createEntry(t)
	_,err := GetEntryById(context.Background(),&Entry,new_Entry.Id,TEST_DB)

	require.NoError(t,err)
	require.NotEmpty(t,Entry)

	require.Equal(t,new_Entry.Account_id,Entry.Account_id)
	require.Equal(t,new_Entry.Amount,Entry.Amount)
	require.Equal(t,new_Entry.Id,Entry.Id)
	
	require.NotZero(t,Entry.CreatedAt)
	
}

func TestUpdateEntry(t *testing.T){
	account := CreateAccounts(t)
	update_Entry :=models.Entry{Amount: 2,Account_id: account.Id }
	Entry := createEntry(t)
	_,err := UpdateEntry(context.Background(),&Entry,&update_Entry,TEST_DB)

	require.NoError(t,err)
	require.NotEmpty(t,Entry)

	require.Equal(t,Entry.Account_id,update_Entry.Account_id)
	require.Equal(t,Entry.Amount,update_Entry.Amount)

	require.NotZero(t,Entry.UpdatedAt)
}

func TestDeleteEntry(t *testing.T){
	
	Entry := createEntry(t)
	deleted_row,err := DeleteEntry(context.Background(),Entry.Id,TEST_DB)

	require.NoError(t,err)
	
	require.Equal(t,deleted_row,int64(1))
}


func TestListAllEntrys(t *testing.T){
	Entrys := []models.Entry{} 
	for i := 0; i < 10; i++ {
		createEntry(t)
	}

	_,err := ListAllEntries(context.Background(),&Entrys,5,5,TEST_DB)

	require.NoError(t,err)
	require.Equal(t,len(Entrys),5)

	for _,i := range Entrys{
		require.NotEmpty(t,i)
	}
}