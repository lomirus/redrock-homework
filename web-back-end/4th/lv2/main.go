package main

import (
	"fmt"
	"time"
)

func main()  {
	exit := make(chan int)
	go printOnTime("没有困难的工作，只有勇敢的打工人！",2,0,0)
	go printOnTime("早安，打工人！",8,0,0)
	go printInterval("芜湖！起飞！", 3600)
	<-exit
}

func getInterval(hour int, minute int, second int) int {
	now := time.Now()
	nowTimeStamp := now.Hour() * 3600 + now.Minute() * 60 + now.Second()
	givenTimeStamp := hour * 3600 + minute * 60 + second
	if givenTimeStamp < nowTimeStamp{
		return 86400 + givenTimeStamp - nowTimeStamp
	} else {
		return givenTimeStamp - nowTimeStamp
	}
}

func printOnTime(content string, hour int, minute int, second int){
	time.AfterFunc(time.Duration(getInterval(hour,minute,second)) * time.Second, func (){
		ticker := time.NewTicker(time.Second * 86400)
		for {
			<-ticker.C
			fmt.Println(content)
		}
	})
}

func printInterval(content string, interval int){
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		<-ticker.C
		fmt.Println(content)
	}
}