package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("1 – Print Covid19 cases in Pakistan")
	fmt.Println("2 – Print Covid19 cases in SouthKorea")
	fmt.Println("3 – Print Covid19 cases in France")
	fmt.Println("4 – Print my personalized message to Coronavirus")
	fmt.Println("0 – Exit")
	options,x:=strconv.Atoi(os.Args[0])
	if x!= nil{
	}
	for ;options!=0;
	{
		if options==1{
			fmt.Println("1 – Print Covid19 cases in Pakistan")
		} else if options==2 {
			fmt.Println("2 – Print Covid19 cases in SouthKorea")
		} else if options==3{
			fmt.Println("3 – Print Covid19 cases in France")
		} else if options==4{
			fmt.Println("4 – Print my personalized message to Coronavirus")
		} else if options==0{
		}
		options,x=strconv.Atoi(os.Args[0])
		if x!= nil{
		}
	}
}
