package organization

import (
	"blackboard/model"
)

func GetAllOrganizations(offset int, limit int) ([]*model.Organization, error) {
	list, err := model.GetAllOrganizations(offset, limit)
	if err != nil {
		return nil, err
	}
	return list, err
}


func GetCreated(ID string,offset int,limit int)([]*model.Organization,error){
	list ,err := model.GetCreated(ID,offset,limit)
	if err !=nil{
		return nil,err
	}
	return list,err 
}

func GetFollowing(ID string,offset int,limit int)([]*model.FollowingOrganization,error){
	list,err := model.GetFollowing(ID,offset,limit)
	if err !=nil{
		return nil,err
	}
	return list ,err
}