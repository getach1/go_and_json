package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


	type Post struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id int `json:"id"`
	Name string `json:"name"`
}
type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main()  {
 post:=Post{
 	Id:1,
 	Content:"Goood morning",
 	Author:Author{Id:2,Name:"Gech"},
 	Comments:[]Comment{{Id:3,Content:"bravooo",Author:"Adam"},{Id:3,Content:"bravooo",Author:"Adam"},},
	}
	jsFile,err:=os.Create("post2.json")
	if err!=nil {
		fmt.Println("Error creareing json file",err)
		return
	}
	defer jsFile.Close()

	/*UING DECODER*/
	//encoder:=json.NewEncoder(jsFile)
 //
	//err=encoder.Encode(&post)
 //if err!=nil{
 //	fmt.Println("error encoding",err)
	// return
 //}

 /*USING MARSHALiNDENT*/
 output,err:=json.MarshalIndent(&post,"","\t\t")
	if err!=nil{
		fmt.Println("Error Mashaling",err)
		return
	}
	_,err=jsFile.Write([]byte(output))
	if err != nil {
		fmt.Println("Error Writting to js file",err)
		return
	}

	/*Other method*/
	err=ioutil.WriteFile("post3.json",[]byte(output),066)
	if err != nil {
		fmt.Println("Error Writting to js file",err)
		return
	}
}
