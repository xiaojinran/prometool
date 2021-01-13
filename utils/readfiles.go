package utils

import  (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/xiaojinran/prometool/models"
)

func ReadRulesFromExcel(filename string)(models.RuleGroups,error){
	f,err := excelize.OpenFile(filename)
	rules := []*models.Rule{}
	rg := []*models.RuleGroup{}
	var g *models.RuleGroup
	var rgs  models.RuleGroups
	if err!=nil{
		return rgs,err
	}
	//读取rules工作表
	rows,err:= f.Rows("rules")
    if err!=nil{
    	return rgs,err
    }
    var name,interval,record,expr,labels string
    for rows.Next(){
    	col,err :=rows.Columns()
    	if err!=nil{
    		return rgs,err
	    }
	    //判断是否为标题
	    if col[0] == "*组名称(name)"{
	   	continue
	   }
	    name = col[0]
	    interval = col[1]
	    record = col[2]
	    expr = col[3]
	    labels = col[4]
	    var rule *models.Rule
	    if labels==""{
		    rule=models.NewRecordRule(record,expr,nil)
	    }else{
		    rule=models.NewRecordRule(record,expr,models.LabelsString2map(labels))
	    }
	    if name!=""{
	    	rules =nil
		    rules = append(rules,rule)
		    g = models.NewGroup(name,interval,rules)
		    rg = append(rg,g)
	    }else{
		    rules = append(rules,rule)
		    g.Rules = rules
	    }
}
    rgs = models.NewGroups(rg)
	return rgs,nil
}

func ReadAlertsFromExcel(filename string)(models.RuleGroups,error){
	f,err := excelize.OpenFile(filename)
	rules := []*models.Rule{}
	rg := []*models.RuleGroup{}
	var g *models.RuleGroup
	var rgs  models.RuleGroups
	if err!=nil{
		return rgs,err
	}
	//读取rules工作表
	rows,err:= f.Rows("alerts")
	if err!=nil{
		return rgs,err
	}
	var name,interval,alert,expr,duration,labels,annotations string
	for rows.Next(){
		col,err :=rows.Columns()
		if err!=nil{
			return rgs,err
		}
		//判断是否为标题
		if col[0] == "*组名称(name)"{
			continue
		}
		name = col[0]
		interval = col[1]
		alert = col[2]
		expr = col[3]
		duration = col[4]
		labels = col[5]
		annotations = col[6]
		var rule *models.Rule
		if labels==""{
			rule=models.NewAlertRule(alert,expr,duration,models.LabelsString2map(annotations),nil)
		}else{
			rule=models.NewAlertRule(alert,expr,duration,models.LabelsString2map(annotations),models.LabelsString2map(labels))
		}
		if name!=""{
			rules =nil
			rules = append(rules,rule)
			g = models.NewGroup(name,interval,rules)
			rg = append(rg,g)
		}else{
			rules = append(rules,rule)
			g.Rules = rules
		}
	}
	rgs = models.NewGroups(rg)
	return rgs,nil
}

func ReadAllCells(filename string){
	f,_ := excelize.OpenFile(filename)
	//读取rules工作表
	rows,_:= f.Rows("alerts")
	for rows.Next() {
		cols,_ := rows.Columns()
		for k,v :=range cols{
			fmt.Println(k,v)
		}
	}
}
