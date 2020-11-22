package cmd

import (
	"bytes"
	"errors"
	"fmt"
)

// ParseBreaks will parse every line breaks then return the byte array
func (m *CasioRewriter) ParseBreaks(slice []byte, reverse bool) []byte {
	// Trim before (in reverse case)
	slice = bytes.TrimSpace(slice)

	// Iterate array
	for i := 0; i < len(slice); i++ {
		// Check for reverse
		if reverse {
			// Replace every '\n' (line break in ASCII)
			if slice[i] == '\n' {
				slice[i] = 0
			}
		} else {
			// Replace every 0 (line break in Casio charset)
			if slice[i] == 0 {
				slice[i] = '\n'
			}
		}
	}

	// Trim after
	return bytes.TrimSpace(slice)
}

// GetModel will generate a type-model from slice
func (m *CasioRewriter) GetModel(slice []byte, limit int) ([]byte, int) {
	// Duplicate
	slice = append([]byte(nil), slice...)

	// Get limit
	if limit > 0 {
		slice = slice[0:limit]
	}

	// Get writable characters
	letters := 0

	// Iterate array
	for i := 0; i < len(slice); i++ {
		// Replace every letters
		if slice[i] >= 32 && slice[i] <= 126 {
			slice[i] = 'l'
			letters++
		}

		// Replace every '\n' by 'b'
		if slice[i] == '\n' {
			slice[i] = 'b'
		}
	}

	return slice, letters
}


// ValidateModel will compare if original slice have the same format than new slice
func (m *CasioRewriter) ValidateModel(original []byte, new []byte) error {
	// Trim slices
	original = bytes.TrimSpace(original)
	new = bytes.TrimSpace(new)

	// Log
	fmt.Printf("[Info] Used size: %d/%d\n", len(new), len(original))

	// Compare length
	if len(new) > len(original) {
		return errors.New("patch is bigger than allowed size")
	}

	// Compare model
	originalModel, originalChars := m.GetModel(original, 0)
	newModel, _ := m.GetModel(new, len(original))
	fmt.Printf("[Info] Amount of writable chars in the model: %d\n", originalChars)

	if bytes.Compare(originalModel, newModel) != 0 {
		return errors.New("models are not same")
	}

	return nil
}