package main

import "testing"

func TestGetSeeds(t *testing.T) {
	lines := ReadFile("../test_input.txt")
	expect := []int{79, 14, 55, 13}
	got := GetSeeds(lines[0])

	for i, r := range got {
		if r != expect[i] {
			t.Errorf("Wrong seed number found. Got %v, expected %v \n",
				r, expect[i])
		}
	}
}

func TestMapInput(t *testing.T) {
	lines := ReadFile("../test_input.txt")
	expect := []ResourceMap{
		ResourceMap{
			name: "seed-to-soil",
			resourceRange: []ResourceRange{
				ResourceRange{
					destinationRangeStart: 50,
					sourceRangeStart:      98,
					rangeLength:           2,
				},
				ResourceRange{
					destinationRangeStart: 52,
					sourceRangeStart:      50,
					rangeLength:           48,
				},
			},
		},
		ResourceMap{
			name: "soil-to-fertilizer",
			resourceRange: []ResourceRange{
				ResourceRange{
					destinationRangeStart: 0,
					sourceRangeStart:      15,
					rangeLength:           37,
				},
				ResourceRange{
					destinationRangeStart: 37,
					sourceRangeStart:      52,
					rangeLength:           2,
				},
				ResourceRange{
					destinationRangeStart: 39,
					sourceRangeStart:      0,
					rangeLength:           15,
				},
			},
		},
		ResourceMap{
			name: "fertilizer-to-water",
			resourceRange: []ResourceRange{
				ResourceRange{
					destinationRangeStart: 49,
					sourceRangeStart:      53,
					rangeLength:           8,
				},
				ResourceRange{
					destinationRangeStart: 0,
					sourceRangeStart:      11,
					rangeLength:           42,
				},
				ResourceRange{
					destinationRangeStart: 42,
					sourceRangeStart:      0,
					rangeLength:           7,
				},
				ResourceRange{
					destinationRangeStart: 57,
					sourceRangeStart:      7,
					rangeLength:           4,
				},
			},
		},
		ResourceMap{
			name: "water-to-light",
			resourceRange: []ResourceRange{
				ResourceRange{
					destinationRangeStart: 88,
					sourceRangeStart:      18,
					rangeLength:           7,
				},
				ResourceRange{
					destinationRangeStart: 18,
					sourceRangeStart:      25,
					rangeLength:           70,
				},
			},
		},
		ResourceMap{
			name: "light-to-temperature",
			resourceRange: []ResourceRange{
				ResourceRange{
					destinationRangeStart: 45,
					sourceRangeStart:      77,
					rangeLength:           23,
				},
				ResourceRange{
					destinationRangeStart: 81,
					sourceRangeStart:      45,
					rangeLength:           19,
				},
				ResourceRange{
					destinationRangeStart: 68,
					sourceRangeStart:      64,
					rangeLength:           13,
				},
			},
		},
		ResourceMap{
			name: "temperature-to-humidity",
			resourceRange: []ResourceRange{
				ResourceRange{
					destinationRangeStart: 0,
					sourceRangeStart:      69,
					rangeLength:           1,
				},
				ResourceRange{
					destinationRangeStart: 1,
					sourceRangeStart:      0,
					rangeLength:           69,
				},
			},
		},
		ResourceMap{
			name: "humidity-to-location",
			resourceRange: []ResourceRange{
				ResourceRange{
					destinationRangeStart: 60,
					sourceRangeStart:      56,
					rangeLength:           37,
				},
				ResourceRange{
					destinationRangeStart: 56,
					sourceRangeStart:      93,
					rangeLength:           4,
				},
			},
		},
	}

	got := MapInput(lines)

	for i, r := range got {
		if r.name != expect[i].name {
			t.Errorf("Wrong ResourceMap name found. Got %v, expected %v \n",
				r, expect[i].name)
		}
		for j, rr := range got[i].resourceRange {
			if rr.destinationRangeStart != expect[i].resourceRange[j].destinationRangeStart {
				t.Errorf("Wrong destinationRangeStart found. Got %v, expected %v \n",
					rr.destinationRangeStart, expect[i].resourceRange[j].destinationRangeStart)
			}
			if rr.sourceRangeStart != expect[i].resourceRange[j].sourceRangeStart {
				t.Errorf("Wrong sourceRangeStart found. Got %v, expected %v \n",
					rr.sourceRangeStart, expect[i].resourceRange[j].sourceRangeStart)
			}
			if rr.rangeLength != expect[i].resourceRange[j].rangeLength {
				t.Errorf("Wrong rangeLength found. Got %v, expected %v \n",
					rr.rangeLength, expect[i].resourceRange[j].rangeLength)
			}
		}
	}
}
