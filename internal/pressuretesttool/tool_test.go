package pressuretesttool

import (
	"testing"
	"time"
)

func TestPressureTestResult_Avg(t *testing.T) {
	type fields struct {
		timeConsumedArr []time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "测试平均值",
			fields: fields{
				timeConsumedArr: []time.Duration{
					time.Duration(100 * time.Millisecond),
					time.Duration(200 * time.Millisecond),
					time.Duration(300 * time.Millisecond),
					time.Duration(400 * time.Millisecond),
					time.Duration(500 * time.Millisecond),
					time.Duration(600 * time.Millisecond),
					time.Duration(700 * time.Millisecond),
					time.Duration(800 * time.Millisecond),
					time.Duration(900 * time.Millisecond),
					time.Duration(1000 * time.Millisecond),
				},
			},
			want: time.Duration(550 * time.Millisecond),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PressureTestResult{
				timeConsumedArr: tt.fields.timeConsumedArr,
			}
			if got := p.Avg(); got != tt.want {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPressureTestResult_Percentile(t *testing.T) {
	type fields struct {
		timeConsumedArr []time.Duration
	}
	type args struct {
		percentage float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Duration
	}{
		{
			name: "测试 95 百分位",
			fields: fields{
				timeConsumedArr: []time.Duration{
					time.Duration(100 * time.Millisecond),
					time.Duration(200 * time.Millisecond),
					time.Duration(300 * time.Millisecond),
					time.Duration(400 * time.Millisecond),
					time.Duration(500 * time.Millisecond),
					time.Duration(600 * time.Millisecond),
					time.Duration(700 * time.Millisecond),
					time.Duration(800 * time.Millisecond),
					time.Duration(900 * time.Millisecond),
					time.Duration(1000 * time.Millisecond),
				},
			},
			args: args{
				percentage: 0.95,
			},
			want: time.Duration(1000 * time.Millisecond),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PressureTestResult{
				timeConsumedArr: tt.fields.timeConsumedArr,
			}
			if got := p.Percentile(tt.args.percentage); got != tt.want {
				t.Errorf("Percentile() = %v, want %v", got, tt.want)
			}
		})
	}
}
