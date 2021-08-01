package bitter

import (
	"reflect"
	"testing"
)

func TestOctetFromByte(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want *Octet
	}{
		{
			name: "simple create",
			args: args{
				b: 0b01101001,
			},
			want: &Octet{bits: 0b01101001},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OctetFromByte(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OctetFromByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctetFromQuartets(t *testing.T) {
	type args struct {
		hi *Quartet
		lo *Quartet
	}
	tests := []struct {
		name string
		args args
		want *Octet
	}{
		{
			name: "simple",
			args: args{
				hi: &Quartet{bits: 0b00000110},
				lo: &Quartet{bits: 0b00001001},
			},
			want: &Octet{bits: 0b01101001},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OctetFromQuartets(tt.args.hi, tt.args.lo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OctetFromQuartets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctet_IsSet(t *testing.T) {
	type fields struct {
		bits byte
	}
	type args struct {
		index OctetIndex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   `simple set`,
			fields: fields{
				bits: 0b00000100,
			},
			args:   args{
				index: Bit2,
			},
			want:   true,
		},
		{
			name:   `simple unset`,
			fields: fields{
				bits: 0b11111011,
			},
			args:   args{
				index: Bit2,
			},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Octet{
				bits: tt.fields.bits,
			}
			if got := o.IsSet(tt.args.index); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctet_OctetMap(t *testing.T) {
	type fields struct {
		bits byte
	}
	tests := []struct {
		name   string
		fields fields
		want   OctetMap
	}{
		{
			name: "simple",
			fields: fields{
				bits: 0b01101111,
			},
			want: OctetMap{
				B0: true,
				B1: true,
				B2: true,
				B3: true,
				B4: false,
				B5: true,
				B6: true,
				B7: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Octet{
				bits: tt.fields.bits,
			}
			if got := o.OctetMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OctetMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctet_OverwriteQuartet(t *testing.T) {
	type fields struct {
		bits byte
	}
	type args struct {
		val *Quartet
		pos QuartetPosition
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Octet
	}{
		{
			name: "simple",
			fields: fields{
				bits: 0b00000000,
			},
			args: args{
				val: &Quartet{
					bits: 0b11100111,
				},
				pos: Lo,
			},
			want: &Octet{
				bits: 0b00000111,
			},
		},{
			name: "simple",
			fields: fields{
				bits: 0b00000000,
			},
			args: args{
				val: &Quartet{
					bits: 0b11100111,
				},
				pos: Hi,
			},
			want: &Octet{
				bits: 0b01110000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Octet{
				bits: tt.fields.bits,
			}
			if got := o.OverwriteQuartet(tt.args.val, tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OverwriteQuartet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctet_QuartetMaps(t *testing.T) {
	type fields struct {
		bits byte
	}
	tests := []struct {
		name   string
		fields fields
		wantHi QuartetMap
		wantLo QuartetMap
	}{
		{
			name: "",
			fields: fields{
				bits: 0b10001110,
			},
			wantHi: QuartetMap{
				B0: false,
				B1: false,
				B2: false,
				B3: true,
			},
			wantLo: QuartetMap{
				B0: false,
				B1: true,
				B2: true,
				B3: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Octet{
				bits: tt.fields.bits,
			}
			gotHi, gotLo := o.QuartetMaps()
			if !reflect.DeepEqual(gotHi, tt.wantHi) {
				t.Errorf("QuartetMaps() gotHi = %v, want %v", gotHi, tt.wantHi)
			}
			if !reflect.DeepEqual(gotLo, tt.wantLo) {
				t.Errorf("QuartetMaps() gotLo = %v, want %v", gotLo, tt.wantLo)
			}
		})
	}
}

func TestOctet_Quartets(t *testing.T) {
	type fields struct {
		bits byte
	}
	tests := []struct {
		name   string
		fields fields
		wantHi *Quartet
		wantLo *Quartet
	}{
		{
			name:   "simple",
			fields: fields{
				bits: 0b11110001,
			},
			wantHi: &Quartet{bits: 0b00001111},
			wantLo: &Quartet{bits: 0b00000001},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Octet{
				bits: tt.fields.bits,
			}
			gotHi, gotLo := o.Quartets()
			if !reflect.DeepEqual(gotHi, tt.wantHi) {
				t.Errorf("Quartets() gotHi = %v, want %v", gotHi, tt.wantHi)
			}
			if !reflect.DeepEqual(gotLo, tt.wantLo) {
				t.Errorf("Quartets() gotLo = %v, want %v", gotLo, tt.wantLo)
			}
		})
	}
}

func TestOctet_Set(t *testing.T) {
	type fields struct {
		bits byte
	}
	type args struct {
		index OctetIndex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Octet
	}{
		{
			name:   "simple change",
			fields: fields{
				bits: 0b11111110,
			},
			args:   args{
				index: Bit0,
			},
			want:   &Octet{bits: 0b11111111},
		},
		{
			name:   "simple unchanged",
			fields: fields{
				bits: 0b11111111,
			},
			args:   args{
				index: Bit0,
			},
			want:   &Octet{bits: 0b11111111},
		},
		{
			name:   "simple change",
			fields: fields{
				bits: 0b01111111,
			},
			args:   args{
				index: Bit7,
			},
			want:   &Octet{bits: 0b11111111},
		},
		{
			name:   "simple unchanged",
			fields: fields{
				bits: 0b11111111,
			},
			args:   args{
				index: Bit7,
			},
			want:   &Octet{bits: 0b11111111},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Octet{
				bits: tt.fields.bits,
			}
			if got := o.Set(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctet_Toggle(t *testing.T) {
	type fields struct {
		bits byte
	}
	type args struct {
		index OctetIndex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Octet
	}{
		{
			name:   "simple unset",
			fields: fields{
				bits: 0b00001111,
			},
			args:   args{
				index: Bit0,
			},
			want:   &Octet{bits: 0b00001110},
		},
		{
			name:   "simple set",
			fields: fields{
				bits: 0b00001110,
			},
			args:   args{
				index: Bit0,
			},
			want:   &Octet{bits: 0b00001111},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Octet{
				bits: tt.fields.bits,
			}
			if got := o.Toggle(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Toggle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOctet_Unset(t *testing.T) {
	type fields struct {
		bits byte
	}
	type args struct {
		index OctetIndex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Octet
	}{
		{
			name:   "simple change",
			fields: fields{
				bits: 0b11111111,
			},
			args:   args{
				index: Bit6,
			},
			want:   &Octet{bits: 0b10111111},
		},
		{
			name:   "simple unchanged",
			fields: fields{
				bits: 0b11111111,
			},
			args:   args{
				index: Bit6,
			},
			want:   &Octet{bits: 0b10111111},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Octet{
				bits: tt.fields.bits,
			}
			if got := o.Unset(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuartet_IsSet(t *testing.T) {
	type fields struct {
		bits byte
	}
	type args struct {
		index QuartetIndex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "check true 0",
			fields: fields{
				bits: 0b11110001,
			},
			args:   args{
				index: QBit0,
			},
			want:   true,
		},{
			name:   "check false 0",
			fields: fields{
				bits: 0b11110000,
			},
			args:   args{
				index: QBit0,
			},
			want:   false,
		},
		{
			name:   "check true 3",
			fields: fields{
				bits: 0b11111001,
			},
			args:   args{
				index: QBit3,
			},
			want:   true,
		},{
			name:   "check false 3",
			fields: fields{
				bits: 0b11110111,
			},
			args:   args{
				index: QBit3,
			},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quartet{
				bits: tt.fields.bits,
			}
			if got := q.IsSet(tt.args.index); got != tt.want {
				t.Errorf("IsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuartet_Set(t *testing.T) {
	type fields struct {
		bits byte
	}
	type args struct {
		index QuartetIndex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Quartet
	}{
		{
			name:   "simple change",
			fields: fields{
				bits: 0b11111110,
			},
			args:   args{
				index: QBit0,
			},
			want:   &Quartet{bits: 0b11111111},
		},
		{
			name:   "simple unchanged",
			fields: fields{
				bits: 0b11111111,
			},
			args:   args{
				index: QBit0,
			},
			want:   &Quartet{bits: 0b11111111},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quartet{
				bits: tt.fields.bits,
			}
			if got := q.Set(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuartet_Toggle(t *testing.T) {
	type fields struct {
		bits byte
	}
	type args struct {
		index QuartetIndex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Quartet
	}{
		{
			name:   "simple unset",
			fields: fields{
				bits: 0b00001111,
			},
			args:   args{
				index: QBit0,
			},
			want:   &Quartet{bits: 0b00001110},
		},
		{
			name:   "simple set",
			fields: fields{
				bits: 0b00001110,
			},
			args:   args{
				index: QBit0,
			},
			want:   &Quartet{bits: 0b00001111},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quartet{
				bits: tt.fields.bits,
			}
			if got := q.Toggle(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Toggle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuartet_Unset(t *testing.T) {
	type fields struct {
		bits byte
	}
	type args struct {
		index QuartetIndex
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Quartet
	}{
		{
			name:   "simple change",
			fields: fields{
				bits: 0b11111111,
			},
			args:   args{
				index: QBit0,
			},
			want:   &Quartet{bits: 0b11111110},
		},
		{
			name:   "simple unchanged",
			fields: fields{
				bits: 0b11111110,
			},
			args:   args{
				index: QBit0,
			},
			want:   &Quartet{bits: 0b11111110},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quartet{
				bits: tt.fields.bits,
			}
			if got := q.Unset(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuartet_toHi(t *testing.T) {
	type fields struct {
		bits byte
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		{
			name:   "simple",
			fields: fields{
				bits: 0b00001111,
			},
			want:   0b11110000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quartet{
				bits: tt.fields.bits,
			}
			if got := q.toHi(); got != tt.want {
				t.Errorf("toHi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuartetsFromByte(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name   string
		args   args
		wantHi *Quartet
		wantLo *Quartet
	}{
		{
			name:   "simple",
			args:   args{
				b: 0b01101001,
			},
			wantHi: &Quartet{bits: 0b00000110},
			wantLo: &Quartet{bits: 0b00001001},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHi, gotLo := QuartetsFromByte(tt.args.b)
			if !reflect.DeepEqual(gotHi, tt.wantHi) {
				t.Errorf("QuartetsFromByte() gotHi = %v, want %v", gotHi, tt.wantHi)
			}
			if !reflect.DeepEqual(gotLo, tt.wantLo) {
				t.Errorf("QuartetsFromByte() gotLo = %v, want %v", gotLo, tt.wantLo)
			}
		})
	}
}
