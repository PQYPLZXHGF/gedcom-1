package main

import (
	"github.com/elliotchance/gedcom"
)

const (
	selectedIndividualsTab = "individuals"
	selectedPlacesTab      = "places"
	selectedExtraTab       = "extra"
)

// header is the tabbed section at the top of each page. The header will be the
// same on all pages except that some pages will use an extra tab for that page.
type header struct {
	document              *gedcom.Document
	extraTab, selectedTab string
}

func newHeader(document *gedcom.Document, extraTab string, selectedTab string) *header {
	return &header{
		document:    document,
		extraTab:    extraTab,
		selectedTab: selectedTab,
	}
}

func (c *header) String() string {
	items := []*navItem{
		newNavItem(
			"People "+newCountBadge(len(c.document.Individuals())).String(),
			c.selectedTab == selectedIndividualsTab,
			pageIndividuals(),
		),
		newNavItem(
			"Places "+newCountBadge(len(c.document.Places())).String(),
			c.selectedTab == selectedPlacesTab,
			pagePlaces(),
		),
	}

	if c.extraTab != "" {
		items = append(items, newNavItem(
			c.extraTab,
			c.selectedTab == selectedExtraTab,
			"#",
		))
	}

	return newComponents(
		newSpace(),
		newNavTabs(items),
		newSpace(),
	).String()
}
