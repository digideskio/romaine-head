package head

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/danopia/romaine-head/common"
	"github.com/danopia/romaine-head/ddp"
	"github.com/kr/pty"
	"github.com/kr/text"
)

const headPath = "~/Downloads/romaine-head.run"

var Leaves = make(map[string]*Leaf)

func GetLeaf(leaf string) (val *Leaf, ok bool) {
	val, ok = Leaves[leaf]
	return
}

var port = 6206

func StartLeaf(leaf string) *Leaf {
	if entry, ok := GetLeaf(leaf); ok && entry.Anchor != nil {
		return entry
	}

	secret := common.GenerateSecret()

	log.Printf("Starting %s under port %d", leaf, port)
	command := fmt.Sprintf("%s -- --mode leaf --port %d --secret %s 2>&1", headPath, port, secret)
	port++

	prefix := []byte(fmt.Sprintf("[%s] ", leaf))
	output := text.NewIndentWriter(os.Stdout, prefix)

	entry := &Leaf{
		Id:     leaf,
		State:  "launching",
		Secret: secret,
		Source: make(chan common.Packet),
		Sink:   make(chan common.Packet),
		Anchor: exec.Command("enter-chroot", "-n", leaf, "sh", "-c", command),
	}
	ddp.Chroots.SetField(leaf, "status", entry.State)

	_, err := entry.Anchor.StdinPipe()
	if err != nil {
		log.Printf("Leaf %s rejected stdin pipe. %+v", leaf, err)
	}

	f, err := pty.Start(entry.Anchor)
	if err != nil {
		// TODO: enough cleanup?
		entry.State = "crashed"
		entry.Anchor = nil
		ddp.Chroots.SetField(leaf, "status", entry.State)
	} else {
		entry.Pty = f
		io.Copy(output, entry.Pty)
	}

	go func() {
		err := entry.Anchor.Wait()
		log.Printf("Leaf %s exited with %+v", leaf, err)

		if err != nil {
			entry.State = "crashed"
		} else {
			entry.State = "stopped"
		}
		entry.Anchor = nil

		ddp.Chroots.SetField(leaf, "status", entry.State)
	}()

	Leaves[leaf] = entry
	return entry
}
