package querier

import (
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/tsdb/chunkenc"

	"github.com/muhammadn/cortex/pkg/chunk"
	"github.com/muhammadn/cortex/pkg/querier/series"
	"github.com/muhammadn/cortex/pkg/util"
)

func mergeChunks(chunks []chunk.Chunk, from, through model.Time) chunkenc.Iterator {
	samples := make([][]model.SamplePair, 0, len(chunks))
	for _, c := range chunks {
		ss, err := c.Samples(from, through)
		if err != nil {
			return series.NewErrIterator(err)
		}

		samples = append(samples, ss)
	}

	merged := util.MergeNSampleSets(samples...)
	return series.NewConcreteSeriesIterator(series.NewConcreteSeries(nil, merged))
}
