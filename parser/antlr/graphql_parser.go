// Code generated from /home/baba/dev/workspace/go/src/github.com/baba2k/graphql-rungen/parser/grammar/GraphQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // GraphQL

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 641,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 3, 2, 6, 2, 138, 10, 2, 13, 2,
	14, 2, 139, 3, 3, 3, 3, 3, 3, 5, 3, 145, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 7, 7, 155, 10, 7, 12, 7, 14, 7, 158, 11, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 7, 8, 164, 10, 8, 12, 8, 14, 8, 167, 11, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 7, 9, 173, 10, 9, 12, 9, 14, 9, 176, 11, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 7, 10, 182, 10, 10, 12, 10, 14, 10, 185, 11, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 6,
	13, 198, 10, 13, 13, 13, 14, 13, 199, 3, 14, 3, 14, 3, 14, 5, 14, 205,
	10, 14, 3, 15, 3, 15, 6, 15, 209, 10, 15, 13, 15, 14, 15, 210, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 229, 10, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 240, 10, 19, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 5, 23, 253,
	10, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 5, 26, 267, 10, 26, 3, 27, 3, 27, 3, 27, 5, 27, 272,
	10, 27, 3, 27, 5, 27, 275, 10, 27, 3, 27, 5, 27, 278, 10, 27, 3, 27, 3,
	27, 5, 27, 282, 10, 27, 3, 28, 3, 28, 6, 28, 286, 10, 28, 13, 28, 14, 28,
	287, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 296, 10, 29, 3, 30,
	3, 30, 6, 30, 300, 10, 30, 13, 30, 14, 30, 301, 3, 30, 3, 30, 3, 31, 3,
	31, 3, 31, 5, 31, 309, 10, 31, 3, 32, 5, 32, 312, 10, 32, 3, 32, 3, 32,
	5, 32, 316, 10, 32, 3, 32, 5, 32, 319, 10, 32, 3, 32, 5, 32, 322, 10, 32,
	3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 5, 34, 330, 10, 34, 3, 35, 3,
	35, 5, 35, 334, 10, 35, 3, 35, 5, 35, 337, 10, 35, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 36, 3, 36, 5, 36, 345, 10, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 3, 39, 5, 39, 355, 10, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	5, 39, 361, 10, 39, 3, 40, 5, 40, 364, 10, 40, 3, 40, 3, 40, 5, 40, 368,
	10, 40, 3, 40, 3, 40, 6, 40, 372, 10, 40, 13, 40, 14, 40, 373, 3, 40, 3,
	40, 3, 41, 5, 41, 379, 10, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42,
	3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 391, 10, 42, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 43, 3, 43, 5, 43, 399, 10, 43, 3, 44, 5, 44, 402, 10, 44, 3, 44,
	3, 44, 3, 44, 5, 44, 407, 10, 44, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45, 413,
	10, 45, 3, 46, 5, 46, 416, 10, 46, 3, 46, 3, 46, 3, 46, 5, 46, 421, 10,
	46, 3, 46, 5, 46, 424, 10, 46, 3, 46, 5, 46, 427, 10, 46, 3, 47, 3, 47,
	3, 47, 3, 47, 5, 47, 433, 10, 47, 3, 47, 5, 47, 436, 10, 47, 3, 47, 5,
	47, 439, 10, 47, 3, 48, 3, 48, 3, 48, 5, 48, 444, 10, 48, 3, 48, 6, 48,
	447, 10, 48, 13, 48, 14, 48, 448, 3, 48, 3, 48, 3, 48, 7, 48, 454, 10,
	48, 12, 48, 14, 48, 457, 11, 48, 3, 49, 3, 49, 6, 49, 461, 10, 49, 13,
	49, 14, 49, 462, 3, 49, 3, 49, 3, 50, 5, 50, 468, 10, 50, 3, 50, 3, 50,
	5, 50, 472, 10, 50, 3, 50, 3, 50, 3, 50, 5, 50, 477, 10, 50, 3, 51, 3,
	51, 6, 51, 481, 10, 51, 13, 51, 14, 51, 482, 3, 51, 3, 51, 3, 52, 5, 52,
	488, 10, 52, 3, 52, 3, 52, 3, 52, 3, 52, 5, 52, 494, 10, 52, 3, 52, 5,
	52, 497, 10, 52, 3, 53, 5, 53, 500, 10, 53, 3, 53, 3, 53, 3, 53, 5, 53,
	505, 10, 53, 3, 53, 5, 53, 508, 10, 53, 3, 54, 3, 54, 3, 54, 3, 54, 5,
	54, 514, 10, 54, 3, 54, 5, 54, 517, 10, 54, 3, 55, 5, 55, 520, 10, 55,
	3, 55, 3, 55, 3, 55, 5, 55, 525, 10, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3,
	56, 3, 56, 5, 56, 533, 10, 56, 3, 56, 5, 56, 536, 10, 56, 3, 57, 3, 57,
	3, 57, 3, 58, 3, 58, 5, 58, 543, 10, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3,
	58, 7, 58, 550, 10, 58, 12, 58, 14, 58, 553, 11, 58, 3, 59, 5, 59, 556,
	10, 59, 3, 59, 3, 59, 3, 59, 5, 59, 561, 10, 59, 3, 59, 3, 59, 3, 60, 3,
	60, 3, 60, 3, 60, 5, 60, 569, 10, 60, 3, 60, 5, 60, 572, 10, 60, 3, 61,
	3, 61, 6, 61, 576, 10, 61, 13, 61, 14, 61, 577, 3, 61, 3, 61, 3, 62, 5,
	62, 583, 10, 62, 3, 62, 3, 62, 5, 62, 587, 10, 62, 3, 63, 5, 63, 590, 10,
	63, 3, 63, 3, 63, 3, 63, 5, 63, 595, 10, 63, 3, 63, 3, 63, 3, 64, 3, 64,
	3, 64, 3, 64, 5, 64, 603, 10, 64, 3, 64, 5, 64, 606, 10, 64, 3, 65, 3,
	65, 6, 65, 610, 10, 65, 13, 65, 14, 65, 611, 3, 65, 3, 65, 3, 66, 5, 66,
	617, 10, 66, 3, 66, 3, 66, 3, 66, 3, 66, 5, 66, 623, 10, 66, 3, 66, 3,
	66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 7, 68,
	636, 10, 68, 12, 68, 14, 68, 639, 11, 68, 3, 68, 2, 5, 94, 114, 134, 69,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
	76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108,
	110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 2, 5,
	3, 2, 19, 21, 3, 2, 18, 32, 3, 2, 42, 43, 2, 683, 2, 137, 3, 2, 2, 2, 4,
	144, 3, 2, 2, 2, 6, 146, 3, 2, 2, 2, 8, 148, 3, 2, 2, 2, 10, 150, 3, 2,
	2, 2, 12, 152, 3, 2, 2, 2, 14, 161, 3, 2, 2, 2, 16, 170, 3, 2, 2, 2, 18,
	179, 3, 2, 2, 2, 20, 188, 3, 2, 2, 2, 22, 192, 3, 2, 2, 2, 24, 197, 3,
	2, 2, 2, 26, 201, 3, 2, 2, 2, 28, 206, 3, 2, 2, 2, 30, 214, 3, 2, 2, 2,
	32, 218, 3, 2, 2, 2, 34, 228, 3, 2, 2, 2, 36, 239, 3, 2, 2, 2, 38, 241,
	3, 2, 2, 2, 40, 244, 3, 2, 2, 2, 42, 247, 3, 2, 2, 2, 44, 252, 3, 2, 2,
	2, 46, 254, 3, 2, 2, 2, 48, 256, 3, 2, 2, 2, 50, 266, 3, 2, 2, 2, 52, 281,
	3, 2, 2, 2, 54, 283, 3, 2, 2, 2, 56, 291, 3, 2, 2, 2, 58, 297, 3, 2, 2,
	2, 60, 308, 3, 2, 2, 2, 62, 311, 3, 2, 2, 2, 64, 323, 3, 2, 2, 2, 66, 326,
	3, 2, 2, 2, 68, 331, 3, 2, 2, 2, 70, 340, 3, 2, 2, 2, 72, 348, 3, 2, 2,
	2, 74, 350, 3, 2, 2, 2, 76, 360, 3, 2, 2, 2, 78, 363, 3, 2, 2, 2, 80, 378,
	3, 2, 2, 2, 82, 390, 3, 2, 2, 2, 84, 398, 3, 2, 2, 2, 86, 401, 3, 2, 2,
	2, 88, 408, 3, 2, 2, 2, 90, 415, 3, 2, 2, 2, 92, 428, 3, 2, 2, 2, 94, 440,
	3, 2, 2, 2, 96, 458, 3, 2, 2, 2, 98, 467, 3, 2, 2, 2, 100, 478, 3, 2, 2,
	2, 102, 487, 3, 2, 2, 2, 104, 499, 3, 2, 2, 2, 106, 509, 3, 2, 2, 2, 108,
	519, 3, 2, 2, 2, 110, 528, 3, 2, 2, 2, 112, 537, 3, 2, 2, 2, 114, 540,
	3, 2, 2, 2, 116, 555, 3, 2, 2, 2, 118, 564, 3, 2, 2, 2, 120, 573, 3, 2,
	2, 2, 122, 582, 3, 2, 2, 2, 124, 589, 3, 2, 2, 2, 126, 598, 3, 2, 2, 2,
	128, 607, 3, 2, 2, 2, 130, 616, 3, 2, 2, 2, 132, 627, 3, 2, 2, 2, 134,
	629, 3, 2, 2, 2, 136, 138, 5, 4, 3, 2, 137, 136, 3, 2, 2, 2, 138, 139,
	3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 3, 3, 2, 2,
	2, 141, 145, 5, 52, 27, 2, 142, 145, 5, 70, 36, 2, 143, 145, 5, 76, 39,
	2, 144, 141, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 143, 3, 2, 2, 2, 145,
	5, 3, 2, 2, 2, 146, 147, 9, 2, 2, 2, 147, 7, 3, 2, 2, 2, 148, 149, 5, 42,
	22, 2, 149, 9, 3, 2, 2, 2, 150, 151, 5, 32, 17, 2, 151, 11, 3, 2, 2, 2,
	152, 156, 7, 3, 2, 2, 153, 155, 5, 34, 18, 2, 154, 153, 3, 2, 2, 2, 155,
	158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 159,
	3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 160, 7, 4, 2, 2, 160, 13, 3, 2,
	2, 2, 161, 165, 7, 3, 2, 2, 162, 164, 5, 36, 19, 2, 163, 162, 3, 2, 2,
	2, 164, 167, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166,
	168, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 168, 169, 7, 4, 2, 2, 169, 15, 3,
	2, 2, 2, 170, 174, 7, 5, 2, 2, 171, 173, 5, 20, 11, 2, 172, 171, 3, 2,
	2, 2, 173, 176, 3, 2, 2, 2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2,
	175, 177, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 177, 178, 7, 6, 2, 2, 178,
	17, 3, 2, 2, 2, 179, 183, 7, 5, 2, 2, 180, 182, 5, 22, 12, 2, 181, 180,
	3, 2, 2, 2, 182, 185, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 183, 184, 3, 2,
	2, 2, 184, 186, 3, 2, 2, 2, 185, 183, 3, 2, 2, 2, 186, 187, 7, 6, 2, 2,
	187, 19, 3, 2, 2, 2, 188, 189, 5, 32, 17, 2, 189, 190, 7, 7, 2, 2, 190,
	191, 5, 34, 18, 2, 191, 21, 3, 2, 2, 2, 192, 193, 5, 32, 17, 2, 193, 194,
	7, 7, 2, 2, 194, 195, 5, 36, 19, 2, 195, 23, 3, 2, 2, 2, 196, 198, 5, 26,
	14, 2, 197, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2,
	199, 200, 3, 2, 2, 2, 200, 25, 3, 2, 2, 2, 201, 202, 7, 8, 2, 2, 202, 204,
	5, 32, 17, 2, 203, 205, 5, 28, 15, 2, 204, 203, 3, 2, 2, 2, 204, 205, 3,
	2, 2, 2, 205, 27, 3, 2, 2, 2, 206, 208, 7, 9, 2, 2, 207, 209, 5, 30, 16,
	2, 208, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210,
	211, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 213, 7, 10, 2, 2, 213, 29,
	3, 2, 2, 2, 214, 215, 5, 32, 17, 2, 215, 216, 7, 7, 2, 2, 216, 217, 5,
	36, 19, 2, 217, 31, 3, 2, 2, 2, 218, 219, 9, 3, 2, 2, 219, 33, 3, 2, 2,
	2, 220, 229, 5, 42, 22, 2, 221, 229, 7, 35, 2, 2, 222, 229, 7, 36, 2, 2,
	223, 229, 7, 33, 2, 2, 224, 229, 7, 34, 2, 2, 225, 229, 5, 10, 6, 2, 226,
	229, 5, 12, 7, 2, 227, 229, 5, 16, 9, 2, 228, 220, 3, 2, 2, 2, 228, 221,
	3, 2, 2, 2, 228, 222, 3, 2, 2, 2, 228, 223, 3, 2, 2, 2, 228, 224, 3, 2,
	2, 2, 228, 225, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 228, 227, 3, 2, 2, 2,
	229, 35, 3, 2, 2, 2, 230, 240, 5, 38, 20, 2, 231, 240, 5, 42, 22, 2, 232,
	240, 7, 35, 2, 2, 233, 240, 7, 36, 2, 2, 234, 240, 7, 33, 2, 2, 235, 240,
	7, 34, 2, 2, 236, 240, 5, 10, 6, 2, 237, 240, 5, 14, 8, 2, 238, 240, 5,
	18, 10, 2, 239, 230, 3, 2, 2, 2, 239, 231, 3, 2, 2, 2, 239, 232, 3, 2,
	2, 2, 239, 233, 3, 2, 2, 2, 239, 234, 3, 2, 2, 2, 239, 235, 3, 2, 2, 2,
	239, 236, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 239, 238, 3, 2, 2, 2, 240,
	37, 3, 2, 2, 2, 241, 242, 7, 11, 2, 2, 242, 243, 5, 32, 17, 2, 243, 39,
	3, 2, 2, 2, 244, 245, 7, 12, 2, 2, 245, 246, 5, 34, 18, 2, 246, 41, 3,
	2, 2, 2, 247, 248, 9, 4, 2, 2, 248, 43, 3, 2, 2, 2, 249, 253, 5, 46, 24,
	2, 250, 253, 5, 48, 25, 2, 251, 253, 5, 50, 26, 2, 252, 249, 3, 2, 2, 2,
	252, 250, 3, 2, 2, 2, 252, 251, 3, 2, 2, 2, 253, 45, 3, 2, 2, 2, 254, 255,
	5, 32, 17, 2, 255, 47, 3, 2, 2, 2, 256, 257, 7, 3, 2, 2, 257, 258, 5, 44,
	23, 2, 258, 259, 7, 4, 2, 2, 259, 49, 3, 2, 2, 2, 260, 261, 5, 46, 24,
	2, 261, 262, 7, 13, 2, 2, 262, 267, 3, 2, 2, 2, 263, 264, 5, 48, 25, 2,
	264, 265, 7, 13, 2, 2, 265, 267, 3, 2, 2, 2, 266, 260, 3, 2, 2, 2, 266,
	263, 3, 2, 2, 2, 267, 51, 3, 2, 2, 2, 268, 282, 5, 58, 30, 2, 269, 271,
	5, 6, 4, 2, 270, 272, 5, 32, 17, 2, 271, 270, 3, 2, 2, 2, 271, 272, 3,
	2, 2, 2, 272, 274, 3, 2, 2, 2, 273, 275, 5, 54, 28, 2, 274, 273, 3, 2,
	2, 2, 274, 275, 3, 2, 2, 2, 275, 277, 3, 2, 2, 2, 276, 278, 5, 24, 13,
	2, 277, 276, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279,
	280, 5, 58, 30, 2, 280, 282, 3, 2, 2, 2, 281, 268, 3, 2, 2, 2, 281, 269,
	3, 2, 2, 2, 282, 53, 3, 2, 2, 2, 283, 285, 7, 9, 2, 2, 284, 286, 5, 56,
	29, 2, 285, 284, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2,
	287, 288, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 290, 7, 10, 2, 2, 290,
	55, 3, 2, 2, 2, 291, 292, 5, 38, 20, 2, 292, 293, 7, 7, 2, 2, 293, 295,
	5, 44, 23, 2, 294, 296, 5, 40, 21, 2, 295, 294, 3, 2, 2, 2, 295, 296, 3,
	2, 2, 2, 296, 57, 3, 2, 2, 2, 297, 299, 7, 5, 2, 2, 298, 300, 5, 60, 31,
	2, 299, 298, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 301,
	302, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 304, 7, 6, 2, 2, 304, 59, 3,
	2, 2, 2, 305, 309, 5, 62, 32, 2, 306, 309, 5, 66, 34, 2, 307, 309, 5, 68,
	35, 2, 308, 305, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 308, 307, 3, 2, 2, 2,
	309, 61, 3, 2, 2, 2, 310, 312, 5, 64, 33, 2, 311, 310, 3, 2, 2, 2, 311,
	312, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 315, 5, 32, 17, 2, 314, 316,
	5, 28, 15, 2, 315, 314, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 318, 3,
	2, 2, 2, 317, 319, 5, 24, 13, 2, 318, 317, 3, 2, 2, 2, 318, 319, 3, 2,
	2, 2, 319, 321, 3, 2, 2, 2, 320, 322, 5, 58, 30, 2, 321, 320, 3, 2, 2,
	2, 321, 322, 3, 2, 2, 2, 322, 63, 3, 2, 2, 2, 323, 324, 5, 32, 17, 2, 324,
	325, 7, 7, 2, 2, 325, 65, 3, 2, 2, 2, 326, 327, 7, 14, 2, 2, 327, 329,
	5, 72, 37, 2, 328, 330, 5, 24, 13, 2, 329, 328, 3, 2, 2, 2, 329, 330, 3,
	2, 2, 2, 330, 67, 3, 2, 2, 2, 331, 333, 7, 14, 2, 2, 332, 334, 5, 74, 38,
	2, 333, 332, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 336, 3, 2, 2, 2, 335,
	337, 5, 24, 13, 2, 336, 335, 3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 338,
	3, 2, 2, 2, 338, 339, 5, 58, 30, 2, 339, 69, 3, 2, 2, 2, 340, 341, 7, 18,
	2, 2, 341, 342, 5, 72, 37, 2, 342, 344, 5, 74, 38, 2, 343, 345, 5, 24,
	13, 2, 344, 343, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2,
	346, 347, 5, 58, 30, 2, 347, 71, 3, 2, 2, 2, 348, 349, 5, 32, 17, 2, 349,
	73, 3, 2, 2, 2, 350, 351, 7, 15, 2, 2, 351, 352, 5, 46, 24, 2, 352, 75,
	3, 2, 2, 2, 353, 355, 5, 8, 5, 2, 354, 353, 3, 2, 2, 2, 354, 355, 3, 2,
	2, 2, 355, 356, 3, 2, 2, 2, 356, 361, 5, 78, 40, 2, 357, 361, 5, 82, 42,
	2, 358, 361, 5, 84, 43, 2, 359, 361, 5, 130, 66, 2, 360, 354, 3, 2, 2,
	2, 360, 357, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 360, 359, 3, 2, 2, 2, 361,
	77, 3, 2, 2, 2, 362, 364, 5, 8, 5, 2, 363, 362, 3, 2, 2, 2, 363, 364, 3,
	2, 2, 2, 364, 365, 3, 2, 2, 2, 365, 367, 7, 22, 2, 2, 366, 368, 5, 24,
	13, 2, 367, 366, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2,
	369, 371, 7, 5, 2, 2, 370, 372, 5, 80, 41, 2, 371, 370, 3, 2, 2, 2, 372,
	373, 3, 2, 2, 2, 373, 371, 3, 2, 2, 2, 373, 374, 3, 2, 2, 2, 374, 375,
	3, 2, 2, 2, 375, 376, 7, 6, 2, 2, 376, 79, 3, 2, 2, 2, 377, 379, 5, 8,
	5, 2, 378, 377, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2,
	380, 381, 5, 6, 4, 2, 381, 382, 7, 7, 2, 2, 382, 383, 5, 46, 24, 2, 383,
	81, 3, 2, 2, 2, 384, 391, 5, 86, 44, 2, 385, 391, 5, 90, 46, 2, 386, 391,
	5, 104, 53, 2, 387, 391, 5, 108, 55, 2, 388, 391, 5, 116, 59, 2, 389, 391,
	5, 124, 63, 2, 390, 384, 3, 2, 2, 2, 390, 385, 3, 2, 2, 2, 390, 386, 3,
	2, 2, 2, 390, 387, 3, 2, 2, 2, 390, 388, 3, 2, 2, 2, 390, 389, 3, 2, 2,
	2, 391, 83, 3, 2, 2, 2, 392, 399, 5, 92, 47, 2, 393, 399, 5, 106, 54, 2,
	394, 399, 5, 110, 56, 2, 395, 399, 5, 88, 45, 2, 396, 399, 5, 118, 60,
	2, 397, 399, 5, 126, 64, 2, 398, 392, 3, 2, 2, 2, 398, 393, 3, 2, 2, 2,
	398, 394, 3, 2, 2, 2, 398, 395, 3, 2, 2, 2, 398, 396, 3, 2, 2, 2, 398,
	397, 3, 2, 2, 2, 399, 85, 3, 2, 2, 2, 400, 402, 5, 8, 5, 2, 401, 400, 3,
	2, 2, 2, 401, 402, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403, 404, 7, 23, 2,
	2, 404, 406, 5, 32, 17, 2, 405, 407, 5, 24, 13, 2, 406, 405, 3, 2, 2, 2,
	406, 407, 3, 2, 2, 2, 407, 87, 3, 2, 2, 2, 408, 409, 7, 30, 2, 2, 409,
	410, 7, 23, 2, 2, 410, 412, 5, 32, 17, 2, 411, 413, 5, 24, 13, 2, 412,
	411, 3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 89, 3, 2, 2, 2, 414, 416, 5,
	8, 5, 2, 415, 414, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 417, 3, 2, 2,
	2, 417, 418, 7, 24, 2, 2, 418, 420, 5, 32, 17, 2, 419, 421, 5, 94, 48,
	2, 420, 419, 3, 2, 2, 2, 420, 421, 3, 2, 2, 2, 421, 423, 3, 2, 2, 2, 422,
	424, 5, 24, 13, 2, 423, 422, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424, 426,
	3, 2, 2, 2, 425, 427, 5, 96, 49, 2, 426, 425, 3, 2, 2, 2, 426, 427, 3,
	2, 2, 2, 427, 91, 3, 2, 2, 2, 428, 429, 7, 30, 2, 2, 429, 430, 7, 24, 2,
	2, 430, 432, 5, 32, 17, 2, 431, 433, 5, 94, 48, 2, 432, 431, 3, 2, 2, 2,
	432, 433, 3, 2, 2, 2, 433, 435, 3, 2, 2, 2, 434, 436, 5, 24, 13, 2, 435,
	434, 3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436, 438, 3, 2, 2, 2, 437, 439,
	5, 96, 49, 2, 438, 437, 3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 93, 3, 2,
	2, 2, 440, 441, 8, 48, 1, 2, 441, 443, 7, 26, 2, 2, 442, 444, 7, 16, 2,
	2, 443, 442, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 446, 3, 2, 2, 2, 445,
	447, 5, 46, 24, 2, 446, 445, 3, 2, 2, 2, 447, 448, 3, 2, 2, 2, 448, 446,
	3, 2, 2, 2, 448, 449, 3, 2, 2, 2, 449, 455, 3, 2, 2, 2, 450, 451, 12, 3,
	2, 2, 451, 452, 7, 16, 2, 2, 452, 454, 5, 46, 24, 2, 453, 450, 3, 2, 2,
	2, 454, 457, 3, 2, 2, 2, 455, 453, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456,
	95, 3, 2, 2, 2, 457, 455, 3, 2, 2, 2, 458, 460, 7, 5, 2, 2, 459, 461, 5,
	98, 50, 2, 460, 459, 3, 2, 2, 2, 461, 462, 3, 2, 2, 2, 462, 460, 3, 2,
	2, 2, 462, 463, 3, 2, 2, 2, 463, 464, 3, 2, 2, 2, 464, 465, 7, 6, 2, 2,
	465, 97, 3, 2, 2, 2, 466, 468, 5, 8, 5, 2, 467, 466, 3, 2, 2, 2, 467, 468,
	3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 471, 5, 32, 17, 2, 470, 472, 5,
	100, 51, 2, 471, 470, 3, 2, 2, 2, 471, 472, 3, 2, 2, 2, 472, 473, 3, 2,
	2, 2, 473, 474, 7, 7, 2, 2, 474, 476, 5, 44, 23, 2, 475, 477, 5, 24, 13,
	2, 476, 475, 3, 2, 2, 2, 476, 477, 3, 2, 2, 2, 477, 99, 3, 2, 2, 2, 478,
	480, 7, 9, 2, 2, 479, 481, 5, 102, 52, 2, 480, 479, 3, 2, 2, 2, 481, 482,
	3, 2, 2, 2, 482, 480, 3, 2, 2, 2, 482, 483, 3, 2, 2, 2, 483, 484, 3, 2,
	2, 2, 484, 485, 7, 10, 2, 2, 485, 101, 3, 2, 2, 2, 486, 488, 5, 8, 5, 2,
	487, 486, 3, 2, 2, 2, 487, 488, 3, 2, 2, 2, 488, 489, 3, 2, 2, 2, 489,
	490, 5, 32, 17, 2, 490, 491, 7, 7, 2, 2, 491, 493, 5, 44, 23, 2, 492, 494,
	5, 40, 21, 2, 493, 492, 3, 2, 2, 2, 493, 494, 3, 2, 2, 2, 494, 496, 3,
	2, 2, 2, 495, 497, 5, 24, 13, 2, 496, 495, 3, 2, 2, 2, 496, 497, 3, 2,
	2, 2, 497, 103, 3, 2, 2, 2, 498, 500, 5, 8, 5, 2, 499, 498, 3, 2, 2, 2,
	499, 500, 3, 2, 2, 2, 500, 501, 3, 2, 2, 2, 501, 502, 7, 25, 2, 2, 502,
	504, 5, 32, 17, 2, 503, 505, 5, 24, 13, 2, 504, 503, 3, 2, 2, 2, 504, 505,
	3, 2, 2, 2, 505, 507, 3, 2, 2, 2, 506, 508, 5, 96, 49, 2, 507, 506, 3,
	2, 2, 2, 507, 508, 3, 2, 2, 2, 508, 105, 3, 2, 2, 2, 509, 510, 7, 30, 2,
	2, 510, 511, 7, 25, 2, 2, 511, 513, 5, 32, 17, 2, 512, 514, 5, 24, 13,
	2, 513, 512, 3, 2, 2, 2, 513, 514, 3, 2, 2, 2, 514, 516, 3, 2, 2, 2, 515,
	517, 5, 96, 49, 2, 516, 515, 3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517, 107,
	3, 2, 2, 2, 518, 520, 5, 8, 5, 2, 519, 518, 3, 2, 2, 2, 519, 520, 3, 2,
	2, 2, 520, 521, 3, 2, 2, 2, 521, 522, 7, 28, 2, 2, 522, 524, 5, 32, 17,
	2, 523, 525, 5, 24, 13, 2, 524, 523, 3, 2, 2, 2, 524, 525, 3, 2, 2, 2,
	525, 526, 3, 2, 2, 2, 526, 527, 5, 112, 57, 2, 527, 109, 3, 2, 2, 2, 528,
	529, 7, 30, 2, 2, 529, 530, 7, 28, 2, 2, 530, 532, 5, 32, 17, 2, 531, 533,
	5, 24, 13, 2, 532, 531, 3, 2, 2, 2, 532, 533, 3, 2, 2, 2, 533, 535, 3,
	2, 2, 2, 534, 536, 5, 112, 57, 2, 535, 534, 3, 2, 2, 2, 535, 536, 3, 2,
	2, 2, 536, 111, 3, 2, 2, 2, 537, 538, 7, 12, 2, 2, 538, 539, 5, 114, 58,
	2, 539, 113, 3, 2, 2, 2, 540, 542, 8, 58, 1, 2, 541, 543, 7, 17, 2, 2,
	542, 541, 3, 2, 2, 2, 542, 543, 3, 2, 2, 2, 543, 544, 3, 2, 2, 2, 544,
	545, 5, 46, 24, 2, 545, 551, 3, 2, 2, 2, 546, 547, 12, 3, 2, 2, 547, 548,
	7, 17, 2, 2, 548, 550, 5, 46, 24, 2, 549, 546, 3, 2, 2, 2, 550, 553, 3,
	2, 2, 2, 551, 549, 3, 2, 2, 2, 551, 552, 3, 2, 2, 2, 552, 115, 3, 2, 2,
	2, 553, 551, 3, 2, 2, 2, 554, 556, 5, 8, 5, 2, 555, 554, 3, 2, 2, 2, 555,
	556, 3, 2, 2, 2, 556, 557, 3, 2, 2, 2, 557, 558, 7, 27, 2, 2, 558, 560,
	5, 32, 17, 2, 559, 561, 5, 24, 13, 2, 560, 559, 3, 2, 2, 2, 560, 561, 3,
	2, 2, 2, 561, 562, 3, 2, 2, 2, 562, 563, 5, 120, 61, 2, 563, 117, 3, 2,
	2, 2, 564, 565, 7, 30, 2, 2, 565, 566, 7, 27, 2, 2, 566, 568, 5, 32, 17,
	2, 567, 569, 5, 24, 13, 2, 568, 567, 3, 2, 2, 2, 568, 569, 3, 2, 2, 2,
	569, 571, 3, 2, 2, 2, 570, 572, 5, 120, 61, 2, 571, 570, 3, 2, 2, 2, 571,
	572, 3, 2, 2, 2, 572, 119, 3, 2, 2, 2, 573, 575, 7, 5, 2, 2, 574, 576,
	5, 122, 62, 2, 575, 574, 3, 2, 2, 2, 576, 577, 3, 2, 2, 2, 577, 575, 3,
	2, 2, 2, 577, 578, 3, 2, 2, 2, 578, 579, 3, 2, 2, 2, 579, 580, 7, 6, 2,
	2, 580, 121, 3, 2, 2, 2, 581, 583, 5, 8, 5, 2, 582, 581, 3, 2, 2, 2, 582,
	583, 3, 2, 2, 2, 583, 584, 3, 2, 2, 2, 584, 586, 5, 10, 6, 2, 585, 587,
	5, 24, 13, 2, 586, 585, 3, 2, 2, 2, 586, 587, 3, 2, 2, 2, 587, 123, 3,
	2, 2, 2, 588, 590, 5, 8, 5, 2, 589, 588, 3, 2, 2, 2, 589, 590, 3, 2, 2,
	2, 590, 591, 3, 2, 2, 2, 591, 592, 7, 29, 2, 2, 592, 594, 5, 32, 17, 2,
	593, 595, 5, 24, 13, 2, 594, 593, 3, 2, 2, 2, 594, 595, 3, 2, 2, 2, 595,
	596, 3, 2, 2, 2, 596, 597, 5, 128, 65, 2, 597, 125, 3, 2, 2, 2, 598, 599,
	7, 30, 2, 2, 599, 600, 7, 29, 2, 2, 600, 602, 5, 32, 17, 2, 601, 603, 5,
	24, 13, 2, 602, 601, 3, 2, 2, 2, 602, 603, 3, 2, 2, 2, 603, 605, 3, 2,
	2, 2, 604, 606, 5, 128, 65, 2, 605, 604, 3, 2, 2, 2, 605, 606, 3, 2, 2,
	2, 606, 127, 3, 2, 2, 2, 607, 609, 7, 5, 2, 2, 608, 610, 5, 102, 52, 2,
	609, 608, 3, 2, 2, 2, 610, 611, 3, 2, 2, 2, 611, 609, 3, 2, 2, 2, 611,
	612, 3, 2, 2, 2, 612, 613, 3, 2, 2, 2, 613, 614, 7, 6, 2, 2, 614, 129,
	3, 2, 2, 2, 615, 617, 5, 8, 5, 2, 616, 615, 3, 2, 2, 2, 616, 617, 3, 2,
	2, 2, 617, 618, 3, 2, 2, 2, 618, 619, 7, 31, 2, 2, 619, 620, 7, 8, 2, 2,
	620, 622, 5, 32, 17, 2, 621, 623, 5, 100, 51, 2, 622, 621, 3, 2, 2, 2,
	622, 623, 3, 2, 2, 2, 623, 624, 3, 2, 2, 2, 624, 625, 7, 15, 2, 2, 625,
	626, 5, 134, 68, 2, 626, 131, 3, 2, 2, 2, 627, 628, 5, 32, 17, 2, 628,
	133, 3, 2, 2, 2, 629, 630, 8, 68, 1, 2, 630, 631, 5, 132, 67, 2, 631, 637,
	3, 2, 2, 2, 632, 633, 12, 3, 2, 2, 633, 634, 7, 17, 2, 2, 634, 636, 5,
	132, 67, 2, 635, 632, 3, 2, 2, 2, 636, 639, 3, 2, 2, 2, 637, 635, 3, 2,
	2, 2, 637, 638, 3, 2, 2, 2, 638, 135, 3, 2, 2, 2, 639, 637, 3, 2, 2, 2,
	86, 139, 144, 156, 165, 174, 183, 199, 204, 210, 228, 239, 252, 266, 271,
	274, 277, 281, 287, 295, 301, 308, 311, 315, 318, 321, 329, 333, 336, 344,
	354, 360, 363, 367, 373, 378, 390, 398, 401, 406, 412, 415, 420, 423, 426,
	432, 435, 438, 443, 448, 455, 462, 467, 471, 476, 482, 487, 493, 496, 499,
	504, 507, 513, 516, 519, 524, 532, 535, 542, 551, 555, 560, 568, 571, 577,
	582, 586, 589, 594, 602, 605, 611, 616, 622, 637,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'['", "']'", "'{'", "'}'", "':'", "'@'", "'('", "')'", "'$'", "'='",
	"'!'", "'...'", "'on'", "'&'", "'|'", "'fragment'", "'query'", "'mutation'",
	"'subscription'", "'schema'", "'scalar'", "'type'", "'interface'", "'implements'",
	"'enum'", "'union'", "'input'", "'extend'", "'directive'", "", "", "'null'",
	"", "", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "FRAGMENT",
	"QUERY", "MUTATION", "SUBSCRIPTION", "SCHEMA", "SCALAR", "TYPE", "INTERFACE",
	"IMPLEMENTS", "ENUM", "UNION", "INPUT", "EXTEND", "DIRECTIVE", "NAME",
	"BooleanValue", "NullValue", "IntValue", "FloatValue", "Sign", "IntegerPart",
	"NonZeroDigit", "ExponentPart", "Digit", "StringValue", "TripleQuotedStringValue",
	"Comment", "Ignored",
}

