package main

import (
	"io"
	"os"

	"encoding/json"
)

// Room represents a location on the map.
type Room struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"` // unused in this implimentation, but i plan to reuse this scheme for rooms that might need
	From        []int  `json:"from,omitempty"`
	To          []int  `json:"to,omitempty"`
}

type Rooms struct {
	Room []Room `json:"rooms"`
}

// Group represents a common Grouping of rooms
type Group struct {
	Name        string `json:"name"`
	Description string `json:"description"` // unused
	Rooms       []int  `json:"rooms"`
}

type Groups struct {
	Group []Group `json:"groups"`
}

func loadFiles() ([]Room, []Group) {
	r, g := make([]Room, 0), make([]Group, 0)
	rfile, err := os.Open("./json/rooms.json")
	if err != nil {
		panic(err)
	}
	gfile, err := os.Open("./json/groups.json")
	if err != nil {
		panic(err)
	}

	rdata, err := io.ReadAll(rfile)
	if err != nil {
		panic(err)
	}
	gdata, err := io.ReadAll(gfile)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(rdata, &r)
	json.Unmarshal(gdata, &g)

	return r, g
}
