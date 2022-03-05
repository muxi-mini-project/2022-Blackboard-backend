package announcement

import(
	"blackboard/model"
)

func GetAnnouncements(offset int,limit int)([]*model.Announcement,error){
	list,err := model.GetAnnouncements(offset,limit)
	if err !=nil{
		return nil,err
	}
	return list,err
}