var ruleNames = []string{
	"document", "definition", "operationType", "description", "enumValue",
	"arrayValue", "arrayValueWithVariable", "objectValue", "objectValueWithVariable",
	"objectField", "objectFieldWithVariable", "directives", "directive", "arguments",
	"argument", "name", "value", "valueWithVariable", "variable", "defaultValue",
	"stringValue", "objType", "typeName", "listType", "nonNullType", "operationDefinition",
	"variableDefinitions", "variableDefinition", "selectionSet", "selection",
	"field", "alias", "fragmentSpread", "inlineFragment", "fragmentDefinition",
	"fragmentName", "typeCondition", "typeSystemDefinition", "schemaDefinition",
	"operationTypeDefinition", "typeDefinition", "typeExtension", "scalarTypeDefinition",
	"scalarTypeExtensionDefinition", "objectTypeDefinition", "objectTypeExtensionDefinition",
	"implementsInterfaces", "fieldsDefinition", "fieldDefinition", "argumentsDefinition",
	"inputValueDefinition", "interfaceTypeDefinition", "interfaceTypeExtensionDefinition",
	"unionTypeDefinition", "unionTypeExtensionDefinition", "unionMembership",
	"unionMembers", "enumTypeDefinition", "enumTypeExtensionDefinition", "enumValueDefinitions",
	"enumValueDefinition", "inputObjectTypeDefinition", "inputObjectTypeExtensionDefinition",
	"inputObjectValueDefinitions", "directiveDefinition", "directiveLocation",
	"directiveLocations",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GraphQLParser struct {
	*antlr.BaseParser
}

func NewGraphQLParser(input antlr.TokenStream) *GraphQLParser {
	this := new(GraphQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GraphQL.g4"

	return this
}

// GraphQLParser tokens.
const (
	GraphQLParserEOF                     = antlr.TokenEOF
	GraphQLParserT__0                    = 1
	GraphQLParserT__1                    = 2
	GraphQLParserT__2                    = 3
	GraphQLParserT__3                    = 4
	GraphQLParserT__4                    = 5
	GraphQLParserT__5                    = 6
	GraphQLParserT__6                    = 7
	GraphQLParserT__7                    = 8
	GraphQLParserT__8                    = 9
	GraphQLParserT__9                    = 10
	GraphQLParserT__10                   = 11
	GraphQLParserT__11                   = 12
	GraphQLParserT__12                   = 13
	GraphQLParserT__13                   = 14
	GraphQLParserT__14                   = 15
	GraphQLParserFRAGMENT                = 16
	GraphQLParserQUERY                   = 17
	GraphQLParserMUTATION                = 18
	GraphQLParserSUBSCRIPTION            = 19
	GraphQLParserSCHEMA                  = 20
	GraphQLParserSCALAR                  = 21
	GraphQLParserTYPE                    = 22
	GraphQLParserINTERFACE               = 23
	GraphQLParserIMPLEMENTS              = 24
	GraphQLParserENUM                    = 25
	GraphQLParserUNION                   = 26
	GraphQLParserINPUT                   = 27
	GraphQLParserEXTEND                  = 28
	GraphQLParserDIRECTIVE               = 29
	GraphQLParserNAME                    = 30
	GraphQLParserBooleanValue            = 31
	GraphQLParserNullValue               = 32
	GraphQLParserIntValue                = 33
	GraphQLParserFloatValue              = 34
	GraphQLParserSign                    = 35
	GraphQLParserIntegerPart             = 36
	GraphQLParserNonZeroDigit            = 37
	GraphQLParserExponentPart            = 38
	GraphQLParserDigit                   = 39
	GraphQLParserStringValue             = 40
	GraphQLParserTripleQuotedStringValue = 41
	GraphQLParserComment                 = 42
	GraphQLParserIgnored                 = 43
)

// GraphQLParser rules.
const (
	GraphQLParserRULE_document                           = 0
	GraphQLParserRULE_definition                         = 1
	GraphQLParserRULE_operationType                      = 2
	GraphQLParserRULE_description                        = 3
	GraphQLParserRULE_enumValue                          = 4
	GraphQLParserRULE_arrayValue                         = 5
	GraphQLParserRULE_arrayValueWithVariable             = 6
	GraphQLParserRULE_objectValue                        = 7
	GraphQLParserRULE_objectValueWithVariable            = 8
	GraphQLParserRULE_objectField                        = 9
	GraphQLParserRULE_objectFieldWithVariable            = 10
	GraphQLParserRULE_directives                         = 11
	GraphQLParserRULE_directive                          = 12
	GraphQLParserRULE_arguments                          = 13
	GraphQLParserRULE_argument                           = 14
	GraphQLParserRULE_name                               = 15
	GraphQLParserRULE_value                              = 16
	GraphQLParserRULE_valueWithVariable                  = 17
	GraphQLParserRULE_variable                           = 18
	GraphQLParserRULE_defaultValue                       = 19
	GraphQLParserRULE_stringValue                        = 20
	GraphQLParserRULE_objType                            = 21
	GraphQLParserRULE_typeName                           = 22
	GraphQLParserRULE_listType                           = 23
	GraphQLParserRULE_nonNullType                        = 24
	GraphQLParserRULE_operationDefinition                = 25
	GraphQLParserRULE_variableDefinitions                = 26
	GraphQLParserRULE_variableDefinition                 = 27
	GraphQLParserRULE_selectionSet                       = 28
	GraphQLParserRULE_selection                          = 29
	GraphQLParserRULE_field                              = 30
	GraphQLParserRULE_alias                              = 31
	GraphQLParserRULE_fragmentSpread                     = 32
	GraphQLParserRULE_inlineFragment                     = 33
	GraphQLParserRULE_fragmentDefinition                 = 34
	GraphQLParserRULE_fragmentName                       = 35
	GraphQLParserRULE_typeCondition                      = 36
	GraphQLParserRULE_typeSystemDefinition               = 37
	GraphQLParserRULE_schemaDefinition                   = 38
	GraphQLParserRULE_operationTypeDefinition            = 39
	GraphQLParserRULE_typeDefinition                     = 40
	GraphQLParserRULE_typeExtension                      = 41
	GraphQLParserRULE_scalarTypeDefinition               = 42
	GraphQLParserRULE_scalarTypeExtensionDefinition      = 43
	GraphQLParserRULE_objectTypeDefinition               = 44
	GraphQLParserRULE_objectTypeExtensionDefinition      = 45
	GraphQLParserRULE_implementsInterfaces               = 46
	GraphQLParserRULE_fieldsDefinition                   = 47
	GraphQLParserRULE_fieldDefinition                    = 48
	GraphQLParserRULE_argumentsDefinition                = 49
	GraphQLParserRULE_inputValueDefinition               = 50
	GraphQLParserRULE_interfaceTypeDefinition            = 51
	GraphQLParserRULE_interfaceTypeExtensionDefinition   = 52
	GraphQLParserRULE_unionTypeDefinition                = 53
	GraphQLParserRULE_unionTypeExtensionDefinition       = 54
	GraphQLParserRULE_unionMembership                    = 55
	GraphQLParserRULE_unionMembers                       = 56
	GraphQLParserRULE_enumTypeDefinition                 = 57
	GraphQLParserRULE_enumTypeExtensionDefinition        = 58
	GraphQLParserRULE_enumValueDefinitions               = 59
	GraphQLParserRULE_enumValueDefinition                = 60
	GraphQLParserRULE_inputObjectTypeDefinition          = 61
	GraphQLParserRULE_inputObjectTypeExtensionDefinition = 62
	GraphQLParserRULE_inputObjectValueDefinitions        = 63
	GraphQLParserRULE_directiveDefinition                = 64
	GraphQLParserRULE_directiveLocation                  = 65
	GraphQLParserRULE_directiveLocations                 = 66
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *DocumentContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *GraphQLParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraphQLParserRULE_document)
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
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__2)|(1<<GraphQLParserFRAGMENT)|(1<<GraphQLParserQUERY)|(1<<GraphQLParserMUTATION)|(1<<GraphQLParserSUBSCRIPTION)|(1<<GraphQLParserSCHEMA)|(1<<GraphQLParserSCALAR)|(1<<GraphQLParserTYPE)|(1<<GraphQLParserINTERFACE)|(1<<GraphQLParserENUM)|(1<<GraphQLParserUNION)|(1<<GraphQLParserINPUT)|(1<<GraphQLParserEXTEND)|(1<<GraphQLParserDIRECTIVE))) != 0) || _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(134)
			p.Definition()
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) OperationDefinition() IOperationDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationDefinitionContext)
}

