package main

import (
	"fmt"
	"os"

	"github.com/curtisnewbie/markdown-text-renderer/renderer"
	"github.com/curtisnewbie/miso/miso"
	"github.com/curtisnewbie/miso/util/flags"
	"github.com/curtisnewbie/miso/util/must"
	"github.com/curtisnewbie/miso/util/strutil"
)

var (
	DebugFlag = flags.Bool("debug", false, "Enable debug log", false)
	FileFlag  = flags.String("file", "", "File", true)
)

func main() {
	flags.Parse()

	if *DebugFlag {
		miso.SetLogLevel("debug")
	}

	buf, err := os.ReadFile(*FileFlag)
	must.Must(err)
	s, err := renderer.RenderMarkdownText(strutil.UnsafeByt2Str(buf))
	must.Must(err)
	fmt.Printf("\n\n%v\n\n", s)
}
