package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	I   = interface{}
	In  = <-chan I
	Out = In
	Bi  = chan I
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	outStream := make(Bi, len(in))

	if len(stages) == 0 {
		return in
	}

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
				select {
				case <-done:
					return
				case outStream <- pipelineValue:
				}
			}
		}
	}()

	return outStream
}
