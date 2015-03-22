/* A simple Debug/trace facility for Go programs
 * It's lineage can be traced back to a set of C preprocessor macros that we used at Dynabase Systems
 * back in the early 90's. 
 *
 * This is better (imo) than debug/log levels because it allows instrumentation to remain in place
 * and be selectively turned on or off as needed.
 *
 * License: MIT
 *
 * (c) 2015 Vince Hodges <vhodges@gmail.com>
 */
package dbg

import (
	"encoding/hex"
	"log"
	"os"
)

// Flags is a bit mask representing the set of traces that are enabled.
var Flags int64 

// Logger is used to write debug messages
var Logger *log.Logger

// init makes sure there's a default logger to write to
func init() {
	// Log instead of fmt so we get datestamped entries.
	Logger = log.New(os.Stdout, "[Debug] ", log.LstdFlags)
}

// Printf will check flag against Flags to see if the bit is set and 
// will print the trace if so
func Printf(flag int64, format string, args ...interface{} ) {
	if (Flags != 0 && Flags & flag == flag) {
		Logger.Printf(format, args...)
	}
}

// HexDump formats a HexDump of data and prints it to Logger if Flags has the
// bit set for flag. 
func HexDump(flag int64, data []byte) {
	if (Flags != 0 && Flags & flag == flag) {
		Logger.Printf("Hex:\n%s", hex.Dump(data))
	}
}

// TODO Add helpers for printing/parsing command line string into debug flag
