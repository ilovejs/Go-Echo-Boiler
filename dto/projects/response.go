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
		Value                float64 `json:"total_item_breakdown,omitempty"`
		ContractorTotalClaim float64 `json:"contractor_total_claim,omitempty"`
		TotalContractValue   float64 `json:"total_contract_value,omitempty"`
		QuantitySurveyor     string  `json:"quantity_surveyor,omitempty"`
	} `json:"project"`
}

func NewProjectResponse(p *Project) *ProjectResponse {
	resp := new(ProjectResponse)
	resp.Project.ID = p.ID
	resp.Project.Name = *p.Name.Ptr()
	resp.Project.ManagerProfileID = p.ManagerID
	resp.Project.CreatorProfileID = p.CreatorID
	return resp
}

/*update*/
type UpdateProjectResponse struct {
	Project struct {
		ID                   int     `json:"id"`
		Name                 string  `json:"name"`
		ManagerProfileID     int     `json:"manager_profile_id"`
		CreatorProfileID     int     `json:"creator_profile_id"`
		SerialNo             string  `json:"serial_no"`
		Details              string  `json:"details"`
		TotalTradeValue      float64 `json:"total_trade_value"`
		ContractorTotalClaim float64 `json:"contractor_total_claim"`
		TotalContractValue   float64 `json:"total_contract_value"`
		QuantitySurveyor     string  `json:"quantity_surveyor"`
	} `json:"project"`
}

func NewUpdateProjectResponse(p *Project) *UpdateProjectResponse {
	resp := new(UpdateProjectResponse)
	resp.Project.ID = p.ID
	resp.Project.Name = *p.Name.Ptr()
	resp.Project.ManagerProfileID = p.ManagerID
	resp.Project.CreatorProfileID = p.CreatorID
	resp.Project.SerialNo = *p.SerialNo.Ptr()
	resp.Project.Details = *p.Details.Ptr()
	resp.Project.TotalTradeValue = *p.TotalTradeValue.Ptr()
	resp.Project.ContractorTotalClaim = *p.ContractorTotalClaim.Ptr()
	resp.Project.TotalContractValue = *p.TotalContractValue.Ptr()
	resp.Project.QuantitySurveyor = *p.QuantitySurveyor.Ptr()
	return resp
}

/* delete */
type DeleteProjectResponse struct {
	Project struct {
		ID int `json:"id"`
	}
}

func NewDeleteProjectResponse(p *Project) *DeleteProjectResponse {
	resp := new(DeleteProjectResponse)
	resp.Project.ID = p.ID
	return resp
}
