package dto

import (
	"github.com/labstack/echo/v4"
	"onsite/model"
)

type CreateProjectRequest struct {
	Project struct {
		Name                string `json:"name"`
		Number              string `json:"number"`
		Details             string `json:"details"`
		ContractorId        string `json:"contractorId"`
		QuantitySurveyor    string `json:"quantitySurveyor"`
		TotalContractAmount string `json:"totalContractAmount"`
	} `json:"project"`
}

func (r *CreateProjectRequest) Bind(c echo.Context, p *model.Project) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	p.Name = r.Project.Name
	p.SerialNo = r.Project.Number
	p.Detail = r.Project.Details
	//1 project could have n contractor (see from login user perspective or different company admin)
	//p.ContractorProjects = []*model.ContractorProject{
	//	{
	//		User: nil, //&
	//
	//	},
	//}
	return nil
}
