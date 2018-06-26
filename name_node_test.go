package gedcom_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/elliotchance/gedcom"
)

var nameTests = []struct {
	node          *gedcom.NameNode
	title         string
	prefix        string
	givenName     string
	surnamePrefix string
	surname       string
	suffix        string
}{
	{
		node:          gedcom.NewNameNode("", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "",
		suffix:        "",
	},
	{
		node:          gedcom.NewNameNode("/Double  Last/", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "Double Last",
		suffix:        "",
	},
	{
		node:          gedcom.NewNameNode("//", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "",
		suffix:        "",
	},
	{
		// This is an invalid case. I don't mind that the data returned seems
		// garbled. It's better than nothing.
		node:          gedcom.NewNameNode("a / b", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "a",
		surnamePrefix: "",
		surname:       "",
		suffix:        "/ b",
	},
	{
		node:          gedcom.NewNameNode("Double First", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "Double First",
		surnamePrefix: "",
		surname:       "",
		suffix:        "",
	},
	{
		node:          gedcom.NewNameNode("First /Last/", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "",
	},
	{
		node:          gedcom.NewNameNode("First   Middle /Last/", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "First Middle",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "",
	},
	{
		node:          gedcom.NewNameNode("First /Last/  Suffix ", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Suffix",
	},
	{
		node:          gedcom.NewNameNode("   /Last/ Suffix", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Suffix",
	},
	{
		// The GivenName overrides the givenName name if provided. When multiple
		// GivenNames are provided then it will always use the first one.
		node: gedcom.NewNameNode("First /Last/ II", "", []gedcom.Node{
			gedcom.NewSimpleNode(gedcom.GivenName, " Other  Name ", "", nil),
			gedcom.NewSimpleNode(gedcom.GivenName, "Uh-oh", "", nil),
		}),
		title:         "",
		prefix:        "",
		givenName:     "Other Name",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "II",
	},
	{
		// The Surname overrides the surname name if provided. When multiple
		// Surnames are provided then it will always use the first one.
		node: gedcom.NewNameNode("First /Last/ II", "", []gedcom.Node{
			gedcom.NewSimpleNode(gedcom.Surname, " Other  name ", "", nil),
			gedcom.NewSimpleNode(gedcom.Surname, "uh-oh", "", nil),
		}),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Other name",
		suffix:        "II",
	},
	{
		node: gedcom.NewNameNode("First /Last/ Esq.", "", []gedcom.Node{
			gedcom.NewSimpleNode(gedcom.NamePrefix, " Mr ", "", nil),
			gedcom.NewSimpleNode(gedcom.NamePrefix, "Dr", "", nil),
		}),
		title:         "",
		prefix:        "Mr",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Esq.",
	},
	{
		// The NameSuffix overrides the suffix in the name if provided.
		// When multiple name suffixes are provided then it will always use the
		// first one.
		node: gedcom.NewNameNode("First /Last/ Suffix", "", []gedcom.Node{
			gedcom.NewSimpleNode(gedcom.NameSuffix, " Esq. ", "", nil),
			gedcom.NewSimpleNode(gedcom.NameSuffix, "Dr", "", nil),
			gedcom.NewSimpleNode(gedcom.NamePrefix, "Sir", "", nil),
		}),
		title:         "",
		prefix:        "Sir",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Esq.",
	},
	{
		node: gedcom.NewNameNode("First /Last/ Esq.", "", []gedcom.Node{
			gedcom.NewSimpleNode(gedcom.SurnamePrefix, " Foo ", "", nil),
			gedcom.NewSimpleNode(gedcom.SurnamePrefix, "Bar", "", nil),
		}),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "Foo",
		surname:       "Last",
		suffix:        "Esq.",
	},
	{
		node: gedcom.NewNameNode("First /Last/ Esq.", "", []gedcom.Node{
			gedcom.NewSimpleNode(gedcom.Title, " Grand  Duke ", "", nil),
			gedcom.NewSimpleNode(gedcom.Title, "Nobody", "", nil),
		}),
		title:         "Grand Duke",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Esq.",
	},
}

func TestNameNode_GivenName(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.GivenName(), test.givenName)
		})
	}
}

func TestNameNode_Surname(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Surname(), test.surname)
		})
	}
}

func TestNameNode_SurnamePrefix(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.SurnamePrefix(), test.surnamePrefix)
		})
	}
}

func TestNameNode_Prefix(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Prefix(), test.prefix)
		})
	}
}

func TestNameNode_Suffix(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Suffix(), test.suffix)
		})
	}
}

func TestNameNode_Title(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Title(), test.title)
		})
	}
}
