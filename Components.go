package main

import(
	"time"
)

type Task[T any] struct{ //the task represents a single unit of work. Needs to know what type of work it is, it's id and the specific data it needs to process
	Type string //is it a map or a reduce task
	T_Id int //the specific Id of the task at hand.
	Data T //the data it needs to process. I used a generic because the map and reduce functios allow different data types.
}

type Master[T any] struct{//the master node that knows what work needs to be done and keeps track of which tasks are assigned and reports back when it's done.
	M_ID int //teh id of the master(naive but necessary)
	Mapper_task []*TaskState //a list of mapper tasks and their state (idle, in progress, or finished)
	Reducer_task []*TaskState // a list of reducer tasks and their state (idle, in progress, or finished)	
	Workers []int  //the ids of the workers
	Results []*ResultState //stores the location of the results
}

type Worker[T any] struct{
	W_Id int                                            
	master *Master[T] //pointer to the master                                                                                                                                                                                                      
	T_ID int //the task id that needs to be done
	Map_Fn func(string)[]KeyValue
	Reduce_Fn func(string, []string)string
}


type TaskState struct {
	T_Id int //the task id
	State string //the state of the tsk
	Assigned time.Time //the time the task was assigned to a worker. For fault tolerance reasons.
}
type ResultState struct{
	State string //intermediary or a final one
	Path []string //the paths that the results is saved at.
}