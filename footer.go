package widgets

import (
	hxcomponents "github.com/TudorHulban/hx-core/components"
	hxhtml "github.com/TudorHulban/hx-core/html"
	pagecss "github.com/TudorHulban/hx-core/page-css"
	hxprimitives "github.com/TudorHulban/hx-core/primitives"
)

func WidgetFooter() hxprimitives.Node {
	form := hxcomponents.NewFormThreeContainers(
		&hxcomponents.ParamsNewFormThreeContainers{
			IDDivEnclosing:  "footer-info",
			IDForm:          "form-footer",
			IDDivContainers: "form-footer-info",

			ElementsInputLeft: []hxprimitives.Node{
				hxhtml.Span(
					hxprimitives.Text(
						"Contact and Open Hours",
					),
				),
				hxhtml.Span(
					hxprimitives.Text(
						"+1234567890",
					),
				),
			},

			ElementsInputMiddle: []hxprimitives.Node{
				hxhtml.Span(
					hxprimitives.Text(
						"Address",
					),
				),
				hxhtml.Span(
					hxprimitives.Text(
						"Lorem ipsum 3",
					),
				),
			},

			ElementsInputRight: []hxprimitives.Node{
				hxhtml.Span(
					hxprimitives.Text(
						"Book an appointment",
					),
				),
			},
		},
	)

	return hxhtml.Div(
		hxprimitives.AttrClass(
			"page-footer",
		),

		form,

		hxhtml.Div(
			hxprimitives.AttrID(
				"copyright",
			),

			hxhtml.H4(
				hxprimitives.Text(
					"Barbershop ©2025",
				),
			),
		),
	)
}

func CSSWidgetFooter() *pagecss.CSSElement {
	return &pagecss.CSSElement{
		CSSAllMedias: `
		.page-footer {
    		background-color: #252525;

		span {
			display: block;
			color: #eeeeee;
		};	

		#form-footer-info {
			display: flex;
		};
		#form-footer {padding: 10px;};
		#copyright {
			text-align: center;

			h4 {color: #eeeeee;};
		};
		}
		`,

		CSSResponsive: []pagecss.CSSMedia{
			{
				InflexionPointPX: _Tablet,
				CSS: `
				#form-footer-info {
				flex-direction: column;
				};
				`,
			},
		},
	}
}
