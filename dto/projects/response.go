package projects

import (
	. "onsite/models"
	"onsite/utils"
)

type ProjectResponse struct {
	Project struct {
		ID               int    `json:"id"`
		ManagerProfileID int    `json:"manager_profile_id" validate:"required"`
		CreatorProfileID int    `json:"creator_profile_id" validate:"required"`
		Name             string `json:"name,omitempty" validate:"required"`
		//
		SerialNo string `json:"serial_no,omitempty"`
		Details  string `json:"details,omitempty"`
		//
		TotalItemBreakdown   float64 `json:"total_item_breakdown,omitempty"`
		ContractorTotalClaim float64 `json:"contractor_total_claim,omitempty"`
		TotalContractValue   float64 `json:"total_contract_value,omitempty"`
		QuantitySurveyor     string  `json:"quantity_surveyor,omitempty"`
	} `json:"project"`
}

func NewProjectResponse(p *Project) *ProjectResponse {
	res := new(ProjectResponse)
	res.Project.ID = p.ID
	pn := new(string)
	err := p.Name.Scan(&pn)
	utils.LogIf(err)
	res.Project.Name = *pn
	return res
}
