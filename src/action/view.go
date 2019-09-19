package action

import scriptService "github.com/easysoft/zentaoatf/src/service/script"

func View(files []string, keywords string) {
	cases := scriptService.GetCaseByDirAndFile(files)

	scriptService.View(cases, keywords)
}
