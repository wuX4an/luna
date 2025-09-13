package stages

import "context"

type Stage interface {
	Name() string
	Run(ctx context.Context) error
}

type Pipeline struct {
	stages []Stage
}

func NewPipeline(stages ...Stage) *Pipeline {
	return &Pipeline{stages: stages}
}

func (p *Pipeline) Run(ctx context.Context) error {
	for _, s := range p.stages {
		if err := s.Run(ctx); err != nil {
			return err
		}
	}
	return nil
}
