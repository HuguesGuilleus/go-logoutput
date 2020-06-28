// Copyright (c) 2020, Hugues GUILLEUS <ghugues@netc.fr>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logoutput

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const day time.Duration = time.Hour * 24

// A logger output. You can use Zero value.
type W struct {
	// The directory taht contains all log file.
	Root  string
	today time.Time
	f     *os.File
}

// Create a W writer
func New(root string) *W {
	w := &W{Root: root}
	w.open()
	return w
}

// Set the output of the standard log.
func SetLog(root string) {
	log.SetOutput(New(root))
}

// Write the content of data to the day file and into os.Stdout
func (w *W) Write(data []byte) (int, error) {
	os.Stdout.Write(data)
	w.open()
	return w.f.Write(data)
}

// Check the date of the opened file, and open it if necessary.
func (w *W) open() {
	now := time.Now().Truncate(day)
	if now.Equal(w.today) && w.f != nil {
		return
	}

	w.today = now
	if w.f != nil {
		w.f.Close()
	}

	name := filepath.Join(w.Root, w.today.Format("2006-01-02.log"))
	os.MkdirAll(filepath.Dir(name), 0777)
	var err error
	w.f, err = os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
