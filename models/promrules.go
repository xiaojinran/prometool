package models

import "strings"

type RuleGroups struct {
	Groups []*RuleGroup `yaml:"groups"`
}

// RuleGroup is a list of sequentially evaluated recording and alerting rules.
type RuleGroup struct {
	Name     string         `yaml:"name"`
	Interval string `yaml:"interval,omitempty"`
	Rules    []*Rule     `yaml:"rules"`
}

// Rule describes an alerting or recording rule.
type Rule struct {
	Record      string            `yaml:"record,omitempty"`
	Alert       string            `yaml:"alert,omitempty"`
	Expr        string            `yaml:"expr"`
	For         string    `yaml:"for,omitempty"`
	Labels      map[string]string `yaml:"labels,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
}
func NewRecordRule(record,expr string,labels map[string]string)*Rule{
	return &Rule{
		Record: record,
		Expr:expr,
		Labels: labels,
	}
}

func NewAlertRule(alert,expr,duration string ,annotations,labels map[string]string)*Rule{
	return &Rule{
		Alert: alert,
		For: duration,
		Expr: expr,
		Labels: labels,
		Annotations: annotations,
	}
}

func NewGroup(name string,interval string,rules []*Rule)*RuleGroup{
	return &RuleGroup{
		Name: name,
		Interval: interval,
		Rules: rules,
	}
}

func NewGroups(groups []*RuleGroup)RuleGroups{
	return RuleGroups{
		Groups: groups,
	}
}
func LabelsString2map(str string)map[string]string{
	//字符串格式
	//  key1:val1
	//  key2:val2
	//
	strs := strings.Split(str, "\n")
	m := make(map[string]string)
	for _,v := range strs{
		ss := strings.Split(v,":")
		m[ss[0]] = ss[1]
	}
	return m
}
