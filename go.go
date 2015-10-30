package gotool

import (
	"log"
	"os/exec"

	"github.com/mason-ci/mason/tool"
)

func init() {

	gotool := tool.NewTool("go", nil)
	/*
		gotool.HandleFunc(tool.PreBuildPhase, func(j tool.Job) {
			cmd := exec.Command("go", "fmt")
			cmd.Dir = j.Workspace
			cmd.Run()
		})
	*/

	gotool.HandleFunc(tool.BuildPhase, func(j tool.Job) {
		log.Printf("Gogogogogog in %v", j.Workspace)
		cmd := exec.Command("go", "build")
		cmd.Dir = j.Workspace
		cmd.Run()
	})

	/*
	   gotool.HandleFunc("build", tool.BuildPhase, func(j tool.Job) {
	   	cmd := exec.Command("go", "build")
	   	cmd.Dir = j.Workspace
	   	cmd.Run()
	   })
	*/
}
