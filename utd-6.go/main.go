package main
import "errors"
type mymap map[string]any

var(
	err1 error =errors.New("nil map")
	err2 error =errors.New("key not exist")
)
func (m mymap)Delete(key string)(bool,error){
	if m==nil{
		return false,err1
	}
	if _,ok:=m[key];!ok{
		return false ,err2
	}else{
		return true,nil
	}
}
func main(){

}