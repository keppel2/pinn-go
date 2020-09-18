// Code generated from Pinn.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Pinn

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 65, 374,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 51, 10, 2, 13, 2, 14, 2, 52, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 63, 10, 3, 12, 3, 14, 3, 66,
	11, 3, 5, 3, 68, 10, 3, 3, 3, 3, 3, 5, 3, 72, 10, 3, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 7, 5, 81, 10, 5, 12, 5, 14, 5, 84, 11, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 95, 10, 7, 3, 7,
	3, 7, 3, 7, 5, 7, 100, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 107,
	10, 8, 3, 8, 5, 8, 110, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 5, 9, 120, 10, 9, 3, 9, 5, 9, 123, 10, 9, 3, 10, 3, 10, 3, 10, 5,
	10, 128, 10, 10, 3, 10, 3, 10, 5, 10, 132, 10, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 5, 10, 140, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	5, 11, 178, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 185, 10,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 202, 10, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 222, 10, 12, 12, 12, 14, 12,
	225, 11, 12, 3, 13, 3, 13, 3, 13, 7, 13, 230, 10, 13, 12, 13, 14, 13, 233,
	11, 13, 3, 14, 3, 14, 5, 14, 237, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 5, 15, 246, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 19, 5, 19, 265, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 289, 10, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 7, 20, 295, 10, 20, 12, 20, 14, 20, 298, 11, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 7, 21, 304, 10, 21, 12, 21, 14, 21, 307, 11, 21, 3, 21,
	3, 21, 3, 21, 7, 21, 312, 10, 21, 12, 21, 14, 21, 315, 11, 21, 5, 21, 317,
	10, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 5, 22, 351, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5,
	23, 358, 10, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23,
	367, 10, 23, 3, 23, 3, 23, 3, 23, 5, 23, 372, 10, 23, 3, 23, 2, 3, 22,
	24, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 2, 8, 3, 2, 11, 12, 4, 2, 47, 47, 55, 55, 3, 2, 20, 23,
	5, 2, 20, 21, 23, 23, 58, 58, 3, 2, 24, 29, 3, 2, 30, 31, 2, 420, 2, 50,
	3, 2, 2, 2, 4, 56, 3, 2, 2, 2, 6, 75, 3, 2, 2, 2, 8, 78, 3, 2, 2, 2, 10,
	87, 3, 2, 2, 2, 12, 99, 3, 2, 2, 2, 14, 109, 3, 2, 2, 2, 16, 122, 3, 2,
	2, 2, 18, 139, 3, 2, 2, 2, 20, 177, 3, 2, 2, 2, 22, 201, 3, 2, 2, 2, 24,
	226, 3, 2, 2, 2, 26, 234, 3, 2, 2, 2, 28, 240, 3, 2, 2, 2, 30, 247, 3,
	2, 2, 2, 32, 252, 3, 2, 2, 2, 34, 256, 3, 2, 2, 2, 36, 288, 3, 2, 2, 2,
	38, 290, 3, 2, 2, 2, 40, 299, 3, 2, 2, 2, 42, 350, 3, 2, 2, 2, 44, 371,
	3, 2, 2, 2, 46, 51, 5, 4, 3, 2, 47, 48, 5, 12, 7, 2, 48, 49, 7, 3, 2, 2,
	49, 51, 3, 2, 2, 2, 50, 46, 3, 2, 2, 2, 50, 47, 3, 2, 2, 2, 51, 52, 3,
	2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54,
	55, 7, 2, 2, 3, 55, 3, 3, 2, 2, 2, 56, 57, 7, 4, 2, 2, 57, 58, 7, 60, 2,
	2, 58, 67, 7, 49, 2, 2, 59, 64, 5, 10, 6, 2, 60, 61, 7, 5, 2, 2, 61, 63,
	5, 10, 6, 2, 62, 60, 3, 2, 2, 2, 63, 66, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2,
	64, 65, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 59, 3,
	2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 71, 7, 6, 2, 2, 70,
	72, 5, 14, 8, 2, 71, 70, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 3, 2,
	2, 2, 73, 74, 5, 8, 5, 2, 74, 5, 3, 2, 2, 2, 75, 76, 7, 60, 2, 2, 76, 77,
	7, 60, 2, 2, 77, 7, 3, 2, 2, 2, 78, 82, 7, 7, 2, 2, 79, 81, 5, 42, 22,
	2, 80, 79, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82, 83,
	3, 2, 2, 2, 83, 85, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 86, 7, 8, 2, 2,
	86, 9, 3, 2, 2, 2, 87, 88, 7, 60, 2, 2, 88, 89, 5, 14, 8, 2, 89, 11, 3,
	2, 2, 2, 90, 91, 7, 60, 2, 2, 91, 94, 5, 14, 8, 2, 92, 93, 7, 9, 2, 2,
	93, 95, 5, 24, 13, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 100,
	3, 2, 2, 2, 96, 97, 7, 60, 2, 2, 97, 98, 7, 56, 2, 2, 98, 100, 5, 22, 12,
	2, 99, 90, 3, 2, 2, 2, 99, 96, 3, 2, 2, 2, 100, 13, 3, 2, 2, 2, 101, 106,
	7, 48, 2, 2, 102, 107, 7, 50, 2, 2, 103, 107, 7, 51, 2, 2, 104, 107, 7,
	52, 2, 2, 105, 107, 5, 22, 12, 2, 106, 102, 3, 2, 2, 2, 106, 103, 3, 2,
	2, 2, 106, 104, 3, 2, 2, 2, 106, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2,
	108, 110, 7, 10, 2, 2, 109, 101, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110,
	111, 3, 2, 2, 2, 111, 112, 7, 53, 2, 2, 112, 15, 3, 2, 2, 2, 113, 123,
	5, 44, 23, 2, 114, 119, 7, 60, 2, 2, 115, 116, 7, 48, 2, 2, 116, 117, 5,
	22, 12, 2, 117, 118, 7, 10, 2, 2, 118, 120, 3, 2, 2, 2, 119, 115, 3, 2,
	2, 2, 119, 120, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 123, 9, 2, 2, 2,
	122, 113, 3, 2, 2, 2, 122, 114, 3, 2, 2, 2, 123, 17, 3, 2, 2, 2, 124, 125,
	7, 60, 2, 2, 125, 127, 7, 48, 2, 2, 126, 128, 5, 22, 12, 2, 127, 126, 3,
	2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 131, 9, 3, 2,
	2, 130, 132, 5, 22, 12, 2, 131, 130, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2,
	132, 133, 3, 2, 2, 2, 133, 140, 7, 10, 2, 2, 134, 135, 7, 60, 2, 2, 135,
	136, 7, 48, 2, 2, 136, 137, 5, 22, 12, 2, 137, 138, 7, 10, 2, 2, 138, 140,
	3, 2, 2, 2, 139, 124, 3, 2, 2, 2, 139, 134, 3, 2, 2, 2, 140, 19, 3, 2,
	2, 2, 141, 142, 7, 13, 2, 2, 142, 143, 7, 49, 2, 2, 143, 144, 5, 24, 13,
	2, 144, 145, 7, 6, 2, 2, 145, 178, 3, 2, 2, 2, 146, 147, 7, 14, 2, 2, 147,
	148, 7, 49, 2, 2, 148, 149, 5, 22, 12, 2, 149, 150, 7, 6, 2, 2, 150, 178,
	3, 2, 2, 2, 151, 152, 7, 15, 2, 2, 152, 153, 7, 49, 2, 2, 153, 154, 5,
	22, 12, 2, 154, 155, 7, 6, 2, 2, 155, 178, 3, 2, 2, 2, 156, 157, 7, 16,
	2, 2, 157, 158, 7, 49, 2, 2, 158, 159, 7, 60, 2, 2, 159, 160, 7, 5, 2,
	2, 160, 161, 5, 22, 12, 2, 161, 162, 7, 6, 2, 2, 162, 178, 3, 2, 2, 2,
	163, 164, 7, 17, 2, 2, 164, 165, 7, 49, 2, 2, 165, 166, 7, 60, 2, 2, 166,
	178, 7, 6, 2, 2, 167, 168, 7, 18, 2, 2, 168, 169, 7, 49, 2, 2, 169, 170,
	5, 22, 12, 2, 170, 171, 7, 6, 2, 2, 171, 178, 3, 2, 2, 2, 172, 173, 7,
	19, 2, 2, 173, 174, 7, 49, 2, 2, 174, 175, 5, 22, 12, 2, 175, 176, 7, 6,
	2, 2, 176, 178, 3, 2, 2, 2, 177, 141, 3, 2, 2, 2, 177, 146, 3, 2, 2, 2,
	177, 151, 3, 2, 2, 2, 177, 156, 3, 2, 2, 2, 177, 163, 3, 2, 2, 2, 177,
	167, 3, 2, 2, 2, 177, 172, 3, 2, 2, 2, 178, 21, 3, 2, 2, 2, 179, 180, 8,
	12, 1, 2, 180, 202, 5, 20, 11, 2, 181, 182, 7, 60, 2, 2, 182, 184, 7, 49,
	2, 2, 183, 185, 5, 24, 13, 2, 184, 183, 3, 2, 2, 2, 184, 185, 3, 2, 2,
	2, 185, 186, 3, 2, 2, 2, 186, 202, 7, 6, 2, 2, 187, 202, 5, 18, 10, 2,
	188, 189, 9, 4, 2, 2, 189, 202, 5, 22, 12, 16, 190, 191, 7, 49, 2, 2, 191,
	192, 5, 22, 12, 2, 192, 193, 7, 6, 2, 2, 193, 202, 3, 2, 2, 2, 194, 202,
	7, 60, 2, 2, 195, 202, 7, 63, 2, 2, 196, 202, 7, 62, 2, 2, 197, 202, 7,
	59, 2, 2, 198, 202, 7, 65, 2, 2, 199, 202, 7, 61, 2, 2, 200, 202, 7, 57,
	2, 2, 201, 179, 3, 2, 2, 2, 201, 181, 3, 2, 2, 2, 201, 187, 3, 2, 2, 2,
	201, 188, 3, 2, 2, 2, 201, 190, 3, 2, 2, 2, 201, 194, 3, 2, 2, 2, 201,
	195, 3, 2, 2, 2, 201, 196, 3, 2, 2, 2, 201, 197, 3, 2, 2, 2, 201, 198,
	3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 201, 200, 3, 2, 2, 2, 202, 223, 3, 2,
	2, 2, 203, 204, 12, 15, 2, 2, 204, 205, 9, 5, 2, 2, 205, 222, 5, 22, 12,
	16, 206, 207, 12, 14, 2, 2, 207, 208, 9, 6, 2, 2, 208, 222, 5, 22, 12,
	15, 209, 210, 12, 13, 2, 2, 210, 211, 9, 7, 2, 2, 211, 222, 5, 22, 12,
	14, 212, 213, 12, 12, 2, 2, 213, 214, 7, 32, 2, 2, 214, 215, 5, 22, 12,
	2, 215, 216, 7, 55, 2, 2, 216, 217, 5, 22, 12, 13, 217, 222, 3, 2, 2, 2,
	218, 219, 12, 10, 2, 2, 219, 220, 9, 3, 2, 2, 220, 222, 5, 22, 12, 11,
	221, 203, 3, 2, 2, 2, 221, 206, 3, 2, 2, 2, 221, 209, 3, 2, 2, 2, 221,
	212, 3, 2, 2, 2, 221, 218, 3, 2, 2, 2, 222, 225, 3, 2, 2, 2, 223, 221,
	3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 23, 3, 2, 2, 2, 225, 223, 3, 2,
	2, 2, 226, 231, 5, 22, 12, 2, 227, 228, 7, 5, 2, 2, 228, 230, 5, 22, 12,
	2, 229, 227, 3, 2, 2, 2, 230, 233, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 231,
	232, 3, 2, 2, 2, 232, 25, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 234, 236, 7,
	33, 2, 2, 235, 237, 5, 22, 12, 2, 236, 235, 3, 2, 2, 2, 236, 237, 3, 2,
	2, 2, 237, 238, 3, 2, 2, 2, 238, 239, 7, 3, 2, 2, 239, 27, 3, 2, 2, 2,
	240, 241, 7, 34, 2, 2, 241, 242, 5, 22, 12, 2, 242, 245, 5, 42, 22, 2,
	243, 244, 7, 35, 2, 2, 244, 246, 5, 42, 22, 2, 245, 243, 3, 2, 2, 2, 245,
	246, 3, 2, 2, 2, 246, 29, 3, 2, 2, 2, 247, 248, 7, 36, 2, 2, 248, 249,
	5, 22, 12, 2, 249, 250, 7, 35, 2, 2, 250, 251, 5, 8, 5, 2, 251, 31, 3,
	2, 2, 2, 252, 253, 7, 37, 2, 2, 253, 254, 5, 22, 12, 2, 254, 255, 5, 8,
	5, 2, 255, 33, 3, 2, 2, 2, 256, 257, 7, 38, 2, 2, 257, 258, 5, 8, 5, 2,
	258, 259, 7, 37, 2, 2, 259, 260, 5, 22, 12, 2, 260, 35, 3, 2, 2, 2, 261,
	264, 7, 39, 2, 2, 262, 265, 5, 12, 7, 2, 263, 265, 5, 16, 9, 2, 264, 262,
	3, 2, 2, 2, 264, 263, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 266, 3, 2,
	2, 2, 266, 267, 7, 3, 2, 2, 267, 268, 5, 22, 12, 2, 268, 269, 7, 3, 2,
	2, 269, 270, 5, 16, 9, 2, 270, 271, 5, 8, 5, 2, 271, 289, 3, 2, 2, 2, 272,
	273, 7, 39, 2, 2, 273, 274, 7, 60, 2, 2, 274, 275, 7, 5, 2, 2, 275, 276,
	7, 60, 2, 2, 276, 277, 7, 9, 2, 2, 277, 278, 7, 54, 2, 2, 278, 279, 5,
	22, 12, 2, 279, 280, 5, 8, 5, 2, 280, 289, 3, 2, 2, 2, 281, 282, 7, 39,
	2, 2, 282, 283, 7, 60, 2, 2, 283, 284, 7, 9, 2, 2, 284, 285, 7, 54, 2,
	2, 285, 286, 5, 22, 12, 2, 286, 287, 5, 8, 5, 2, 287, 289, 3, 2, 2, 2,
	288, 261, 3, 2, 2, 2, 288, 272, 3, 2, 2, 2, 288, 281, 3, 2, 2, 2, 289,
	37, 3, 2, 2, 2, 290, 291, 7, 40, 2, 2, 291, 292, 5, 24, 13, 2, 292, 296,
	7, 55, 2, 2, 293, 295, 5, 42, 22, 2, 294, 293, 3, 2, 2, 2, 295, 298, 3,
	2, 2, 2, 296, 294, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 39, 3, 2, 2,
	2, 298, 296, 3, 2, 2, 2, 299, 300, 7, 41, 2, 2, 300, 301, 5, 22, 12, 2,
	301, 305, 7, 7, 2, 2, 302, 304, 5, 38, 20, 2, 303, 302, 3, 2, 2, 2, 304,
	307, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 316,
	3, 2, 2, 2, 307, 305, 3, 2, 2, 2, 308, 309, 7, 42, 2, 2, 309, 313, 7, 55,
	2, 2, 310, 312, 5, 42, 22, 2, 311, 310, 3, 2, 2, 2, 312, 315, 3, 2, 2,
	2, 313, 311, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 317, 3, 2, 2, 2, 315,
	313, 3, 2, 2, 2, 316, 308, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 318,
	3, 2, 2, 2, 318, 319, 7, 8, 2, 2, 319, 41, 3, 2, 2, 2, 320, 321, 5, 22,
	12, 2, 321, 322, 7, 3, 2, 2, 322, 351, 3, 2, 2, 2, 323, 324, 5, 12, 7,
	2, 324, 325, 7, 3, 2, 2, 325, 351, 3, 2, 2, 2, 326, 327, 5, 16, 9, 2, 327,
	328, 7, 3, 2, 2, 328, 351, 3, 2, 2, 2, 329, 351, 5, 28, 15, 2, 330, 351,
	5, 30, 16, 2, 331, 351, 5, 32, 17, 2, 332, 333, 5, 34, 18, 2, 333, 334,
	7, 3, 2, 2, 334, 351, 3, 2, 2, 2, 335, 351, 5, 40, 21, 2, 336, 351, 5,
	26, 14, 2, 337, 351, 5, 36, 19, 2, 338, 351, 5, 8, 5, 2, 339, 340, 7, 43,
	2, 2, 340, 351, 7, 3, 2, 2, 341, 342, 7, 44, 2, 2, 342, 351, 7, 3, 2, 2,
	343, 344, 7, 45, 2, 2, 344, 351, 7, 3, 2, 2, 345, 346, 7, 46, 2, 2, 346,
	347, 7, 49, 2, 2, 347, 348, 7, 6, 2, 2, 348, 351, 7, 3, 2, 2, 349, 351,
	7, 3, 2, 2, 350, 320, 3, 2, 2, 2, 350, 323, 3, 2, 2, 2, 350, 326, 3, 2,
	2, 2, 350, 329, 3, 2, 2, 2, 350, 330, 3, 2, 2, 2, 350, 331, 3, 2, 2, 2,
	350, 332, 3, 2, 2, 2, 350, 335, 3, 2, 2, 2, 350, 336, 3, 2, 2, 2, 350,
	337, 3, 2, 2, 2, 350, 338, 3, 2, 2, 2, 350, 339, 3, 2, 2, 2, 350, 341,
	3, 2, 2, 2, 350, 343, 3, 2, 2, 2, 350, 345, 3, 2, 2, 2, 350, 349, 3, 2,
	2, 2, 351, 43, 3, 2, 2, 2, 352, 357, 7, 60, 2, 2, 353, 354, 7, 48, 2, 2,
	354, 355, 5, 22, 12, 2, 355, 356, 7, 10, 2, 2, 356, 358, 3, 2, 2, 2, 357,
	353, 3, 2, 2, 2, 357, 358, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 360,
	7, 9, 2, 2, 360, 372, 5, 22, 12, 2, 361, 366, 7, 60, 2, 2, 362, 363, 7,
	48, 2, 2, 363, 364, 5, 22, 12, 2, 364, 365, 7, 10, 2, 2, 365, 367, 3, 2,
	2, 2, 366, 362, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2,
	368, 369, 9, 5, 2, 2, 369, 370, 7, 9, 2, 2, 370, 372, 5, 22, 12, 2, 371,
	352, 3, 2, 2, 2, 371, 361, 3, 2, 2, 2, 372, 45, 3, 2, 2, 2, 35, 50, 52,
	64, 67, 71, 82, 94, 99, 106, 109, 119, 122, 127, 131, 139, 177, 184, 201,
	221, 223, 231, 236, 245, 264, 288, 296, 305, 313, 316, 350, 357, 366, 371,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'func'", "','", "')'", "'{'", "'}'", "'='", "']'", "'++'",
	"'--'", "'print'", "'printB'", "'printH'", "'delete'", "'len'", "'strLen'",
	"'stringValue'", "'+'", "'-'", "'!'", "'^'", "'=='", "'!='", "'>'", "'<'",
	"'>='", "'<='", "'&&'", "'||'", "'?'", "'return'", "'if'", "'else'", "'guard'",
	"'while'", "'repeat'", "'for'", "'case'", "'switch'", "'default'", "'break'",
	"'continue'", "'fallthrough'", "'debug'", "'@'", "'['", "'('", "'map'",
	"'slice'", "'fill'", "", "'range'", "':'", "':='", "'iota'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "TWODOTS", "LSQUARE", "LPAREN", "MAP",
	"SLICE", "FILL", "TYPES", "RANGE", "COLON", "CE", "IOTA", "BINOP", "BOOL",
	"ID", "CHAR", "INT", "FLOAT", "WS", "STRING",
}

