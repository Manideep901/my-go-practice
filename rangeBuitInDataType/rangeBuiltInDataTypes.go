package main
import ("fmt")

func main(){
	for idx,val := range "manideep"{
		fmt.Println(idx,val)
	}
	nums := []int{1,2,3,4,5,6}
	sum:= 0;
	for _,num := range nums{
		sum += num
	}
	fmt.Printf("%v",sum)
	fmt.Println()
	for i, num := range nums {
		if num == 3 {
				fmt.Println("index:", i)
		}
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
		for k := range kvs {
			fmt.Println("key:", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
}
}