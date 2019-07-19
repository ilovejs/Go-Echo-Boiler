package dto

import "onsite/model"

//Read Project
type readProjectResponse struct {
	Project struct {
		Name string `json:"name"`
	} `json:"project"`
}

func NewReadProjectResponse(p *model.Project) *readProjectResponse {
	r := new(readProjectResponse)
	r.Project.Name = p.Name
	return r
}

//Create Project
type CreateProjectResponse struct {
	Project struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	}
}

func NewCreateProjectResponse(p *model.Project) *CreateProjectResponse {
	r := new(CreateProjectResponse)
	r.Project.Name = p.Name
	r.Project.Id = int(p.ID) //TODO
	return r
}