var ruleNames = []string{
	"file", "function", "testRule", "block", "fvarDecl", "varDecl", "kind",
	"simpleStatement", "indexExpr", "funcExpr", "expr", "exprList", "returnStatement",
	"ifStatement", "guardStatement", "whStatement", "repeatStatement", "foStatement",
	"caseStatement", "switchStatement", "statement", "pset",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PinnParser struct {
	*antlr.BaseParser
}

func NewPinnParser(input antlr.TokenStream) *PinnParser {
	this := new(PinnParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Pinn.g4"

	return this
}

// PinnParser tokens.
const (
	PinnParserEOF     = antlr.TokenEOF
	PinnParserT__0    = 1
	PinnParserT__1    = 2
	PinnParserT__2    = 3
	PinnParserT__3    = 4
	PinnParserT__4    = 5
	PinnParserT__5    = 6
	PinnParserT__6    = 7
	PinnParserT__7    = 8
	PinnParserT__8    = 9
	PinnParserT__9    = 10
	PinnParserT__10   = 11
	PinnParserT__11   = 12
	PinnParserT__12   = 13
	PinnParserT__13   = 14
	PinnParserT__14   = 15
	PinnParserT__15   = 16
	PinnParserT__16   = 17
	PinnParserT__17   = 18
	PinnParserT__18   = 19
	PinnParserT__19   = 20
	PinnParserT__20   = 21
	PinnParserT__21   = 22
	PinnParserT__22   = 23
	PinnParserT__23   = 24
	PinnParserT__24   = 25
	PinnParserT__25   = 26
	PinnParserT__26   = 27
	PinnParserT__27   = 28
	PinnParserT__28   = 29
	PinnParserT__29   = 30
	PinnParserT__30   = 31
	PinnParserT__31   = 32
	PinnParserT__32   = 33
	PinnParserT__33   = 34
	PinnParserT__34   = 35
	PinnParserT__35   = 36
	PinnParserT__36   = 37
	PinnParserT__37   = 38
	PinnParserT__38   = 39
	PinnParserT__39   = 40
	PinnParserT__40   = 41
	PinnParserT__41   = 42
	PinnParserT__42   = 43
	PinnParserT__43   = 44
	PinnParserTWODOTS = 45
	PinnParserLSQUARE = 46
	PinnParserLPAREN  = 47
	PinnParserMAP     = 48
	PinnParserSLICE   = 49
	PinnParserFILL    = 50
	PinnParserTYPES   = 51
	PinnParserRANGE   = 52
	PinnParserCOLON   = 53
	PinnParserCE      = 54
	PinnParserIOTA    = 55
	PinnParserBINOP   = 56
	PinnParserBOOL    = 57
	PinnParserID      = 58
	PinnParserCHAR    = 59
	PinnParserINT     = 60
	PinnParserFLOAT   = 61
	PinnParserWS      = 62
	PinnParserSTRING  = 63
)

// PinnParser rules.
const (
	PinnParserRULE_file            = 0
	PinnParserRULE_function        = 1
	PinnParserRULE_testRule        = 2
	PinnParserRULE_block           = 3
	PinnParserRULE_fvarDecl        = 4
	PinnParserRULE_varDecl         = 5
	PinnParserRULE_kind            = 6
	PinnParserRULE_simpleStatement = 7
	PinnParserRULE_indexExpr       = 8
	PinnParserRULE_funcExpr        = 9
	PinnParserRULE_expr            = 10
	PinnParserRULE_exprList        = 11
	PinnParserRULE_returnStatement = 12
	PinnParserRULE_ifStatement     = 13
	PinnParserRULE_guardStatement  = 14
	PinnParserRULE_whStatement     = 15
	PinnParserRULE_repeatStatement = 16
	PinnParserRULE_foStatement     = 17
	PinnParserRULE_caseStatement   = 18
	PinnParserRULE_switchStatement = 19
	PinnParserRULE_statement       = 20
	PinnParserRULE_pset            = 21
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(PinnParserEOF, 0)
}

func (s *FileContext) AllFunction() []IFunctionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionContext)(nil)).Elem())
	var tst = make([]IFunctionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionContext)
		}
	}

	return tst
}

