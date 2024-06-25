package entity_test

import (
	"testing"

	entity "github.com/go-frame/internals/entity"
	"github.com/stretchr/testify/require"
)

const (
	fakeBusinessTypeForTaxID   string = "114"
	fakeBusinessTypeForTaxName string = "any_titlebbb"
	fakeDescription            string = "any_author"
	fakeIsActive               bool   = true
)

func makeFakeBusinessType() *entity.BusinessType {
	buisnessType := entity.BusinessType{
		BusinessTypeForTaxID:   "1144555",
		BusinessTypeForTaxName: "any_title",
		Description:            "any_author",
		IsActive:               true,
	}

	return &buisnessType
}

func TestNewBusinessType(t *testing.T) {
	makeFakeBusinessType()
	newBusinessType, err := entity.NewBusinessType(fakeBusinessTypeForTaxID, fakeDescription, fakeBusinessTypeForTaxName, fakeIsActive)
	require.Nil(t, err)
	require.Equal(t, newBusinessType.BusinessTypeForTaxID, fakeBusinessTypeForTaxID)
	require.Equal(t, newBusinessType.BusinessTypeForTaxName, fakeBusinessTypeForTaxName)
	require.Equal(t, newBusinessType.Description, fakeDescription)
	require.Equal(t, newBusinessType.IsActive, fakeIsActive)

}
