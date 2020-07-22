package main
 
import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)
 
type Contact struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}
 
func file_get_contents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
 
func main() {
	var c Contact
	var content []byte
	var err error
	
	content, err = file_get_contents("gopher.json")
	if err != nil {
		fmt.Println("读取json文件失败"+ err.Error())
		return
	}
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		fmt.Println("错误: ", err.Error())
		return
	}
	fmt.Println(c)
	fmt.Println(c.Name)
}