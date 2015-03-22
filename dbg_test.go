/* A simple Debug/trace facility for Go programs
  License: MIT
  (c) 2015 Vince Hodges <vhodges@gmail.com>
 */
package dbg

import "testing"

// Example dbg flags - Upto 64 of them
const (
    _ = iota  // ignore first value by assigning to blank identifier
    NETWORK int64 = 1 << (1*iota)  
    COMMAND_LINE_FLAGS
    PACKET_READER
    PACKET_PROCESSING
    PERSISTENCE
    JOB_HANDLER
    ETC
)

// Setting dbg.Flags.  Typically you would do this by reading a command line flag or
// by providing a mechanism to set the value at runtime.
func ExampleFlags() {

	// Just || the flags together you want enabled.
	Flags =  -1 // All dbgs enabled (all bits are set)

	// Only traces with these flags would be displayed.
	Flags = PACKET_PROCESSING | PACKET_READER | PERSISTENCE // 28

	Flags = JOB_HANDLER | ETC  // 96
}

// Changing the Logger settings
func ExampleLoggers() {
	// TODO Write examples - The default will work for all my use cases so far.
	// but something like: Logger  = log.New(...)
	// See http://golang.org/pkg/log/ for details
}

func ExampleDbgs() {

	Printf(PACKET_PROCESSING, "A trace about packet processing")
	Printf(PERSISTENCE, "A trace about persistence")
	Printf(JOB_HANDLER, "A trace from the job handler")

	Printf(ETC, "And so on...")

	var packet []byte
	HexDump(PACKET_READER, packet)
}
