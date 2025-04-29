package widgets

import (
	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	pagecss "github.com/TudorHulban/hx-core/page-css"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

type ParamsWidgetAppointment struct {
	ParamsWidgetSlots
	ParamsWidgetInputDate

	hxcomponents.ParamsButtonSubmit
}

type ResponseWidgetAppointment struct {
	HTML           hxprimitives.Node
	LinkJavascript hxprimitives.Node

	CSS []func() *pagecss.CSSElement
}

func WidgetAppointment(params *ParamsWidgetAppointment) *ResponseWidgetAppointment {
	nodesInputDate := WidgetInputDate(
		&params.ParamsWidgetInputDate,
	)

	return &ResponseWidgetAppointment{
		LinkJavascript: nodesInputDate.LinkJavascript,
		CSS: []func() *pagecss.CSSElement{
			CSSInputDate,
			CSSWidgetSlots,
			CSSAppointment,
		},

		HTML: hxhtml.Div(
			hxprimitives.AttrClass(
				"appointment-container",
			),

			hxhtml.Div(
				hxprimitives.AttrCSS(
					`display: flex; flex-wrap: nowrap; gap: 0.2em;`,
				),

				nodesInputDate.HTML,

				WidgetSlots(
					&params.ParamsWidgetSlots,
				),
			),

			hxcomponents.ButtonSubmit(
				&params.ParamsButtonSubmit,
			),
		),
	}
}

func CSSAppointment() *pagecss.CSSElement {
	return &pagecss.CSSElement{
		CSSAllMedias: `
		.appointment-container {
			padding: 0.3em;
			text-align: right;
			width: fit-content;
			background-color:rgb(134, 146, 138);

			.hours-grid {
				width: 100%;
				padding-top: 2.1em;
			}
		}
		`,
	}
}
