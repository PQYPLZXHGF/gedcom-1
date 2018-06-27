package gedcom_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/elliotchance/gedcom"
)

var documentTests = []struct {
	doc         *gedcom.Document
	individuals []*gedcom.IndividualNode
}{
	{
		doc:         &gedcom.Document{},
		individuals: []*gedcom.IndividualNode{},
	},
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
				}),
			},
		},
		individuals: []*gedcom.IndividualNode{
			gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
			}),
		},
	},
	{
		doc: &gedcom.Document{
			Nodes: []gedcom.Node{
				gedcom.NewIndividualNode("", "P1", []gedcom.Node{
					gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
				}),
				gedcom.NewSimpleNode(gedcom.Version, "", "", []gedcom.Node{}),
				gedcom.NewIndividualNode("", "P2", []gedcom.Node{
					gedcom.NewNameNode("John /Doe/", "", []gedcom.Node{}),
				}),
			},
		},
		individuals: []*gedcom.IndividualNode{
			gedcom.NewIndividualNode("", "P1", []gedcom.Node{
				gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{}),
			}),
			gedcom.NewIndividualNode("", "P2", []gedcom.Node{
				gedcom.NewNameNode("John /Doe/", "", []gedcom.Node{}),
			}),
		},
	},
}

func TestDocument_Individuals(t *testing.T) {
	for _, test := range documentTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.doc.Individuals(), test.individuals)
		})
	}
}
