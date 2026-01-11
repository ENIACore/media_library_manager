package metadata

import (
	"testing"
)


func TestHeight(t *testing.T) {

	tests := []struct{
		name		string
		input		Entry
		expected	int
	}{
		{
			name:		"root only",
			input:		Entry{
				Parent: nil,
				Children: nil,
				Depth: 0,
			},
			expected:	0,
		},
		{
			name:		"one child entry",
			input:		Entry{
				Parent: nil,
				Children: []*Entry{
					{
						Parent: nil,
						Children: nil,
						Depth: 1,
					},
				},
				Depth: 0,
			},
			expected:	1,
		},
		{
			name:		"grandchild entry",
			input:		Entry{
				Parent: nil,
				Children: []*Entry{
					{
						Parent: nil,
						Children: []*Entry{
							{
								Parent: nil,
								Children: nil,
								Depth: 2,
							},
						},
						Depth: 1,
					},
				},
				Depth: 0,
			},
			expected:	2,
		},
		{
			name:		"grandchild entry and two child entries",
			input:		Entry{
				Parent: nil,
				Children: []*Entry{
					{
						Parent: nil,
						Children: nil,
						Depth: 1,
					},
					{
						Parent: nil,
						Children: []*Entry{
							{
								Parent: nil,
								Children: nil,
								Depth: 2,
							},
						},
						Depth: 1,
					},
				},
				Depth: 0,
			},
			expected:	2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			height := test.input.Height()
			if height != test.expected {
				t.Errorf("Height = %v, want %v", height, test.expected)
			}
		})
	}
}