func (s *FileContext) Function(i int) IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *FileContext) AllVarDecl() []IVarDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarDeclContext)(nil)).Elem())
	var tst = make([]IVarDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarDeclContext)
		}
	}

	return tst
}

func (s *FileContext) VarDecl(i int) IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PinnParserRULE_file)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PinnParserT__1 || _la == PinnParserID {
		p.SetState(48)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PinnParserT__1:
			{
				p.SetState(44)
				p.Function()
			}

		case PinnParserID:
			{
				p.SetState(45)
				p.VarDecl()
			}
			{
				p.SetState(46)
				p.Match(PinnParserT__0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
		p.Match(PinnParserEOF)
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(PinnParserID, 0)
}

func (s *FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PinnParserLPAREN, 0)
}

func (s *FunctionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionContext) AllFvarDecl() []IFvarDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFvarDeclContext)(nil)).Elem())
	var tst = make([]IFvarDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFvarDeclContext)
		}
	}

	return tst
}

func (s *FunctionContext) FvarDecl(i int) IFvarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFvarDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFvarDeclContext)
}

func (s *FunctionContext) Kind() IKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKindContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PinnParserRULE_function)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)
		p.Match(PinnParserT__1)
	}
	{
		p.SetState(55)
		p.Match(PinnParserID)
	}
	{
		p.SetState(56)
		p.Match(PinnParserLPAREN)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PinnParserID {
		{
			p.SetState(57)
			p.FvarDecl()
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PinnParserT__2 {
			{
				p.SetState(58)
				p.Match(PinnParserT__2)
			}
			{
				p.SetState(59)
				p.FvarDecl()
			}

			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(67)
		p.Match(PinnParserT__3)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PinnParserLSQUARE || _la == PinnParserTYPES {
		{
			p.SetState(68)
			p.Kind()
		}

	}
	{
		p.SetState(71)
		p.Block()
	}

	return localctx
}

// ITestRuleContext is an interface to support dynamic dispatch.
type ITestRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestRuleContext differentiates from other interfaces.
	IsTestRuleContext()
}

type TestRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestRuleContext() *TestRuleContext {
	var p = new(TestRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_testRule
	return p
}

func (*TestRuleContext) IsTestRuleContext() {}

func NewTestRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestRuleContext {
	var p = new(TestRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_testRule

	return p
}

func (s *TestRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TestRuleContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PinnParserID)
}

func (s *TestRuleContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PinnParserID, i)
}

func (s *TestRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) TestRule() (localctx ITestRuleContext) {
	localctx = NewTestRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PinnParserRULE_testRule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(PinnParserID)
	}
	{
		p.SetState(74)
		p.Match(PinnParserID)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PinnParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(PinnParserT__4)
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PinnParserT__0)|(1<<PinnParserT__4)|(1<<PinnParserT__10)|(1<<PinnParserT__11)|(1<<PinnParserT__12)|(1<<PinnParserT__13)|(1<<PinnParserT__14)|(1<<PinnParserT__15)|(1<<PinnParserT__16)|(1<<PinnParserT__17)|(1<<PinnParserT__18)|(1<<PinnParserT__19)|(1<<PinnParserT__20)|(1<<PinnParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(PinnParserT__31-32))|(1<<(PinnParserT__33-32))|(1<<(PinnParserT__34-32))|(1<<(PinnParserT__35-32))|(1<<(PinnParserT__36-32))|(1<<(PinnParserT__38-32))|(1<<(PinnParserT__40-32))|(1<<(PinnParserT__41-32))|(1<<(PinnParserT__42-32))|(1<<(PinnParserT__43-32))|(1<<(PinnParserLPAREN-32))|(1<<(PinnParserIOTA-32))|(1<<(PinnParserBOOL-32))|(1<<(PinnParserID-32))|(1<<(PinnParserCHAR-32))|(1<<(PinnParserINT-32))|(1<<(PinnParserFLOAT-32))|(1<<(PinnParserSTRING-32)))) != 0) {
		{
			p.SetState(77)
			p.Statement()
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(83)
		p.Match(PinnParserT__5)
	}

	return localctx
}

