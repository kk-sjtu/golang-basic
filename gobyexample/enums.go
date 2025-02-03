package main

import "fmt"

type ServerState int

const(
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle: "idle",
	StateConnected: "connected",
	StateError: "error",
	StateRetrying: "retrying",
}

func (ss ServerState) String() string{
	return stateName[ss]
}

/*
func main(){
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}*/
func main() {
    ns := StateIdle.transition()
    fmt.Println(ns)

    ns2 := ns.transition()
    fmt.Println(ns2)
}
// 这样写的逻辑好像怪怪的
func (s ServerState) transition() ServerState{
	switch s{
	case StateIdle:
		return StateConnected
	case StateConnected,StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unkown state %d",s))
	}
}