func (s *DefinitionContext) FragmentDefinition() IFragmentDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentDefinitionContext)
}

func (s *DefinitionContext) TypeSystemDefinition() ITypeSystemDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSystemDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSystemDefinitionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *GraphQLParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraphQLParserRULE_definition)

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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__2, GraphQLParserQUERY, GraphQLParserMUTATION, GraphQLParserSUBSCRIPTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.OperationDefinition()
		}

	case GraphQLParserFRAGMENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.FragmentDefinition()
		}

	case GraphQLParserSCHEMA, GraphQLParserSCALAR, GraphQLParserTYPE, GraphQLParserINTERFACE, GraphQLParserENUM, GraphQLParserUNION, GraphQLParserINPUT, GraphQLParserEXTEND, GraphQLParserDIRECTIVE, GraphQLParserStringValue, GraphQLParserTripleQuotedStringValue:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(141)
			p.TypeSystemDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationTypeContext is an interface to support dynamic dispatch.
type IOperationTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationTypeContext differentiates from other interfaces.
	IsOperationTypeContext()
}

type OperationTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTypeContext() *OperationTypeContext {
	var p = new(OperationTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationType
	return p
}

func (*OperationTypeContext) IsOperationTypeContext() {}

func NewOperationTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTypeContext {
	var p = new(OperationTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationType

	return p
}

func (s *OperationTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationTypeContext) SUBSCRIPTION() antlr.TerminalNode {
	return s.GetToken(GraphQLParserSUBSCRIPTION, 0)
}

func (s *OperationTypeContext) MUTATION() antlr.TerminalNode {
	return s.GetToken(GraphQLParserMUTATION, 0)
}

func (s *OperationTypeContext) QUERY() antlr.TerminalNode {
	return s.GetToken(GraphQLParserQUERY, 0)
}

func (s *OperationTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationType(s)
	}
}

func (s *OperationTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationType(s)
	}
}

func (p *GraphQLParser) OperationType() (localctx IOperationTypeContext) {
	localctx = NewOperationTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GraphQLParserRULE_operationType)
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
		p.SetState(144)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserQUERY)|(1<<GraphQLParserMUTATION)|(1<<GraphQLParserSUBSCRIPTION))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_description
	return p
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (p *GraphQLParser) Description() (localctx IDescriptionContext) {
	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GraphQLParserRULE_description)

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
		p.SetState(146)
		p.StringValue()
	}

	return localctx
}

// IEnumValueContext is an interface to support dynamic dispatch.
type IEnumValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueContext differentiates from other interfaces.
	IsEnumValueContext()
}

type EnumValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueContext() *EnumValueContext {
	var p = new(EnumValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValue
	return p
}

func (*EnumValueContext) IsEnumValueContext() {}

func NewEnumValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueContext {
	var p = new(EnumValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValue

	return p
}

func (s *EnumValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValue(s)
	}
}

func (s *EnumValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValue(s)
	}
}

func (p *GraphQLParser) EnumValue() (localctx IEnumValueContext) {
	localctx = NewEnumValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GraphQLParserRULE_enumValue)

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
		p.SetState(148)
		p.Name()
	}

	return localctx
}

// IArrayValueContext is an interface to support dynamic dispatch.
type IArrayValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayValueContext differentiates from other interfaces.
	IsArrayValueContext()
}

type ArrayValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayValueContext() *ArrayValueContext {
	var p = new(ArrayValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_arrayValue
	return p
}

func (*ArrayValueContext) IsArrayValueContext() {}

func NewArrayValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayValueContext {
	var p = new(ArrayValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_arrayValue

	return p
}

func (s *ArrayValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayValueContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrayValueContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

func (p *GraphQLParser) ArrayValue() (localctx IArrayValueContext) {
	localctx = NewArrayValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GraphQLParserRULE_arrayValue)
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
		p.SetState(150)
		p.Match(GraphQLParserT__0)
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__2)|(1<<GraphQLParserFRAGMENT)|(1<<GraphQLParserQUERY)|(1<<GraphQLParserMUTATION)|(1<<GraphQLParserSUBSCRIPTION)|(1<<GraphQLParserSCHEMA)|(1<<GraphQLParserSCALAR)|(1<<GraphQLParserTYPE)|(1<<GraphQLParserINTERFACE)|(1<<GraphQLParserIMPLEMENTS)|(1<<GraphQLParserENUM)|(1<<GraphQLParserUNION)|(1<<GraphQLParserINPUT)|(1<<GraphQLParserEXTEND)|(1<<GraphQLParserDIRECTIVE)|(1<<GraphQLParserNAME)|(1<<GraphQLParserBooleanValue))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GraphQLParserNullValue-32))|(1<<(GraphQLParserIntValue-32))|(1<<(GraphQLParserFloatValue-32))|(1<<(GraphQLParserStringValue-32))|(1<<(GraphQLParserTripleQuotedStringValue-32)))) != 0) {
		{
			p.SetState(151)
			p.Value()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(157)
		p.Match(GraphQLParserT__1)
	}

	return localctx
}

// IArrayValueWithVariableContext is an interface to support dynamic dispatch.
type IArrayValueWithVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayValueWithVariableContext differentiates from other interfaces.
	IsArrayValueWithVariableContext()
}

type ArrayValueWithVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayValueWithVariableContext() *ArrayValueWithVariableContext {
	var p = new(ArrayValueWithVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_arrayValueWithVariable
	return p
}

func (*ArrayValueWithVariableContext) IsArrayValueWithVariableContext() {}

func NewArrayValueWithVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayValueWithVariableContext {
	var p = new(ArrayValueWithVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_arrayValueWithVariable

	return p
}

func (s *ArrayValueWithVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayValueWithVariableContext) AllValueWithVariable() []IValueWithVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueWithVariableContext)(nil)).Elem())
	var tst = make([]IValueWithVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueWithVariableContext)
		}
	}

	return tst
}

func (s *ArrayValueWithVariableContext) ValueWithVariable(i int) IValueWithVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueWithVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueWithVariableContext)
}

func (s *ArrayValueWithVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueWithVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayValueWithVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArrayValueWithVariable(s)
	}
}

func (s *ArrayValueWithVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArrayValueWithVariable(s)
	}
}

func (p *GraphQLParser) ArrayValueWithVariable() (localctx IArrayValueWithVariableContext) {
	localctx = NewArrayValueWithVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GraphQLParserRULE_arrayValueWithVariable)
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
		p.SetState(159)
		p.Match(GraphQLParserT__0)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__8)|(1<<GraphQLParserFRAGMENT)|(1<<GraphQLParserQUERY)|(1<<GraphQLParserMUTATION)|(1<<GraphQLParserSUBSCRIPTION)|(1<<GraphQLParserSCHEMA)|(1<<GraphQLParserSCALAR)|(1<<GraphQLParserTYPE)|(1<<GraphQLParserINTERFACE)|(1<<GraphQLParserIMPLEMENTS)|(1<<GraphQLParserENUM)|(1<<GraphQLParserUNION)|(1<<GraphQLParserINPUT)|(1<<GraphQLParserEXTEND)|(1<<GraphQLParserDIRECTIVE)|(1<<GraphQLParserNAME)|(1<<GraphQLParserBooleanValue))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GraphQLParserNullValue-32))|(1<<(GraphQLParserIntValue-32))|(1<<(GraphQLParserFloatValue-32))|(1<<(GraphQLParserStringValue-32))|(1<<(GraphQLParserTripleQuotedStringValue-32)))) != 0) {
		{
			p.SetState(160)
			p.ValueWithVariable()
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.Match(GraphQLParserT__1)
	}

	return localctx
}

// IObjectValueContext is an interface to support dynamic dispatch.
type IObjectValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectValueContext differentiates from other interfaces.
	IsObjectValueContext()
}

type ObjectValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectValueContext() *ObjectValueContext {
	var p = new(ObjectValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectValue
	return p
}

func (*ObjectValueContext) IsObjectValueContext() {}

func NewObjectValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectValueContext {
	var p = new(ObjectValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectValue

	return p
}

func (s *ObjectValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectValueContext) AllObjectField() []IObjectFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectFieldContext)(nil)).Elem())
	var tst = make([]IObjectFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectFieldContext)
		}
	}

	return tst
}

func (s *ObjectValueContext) ObjectField(i int) IObjectFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjectFieldContext)
}

func (s *ObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectValue(s)
	}
}

func (s *ObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectValue(s)
	}
}

