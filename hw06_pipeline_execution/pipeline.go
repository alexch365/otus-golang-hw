package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return in
	}

	outStream := make(Bi, len(in))
	if in == nil {
		close(outStream)
		return outStream
	}

	pipeline := stages[0](in)
	for _, stage := range stages[1:] {
		pipeline = stage(pipeline)
	}

	go func() {
		defer close(outStream)
		for {
			select {
			case <-done:
				return
			case pipelineValue, ok := <-pipeline:
				if !ok {
					return
				}
				outStream <- pipelineValue
			}
		}
	}()

	return outStream
}
