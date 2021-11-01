package problem1122

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	arr1 []int
	arr2 []int
	ans  []int
}{

	{
		[]int{940, 779, 194, 460, 545, 15, 143, 228, 733, 274, 652, 234, 831, 932, 939, 378, 159, 824, 614, 152, 989, 256, 919, 133, 922, 491, 208, 702, 503, 305, 906, 259, 346, 535, 683, 657, 929, 775, 384, 787, 966, 199, 623, 93, 507, 108, 643, 456, 897, 394, 628, 529, 277, 190, 438, 837, 341, 631, 988, 409, 629, 540, 98, 863, 678, 845, 452, 156, 600, 77, 360, 347, 522, 221, 325, 949, 162, 810, 688, 892, 665, 734, 475, 209, 816, 636, 421, 374, 792, 499, 386, 213, 442, 977, 310, 367, 379, 671, 626, 243, 482, 513, 128, 541, 705, 297, 290, 649, 265, 289, 762, 302, 418, 759, 204, 331, 288, 260, 389, 195, 60, 765, 107, 763, 441, 106, 743, 595, 461, 299, 616, 68, 19, 909, 978, 230, 925, 605, 27, 856, 180, 771, 361, 12, 911, 800, 655, 574, 3, 834, 6, 313, 84, 236, 923, 524, 884, 562, 86, 684, 757, 520, 927, 891, 240, 385, 565, 75, 33, 437, 1000, 537, 492, 607, 139, 877, 67, 53, 191, 263, 376, 741, 24, 487, 21, 687, 292, 92, 801, 161, 183, 716, 822, 231, 672, 189, 90, 343, 566, 8, 345, 865, 16, 560, 166, 575, 275, 878, 812, 739, 18, 531, 451, 171, 330, 679, 321, 89, 836, 550, 314, 425, 729, 309, 868, 916, 262, 176, 692, 711, 695, 797, 381, 375, 212, 814, 153, 366, 589, 690, 419, 738, 168, 337, 601, 899, 914, 453, 745, 145, 480, 23, 47, 510, 844, 71, 848, 114, 669, 157, 673, 591, 872, 205, 658, 774, 501, 323, 653, 606, 542, 434, 272, 901, 383, 39, 990, 728, 87, 693, 268, 890, 364, 70, 196, 617, 553, 766, 782, 273, 342, 188, 319, 704, 694, 584, 178, 430, 610, 650, 789, 903, 637, 455, 144, 536, 449, 339, 770, 777, 109, 724, 760, 737, 564, 250, 498, 284, 397, 357, 559, 829, 630, 594, 663, 353, 120, 618, 938, 767, 862, 577, 817, 276, 311, 576, 751, 28, 432, 833, 225, 468, 956, 710, 583, 5, 538, 363, 187, 700, 850, 502, 283, 459, 14, 981, 202, 484, 37, 422, 937, 736, 888, 490, 613, 373, 136, 568, 943, 344, 680, 900, 985, 995, 511, 44, 233, 134, 975, 883, 338, 497, 61, 905, 867, 110, 398, 917, 0, 298, 185, 644, 997, 244, 846, 137, 802, 926, 223, 62, 226, 315, 615, 138, 854, 852, 857, 725, 656, 703, 586, 744, 783, 318, 847, 301, 674, 219, 448, 48, 181, 882, 963, 20, 549, 116, 457, 602, 945, 111, 80, 586, 433, 184, 835, 991, 676, 181, 417, 944, 750, 494, 100, 640, 843, 963, 783, 819, 571, 670, 886, 493, 182, 796, 371, 933, 248, 639, 924, 706, 443, 635, 264, 611, 772, 875, 131, 755, 838, 387, 950, 660, 504, 662, 786, 76, 151, 206, 192, 349, 119, 295, 902, 124, 294, 532, 116, 329, 135, 63, 596, 971, 91, 286, 912, 698, 509, 72, 962, 994, 426, 95, 411, 151, 668, 980, 974, 474, 500, 555, 509, 830, 621, 783, 793, 898, 446, 640, 489, 556, 772, 201, 942, 113, 533, 619, 791, 552, 746, 675, 942, 251, 708, 908, 740, 521, 952, 380, 141, 554, 420, 494, 404, 676, 439, 619, 593, 668, 795, 942, 963, 279, 624, 380, 958, 224, 457, 172, 388, 764, 182, 454, 1, 151, 747, 970, 222, 130, 790, 983, 348, 332, 286, 840, 467, 388, 251, 811, 811, 11, 855, 718, 135, 718, 612, 462, 226, 372, 173, 828, 597, 640, 861, 40, 504, 962, 744, 543, 887, 509, 450, 830, 326, 329, 242, 532, 46, 192, 13, 436, 450, 241, 746, 436, 138, 556, 667, 561, 732, 570, 854, 539, 306, 10, 102, 445, 753, 740, 880, 224, 56, 795, 428, 101, 413, 586, 835, 973, 332, 348, 42, 132, 380, 291, 186, 382, 935, 885, 117, 112, 869, 115, 216, 485, 769, 732, 141, 596, 349, 462, 473, 731, 707, 410, 65, 838, 414, 355, 992, 570, 662, 790, 201, 95, 622, 795, 670, 206, 258, 811, 182, 281, 165, 572, 970, 352, 632, 968, 500, 232, 135, 795, 362, 755, 979, 116, 247, 847, 902, 744, 124, 924, 402, 428, 239, 29, 723, 13, 864, 65, 698, 805, 567, 785, 270, 910, 35, 526, 51, 406, 718, 307, 316, 13, 835, 588, 219, 320, 203, 908, 748, 443, 141, 612, 404, 160, 270, 982, 720, 193, 879, 849, 732, 962, 154, 181, 392, 886, 91, 252, 76, 681, 572, 876, 677, 328, 457, 951, 26, 880, 876, 252, 588, 457, 229, 426, 388, 721, 295, 481, 967, 934, 103, 308, 293, 326, 123, 76, 463, 715, 563, 993, 962, 296, 961, 94, 881, 151, 706, 179, 749, 660, 46, 544, 72, 443, 85, 715, 786, 278},
		[]int{61, 153, 262, 606, 917, 434, 47, 409, 190, 899, 166, 503, 183, 156, 694, 911, 24, 765, 451, 244, 234, 480, 375, 438, 16, 290, 653, 511, 498, 814, 834, 128, 367, 688, 344, 501, 739, 339, 702, 687, 385, 939, 268, 27, 176, 468, 833, 204, 607, 421, 817, 87, 777, 824, 484, 724, 92, 455, 452, 692, 243, 779, 535, 789, 360, 628, 430, 679, 18, 460, 265, 600, 800, 8, 274, 644, 836, 196, 529, 774, 314, 575, 338, 977, 614, 110, 831, 381, 325, 995, 771, 613, 187, 848, 865, 273, 364, 39, 502, 705, 84, 845, 905, 704, 134, 292, 311, 925, 497, 671, 67, 566, 114, 693, 610, 510, 379, 975, 741, 762, 93, 213, 901, 678, 695, 751, 230, 492, 892, 212, 283, 161, 323, 536, 700, 914, 760, 297, 342, 309, 490, 6, 932, 792, 923, 199, 846, 345, 376, 816, 337, 331, 770, 98, 23, 862, 108, 711, 374, 966, 812, 649, 728, 617, 277, 903, 383, 194, 956, 574, 137, 595, 766, 185, 289, 168, 672, 346, 937, 856, 145, 658, 763, 343, 33, 171, 631, 545, 906, 542, 397, 900, 144, 456, 180, 21, 90, 441, 988, 537, 106, 208, 989, 366, 680, 890, 5, 652, 550, 298, 143, 629, 449, 782, 767, 276, 520, 710, 475, 940, 188, 442, 560, 341, 432, 883, 690, 891, 191, 231, 513, 373, 347, 12, 797, 288, 195, 386, 990, 636, 225, 884, 1000, 482, 487, 418, 319, 109, 863, 284, 745, 759, 801, 394, 829, 305, 363, 202, 259, 531, 616, 736, 44, 540, 321, 844, 152, 353, 565, 657, 802, 453, 663, 938, 68, 655, 729, 14, 491, 310, 228, 263, 605, 507, 637, 665, 205, 499, 949, 743, 589, 909, 878, 978, 53, 419, 330, 302, 256, 837, 757, 522, 221, 576, 240, 897, 981, 3, 15, 236, 524, 425, 583, 19, 60, 929, 716, 623, 107, 461, 684, 553, 868, 673, 361, 822, 275, 357, 422, 775, 926, 643, 594, 299, 398, 260, 71, 601, 538, 28, 564, 62, 922, 626, 568, 389, 120, 669, 162, 877, 810, 888, 133, 384, 541, 75, 139, 209, 70, 734, 733, 37, 159, 562, 916, 997, 683, 919, 927, 591, 787, 77, 136, 850, 189, 157, 559, 378, 738, 618, 233, 650, 584, 867, 86, 630, 985, 313, 0, 272, 250, 223, 943, 178, 437, 89, 737, 872, 577, 459},
		[]int{61, 153, 262, 606, 917, 434, 47, 409, 190, 899, 166, 503, 183, 156, 694, 911, 24, 765, 451, 244, 234, 480, 375, 438, 16, 290, 653, 511, 498, 814, 834, 128, 367, 688, 344, 501, 739, 339, 702, 687, 385, 939, 268, 27, 176, 468, 833, 204, 607, 421, 817, 87, 777, 824, 484, 724, 92, 455, 452, 692, 243, 779, 535, 789, 360, 628, 430, 679, 18, 460, 265, 600, 800, 8, 274, 644, 836, 196, 529, 774, 314, 575, 338, 977, 614, 110, 831, 381, 325, 995, 771, 613, 187, 848, 865, 273, 364, 39, 502, 705, 84, 845, 905, 704, 134, 292, 311, 925, 497, 671, 67, 566, 114, 693, 610, 510, 379, 975, 741, 762, 93, 213, 901, 678, 695, 751, 230, 492, 892, 212, 283, 161, 323, 536, 700, 914, 760, 297, 342, 309, 490, 6, 932, 792, 923, 199, 846, 345, 376, 816, 337, 331, 770, 98, 23, 862, 108, 711, 374, 966, 812, 649, 728, 617, 277, 903, 383, 194, 956, 574, 137, 595, 766, 185, 289, 168, 672, 346, 937, 856, 145, 658, 763, 343, 33, 171, 631, 545, 906, 542, 397, 900, 144, 456, 180, 21, 90, 441, 988, 537, 106, 208, 989, 366, 680, 890, 5, 652, 550, 298, 143, 629, 449, 782, 767, 276, 520, 710, 475, 940, 188, 442, 560, 341, 432, 883, 690, 891, 191, 231, 513, 373, 347, 12, 797, 288, 195, 386, 990, 636, 225, 884, 1000, 482, 487, 418, 319, 109, 863, 284, 745, 759, 801, 394, 829, 305, 363, 202, 259, 531, 616, 736, 44, 540, 321, 844, 152, 353, 565, 657, 802, 453, 663, 938, 68, 655, 729, 14, 491, 310, 228, 263, 605, 507, 637, 665, 205, 499, 949, 743, 589, 909, 878, 978, 53, 419, 330, 302, 256, 837, 757, 522, 221, 576, 240, 897, 981, 3, 15, 236, 524, 425, 583, 19, 60, 929, 716, 623, 107, 461, 684, 553, 868, 673, 361, 822, 275, 357, 422, 775, 926, 643, 594, 299, 398, 260, 71, 601, 538, 28, 564, 62, 922, 626, 568, 389, 120, 669, 162, 877, 810, 888, 133, 384, 541, 75, 139, 209, 70, 734, 733, 37, 159, 562, 916, 997, 683, 919, 927, 591, 787, 77, 136, 850, 189, 157, 559, 378, 738, 618, 233, 650, 584, 867, 86, 630, 985, 313, 0, 272, 250, 223, 943, 178, 437, 89, 737, 872, 577, 459, 1, 10, 11, 13, 13, 13, 20, 26, 29, 35, 40, 42, 46, 46, 48, 51, 56, 63, 65, 65, 72, 72, 76, 76, 76, 80, 85, 91, 91, 94, 95, 95, 100, 101, 102, 103, 111, 112, 113, 115, 116, 116, 116, 117, 119, 123, 124, 124, 130, 131, 132, 135, 135, 135, 138, 138, 141, 141, 141, 151, 151, 151, 151, 154, 160, 165, 172, 173, 179, 181, 181, 181, 182, 182, 182, 184, 186, 192, 192, 193, 201, 201, 203, 206, 206, 216, 219, 219, 222, 224, 224, 226, 226, 229, 232, 239, 241, 242, 247, 248, 251, 251, 252, 252, 258, 264, 270, 270, 278, 279, 281, 286, 286, 291, 293, 294, 295, 295, 296, 301, 306, 307, 308, 315, 316, 318, 320, 326, 326, 328, 329, 329, 332, 332, 348, 348, 349, 349, 352, 355, 362, 371, 372, 380, 380, 380, 382, 387, 388, 388, 388, 392, 402, 404, 404, 406, 410, 411, 413, 414, 417, 420, 426, 426, 428, 428, 433, 436, 436, 439, 443, 443, 443, 445, 446, 448, 450, 450, 454, 457, 457, 457, 457, 462, 462, 463, 467, 473, 474, 481, 485, 489, 493, 494, 494, 500, 500, 504, 504, 509, 509, 509, 521, 526, 532, 532, 533, 539, 543, 544, 549, 552, 554, 555, 556, 556, 561, 563, 567, 570, 570, 571, 572, 572, 586, 586, 586, 588, 588, 593, 596, 596, 597, 602, 611, 612, 612, 615, 619, 619, 621, 622, 624, 632, 635, 639, 640, 640, 640, 656, 660, 660, 662, 662, 667, 668, 668, 670, 670, 674, 675, 676, 676, 677, 681, 698, 698, 703, 706, 706, 707, 708, 715, 715, 718, 718, 718, 720, 721, 723, 725, 731, 732, 732, 732, 740, 740, 744, 744, 744, 746, 746, 747, 748, 749, 750, 753, 755, 755, 764, 769, 772, 772, 783, 783, 783, 785, 786, 786, 790, 790, 791, 793, 795, 795, 795, 795, 796, 805, 811, 811, 811, 819, 828, 830, 830, 835, 835, 835, 838, 838, 840, 843, 847, 847, 849, 852, 854, 854, 855, 857, 861, 864, 869, 875, 876, 876, 879, 880, 880, 881, 882, 885, 886, 886, 887, 898, 902, 902, 908, 908, 910, 912, 924, 924, 933, 934, 935, 942, 942, 942, 944, 945, 950, 951, 952, 958, 961, 962, 962, 962, 962, 963, 963, 963, 967, 968, 970, 970, 971, 973, 974, 979, 980, 982, 983, 991, 992, 993, 994},
	},

	{
		[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
		[]int{2, 1, 4, 3, 9, 6},
		[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
	},

	// 可以有多个 testcase
}

func Test_relativeSortArray(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, relativeSortArray(tc.arr1, tc.arr2), "输入:%v", tc)
	}
}

func Benchmark_relativeSortArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			relativeSortArray(tc.arr1, tc.arr2)
		}
	}
}
