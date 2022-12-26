package date2023

import "PackageFive/date2023/model"

func main(){
	 use:=model.UserInfo{Name: "小张",Age: 16,Hobby: []string{"运动","🏊‍♀️"},}
	 println(use.Hobby)
}
