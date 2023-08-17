package main

func main() {
	wg := new(sync.waitGroup)
	i:=1
	go func(){
		wg.Add(1)
		for i:=1;i<=100;i++{
			
		}
	}
}
