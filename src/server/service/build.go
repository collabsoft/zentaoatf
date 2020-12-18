package service

import (
	"encoding/json"
	"fmt"
	"github.com/easysoft/zentaoatf/src/server/domain"
	logUtils "github.com/easysoft/zentaoatf/src/utils/log"
)

type BuildService struct {
	taskService *TaskService
}

func NewBuildService(taskService *TaskService) *BuildService {
	return &BuildService{taskService: taskService}
}

func (s *BuildService) Add(req domain.ReqData) (reply domain.OptResult) {
	build := domain.Build{}

	reqStr, _ := json.Marshal(req.Data)
	err := json.Unmarshal(reqStr, &build)
	if err != nil {
		logUtils.PrintTo(fmt.Sprintf("error: %v", err))
		return
	}

	size := s.taskService.GetSize()
	if size == 0 {
		s.taskService.Add(build)
		reply.Success("Success to add task.")
	} else {
		reply.Fail(fmt.Sprintf("already has %d jobs to be done.", size))
	}

	return
}
