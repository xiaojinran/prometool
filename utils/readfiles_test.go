package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestReadRulesExcel(t *testing.T) {
	rgs, err := ReadRulesFromExcel("prometheus.xlsx")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	d, err := yaml.Marshal(rgs)
	if err!=nil{
		t.Error(err.Error())
	}
	t.Logf("\n%s",string(d))
}

func TestReadAlertsFromExcel(t *testing.T) {
	rgs, err := ReadAlertsFromExcel("prometheus.xlsx")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	d, err := yaml.Marshal(rgs)
	if err!=nil{
		t.Error(err.Error())
	}
	t.Logf("\n%s",string(d))
}

func TestReadAllCellst(t *testing.T) {
	ReadAllCells("prometheus.xlsx")
}