func (p *GraphQLParser) ObjectValue() (localctx IObjectValueContext) {
	localctx = NewObjectValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GraphQLParserRULE_objectValue)
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
		p.SetState(168)
		p.Match(GraphQLParserT__2)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserFRAGMENT)|(1<<GraphQLParserQUERY)|(1<<GraphQLParserMUTATION)|(1<<GraphQLParserSUBSCRIPTION)|(1<<GraphQLParserSCHEMA)|(1<<GraphQLParserSCALAR)|(1<<GraphQLParserTYPE)|(1<<GraphQLParserINTERFACE)|(1<<GraphQLParserIMPLEMENTS)|(1<<GraphQLParserENUM)|(1<<GraphQLParserUNION)|(1<<GraphQLParserINPUT)|(1<<GraphQLParserEXTEND)|(1<<GraphQLParserDIRECTIVE)|(1<<GraphQLParserNAME))) != 0 {
		{
			p.SetState(169)
			p.ObjectField()
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(175)
		p.Match(GraphQLParserT__3)
	}

	return localctx
}

// IObjectValueWithVariableContext is an interface to support dynamic dispatch.
type IObjectValueWithVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectValueWithVariableContext differentiates from other interfaces.
	IsObjectValueWithVariableContext()
}

type ObjectValueWithVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectValueWithVariableContext() *ObjectValueWithVariableContext {
	var p = new(ObjectValueWithVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectValueWithVariable
	return p
}

func (*ObjectValueWithVariableContext) IsObjectValueWithVariableContext() {}

func NewObjectValueWithVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectValueWithVariableContext {
	var p = new(ObjectValueWithVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectValueWithVariable

	return p
}

func (s *ObjectValueWithVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectValueWithVariableContext) AllObjectFieldWithVariable() []IObjectFieldWithVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectFieldWithVariableContext)(nil)).Elem())
	var tst = make([]IObjectFieldWithVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectFieldWithVariableContext)
		}
	}

	return tst
}

func (s *ObjectValueWithVariableContext) ObjectFieldWithVariable(i int) IObjectFieldWithVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectFieldWithVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjectFieldWithVariableContext)
}

func (s *ObjectValueWithVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueWithVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectValueWithVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectValueWithVariable(s)
	}
}

func (s *ObjectValueWithVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectValueWithVariable(s)
	}
}

func (p *GraphQLParser) ObjectValueWithVariable() (localctx IObjectValueWithVariableContext) {
	localctx = NewObjectValueWithVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GraphQLParserRULE_objectValueWithVariable)
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
		p.SetState(177)
		p.Match(GraphQLParserT__2)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserFRAGMENT)|(1<<GraphQLParserQUERY)|(1<<GraphQLParserMUTATION)|(1<<GraphQLParserSUBSCRIPTION)|(1<<GraphQLParserSCHEMA)|(1<<GraphQLParserSCALAR)|(1<<GraphQLParserTYPE)|(1<<GraphQLParserINTERFACE)|(1<<GraphQLParserIMPLEMENTS)|(1<<GraphQLParserENUM)|(1<<GraphQLParserUNION)|(1<<GraphQLParserINPUT)|(1<<GraphQLParserEXTEND)|(1<<GraphQLParserDIRECTIVE)|(1<<GraphQLParserNAME))) != 0 {
		{
			p.SetState(178)
			p.ObjectFieldWithVariable()
		}

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
		p.Match(GraphQLParserT__3)
	}

	return localctx
}

// IObjectFieldContext is an interface to support dynamic dispatch.
type IObjectFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectFieldContext differentiates from other interfaces.
	IsObjectFieldContext()
}

type ObjectFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectFieldContext() *ObjectFieldContext {
	var p = new(ObjectFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectField
	return p
}

func (*ObjectFieldContext) IsObjectFieldContext() {}

func NewObjectFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectFieldContext {
	var p = new(ObjectFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectField

	return p
}

func (s *ObjectFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectFieldContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectFieldContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ObjectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectField(s)
	}
}

func (s *ObjectFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectField(s)
	}
}

func (p *GraphQLParser) ObjectField() (localctx IObjectFieldContext) {
	localctx = NewObjectFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GraphQLParserRULE_objectField)

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
		p.SetState(186)
		p.Name()
	}
	{
		p.SetState(187)
		p.Match(GraphQLParserT__4)
	}
	{
		p.SetState(188)
		p.Value()
	}

	return localctx
}

// IObjectFieldWithVariableContext is an interface to support dynamic dispatch.
type IObjectFieldWithVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectFieldWithVariableContext differentiates from other interfaces.
	IsObjectFieldWithVariableContext()
}

type ObjectFieldWithVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectFieldWithVariableContext() *ObjectFieldWithVariableContext {
	var p = new(ObjectFieldWithVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectFieldWithVariable
	return p
}

func (*ObjectFieldWithVariableContext) IsObjectFieldWithVariableContext() {}

func NewObjectFieldWithVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectFieldWithVariableContext {
	var p = new(ObjectFieldWithVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectFieldWithVariable

	return p
}

func (s *ObjectFieldWithVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectFieldWithVariableContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectFieldWithVariableContext) ValueWithVariable() IValueWithVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueWithVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueWithVariableContext)
}

func (s *ObjectFieldWithVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectFieldWithVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectFieldWithVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectFieldWithVariable(s)
	}
}

func (s *ObjectFieldWithVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectFieldWithVariable(s)
	}
}

func (p *GraphQLParser) ObjectFieldWithVariable() (localctx IObjectFieldWithVariableContext) {
	localctx = NewObjectFieldWithVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GraphQLParserRULE_objectFieldWithVariable)

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
		p.SetState(190)
		p.Name()
	}
	{
		p.SetState(191)
		p.Match(GraphQLParserT__4)
	}
	{
		p.SetState(192)
		p.ValueWithVariable()
	}

	return localctx
}

// IDirectivesContext is an interface to support dynamic dispatch.
type IDirectivesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectivesContext differentiates from other interfaces.
	IsDirectivesContext()
}

type DirectivesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectivesContext() *DirectivesContext {
	var p = new(DirectivesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directives
	return p
}

func (*DirectivesContext) IsDirectivesContext() {}

func NewDirectivesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectivesContext {
	var p = new(DirectivesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directives

	return p
}

func (s *DirectivesContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectivesContext) AllDirective() []IDirectiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveContext)(nil)).Elem())
	var tst = make([]IDirectiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveContext)
		}
	}

	return tst
}

func (s *DirectivesContext) Directive(i int) IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *DirectivesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectivesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectivesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectives(s)
	}
}

func (s *DirectivesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectives(s)
	}
}

func (p *GraphQLParser) Directives() (localctx IDirectivesContext) {
	localctx = NewDirectivesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GraphQLParserRULE_directives)
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
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserT__5 {
		{
			p.SetState(194)
			p.Directive()
		}

		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DirectiveContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (p *GraphQLParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GraphQLParserRULE_directive)
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
		p.SetState(199)
		p.Match(GraphQLParserT__5)
	}
	{
		p.SetState(200)
		p.Name()
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__6 {
		{
			p.SetState(201)
			p.Arguments()
		}

	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *GraphQLParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GraphQLParserRULE_arguments)
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
		p.SetState(204)
		p.Match(GraphQLParserT__6)
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserFRAGMENT)|(1<<GraphQLParserQUERY)|(1<<GraphQLParserMUTATION)|(1<<GraphQLParserSUBSCRIPTION)|(1<<GraphQLParserSCHEMA)|(1<<GraphQLParserSCALAR)|(1<<GraphQLParserTYPE)|(1<<GraphQLParserINTERFACE)|(1<<GraphQLParserIMPLEMENTS)|(1<<GraphQLParserENUM)|(1<<GraphQLParserUNION)|(1<<GraphQLParserINPUT)|(1<<GraphQLParserEXTEND)|(1<<GraphQLParserDIRECTIVE)|(1<<GraphQLParserNAME))) != 0) {
		{
			p.SetState(205)
			p.Argument()
		}

		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(210)
		p.Match(GraphQLParserT__7)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ArgumentContext) ValueWithVariable() IValueWithVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueWithVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueWithVariableContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *GraphQLParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GraphQLParserRULE_argument)

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
		p.SetState(212)
		p.Name()
	}
	{
		p.SetState(213)
		p.Match(GraphQLParserT__4)
	}
	{
		p.SetState(214)
		p.ValueWithVariable()
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, 0)
}

func (s *NameContext) FRAGMENT() antlr.TerminalNode {
	return s.GetToken(GraphQLParserFRAGMENT, 0)
}

func (s *NameContext) QUERY() antlr.TerminalNode {
	return s.GetToken(GraphQLParserQUERY, 0)
}

func (s *NameContext) MUTATION() antlr.TerminalNode {
	return s.GetToken(GraphQLParserMUTATION, 0)
}

func (s *NameContext) SUBSCRIPTION() antlr.TerminalNode {
	return s.GetToken(GraphQLParserSUBSCRIPTION, 0)
}

func (s *NameContext) SCHEMA() antlr.TerminalNode {
	return s.GetToken(GraphQLParserSCHEMA, 0)
}

func (s *NameContext) SCALAR() antlr.TerminalNode {
	return s.GetToken(GraphQLParserSCALAR, 0)
}

func (s *NameContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GraphQLParserTYPE, 0)
}

func (s *NameContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(GraphQLParserINTERFACE, 0)
}

func (s *NameContext) IMPLEMENTS() antlr.TerminalNode {
	return s.GetToken(GraphQLParserIMPLEMENTS, 0)
}

func (s *NameContext) ENUM() antlr.TerminalNode {
	return s.GetToken(GraphQLParserENUM, 0)
}

func (s *NameContext) UNION() antlr.TerminalNode {
	return s.GetToken(GraphQLParserUNION, 0)
}

func (s *NameContext) INPUT() antlr.TerminalNode {
	return s.GetToken(GraphQLParserINPUT, 0)
}

func (s *NameContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(GraphQLParserEXTEND, 0)
}

func (s *NameContext) DIRECTIVE() antlr.TerminalNode {
	return s.GetToken(GraphQLParserDIRECTIVE, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *GraphQLParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GraphQLParserRULE_name)
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
		p.SetState(216)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserFRAGMENT)|(1<<GraphQLParserQUERY)|(1<<GraphQLParserMUTATION)|(1<<GraphQLParserSUBSCRIPTION)|(1<<GraphQLParserSCHEMA)|(1<<GraphQLParserSCALAR)|(1<<GraphQLParserTYPE)|(1<<GraphQLParserINTERFACE)|(1<<GraphQLParserIMPLEMENTS)|(1<<GraphQLParserENUM)|(1<<GraphQLParserUNION)|(1<<GraphQLParserINPUT)|(1<<GraphQLParserEXTEND)|(1<<GraphQLParserDIRECTIVE)|(1<<GraphQLParserNAME))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *ValueContext) IntValue() antlr.TerminalNode {
	return s.GetToken(GraphQLParserIntValue, 0)
}

func (s *ValueContext) FloatValue() antlr.TerminalNode {
	return s.GetToken(GraphQLParserFloatValue, 0)
}

func (s *ValueContext) BooleanValue() antlr.TerminalNode {
	return s.GetToken(GraphQLParserBooleanValue, 0)
}

func (s *ValueContext) NullValue() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNullValue, 0)
}

func (s *ValueContext) EnumValue() IEnumValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *ValueContext) ArrayValue() IArrayValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayValueContext)
}

func (s *ValueContext) ObjectValue() IObjectValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectValueContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *GraphQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GraphQLParserRULE_value)

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

	p.SetState(226)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserStringValue, GraphQLParserTripleQuotedStringValue:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.StringValue()
		}

	case GraphQLParserIntValue:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.Match(GraphQLParserIntValue)
		}

	case GraphQLParserFloatValue:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.Match(GraphQLParserFloatValue)
		}

	case GraphQLParserBooleanValue:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(221)
			p.Match(GraphQLParserBooleanValue)
		}

	case GraphQLParserNullValue:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(222)
			p.Match(GraphQLParserNullValue)
		}

	case GraphQLParserFRAGMENT, GraphQLParserQUERY, GraphQLParserMUTATION, GraphQLParserSUBSCRIPTION, GraphQLParserSCHEMA, GraphQLParserSCALAR, GraphQLParserTYPE, GraphQLParserINTERFACE, GraphQLParserIMPLEMENTS, GraphQLParserENUM, GraphQLParserUNION, GraphQLParserINPUT, GraphQLParserEXTEND, GraphQLParserDIRECTIVE, GraphQLParserNAME:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(223)
			p.EnumValue()
		}

	case GraphQLParserT__0:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(224)
			p.ArrayValue()
		}

	case GraphQLParserT__2:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(225)
			p.ObjectValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueWithVariableContext is an interface to support dynamic dispatch.
type IValueWithVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueWithVariableContext differentiates from other interfaces.
	IsValueWithVariableContext()
}

type ValueWithVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueWithVariableContext() *ValueWithVariableContext {
	var p = new(ValueWithVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_valueWithVariable
	return p
}

func (*ValueWithVariableContext) IsValueWithVariableContext() {}

func NewValueWithVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueWithVariableContext {
	var p = new(ValueWithVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_valueWithVariable

	return p
}

func (s *ValueWithVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueWithVariableContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ValueWithVariableContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *ValueWithVariableContext) IntValue() antlr.TerminalNode {
	return s.GetToken(GraphQLParserIntValue, 0)
}

func (s *ValueWithVariableContext) FloatValue() antlr.TerminalNode {
	return s.GetToken(GraphQLParserFloatValue, 0)
}

func (s *ValueWithVariableContext) BooleanValue() antlr.TerminalNode {
	return s.GetToken(GraphQLParserBooleanValue, 0)
}

func (s *ValueWithVariableContext) NullValue() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNullValue, 0)
}

func (s *ValueWithVariableContext) EnumValue() IEnumValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *ValueWithVariableContext) ArrayValueWithVariable() IArrayValueWithVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayValueWithVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayValueWithVariableContext)
}

func (s *ValueWithVariableContext) ObjectValueWithVariable() IObjectValueWithVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectValueWithVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectValueWithVariableContext)
}

func (s *ValueWithVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueWithVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueWithVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterValueWithVariable(s)
	}
}

func (s *ValueWithVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitValueWithVariable(s)
	}
}

func (p *GraphQLParser) ValueWithVariable() (localctx IValueWithVariableContext) {
	localctx = NewValueWithVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GraphQLParserRULE_valueWithVariable)

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

	p.SetState(237)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__8:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.Variable()
		}

	case GraphQLParserStringValue, GraphQLParserTripleQuotedStringValue:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.StringValue()
		}

	case GraphQLParserIntValue:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.Match(GraphQLParserIntValue)
		}

	case GraphQLParserFloatValue:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(231)
			p.Match(GraphQLParserFloatValue)
		}

	case GraphQLParserBooleanValue:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(232)
			p.Match(GraphQLParserBooleanValue)
		}

	case GraphQLParserNullValue:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(233)
			p.Match(GraphQLParserNullValue)
		}

	case GraphQLParserFRAGMENT, GraphQLParserQUERY, GraphQLParserMUTATION, GraphQLParserSUBSCRIPTION, GraphQLParserSCHEMA, GraphQLParserSCALAR, GraphQLParserTYPE, GraphQLParserINTERFACE, GraphQLParserIMPLEMENTS, GraphQLParserENUM, GraphQLParserUNION, GraphQLParserINPUT, GraphQLParserEXTEND, GraphQLParserDIRECTIVE, GraphQLParserNAME:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(234)
			p.EnumValue()
		}

	case GraphQLParserT__0:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(235)
			p.ArrayValueWithVariable()
		}

	case GraphQLParserT__2:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(236)
			p.ObjectValueWithVariable()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *GraphQLParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GraphQLParserRULE_variable)

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
		p.SetState(239)
		p.Match(GraphQLParserT__8)
	}
	{
		p.SetState(240)
		p.Name()
	}

	return localctx
}

// IDefaultValueContext is an interface to support dynamic dispatch.
type IDefaultValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultValueContext differentiates from other interfaces.
	IsDefaultValueContext()
}

type DefaultValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValueContext() *DefaultValueContext {
	var p = new(DefaultValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_defaultValue
	return p
}

func (*DefaultValueContext) IsDefaultValueContext() {}

func NewDefaultValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValueContext {
	var p = new(DefaultValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_defaultValue

	return p
}

func (s *DefaultValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValueContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DefaultValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDefaultValue(s)
	}
}

func (s *DefaultValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDefaultValue(s)
	}
}

func (p *GraphQLParser) DefaultValue() (localctx IDefaultValueContext) {
	localctx = NewDefaultValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GraphQLParserRULE_defaultValue)

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
		p.SetState(242)
		p.Match(GraphQLParserT__9)
	}
	{
		p.SetState(243)
		p.Value()
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) TripleQuotedStringValue() antlr.TerminalNode {
	return s.GetToken(GraphQLParserTripleQuotedStringValue, 0)
}

func (s *StringValueContext) StringValue() antlr.TerminalNode {
	return s.GetToken(GraphQLParserStringValue, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (p *GraphQLParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GraphQLParserRULE_stringValue)
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
		p.SetState(245)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IObjTypeContext is an interface to support dynamic dispatch.
type IObjTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjTypeContext differentiates from other interfaces.
	IsObjTypeContext()
}

type ObjTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjTypeContext() *ObjTypeContext {
	var p = new(ObjTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objType
	return p
}

func (*ObjTypeContext) IsObjTypeContext() {}

func NewObjTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjTypeContext {
	var p = new(ObjTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objType

	return p
}

func (s *ObjTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjTypeContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *ObjTypeContext) ListType() IListTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *ObjTypeContext) NonNullType() INonNullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonNullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonNullTypeContext)
}

