package projects

import (
	. "onsite/models"
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
	resp := new(ProjectResponse)

	var pn *string
	pn = p.Name.Ptr()

	resp.Project.ID = p.ID
	resp.Project.Name = *pn
	//log.EnableColor()
	//log.Errorf("%T %v", pn, pn)
	//log.Errorf("%T %v", *pn, *pn)
	resp.Project.ManagerProfileID = p.ManagerID
	resp.Project.CreatorProfileID = p.CreatorID
	return resp
}
