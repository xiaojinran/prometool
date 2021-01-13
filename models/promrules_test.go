package models

import (
	"gopkg.in/yaml.v2"
	"testing"
)

func TestNewGroups(t *testing.T) {
	r:=NewRecordRule("job:http_inprogress_requests:sum",
		  "round(rate(node_cpu_seconds_total[5m]),2)",
		  nil,)
	g:=NewGroup("example","2m",[]*Rule{r})
	gs:=NewGroups([]*RuleGroup{g})
	d, err := yaml.Marshal(gs)
	if err!=nil{
		t.Error(err.Error())
	}
	t.Logf("\n%s",string(d))
}
