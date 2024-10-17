package ctx

import (
	"context"
	"fmt"
	"time"
)


func RunContext(){

	parentCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	go perfomBgTask(parentCtx)

	select {
	case <-parentCtx.Done():
		fmt.Println("Parent Task timed out")
}

fmt.Println("Returning from RunContext Func")
}

func perfomBgTask(ctx context.Context){
	childctx, cancel := context.WithCancel(ctx)
	defer cancel();

	go childTask(childctx)

	select {
	case <-childctx.Done():
		fmt.Println("Child function cancelled")
	}
}

func childTask(ctx context.Context){
	
}