// IFvarDeclContext is an interface to support dynamic dispatch.
type IFvarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFvarDeclContext differentiates from other interfaces.
	IsFvarDeclContext()
}

type FvarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFvarDeclContext() *FvarDeclContext {
	var p = new(FvarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_fvarDecl
	return p
}

func (*FvarDeclContext) IsFvarDeclContext() {}

func NewFvarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FvarDeclContext {
	var p = new(FvarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_fvarDecl

	return p
}

func (s *FvarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FvarDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(PinnParserID, 0)
}

func (s *FvarDeclContext) Kind() IKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKindContext)
}

func (s *FvarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FvarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) FvarDecl() (localctx IFvarDeclContext) {
	localctx = NewFvarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PinnParserRULE_fvarDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(PinnParserID)
	}
	{
		p.SetState(86)
		p.Kind()
	}

	return localctx
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_varDecl
	return p
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) ID() antlr.TerminalNode {
	return s.GetToken(PinnParserID, 0)
}

func (s *VarDeclContext) Kind() IKindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKindContext)
}

func (s *VarDeclContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *VarDeclContext) CE() antlr.TerminalNode {
	return s.GetToken(PinnParserCE, 0)
}

func (s *VarDeclContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PinnParserRULE_varDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(PinnParserID)
		}
		{
			p.SetState(89)
			p.Kind()
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PinnParserT__6 {
			{
				p.SetState(90)
				p.Match(PinnParserT__6)
			}
			{
				p.SetState(91)
				p.ExprList()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(PinnParserID)
		}
		{
			p.SetState(95)
			p.Match(PinnParserCE)
		}
		{
			p.SetState(96)
			p.expr(0)
		}

	}

	return localctx
}

// IKindContext is an interface to support dynamic dispatch.
type IKindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKindContext differentiates from other interfaces.
	IsKindContext()
}

type KindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKindContext() *KindContext {
	var p = new(KindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_kind
	return p
}

func (*KindContext) IsKindContext() {}

func NewKindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KindContext {
	var p = new(KindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_kind

	return p
}

func (s *KindContext) GetParser() antlr.Parser { return s.parser }

func (s *KindContext) TYPES() antlr.TerminalNode {
	return s.GetToken(PinnParserTYPES, 0)
}

func (s *KindContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(PinnParserLSQUARE, 0)
}

func (s *KindContext) MAP() antlr.TerminalNode {
	return s.GetToken(PinnParserMAP, 0)
}

func (s *KindContext) SLICE() antlr.TerminalNode {
	return s.GetToken(PinnParserSLICE, 0)
}

func (s *KindContext) FILL() antlr.TerminalNode {
	return s.GetToken(PinnParserFILL, 0)
}

func (s *KindContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *KindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) Kind() (localctx IKindContext) {
	localctx = NewKindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PinnParserRULE_kind)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PinnParserLSQUARE {
		{
			p.SetState(99)
			p.Match(PinnParserLSQUARE)
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PinnParserMAP:
			{
				p.SetState(100)
				p.Match(PinnParserMAP)
			}

		case PinnParserSLICE:
			{
				p.SetState(101)
				p.Match(PinnParserSLICE)
			}

		case PinnParserFILL:
			{
				p.SetState(102)
				p.Match(PinnParserFILL)
			}

		case PinnParserT__10, PinnParserT__11, PinnParserT__12, PinnParserT__13, PinnParserT__14, PinnParserT__15, PinnParserT__16, PinnParserT__17, PinnParserT__18, PinnParserT__19, PinnParserT__20, PinnParserLPAREN, PinnParserIOTA, PinnParserBOOL, PinnParserID, PinnParserCHAR, PinnParserINT, PinnParserFLOAT, PinnParserSTRING:
			{
				p.SetState(103)
				p.expr(0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(106)
			p.Match(PinnParserT__7)
		}

	}
	{
		p.SetState(109)
		p.Match(PinnParserTYPES)
	}

	return localctx
}

// ISimpleStatementContext is an interface to support dynamic dispatch.
type ISimpleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleStatementContext differentiates from other interfaces.
	IsSimpleStatementContext()
}

type SimpleStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStatementContext() *SimpleStatementContext {
	var p = new(SimpleStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_simpleStatement
	return p
}

func (*SimpleStatementContext) IsSimpleStatementContext() {}

func NewSimpleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStatementContext {
	var p = new(SimpleStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_simpleStatement

	return p
}

func (s *SimpleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStatementContext) Pset() IPsetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPsetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPsetContext)
}

func (s *SimpleStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(PinnParserID, 0)
}

func (s *SimpleStatementContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(PinnParserLSQUARE, 0)
}

func (s *SimpleStatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SimpleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) SimpleStatement() (localctx ISimpleStatementContext) {
	localctx = NewSimpleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PinnParserRULE_simpleStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Pset()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(PinnParserID)
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PinnParserLSQUARE {
			{
				p.SetState(113)
				p.Match(PinnParserLSQUARE)
			}
			{
				p.SetState(114)
				p.expr(0)
			}
			{
				p.SetState(115)
				p.Match(PinnParserT__7)
			}

		}
		{
			p.SetState(119)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PinnParserT__8 || _la == PinnParserT__9) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IIndexExprContext is an interface to support dynamic dispatch.
type IIndexExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirst returns the first rule contexts.
	GetFirst() IExprContext

	// GetSecond returns the second rule contexts.
	GetSecond() IExprContext

	// SetFirst sets the first rule contexts.
	SetFirst(IExprContext)

	// SetSecond sets the second rule contexts.
	SetSecond(IExprContext)

	// IsIndexExprContext differentiates from other interfaces.
	IsIndexExprContext()
}

type IndexExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	first  IExprContext
	second IExprContext
}

func NewEmptyIndexExprContext() *IndexExprContext {
	var p = new(IndexExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_indexExpr
	return p
}

func (*IndexExprContext) IsIndexExprContext() {}

func NewIndexExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexExprContext {
	var p = new(IndexExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_indexExpr

	return p
}

func (s *IndexExprContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexExprContext) GetFirst() IExprContext { return s.first }

func (s *IndexExprContext) GetSecond() IExprContext { return s.second }

func (s *IndexExprContext) SetFirst(v IExprContext) { s.first = v }

func (s *IndexExprContext) SetSecond(v IExprContext) { s.second = v }

func (s *IndexExprContext) ID() antlr.TerminalNode {
	return s.GetToken(PinnParserID, 0)
}

func (s *IndexExprContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(PinnParserLSQUARE, 0)
}

func (s *IndexExprContext) TWODOTS() antlr.TerminalNode {
	return s.GetToken(PinnParserTWODOTS, 0)
}

func (s *IndexExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(PinnParserCOLON, 0)
}

func (s *IndexExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IndexExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) IndexExpr() (localctx IIndexExprContext) {
	localctx = NewIndexExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PinnParserRULE_indexExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(PinnParserID)
		}
		{
			p.SetState(123)
			p.Match(PinnParserLSQUARE)
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PinnParserT__10)|(1<<PinnParserT__11)|(1<<PinnParserT__12)|(1<<PinnParserT__13)|(1<<PinnParserT__14)|(1<<PinnParserT__15)|(1<<PinnParserT__16)|(1<<PinnParserT__17)|(1<<PinnParserT__18)|(1<<PinnParserT__19)|(1<<PinnParserT__20))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(PinnParserLPAREN-47))|(1<<(PinnParserIOTA-47))|(1<<(PinnParserBOOL-47))|(1<<(PinnParserID-47))|(1<<(PinnParserCHAR-47))|(1<<(PinnParserINT-47))|(1<<(PinnParserFLOAT-47))|(1<<(PinnParserSTRING-47)))) != 0) {
			{
				p.SetState(124)

				var _x = p.expr(0)

				localctx.(*IndexExprContext).first = _x
			}

		}
		{
			p.SetState(127)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PinnParserTWODOTS || _la == PinnParserCOLON) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PinnParserT__10)|(1<<PinnParserT__11)|(1<<PinnParserT__12)|(1<<PinnParserT__13)|(1<<PinnParserT__14)|(1<<PinnParserT__15)|(1<<PinnParserT__16)|(1<<PinnParserT__17)|(1<<PinnParserT__18)|(1<<PinnParserT__19)|(1<<PinnParserT__20))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(PinnParserLPAREN-47))|(1<<(PinnParserIOTA-47))|(1<<(PinnParserBOOL-47))|(1<<(PinnParserID-47))|(1<<(PinnParserCHAR-47))|(1<<(PinnParserINT-47))|(1<<(PinnParserFLOAT-47))|(1<<(PinnParserSTRING-47)))) != 0) {
			{
				p.SetState(128)

				var _x = p.expr(0)

				localctx.(*IndexExprContext).second = _x
			}

		}
		{
			p.SetState(131)
			p.Match(PinnParserT__7)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(PinnParserID)
		}
		{
			p.SetState(133)
			p.Match(PinnParserLSQUARE)
		}
		{
			p.SetState(134)
			p.expr(0)
		}
		{
			p.SetState(135)
			p.Match(PinnParserT__7)
		}

	}

	return localctx
}

// IFuncExprContext is an interface to support dynamic dispatch.
type IFuncExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncExprContext differentiates from other interfaces.
	IsFuncExprContext()
}

type FuncExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExprContext() *FuncExprContext {
	var p = new(FuncExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_funcExpr
	return p
}

func (*FuncExprContext) IsFuncExprContext() {}

func NewFuncExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExprContext {
	var p = new(FuncExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_funcExpr

	return p
}

func (s *FuncExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PinnParserLPAREN, 0)
}

func (s *FuncExprContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *FuncExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncExprContext) ID() antlr.TerminalNode {
	return s.GetToken(PinnParserID, 0)
}

func (s *FuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) FuncExpr() (localctx IFuncExprContext) {
	localctx = NewFuncExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PinnParserRULE_funcExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PinnParserT__10:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.Match(PinnParserT__10)
		}
		{
			p.SetState(140)
			p.Match(PinnParserLPAREN)
		}
		{
			p.SetState(141)
			p.ExprList()
		}
		{
			p.SetState(142)
			p.Match(PinnParserT__3)
		}

	case PinnParserT__11:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.Match(PinnParserT__11)
		}
		{
			p.SetState(145)
			p.Match(PinnParserLPAREN)
		}
		{
			p.SetState(146)
			p.expr(0)
		}
		{
			p.SetState(147)
			p.Match(PinnParserT__3)
		}

	case PinnParserT__12:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(149)
			p.Match(PinnParserT__12)
		}
		{
			p.SetState(150)
			p.Match(PinnParserLPAREN)
		}
		{
			p.SetState(151)
			p.expr(0)
		}
		{
			p.SetState(152)
			p.Match(PinnParserT__3)
		}

	case PinnParserT__13:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(154)
			p.Match(PinnParserT__13)
		}
		{
			p.SetState(155)
			p.Match(PinnParserLPAREN)
		}
		{
			p.SetState(156)
			p.Match(PinnParserID)
		}
		{
			p.SetState(157)
			p.Match(PinnParserT__2)
		}
		{
			p.SetState(158)
			p.expr(0)
		}
		{
			p.SetState(159)
			p.Match(PinnParserT__3)
		}

	case PinnParserT__14:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(161)
			p.Match(PinnParserT__14)
		}
		{
			p.SetState(162)
			p.Match(PinnParserLPAREN)
		}
		{
			p.SetState(163)
			p.Match(PinnParserID)
		}
		{
			p.SetState(164)
			p.Match(PinnParserT__3)
		}

	case PinnParserT__15:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(165)
			p.Match(PinnParserT__15)
		}
		{
			p.SetState(166)
			p.Match(PinnParserLPAREN)
		}
		{
			p.SetState(167)
			p.expr(0)
		}
		{
			p.SetState(168)
			p.Match(PinnParserT__3)
		}

	case PinnParserT__16:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(170)
			p.Match(PinnParserT__16)
		}
		{
			p.SetState(171)
			p.Match(PinnParserLPAREN)
		}
		{
			p.SetState(172)
			p.expr(0)
		}
		{
			p.SetState(173)
			p.Match(PinnParserT__3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirstExpr returns the firstExpr rule contexts.
	GetFirstExpr() IExprContext

	// GetSecondExpr returns the secondExpr rule contexts.
	GetSecondExpr() IExprContext

	// SetFirstExpr sets the firstExpr rule contexts.
	SetFirstExpr(IExprContext)

	// SetSecondExpr sets the secondExpr rule contexts.
	SetSecondExpr(IExprContext)

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	firstExpr  IExprContext
	secondExpr IExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetFirstExpr() IExprContext { return s.firstExpr }

func (s *ExprContext) GetSecondExpr() IExprContext { return s.secondExpr }

func (s *ExprContext) SetFirstExpr(v IExprContext) { s.firstExpr = v }

func (s *ExprContext) SetSecondExpr(v IExprContext) { s.secondExpr = v }

func (s *ExprContext) FuncExpr() IFuncExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncExprContext)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(PinnParserID, 0)
}

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PinnParserLPAREN, 0)
}

