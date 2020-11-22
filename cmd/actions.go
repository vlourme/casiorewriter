package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// FetchSlice will create a slice from user defined offsets
func (m *CasioRewriter) FetchSlice(data []byte) []byte {
	// Get slice
	fmt.Println("[Info] Slicing bytes")
	slice := data[m.In:m.Out]

	// Parse line breaks
	fmt.Println("[Info] Fixing charset")
	formatted := m.ParseBreaks(slice, false)

	// Return formatted slice
	return formatted
}

// Apply will patch the OS
func (m *CasioRewriter) Apply(data []byte) error {
	// Get slice
	original := m.FetchSlice(data)

	// Load text
	fmt.Printf("[Info] Loading patch file (%s)\n", m.Patch)
	patch, err := ioutil.ReadFile(m.Patch)
	if err != nil {
		return err
	}
	fmt.Println("[Ok] Patch loaded")

	// Compare models
	fmt.Println("[Info] Validating patch model")
	err = m.ValidateModel(original, patch)
	if err != nil {
		return err
	}

	fmt.Println("[Ok] Patch is valid")

	// Replace in OS
	fmt.Println("[Info] Rewriting bytes...")
	os := bytes.ReplaceAll(data, data[m.In:m.Out], m.ParseBreaks(patch, true))

	// Store OS
	fmt.Printf("[Info] Writing new OS file (%s)\n", m.Output)
	return ioutil.WriteFile(m.Output, os, 0664)
}