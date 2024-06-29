package main

import (
	"fmt"
	"math/rand"
	"github.com/ElwinCabrera/go-containers/list"
)


func main(){
	
	
	sll := list.InitSinglyLinkedList()
	var nodes []*list.Node
	//var seen_map map[int]bool
	seen_map := make(map[int]bool)
	for i := 0; i < 10; i++ {
		ran_num := rand.Intn(10)
		_, seen := seen_map[ran_num]
		for seen{
			ran_num = rand.Intn(10)
			_, seen = seen_map[ran_num]
		}
		seen_map[ran_num] = true
		
		n := sll.InsertEnd(ran_num)
		nodes = append(nodes, n)
	}
	
	fmt.Println(sll)
	clear(seen_map)

	for i := 0; i < 10; i++ {
		ran_num := rand.Intn(10)
		_, seen := seen_map[ran_num]
		for seen{
			ran_num = rand.Intn(10)
			_, seen = seen_map[ran_num]
		}
		seen_map[ran_num] = true
		fmt.Printf("Removed %v ", nodes[ran_num].Value)
		sll.Remove(nodes[ran_num])
		fmt.Print(sll)
		fmt.Print("\n")
		
	}



	
	fmt.Println("Done!")
}