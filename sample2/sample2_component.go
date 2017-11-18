package sample2

type Sample2Component interface {
	Exec() error
	StrExec(string) string
}

type sampleComponent struct {
}

func NewSample2Component() Sample2Component {
	return &sampleComponent{}
}

func (s *sampleComponent) Exec() error {
	return nil
}

func (s *sampleComponent) StrExec(str string) string {
	return str
}
