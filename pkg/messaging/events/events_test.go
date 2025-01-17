package events

import (
	"testing"
	"time"

	"github.com/farshidtz/senml/v2"
	"github.com/matryer/is"
)

func TestGetValuesFromPack(t *testing.T) {
	is := testSetup(t)
	var v float64 = 1.0
	var b bool = true

	dt := time.Date(2022, time.January, 1, 12, 0, 0, 0, time.UTC)

	evt := NewMessageAccepted("sensor", senml.Pack{}, Rec("withValues", "str", &v, &b, float64(dt.Unix()), nil))

	b, ok := evt.GetBool("withValues")
	is.True(ok)
	v, ok = evt.GetFloat64("withValues")
	is.True(ok)
	str, ok := evt.GetString("withValues")
	is.True(ok)
	date, ok := evt.GetTime("withValues")
	is.True(ok)

	is.True(b)
	is.Equal(v, 1.0)
	is.Equal(str, "str")
	is.Equal(float64(dt.Unix()), date)
}

func TestNilValues(t *testing.T) {
	is := testSetup(t)

	evt := NewMessageAccepted("sensor", senml.Pack{}, Rec("nil", "", nil, nil, 0, nil))
	v, ok := evt.GetFloat64("nil")
	is.True(!ok)
	s, ok := evt.GetString("nil")
	is.True(ok)
	b, ok := evt.GetBool("nil")
	is.True(!ok)

	is.Equal(v, 0.0)
	is.Equal(s, "")
	is.True(!b)
}

func TestGetValuesFromPack2(t *testing.T) {
	is := testSetup(t)
	var v float64 = 1.0
	var b bool = true
	dt := time.Date(2022, time.January, 1, 12, 0, 0, 0, time.UTC)
	baseRec := senml.Pack{
		senml.Record{
			Name: "0",
			BaseName: "basename",
		},
	}

	evt := NewMessageAccepted("sensor", baseRec, Rec("1", "str", &v, &b, float64(dt.Unix()), nil))

	f, _ := Get[float64](*evt, "basename", 1)
	is.Equal(1.0, f)
	s, _ := Get[string](*evt, "basename", 1)
	is.Equal(s, "str")
	b2, _ := Get[bool](*evt, "basename", 1)
	is.Equal(b2, true)	
}

func testSetup(t *testing.T) *is.I {
	is := is.New(t)
	return is
}
