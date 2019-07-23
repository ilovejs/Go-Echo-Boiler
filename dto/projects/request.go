package projects

import (
	"github.com/labstack/echo"
	m "onsite/models"
)

type CreateProjectRequest struct {
	Project struct {
		ManagerProfileID int    `json:"manager_profile_id" validate:"required"`
		CreatorProfileID int    `json:"creator_profile_id" validate:"required"`
		Name             string `json:"name,omitempty" validate:"required"`
		// optional text
		SerialNo string `json:"serial_no,omitempty"`
		Details  string `json:"details,omitempty"`
		// for future input
		TotalItemBreakdown   float64 `json:"total_item_breakdown,omitempty"`
		ContractorTotalClaim float64 `json:"contractor_total_claim,omitempty"`
		TotalContractValue   float64 `json:"total_contract_value,omitempty"`
		QuantitySurveyor     string  `json:"quantity_surveyor,omitempty"`
	} `json:"project"`
}

func (r *CreateProjectRequest) Bind(c echo.Context, p *m.Project) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	p.ManagerProfileID = r.Project.ManagerProfileID
	p.CreatorProfileID = r.Project.CreatorProfileID
	//optional
	p.Name.SetValid(r.Project.Name)
	p.SerialNo.SetValid(r.Project.SerialNo)
	p.Details.SetValid(r.Project.Details)
	return nil
}
