// Code generated from /home/baba/dev/workspace/go/src/github.com/baba2k/graphql-rungen/parser/grammar/GraphQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 45, 393,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 7, 31, 256, 10, 31,
	12, 31, 14, 31, 259, 11, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 5, 32, 270, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 34, 5, 34, 278, 10, 34, 3, 34, 3, 34, 3, 35, 5, 35, 283, 10, 35,
	3, 35, 3, 35, 3, 35, 6, 35, 288, 10, 35, 13, 35, 14, 35, 289, 5, 35, 292,
	10, 35, 3, 35, 5, 35, 295, 10, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3,
	37, 6, 37, 303, 10, 37, 13, 37, 14, 37, 304, 5, 37, 307, 10, 37, 3, 38,
	3, 38, 3, 39, 3, 39, 5, 39, 313, 10, 39, 3, 39, 6, 39, 316, 10, 39, 13,
	39, 14, 39, 317, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 7, 41, 325, 10, 41,
	12, 41, 14, 41, 328, 11, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 42, 5, 42, 337, 10, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 6,
	43, 345, 10, 43, 13, 43, 14, 43, 346, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 5, 46, 359, 10, 46, 3, 47, 3, 47, 3,
	47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51,
	3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 7, 53, 379, 10, 53, 12, 53, 14, 53,
	382, 11, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 5, 54, 390, 10,
	54, 3, 54, 3, 54, 3, 346, 2, 55, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 2,
	87, 2, 89, 2, 91, 2, 93, 2, 95, 2, 97, 2, 99, 2, 101, 2, 103, 2, 105, 44,
	107, 45, 3, 2, 12, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92,
	97, 97, 99, 124, 4, 2, 71, 71, 103, 103, 7, 2, 12, 12, 15, 15, 36, 36,
	94, 94, 8234, 8235, 5, 2, 11, 12, 15, 15, 34, 1, 10, 2, 36, 36, 49, 49,
	94, 94, 100, 100, 104, 104, 112, 112, 116, 116, 118, 118, 5, 2, 50, 59,
	67, 72, 99, 104, 5, 2, 12, 12, 15, 15, 8234, 8235, 4, 2, 11, 11, 34, 34,
	3, 2, 65281, 65281, 2, 404, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2,
	2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3,
	2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61,
	3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2,
	69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2,
	2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2,
	2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 3, 109, 3, 2, 2, 2, 5, 111,
	3, 2, 2, 2, 7, 113, 3, 2, 2, 2, 9, 115, 3, 2, 2, 2, 11, 117, 3, 2, 2, 2,
	13, 119, 3, 2, 2, 2, 15, 121, 3, 2, 2, 2, 17, 123, 3, 2, 2, 2, 19, 125,
	3, 2, 2, 2, 21, 127, 3, 2, 2, 2, 23, 129, 3, 2, 2, 2, 25, 131, 3, 2, 2,
	2, 27, 135, 3, 2, 2, 2, 29, 138, 3, 2, 2, 2, 31, 140, 3, 2, 2, 2, 33, 142,
	3, 2, 2, 2, 35, 151, 3, 2, 2, 2, 37, 157, 3, 2, 2, 2, 39, 166, 3, 2, 2,
	2, 41, 179, 3, 2, 2, 2, 43, 186, 3, 2, 2, 2, 45, 193, 3, 2, 2, 2, 47, 198,
	3, 2, 2, 2, 49, 208, 3, 2, 2, 2, 51, 219, 3, 2, 2, 2, 53, 224, 3, 2, 2,
	2, 55, 230, 3, 2, 2, 2, 57, 236, 3, 2, 2, 2, 59, 243, 3, 2, 2, 2, 61, 253,
	3, 2, 2, 2, 63, 269, 3, 2, 2, 2, 65, 271, 3, 2, 2, 2, 67, 277, 3, 2, 2,
	2, 69, 282, 3, 2, 2, 2, 71, 296, 3, 2, 2, 2, 73, 306, 3, 2, 2, 2, 75, 308,
	3, 2, 2, 2, 77, 310, 3, 2, 2, 2, 79, 319, 3, 2, 2, 2, 81, 321, 3, 2, 2,
	2, 83, 331, 3, 2, 2, 2, 85, 344, 3, 2, 2, 2, 87, 348, 3, 2, 2, 2, 89, 353,
	3, 2, 2, 2, 91, 355, 3, 2, 2, 2, 93, 360, 3, 2, 2, 2, 95, 366, 3, 2, 2,
	2, 97, 368, 3, 2, 2, 2, 99, 370, 3, 2, 2, 2, 101, 372, 3, 2, 2, 2, 103,
	374, 3, 2, 2, 2, 105, 376, 3, 2, 2, 2, 107, 389, 3, 2, 2, 2, 109, 110,
	7, 93, 2, 2, 110, 4, 3, 2, 2, 2, 111, 112, 7, 95, 2, 2, 112, 6, 3, 2, 2,
	2, 113, 114, 7, 125, 2, 2, 114, 8, 3, 2, 2, 2, 115, 116, 7, 127, 2, 2,
	116, 10, 3, 2, 2, 2, 117, 118, 7, 60, 2, 2, 118, 12, 3, 2, 2, 2, 119, 120,
	7, 66, 2, 2, 120, 14, 3, 2, 2, 2, 121, 122, 7, 42, 2, 2, 122, 16, 3, 2,
	2, 2, 123, 124, 7, 43, 2, 2, 124, 18, 3, 2, 2, 2, 125, 126, 7, 38, 2, 2,
	126, 20, 3, 2, 2, 2, 127, 128, 7, 63, 2, 2, 128, 22, 3, 2, 2, 2, 129, 130,
	7, 35, 2, 2, 130, 24, 3, 2, 2, 2, 131, 132, 7, 48, 2, 2, 132, 133, 7, 48,
	2, 2, 133, 134, 7, 48, 2, 2, 134, 26, 3, 2, 2, 2, 135, 136, 7, 113, 2,
	2, 136, 137, 7, 112, 2, 2, 137, 28, 3, 2, 2, 2, 138, 139, 7, 40, 2, 2,
	139, 30, 3, 2, 2, 2, 140, 141, 7, 126, 2, 2, 141, 32, 3, 2, 2, 2, 142,
	143, 7, 104, 2, 2, 143, 144, 7, 116, 2, 2, 144, 145, 7, 99, 2, 2, 145,
	146, 7, 105, 2, 2, 146, 147, 7, 111, 2, 2, 147, 148, 7, 103, 2, 2, 148,
	149, 7, 112, 2, 2, 149, 150, 7, 118, 2, 2, 150, 34, 3, 2, 2, 2, 151, 152,
	7, 115, 2, 2, 152, 153, 7, 119, 2, 2, 153, 154, 7, 103, 2, 2, 154, 155,
	7, 116, 2, 2, 155, 156, 7, 123, 2, 2, 156, 36, 3, 2, 2, 2, 157, 158, 7,
	111, 2, 2, 158, 159, 7, 119, 2, 2, 159, 160, 7, 118, 2, 2, 160, 161, 7,
	99, 2, 2, 161, 162, 7, 118, 2, 2, 162, 163, 7, 107, 2, 2, 163, 164, 7,
	113, 2, 2, 164, 165, 7, 112, 2, 2, 165, 38, 3, 2, 2, 2, 166, 167, 7, 117,
	2, 2, 167, 168, 7, 119, 2, 2, 168, 169, 7, 100, 2, 2, 169, 170, 7, 117,
	2, 2, 170, 171, 7, 101, 2, 2, 171, 172, 7, 116, 2, 2, 172, 173, 7, 107,
	2, 2, 173, 174, 7, 114, 2, 2, 174, 175, 7, 118, 2, 2, 175, 176, 7, 107,
	2, 2, 176, 177, 7, 113, 2, 2, 177, 178, 7, 112, 2, 2, 178, 40, 3, 2, 2,
	2, 179, 180, 7, 117, 2, 2, 180, 181, 7, 101, 2, 2, 181, 182, 7, 106, 2,
	2, 182, 183, 7, 103, 2, 2, 183, 184, 7, 111, 2, 2, 184, 185, 7, 99, 2,
	2, 185, 42, 3, 2, 2, 2, 186, 187, 7, 117, 2, 2, 187, 188, 7, 101, 2, 2,
	188, 189, 7, 99, 2, 2, 189, 190, 7, 110, 2, 2, 190, 191, 7, 99, 2, 2, 191,
	192, 7, 116, 2, 2, 192, 44, 3, 2, 2, 2, 193, 194, 7, 118, 2, 2, 194, 195,
	7, 123, 2, 2, 195, 196, 7, 114, 2, 2, 196, 197, 7, 103, 2, 2, 197, 46,
	3, 2, 2, 2, 198, 199, 7, 107, 2, 2, 199, 200, 7, 112, 2, 2, 200, 201, 7,
	118, 2, 2, 201, 202, 7, 103, 2, 2, 202, 203, 7, 116, 2, 2, 203, 204, 7,
	104, 2, 2, 204, 205, 7, 99, 2, 2, 205, 206, 7, 101, 2, 2, 206, 207, 7,
	103, 2, 2, 207, 48, 3, 2, 2, 2, 208, 209, 7, 107, 2, 2, 209, 210, 7, 111,
	2, 2, 210, 211, 7, 114, 2, 2, 211, 212, 7, 110, 2, 2, 212, 213, 7, 103,
	2, 2, 213, 214, 7, 111, 2, 2, 214, 215, 7, 103, 2, 2, 215, 216, 7, 112,
	2, 2, 216, 217, 7, 118, 2, 2, 217, 218, 7, 117, 2, 2, 218, 50, 3, 2, 2,
	2, 219, 220, 7, 103, 2, 2, 220, 221, 7, 112, 2, 2, 221, 222, 7, 119, 2,
	2, 222, 223, 7, 111, 2, 2, 223, 52, 3, 2, 2, 2, 224, 225, 7, 119, 2, 2,
	225, 226, 7, 112, 2, 2, 226, 227, 7, 107, 2, 2, 227, 228, 7, 113, 2, 2,
	228, 229, 7, 112, 2, 2, 229, 54, 3, 2, 2, 2, 230, 231, 7, 107, 2, 2, 231,
	232, 7, 112, 2, 2, 232, 233, 7, 114, 2, 2, 233, 234, 7, 119, 2, 2, 234,
	235, 7, 118, 2, 2, 235, 56, 3, 2, 2, 2, 236, 237, 7, 103, 2, 2, 237, 238,
	7, 122, 2, 2, 238, 239, 7, 118, 2, 2, 239, 240, 7, 103, 2, 2, 240, 241,
	7, 112, 2, 2, 241, 242, 7, 102, 2, 2, 242, 58, 3, 2, 2, 2, 243, 244, 7,
	102, 2, 2, 244, 245, 7, 107, 2, 2, 245, 246, 7, 116, 2, 2, 246, 247, 7,
	103, 2, 2, 247, 248, 7, 101, 2, 2, 248, 249, 7, 118, 2, 2, 249, 250, 7,
	107, 2, 2, 250, 251, 7, 120, 2, 2, 251, 252, 7, 103, 2, 2, 252, 60, 3,
	2, 2, 2, 253, 257, 9, 2, 2, 2, 254, 256, 9, 3, 2, 2, 255, 254, 3, 2, 2,
	2, 256, 259, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258,
	62, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 260, 261, 7, 118, 2, 2, 261, 262,
	7, 116, 2, 2, 262, 263, 7, 119, 2, 2, 263, 270, 7, 103, 2, 2, 264, 265,
	7, 104, 2, 2, 265, 266, 7, 99, 2, 2, 266, 267, 7, 110, 2, 2, 267, 268,
	7, 117, 2, 2, 268, 270, 7, 103, 2, 2, 269, 260, 3, 2, 2, 2, 269, 264, 3,
	2, 2, 2, 270, 64, 3, 2, 2, 2, 271, 272, 7, 112, 2, 2, 272, 273, 7, 119,
	2, 2, 273, 274, 7, 110, 2, 2, 274, 275, 7, 110, 2, 2, 275, 66, 3, 2, 2,
	2, 276, 278, 5, 71, 36, 2, 277, 276, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2,
	278, 279, 3, 2, 2, 2, 279, 280, 5, 73, 37, 2, 280, 68, 3, 2, 2, 2, 281,
	283, 5, 71, 36, 2, 282, 281, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 284,
	3, 2, 2, 2, 284, 291, 5, 73, 37, 2, 285, 287, 7, 48, 2, 2, 286, 288, 5,
	79, 40, 2, 287, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 287, 3, 2,
	2, 2, 289, 290, 3, 2, 2, 2, 290, 292, 3, 2, 2, 2, 291, 285, 3, 2, 2, 2,
	291, 292, 3, 2, 2, 2, 292, 294, 3, 2, 2, 2, 293, 295, 5, 77, 39, 2, 294,
	293, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 70, 3, 2, 2, 2, 296, 297, 7,
	47, 2, 2, 297, 72, 3, 2, 2, 2, 298, 307, 7, 50, 2, 2, 299, 307, 5, 75,
	38, 2, 300, 302, 5, 75, 38, 2, 301, 303, 5, 79, 40, 2, 302, 301, 3, 2,
	2, 2, 303, 304, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2,
	305, 307, 3, 2, 2, 2, 306, 298, 3, 2, 2, 2, 306, 299, 3, 2, 2, 2, 306,
	300, 3, 2, 2, 2, 307, 74, 3, 2, 2, 2, 308, 309, 4, 51, 59, 2, 309, 76,
	3, 2, 2, 2, 310, 312, 9, 4, 2, 2, 311, 313, 5, 71, 36, 2, 312, 311, 3,
	2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 315, 3, 2, 2, 2, 314, 316, 5, 79, 40,
	2, 315, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 315, 3, 2, 2, 2, 317,
	318, 3, 2, 2, 2, 318, 78, 3, 2, 2, 2, 319, 320, 4, 50, 59, 2, 320, 80,
	3, 2, 2, 2, 321, 326, 7, 36, 2, 2, 322, 325, 10, 5, 2, 2, 323, 325, 5,
	91, 46, 2, 324, 322, 3, 2, 2, 2, 324, 323, 3, 2, 2, 2, 325, 328, 3, 2,
	2, 2, 326, 324, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327, 329, 3, 2, 2, 2,
	328, 326, 3, 2, 2, 2, 329, 330, 7, 36, 2, 2, 330, 82, 3, 2, 2, 2, 331,
	332, 7, 36, 2, 2, 332, 333, 7, 36, 2, 2, 333, 334, 7, 36, 2, 2, 334, 336,
	3, 2, 2, 2, 335, 337, 5, 85, 43, 2, 336, 335, 3, 2, 2, 2, 336, 337, 3,
	2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 339, 7, 36, 2, 2, 339, 340, 7, 36,
	2, 2, 340, 341, 7, 36, 2, 2, 341, 84, 3, 2, 2, 2, 342, 345, 5, 87, 44,
	2, 343, 345, 5, 89, 45, 2, 344, 342, 3, 2, 2, 2, 344, 343, 3, 2, 2, 2,
	345, 346, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 347,
	86, 3, 2, 2, 2, 348, 349, 7, 94, 2, 2, 349, 350, 7, 36, 2, 2, 350, 351,
	7, 36, 2, 2, 351, 352, 7, 36, 2, 2, 352, 88, 3, 2, 2, 2, 353, 354, 9, 6,
	2, 2, 354, 90, 3, 2, 2, 2, 355, 358, 7, 94, 2, 2, 356, 359, 9, 7, 2, 2,
	357, 359, 5, 93, 47, 2, 358, 356, 3, 2, 2, 2, 358, 357, 3, 2, 2, 2, 359,
	92, 3, 2, 2, 2, 360, 361, 7, 119, 2, 2, 361, 362, 5, 95, 48, 2, 362, 363,
	5, 95, 48, 2, 363, 364, 5, 95, 48, 2, 364, 365, 5, 95, 48, 2, 365, 94,
	3, 2, 2, 2, 366, 367, 9, 8, 2, 2, 367, 96, 3, 2, 2, 2, 368, 369, 9, 9,
	2, 2, 369, 98, 3, 2, 2, 2, 370, 371, 9, 10, 2, 2, 371, 100, 3, 2, 2, 2,
	372, 373, 7, 46, 2, 2, 373, 102, 3, 2, 2, 2, 374, 375, 9, 11, 2, 2, 375,
	104, 3, 2, 2, 2, 376, 380, 7, 37, 2, 2, 377, 379, 10, 9, 2, 2, 378, 377,
	3, 2, 2, 2, 379, 382, 3, 2, 2, 2, 380, 378, 3, 2, 2, 2, 380, 381, 3, 2,
	2, 2, 381, 383, 3, 2, 2, 2, 382, 380, 3, 2, 2, 2, 383, 384, 8, 53, 2, 2,
	384, 106, 3, 2, 2, 2, 385, 390, 5, 103, 52, 2, 386, 390, 5, 99, 50, 2,
	387, 390, 5, 97, 49, 2, 388, 390, 5, 101, 51, 2, 389, 385, 3, 2, 2, 2,
	389, 386, 3, 2, 2, 2, 389, 387, 3, 2, 2, 2, 389, 388, 3, 2, 2, 2, 390,
	391, 3, 2, 2, 2, 391, 392, 8, 54, 3, 2, 392, 108, 3, 2, 2, 2, 22, 2, 257,
	269, 277, 282, 289, 291, 294, 304, 306, 312, 317, 324, 326, 336, 344, 346,
	358, 380, 389, 4, 2, 4, 2, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "']'", "'{'", "'}'", "':'", "'@'", "'('", "')'", "'$'", "'='",
	"'!'", "'...'", "'on'", "'&'", "'|'", "'fragment'", "'query'", "'mutation'",
	"'subscription'", "'schema'", "'scalar'", "'type'", "'interface'", "'implements'",
	"'enum'", "'union'", "'input'", "'extend'", "'directive'", "", "", "'null'",
	"", "", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "FRAGMENT",
	"QUERY", "MUTATION", "SUBSCRIPTION", "SCHEMA", "SCALAR", "TYPE", "INTERFACE",
	"IMPLEMENTS", "ENUM", "UNION", "INPUT", "EXTEND", "DIRECTIVE", "NAME",
	"BooleanValue", "NullValue", "IntValue", "FloatValue", "Sign", "IntegerPart",
	"NonZeroDigit", "ExponentPart", "Digit", "StringValue", "TripleQuotedStringValue",
	"Comment", "Ignored",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "FRAGMENT", "QUERY",
	"MUTATION", "SUBSCRIPTION", "SCHEMA", "SCALAR", "TYPE", "INTERFACE", "IMPLEMENTS",
	"ENUM", "UNION", "INPUT", "EXTEND", "DIRECTIVE", "NAME", "BooleanValue",
	"NullValue", "IntValue", "FloatValue", "Sign", "IntegerPart", "NonZeroDigit",
	"ExponentPart", "Digit", "StringValue", "TripleQuotedStringValue", "TripleQuotedStringPart",
	"EscapedTripleQuote", "SourceCharacter", "EscapedChar", "Unicode", "Hex",
	"LineTerminator", "Whitespace", "Comma", "UnicodeBOM", "Comment", "Ignored",
}

type GraphQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewGraphQLLexer(input antlr.CharStream) *GraphQLLexer {

	l := new(GraphQLLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "GraphQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GraphQLLexer tokens.
const (
	GraphQLLexerT__0                    = 1
	GraphQLLexerT__1                    = 2
	GraphQLLexerT__2                    = 3
	GraphQLLexerT__3                    = 4
	GraphQLLexerT__4                    = 5
	GraphQLLexerT__5                    = 6
	GraphQLLexerT__6                    = 7
	GraphQLLexerT__7                    = 8
	GraphQLLexerT__8                    = 9
	GraphQLLexerT__9                    = 10
	GraphQLLexerT__10                   = 11
	GraphQLLexerT__11                   = 12
	GraphQLLexerT__12                   = 13
	GraphQLLexerT__13                   = 14
	GraphQLLexerT__14                   = 15
	GraphQLLexerFRAGMENT                = 16
	GraphQLLexerQUERY                   = 17
	GraphQLLexerMUTATION                = 18
	GraphQLLexerSUBSCRIPTION            = 19
	GraphQLLexerSCHEMA                  = 20
	GraphQLLexerSCALAR                  = 21
	GraphQLLexerTYPE                    = 22
	GraphQLLexerINTERFACE               = 23
	GraphQLLexerIMPLEMENTS              = 24
	GraphQLLexerENUM                    = 25
	GraphQLLexerUNION                   = 26
	GraphQLLexerINPUT                   = 27
	GraphQLLexerEXTEND                  = 28
	GraphQLLexerDIRECTIVE               = 29
	GraphQLLexerNAME                    = 30
	GraphQLLexerBooleanValue            = 31
	GraphQLLexerNullValue               = 32
	GraphQLLexerIntValue                = 33
	GraphQLLexerFloatValue              = 34
	GraphQLLexerSign                    = 35
	GraphQLLexerIntegerPart             = 36
	GraphQLLexerNonZeroDigit            = 37
	GraphQLLexerExponentPart            = 38
	GraphQLLexerDigit                   = 39
	GraphQLLexerStringValue             = 40
	GraphQLLexerTripleQuotedStringValue = 41
	GraphQLLexerComment                 = 42
	GraphQLLexerIgnored                 = 43
)
