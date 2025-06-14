package static_data

type ContainerCapacity struct {
	Name                 string
	VolumeCapacityAmount float64
	VolumeCapacityUnit   string
	WeightCapacityLbs    float64
}

var ContainerCapacities = func() map[string]ContainerCapacity {
	return map[string]ContainerCapacity{
		"backpack": {
			Name:                 "Backpack",
			VolumeCapacityAmount: 1,
			VolumeCapacityUnit:   "cubic feet",
			WeightCapacityLbs:    30,
		},
		"barrel": {
			Name:                 "Barrel",
			VolumeCapacityAmount: 4,
			VolumeCapacityUnit:   "cubic feet",
			WeightCapacityLbs:    300,
		},
		"basket": {
			Name:                 "Basket",
			VolumeCapacityAmount: 2,
			VolumeCapacityUnit:   "cubic feet",
			WeightCapacityLbs:    40,
		},
		"bottle": {
			Name:                 "Bottle",
			VolumeCapacityAmount: 1,
			VolumeCapacityUnit:   "pint",
			WeightCapacityLbs:    1,
		},
		"flask": {
			Name:                 "Flask",
			VolumeCapacityAmount: 1,
			VolumeCapacityUnit:   "pint",
			WeightCapacityLbs:    1,
		},
		"tankard": {
			Name:                 "Tankard",
			VolumeCapacityAmount: 1,
			VolumeCapacityUnit:   "pint",
			WeightCapacityLbs:    1,
		},
		"chest": {
			Name:                 "Chest",
			VolumeCapacityAmount: 12,
			VolumeCapacityUnit:   "cubic feet",
			WeightCapacityLbs:    300,
		},
		"jug": {
			Name:                 "Jug",
			VolumeCapacityAmount: 1,
			VolumeCapacityUnit:   "gallon",
			WeightCapacityLbs:    8,
		},
		"pitcher": {
			Name:                 "Pitcher",
			VolumeCapacityAmount: 1,
			VolumeCapacityUnit:   "gallon",
			WeightCapacityLbs:    8,
		},
		"pot": {
			Name:                 "Pot",
			VolumeCapacityAmount: 1,
			VolumeCapacityUnit:   "gallon",
			WeightCapacityLbs:    8,
		},
		"pouch": {
			Name:                 "Pouch",
			VolumeCapacityAmount: 0.20,
			VolumeCapacityUnit:   "cubic feet",
			WeightCapacityLbs:    6,
		},
		"sack": {
			Name:                 "Sack",
			VolumeCapacityAmount: 1,
			VolumeCapacityUnit:   "cubic feet",
			WeightCapacityLbs:    30,
		},
		"vial": {
			Name:                 "Vial",
			VolumeCapacityAmount: 4,
			VolumeCapacityUnit:   "ounces",
			WeightCapacityLbs:    0.25,
		},
		"waterskin": {
			Name:                 "Waterskin",
			VolumeCapacityAmount: 4,
			VolumeCapacityUnit:   "pints",
			WeightCapacityLbs:    4,
		},
	}
}
