package app

// Version returns the build version string (may be set via -ldflags).
func Version() string {
	return version
}
