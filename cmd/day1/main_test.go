package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDial_Move(t *testing.T) {
	type fields struct {
		Position int
		Size     int
		Moves    []*Move
		Zeros    int
	}
	type args struct {
		moves []*Move
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantZeros int
		wantPos   int
	}{
		{
			name: "test valid right non zero move",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "R", Distance: 3}},
			},
			wantZeros: 0,
			wantPos:   53,
		}, {
			name: "test valid right zero move",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "R", Distance: 65}},
			},
			wantZeros: 1,
			wantPos:   15,
		}, {
			name: "test valid right zero moves",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "R", Distance: 150}},
			},
			wantZeros: 2,
			wantPos:   0,
		}, {
			name: "test valid right zero move edge case",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "R", Distance: 149}},
			},
			wantZeros: 1,
			wantPos:   99,
		}, {
			name: "test valid left non zero move",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "L", Distance: 3}},
			},
			wantZeros: 0,
			wantPos:   47,
		}, {
			name: "test valid left zero move",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "L", Distance: 65}},
			},
			wantZeros: 1,
			wantPos:   85,
		}, {
			name: "test valid left zero moves",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "L", Distance: 150}},
			},
			wantZeros: 2,
			wantPos:   00,
		}, {
			name: "test valid left zero move edge case",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "L", Distance: 149}},
			},
			wantZeros: 1,
			wantPos:   1,
		}, {
			name: "move 3 positions 0 zeros",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "R", Distance: 3}, {Direction: "R", Distance: 3}, {Direction: "R", Distance: 3}},
			},
			wantZeros: 0,
			wantPos:   59,
		}, {
			name: "move 3 positions 2 zeros",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "R", Distance: 50}, {Direction: "R", Distance: 60}, {Direction: "R", Distance: 60}},
			},
			wantZeros: 2,
			wantPos:   20,
		}, {
			name: "move 6 positions 3 zeros",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "R", Distance: 50}, {Direction: "L", Distance: 60}, {Direction: "L", Distance: 60},
					{Direction: "R", Distance: 50}, {Direction: "R", Distance: 60}, {Direction: "L", Distance: 60}},
			},
			wantZeros: 3,
			wantPos:   30,
		}, {
			name: "move 3 positions 2 zeros",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "L", Distance: 50}, {Direction: "R", Distance: 100}},
			},
			wantZeros: 2,
			wantPos:   0,
		}, {
			name: "basic single left pass",
			fields: fields{
				Position: 50,
				Size:     100,
				Moves:    []*Move{},
				Zeros:    0,
			},
			args: args{
				moves: []*Move{{Direction: "L", Distance: 200}},
			},
			wantZeros: 2,
			wantPos:   50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dial{
				Position: tt.fields.Position,
				Size:     tt.fields.Size,
				Moves:    tt.fields.Moves,
				Zeros:    tt.fields.Zeros,
			}
			for _, m := range tt.args.moves {
				d.Move(m)
			}
			assert.Equal(t, tt.wantZeros, d.Zeros)
			assert.Equal(t, tt.wantPos, d.Position)
		})
	}
}
