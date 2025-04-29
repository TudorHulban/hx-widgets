package widgets

import (
	"fmt"

	hxhelpers "github.com/TudorHulban/hx-core/helpers"
	hxhtml "github.com/TudorHulban/hx-core/html"
	pagecss "github.com/TudorHulban/hx-core/page-css"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type InfoSlot struct {
	Caption string

	ResourceID int64
	SlotID     int64
}

func (slot *InfoSlot) URL() string {
	return fmt.Sprintf(
		"xxx/%d/%d",

		slot.ResourceID,
		slot.SlotID,
	)
}

type ParamsWidgetSlots struct {
	SlotsInfo     []*InfoSlot
	NumberColumns uint8
}

func WidgetSlots(params *ParamsWidgetSlots) hxprimitives.Node {
	element := func(slotURL string, caption string) hxprimitives.Node {
		return hxprimitives.Raw(
			hxhelpers.Sprintf(
				`<button class="time-slot" type="button" onclick="handletimeclick('%s')">%s</button>`,

				slotURL,
				caption,
			),
		)
	}

	rows := []hxprimitives.Node{
		hxprimitives.Raw(
			`<script>
			function handletimeclick(slot){console.log('time slot clicked:', slot);};
			</script>`,
		),
	}

	currentRow := make([]hxprimitives.Node, 0)

	for ix, slot := range params.SlotsInfo {
		currentRow = append(
			currentRow,
			element(
				slot.URL(),
				slot.Caption,
			),
		)

		if (ix+1)%int(params.NumberColumns) == 0 || ix == len(params.SlotsInfo)-1 {
			rows = append(
				rows,
				hxhtml.Div(
					append(
						[]hxprimitives.Node{
							hxprimitives.AttrClass("hours-row"),
						},

						currentRow...,
					)...,
				),
			)

			currentRow = currentRow[:0]
		}
	}

	return hxhtml.Div(
		append(
			[]hxprimitives.Node{
				hxprimitives.AttrClass("hours-grid"),
			},
			rows...,
		)...,
	)
}

func CSSWidgetSlots() *pagecss.CSSElement {
	return &pagecss.CSSElement{
		CSSAllMedias: `
		.hours-grid {
    	width: 40%;
		}

		.hours-row {
    	display: flex;
		}

		.time-slot {
    	flex-grow: 1;
		}
		`,
	}
}
