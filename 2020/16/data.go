package main

import "fmt"

type intv struct {
	low, high int
}

func (iv *intv) String() string {
	return fmt.Sprintf("[%v-%v]", iv.low, iv.high)
}

func (iv *intv) isValid(v int) bool {
	return v >= iv.low && v <= iv.high
}

var sample2Intvs = map[string][]*intv{
	"class": {{0, 1}, {4, 19}},
	"row":   {{0, 5}, {8, 19}},
	"seat":  {{0, 13}, {16, 19}},
}

// your ticket:
// 11,12,13

var sample2Tickets = [][]int{
	{3, 9, 18},
	{15, 1, 5},
	{5, 14, 9},
}

var sampleIntvs = map[string][]*intv{
	"class": {{1, 3}, {5, 7}},
	"row":   {{6, 11}, {33, 44}},
	"seat":  {{13, 40}, {45, 50}},
}

var sampleTickets = [][]int{
	{7, 3, 47},
	{40, 4, 50},
	{55, 2, 20},
	{38, 6, 12},
}

var inputIntvs = map[string][]*intv{
	"departure location": {{33, 224}, {230, 954}},
	"departure station":  {{32, 358}, {371, 974}},
	"departure platform": {{42, 411}, {417, 967}},
	"departure track":    {{30, 323}, {330, 964}},
	"departure date":     {{37, 608}, {624, 953}},
	"departure time":     {{43, 838}, {848, 954}},
	"arrival location":   {{40, 104}, {111, 955}},
	"arrival station":    {{43, 301}, {309, 961}},
	"arrival platform":   {{45, 253}, {275, 964}},
	"arrival track":      {{28, 79}, {97, 973}},
	"class":              {{31, 920}, {944, 960}},
	"duration":           {{35, 664}, {676, 951}},
	"price":              {{35, 499}, {521, 949}},
	"route":              {{38, 276}, {295, 974}},
	"row":                {{38, 582}, {599, 950}},
	"seat":               {{31, 646}, {657, 953}},
	"train":              {{40, 864}, {878, 955}},
	"type":               {{33, 430}, {443, 961}},
	"wagon":              {{30, 773}, {790, 956}},
	"zone":               {{48, 148}, {172, 962}},
}

var myTicket = []int{61, 101, 131, 127, 103, 191, 67, 181, 79, 71, 113, 97, 173, 59, 73, 137, 139, 53, 193, 179}