func (s *ObjTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjType(s)
	}
}

func (s *ObjTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjType(s)
	}
}

func (p *GraphQLParser) ObjType() (localctx IObjTypeContext) {
	localctx = NewObjTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GraphQLParserRULE_objType)

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

	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(247)
			p.TypeName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(248)
			p.ListType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(249)
			p.NonNullType()
		}

	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (p *GraphQLParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GraphQLParserRULE_typeName)

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
		p.SetState(252)
		p.Name()
	}

	return localctx
}

// IListTypeContext is an interface to support dynamic dispatch.
type IListTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListTypeContext differentiates from other interfaces.
	IsListTypeContext()
}

type ListTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListTypeContext() *ListTypeContext {
	var p = new(ListTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_listType
	return p
}

func (*ListTypeContext) IsListTypeContext() {}

func NewListTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeContext {
	var p = new(ListTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_listType

	return p
}

func (s *ListTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeContext) ObjType() IObjTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjTypeContext)
}

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitListType(s)
	}
}

func (p *GraphQLParser) ListType() (localctx IListTypeContext) {
	localctx = NewListTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GraphQLParserRULE_listType)

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
		p.Match(GraphQLParserT__0)
	}
	{
		p.SetState(255)
		p.ObjType()
	}
	{
		p.SetState(256)
		p.Match(GraphQLParserT__1)
	}

	return localctx
}

// INonNullTypeContext is an interface to support dynamic dispatch.
type INonNullTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNonNullTypeContext differentiates from other interfaces.
	IsNonNullTypeContext()
}

type NonNullTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonNullTypeContext() *NonNullTypeContext {
	var p = new(NonNullTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_nonNullType
	return p
}

func (*NonNullTypeContext) IsNonNullTypeContext() {}

func NewNonNullTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonNullTypeContext {
	var p = new(NonNullTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_nonNullType

	return p
}

func (s *NonNullTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NonNullTypeContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *NonNullTypeContext) ListType() IListTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *NonNullTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonNullTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonNullTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterNonNullType(s)
	}
}

func (s *NonNullTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitNonNullType(s)
	}
}

func (p *GraphQLParser) NonNullType() (localctx INonNullTypeContext) {
	localctx = NewNonNullTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GraphQLParserRULE_nonNullType)

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

	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserFRAGMENT, GraphQLParserQUERY, GraphQLParserMUTATION, GraphQLParserSUBSCRIPTION, GraphQLParserSCHEMA, GraphQLParserSCALAR, GraphQLParserTYPE, GraphQLParserINTERFACE, GraphQLParserIMPLEMENTS, GraphQLParserENUM, GraphQLParserUNION, GraphQLParserINPUT, GraphQLParserEXTEND, GraphQLParserDIRECTIVE, GraphQLParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.TypeName()
		}
		{
			p.SetState(259)
			p.Match(GraphQLParserT__10)
		}

	case GraphQLParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.ListType()
		}
		{
			p.SetState(262)
			p.Match(GraphQLParserT__10)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationDefinitionContext is an interface to support dynamic dispatch.
type IOperationDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationDefinitionContext differentiates from other interfaces.
	IsOperationDefinitionContext()
}

type OperationDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationDefinitionContext() *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationDefinition
	return p
}

func (*OperationDefinitionContext) IsOperationDefinitionContext() {}

func NewOperationDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationDefinition

	return p
}

func (s *OperationDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationDefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *OperationDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *OperationDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *OperationDefinitionContext) VariableDefinitions() IVariableDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionsContext)
}

func (s *OperationDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *OperationDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationDefinition(s)
	}
}

func (s *OperationDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationDefinition(s)
	}
}

func (p *GraphQLParser) OperationDefinition() (localctx IOperationDefinitionContext) {
	localctx = NewOperationDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GraphQLParserRULE_operationDefinition)
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

	p.SetState(279)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(266)
			p.SelectionSet()
		}

	case GraphQLParserQUERY, GraphQLParserMUTATION, GraphQLParserSUBSCRIPTION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.OperationType()
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserFRAGMENT)|(1<<GraphQLParserQUERY)|(1<<GraphQLParserMUTATION)|(1<<GraphQLParserSUBSCRIPTION)|(1<<GraphQLParserSCHEMA)|(1<<GraphQLParserSCALAR)|(1<<GraphQLParserTYPE)|(1<<GraphQLParserINTERFACE)|(1<<GraphQLParserIMPLEMENTS)|(1<<GraphQLParserENUM)|(1<<GraphQLParserUNION)|(1<<GraphQLParserINPUT)|(1<<GraphQLParserEXTEND)|(1<<GraphQLParserDIRECTIVE)|(1<<GraphQLParserNAME))) != 0 {
			{
				p.SetState(268)
				p.Name()
			}

		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__6 {
			{
				p.SetState(271)
				p.VariableDefinitions()
			}

		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__5 {
			{
				p.SetState(274)
				p.Directives()
			}

		}
		{
			p.SetState(277)
			p.SelectionSet()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableDefinitionsContext is an interface to support dynamic dispatch.
type IVariableDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDefinitionsContext differentiates from other interfaces.
	IsVariableDefinitionsContext()
}

type VariableDefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionsContext() *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinitions
	return p
}

func (*VariableDefinitionsContext) IsVariableDefinitionsContext() {}

func NewVariableDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variableDefinitions

	return p
}

func (s *VariableDefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionsContext) AllVariableDefinition() []IVariableDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDefinitionContext)(nil)).Elem())
	var tst = make([]IVariableDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDefinitionContext)
		}
	}

	return tst
}

func (s *VariableDefinitionsContext) VariableDefinition(i int) IVariableDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionContext)
}

func (s *VariableDefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariableDefinitions(s)
	}
}

func (s *VariableDefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariableDefinitions(s)
	}
}

func (p *GraphQLParser) VariableDefinitions() (localctx IVariableDefinitionsContext) {
	localctx = NewVariableDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GraphQLParserRULE_variableDefinitions)
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
		p.SetState(281)
		p.Match(GraphQLParserT__6)
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserT__8 {
		{
			p.SetState(282)
			p.VariableDefinition()
		}

		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(287)
		p.Match(GraphQLParserT__7)
	}

	return localctx
}

// IVariableDefinitionContext is an interface to support dynamic dispatch.
type IVariableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDefinitionContext differentiates from other interfaces.
	IsVariableDefinitionContext()
}

type VariableDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionContext() *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinition
	return p
}

func (*VariableDefinitionContext) IsVariableDefinitionContext() {}

func NewVariableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variableDefinition

	return p
}

func (s *VariableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableDefinitionContext) ObjType() IObjTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjTypeContext)
}

func (s *VariableDefinitionContext) DefaultValue() IDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *VariableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariableDefinition(s)
	}
}

func (p *GraphQLParser) VariableDefinition() (localctx IVariableDefinitionContext) {
	localctx = NewVariableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GraphQLParserRULE_variableDefinition)
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
		p.SetState(289)
		p.Variable()
	}
	{
		p.SetState(290)
		p.Match(GraphQLParserT__4)
	}
	{
		p.SetState(291)
		p.ObjType()
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__9 {
		{
			p.SetState(292)
			p.DefaultValue()
		}

	}

	return localctx
}

// ISelectionSetContext is an interface to support dynamic dispatch.
type ISelectionSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionSetContext differentiates from other interfaces.
	IsSelectionSetContext()
}

type SelectionSetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionSetContext() *SelectionSetContext {
	var p = new(SelectionSetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_selectionSet
	return p
}

func (*SelectionSetContext) IsSelectionSetContext() {}

func NewSelectionSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionSetContext {
	var p = new(SelectionSetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_selectionSet

	return p
}

func (s *SelectionSetContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionSetContext) AllSelection() []ISelectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectionContext)(nil)).Elem())
	var tst = make([]ISelectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectionContext)
		}
	}

	return tst
}

func (s *SelectionSetContext) Selection(i int) ISelectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectionContext)
}

func (s *SelectionSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSelectionSet(s)
	}
}

func (s *SelectionSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSelectionSet(s)
	}
}

func (p *GraphQLParser) SelectionSet() (localctx ISelectionSetContext) {
	localctx = NewSelectionSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GraphQLParserRULE_selectionSet)
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
		p.SetState(295)
		p.Match(GraphQLParserT__2)
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__11)|(1<<GraphQLParserFRAGMENT)|(1<<GraphQLParserQUERY)|(1<<GraphQLParserMUTATION)|(1<<GraphQLParserSUBSCRIPTION)|(1<<GraphQLParserSCHEMA)|(1<<GraphQLParserSCALAR)|(1<<GraphQLParserTYPE)|(1<<GraphQLParserINTERFACE)|(1<<GraphQLParserIMPLEMENTS)|(1<<GraphQLParserENUM)|(1<<GraphQLParserUNION)|(1<<GraphQLParserINPUT)|(1<<GraphQLParserEXTEND)|(1<<GraphQLParserDIRECTIVE)|(1<<GraphQLParserNAME))) != 0) {
		{
			p.SetState(296)
			p.Selection()
		}

		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(301)
		p.Match(GraphQLParserT__3)
	}

	return localctx
}

// ISelectionContext is an interface to support dynamic dispatch.
type ISelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionContext differentiates from other interfaces.
	IsSelectionContext()
}

type SelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionContext() *SelectionContext {
	var p = new(SelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_selection
	return p
}

func (*SelectionContext) IsSelectionContext() {}

func NewSelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionContext {
	var p = new(SelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_selection

	return p
}

func (s *SelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *SelectionContext) FragmentSpread() IFragmentSpreadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentSpreadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentSpreadContext)
}

func (s *SelectionContext) InlineFragment() IInlineFragmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineFragmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineFragmentContext)
}

func (s *SelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSelection(s)
	}
}

func (s *SelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSelection(s)
	}
}

func (p *GraphQLParser) Selection() (localctx ISelectionContext) {
	localctx = NewSelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GraphQLParserRULE_selection)

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

	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(303)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(304)
			p.FragmentSpread()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(305)
			p.InlineFragment()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *FieldContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FieldContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FieldContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *GraphQLParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GraphQLParserRULE_field)
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
	p.SetState(309)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(308)
			p.Alias()
		}

	}
	{
		p.SetState(311)
		p.Name()
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__6 {
		{
			p.SetState(312)
			p.Arguments()
		}

	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(315)
			p.Directives()
		}

	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__2 {
		{
			p.SetState(318)
			p.SelectionSet()
		}

	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (p *GraphQLParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GraphQLParserRULE_alias)

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
		p.SetState(321)
		p.Name()
	}
	{
		p.SetState(322)
		p.Match(GraphQLParserT__4)
	}

	return localctx
}

// IFragmentSpreadContext is an interface to support dynamic dispatch.
type IFragmentSpreadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentSpreadContext differentiates from other interfaces.
	IsFragmentSpreadContext()
}

type FragmentSpreadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentSpreadContext() *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentSpread
	return p
}

func (*FragmentSpreadContext) IsFragmentSpreadContext() {}

func NewFragmentSpreadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentSpread

	return p
}

func (s *FragmentSpreadContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentSpreadContext) FragmentName() IFragmentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentSpreadContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentSpreadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentSpreadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentSpreadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentSpread(s)
	}
}

func (s *FragmentSpreadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentSpread(s)
	}
}

func (p *GraphQLParser) FragmentSpread() (localctx IFragmentSpreadContext) {
	localctx = NewFragmentSpreadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GraphQLParserRULE_fragmentSpread)
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
		p.SetState(324)
		p.Match(GraphQLParserT__11)
	}
	{
		p.SetState(325)
		p.FragmentName()
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(326)
			p.Directives()
		}

	}

	return localctx
}

// IInlineFragmentContext is an interface to support dynamic dispatch.
type IInlineFragmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineFragmentContext differentiates from other interfaces.
	IsInlineFragmentContext()
}

type InlineFragmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineFragmentContext() *InlineFragmentContext {
	var p = new(InlineFragmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inlineFragment
	return p
}

func (*InlineFragmentContext) IsInlineFragmentContext() {}

func NewInlineFragmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineFragmentContext {
	var p = new(InlineFragmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inlineFragment

	return p
}

func (s *InlineFragmentContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineFragmentContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *InlineFragmentContext) TypeCondition() ITypeConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *InlineFragmentContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InlineFragmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineFragmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineFragmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInlineFragment(s)
	}
}

func (s *InlineFragmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInlineFragment(s)
	}
}

func (p *GraphQLParser) InlineFragment() (localctx IInlineFragmentContext) {
	localctx = NewInlineFragmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GraphQLParserRULE_inlineFragment)
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
		p.SetState(329)
		p.Match(GraphQLParserT__11)
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__12 {
		{
			p.SetState(330)
			p.TypeCondition()
		}

	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(333)
			p.Directives()
		}

	}
	{
		p.SetState(336)
		p.SelectionSet()
	}

	return localctx
}

// IFragmentDefinitionContext is an interface to support dynamic dispatch.
type IFragmentDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentDefinitionContext differentiates from other interfaces.
	IsFragmentDefinitionContext()
}

type FragmentDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentDefinitionContext() *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentDefinition
	return p
}

func (*FragmentDefinitionContext) IsFragmentDefinitionContext() {}

func NewFragmentDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentDefinition

	return p
}

func (s *FragmentDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentDefinitionContext) FragmentName() IFragmentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentDefinitionContext) TypeCondition() ITypeConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *FragmentDefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FragmentDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentDefinition(s)
	}
}

func (s *FragmentDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentDefinition(s)
	}
}

func (p *GraphQLParser) FragmentDefinition() (localctx IFragmentDefinitionContext) {
	localctx = NewFragmentDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GraphQLParserRULE_fragmentDefinition)
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
		p.SetState(338)
		p.Match(GraphQLParserFRAGMENT)
	}
	{
		p.SetState(339)
		p.FragmentName()
	}
	{
		p.SetState(340)
		p.TypeCondition()
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(341)
			p.Directives()
		}

	}
	{
		p.SetState(344)
		p.SelectionSet()
	}

	return localctx
}

// IFragmentNameContext is an interface to support dynamic dispatch.
type IFragmentNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentNameContext differentiates from other interfaces.
	IsFragmentNameContext()
}

type FragmentNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentNameContext() *FragmentNameContext {
	var p = new(FragmentNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentName
	return p
}

func (*FragmentNameContext) IsFragmentNameContext() {}

func NewFragmentNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentNameContext {
	var p = new(FragmentNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentName

	return p
}

func (s *FragmentNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentNameContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FragmentNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentName(s)
	}
}

func (s *FragmentNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentName(s)
	}
}

func (p *GraphQLParser) FragmentName() (localctx IFragmentNameContext) {
	localctx = NewFragmentNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GraphQLParserRULE_fragmentName)

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
		p.SetState(346)
		p.Name()
	}

	return localctx
}

// ITypeConditionContext is an interface to support dynamic dispatch.
type ITypeConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeConditionContext differentiates from other interfaces.
	IsTypeConditionContext()
}

type TypeConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeConditionContext() *TypeConditionContext {
	var p = new(TypeConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeCondition
	return p
}

func (*TypeConditionContext) IsTypeConditionContext() {}

func NewTypeConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeConditionContext {
	var p = new(TypeConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeCondition

	return p
}

func (s *TypeConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeConditionContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TypeConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeCondition(s)
	}
}

func (s *TypeConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeCondition(s)
	}
}

func (p *GraphQLParser) TypeCondition() (localctx ITypeConditionContext) {
	localctx = NewTypeConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GraphQLParserRULE_typeCondition)

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
		p.SetState(348)
		p.Match(GraphQLParserT__12)
	}
	{
		p.SetState(349)
		p.TypeName()
	}

	return localctx
}

// ITypeSystemDefinitionContext is an interface to support dynamic dispatch.
type ITypeSystemDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSystemDefinitionContext differentiates from other interfaces.
	IsTypeSystemDefinitionContext()
}

type TypeSystemDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemDefinitionContext() *TypeSystemDefinitionContext {
	var p = new(TypeSystemDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemDefinition
	return p
}

func (*TypeSystemDefinitionContext) IsTypeSystemDefinitionContext() {}

func NewTypeSystemDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemDefinitionContext {
	var p = new(TypeSystemDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeSystemDefinition

	return p
}

func (s *TypeSystemDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSystemDefinitionContext) SchemaDefinition() ISchemaDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaDefinitionContext)
}

func (s *TypeSystemDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *TypeSystemDefinitionContext) TypeDefinition() ITypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefinitionContext)
}

func (s *TypeSystemDefinitionContext) TypeExtension() ITypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExtensionContext)
}

func (s *TypeSystemDefinitionContext) DirectiveDefinition() IDirectiveDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveDefinitionContext)
}

func (s *TypeSystemDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeSystemDefinition(s)
	}
}

func (s *TypeSystemDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeSystemDefinition(s)
	}
}

func (p *GraphQLParser) TypeSystemDefinition() (localctx ITypeSystemDefinitionContext) {
	localctx = NewTypeSystemDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GraphQLParserRULE_typeSystemDefinition)

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

	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(352)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(351)
				p.Description()
			}

		}
		{
			p.SetState(354)
			p.SchemaDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(355)
			p.TypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(356)
			p.TypeExtension()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(357)
			p.DirectiveDefinition()
		}

	}

	return localctx
}

// ISchemaDefinitionContext is an interface to support dynamic dispatch.
type ISchemaDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaDefinitionContext differentiates from other interfaces.
	IsSchemaDefinitionContext()
}

type SchemaDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaDefinitionContext() *SchemaDefinitionContext {
	var p = new(SchemaDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_schemaDefinition
	return p
}

func (*SchemaDefinitionContext) IsSchemaDefinitionContext() {}

func NewSchemaDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaDefinitionContext {
	var p = new(SchemaDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_schemaDefinition

	return p
}

func (s *SchemaDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaDefinitionContext) SCHEMA() antlr.TerminalNode {
	return s.GetToken(GraphQLParserSCHEMA, 0)
}

func (s *SchemaDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *SchemaDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *SchemaDefinitionContext) AllOperationTypeDefinition() []IOperationTypeDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperationTypeDefinitionContext)(nil)).Elem())
	var tst = make([]IOperationTypeDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperationTypeDefinitionContext)
		}
	}

	return tst
}

func (s *SchemaDefinitionContext) OperationTypeDefinition(i int) IOperationTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeDefinitionContext)
}

func (s *SchemaDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSchemaDefinition(s)
	}
}

func (s *SchemaDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSchemaDefinition(s)
	}
}

func (p *GraphQLParser) SchemaDefinition() (localctx ISchemaDefinitionContext) {
	localctx = NewSchemaDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GraphQLParserRULE_schemaDefinition)
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
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(360)
			p.Description()
		}

	}
	{
		p.SetState(363)
		p.Match(GraphQLParserSCHEMA)
	}
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(364)
			p.Directives()
		}

	}
	{
		p.SetState(367)
		p.Match(GraphQLParserT__2)
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-17)&-(0x1f+1)) == 0 && ((1<<uint((_la-17)))&((1<<(GraphQLParserQUERY-17))|(1<<(GraphQLParserMUTATION-17))|(1<<(GraphQLParserSUBSCRIPTION-17))|(1<<(GraphQLParserStringValue-17))|(1<<(GraphQLParserTripleQuotedStringValue-17)))) != 0) {
		{
			p.SetState(368)
			p.OperationTypeDefinition()
		}

		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(373)
		p.Match(GraphQLParserT__3)
	}

	return localctx
}

// IOperationTypeDefinitionContext is an interface to support dynamic dispatch.
type IOperationTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationTypeDefinitionContext differentiates from other interfaces.
	IsOperationTypeDefinitionContext()
}

type OperationTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTypeDefinitionContext() *OperationTypeDefinitionContext {
	var p = new(OperationTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationTypeDefinition
	return p
}

func (*OperationTypeDefinitionContext) IsOperationTypeDefinitionContext() {}

func NewOperationTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTypeDefinitionContext {
	var p = new(OperationTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationTypeDefinition

	return p
}

func (s *OperationTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationTypeDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *OperationTypeDefinitionContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *OperationTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *OperationTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationTypeDefinition(s)
	}
}

func (s *OperationTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationTypeDefinition(s)
	}
}

func (p *GraphQLParser) OperationTypeDefinition() (localctx IOperationTypeDefinitionContext) {
	localctx = NewOperationTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GraphQLParserRULE_operationTypeDefinition)
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
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(375)
			p.Description()
		}

	}
	{
		p.SetState(378)
		p.OperationType()
	}
	{
		p.SetState(379)
		p.Match(GraphQLParserT__4)
	}
	{
		p.SetState(380)
		p.TypeName()
	}

	return localctx
}

// ITypeDefinitionContext is an interface to support dynamic dispatch.
type ITypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefinitionContext differentiates from other interfaces.
	IsTypeDefinitionContext()
}

type TypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefinitionContext() *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeDefinition
	return p
}

func (*TypeDefinitionContext) IsTypeDefinitionContext() {}

func NewTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeDefinition

	return p
}

func (s *TypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefinitionContext) ScalarTypeDefinition() IScalarTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarTypeDefinitionContext)
}

func (s *TypeDefinitionContext) ObjectTypeDefinition() IObjectTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectTypeDefinitionContext)
}

func (s *TypeDefinitionContext) InterfaceTypeDefinition() IInterfaceTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeDefinitionContext)
}

func (s *TypeDefinitionContext) UnionTypeDefinition() IUnionTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionTypeDefinitionContext)
}

func (s *TypeDefinitionContext) EnumTypeDefinition() IEnumTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeDefinitionContext)
}

func (s *TypeDefinitionContext) InputObjectTypeDefinition() IInputObjectTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputObjectTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputObjectTypeDefinitionContext)
}

func (s *TypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeDefinition(s)
	}
}

func (s *TypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeDefinition(s)
	}
}

func (p *GraphQLParser) TypeDefinition() (localctx ITypeDefinitionContext) {
	localctx = NewTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GraphQLParserRULE_typeDefinition)

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

	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(382)
			p.ScalarTypeDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(383)
			p.ObjectTypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(384)
			p.InterfaceTypeDefinition()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(385)
			p.UnionTypeDefinition()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(386)
			p.EnumTypeDefinition()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(387)
			p.InputObjectTypeDefinition()
		}

	}

	return localctx
}

// ITypeExtensionContext is an interface to support dynamic dispatch.
type ITypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeExtensionContext differentiates from other interfaces.
	IsTypeExtensionContext()
}

type TypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExtensionContext() *TypeExtensionContext {
	var p = new(TypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeExtension
	return p
}

func (*TypeExtensionContext) IsTypeExtensionContext() {}

func NewTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExtensionContext {
	var p = new(TypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeExtension

	return p
}

func (s *TypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExtensionContext) ObjectTypeExtensionDefinition() IObjectTypeExtensionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectTypeExtensionDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectTypeExtensionDefinitionContext)
}

func (s *TypeExtensionContext) InterfaceTypeExtensionDefinition() IInterfaceTypeExtensionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceTypeExtensionDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeExtensionDefinitionContext)
}

func (s *TypeExtensionContext) UnionTypeExtensionDefinition() IUnionTypeExtensionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionTypeExtensionDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionTypeExtensionDefinitionContext)
}

func (s *TypeExtensionContext) ScalarTypeExtensionDefinition() IScalarTypeExtensionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarTypeExtensionDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarTypeExtensionDefinitionContext)
}

func (s *TypeExtensionContext) EnumTypeExtensionDefinition() IEnumTypeExtensionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeExtensionDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeExtensionDefinitionContext)
}

func (s *TypeExtensionContext) InputObjectTypeExtensionDefinition() IInputObjectTypeExtensionDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputObjectTypeExtensionDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputObjectTypeExtensionDefinitionContext)
}

func (s *TypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeExtension(s)
	}
}

func (s *TypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeExtension(s)
	}
}

func (p *GraphQLParser) TypeExtension() (localctx ITypeExtensionContext) {
	localctx = NewTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GraphQLParserRULE_typeExtension)

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

	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(390)
			p.ObjectTypeExtensionDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(391)
			p.InterfaceTypeExtensionDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(392)
			p.UnionTypeExtensionDefinition()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(393)
			p.ScalarTypeExtensionDefinition()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(394)
			p.EnumTypeExtensionDefinition()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(395)
			p.InputObjectTypeExtensionDefinition()
		}

	}

	return localctx
}

// IScalarTypeDefinitionContext is an interface to support dynamic dispatch.
type IScalarTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarTypeDefinitionContext differentiates from other interfaces.
	IsScalarTypeDefinitionContext()
}

type ScalarTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeDefinitionContext() *ScalarTypeDefinitionContext {
	var p = new(ScalarTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_scalarTypeDefinition
	return p
}

func (*ScalarTypeDefinitionContext) IsScalarTypeDefinitionContext() {}

func NewScalarTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeDefinitionContext {
	var p = new(ScalarTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_scalarTypeDefinition

	return p
}

func (s *ScalarTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeDefinitionContext) SCALAR() antlr.TerminalNode {
	return s.GetToken(GraphQLParserSCALAR, 0)
}

func (s *ScalarTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ScalarTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *ScalarTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ScalarTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterScalarTypeDefinition(s)
	}
}

func (s *ScalarTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitScalarTypeDefinition(s)
	}
}

func (p *GraphQLParser) ScalarTypeDefinition() (localctx IScalarTypeDefinitionContext) {
	localctx = NewScalarTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GraphQLParserRULE_scalarTypeDefinition)
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
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(398)
			p.Description()
		}

	}
	{
		p.SetState(401)
		p.Match(GraphQLParserSCALAR)
	}
	{
		p.SetState(402)
		p.Name()
	}
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(403)
			p.Directives()
		}

	}

	return localctx
}

// IScalarTypeExtensionDefinitionContext is an interface to support dynamic dispatch.
type IScalarTypeExtensionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarTypeExtensionDefinitionContext differentiates from other interfaces.
	IsScalarTypeExtensionDefinitionContext()
}

type ScalarTypeExtensionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeExtensionDefinitionContext() *ScalarTypeExtensionDefinitionContext {
	var p = new(ScalarTypeExtensionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_scalarTypeExtensionDefinition
	return p
}

func (*ScalarTypeExtensionDefinitionContext) IsScalarTypeExtensionDefinitionContext() {}

func NewScalarTypeExtensionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeExtensionDefinitionContext {
	var p = new(ScalarTypeExtensionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_scalarTypeExtensionDefinition

	return p
}

func (s *ScalarTypeExtensionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeExtensionDefinitionContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(GraphQLParserEXTEND, 0)
}

func (s *ScalarTypeExtensionDefinitionContext) SCALAR() antlr.TerminalNode {
	return s.GetToken(GraphQLParserSCALAR, 0)
}

func (s *ScalarTypeExtensionDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ScalarTypeExtensionDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ScalarTypeExtensionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeExtensionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarTypeExtensionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterScalarTypeExtensionDefinition(s)
	}
}

func (s *ScalarTypeExtensionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitScalarTypeExtensionDefinition(s)
	}
}

func (p *GraphQLParser) ScalarTypeExtensionDefinition() (localctx IScalarTypeExtensionDefinitionContext) {
	localctx = NewScalarTypeExtensionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, GraphQLParserRULE_scalarTypeExtensionDefinition)
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
		p.SetState(406)
		p.Match(GraphQLParserEXTEND)
	}
	{
		p.SetState(407)
		p.Match(GraphQLParserSCALAR)
	}
	{
		p.SetState(408)
		p.Name()
	}
	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(409)
			p.Directives()
		}

	}

	return localctx
}

// IObjectTypeDefinitionContext is an interface to support dynamic dispatch.
type IObjectTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectTypeDefinitionContext differentiates from other interfaces.
	IsObjectTypeDefinitionContext()
}

type ObjectTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectTypeDefinitionContext() *ObjectTypeDefinitionContext {
	var p = new(ObjectTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectTypeDefinition
	return p
}

func (*ObjectTypeDefinitionContext) IsObjectTypeDefinitionContext() {}

func NewObjectTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectTypeDefinitionContext {
	var p = new(ObjectTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectTypeDefinition

	return p
}

func (s *ObjectTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectTypeDefinitionContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GraphQLParserTYPE, 0)
}

func (s *ObjectTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *ObjectTypeDefinitionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ObjectTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ObjectTypeDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *ObjectTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectTypeDefinition(s)
	}
}

func (s *ObjectTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectTypeDefinition(s)
	}
}

func (p *GraphQLParser) ObjectTypeDefinition() (localctx IObjectTypeDefinitionContext) {
	localctx = NewObjectTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, GraphQLParserRULE_objectTypeDefinition)
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
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(412)
			p.Description()
		}

	}
	{
		p.SetState(415)
		p.Match(GraphQLParserTYPE)
	}
	{
		p.SetState(416)
		p.Name()
	}
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserIMPLEMENTS {
		{
			p.SetState(417)
			p.implementsInterfaces(0)
		}

	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(420)
			p.Directives()
		}

	}
	p.SetState(424)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(423)
			p.FieldsDefinition()
		}

	}

	return localctx
}

// IObjectTypeExtensionDefinitionContext is an interface to support dynamic dispatch.
type IObjectTypeExtensionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectTypeExtensionDefinitionContext differentiates from other interfaces.
	IsObjectTypeExtensionDefinitionContext()
}

type ObjectTypeExtensionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectTypeExtensionDefinitionContext() *ObjectTypeExtensionDefinitionContext {
	var p = new(ObjectTypeExtensionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectTypeExtensionDefinition
	return p
}

func (*ObjectTypeExtensionDefinitionContext) IsObjectTypeExtensionDefinitionContext() {}

func NewObjectTypeExtensionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectTypeExtensionDefinitionContext {
	var p = new(ObjectTypeExtensionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectTypeExtensionDefinition

	return p
}

func (s *ObjectTypeExtensionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectTypeExtensionDefinitionContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(GraphQLParserEXTEND, 0)
}

func (s *ObjectTypeExtensionDefinitionContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GraphQLParserTYPE, 0)
}

func (s *ObjectTypeExtensionDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectTypeExtensionDefinitionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ObjectTypeExtensionDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ObjectTypeExtensionDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *ObjectTypeExtensionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectTypeExtensionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectTypeExtensionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectTypeExtensionDefinition(s)
	}
}

func (s *ObjectTypeExtensionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectTypeExtensionDefinition(s)
	}
}

func (p *GraphQLParser) ObjectTypeExtensionDefinition() (localctx IObjectTypeExtensionDefinitionContext) {
	localctx = NewObjectTypeExtensionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, GraphQLParserRULE_objectTypeExtensionDefinition)
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
		p.SetState(426)
		p.Match(GraphQLParserEXTEND)
	}
	{
		p.SetState(427)
		p.Match(GraphQLParserTYPE)
	}
	{
		p.SetState(428)
		p.Name()
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserIMPLEMENTS {
		{
			p.SetState(429)
			p.implementsInterfaces(0)
		}

	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(432)
			p.Directives()
		}

	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(435)
			p.FieldsDefinition()
		}

	}

	return localctx
}

// IImplementsInterfacesContext is an interface to support dynamic dispatch.
type IImplementsInterfacesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImplementsInterfacesContext differentiates from other interfaces.
	IsImplementsInterfacesContext()
}

type ImplementsInterfacesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplementsInterfacesContext() *ImplementsInterfacesContext {
	var p = new(ImplementsInterfacesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_implementsInterfaces
	return p
}

func (*ImplementsInterfacesContext) IsImplementsInterfacesContext() {}

func NewImplementsInterfacesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplementsInterfacesContext {
	var p = new(ImplementsInterfacesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_implementsInterfaces

	return p
}

func (s *ImplementsInterfacesContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplementsInterfacesContext) IMPLEMENTS() antlr.TerminalNode {
	return s.GetToken(GraphQLParserIMPLEMENTS, 0)
}

func (s *ImplementsInterfacesContext) AllTypeName() []ITypeNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeNameContext)(nil)).Elem())
	var tst = make([]ITypeNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeNameContext)
		}
	}

	return tst
}

func (s *ImplementsInterfacesContext) TypeName(i int) ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *ImplementsInterfacesContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ImplementsInterfacesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplementsInterfacesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplementsInterfacesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterImplementsInterfaces(s)
	}
}

func (s *ImplementsInterfacesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitImplementsInterfaces(s)
	}
}

func (p *GraphQLParser) ImplementsInterfaces() (localctx IImplementsInterfacesContext) {
	return p.implementsInterfaces(0)
}

func (p *GraphQLParser) implementsInterfaces(_p int) (localctx IImplementsInterfacesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewImplementsInterfacesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IImplementsInterfacesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 92
	p.EnterRecursionRule(localctx, 92, GraphQLParserRULE_implementsInterfaces, _p)
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
	{
		p.SetState(439)
		p.Match(GraphQLParserIMPLEMENTS)
	}
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__13 {
		{
			p.SetState(440)
			p.Match(GraphQLParserT__13)
		}

	}
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(443)
				p.TypeName()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(446)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewImplementsInterfacesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GraphQLParserRULE_implementsInterfaces)
			p.SetState(448)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(449)
				p.Match(GraphQLParserT__13)
			}
			{
				p.SetState(450)
				p.TypeName()
			}

		}
		p.SetState(455)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext())
	}

	return localctx
}

// IFieldsDefinitionContext is an interface to support dynamic dispatch.
type IFieldsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsDefinitionContext differentiates from other interfaces.
	IsFieldsDefinitionContext()
}

type FieldsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsDefinitionContext() *FieldsDefinitionContext {
	var p = new(FieldsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fieldsDefinition
	return p
}

func (*FieldsDefinitionContext) IsFieldsDefinitionContext() {}

func NewFieldsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsDefinitionContext {
	var p = new(FieldsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fieldsDefinition

	return p
}

func (s *FieldsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsDefinitionContext) AllFieldDefinition() []IFieldDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldDefinitionContext)(nil)).Elem())
	var tst = make([]IFieldDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldDefinitionContext)
		}
	}

	return tst
}

func (s *FieldsDefinitionContext) FieldDefinition(i int) IFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldDefinitionContext)
}

func (s *FieldsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFieldsDefinition(s)
	}
}

func (s *FieldsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFieldsDefinition(s)
	}
}

