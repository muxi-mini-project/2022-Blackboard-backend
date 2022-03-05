package collect

import (
	"blackboard/model"


)

func GetCollection(ID string,offset int,limit int)([]*model.Collection,error){
	list,err := model.GetCollection(ID,offset,limit)
	if err !=nil{
		return nil,err
	}
	return list,err
}