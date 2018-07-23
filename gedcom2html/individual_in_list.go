package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

// individualInList is a single row in the table of individuals on the list
// page.
type individualInList struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
}

func newIndividualInList(document *gedcom.Document, individual *gedcom.IndividualNode) *individualInList {
	return &individualInList{
		individual: individual,
		document:   document,
	}
}

func (c *individualInList) String() string {
	birthDate, birthPlace := getBirth(c.individual)
	deathDate, deathPlace := getBirth(c.individual)

	birthPlace = prettyPlaceName(birthPlace)
	deathPlace = prettyPlaceName(deathPlace)

	return fmt.Sprintf(fmt.Sprintf(`
		<tr>
			<td nowrap="nowrap">%s</td>
			<td nowrap="nowrap">%s</td>
			<td>%s</td>
			<td nowrap="nowrap">%s</td>
			<td>%s</td>
		</tr>`,
		newIndividualLink(c.document, c.individual),
		birthDate, birthPlace, deathDate, deathPlace))
}
