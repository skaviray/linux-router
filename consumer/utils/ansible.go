package utils

import (
	"bytes"
	"context"
	"gateway-router-consumer/consts"
	"io"
	"log"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	results "github.com/apenella/go-ansible/v2/pkg/execute/result/json"
	"github.com/apenella/go-ansible/v2/pkg/execute/stdoutcallback"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

func Process(vars interface{}, action string) map[string]error {
	variables, _ := vars.(map[string]interface{})
	variables["action"] = action
	log.Println(variables)
	var res *results.AnsiblePlaybookJSONResults
	buff := new(bytes.Buffer)
	ansiblePlaybookOptions := &playbook.AnsiblePlaybookOptions{
		// ExtraVars: map[string]interface{}{
		// 	"extravar1":    "value11",
		// 	"extravar2":    "value12",
		// 	"ansible_port": "22225",
		// },
		ExtraVars:     variables,
		Inventory:     "/dev/null",
		SSHCommonArgs: "-o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null",
		Tags: "vxlan",
		Connection:    "local",
		// Inventory:     "127.0.0.1",
	}
	playbookCmd := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(consts.PlayBookPath),
		playbook.WithPlaybookOptions(ansiblePlaybookOptions),
	)
	exec := stdoutcallback.NewJSONStdoutCallbackExecute(
		execute.NewDefaultExecute(
			execute.WithCmd(playbookCmd),
			execute.WithErrorEnrich(playbook.NewAnsiblePlaybookErrorEnrich()),
			execute.WithWrite(io.Writer(buff)),
		),
	)
	err := exec.Execute(context.TODO())
	if err != nil {
		res, err = results.ParseJSONResultsStream(io.Reader(buff))
		if err != nil {
			// log.Println(err.Error())
			return map[string]error{
				"error": err,
			}
		}
		hosts := GetFailedHosts(res.Stats)
		// log.Println(hosts)
		err := GetFailureMessage(res.Plays, hosts)
		return err
		// log.Print(err)
	}
	return nil
}
