package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	
	"github.com/xiaojinran/prometool/utils"
)

func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetCursor(0, 0)
	termbox.HideCursor()
}

func main(){
	fmt.Println("This tool helps to generate prometheus rules.")
	pause()
	err:=writeRecordfile()
	if err!=nil{
		fmt.Println(err.Error())
		pause()
		os.Exit(127)
	}
	err = writeAlertstfile()
	if err!=nil{
		fmt.Println(err.Error())
		pause()
		os.Exit(127)
	}
	fmt.Println("Generate prometheus rules successfully.")
	pause()
	os.Exit(0)
}

func writeRecordfile() error{
	rgs, err := utils.ReadRulesFromExcel("prometheus.xlsx")
	if err!=nil{
		return err
	}
	d, err := yaml.Marshal(rgs)
	if err!=nil{
		return err
	}
	err = ioutil.WriteFile("my_records.rules", d, 0666)
	if err != nil {
		return err
	}
	return nil
}

func writeAlertstfile()error{
	rgs, err := utils.ReadAlertsFromExcel("prometheus.xlsx")
	if err!=nil{
		return err
	}
	d, err := yaml.Marshal(rgs)
	if err!=nil{
		return err
	}
	err = ioutil.WriteFile("my_alerts.rules", d, 0666)
	if err != nil {
		return err
	}
	return nil
}

func pause() {
	fmt.Println("请按任意键继续...")
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			break Loop
		}
	}
}