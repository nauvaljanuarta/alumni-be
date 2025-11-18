package main

import (
	"testing"
)

// Import semua test packages (jika beda folder)
func TestMainServices(t *testing.T) {
	t.Run("AlumniServiceTests", func(t *testing.T) {
		// semua *_test.go dalam package ini otomatis ke-run
	})

	t.Run("FileServiceTests", func(t *testing.T) {
		// semua *_test.go dalam package ini otomatis ke-run
	})

	t.Run("PekerjaanServiceTests", func(t *testing.T) {
		// semua *_test.go dalam package ini otomatis ke-run
	})
}
