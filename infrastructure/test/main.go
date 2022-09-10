package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/kmaguswira/coinbit/proto"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
	"github.com/lovoo/goka/storage"
	"google.golang.org/protobuf/proto"
)

var (
	brokers                         = []string{"localhost:9092"}
	topic               goka.Stream = "deposits"
	balanceGroup        goka.Group  = "balance1.group"
	aboveThresholdGroup goka.Group  = "aboveThreshold.group"
)

// Emit messages forever every second
func runEmitter() {
	emitter, err := goka.NewEmitter(brokers, topic, new(codec.String))
	if err != nil {
		log.Fatalf("error creating emitter: %v", err)
	}
	defer emitter.Finish()
	for {
		time.Sleep(1 * time.Second)
		err = emitter.EmitSync("some-key", "some-value")
		if err != nil {
			log.Fatalf("error emitting message: %v", err)
		}
	}
}

// process messages until ctrl-c is pressed
func runProcessor() {
	// process callback is invoked for each message delivered from
	// "example-stream" topic.
	cb := func(ctx goka.Context, msg interface{}) {
		var counter int64
		// ctx.Value() gets from the group table the value that is stored for
		// the message's key.
		if val := ctx.Value(); val != nil {
			counter = val.(int64)
		}
		counter++

		event := &pb.DepositAmount{}
		stringMsg, _ := msg.(string)
		// fmt.Println(err)

		if err := proto.Unmarshal([]byte(stringMsg), event); err != nil {
			log.Fatalln("Failed to parse address book:", err)
		}
		// SetValue stores the incremented counter in the group table for in
		// the message's key.

		local := storage.New()

		ctx.SetValue(counter)
		log.Printf("key = %s, counter = %v, msg = %v", ctx.Key(), counter, event)
	}

	// Define a new processor group. The group defines all inputs, outputs, and
	// serialization formats. The group-table topic is "example-group-table".
	g := goka.DefineGroup(balanceGroup,
		goka.Input(topic, new(codec.String), cb),
		goka.Persist(new(codec.Int64)),
	)

	p, err := goka.NewProcessor(brokers, g)
	if err != nil {
		log.Fatalf("error creating processor: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)
	go func() {
		defer close(done)
		if err = p.Run(ctx); err != nil {
			log.Fatalf("error running processor: %v", err)
		} else {
			log.Printf("Processor shutdown cleanly")
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<-wait   // wait for SIGINT/SIGTERM
	cancel() // gracefully stop processor
	<-done
}

func main() {
	// go runEmitter() // emits one message and stops
	runProcessor() // press ctrl-c to stop

	// wait := make(chan os.Signal, 1)
	// signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	// <-wait // wait for SIGINT/SIGTERM
	// cancel() // gracefully stop processor
	// <-done
}
