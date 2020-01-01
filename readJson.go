package main
import(
	"encoding/json"
	"fmt"
	"io"
	//"io/ioutil"
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
	jsFile,err:=os.Open("post3.json")
	if err!=nil{
		fmt.Println("Error opening json file",err)
		return
	}
	defer jsFile.Close()

	/*Using Decoder  is json is dynamic*/
	decoder:=json.NewDecoder(jsFile)

	for{
		var post Post
		err=decoder.Decode(&post)
		if err==io.EOF {
			break
		}
		if err!=nil{
			fmt.Println("Error reading json data",err)
			return
		}

		fmt.Println(post)

	}



	/* using Read all if the json data is static*/
	//jsData,err:=ioutil.ReadAll(jsFile)
	//
	//if err!=nil{
	//	fmt.Println("Error reading json data",err)
	//	return
	//}

//	var post Post
//	json.Unmarshal(jsData,&post)
//	fmt.Println(post)
}