func (s *ExprContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *ExprContext) IndexExpr() IIndexExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIndexExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIndexExprContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(PinnParserFLOAT, 0)
}

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(PinnParserINT, 0)
}

func (s *ExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(PinnParserBOOL, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(PinnParserSTRING, 0)
}

func (s *ExprContext) CHAR() antlr.TerminalNode {
	return s.GetToken(PinnParserCHAR, 0)
}

func (s *ExprContext) IOTA() antlr.TerminalNode {
	return s.GetToken(PinnParserIOTA, 0)
}

func (s *ExprContext) BINOP() antlr.TerminalNode {
	return s.GetToken(PinnParserBINOP, 0)
}

func (s *ExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(PinnParserCOLON, 0)
}

func (s *ExprContext) TWODOTS() antlr.TerminalNode {
	return s.GetToken(PinnParserTWODOTS, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *PinnParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, PinnParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(178)
			p.FuncExpr()
		}

	case 2:
		{
			p.SetState(179)
			p.Match(PinnParserID)
		}
		{
			p.SetState(180)
			p.Match(PinnParserLPAREN)
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PinnParserT__10)|(1<<PinnParserT__11)|(1<<PinnParserT__12)|(1<<PinnParserT__13)|(1<<PinnParserT__14)|(1<<PinnParserT__15)|(1<<PinnParserT__16)|(1<<PinnParserT__17)|(1<<PinnParserT__18)|(1<<PinnParserT__19)|(1<<PinnParserT__20))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(PinnParserLPAREN-47))|(1<<(PinnParserIOTA-47))|(1<<(PinnParserBOOL-47))|(1<<(PinnParserID-47))|(1<<(PinnParserCHAR-47))|(1<<(PinnParserINT-47))|(1<<(PinnParserFLOAT-47))|(1<<(PinnParserSTRING-47)))) != 0) {
			{
				p.SetState(181)
				p.ExprList()
			}

		}
		{
			p.SetState(184)
			p.Match(PinnParserT__3)
		}

	case 3:
		{
			p.SetState(185)
			p.IndexExpr()
		}

	case 4:
		{
			p.SetState(186)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PinnParserT__17)|(1<<PinnParserT__18)|(1<<PinnParserT__19)|(1<<PinnParserT__20))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(187)
			p.expr(14)
		}

	case 5:
		{
			p.SetState(188)
			p.Match(PinnParserLPAREN)
		}
		{
			p.SetState(189)
			p.expr(0)
		}
		{
			p.SetState(190)
			p.Match(PinnParserT__3)
		}

	case 6:
		{
			p.SetState(192)
			p.Match(PinnParserID)
		}

	case 7:
		{
			p.SetState(193)
			p.Match(PinnParserFLOAT)
		}

	case 8:
		{
			p.SetState(194)
			p.Match(PinnParserINT)
		}

	case 9:
		{
			p.SetState(195)
			p.Match(PinnParserBOOL)
		}

	case 10:
		{
			p.SetState(196)
			p.Match(PinnParserSTRING)
		}

	case 11:
		{
			p.SetState(197)
			p.Match(PinnParserCHAR)
		}

	case 12:
		{
			p.SetState(198)
			p.Match(PinnParserIOTA)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PinnParserRULE_expr)
				p.SetState(201)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(202)
					_la = p.GetTokenStream().LA(1)

					if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PinnParserT__17)|(1<<PinnParserT__18)|(1<<PinnParserT__20))) != 0) || _la == PinnParserBINOP) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(203)
					p.expr(14)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PinnParserRULE_expr)
				p.SetState(204)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(205)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PinnParserT__21)|(1<<PinnParserT__22)|(1<<PinnParserT__23)|(1<<PinnParserT__24)|(1<<PinnParserT__25)|(1<<PinnParserT__26))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(206)
					p.expr(13)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PinnParserRULE_expr)
				p.SetState(207)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(208)
					_la = p.GetTokenStream().LA(1)

					if !(_la == PinnParserT__27 || _la == PinnParserT__28) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(209)
					p.expr(12)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PinnParserRULE_expr)
				p.SetState(210)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(211)
					p.Match(PinnParserT__29)
				}
				{
					p.SetState(212)
					p.expr(0)
				}
				{
					p.SetState(213)
					p.Match(PinnParserCOLON)
				}
				{
					p.SetState(214)
					p.expr(11)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).firstExpr = _prevctx
				p.PushNewRecursionContext(localctx, _startState, PinnParserRULE_expr)
				p.SetState(216)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(217)
					_la = p.GetTokenStream().LA(1)

					if !(_la == PinnParserTWODOTS || _la == PinnParserCOLON) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(218)

					var _x = p.expr(9)

					localctx.(*ExprContext).secondExpr = _x
				}

			}

		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}

	return localctx
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_exprList
	return p
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PinnParserRULE_exprList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.expr(0)
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PinnParserT__2 {
		{
			p.SetState(225)
			p.Match(PinnParserT__2)
		}
		{
			p.SetState(226)
			p.expr(0)
		}

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PinnParserRULE_returnStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(PinnParserT__30)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PinnParserT__10)|(1<<PinnParserT__11)|(1<<PinnParserT__12)|(1<<PinnParserT__13)|(1<<PinnParserT__14)|(1<<PinnParserT__15)|(1<<PinnParserT__16)|(1<<PinnParserT__17)|(1<<PinnParserT__18)|(1<<PinnParserT__19)|(1<<PinnParserT__20))) != 0) || (((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(PinnParserLPAREN-47))|(1<<(PinnParserIOTA-47))|(1<<(PinnParserBOOL-47))|(1<<(PinnParserID-47))|(1<<(PinnParserCHAR-47))|(1<<(PinnParserINT-47))|(1<<(PinnParserFLOAT-47))|(1<<(PinnParserSTRING-47)))) != 0) {
		{
			p.SetState(233)
			p.expr(0)
		}

	}
	{
		p.SetState(236)
		p.Match(PinnParserT__0)
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PinnParserRULE_ifStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(PinnParserT__31)
	}
	{
		p.SetState(239)
		p.expr(0)
	}
	{
		p.SetState(240)
		p.Statement()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(241)
			p.Match(PinnParserT__32)
		}
		{
			p.SetState(242)
			p.Statement()
		}

	}

	return localctx
}

