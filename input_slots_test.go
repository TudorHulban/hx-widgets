package widgets

import (
	"fmt"
	"testing"

	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	pagecss "github.com/TudorHulban/hx-core/page-css"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
	"github.com/stretchr/testify/require"
)

func TestSlots(t *testing.T) {
	fragment := WidgetSlots(
		&ParamsWidgetSlots{
			NumberColumns: 1,

			SlotsInfo: []*InfoSlot{
				{
					ResourceID: 1,
					SlotID:     1000,
					Caption:    "10 00 - dr. John Smith",
				},
				{
					ResourceID: 2,
					SlotID:     1030,
					Caption:    "10 30 - dr. Martha Doe",
				},
				{
					ResourceID: 1,
					SlotID:     1100,
					Caption:    "11 00 - dr. John Smith",
				},
				{
					ResourceID: 2,
					SlotID:     1100,
					Caption:    "11 00 - dr. Martha Doe",
				},
			},
		},
	)

	page := hxcomponents.Page{
		Title: t.Name(),

		Head: []hxprimitives.Node{
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("generated.css"),
			),
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("https://cdn.jsdelivr.net/npm/flatpickr/dist/flatpickr.min.css"),
			),
			hxhtml.Link(
				hxprimitives.Rel("stylesheet"),
				hxprimitives.Href("https://npmcdn.com/flatpickr/dist/themes/dark.css"),
			),
		},

		Body: []hxprimitives.Node{
			fragment,
		},
	}

	writerCSS, errWriterCSS := getFileWriter("generated.css")
	require.NoError(t, errWriterCSS)

	defer writerCSS.Close()

	cssPage := pagecss.NewCSSPage(
		CSSBase,
		CSSSite,
		CSSWidgetSlots,
	)

	cssPage.GetCSSTo(writerCSS)

	writerHTML, errWriterHTML := getFileWriter(t.Name() + ".html")
	require.NoError(t, errWriterHTML)

	defer writerHTML.Close()

	fmt.Println(
		page.Build().Render(writerHTML),
	)
}
