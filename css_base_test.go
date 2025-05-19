package widgets

import (
	"testing"

	pagecss "github.com/TudorHulban/hx-core/page-css"
	"github.com/stretchr/testify/require"
)

func TestBase(t *testing.T) {
	cssPage := pagecss.NewCSSPage(
		CSSBase,
	)

	writerCSS, errWriterCSS := getFileWriter("generated_base.css")
	require.NoError(t, errWriterCSS)

	defer writerCSS.Close()

	cssPage.GetCSSAccurateBeautifiedTo(
		writerCSS,
		&pagecss.ParamsSpaces{
			NumberSpaces:              5,
			IncrementWithNumberSpaces: 2,
		},
	)

}