func (p *GraphQLParser) FieldsDefinition() (localctx IFieldsDefinitionContext) {
	localctx = NewFieldsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, GraphQLParserRULE_fieldsDefinition)
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
		p.SetState(456)
		p.Match(GraphQLParserT__2)
	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(GraphQLParserFRAGMENT-16))|(1<<(GraphQLParserQUERY-16))|(1<<(GraphQLParserMUTATION-16))|(1<<(GraphQLParserSUBSCRIPTION-16))|(1<<(GraphQLParserSCHEMA-16))|(1<<(GraphQLParserSCALAR-16))|(1<<(GraphQLParserTYPE-16))|(1<<(GraphQLParserINTERFACE-16))|(1<<(GraphQLParserIMPLEMENTS-16))|(1<<(GraphQLParserENUM-16))|(1<<(GraphQLParserUNION-16))|(1<<(GraphQLParserINPUT-16))|(1<<(GraphQLParserEXTEND-16))|(1<<(GraphQLParserDIRECTIVE-16))|(1<<(GraphQLParserNAME-16))|(1<<(GraphQLParserStringValue-16))|(1<<(GraphQLParserTripleQuotedStringValue-16)))) != 0) {
		{
			p.SetState(457)
			p.FieldDefinition()
		}

		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(462)
		p.Match(GraphQLParserT__3)
	}

	return localctx
}

// IFieldDefinitionContext is an interface to support dynamic dispatch.
type IFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDefinitionContext differentiates from other interfaces.
	IsFieldDefinitionContext()
}

type FieldDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDefinitionContext() *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fieldDefinition
	return p
}

func (*FieldDefinitionContext) IsFieldDefinitionContext() {}

func NewFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fieldDefinition

	return p
}

func (s *FieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldDefinitionContext) ObjType() IObjTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjTypeContext)
}

func (s *FieldDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *FieldDefinitionContext) ArgumentsDefinition() IArgumentsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsDefinitionContext)
}

func (s *FieldDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFieldDefinition(s)
	}
}

func (s *FieldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFieldDefinition(s)
	}
}

func (p *GraphQLParser) FieldDefinition() (localctx IFieldDefinitionContext) {
	localctx = NewFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, GraphQLParserRULE_fieldDefinition)
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
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(464)
			p.Description()
		}

	}
	{
		p.SetState(467)
		p.Name()
	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__6 {
		{
			p.SetState(468)
			p.ArgumentsDefinition()
		}

	}
	{
		p.SetState(471)
		p.Match(GraphQLParserT__4)
	}
	{
		p.SetState(472)
		p.ObjType()
	}
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(473)
			p.Directives()
		}

	}

	return localctx
}

// IArgumentsDefinitionContext is an interface to support dynamic dispatch.
type IArgumentsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsDefinitionContext differentiates from other interfaces.
	IsArgumentsDefinitionContext()
}

type ArgumentsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsDefinitionContext() *ArgumentsDefinitionContext {
	var p = new(ArgumentsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_argumentsDefinition
	return p
}

func (*ArgumentsDefinitionContext) IsArgumentsDefinitionContext() {}

func NewArgumentsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsDefinitionContext {
	var p = new(ArgumentsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_argumentsDefinition

	return p
}

func (s *ArgumentsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsDefinitionContext) AllInputValueDefinition() []IInputValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem())
	var tst = make([]IInputValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInputValueDefinitionContext)
		}
	}

	return tst
}

func (s *ArgumentsDefinitionContext) InputValueDefinition(i int) IInputValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInputValueDefinitionContext)
}

func (s *ArgumentsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArgumentsDefinition(s)
	}
}

func (s *ArgumentsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArgumentsDefinition(s)
	}
}

func (p *GraphQLParser) ArgumentsDefinition() (localctx IArgumentsDefinitionContext) {
	localctx = NewArgumentsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, GraphQLParserRULE_argumentsDefinition)
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
		p.SetState(476)
		p.Match(GraphQLParserT__6)
	}
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(GraphQLParserFRAGMENT-16))|(1<<(GraphQLParserQUERY-16))|(1<<(GraphQLParserMUTATION-16))|(1<<(GraphQLParserSUBSCRIPTION-16))|(1<<(GraphQLParserSCHEMA-16))|(1<<(GraphQLParserSCALAR-16))|(1<<(GraphQLParserTYPE-16))|(1<<(GraphQLParserINTERFACE-16))|(1<<(GraphQLParserIMPLEMENTS-16))|(1<<(GraphQLParserENUM-16))|(1<<(GraphQLParserUNION-16))|(1<<(GraphQLParserINPUT-16))|(1<<(GraphQLParserEXTEND-16))|(1<<(GraphQLParserDIRECTIVE-16))|(1<<(GraphQLParserNAME-16))|(1<<(GraphQLParserStringValue-16))|(1<<(GraphQLParserTripleQuotedStringValue-16)))) != 0) {
		{
			p.SetState(477)
			p.InputValueDefinition()
		}

		p.SetState(480)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(482)
		p.Match(GraphQLParserT__7)
	}

	return localctx
}

// IInputValueDefinitionContext is an interface to support dynamic dispatch.
type IInputValueDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputValueDefinitionContext differentiates from other interfaces.
	IsInputValueDefinitionContext()
}

type InputValueDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputValueDefinitionContext() *InputValueDefinitionContext {
	var p = new(InputValueDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputValueDefinition
	return p
}

func (*InputValueDefinitionContext) IsInputValueDefinitionContext() {}

func NewInputValueDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputValueDefinitionContext {
	var p = new(InputValueDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputValueDefinition

	return p
}

func (s *InputValueDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputValueDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputValueDefinitionContext) ObjType() IObjTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjTypeContext)
}

func (s *InputValueDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InputValueDefinitionContext) DefaultValue() IDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *InputValueDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputValueDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputValueDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputValueDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputValueDefinition(s)
	}
}

func (s *InputValueDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputValueDefinition(s)
	}
}

func (p *GraphQLParser) InputValueDefinition() (localctx IInputValueDefinitionContext) {
	localctx = NewInputValueDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, GraphQLParserRULE_inputValueDefinition)
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
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(484)
			p.Description()
		}

	}
	{
		p.SetState(487)
		p.Name()
	}
	{
		p.SetState(488)
		p.Match(GraphQLParserT__4)
	}
	{
		p.SetState(489)
		p.ObjType()
	}
	p.SetState(491)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__9 {
		{
			p.SetState(490)
			p.DefaultValue()
		}

	}
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(493)
			p.Directives()
		}

	}

	return localctx
}

// IInterfaceTypeDefinitionContext is an interface to support dynamic dispatch.
type IInterfaceTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceTypeDefinitionContext differentiates from other interfaces.
	IsInterfaceTypeDefinitionContext()
}

type InterfaceTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeDefinitionContext() *InterfaceTypeDefinitionContext {
	var p = new(InterfaceTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_interfaceTypeDefinition
	return p
}

func (*InterfaceTypeDefinitionContext) IsInterfaceTypeDefinitionContext() {}

func NewInterfaceTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeDefinitionContext {
	var p = new(InterfaceTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_interfaceTypeDefinition

	return p
}

func (s *InterfaceTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeDefinitionContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(GraphQLParserINTERFACE, 0)
}

func (s *InterfaceTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InterfaceTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InterfaceTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InterfaceTypeDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *InterfaceTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInterfaceTypeDefinition(s)
	}
}

func (s *InterfaceTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInterfaceTypeDefinition(s)
	}
}

func (p *GraphQLParser) InterfaceTypeDefinition() (localctx IInterfaceTypeDefinitionContext) {
	localctx = NewInterfaceTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, GraphQLParserRULE_interfaceTypeDefinition)
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
	p.SetState(497)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(496)
			p.Description()
		}

	}
	{
		p.SetState(499)
		p.Match(GraphQLParserINTERFACE)
	}
	{
		p.SetState(500)
		p.Name()
	}
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(501)
			p.Directives()
		}

	}
	p.SetState(505)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(504)
			p.FieldsDefinition()
		}

	}

	return localctx
}

// IInterfaceTypeExtensionDefinitionContext is an interface to support dynamic dispatch.
type IInterfaceTypeExtensionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceTypeExtensionDefinitionContext differentiates from other interfaces.
	IsInterfaceTypeExtensionDefinitionContext()
}

type InterfaceTypeExtensionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeExtensionDefinitionContext() *InterfaceTypeExtensionDefinitionContext {
	var p = new(InterfaceTypeExtensionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_interfaceTypeExtensionDefinition
	return p
}

func (*InterfaceTypeExtensionDefinitionContext) IsInterfaceTypeExtensionDefinitionContext() {}

func NewInterfaceTypeExtensionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeExtensionDefinitionContext {
	var p = new(InterfaceTypeExtensionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_interfaceTypeExtensionDefinition

	return p
}

func (s *InterfaceTypeExtensionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeExtensionDefinitionContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(GraphQLParserEXTEND, 0)
}

func (s *InterfaceTypeExtensionDefinitionContext) INTERFACE() antlr.TerminalNode {
	return s.GetToken(GraphQLParserINTERFACE, 0)
}

func (s *InterfaceTypeExtensionDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InterfaceTypeExtensionDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InterfaceTypeExtensionDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *InterfaceTypeExtensionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeExtensionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeExtensionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInterfaceTypeExtensionDefinition(s)
	}
}

func (s *InterfaceTypeExtensionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInterfaceTypeExtensionDefinition(s)
	}
}

func (p *GraphQLParser) InterfaceTypeExtensionDefinition() (localctx IInterfaceTypeExtensionDefinitionContext) {
	localctx = NewInterfaceTypeExtensionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, GraphQLParserRULE_interfaceTypeExtensionDefinition)
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
		p.SetState(507)
		p.Match(GraphQLParserEXTEND)
	}
	{
		p.SetState(508)
		p.Match(GraphQLParserINTERFACE)
	}
	{
		p.SetState(509)
		p.Name()
	}
	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(510)
			p.Directives()
		}

	}
	p.SetState(514)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(513)
			p.FieldsDefinition()
		}

	}

	return localctx
}

// IUnionTypeDefinitionContext is an interface to support dynamic dispatch.
type IUnionTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeDefinitionContext differentiates from other interfaces.
	IsUnionTypeDefinitionContext()
}

type UnionTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeDefinitionContext() *UnionTypeDefinitionContext {
	var p = new(UnionTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionTypeDefinition
	return p
}

func (*UnionTypeDefinitionContext) IsUnionTypeDefinitionContext() {}

func NewUnionTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeDefinitionContext {
	var p = new(UnionTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionTypeDefinition

	return p
}

func (s *UnionTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeDefinitionContext) UNION() antlr.TerminalNode {
	return s.GetToken(GraphQLParserUNION, 0)
}

func (s *UnionTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UnionTypeDefinitionContext) UnionMembership() IUnionMembershipContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionMembershipContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionMembershipContext)
}

func (s *UnionTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *UnionTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *UnionTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionTypeDefinition(s)
	}
}

func (s *UnionTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionTypeDefinition(s)
	}
}

func (p *GraphQLParser) UnionTypeDefinition() (localctx IUnionTypeDefinitionContext) {
	localctx = NewUnionTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, GraphQLParserRULE_unionTypeDefinition)
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
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(516)
			p.Description()
		}

	}
	{
		p.SetState(519)
		p.Match(GraphQLParserUNION)
	}
	{
		p.SetState(520)
		p.Name()
	}
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(521)
			p.Directives()
		}

	}
	{
		p.SetState(524)
		p.UnionMembership()
	}

	return localctx
}

// IUnionTypeExtensionDefinitionContext is an interface to support dynamic dispatch.
type IUnionTypeExtensionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeExtensionDefinitionContext differentiates from other interfaces.
	IsUnionTypeExtensionDefinitionContext()
}

type UnionTypeExtensionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeExtensionDefinitionContext() *UnionTypeExtensionDefinitionContext {
	var p = new(UnionTypeExtensionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionTypeExtensionDefinition
	return p
}

func (*UnionTypeExtensionDefinitionContext) IsUnionTypeExtensionDefinitionContext() {}

func NewUnionTypeExtensionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeExtensionDefinitionContext {
	var p = new(UnionTypeExtensionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionTypeExtensionDefinition

	return p
}

func (s *UnionTypeExtensionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeExtensionDefinitionContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(GraphQLParserEXTEND, 0)
}

func (s *UnionTypeExtensionDefinitionContext) UNION() antlr.TerminalNode {
	return s.GetToken(GraphQLParserUNION, 0)
}

func (s *UnionTypeExtensionDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UnionTypeExtensionDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *UnionTypeExtensionDefinitionContext) UnionMembership() IUnionMembershipContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionMembershipContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionMembershipContext)
}

func (s *UnionTypeExtensionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeExtensionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeExtensionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionTypeExtensionDefinition(s)
	}
}

func (s *UnionTypeExtensionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionTypeExtensionDefinition(s)
	}
}

func (p *GraphQLParser) UnionTypeExtensionDefinition() (localctx IUnionTypeExtensionDefinitionContext) {
	localctx = NewUnionTypeExtensionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, GraphQLParserRULE_unionTypeExtensionDefinition)
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
		p.SetState(526)
		p.Match(GraphQLParserEXTEND)
	}
	{
		p.SetState(527)
		p.Match(GraphQLParserUNION)
	}
	{
		p.SetState(528)
		p.Name()
	}
	p.SetState(530)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(529)
			p.Directives()
		}

	}
	p.SetState(533)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__9 {
		{
			p.SetState(532)
			p.UnionMembership()
		}

	}

	return localctx
}

// IUnionMembershipContext is an interface to support dynamic dispatch.
type IUnionMembershipContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionMembershipContext differentiates from other interfaces.
	IsUnionMembershipContext()
}

type UnionMembershipContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionMembershipContext() *UnionMembershipContext {
	var p = new(UnionMembershipContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionMembership
	return p
}

func (*UnionMembershipContext) IsUnionMembershipContext() {}

func NewUnionMembershipContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionMembershipContext {
	var p = new(UnionMembershipContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionMembership

	return p
}

func (s *UnionMembershipContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionMembershipContext) UnionMembers() IUnionMembersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionMembersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionMembersContext)
}

func (s *UnionMembershipContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionMembershipContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionMembershipContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionMembership(s)
	}
}

func (s *UnionMembershipContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionMembership(s)
	}
}

func (p *GraphQLParser) UnionMembership() (localctx IUnionMembershipContext) {
	localctx = NewUnionMembershipContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, GraphQLParserRULE_unionMembership)

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
		p.SetState(535)
		p.Match(GraphQLParserT__9)
	}
	{
		p.SetState(536)
		p.unionMembers(0)
	}

	return localctx
}

// IUnionMembersContext is an interface to support dynamic dispatch.
type IUnionMembersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionMembersContext differentiates from other interfaces.
	IsUnionMembersContext()
}

type UnionMembersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionMembersContext() *UnionMembersContext {
	var p = new(UnionMembersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionMembers
	return p
}

func (*UnionMembersContext) IsUnionMembersContext() {}

func NewUnionMembersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionMembersContext {
	var p = new(UnionMembersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionMembers

	return p
}

func (s *UnionMembersContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionMembersContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *UnionMembersContext) UnionMembers() IUnionMembersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionMembersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionMembersContext)
}

func (s *UnionMembersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionMembersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionMembersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionMembers(s)
	}
}

func (s *UnionMembersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionMembers(s)
	}
}

func (p *GraphQLParser) UnionMembers() (localctx IUnionMembersContext) {
	return p.unionMembers(0)
}

func (p *GraphQLParser) unionMembers(_p int) (localctx IUnionMembersContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewUnionMembersContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IUnionMembersContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 112
	p.EnterRecursionRule(localctx, 112, GraphQLParserRULE_unionMembers, _p)
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
	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__14 {
		{
			p.SetState(539)
			p.Match(GraphQLParserT__14)
		}

	}
	{
		p.SetState(542)
		p.TypeName()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(549)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewUnionMembersContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GraphQLParserRULE_unionMembers)
			p.SetState(544)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(545)
				p.Match(GraphQLParserT__14)
			}
			{
				p.SetState(546)
				p.TypeName()
			}

		}
		p.SetState(551)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 68, p.GetParserRuleContext())
	}

	return localctx
}

// IEnumTypeDefinitionContext is an interface to support dynamic dispatch.
type IEnumTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeDefinitionContext differentiates from other interfaces.
	IsEnumTypeDefinitionContext()
}

type EnumTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeDefinitionContext() *EnumTypeDefinitionContext {
	var p = new(EnumTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumTypeDefinition
	return p
}

func (*EnumTypeDefinitionContext) IsEnumTypeDefinitionContext() {}

func NewEnumTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeDefinitionContext {
	var p = new(EnumTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumTypeDefinition

	return p
}

func (s *EnumTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeDefinitionContext) ENUM() antlr.TerminalNode {
	return s.GetToken(GraphQLParserENUM, 0)
}

func (s *EnumTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumTypeDefinitionContext) EnumValueDefinitions() IEnumValueDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueDefinitionsContext)
}

func (s *EnumTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *EnumTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumTypeDefinition(s)
	}
}

func (s *EnumTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumTypeDefinition(s)
	}
}

func (p *GraphQLParser) EnumTypeDefinition() (localctx IEnumTypeDefinitionContext) {
	localctx = NewEnumTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, GraphQLParserRULE_enumTypeDefinition)
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
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(552)
			p.Description()
		}

	}
	{
		p.SetState(555)
		p.Match(GraphQLParserENUM)
	}
	{
		p.SetState(556)
		p.Name()
	}
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(557)
			p.Directives()
		}

	}
	{
		p.SetState(560)
		p.EnumValueDefinitions()
	}

	return localctx
}

// IEnumTypeExtensionDefinitionContext is an interface to support dynamic dispatch.
type IEnumTypeExtensionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeExtensionDefinitionContext differentiates from other interfaces.
	IsEnumTypeExtensionDefinitionContext()
}

type EnumTypeExtensionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeExtensionDefinitionContext() *EnumTypeExtensionDefinitionContext {
	var p = new(EnumTypeExtensionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumTypeExtensionDefinition
	return p
}

func (*EnumTypeExtensionDefinitionContext) IsEnumTypeExtensionDefinitionContext() {}

func NewEnumTypeExtensionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeExtensionDefinitionContext {
	var p = new(EnumTypeExtensionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumTypeExtensionDefinition

	return p
}

func (s *EnumTypeExtensionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeExtensionDefinitionContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(GraphQLParserEXTEND, 0)
}

func (s *EnumTypeExtensionDefinitionContext) ENUM() antlr.TerminalNode {
	return s.GetToken(GraphQLParserENUM, 0)
}

func (s *EnumTypeExtensionDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumTypeExtensionDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumTypeExtensionDefinitionContext) EnumValueDefinitions() IEnumValueDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueDefinitionsContext)
}

func (s *EnumTypeExtensionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeExtensionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeExtensionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumTypeExtensionDefinition(s)
	}
}

func (s *EnumTypeExtensionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumTypeExtensionDefinition(s)
	}
}

func (p *GraphQLParser) EnumTypeExtensionDefinition() (localctx IEnumTypeExtensionDefinitionContext) {
	localctx = NewEnumTypeExtensionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, GraphQLParserRULE_enumTypeExtensionDefinition)
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
		p.SetState(562)
		p.Match(GraphQLParserEXTEND)
	}
	{
		p.SetState(563)
		p.Match(GraphQLParserENUM)
	}
	{
		p.SetState(564)
		p.Name()
	}
	p.SetState(566)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(565)
			p.Directives()
		}

	}
	p.SetState(569)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(568)
			p.EnumValueDefinitions()
		}

	}

	return localctx
}

// IEnumValueDefinitionsContext is an interface to support dynamic dispatch.
type IEnumValueDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueDefinitionsContext differentiates from other interfaces.
	IsEnumValueDefinitionsContext()
}

type EnumValueDefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueDefinitionsContext() *EnumValueDefinitionsContext {
	var p = new(EnumValueDefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValueDefinitions
	return p
}

func (*EnumValueDefinitionsContext) IsEnumValueDefinitionsContext() {}

func NewEnumValueDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueDefinitionsContext {
	var p = new(EnumValueDefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValueDefinitions

	return p
}

func (s *EnumValueDefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueDefinitionsContext) AllEnumValueDefinition() []IEnumValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumValueDefinitionContext)(nil)).Elem())
	var tst = make([]IEnumValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumValueDefinitionContext)
		}
	}

	return tst
}

func (s *EnumValueDefinitionsContext) EnumValueDefinition(i int) IEnumValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumValueDefinitionContext)
}

func (s *EnumValueDefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueDefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueDefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValueDefinitions(s)
	}
}

func (s *EnumValueDefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValueDefinitions(s)
	}
}

func (p *GraphQLParser) EnumValueDefinitions() (localctx IEnumValueDefinitionsContext) {
	localctx = NewEnumValueDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, GraphQLParserRULE_enumValueDefinitions)
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
		p.SetState(571)
		p.Match(GraphQLParserT__2)
	}
	p.SetState(573)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(GraphQLParserFRAGMENT-16))|(1<<(GraphQLParserQUERY-16))|(1<<(GraphQLParserMUTATION-16))|(1<<(GraphQLParserSUBSCRIPTION-16))|(1<<(GraphQLParserSCHEMA-16))|(1<<(GraphQLParserSCALAR-16))|(1<<(GraphQLParserTYPE-16))|(1<<(GraphQLParserINTERFACE-16))|(1<<(GraphQLParserIMPLEMENTS-16))|(1<<(GraphQLParserENUM-16))|(1<<(GraphQLParserUNION-16))|(1<<(GraphQLParserINPUT-16))|(1<<(GraphQLParserEXTEND-16))|(1<<(GraphQLParserDIRECTIVE-16))|(1<<(GraphQLParserNAME-16))|(1<<(GraphQLParserStringValue-16))|(1<<(GraphQLParserTripleQuotedStringValue-16)))) != 0) {
		{
			p.SetState(572)
			p.EnumValueDefinition()
		}

		p.SetState(575)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(577)
		p.Match(GraphQLParserT__3)
	}

	return localctx
}

// IEnumValueDefinitionContext is an interface to support dynamic dispatch.
type IEnumValueDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueDefinitionContext differentiates from other interfaces.
	IsEnumValueDefinitionContext()
}

type EnumValueDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueDefinitionContext() *EnumValueDefinitionContext {
	var p = new(EnumValueDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValueDefinition
	return p
}

func (*EnumValueDefinitionContext) IsEnumValueDefinitionContext() {}

func NewEnumValueDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueDefinitionContext {
	var p = new(EnumValueDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValueDefinition

	return p
}

func (s *EnumValueDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueDefinitionContext) EnumValue() IEnumValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *EnumValueDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *EnumValueDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumValueDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValueDefinition(s)
	}
}

func (s *EnumValueDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValueDefinition(s)
	}
}

func (p *GraphQLParser) EnumValueDefinition() (localctx IEnumValueDefinitionContext) {
	localctx = NewEnumValueDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, GraphQLParserRULE_enumValueDefinition)
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
	p.SetState(580)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(579)
			p.Description()
		}

	}
	{
		p.SetState(582)
		p.EnumValue()
	}
	p.SetState(584)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(583)
			p.Directives()
		}

	}

	return localctx
}

// IInputObjectTypeDefinitionContext is an interface to support dynamic dispatch.
type IInputObjectTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputObjectTypeDefinitionContext differentiates from other interfaces.
	IsInputObjectTypeDefinitionContext()
}

type InputObjectTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputObjectTypeDefinitionContext() *InputObjectTypeDefinitionContext {
	var p = new(InputObjectTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeDefinition
	return p
}

func (*InputObjectTypeDefinitionContext) IsInputObjectTypeDefinitionContext() {}

func NewInputObjectTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputObjectTypeDefinitionContext {
	var p = new(InputObjectTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeDefinition

	return p
}

func (s *InputObjectTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputObjectTypeDefinitionContext) INPUT() antlr.TerminalNode {
	return s.GetToken(GraphQLParserINPUT, 0)
}

func (s *InputObjectTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputObjectTypeDefinitionContext) InputObjectValueDefinitions() IInputObjectValueDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputObjectValueDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputObjectValueDefinitionsContext)
}

func (s *InputObjectTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InputObjectTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputObjectTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputObjectTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputObjectTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputObjectTypeDefinition(s)
	}
}

func (s *InputObjectTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputObjectTypeDefinition(s)
	}
}

func (p *GraphQLParser) InputObjectTypeDefinition() (localctx IInputObjectTypeDefinitionContext) {
	localctx = NewInputObjectTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, GraphQLParserRULE_inputObjectTypeDefinition)
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
	p.SetState(587)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(586)
			p.Description()
		}

	}
	{
		p.SetState(589)
		p.Match(GraphQLParserINPUT)
	}
	{
		p.SetState(590)
		p.Name()
	}
	p.SetState(592)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(591)
			p.Directives()
		}

	}
	{
		p.SetState(594)
		p.InputObjectValueDefinitions()
	}

	return localctx
}

// IInputObjectTypeExtensionDefinitionContext is an interface to support dynamic dispatch.
type IInputObjectTypeExtensionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputObjectTypeExtensionDefinitionContext differentiates from other interfaces.
	IsInputObjectTypeExtensionDefinitionContext()
}

type InputObjectTypeExtensionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputObjectTypeExtensionDefinitionContext() *InputObjectTypeExtensionDefinitionContext {
	var p = new(InputObjectTypeExtensionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeExtensionDefinition
	return p
}

func (*InputObjectTypeExtensionDefinitionContext) IsInputObjectTypeExtensionDefinitionContext() {}

func NewInputObjectTypeExtensionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputObjectTypeExtensionDefinitionContext {
	var p = new(InputObjectTypeExtensionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeExtensionDefinition

	return p
}

func (s *InputObjectTypeExtensionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputObjectTypeExtensionDefinitionContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(GraphQLParserEXTEND, 0)
}

func (s *InputObjectTypeExtensionDefinitionContext) INPUT() antlr.TerminalNode {
	return s.GetToken(GraphQLParserINPUT, 0)
}

func (s *InputObjectTypeExtensionDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputObjectTypeExtensionDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputObjectTypeExtensionDefinitionContext) InputObjectValueDefinitions() IInputObjectValueDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputObjectValueDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputObjectValueDefinitionsContext)
}

func (s *InputObjectTypeExtensionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputObjectTypeExtensionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputObjectTypeExtensionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputObjectTypeExtensionDefinition(s)
	}
}

func (s *InputObjectTypeExtensionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputObjectTypeExtensionDefinition(s)
	}
}

func (p *GraphQLParser) InputObjectTypeExtensionDefinition() (localctx IInputObjectTypeExtensionDefinitionContext) {
	localctx = NewInputObjectTypeExtensionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, GraphQLParserRULE_inputObjectTypeExtensionDefinition)
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
		p.SetState(596)
		p.Match(GraphQLParserEXTEND)
	}
	{
		p.SetState(597)
		p.Match(GraphQLParserINPUT)
	}
	{
		p.SetState(598)
		p.Name()
	}
	p.SetState(600)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__5 {
		{
			p.SetState(599)
			p.Directives()
		}

	}
	p.SetState(603)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 79, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(602)
			p.InputObjectValueDefinitions()
		}

	}

	return localctx
}

// IInputObjectValueDefinitionsContext is an interface to support dynamic dispatch.
type IInputObjectValueDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputObjectValueDefinitionsContext differentiates from other interfaces.
	IsInputObjectValueDefinitionsContext()
}

type InputObjectValueDefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputObjectValueDefinitionsContext() *InputObjectValueDefinitionsContext {
	var p = new(InputObjectValueDefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputObjectValueDefinitions
	return p
}

func (*InputObjectValueDefinitionsContext) IsInputObjectValueDefinitionsContext() {}

func NewInputObjectValueDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputObjectValueDefinitionsContext {
	var p = new(InputObjectValueDefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputObjectValueDefinitions

	return p
}

func (s *InputObjectValueDefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *InputObjectValueDefinitionsContext) AllInputValueDefinition() []IInputValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem())
	var tst = make([]IInputValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInputValueDefinitionContext)
		}
	}

	return tst
}

func (s *InputObjectValueDefinitionsContext) InputValueDefinition(i int) IInputValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInputValueDefinitionContext)
}

func (s *InputObjectValueDefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputObjectValueDefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputObjectValueDefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputObjectValueDefinitions(s)
	}
}

func (s *InputObjectValueDefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputObjectValueDefinitions(s)
	}
}

func (p *GraphQLParser) InputObjectValueDefinitions() (localctx IInputObjectValueDefinitionsContext) {
	localctx = NewInputObjectValueDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, GraphQLParserRULE_inputObjectValueDefinitions)
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
		p.SetState(605)
		p.Match(GraphQLParserT__2)
	}
	p.SetState(607)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-16)&-(0x1f+1)) == 0 && ((1<<uint((_la-16)))&((1<<(GraphQLParserFRAGMENT-16))|(1<<(GraphQLParserQUERY-16))|(1<<(GraphQLParserMUTATION-16))|(1<<(GraphQLParserSUBSCRIPTION-16))|(1<<(GraphQLParserSCHEMA-16))|(1<<(GraphQLParserSCALAR-16))|(1<<(GraphQLParserTYPE-16))|(1<<(GraphQLParserINTERFACE-16))|(1<<(GraphQLParserIMPLEMENTS-16))|(1<<(GraphQLParserENUM-16))|(1<<(GraphQLParserUNION-16))|(1<<(GraphQLParserINPUT-16))|(1<<(GraphQLParserEXTEND-16))|(1<<(GraphQLParserDIRECTIVE-16))|(1<<(GraphQLParserNAME-16))|(1<<(GraphQLParserStringValue-16))|(1<<(GraphQLParserTripleQuotedStringValue-16)))) != 0) {
		{
			p.SetState(606)
			p.InputValueDefinition()
		}

		p.SetState(609)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(611)
		p.Match(GraphQLParserT__3)
	}

	return localctx
}

// IDirectiveDefinitionContext is an interface to support dynamic dispatch.
type IDirectiveDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveDefinitionContext differentiates from other interfaces.
	IsDirectiveDefinitionContext()
}

type DirectiveDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveDefinitionContext() *DirectiveDefinitionContext {
	var p = new(DirectiveDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveDefinition
	return p
}

func (*DirectiveDefinitionContext) IsDirectiveDefinitionContext() {}

func NewDirectiveDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveDefinitionContext {
	var p = new(DirectiveDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directiveDefinition

	return p
}

func (s *DirectiveDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveDefinitionContext) DIRECTIVE() antlr.TerminalNode {
	return s.GetToken(GraphQLParserDIRECTIVE, 0)
}

func (s *DirectiveDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DirectiveDefinitionContext) DirectiveLocations() IDirectiveLocationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveLocationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveLocationsContext)
}

func (s *DirectiveDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *DirectiveDefinitionContext) ArgumentsDefinition() IArgumentsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsDefinitionContext)
}

func (s *DirectiveDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectiveDefinition(s)
	}
}

func (s *DirectiveDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectiveDefinition(s)
	}
}

func (p *GraphQLParser) DirectiveDefinition() (localctx IDirectiveDefinitionContext) {
	localctx = NewDirectiveDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, GraphQLParserRULE_directiveDefinition)
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
	p.SetState(614)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringValue || _la == GraphQLParserTripleQuotedStringValue {
		{
			p.SetState(613)
			p.Description()
		}

	}
	{
		p.SetState(616)
		p.Match(GraphQLParserDIRECTIVE)
	}
	{
		p.SetState(617)
		p.Match(GraphQLParserT__5)
	}
	{
		p.SetState(618)
		p.Name()
	}
	p.SetState(620)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__6 {
		{
			p.SetState(619)
			p.ArgumentsDefinition()
		}

	}
	{
		p.SetState(622)
		p.Match(GraphQLParserT__12)
	}
	{
		p.SetState(623)
		p.directiveLocations(0)
	}

	return localctx
}

// IDirectiveLocationContext is an interface to support dynamic dispatch.
type IDirectiveLocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveLocationContext differentiates from other interfaces.
	IsDirectiveLocationContext()
}

type DirectiveLocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveLocationContext() *DirectiveLocationContext {
	var p = new(DirectiveLocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveLocation
	return p
}

func (*DirectiveLocationContext) IsDirectiveLocationContext() {}

func NewDirectiveLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveLocationContext {
	var p = new(DirectiveLocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directiveLocation

	return p
}

func (s *DirectiveLocationContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveLocationContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DirectiveLocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveLocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveLocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectiveLocation(s)
	}
}

func (s *DirectiveLocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectiveLocation(s)
	}
}

func (p *GraphQLParser) DirectiveLocation() (localctx IDirectiveLocationContext) {
	localctx = NewDirectiveLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, GraphQLParserRULE_directiveLocation)

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
		p.SetState(625)
		p.Name()
	}

	return localctx
}

// IDirectiveLocationsContext is an interface to support dynamic dispatch.
type IDirectiveLocationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveLocationsContext differentiates from other interfaces.
	IsDirectiveLocationsContext()
}

type DirectiveLocationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveLocationsContext() *DirectiveLocationsContext {
	var p = new(DirectiveLocationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveLocations
	return p
}

func (*DirectiveLocationsContext) IsDirectiveLocationsContext() {}

func NewDirectiveLocationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveLocationsContext {
	var p = new(DirectiveLocationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directiveLocations

	return p
}

func (s *DirectiveLocationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveLocationsContext) DirectiveLocation() IDirectiveLocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveLocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveLocationContext)
}

func (s *DirectiveLocationsContext) DirectiveLocations() IDirectiveLocationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveLocationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveLocationsContext)
}

func (s *DirectiveLocationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveLocationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveLocationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectiveLocations(s)
	}
}

func (s *DirectiveLocationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectiveLocations(s)
	}
}

func (p *GraphQLParser) DirectiveLocations() (localctx IDirectiveLocationsContext) {
	return p.directiveLocations(0)
}

func (p *GraphQLParser) directiveLocations(_p int) (localctx IDirectiveLocationsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDirectiveLocationsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDirectiveLocationsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 132
	p.EnterRecursionRule(localctx, 132, GraphQLParserRULE_directiveLocations, _p)

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
	{
		p.SetState(628)
		p.DirectiveLocation()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDirectiveLocationsContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GraphQLParserRULE_directiveLocations)
			p.SetState(630)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(631)
				p.Match(GraphQLParserT__14)
			}
			{
				p.SetState(632)
				p.DirectiveLocation()
			}

		}
		p.SetState(637)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 83, p.GetParserRuleContext())
	}

	return localctx
}

func (p *GraphQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 46:
		var t *ImplementsInterfacesContext = nil
		if localctx != nil {
			t = localctx.(*ImplementsInterfacesContext)
		}
		return p.ImplementsInterfaces_Sempred(t, predIndex)

	case 56:
		var t *UnionMembersContext = nil
		if localctx != nil {
			t = localctx.(*UnionMembersContext)
		}
		return p.UnionMembers_Sempred(t, predIndex)

	case 66:
		var t *DirectiveLocationsContext = nil
		if localctx != nil {
			t = localctx.(*DirectiveLocationsContext)
		}
		return p.DirectiveLocations_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GraphQLParser) ImplementsInterfaces_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GraphQLParser) UnionMembers_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GraphQLParser) DirectiveLocations_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
