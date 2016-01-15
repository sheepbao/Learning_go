package main

import (
	"bytes"
	"fmt"
	"encoding/gob"
	"log"
)

type P struct {
	X,Y,Z int
	Name string
}

type Q struct {
	X,Y *int32
	Name string
}

func main () {
	// Initialize the encoder and decoder.  Normally enc and dec would be      
	// bound to network connections and the encoder and decoder would      
	// run in different processes.
	var network bytes.Buffer
	enc:=gob.NewEncoder(&network)
	dec:=gob.NewDecoder(&network)
	err:=enc.Encode(P{3,4,5,"pythagoras"})
	if err != nil {
		log.Fatal("encode error:",err)
	}

	var q Q
	err =dec.Decode(&q)
	if err !=nil {
		log.Fatal("decode error:",err)
	}
	fmt.Printf("%q:{%d,%d}\n",q.Name,*q.X,*q.Y)
}

