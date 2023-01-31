package guPlot

import (
	// "math"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
	"github.com/MetalBlueberry/go-plotly/offline"
)

type ScatterTrace struct {
	Xs    []float64 // x좌표 배열
	Ys    []float64 // y좌표 배열
	Mode  string
	Color string
	Size  float64
	// Mode   grob.ScatterMode    // "markers" "lines" "lines+markers"
	// Marker *grob.ScatterMarker // https://github.com/MetalBlueberry/go-plotly/blob/135f6aad2ff763a003f200ea59d5da9c13ac04d5/graph_objects/scatter_gen.go
}

func PlotScatters(scatterTraces []ScatterTrace) {
	scatters := []grob.Trace{}
	for _, scatterTrace := range scatterTraces {
		scatter := &grob.Scatter{
			Type:   grob.TraceTypeScatter,
			X:      scatterTrace.Xs,
			Y:      scatterTrace.Ys,
			Mode:   grob.ScatterMode(scatterTrace.Mode),
			Marker: &grob.ScatterMarker{Color: scatterTrace.Color, Size: scatterTrace.Size},
		}
		scatters = append(scatters, scatter)
	}

	fig := &grob.Fig{
		Data: grob.Traces{},
	}

	for _, scatter := range scatters {
		fig.Data = append(fig.Data, scatter)
	}

	// offline.ToHtml(fig, "scatter.html")
	offline.Show(fig)
}
