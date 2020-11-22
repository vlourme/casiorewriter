package cmd

import (
	"fmt"
	"io/ioutil"
)

// CasioRewriter defines the structure of command line
type CasioRewriter struct {
	Input  string `help:"Input file (should be a .bin file)"`
	Patch  string `help:"Patched text file used for fetch and apply (default: slice.txt)"`
	Output string `help:"Output file (default: patched_os.bin)"`
	Fetch  bool   `help:"Fetch texts in the selected offsets"`
	In     int    `help:"Start offset (default: enter menu, 0x0002F82E)"`
	Out    int    `help:"Ending offset (default: enter menu, 0x0002F8CD)"`
}

// NewMain instance a new CasioRewriter instance
func NewMain() *CasioRewriter {
	return &CasioRewriter{
		Output: "patched_os.bin",
		Patch:  "slice.txt",
		Fetch:  false,
		In:     0x0002F82E,
		Out:    0x0002F8CD,
	}
}

// Run is the main entry to the program
func (m *CasioRewriter) Run() error {
	// Authoring
	fmt.Println("[+] Casio OS Rewriter made by Victor Lourme")
	fmt.Println("[+] https://github.com/vlourme/casiorewriter")

	// Offset log
	fmt.Printf("[Info] Offset size: %d bytes\n", m.Out - m.In)
	fmt.Printf("[Info] Offset entry: %#x\n", m.In)
	fmt.Printf("[Info] Offset end: %#x\n", m.Out)

	// Load binary
	fmt.Println("[Info] Loading OS binary")
	data, err := ioutil.ReadFile(m.Input)

	if err != nil {
		return err
	}

	fmt.Println("[Ok] OS binary file loaded")

	// Fetch
	if m.Fetch {
		formatted := m.FetchSlice(data)

		// Store formatted
		fmt.Println("[Info] Storing into slice.txt")
		err := ioutil.WriteFile(m.Patch, formatted, 0664)

		return err
	}

	// Apply
	return m.Apply(data)
}
