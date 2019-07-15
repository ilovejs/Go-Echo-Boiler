package handler

import "onsite/model"

//Read Project
type readProjectResponse struct {
	Project struct {
		Name string `json:"name"`
	} `json:"project"`
}

func newReadProjectResponse(p *model.Project) *readProjectResponse {
	r := new(readProjectResponse)
	r.Project.Name = p.Name
	return r
}

//Create Project
type createProjectResponse struct {
	Project struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	}
}

func newCreateProjectResponse(p *model.Project) *createProjectResponse {
	r := new(createProjectResponse)
	r.Project.Name = p.Name
	r.Project.Id = int(p.ID) //TODO
	return r
}
