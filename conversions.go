package gogenutils

import "fmt"

// from https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
// ByteCountSI returns the "human readable" number of bytes in SI units (as a string)
func ByteCountSI(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
