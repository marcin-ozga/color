package color

import (
	"reflect"
	"testing"
)

func TestColor_String(t *testing.T) {
	tests := []struct {
		name  string
		color Color
		want  string
	}{
		{
			name:  "White",
			color: White,
			want:  "#FFFFFFFF",
		},
		{
			name:  "Red",
			color: Red,
			want:  "#FFFF0000",
		},
		{
			name:  "Green",
			color: Lime,
			want:  "#FF00FF00",
		},
		{
			name:  "Blue",
			color: Blue,
			want:  "#FF0000FF",
		},
		{
			name:  "Black",
			color: Black,
			want:  "#FF000000",
		},
		{
			name:  "Aqua",
			color: Aqua,
			want:  "#FF00FFFF",
		},
		{
			name:  "Magenta",
			color: Magenta,
			want:  "#FFFF00FF",
		},
		{
			name:  "Yellow",
			color: Yellow,
			want:  "#FFFFFF00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.color.String(); got != tt.want {
				t.Errorf("Color.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_ToString(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name  string
		color Color
		args  args
		want  string
	}{
		{
			name:  "ARGB",
			color: FromARGB(0xAA, 0xBB, 0xCC, 0xDD),
			args: args{
				format: FormatARGB,
			},
			want: "#AABBCCDD",
		},
		{
			name:  "RGB",
			color: FromARGB(0xAA, 0xBB, 0xCC, 0xDD),
			args: args{
				format: FormatRGB,
			},
			want: "#BBCCDD",
		},
		{
			name:  "RGBA",
			color: FromARGB(0xAA, 0xBB, 0xCC, 0xDD),
			args: args{
				format: FormatRGBA,
			},
			want: "#BBCCDDAA",
		},
		{
			name:  "Any",
			color: FromARGB(0xAA, 0xBB, 0xCC, 0xDD),
			args: args{
				format: "%08X",
			},
			want: "AABBCCDD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.color.ToString(tt.args.format); got != tt.want {
				t.Errorf("Color.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		color   Color
		want    []byte
		wantErr bool
	}{
		{
			name:    "Gray",
			color:   Gray,
			want:    []byte(`"#FF808080"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.color.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Color.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestColor_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name      string
		color     *Color
		args      args
		wantColor Color
		wantErr   bool
	}{
		{
			name:  "#AABBCCDD",
			color: new(Color),
			args: args{
				data: []byte(`"#AABBCCDD"`),
			},
			wantColor: Color(0xAABBCCDD),
			wantErr:   false,
		},
		{
			name:  "ZeroColor",
			color: new(Color),
			args: args{
				data: []byte(`""`),
			},
			wantColor: Color(0),
			wantErr:   false,
		},
		{
			name:  "Bad syntax",
			color: new(Color),
			args: args{
				data: []byte(`"AABBCCDD"`),
			},
			wantColor: Color(0),
			wantErr:   true,
		},
		{
			name:  "Bad syntax",
			color: new(Color),
			args: args{
				data: []byte(`"qwerty"`),
			},
			wantColor: Color(0),
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.color.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Color.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*tt.color, tt.wantColor) {
				t.Errorf("Color.UnmarshalJSON() = %v, want %v", *tt.color, tt.wantColor)
			}
		})
	}
}
