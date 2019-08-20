package projects

import (
	. "onsite/models"
)

type ProjectResponse struct {
	Project struct {
		ID               int    `json:"id"`
		Name             string `json:"name,omitempty" validate:"required"`
		ManagerProfileID int    `json:"manager_id" validate:"required"`
		CreatorProfileID int    `json:"creator_id" validate:"required"`
		SerialNo         string `json:"serial_no"`
		Address          string `json:"address"`
		// Value                float64 `json:"total_trade_value,omitempty"`
		// ContractorTotalClaim float64 `json:"contractor_total_claim,omitempty"`
		TotalContractValue float64 `json:"total_contract_value"`
		QuantitySurveyor   string  `json:"quantity_surveyor"`
		Notes              string  `json:"notes"`
	} `json:"project"`
}

/* Create */
func NewProjectResponse(p *Project) *ProjectResponse {
	resp := new(ProjectResponse)
	resp.Project.ID = p.ID
	resp.Project.Name = *p.Name.Ptr()
	resp.Project.ManagerProfileID = p.ManagerID
	resp.Project.CreatorProfileID = p.CreatorID

	resp.Project.SerialNo = p.SerialNo.String
	resp.Project.Address = p.Address.String

	if p.TotalContractValue.Valid {
		resp.Project.TotalContractValue = p.TotalContractValue.Float64
	}
	if p.QuantitySurveyor.Valid {
		resp.Project.QuantitySurveyor = p.QuantitySurveyor.String
	}
	if p.Notes.Valid {
		resp.Project.Notes = p.Notes.String
	}
	return resp
}

/* List */
type ProjectListResponse struct {
	Projects []*ProjectResponse `json:"projects"`
	Count    int                `json:"count"`
}

func NewProjectListResponse(projectsSrc []*Project, count int) *ProjectListResponse {
	r := new(ProjectListResponse)
	for _, src := range projectsSrc {
		item := NewProjectResponse(src)
		r.Projects = append(r.Projects, item)
	}
	r.Count = count
	return r
}

/* update */
type UpdateProjectResponse struct {
	Project struct {
		ID               int    `json:"id"`
		Name             string `json:"name"`
		ManagerProfileID int    `json:"manager_profile_id"`
		CreatorProfileID int    `json:"creator_profile_id"`
		SerialNo         string `json:"serial_no"`
		Address          string `json:"address"`
		// TotalTradeValue      float64 `json:"total_trade_value"`
		// ContractorTotalClaim float64 `json:"contractor_total_claim"`
		TotalContractValue float64 `json:"total_contract_value"`
		QuantitySurveyor   string  `json:"quantity_surveyor"`
		Notes              string  `json:"notes"`
	} `json:"project"`
}

func NewUpdateProjectResponse(p *Project) *UpdateProjectResponse {
	resp := new(UpdateProjectResponse)
	resp.Project.ID = p.ID
	resp.Project.Name = *p.Name.Ptr()
	resp.Project.ManagerProfileID = p.ManagerID
	resp.Project.CreatorProfileID = p.CreatorID
	resp.Project.SerialNo = *p.SerialNo.Ptr()
	resp.Project.Address = *p.Address.Ptr()
	// resp.Project.TotalTradeValue = *p.TotalTradeValue.Ptr()
	// resp.Project.ContractorTotalClaim = *p.ContractorTotalClaim.Ptr()
	resp.Project.TotalContractValue = *p.TotalContractValue.Ptr()
	resp.Project.QuantitySurveyor = *p.QuantitySurveyor.Ptr()
	resp.Project.Notes = *p.Notes.Ptr()
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
