package ams

import (
	"testing"
	"time"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestAmsDate_MarshalJSON(t *testing.T) {
	var a AmsDate
	b, e := json.Marshal(&a)
	assert.Equal(t, nil, e)
	println(string(b))

	a = AmsDate(time.Now())
	b, e = json.Marshal(&a)
	assert.Equal(t, nil, e)
	println(string(b))
}

func TestAmsDate_Time(t *testing.T) {
	a := AmsDate(time.Now())
	b := a.Time().String()
	assert.NotEqual(t, "", b)
}

func TestAmsDate_String(t *testing.T) {
	a := AmsDate(time.Now())
	println(a.String())
	assert.NotEqual(t, "", a.String())
}

func TestAmsDate_UnmarshalJSON(t *testing.T) {
	var a AmsDate
	e :=json.Unmarshal([]byte(`"19810509000000"`), &a)
	println(a.Time().String())
	assert.Equal(t, nil, e)
	assert.Equal(t, "19810509000000", a.String())
}
