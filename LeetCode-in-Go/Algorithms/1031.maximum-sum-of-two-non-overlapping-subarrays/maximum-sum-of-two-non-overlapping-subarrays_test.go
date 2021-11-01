package problem1031

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	L   int
	M   int
	ans int
}{

	{
		[]int{360, 629, 175, 292, 546, 681, 708, 458, 411, 596, 978, 592, 7, 830, 968, 589, 234, 743, 997, 542, 431, 496, 545, 97, 781, 978, 31, 531, 651, 855, 692, 852, 137, 706, 458, 47, 963, 993, 205, 223, 180, 661, 418, 458, 780, 365, 510, 344, 945, 988, 744, 498, 71, 927, 971, 240, 829, 546, 629, 944, 248, 116, 390, 667, 343, 390, 295, 701, 143, 197, 811, 608, 190, 858, 954, 905, 335, 373, 797, 598, 817, 965, 1000, 861, 612, 306, 783, 772, 755, 439, 256, 82, 993, 207, 695, 613, 930, 236, 0, 439, 348, 400, 825, 12, 306, 846, 792, 334, 339, 814, 631, 51, 633, 666, 852, 686, 939, 139, 264, 508, 675, 784, 60, 89, 730, 958, 879, 356, 550, 135, 755, 517, 390, 945, 536, 406, 383, 110, 389, 649, 802, 682, 257, 770, 420, 165, 101, 126, 475, 42, 714, 305, 547, 750, 640, 61, 914, 358, 617, 266, 69, 543, 771, 509, 162, 738, 806, 668, 734, 845, 718, 795, 956, 203, 368, 523, 440, 674, 586, 552, 928, 954, 414, 623, 960, 673, 279, 76, 814, 80, 881, 834, 377, 434, 84, 737, 166, 173, 429, 434, 925, 252, 1, 172, 209, 340, 242, 173, 460, 520, 632, 618, 260, 401, 300, 879, 74, 228, 5, 694, 976, 28, 448, 131, 431, 702, 217, 766, 841, 513, 647, 541, 48, 321, 986, 200, 998, 657, 338, 365, 276, 562, 897, 161, 586, 381, 775, 140, 355, 558, 237, 683, 953, 350, 122, 295, 606, 935, 736, 641, 239, 415, 999, 834, 935, 637, 940, 972, 670, 778, 257, 446, 118, 799, 28, 921, 987, 99, 415, 844, 827, 294, 454, 993, 58, 916, 312, 571, 217, 487, 619, 191, 616, 450, 68, 688, 270, 812, 468, 239, 762, 636, 267, 745, 560, 212, 228, 425, 364, 670, 231, 512, 639, 732, 450, 472, 78, 562, 145, 963, 331, 325, 323, 804, 476, 149, 247, 88, 237, 370, 82, 392, 399, 432, 135, 890, 731, 871, 601, 37, 124, 691, 249, 240, 518, 69, 430, 761, 963, 803, 397, 575, 535, 199, 311, 39, 312, 317, 123, 103, 249, 77, 131, 730, 712, 405, 134, 662, 489, 657, 846, 824, 788, 257, 332, 891, 31, 618, 584, 482, 141, 300, 915, 493, 760, 954, 308, 584, 704, 857, 464, 69, 601, 683, 918, 456, 24, 81, 39, 23, 829, 567, 402, 84, 619, 247, 437, 529, 190, 178, 195, 57, 818, 884, 498, 204, 509, 354, 86, 105, 663, 627, 913, 172, 330, 828, 763, 508, 306, 861, 962, 874, 806, 891, 304, 927, 700, 77, 811, 136, 459, 485, 29, 778, 344, 971, 377, 588, 601, 429, 282, 710, 794, 918, 312, 250, 950, 143, 947, 907, 392, 330, 270, 326, 439, 380, 155, 706, 962, 683, 877, 711, 459, 47, 317, 691, 132, 989, 816, 748, 765, 989, 970, 178, 546, 831, 572, 651, 664, 221, 424, 992, 90, 889, 418, 11, 278, 776, 570, 57, 970, 391, 623, 177, 797, 123, 693, 449, 662, 679, 296, 987, 475, 918, 486, 624, 218, 512, 185, 526, 831, 433, 927, 447, 381, 928, 754, 459, 607, 649, 273, 509, 163, 591, 734, 573, 664, 501, 607, 191, 675, 939, 185, 465, 644, 307, 695, 479, 674, 751, 922, 803, 860, 932, 768, 272, 294, 223, 977, 536, 150, 602, 612, 863, 768, 154, 477, 156, 331, 404, 345, 401, 4, 307, 606, 60, 75, 2, 295, 297, 848, 612, 661, 640, 635, 841, 800, 174, 366, 54, 539, 786, 879, 469, 322, 32, 588, 904, 411, 577, 97, 621, 32, 4, 784, 346, 791, 659, 62, 228, 422, 532, 961, 889, 859, 441, 978, 515, 311, 665, 728, 892, 299, 22, 573, 473, 963, 150, 831, 858, 156, 139, 472, 53, 579, 105, 337, 576, 83, 409, 762, 556, 906, 836, 782, 825, 380, 416, 2, 904, 676, 265, 637, 757, 970, 911, 62, 816, 784, 511, 372, 572, 889, 25, 94, 722, 514, 137, 147, 863, 674, 236, 494, 351, 706, 962, 192, 930, 833, 580, 992, 239, 14, 786, 749, 746, 166, 321, 951, 820, 387, 157, 852, 376, 258, 682, 882, 742, 995, 291, 304, 920, 131, 800, 78, 264, 164, 428, 301, 952, 510, 408, 336, 258, 442, 933, 866, 366, 991, 314, 323, 190, 326, 315, 195, 848, 583, 357, 462, 336, 417, 838, 208, 969, 559, 48, 938, 248, 702, 5, 928, 936, 378, 792, 325, 215, 240, 166, 643, 395, 3, 158, 89, 742, 448, 90, 348, 694, 797, 374, 886, 852, 19, 393, 202, 906, 693, 308, 227, 558, 379, 385, 691, 953, 881, 973, 762, 663, 539, 63, 983, 633, 310, 798, 389, 140, 411, 100, 40, 826, 328, 997, 856, 433, 586, 204, 724, 340, 773, 656, 694, 498, 257, 658, 468, 8, 486, 380, 118, 123, 901, 108, 539, 763, 461, 294, 327, 24, 46, 495, 744, 417, 823, 844, 511, 40, 245, 854, 1, 210, 386, 64, 688, 516, 461, 672, 371, 946, 155, 524, 56, 825, 23, 579, 580, 198, 312, 179, 445, 47, 644, 759, 89, 676, 314, 567, 298, 907, 555, 646, 544, 403, 643, 194, 485, 428, 139, 296, 808, 445, 343, 236, 51, 132, 945, 9, 928, 929, 90, 313, 458, 231, 429, 662, 707, 833, 407, 589, 708, 555, 762, 806, 638, 553, 489, 577, 497, 507, 905, 51, 204, 297, 575, 876, 849, 145, 363, 68, 431, 532, 881, 15, 355, 956, 575, 657, 809, 544, 343, 80, 641, 93, 619, 210, 834, 764, 26, 685, 312, 434, 974, 999, 384, 704, 590, 547, 234, 782, 448, 22, 866, 614, 683, 852, 564, 846, 430, 440, 899, 383, 244, 644, 809, 304, 651, 847, 314, 378, 495, 488, 418, 324, 838, 571, 800, 461, 2, 904, 83, 226, 836, 146, 667, 154, 603, 706, 336, 40, 802, 19, 449, 726, 539, 73, 511, 695, 527, 79, 69, 395, 876, 415, 943, 756, 398, 348, 900, 228, 953, 119},
		204,
		648,
		434985,
	},

	{
		[]int{4, 5, 14, 16, 16, 20, 7, 13, 8, 15},
		3,
		5,
		109,
	},

	{
		[]int{0, 6, 5, 2, 2, 5, 1, 9, 4},
		1,
		2,
		20,
	},

	{
		[]int{3, 8, 1, 3, 2, 1, 8, 9, 0},
		3,
		2,
		29,
	},

	{
		[]int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8},
		4,
		3,
		31,
	},

	// 可以有多个 testcase
}

func Test_maxSumTwoNoOverlap(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxSumTwoNoOverlap(tc.A, tc.L, tc.M), "输入:%v", tc)
	}
}

func Benchmark_maxSumTwoNoOverlap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxSumTwoNoOverlap(tc.A, tc.L, tc.M)
		}
	}
}