// IGuardStatementContext is an interface to support dynamic dispatch.
type IGuardStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGuardStatementContext differentiates from other interfaces.
	IsGuardStatementContext()
}

type GuardStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardStatementContext() *GuardStatementContext {
	var p = new(GuardStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_guardStatement
	return p
}

func (*GuardStatementContext) IsGuardStatementContext() {}

func NewGuardStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardStatementContext {
	var p = new(GuardStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_guardStatement

	return p
}

func (s *GuardStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardStatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *GuardStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) GuardStatement() (localctx IGuardStatementContext) {
	localctx = NewGuardStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PinnParserRULE_guardStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Match(PinnParserT__33)
	}
	{
		p.SetState(246)
		p.expr(0)
	}
	{
		p.SetState(247)
		p.Match(PinnParserT__32)
	}
	{
		p.SetState(248)
		p.Block()
	}

	return localctx
}

// IWhStatementContext is an interface to support dynamic dispatch.
type IWhStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhStatementContext differentiates from other interfaces.
	IsWhStatementContext()
}

type WhStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhStatementContext() *WhStatementContext {
	var p = new(WhStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_whStatement
	return p
}

func (*WhStatementContext) IsWhStatementContext() {}

func NewWhStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhStatementContext {
	var p = new(WhStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_whStatement

	return p
}

func (s *WhStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhStatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) WhStatement() (localctx IWhStatementContext) {
	localctx = NewWhStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PinnParserRULE_whStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Match(PinnParserT__34)
	}
	{
		p.SetState(251)
		p.expr(0)
	}
	{
		p.SetState(252)
		p.Block()
	}

	return localctx
}

// IRepeatStatementContext is an interface to support dynamic dispatch.
type IRepeatStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRepeatStatementContext differentiates from other interfaces.
	IsRepeatStatementContext()
}

type RepeatStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepeatStatementContext() *RepeatStatementContext {
	var p = new(RepeatStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_repeatStatement
	return p
}

func (*RepeatStatementContext) IsRepeatStatementContext() {}

func NewRepeatStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatStatementContext {
	var p = new(RepeatStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_repeatStatement

	return p
}

func (s *RepeatStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RepeatStatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RepeatStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) RepeatStatement() (localctx IRepeatStatementContext) {
	localctx = NewRepeatStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PinnParserRULE_repeatStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.Match(PinnParserT__35)
	}
	{
		p.SetState(255)
		p.Block()
	}
	{
		p.SetState(256)
		p.Match(PinnParserT__34)
	}
	{
		p.SetState(257)
		p.expr(0)
	}

	return localctx
}

// IFoStatementContext is an interface to support dynamic dispatch.
type IFoStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFss returns the fss rule contexts.
	GetFss() ISimpleStatementContext

	// GetSss returns the sss rule contexts.
	GetSss() ISimpleStatementContext

	// SetFss sets the fss rule contexts.
	SetFss(ISimpleStatementContext)

	// SetSss sets the sss rule contexts.
	SetSss(ISimpleStatementContext)

	// IsFoStatementContext differentiates from other interfaces.
	IsFoStatementContext()
}

type FoStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	fss    ISimpleStatementContext
	sss    ISimpleStatementContext
}

func NewEmptyFoStatementContext() *FoStatementContext {
	var p = new(FoStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_foStatement
	return p
}

func (*FoStatementContext) IsFoStatementContext() {}

func NewFoStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FoStatementContext {
	var p = new(FoStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_foStatement

	return p
}

func (s *FoStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *FoStatementContext) GetFss() ISimpleStatementContext { return s.fss }

func (s *FoStatementContext) GetSss() ISimpleStatementContext { return s.sss }

func (s *FoStatementContext) SetFss(v ISimpleStatementContext) { s.fss = v }

func (s *FoStatementContext) SetSss(v ISimpleStatementContext) { s.sss = v }

func (s *FoStatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FoStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FoStatementContext) AllSimpleStatement() []ISimpleStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISimpleStatementContext)(nil)).Elem())
	var tst = make([]ISimpleStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISimpleStatementContext)
		}
	}

	return tst
}

func (s *FoStatementContext) SimpleStatement(i int) ISimpleStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *FoStatementContext) VarDecl() IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *FoStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PinnParserID)
}

func (s *FoStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PinnParserID, i)
}

func (s *FoStatementContext) RANGE() antlr.TerminalNode {
	return s.GetToken(PinnParserRANGE, 0)
}

func (s *FoStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FoStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) FoStatement() (localctx IFoStatementContext) {
	localctx = NewFoStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PinnParserRULE_foStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Match(PinnParserT__36)
		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(260)
				p.VarDecl()
			}

		} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(261)

				var _x = p.SimpleStatement()

				localctx.(*FoStatementContext).fss = _x
			}

		}
		{
			p.SetState(264)
			p.Match(PinnParserT__0)
		}
		{
			p.SetState(265)
			p.expr(0)
		}
		{
			p.SetState(266)
			p.Match(PinnParserT__0)
		}
		{
			p.SetState(267)

			var _x = p.SimpleStatement()

			localctx.(*FoStatementContext).sss = _x
		}
		{
			p.SetState(268)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			p.Match(PinnParserT__36)
		}
		{
			p.SetState(271)
			p.Match(PinnParserID)
		}
		{
			p.SetState(272)
			p.Match(PinnParserT__2)
		}
		{
			p.SetState(273)
			p.Match(PinnParserID)
		}
		{
			p.SetState(274)
			p.Match(PinnParserT__6)
		}
		{
			p.SetState(275)
			p.Match(PinnParserRANGE)
		}
		{
			p.SetState(276)
			p.expr(0)
		}
		{
			p.SetState(277)
			p.Block()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(279)
			p.Match(PinnParserT__36)
		}
		{
			p.SetState(280)
			p.Match(PinnParserID)
		}
		{
			p.SetState(281)
			p.Match(PinnParserT__6)
		}
		{
			p.SetState(282)
			p.Match(PinnParserRANGE)
		}
		{
			p.SetState(283)
			p.expr(0)
		}
		{
			p.SetState(284)
			p.Block()
		}

	}

	return localctx
}

// ICaseStatementContext is an interface to support dynamic dispatch.
type ICaseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseStatementContext differentiates from other interfaces.
	IsCaseStatementContext()
}

type CaseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseStatementContext() *CaseStatementContext {
	var p = new(CaseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_caseStatement
	return p
}

func (*CaseStatementContext) IsCaseStatementContext() {}

func NewCaseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseStatementContext {
	var p = new(CaseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_caseStatement

	return p
}

func (s *CaseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseStatementContext) ExprList() IExprListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprListContext)
}

func (s *CaseStatementContext) COLON() antlr.TerminalNode {
	return s.GetToken(PinnParserCOLON, 0)
}

func (s *CaseStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *CaseStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CaseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) CaseStatement() (localctx ICaseStatementContext) {
	localctx = NewCaseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PinnParserRULE_caseStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(PinnParserT__37)
	}
	{
		p.SetState(289)
		p.ExprList()
	}
	{
		p.SetState(290)
		p.Match(PinnParserCOLON)
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PinnParserT__0)|(1<<PinnParserT__4)|(1<<PinnParserT__10)|(1<<PinnParserT__11)|(1<<PinnParserT__12)|(1<<PinnParserT__13)|(1<<PinnParserT__14)|(1<<PinnParserT__15)|(1<<PinnParserT__16)|(1<<PinnParserT__17)|(1<<PinnParserT__18)|(1<<PinnParserT__19)|(1<<PinnParserT__20)|(1<<PinnParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(PinnParserT__31-32))|(1<<(PinnParserT__33-32))|(1<<(PinnParserT__34-32))|(1<<(PinnParserT__35-32))|(1<<(PinnParserT__36-32))|(1<<(PinnParserT__38-32))|(1<<(PinnParserT__40-32))|(1<<(PinnParserT__41-32))|(1<<(PinnParserT__42-32))|(1<<(PinnParserT__43-32))|(1<<(PinnParserLPAREN-32))|(1<<(PinnParserIOTA-32))|(1<<(PinnParserBOOL-32))|(1<<(PinnParserID-32))|(1<<(PinnParserCHAR-32))|(1<<(PinnParserINT-32))|(1<<(PinnParserFLOAT-32))|(1<<(PinnParserSTRING-32)))) != 0) {
		{
			p.SetState(291)
			p.Statement()
		}

		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_switchStatement
	return p
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchStatementContext) AllCaseStatement() []ICaseStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseStatementContext)(nil)).Elem())
	var tst = make([]ICaseStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseStatementContext)
		}
	}

	return tst
}

