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
		},

		HTML: hxhtml.Div(
			hxprimitives.AttrClass(
				"appointment-container",
			),

			hxprimitives.AttrCSS(
				`display: flex; flex-wrap: nowrap; gap: 0;`,
			),

			nodesInputDate.HTML,

			WidgetSlots(&params.ParamsWidgetSlots),

			hxcomponents.ButtonSubmit(
				&params.ParamsButtonSubmit,
			),
		),
	}
}
