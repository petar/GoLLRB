package main

import (
	"log"

	"github.com/petar/GoLLRB/llrb"
)

// SampleData will be sorted by Text field
type SampleData struct {
	Text string
	Data []byte
}

// Less method, implements llrb.Item interface for SampleData structure
func (a *SampleData) Less(b llrb.Item) bool { return a.Text < b.(*SampleData).Text }

func main() {
	tree := llrb.New()

	// Inserting sample data
	neo := &SampleData{Text: "Neo", Data: []byte("The Matrix")}
	bob := &SampleData{Text: "Bob", Data: []byte{42}}
	tree.ReplaceOrInsert(neo)
	tree.ReplaceOrInsert(bob)
	tree.ReplaceOrInsert(&SampleData{Text: "Alice", Data: []byte("Please, delete me!")})
	tree.ReplaceOrInsert(&SampleData{Text: "Boris", Data: []byte(bob.Text)})
	tree.ReplaceOrInsert(&SampleData{Text: "Petar", Data: []byte{'G', 'o', 0x4c, 0x4c, 0x52, 0x42}})

	tree.DeleteMin() // Alice will be deleted
	tree.Delete(neo) // Neo is deleted

	// Filtering SampleData items so that
	// SampleData.Text should start with "Bo"
	filter := &SampleData{Text: "Bo", Data: nil}
	tree.AscendGreaterOrEqual(filter, func(i llrb.Item) bool {
		s := i.(*SampleData)
		log.Printf("%q: %q", s.Text, s.Data)
		return true
	})
}
