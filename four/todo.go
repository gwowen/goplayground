package main

// This is our model!

import "time"

type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

type Todos []Todo
