package connector

import (
	"blackboard/services"
	"blackboard/services/flag_handle"
	"blackboard/services/gitee"
)

//定义serve的映射关系
var serveMap = map[string]services.RepoInterface{
	"gitee": &gitee.GiteeServe{},
	// "github": &github.GithubServe{},
}

func RepoCreate() services.RepoInterface {
	return serveMap[flag_handle.PLATFORM]
}