var inputTickets = [][]int{
	{390, 567, 702, 704, 825, 194, 543, 17, 472, 540, 687, 230, 235, 864, 884, 692, 375, 206, 920, 806},
	{339, 763, 628, 299, 191, 627, 793, 483, 645, 76, 731, 675, 172, 729, 945, 713, 350, 918, 720, 393},
	{206, 905, 97, 345, 392, 882, 768, 446, 924, 541, 145, 493, 101, 773, 232, 50, 635, 638, 499, 471},
	{352, 306, 382, 718, 338, 848, 944, 658, 646, 825, 754, 884, 711, 221, 146, 396, 140, 103, 894, 138},
	{419, 542, 487, 78, 66, 181, 424, 760, 490, 718, 664, 819, 527, 189, 327, 683, 57, 878, 703, 203},
	{71, 421, 734, 233, 536, 702, 77, 642, 445, 812, 335, 720, 735, 218, 936, 348, 352, 678, 632, 748},
	{427, 488, 681, 544, 495, 324, 58, 352, 295, 138, 731, 183, 212, 491, 660, 484, 79, 912, 636, 529},
	{568, 478, 339, 456, 702, 628, 56, 572, 528, 813, 708, 628, 133, 307, 633, 491, 900, 572, 911, 122},
	{707, 561, 811, 348, 790, 122, 189, 457, 496, 911, 812, 836, 350, 568, 346, 925, 555, 807, 724, 729},
	{448, 445, 269, 549, 351, 638, 678, 661, 570, 230, 709, 559, 222, 699, 240, 810, 479, 398, 464, 343},
	{736, 677, 572, 761, 769, 424, 629, 204, 681, 886, 542, 714, 317, 452, 811, 174, 220, 597, 556, 901},
	{472, 657, 448, 181, 397, 219, 810, 342, 312, 103, 355, 723, 632, 326, 346, 769, 67, 346, 601, 422},
	{852, 907, 878, 575, 59, 661, 816, 120, 531, 243, 831, 636, 548, 425, 875, 691, 61, 686, 184, 730},
	{397, 871, 757, 320, 118, 236, 474, 792, 197, 301, 309, 738, 99, 626, 579, 548, 756, 529, 600, 816},
	{644, 230, 983, 233, 658, 851, 891, 578, 460, 388, 127, 418, 560, 117, 197, 242, 533, 471, 582, 552},
	{99, 744, 134, 195, 944, 213, 207, 913, 920, 537, 705, 708, 750, 198, 916, 133, 900, 127, 859, 588},
	{559, 470, 495, 721, 568, 64, 722, 891, 449, 595, 533, 705, 535, 639, 902, 863, 524, 174, 471, 555},
	{680, 276, 404, 860, 447, 672, 128, 456, 696, 321, 173, 373, 117, 315, 429, 311, 708, 759, 801, 405},
	{663, 677, 212, 222, 6, 51, 736, 804, 537, 537, 447, 568, 420, 53, 74, 230, 886, 815, 115, 463},
	{577, 410, 556, 793, 799, 417, 455, 499, 444, 678, 295, 405, 917, 5, 221, 890, 720, 736, 605, 409},
	{375, 752, 477, 658, 495, 230, 851, 535, 618, 768, 381, 422, 531, 910, 699, 211, 136, 189, 906, 448},
	{187, 121, 322, 347, 60, 545, 355, 402, 818, 97, 607, 801, 797, 656, 795, 247, 713, 809, 902, 65},
	{129, 127, 899, 903, 608, 241, 245, 818, 388, 554, 525, 693, 578, 380, 982, 886, 856, 134, 135, 754},
	{681, 828, 392, 234, 313, 390, 706, 797, 752, 828, 495, 763, 229, 523, 497, 380, 571, 848, 820, 746},
	{234, 144, 384, 828, 99, 602, 405, 204, 854, 558, 815, 949, 338, 524, 229, 739, 692, 862, 398, 376},
	{825, 414, 543, 466, 906, 702, 710, 499, 400, 145, 319, 748, 497, 117, 682, 201, 191, 836, 570, 948},
	{448, 678, 417, 445, 729, 536, 236, 676, 275, 831, 663, 221, 421, 531, 843, 252, 726, 404, 762, 681},
	{558, 102, 448, 317, 206, 388, 253, 639, 630, 632, 632, 339, 400, 125, 23, 850, 189, 561, 449, 470},
	{528, 561, 250, 877, 182, 55, 704, 102, 571, 457, 337, 56, 750, 320, 821, 762, 727, 406, 914, 111},
	{271, 555, 112, 804, 885, 147, 218, 127, 300, 494, 693, 139, 185, 372, 892, 57, 713, 944, 202, 563},
	{569, 468, 851, 493, 851, 188, 239, 772, 669, 905, 807, 727, 387, 747, 494, 679, 948, 601, 691, 488},
	{184, 12, 449, 698, 491, 61, 76, 554, 733, 236, 823, 323, 604, 298, 50, 662, 917, 196, 350, 566},
	{426, 769, 638, 904, 635, 587, 393, 791, 724, 643, 708, 607, 913, 528, 73, 340, 309, 232, 322, 348},
	{383, 56, 627, 544, 296, 732, 556, 728, 321, 819, 63, 911, 482, 336, 364, 130, 388, 60, 67, 301},
	{625, 722, 731, 228, 394, 688, 185, 99, 320, 860, 527, 97, 313, 641, 549, 132, 75, 945, 70, 522},
	{54, 67, 322, 315, 56, 578, 301, 323, 720, 745, 443, 692, 204, 879, 832, 743, 242, 9, 917, 546},
	{794, 202, 628, 393, 806, 715, 301, 472, 486, 229, 466, 131, 731, 572, 184, 231, 221, 247, 632, 741},
	{836, 522, 945, 296, 529, 479, 449, 275, 559, 474, 121, 878, 268, 688, 849, 331, 499, 395, 194, 578},
	{66, 132, 317, 380, 710, 745, 231, 429, 693, 495, 352, 896, 794, 64, 719, 842, 605, 419, 172, 804},
	{834, 295, 147, 631, 765, 330, 946, 78, 68, 656, 772, 122, 238, 827, 480, 912, 176, 722, 492, 390},
	{276, 748, 58, 461, 736, 825, 691, 704, 116, 357, 608, 892, 876, 744, 834, 698, 78, 103, 576, 602},
	{688, 491, 191, 429, 908, 376, 483, 10, 234, 485, 639, 601, 864, 475, 661, 253, 331, 189, 492, 174},
	{344, 476, 717, 219, 385, 740, 321, 947, 690, 131, 946, 318, 605, 689, 99, 111, 17, 235, 398, 410},
	{916, 179, 542, 64, 495, 211, 187, 197, 747, 548, 338, 532, 181, 844, 551, 51, 187, 663, 493, 887},
	{203, 120, 857, 188, 847, 817, 715, 353, 657, 125, 143, 444, 903, 358, 213, 556, 381, 53, 737, 418},
	{792, 206, 582, 69, 749, 129, 812, 764, 326, 406, 707, 637, 906, 398, 61, 178, 253, 878, 895, 836},
	{542, 946, 524, 498, 688, 205, 581, 452, 799, 716, 898, 979, 579, 480, 471, 536, 447, 183, 888, 407},
	{124, 549, 822, 311, 389, 793, 828, 530, 445, 145, 810, 129, 252, 348, 115, 367, 62, 527, 911, 817},
	{645, 147, 198, 994, 883, 545, 111, 660, 663, 447, 767, 475, 822, 631, 793, 136, 660, 253, 733, 375},
	{690, 211, 195, 354, 366, 455, 73, 172, 209, 852, 252, 772, 316, 197, 799, 409, 791, 718, 582, 946},
	{382, 914, 340, 696, 547, 683, 462, 335, 65, 202, 211, 539, 564, 990, 861, 733, 113, 336, 119, 251},
	{465, 69, 471, 766, 211, 61, 384, 347, 180, 421, 179, 201, 915, 93, 947, 120, 746, 190, 727, 680},
	{571, 893, 319, 522, 249, 805, 457, 534, 860, 862, 646, 354, 728, 406, 67, 611, 314, 747, 681, 465},
	{237, 468, 751, 714, 692, 551, 641, 351, 547, 851, 478, 538, 737, 815, 224, 675, 411, 793, 205, 102},
	{660, 54, 550, 422, 211, 864, 497, 686, 494, 637, 319, 388, 881, 819, 72, 884, 813, 651, 863, 534},
	{201, 496, 537, 468, 412, 175, 114, 337, 829, 137, 79, 406, 386, 427, 139, 342, 275, 879, 183, 129},
	{571, 466, 892, 194, 50, 768, 178, 97, 419, 713, 947, 214, 880, 54, 685, 833, 527, 329, 736, 467},
	{919, 348, 684, 143, 897, 805, 411, 555, 208, 207, 664, 373, 310, 667, 425, 812, 813, 235, 462, 538},
	{179, 194, 882, 220, 434, 795, 379, 485, 462, 146, 479, 216, 62, 207, 563, 338, 486, 319, 754, 820},
	{637, 698, 388, 812, 705, 423, 701, 252, 803, 529, 534, 120, 545, 366, 113, 199, 315, 119, 195, 601},
	{799, 58, 852, 443, 793, 825, 142, 694, 58, 581, 459, 136, 448, 900, 416, 375, 65, 765, 646, 235},
	{54, 825, 728, 918, 663, 379, 856, 562, 259, 71, 67, 176, 899, 720, 799, 132, 532, 899, 903, 97},
	{568, 449, 711, 772, 590, 235, 749, 387, 75, 72, 399, 401, 411, 428, 707, 317, 552, 525, 58, 347},
	{717, 485, 696, 483, 862, 229, 899, 796, 141, 749, 231, 317, 465, 800, 530, 67, 836, 98, 197, 599},
	{632, 459, 542, 318, 95, 63, 634, 388, 803, 900, 555, 756, 50, 752, 574, 219, 628, 491, 902, 352},
	{911, 757, 708, 66, 102, 549, 762, 101, 249, 617, 849, 824, 643, 544, 447, 765, 491, 132, 104, 893},
	{180, 187, 721, 894, 53, 7, 241, 425, 645, 815, 396, 126, 176, 542, 377, 767, 897, 727, 822, 323},
	{491, 206, 122, 133, 753, 118, 714, 119, 77, 740, 298, 131, 839, 417, 275, 193, 56, 449, 657, 489},
	{853, 712, 315, 706, 945, 298, 197, 724, 750, 873, 339, 119, 461, 405, 356, 443, 816, 144, 602, 449},
	{550, 631, 421, 559, 353, 443, 920, 791, 728, 357, 181, 872, 239, 497, 460, 462, 543, 334, 628, 818},
	{494, 344, 765, 356, 344, 450, 237, 894, 656, 387, 827, 230, 135, 117, 539, 311, 112, 135, 387, 809},
	{210, 848, 253, 842, 894, 421, 739, 638, 483, 315, 189, 550, 136, 210, 74, 145, 626, 455, 378, 882},
	{450, 127, 589, 636, 855, 221, 631, 388, 838, 853, 523, 910, 50, 353, 117, 691, 864, 333, 231, 642},
	{480, 759, 647, 795, 599, 410, 77, 180, 820, 176, 316, 848, 816, 130, 332, 474, 497, 879, 247, 817},
	{226, 711, 555, 209, 558, 234, 175, 124, 718, 331, 799, 531, 559, 213, 825, 560, 761, 758, 409, 791},
	{322, 864, 67, 607, 812, 421, 741, 556, 357, 854, 381, 219, 835, 527, 247, 599, 76, 101, 297, 977},
	{72, 491, 368, 138, 144, 543, 208, 190, 237, 445, 748, 78, 856, 51, 404, 389, 534, 201, 629, 827},
	{349, 831, 336, 131, 320, 119, 700, 572, 55, 148, 18, 848, 810, 243, 762, 822, 548, 626, 540, 457},
	{114, 764, 719, 884, 659, 450, 350, 769, 797, 346, 910, 910, 752, 196, 323, 481, 765, 303, 376, 378},
	{55, 605, 190, 466, 350, 857, 817, 213, 640, 862, 11, 192, 734, 759, 72, 751, 425, 822, 484, 575},
	{200, 321, 206, 586, 576, 678, 685, 832, 54, 341, 830, 172, 459, 479, 71, 122, 546, 200, 201, 53},
	{798, 764, 348, 201, 129, 533, 212, 103, 176, 200, 145, 758, 114, 199, 132, 71, 707, 614, 691, 749},
	{457, 317, 989, 341, 657, 582, 566, 246, 898, 182, 806, 715, 659, 450, 576, 349, 463, 747, 74, 475},
	{816, 822, 869, 687, 198, 382, 746, 208, 864, 119, 659, 906, 882, 123, 703, 422, 250, 55, 309, 681},
	{705, 181, 389, 137, 918, 305, 545, 484, 59, 98, 663, 801, 392, 689, 449, 465, 479, 428, 683, 219},
	{401, 557, 667, 771, 752, 681, 355, 555, 385, 218, 497, 497, 63, 461, 882, 606, 897, 403, 236, 97},
	{885, 530, 241, 827, 208, 313, 722, 174, 948, 403, 772, 480, 862, 984, 600, 355, 602, 644, 126, 388},
	{405, 348, 556, 267, 473, 754, 334, 945, 683, 770, 543, 497, 542, 251, 98, 334, 123, 705, 879, 526},
	{303, 378, 637, 636, 646, 211, 899, 914, 662, 130, 554, 748, 392, 276, 902, 944, 206, 296, 790, 707},
	{883, 676, 600, 127, 253, 345, 596, 860, 678, 191, 249, 447, 116, 901, 685, 624, 645, 714, 862, 377},
	{492, 63, 932, 854, 563, 646, 344, 833, 715, 809, 392, 173, 121, 728, 606, 390, 189, 769, 731, 425},
	{1, 420, 579, 55, 345, 833, 184, 429, 317, 853, 541, 539, 770, 52, 714, 356, 539, 386, 827, 113},
	{231, 946, 317, 859, 479, 246, 638, 357, 542, 358, 492, 764, 880, 445, 275, 179, 470, 242, 915, 670},
	{209, 738, 657, 890, 357, 711, 422, 678, 631, 190, 383, 680, 604, 704, 642, 511, 384, 529, 830, 66},
	{400, 749, 478, 915, 984, 487, 690, 864, 895, 193, 183, 142, 444, 897, 749, 185, 909, 389, 680, 580},
	{193, 732, 187, 212, 548, 353, 856, 498, 472, 558, 910, 460, 459, 374, 722, 882, 874, 383, 690, 773},
	{335, 607, 899, 888, 349, 712, 129, 249, 22, 822, 357, 410, 739, 334, 465, 748, 860, 657, 351, 602},
	{859, 643, 202, 299, 75, 900, 800, 540, 227, 806, 466, 392, 197, 834, 300, 637, 241, 78, 68, 882},
	{809, 299, 64, 729, 994, 313, 795, 207, 769, 886, 705, 705, 879, 579, 735, 479, 323, 814, 448, 495},
	{906, 386, 572, 401, 698, 886, 196, 403, 479, 477, 571, 147, 530, 176, 675, 630, 898, 737, 602, 811},
	{572, 125, 59, 377, 881, 250, 573, 746, 590, 116, 561, 71, 880, 544, 148, 852, 462, 910, 689, 661},
	{253, 98, 849, 345, 219, 173, 317, 601, 307, 522, 221, 849, 771, 331, 744, 371, 530, 582, 334, 735},
	{534, 218, 378, 890, 75, 574, 692, 817, 253, 208, 147, 761, 302, 757, 735, 340, 898, 604, 482, 562},
	{560, 734, 316, 147, 907, 137, 376, 428, 808, 52, 248, 883, 815, 701, 317, 353, 244, 279, 403, 132},
	{457, 337, 347, 366, 209, 253, 483, 411, 50, 411, 911, 830, 463, 136, 790, 821, 657, 645, 725, 64},
	{822, 65, 756, 692, 409, 493, 177, 454, 887, 397, 894, 624, 444, 427, 387, 457, 319, 758, 773, 306},
	{50, 197, 200, 327, 336, 316, 239, 300, 454, 378, 858, 133, 858, 410, 773, 395, 720, 476, 907, 948},
	{558, 103, 604, 495, 429, 760, 550, 458, 362, 209, 380, 65, 498, 558, 335, 747, 855, 771, 770, 581},
	{297, 582, 769, 208, 751, 791, 803, 759, 658, 738, 818, 684, 298, 143, 338, 652, 554, 693, 757, 575},
	{523, 119, 751, 412, 893, 221, 899, 710, 358, 695, 542, 916, 392, 747, 811, 317, 460, 418, 679, 348},
	{574, 409, 535, 216, 668, 241, 132, 557, 498, 637, 372, 817, 243, 707, 901, 295, 608, 524, 823, 878},
	{391, 67, 101, 63, 600, 634, 186, 657, 891, 719, 423, 224, 603, 240, 178, 334, 386, 15, 554, 800},
	{62, 497, 767, 533, 486, 198, 661, 341, 216, 802, 824, 52, 887, 643, 483, 225, 691, 947, 608, 800},
	{178, 681, 529, 638, 602, 873, 406, 907, 210, 240, 566, 540, 483, 350, 819, 580, 77, 714, 738, 759},
	{330, 825, 566, 852, 548, 66, 103, 481, 664, 763, 143, 312, 639, 396, 443, 354, 835, 912, 426, 647},
	{562, 325, 187, 684, 794, 813, 104, 608, 444, 576, 447, 905, 485, 599, 354, 189, 605, 486, 552, 399},
	{662, 296, 538, 488, 481, 429, 795, 230, 822, 444, 838, 351, 174, 173, 896, 215, 977, 906, 235, 276},
	{384, 687, 406, 693, 335, 391, 357, 71, 691, 577, 14, 580, 754, 210, 749, 448, 822, 829, 336, 137},
	{998, 688, 838, 117, 192, 808, 463, 353, 561, 410, 744, 133, 194, 812, 403, 752, 532, 221, 182, 140},
	{949, 462, 754, 242, 802, 116, 663, 460, 106, 819, 137, 209, 233, 330, 535, 542, 253, 563, 496, 498},
	{907, 351, 542, 183, 404, 72, 389, 532, 562, 740, 856, 863, 415, 111, 911, 825, 403, 475, 636, 379},
	{890, 139, 398, 629, 755, 187, 829, 949, 299, 475, 481, 907, 572, 456, 884, 210, 899, 426, 220, 868},
	{340, 813, 720, 335, 602, 388, 906, 183, 707, 202, 143, 546, 806, 200, 66, 689, 759, 987, 646, 858},
	{867, 581, 569, 176, 717, 686, 757, 187, 704, 758, 541, 231, 733, 373, 749, 790, 448, 183, 380, 646},
	{144, 560, 680, 615, 853, 855, 915, 578, 470, 477, 754, 827, 579, 746, 746, 535, 56, 640, 916, 137},
	{138, 486, 618, 709, 765, 693, 684, 880, 553, 807, 58, 914, 444, 920, 906, 353, 338, 339, 353, 889},
	{663, 336, 114, 558, 975, 470, 396, 358, 600, 197, 801, 430, 731, 707, 338, 722, 832, 627, 852, 521},
	{397, 751, 662, 426, 810, 526, 358, 411, 702, 628, 762, 823, 141, 912, 489, 99, 711, 994, 244, 764},
	{447, 695, 521, 762, 227, 560, 576, 190, 813, 246, 202, 399, 580, 759, 476, 689, 473, 553, 126, 694},
	{487, 578, 71, 802, 737, 920, 550, 60, 745, 791, 141, 333, 900, 989, 539, 773, 607, 338, 295, 892},
	{818, 636, 742, 696, 476, 181, 795, 300, 493, 795, 809, 127, 493, 805, 305, 681, 56, 899, 882, 465},
	{393, 464, 913, 403, 878, 573, 212, 696, 147, 919, 829, 602, 558, 473, 766, 709, 386, 942, 444, 353},
	{550, 699, 905, 879, 625, 813, 350, 472, 122, 197, 419, 652, 698, 527, 102, 704, 554, 252, 792, 390},
	{342, 465, 476, 767, 296, 475, 977, 821, 900, 185, 341, 172, 948, 697, 377, 219, 812, 552, 348, 918},
	{241, 315, 307, 680, 102, 554, 709, 824, 910, 700, 235, 459, 901, 602, 878, 564, 231, 791, 534, 815},
	{103, 855, 248, 250, 459, 144, 717, 469, 338, 97, 348, 397, 218, 323, 300, 329, 223, 344, 553, 132},
	{66, 488, 635, 461, 539, 756, 447, 718, 737, 220, 707, 422, 636, 534, 475, 594, 836, 816, 742, 246},
	{138, 402, 602, 212, 444, 258, 766, 138, 102, 816, 216, 713, 914, 354, 178, 522, 803, 199, 481, 770},
	{568, 411, 201, 479, 330, 829, 189, 715, 582, 460, 696, 790, 817, 638, 321, 267, 356, 749, 345, 727},
	{816, 135, 558, 455, 472, 653, 531, 384, 599, 716, 764, 690, 64, 395, 301, 529, 427, 334, 705, 425},
	{919, 143, 536, 345, 580, 335, 119, 712, 453, 265, 739, 470, 224, 581, 250, 659, 792, 353, 683, 242},
	{645, 722, 147, 404, 646, 815, 818, 147, 316, 190, 791, 373, 529, 554, 653, 695, 478, 389, 697, 59},
	{341, 711, 193, 392, 892, 771, 696, 309, 415, 464, 825, 148, 723, 394, 677, 465, 797, 371, 377, 579},
	{849, 409, 692, 179, 905, 139, 129, 666, 124, 627, 704, 883, 914, 126, 899, 134, 713, 345, 119, 572},
	{184, 190, 412, 884, 848, 395, 659, 736, 894, 605, 688, 761, 126, 895, 823, 629, 551, 206, 375, 765},
	{827, 481, 94, 638, 795, 628, 856, 449, 727, 708, 218, 825, 50, 103, 199, 694, 344, 313, 581, 920},
	{123, 598, 602, 746, 773, 858, 111, 627, 945, 878, 349, 813, 321, 340, 885, 295, 886, 192, 639, 636},
	{750, 536, 566, 71, 186, 444, 498, 408, 129, 419, 723, 734, 420, 275, 542, 305, 729, 829, 582, 743},
	{378, 58, 531, 882, 711, 803, 50, 527, 238, 540, 905, 575, 452, 744, 689, 121, 711, 665, 749, 707},
	{684, 764, 979, 237, 402, 644, 945, 145, 894, 827, 860, 372, 527, 857, 473, 429, 188, 792, 346, 456},
	{898, 321, 708, 822, 238, 202, 813, 451, 497, 717, 341, 400, 140, 405, 418, 350, 671, 581, 573, 443},
	{211, 183, 187, 223, 230, 888, 676, 573, 751, 675, 578, 736, 884, 201, 445, 318, 542, 686, 312, 127},
	{757, 485, 831, 796, 651, 625, 765, 638, 625, 724, 173, 565, 315, 64, 181, 219, 727, 736, 147, 698},
	{561, 116, 624, 945, 874, 571, 573, 563, 468, 340, 181, 301, 322, 491, 111, 353, 904, 129, 489, 299},
	{424, 821, 179, 236, 677, 252, 606, 689, 116, 836, 126, 714, 495, 380, 908, 282, 379, 754, 309, 475},
	{63, 572, 639, 238, 629, 68, 296, 640, 844, 677, 320, 569, 533, 53, 494, 882, 699, 685, 476, 801},
	{760, 768, 230, 545, 314, 460, 855, 548, 526, 302, 72, 198, 849, 525, 736, 399, 640, 297, 758, 373},
	{352, 248, 343, 698, 528, 792, 219, 466, 330, 766, 424, 73, 637, 606, 454, 412, 808, 56, 534, 101},
	{901, 478, 116, 240, 348, 353, 570, 531, 472, 685, 493, 900, 253, 994, 187, 186, 64, 400, 119, 142},
	{117, 634, 818, 375, 192, 612, 464, 119, 213, 916, 680, 373, 710, 686, 678, 451, 914, 245, 249, 704},
	{663, 642, 465, 346, 678, 137, 220, 837, 849, 477, 878, 189, 805, 876, 909, 133, 745, 249, 407, 849},
	{337, 541, 409, 802, 891, 246, 79, 802, 528, 741, 690, 526, 767, 400, 141, 13, 768, 398, 659, 719},
	{334, 904, 161, 206, 860, 657, 379, 238, 310, 312, 832, 493, 488, 466, 400, 559, 947, 134, 645, 564},
	{854, 635, 835, 423, 540, 275, 486, 561, 637, 730, 671, 140, 910, 773, 76, 733, 424, 354, 248, 805},
	{136, 827, 227, 624, 456, 913, 738, 639, 98, 862, 401, 606, 67, 748, 802, 700, 903, 122, 911, 887},
	{486, 79, 802, 632, 605, 981, 123, 734, 489, 174, 815, 204, 216, 689, 485, 632, 797, 804, 61, 824},
	{878, 196, 420, 495, 106, 443, 346, 898, 712, 330, 580, 555, 604, 449, 797, 118, 495, 739, 79, 716},
	{195, 79, 387, 911, 284, 203, 469, 564, 525, 498, 731, 74, 750, 59, 561, 605, 465, 421, 185, 221},
	{552, 206, 138, 857, 444, 773, 793, 430, 707, 817, 295, 854, 329, 343, 837, 190, 894, 728, 682, 249},
	{358, 224, 687, 742, 816, 51, 409, 857, 757, 275, 196, 317, 645, 609, 443, 864, 574, 417, 763, 207},
	{485, 465, 735, 712, 629, 223, 731, 718, 146, 376, 864, 634, 237, 113, 58, 637, 595, 822, 795, 126},
	{793, 297, 679, 687, 326, 63, 854, 371, 915, 458, 354, 402, 716, 248, 179, 338, 919, 142, 195, 336},
	{672, 755, 445, 490, 187, 323, 491, 479, 461, 75, 371, 830, 352, 189, 128, 395, 626, 726, 887, 795},
	{61, 264, 772, 637, 553, 351, 859, 446, 253, 702, 235, 469, 755, 859, 729, 807, 69, 331, 643, 343},
	{999, 121, 550, 736, 474, 238, 809, 72, 863, 577, 392, 451, 719, 485, 74, 450, 51, 78, 733, 391},
	{422, 177, 769, 904, 931, 632, 546, 199, 545, 863, 295, 635, 239, 729, 133, 250, 350, 913, 691, 321},
	{662, 300, 480, 173, 269, 243, 252, 807, 678, 450, 765, 548, 714, 489, 736, 469, 187, 803, 556, 662},
	{313, 208, 732, 745, 772, 906, 78, 60, 400, 137, 860, 886, 917, 187, 529, 865, 945, 486, 826, 333},
	{131, 540, 560, 405, 730, 447, 353, 352, 145, 351, 477, 946, 585, 73, 564, 116, 762, 67, 397, 641},
	{424, 465, 573, 826, 794, 318, 78, 342, 240, 386, 338, 450, 546, 260, 580, 629, 799, 381, 495, 211},
	{858, 905, 63, 485, 614, 808, 602, 702, 709, 730, 315, 581, 553, 730, 567, 748, 576, 221, 566, 137},
	{57, 909, 142, 580, 831, 316, 803, 900, 290, 690, 60, 409, 356, 332, 371, 699, 827, 848, 857, 563},
	{449, 445, 834, 574, 756, 422, 67, 540, 95, 698, 339, 77, 483, 556, 910, 747, 800, 920, 114, 299},
	{102, 379, 68, 419, 740, 538, 332, 553, 642, 532, 451, 54, 116, 311, 145, 312, 299, 866, 582, 474},
	{391, 404, 126, 420, 138, 720, 796, 141, 372, 524, 421, 755, 496, 837, 634, 182, 692, 229, 187, 688},
	{655, 205, 451, 122, 219, 887, 125, 828, 856, 555, 522, 126, 251, 557, 912, 805, 852, 394, 891, 332},
	{100, 230, 689, 981, 352, 540, 493, 56, 342, 755, 174, 945, 234, 820, 857, 276, 823, 124, 603, 335},
	{894, 224, 103, 525, 498, 406, 143, 318, 528, 21, 684, 641, 192, 521, 707, 348, 852, 356, 384, 658},
	{580, 228, 213, 910, 945, 466, 124, 763, 533, 199, 732, 235, 376, 769, 348, 454, 69, 203, 634, 796},
	{316, 819, 379, 312, 430, 75, 339, 378, 733, 209, 137, 408, 189, 585, 857, 884, 646, 737, 728, 335},
	{311, 397, 54, 311, 124, 642, 209, 827, 546, 180, 749, 385, 613, 98, 205, 636, 566, 445, 578, 714},
	{566, 895, 183, 580, 919, 310, 903, 626, 59, 187, 729, 677, 114, 400, 595, 680, 817, 490, 723, 606},
	{232, 184, 679, 52, 658, 211, 103, 827, 176, 183, 681, 421, 405, 752, 383, 639, 257, 898, 475, 222},
	{451, 726, 944, 813, 132, 760, 380, 904, 142, 919, 681, 635, 647, 728, 474, 375, 420, 310, 699, 690},
	{918, 633, 147, 849, 822, 330, 629, 947, 804, 630, 487, 394, 559, 636, 14, 686, 410, 247, 713, 458},
	{724, 142, 234, 735, 889, 944, 314, 759, 875, 743, 181, 133, 903, 897, 309, 349, 761, 773, 295, 481},
	{700, 316, 51, 819, 429, 915, 176, 730, 148, 551, 981, 750, 249, 131, 521, 947, 554, 250, 819, 75},
	{567, 761, 371, 897, 475, 451, 318, 601, 790, 632, 111, 836, 739, 859, 356, 484, 468, 80, 242, 354},
	{542, 454, 246, 247, 542, 767, 317, 137, 384, 479, 51, 426, 183, 97, 607, 700, 113, 158, 461, 340},
	{860, 853, 830, 688, 545, 320, 732, 74, 130, 539, 712, 331, 747, 415, 527, 240, 425, 479, 708, 521},
	{582, 239, 298, 50, 307, 699, 346, 380, 313, 526, 561, 604, 350, 919, 403, 342, 429, 815, 640, 320},
	{319, 698, 136, 805, 687, 552, 187, 993, 606, 190, 800, 103, 197, 543, 628, 704, 761, 389, 642, 575},
	{710, 863, 604, 383, 630, 175, 463, 640, 603, 389, 606, 348, 676, 913, 482, 395, 16, 570, 69, 223},
	{174, 332, 382, 851, 835, 62, 250, 521, 639, 892, 913, 497, 600, 479, 123, 886, 131, 723, 334, 984},
	{757, 524, 841, 735, 748, 446, 400, 490, 731, 628, 483, 814, 560, 769, 633, 677, 538, 190, 61, 401},
	{752, 655, 628, 832, 658, 392, 71, 920, 482, 115, 713, 638, 792, 444, 499, 905, 633, 487, 123, 496},
	{545, 523, 766, 791, 749, 212, 213, 820, 798, 317, 875, 745, 135, 920, 74, 221, 172, 693, 854, 385},
	{182, 394, 326, 525, 195, 798, 68, 202, 297, 768, 701, 700, 374, 492, 124, 827, 450, 142, 221, 758},
	{426, 523, 485, 655, 317, 141, 800, 806, 862, 453, 602, 379, 481, 445, 202, 890, 235, 883, 346, 696},
	{276, 920, 417, 540, 339, 570, 465, 827, 574, 689, 299, 269, 62, 664, 126, 565, 710, 71, 377, 253},
	{309, 346, 209, 761, 125, 552, 394, 582, 918, 490, 404, 736, 72, 830, 634, 399, 743, 846, 69, 700},
	{214, 390, 887, 794, 854, 240, 121, 657, 641, 74, 181, 574, 135, 232, 564, 691, 206, 264, 217, 319},
	{448, 424, 644, 746, 388, 850, 522, 690, 718, 634, 750, 55, 241, 773, 714, 109, 600, 223, 635, 320},
	{757, 850, 752, 146, 347, 915, 732, 243, 161, 605, 644, 714, 753, 546, 377, 658, 204, 808, 753, 896},
	{903, 524, 801, 349, 333, 55, 333, 496, 129, 879, 253, 596, 56, 487, 677, 398, 127, 323, 946, 746},
	{396, 142, 858, 793, 391, 449, 321, 236, 862, 548, 598, 864, 383, 575, 202, 342, 140, 766, 690, 224},
	{317, 622, 882, 443, 246, 355, 53, 219, 850, 60, 177, 759, 522, 693, 53, 576, 113, 817, 561, 350},
	{217, 490, 523, 122, 856, 447, 492, 429, 529, 748, 64, 765, 555, 573, 793, 753, 640, 676, 997, 540},
	{810, 688, 691, 698, 726, 484, 768, 358, 681, 344, 707, 886, 977, 860, 819, 948, 739, 194, 640, 809},
	{73, 51, 555, 737, 460, 179, 674, 563, 886, 66, 410, 702, 99, 720, 471, 628, 538, 729, 737, 423},
	{739, 643, 801, 417, 72, 911, 663, 276, 740, 854, 497, 999, 538, 102, 394, 664, 209, 241, 446, 345},
	{909, 357, 387, 178, 214, 898, 411, 407, 241, 295, 914, 636, 895, 227, 470, 734, 743, 919, 79, 894},
	{224, 456, 386, 125, 177, 397, 114, 343, 791, 709, 71, 903, 443, 276, 610, 629, 403, 802, 74, 486},
	{219, 247, 352, 457, 911, 338, 398, 394, 71, 148, 662, 130, 0, 77, 702, 713, 391, 658, 718, 803},
	{252, 239, 904, 802, 535, 74, 521, 722, 471, 481, 146, 860, 552, 755, 175, 55, 211, 859, 147, 261},
	{803, 67, 311, 665, 525, 526, 346, 221, 795, 753, 444, 553, 103, 311, 733, 200, 856, 769, 745, 537},
	{646, 568, 426, 310, 918, 334, 54, 862, 238, 761, 904, 769, 238, 240, 12, 194, 429, 731, 630, 224},
	{910, 529, 802, 890, 354, 680, 455, 637, 207, 68, 493, 377, 371, 715, 198, 83, 410, 59, 864, 148},
	{558, 298, 212, 355, 907, 561, 553, 803, 600, 178, 692, 858, 819, 452, 255, 488, 534, 298, 816, 451},
	{72, 627, 175, 305, 462, 300, 555, 903, 833, 758, 51, 801, 380, 550, 371, 556, 888, 697, 50, 603},
	{914, 180, 889, 452, 475, 825, 496, 299, 480, 145, 475, 317, 577, 497, 175, 169, 601, 445, 145, 569},
	{472, 678, 249, 224, 736, 176, 345, 238, 231, 854, 549, 680, 71, 529, 67, 381, 316, 367, 827, 61},
	{245, 863, 570, 472, 68, 748, 205, 919, 463, 690, 78, 405, 540, 931, 659, 814, 556, 341, 685, 142},
	{400, 710, 484, 577, 100, 768, 127, 331, 582, 231, 662, 637, 900, 529, 173, 580, 901, 415, 97, 75},
	{731, 206, 553, 884, 71, 572, 729, 486, 217, 313, 194, 483, 949, 695, 472, 435, 749, 253, 371, 912},
	{204, 721, 601, 548, 161, 179, 567, 380, 879, 299, 428, 346, 540, 528, 458, 447, 719, 180, 608, 207},
	{477, 665, 52, 75, 352, 815, 485, 340, 125, 601, 911, 67, 690, 397, 805, 71, 425, 577, 408, 661},
	{590, 561, 221, 216, 244, 689, 703, 909, 769, 142, 564, 806, 418, 72, 564, 102, 659, 343, 420, 553},
	{828, 195, 444, 398, 221, 719, 237, 639, 945, 703, 407, 766, 852, 734, 410, 205, 849, 524, 465, 984},
	{182, 754, 180, 902, 676, 476, 664, 834, 729, 234, 378, 223, 667, 275, 682, 247, 350, 879, 244, 353},
	{711, 628, 810, 399, 387, 488, 197, 224, 686, 424, 70, 758, 8, 380, 68, 764, 853, 919, 448, 542},
}
