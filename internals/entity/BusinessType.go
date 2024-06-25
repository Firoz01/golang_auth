package entity

type BusinessType struct {
	BusinessTypeForTaxID   string `json:"business_type_for_tax_id" bun:"business_type_for_tax_id"`
	BusinessTypeForTaxName string `json:"business_type_for_tax_name" bun:"business_type_for_tax_name"`
	Description            string `json:"description" bun:"description"`
	IsActive               bool   `json:"is_active" bun:"is_active"`
}

func NewBusinessType(BusinessTypeForTaxID string, Description string, BusinessTypeForTaxName string, IsActive bool) (*BusinessType, error) {
	businessType := BusinessType{
		BusinessTypeForTaxID:   BusinessTypeForTaxID,
		BusinessTypeForTaxName: BusinessTypeForTaxName,
		Description:            Description,
		IsActive:               IsActive,
	}

	return &businessType, nil
}
