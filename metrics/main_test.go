package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Uint64Ptr(v uint64) *uint64 {
	return &v
}

func Test_buildMTree(t *testing.T) {
	type args struct {
		parent    *node
		nameParts []string
		value     uint64
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{
			name: "",
			args: args{
				parent:    &node{},
				nameParts: []string{"x", "y", "b", "c", "d"},
				value:     42,
			},
			want: &node{
				name:  "x",
				level: 0,
				value: nil,
				children: map[string]*node{
					"y": {
						name:  "y",
						level: 1,
						value: nil,
						children: map[string]*node{
							"b": {
								name:  "b",
								level: 2,
								value: nil,
								children: map[string]*node{
									"c": {
										name:  "c",
										level: 3,
										value: nil,
										children: map[string]*node{
											"d": {
												name:     "d",
												level:    4,
												value:    Uint64Ptr(42),
												children: nil,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res := buildMetricsTree(tt.args.parent, tt.args.nameParts, tt.args.value)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestSafeMetrics_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		m       *SafeMetrics
		want    []byte
		wantErr bool
	}{
		{
			name: "",
			m: &SafeMetrics{
				Data: map[MKey]uint64{
					"test": 1,
				},
				PrettyPrint: false,
			},
			want:    []byte(`{"test":1}`),
			wantErr: false,
		},
		{
			name: "",
			m: &SafeMetrics{
				Data: map[MKey]uint64{
					"test.alpha": 12,
					"test.bb":    3,
					"test.beta":  32,
				},
				PrettyPrint: true,
			},
			want:    []byte(`{"test":{"alpha":{"value":12},"bb":{"value":3},"beta":{"value":32}}}`),
			wantErr: false,
		},
		{
			name: "",
			m: &SafeMetrics{
				Data: map[MKey]uint64{
					"test.alpha":     12,
					"test.bb":        3,
					"test.beta":      32,
					"omega.test1":    55,
					"omega.test1.bp": 56,
					"omega.test1.db": 85,
					"omega.test1.hp": 59,
				},
				PrettyPrint: true,
			},
			want:    []byte(`{"omega":{"test1":{"bp":{"value":56},"db":{"value":85},"hp":{"value":59},"value":55}},"test":{"alpha":{"value":12},"bb":{"value":3},"beta":{"value":32}}}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("SafeMetrics.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
