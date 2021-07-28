package trace

import(
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("New should not return nil")
	} else {
		tracer.Trace("Hello, Trace!")
		if buf.String() != "Hello, Trace!\n" {
			t.Errorf("Unexpected output from Trace: '%s'", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	silentTracer := Off()
	silentTracer.Trace("YOU SHOULD NEVER SEE THIS!!!")
}