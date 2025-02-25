package leaderelection

type SimpleListener struct {
	do         func(cancelChan <-chan struct{})
	isRunning  bool
	cancelChan chan struct{}
}

func NewListener(do func(cancelChan <-chan struct{})) Listener {
	return &SimpleListener{
		do:         do,
		isRunning:  false,
		cancelChan: make(chan struct{}),
	}
}

func (s *SimpleListener) Start() {
	if !s.isRunning {
		s.isRunning = true
		s.do(s.cancelChan)
	}
}

func (s *SimpleListener) Stop() {
	if s.isRunning {
		s.isRunning = false
		s.cancelChan <- struct{}{}
	}
}

func (s *SimpleListener) IsRunning() bool {
    return s.isRunning
}