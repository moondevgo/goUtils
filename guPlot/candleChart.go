package guPlot

import (
	// "math"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
	"github.com/MetalBlueberry/go-plotly/offline"
)

func PlotCandles(opens, highs, lows, closes []float64) {

	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Candlestick{
				Type:  grob.TraceTypeCandlestick,
				Open:  opens,
				High:  highs,
				Low:   lows,
				Close: closes,
			},
		},
	}
	// offline.ToHtml(fig, "scatter.html")
	offline.Show(fig)
}
