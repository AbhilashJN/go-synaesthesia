package color

var Palettes = []ColorPalette{
	{
		BgPrimary:   newHSBColor(60, 12, 100),
		BgSecondary: newHSBColor(60, 12, 100),
		Colors: []HSBColor{
			newHSBColor(88, 100, 54),
			newHSBColor(1, 100, 54),
			newHSBColor(215, 100, 54),
			newHSBColor(54, 0, 32),
		},
		GetColorAt: Lerp4Corners,
	},
	{
		BgPrimary:   newHSBColor(55, 0, 3),
		BgSecondary: newHSBColor(55, 0, 3),
		Colors: []HSBColor{
			newHSBColor(34, 67, 99),
			newHSBColor(200, 61, 74),
			newHSBColor(146, 46, 52),
			newHSBColor(359, 54, 92),
		},
		GetColorAt: GridDistribution,
	},
	{
		BgPrimary:   newHSBColor(40, 10, 94),
		BgSecondary: newHSBColor(40, 10, 94),
		Colors: []HSBColor{
			newHSBColor(181, 87, 61),
			newHSBColor(11, 86, 92),
			newHSBColor(231, 15, 17),
			newHSBColor(181, 87, 61),
		},
		GetColorAt: GridDistribution,
	},
	{
		BgPrimary:   newHSBColor(60, 3, 92),
		BgSecondary: newHSBColor(60, 3, 92),
		Colors: []HSBColor{
			newHSBColor(352, 90, 83),
			newHSBColor(105, 78, 69),
			newHSBColor(199, 48, 71),
		},
		GetColorAt: ThreeColorDistribution,
	},
	{
		BgPrimary:   newHSBColor(353, 50, 100),
		BgSecondary: newHSBColor(353, 50, 100),
		Colors: []HSBColor{
			newHSBColor(216, 13, 74),
			newHSBColor(351, 81, 81),
			newHSBColor(228, 10, 20),
		},
		GetColorAt: ThreeColorDistribution,
	},
	{
		BgPrimary:   newHSBColor(0, 0, 100),
		BgSecondary: newHSBColor(0, 0, 100),
		Colors: []HSBColor{
			newHSBColor(0, 0, 0),
			newHSBColor(0, 0, 0),
			newHSBColor(0, 0, 0),
		},
		GetColorAt: ThreeColorDistribution,
	},
	{
		BgPrimary:   newHSBColor(214, 50, 45),
		BgSecondary: newHSBColor(214, 50, 45),
		Colors: []HSBColor{
			newHSBColor(168, 100, 73),
			newHSBColor(53, 78, 84),
			newHSBColor(103, 52, 74),
		},
		GetColorAt: ThreeColorDistribution,
	},
	{
		BgPrimary:   newHSBColor(252, 94, 56),
		BgSecondary: newHSBColor(252, 94, 56),
		Colors: []HSBColor{
			newHSBColor(0, 0, 100),
			newHSBColor(0, 0, 90),
			newHSBColor(0, 0, 80),
			newHSBColor(0, 0, 95),
		},
		GetColorAt: GridDistribution,
	},
}
