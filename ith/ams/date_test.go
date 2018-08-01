package ams

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAmsDate_MarshalJSON(t *testing.T) {
	var a *AmsDate
	println("Empty:", a.Empty())
	b, e := json.Marshal(a)
	assert.Equal(t, nil, e)
	println(string(b))

	x := AmsDate(time.Now())
	a = &x
	b, e = json.Marshal(a)
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

func TestAmsDateFromInt(t *testing.T) {
	var x int64
	cx := time.Now().UTC()
	x = cx.Unix()
	ad := DateFromInt(x)
	assert.Equal(t, ad.Time().Unix(), x)
}

func TestAmsDate_UnmarshalJSON(t *testing.T) {
	var a AmsDate
	e := json.Unmarshal([]byte(`"19810509000000"`), &a)
	println(a.Time().String())
	assert.Equal(t, nil, e)
	assert.Equal(t, "19810509000000", a.String())
}

func Test_OmitemptyMarshalJSON(t *testing.T) {
	a := struct {
		BirthDate *AmsDate `json:"birthDate,omitempty"`
		Id        int64    `json:"id"`
	}{Id: 1}
	b, e := json.MarshalIndent(&a, "", "  ")
	assert.NoError(t, e)
	assert.NotEmpty(t, b)
	println(string(b))
}