func (s *SwitchStatementContext) CaseStatement(i int) ICaseStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseStatementContext)
}

func (s *SwitchStatementContext) COLON() antlr.TerminalNode {
	return s.GetToken(PinnParserCOLON, 0)
}

func (s *SwitchStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *SwitchStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PinnParserRULE_switchStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(PinnParserT__38)
	}
	{
		p.SetState(298)
		p.expr(0)
	}
	{
		p.SetState(299)
		p.Match(PinnParserT__4)
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PinnParserT__37 {
		{
			p.SetState(300)
			p.CaseStatement()
		}

		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PinnParserT__39 {
		{
			p.SetState(306)
			p.Match(PinnParserT__39)
		}
		{
			p.SetState(307)
			p.Match(PinnParserCOLON)
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PinnParserT__0)|(1<<PinnParserT__4)|(1<<PinnParserT__10)|(1<<PinnParserT__11)|(1<<PinnParserT__12)|(1<<PinnParserT__13)|(1<<PinnParserT__14)|(1<<PinnParserT__15)|(1<<PinnParserT__16)|(1<<PinnParserT__17)|(1<<PinnParserT__18)|(1<<PinnParserT__19)|(1<<PinnParserT__20)|(1<<PinnParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(PinnParserT__31-32))|(1<<(PinnParserT__33-32))|(1<<(PinnParserT__34-32))|(1<<(PinnParserT__35-32))|(1<<(PinnParserT__36-32))|(1<<(PinnParserT__38-32))|(1<<(PinnParserT__40-32))|(1<<(PinnParserT__41-32))|(1<<(PinnParserT__42-32))|(1<<(PinnParserT__43-32))|(1<<(PinnParserLPAREN-32))|(1<<(PinnParserIOTA-32))|(1<<(PinnParserBOOL-32))|(1<<(PinnParserID-32))|(1<<(PinnParserCHAR-32))|(1<<(PinnParserINT-32))|(1<<(PinnParserFLOAT-32))|(1<<(PinnParserSTRING-32)))) != 0) {
			{
				p.SetState(308)
				p.Statement()
			}

			p.SetState(313)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(316)
		p.Match(PinnParserT__5)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatementContext) VarDecl() IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *StatementContext) SimpleStatement() ISimpleStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) GuardStatement() IGuardStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGuardStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGuardStatementContext)
}

func (s *StatementContext) WhStatement() IWhStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhStatementContext)
}

func (s *StatementContext) RepeatStatement() IRepeatStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepeatStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRepeatStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) FoStatement() IFoStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFoStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFoStatementContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PinnParserLPAREN, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *PinnParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PinnParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.expr(0)
		}
		{
			p.SetState(319)
			p.Match(PinnParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(321)
			p.VarDecl()
		}
		{
			p.SetState(322)
			p.Match(PinnParserT__0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(324)
			p.SimpleStatement()
		}
		{
			p.SetState(325)
			p.Match(PinnParserT__0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(327)
			p.IfStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(328)
			p.GuardStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(329)
			p.WhStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(330)
			p.RepeatStatement()
		}
		{
			p.SetState(331)
			p.Match(PinnParserT__0)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(333)
			p.SwitchStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(334)
			p.ReturnStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(335)
			p.FoStatement()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(336)
			p.Block()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(337)
			p.Match(PinnParserT__40)
		}
		{
			p.SetState(338)
			p.Match(PinnParserT__0)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(339)
			p.Match(PinnParserT__41)
		}
		{
			p.SetState(340)
			p.Match(PinnParserT__0)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(341)
			p.Match(PinnParserT__42)
		}
		{
			p.SetState(342)
			p.Match(PinnParserT__0)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(343)
			p.Match(PinnParserT__43)
		}
		{
			p.SetState(344)
			p.Match(PinnParserLPAREN)
		}
		{
			p.SetState(345)
			p.Match(PinnParserT__3)
		}
		{
			p.SetState(346)
			p.Match(PinnParserT__0)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(347)
			p.Match(PinnParserT__0)
		}

	}

	return localctx
}

// IPsetContext is an interface to support dynamic dispatch.
type IPsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPsetContext differentiates from other interfaces.
	IsPsetContext()
}

type PsetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPsetContext() *PsetContext {
	var p = new(PsetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PinnParserRULE_pset
	return p
}

func (*PsetContext) IsPsetContext() {}

func NewPsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PsetContext {
	var p = new(PsetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PinnParserRULE_pset

	return p
}

func (s *PsetContext) GetParser() antlr.Parser { return s.parser }

func (s *PsetContext) CopyFrom(ctx *PsetContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SimpleSetContext struct {
	*PsetContext
}

func NewSimpleSetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleSetContext {
	var p = new(SimpleSetContext)

	p.PsetContext = NewEmptyPsetContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PsetContext))

	return p
}

func (s *SimpleSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleSetContext) ID() antlr.TerminalNode {
	return s.GetToken(PinnParserID, 0)
}

func (s *SimpleSetContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SimpleSetContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SimpleSetContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(PinnParserLSQUARE, 0)
}

type CompoundSetContext struct {
	*PsetContext
}

func NewCompoundSetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompoundSetContext {
	var p = new(CompoundSetContext)

	p.PsetContext = NewEmptyPsetContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PsetContext))

	return p
}

func (s *CompoundSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundSetContext) ID() antlr.TerminalNode {
	return s.GetToken(PinnParserID, 0)
}

func (s *CompoundSetContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CompoundSetContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompoundSetContext) BINOP() antlr.TerminalNode {
	return s.GetToken(PinnParserBINOP, 0)
}

func (s *CompoundSetContext) LSQUARE() antlr.TerminalNode {
	return s.GetToken(PinnParserLSQUARE, 0)
}

func (p *PinnParser) Pset() (localctx IPsetContext) {
	localctx = NewPsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PinnParserRULE_pset)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleSetContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)
			p.Match(PinnParserID)
		}
		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PinnParserLSQUARE {
			{
				p.SetState(351)
				p.Match(PinnParserLSQUARE)
			}
			{
				p.SetState(352)
				p.expr(0)
			}
			{
				p.SetState(353)
				p.Match(PinnParserT__7)
			}

		}
		{
			p.SetState(357)
			p.Match(PinnParserT__6)
		}
		{
			p.SetState(358)
			p.expr(0)
		}

	case 2:
		localctx = NewCompoundSetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)
			p.Match(PinnParserID)
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PinnParserLSQUARE {
			{
				p.SetState(360)
				p.Match(PinnParserLSQUARE)
			}
			{
				p.SetState(361)
				p.expr(0)
			}
			{
				p.SetState(362)
				p.Match(PinnParserT__7)
			}

		}
		{
			p.SetState(366)
			_la = p.GetTokenStream().LA(1)

			if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PinnParserT__17)|(1<<PinnParserT__18)|(1<<PinnParserT__20))) != 0) || _la == PinnParserBINOP) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(367)
			p.Match(PinnParserT__6)
		}
		{
			p.SetState(368)
			p.expr(0)
		}

	}

	return localctx
}

func (p *PinnParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PinnParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
