package unix

import (
	"fmt"
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"github.com/project-flogo/core/support/log"
)

func init() {
	_ = function.Register(&fntimestamp{})
	_ = function.Register(&fntimestampMilli{})
	_ = function.Register(&fntimestampMicro{})
}

type fntimestamp struct{
}

type fntimestampMilli struct{
}

type fntimestampMicro struct{
}

func (fntimestamp) Name() string {
	return "timestamp"
}

func (fntimestampMilli) Name() string {
	return "timestampMilli"
}

func (fntimestampMicro) Name() string {
	return "timestampMicro"
}

func (fntimestamp) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (fntimestampMilli) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (fntimestampMicro) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{}, false
}

func (fntimestamp) Eval(params ...interface{}) (interface{}, error) {
	var fName string = (&fntimestamp{}).Name()
	var logger = log.ChildLogger(unixLogger, fName)
	
	if len(params) > 0 {
		logger.Errorf("expected no parameters but got %d", len(params))
		return nil, fmt.Errorf("%s: expected no parameters but got %d", fName, len(params))
	}

	logger.Debugf("%s function called", fName)

	var timestamp = time.Now().UTC().Unix()
	return timestamp, nil
}

func (fntimestampMilli) Eval(params ...interface{}) (interface{}, error) {
	var fName string = (&fntimestampMilli{}).Name()
	var logger = log.ChildLogger(unixLogger, fName)

	if len(params) > 0 {
		logger.Errorf("expected no parameters but got %d", len(params))
		return nil, fmt.Errorf("%s: expected no parameters but got %d", fName, len(params))
	}

	logger.Debugf("%s function called", fName)

	var timestamp = time.Now().UTC().UnixMilli()
	return timestamp, nil
}

func (fntimestampMicro) Eval(params ...interface{}) (interface{}, error) {
	var fName string = (&fntimestampMicro{}).Name()
	var logger = log.ChildLogger(unixLogger, fName)
	
	if len(params) > 0 {
		logger.Errorf("expected no parameters but got %d", len(params))
		return nil, fmt.Errorf("%s: expected no parameters but got %d", fName, len(params))
	}

	logger.Debugf("%s function called", fName)

	var timestamp = time.Now().UTC().UnixMicro()
	return timestamp, nil
}
