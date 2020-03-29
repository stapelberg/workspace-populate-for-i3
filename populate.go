// Program workspace-populate-for-i3 restores a 50/50 split layout and starts 2
// urxvt terminals when a new workspace is created.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go.i3wm.org/i3/v4"
)

func logic() error {
	layoutPath := filepath.Join(os.Getenv("HOME"), ".config", "i3", "urxvt-50-50.json")
	recv := i3.Subscribe(i3.WorkspaceEventType)
	for recv.Next() {
		ev := recv.Event().(*i3.WorkspaceEvent)
		if ev.Change != "init" {
			continue
		}
		_, err := i3.RunCommand(fmt.Sprintf("append_layout %s", layoutPath))
		if err != nil {
			log.Printf("append_layout: %v", err)
			continue
		}
		for i := 0; i < 2; i++ {
			_, err := i3.RunCommand(fmt.Sprintf("[con_id=%d] exec exec urxvtc -title 'midna: ~'", ev.Current.ID))
			if err != nil {
				log.Printf("exec urxvtc: %v", err)
			}
		}
	}
	return recv.Close()
}

func main() {
	if err := logic(); err != nil {
		log.Fatal(err)
	}
}
