package unix

import (
	"testing"

	"github.com/project-flogo/core/data/expression/function"
	"github.com/stretchr/testify/assert"
)

func TestFntimestamp_Eval(t *testing.T) {
	f := &fntimestamp{}
	v, err := function.Eval(f)

	if err == nil {
		t.Logf("timestamp (UTC) in seconds: %+v", v)
	} else {
		t.Logf("error: %+v",err)
	}
	assert.Nil(t, err)
	assert.IsType(t, int64(1), v)
}

func TestFntimestampMilli_Eval(t *testing.T) {
	f := &fntimestampMilli{}
	v, err := function.Eval(f)

	if err == nil {
		t.Logf("timestamp (UTC) in milliseconds: %+v", v)
	} else {
		t.Logf("error: %+v",err)
	}
	assert.Nil(t, err)
	assert.IsType(t, int64(1), v)
}

func TestFntimestampMicro_Eval(t *testing.T) {
	f := &fntimestampMicro{}
	v, err := function.Eval(f)

	if err == nil {
		t.Logf("timestamp (UTC) in microseconds: %+v", v)
	} else {
		t.Logf("error: %+v",err)
	}
	assert.Nil(t, err)
	assert.IsType(t, int64(1), v)
}
