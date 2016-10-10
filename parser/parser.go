//line parser/parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser/parser.go.y:2
//line parser/parser.go.y:5
type yySymType struct {
	yys     int
	integer int
	symbol  string
}

const NUMBER = 57346
const ID = 57347
const STRING = 57348
const EMAIL = 57349
const IPV4 = 57350
const IPV6 = 57351
const IP_CIDR = 57352
const IPADDR_RANGE = 57353
const HEX32 = 57354
const PATHSTR = 57355
const LB = 57356
const RB = 57357
const GLOBALDEFS = 57358
const NOTIFICATION_EMAIL = 57359
const NOTIFICATION_EMAIL_FROM = 57360
const SMTP_SERVER = 57361
const SMTP_CONNECT_TIMEOUT = 57362
const ROUTER_ID = 57363
const LVS_ID = 57364
const VRRP_MCAST_GROUP4 = 57365
const VRRP_MCAST_GROUP6 = 57366
const VRRP_GARP_MASTER_DELAY = 57367
const VRRP_GARP_MASTER_REPEAT = 57368
const VRRP_GARP_MASTER_REFRESH = 57369
const VRRP_GARP_MASTER_REFRESH_REPEAT = 57370
const VRRP_VERSION = 57371
const STATIC_IPADDRESS = 57372
const STATIC_ROUTES = 57373
const STATIC_RULES = 57374
const VRRP_SYNC_GROUP = 57375
const GROUP = 57376
const VRRP_INSTANCE = 57377
const USE_VMAC = 57378
const VERSION = 57379
const VMAC_XMIT_BASE = 57380
const NATIVE_IPV6 = 57381
const INTERFACE = 57382
const MCAST_SRC_IP = 57383
const UNICAST_SRC_IP = 57384
const UNICAST_PEER = 57385
const LVS_SYNC_DAEMON_INTERFACE = 57386
const VIRTUAL_ROUTER_ID = 57387
const NOPREEMPT = 57388
const PRIORITY = 57389
const ADVERT_INT = 57390
const VIRTUAL_IPADDRESS = 57391
const VIRTUAL_IPADDRESS_EXCLUDED = 57392
const VIRTUAL_ROUTES = 57393
const STATE = 57394
const MASTER = 57395
const BACKUP = 57396
const GARP_MASTER_DELAY = 57397
const SMTP_ALERT = 57398
const AUTHENTICATION = 57399
const AUTH_TYPE = 57400
const AUTH_PASS = 57401
const PASS = 57402
const AH = 57403
const LABEL = 57404
const DEV = 57405
const SCOPE = 57406
const SITE = 57407
const LINK = 57408
const HOST = 57409
const NOWHERE = 57410
const GLOBAL = 57411
const BRD = 57412
const SRC = 57413
const FROM = 57414
const TO = 57415
const VIA = 57416
const GW = 57417
const OR = 57418
const TABLE = 57419
const METRIC = 57420
const TRACK_INTERFACE = 57421
const TRACK_SCRIPT = 57422
const DONT_TRACK_PRIMARY = 57423
const NOTIFY_MASTER = 57424
const NOTIFY_BACKUP = 57425
const NOTIFY_FAULT = 57426
const NOTIFY_STOP = 57427
const NOTIFY = 57428
const BLACKHOLE = 57429
const VRRP_SCRIPT = 57430
const SCRIPT = 57431
const INTERVAL = 57432
const TIMEOUT = 57433
const WEIGHT = 57434
const FALL = 57435
const RISE = 57436
const VIRTUAL_SERVER_GROUP = 57437
const VIRTUAL_SERVER = 57438
const DELAY_LOOP = 57439
const LB_ALGO = 57440
const LB_KIND = 57441
const LVS_SCHED = 57442
const LVS_METHOD = 57443
const RR = 57444
const WRR = 57445
const LC = 57446
const WLC = 57447
const LBLC = 57448
const SH = 57449
const DH = 57450
const NAT = 57451
const DR = 57452
const TUN = 57453
const PERSISTENCE_TIMEOUT = 57454
const PROTOCOL = 57455
const TCP = 57456
const UDP = 57457
const SORRY_SERVER = 57458
const REAL_SERVER = 57459
const FWMARK = 57460
const INHIBIT_ON_FAILURE = 57461
const TCP_CHECK = 57462
const HTTP_GET = 57463
const SSL_GET = 57464
const SMTP_CHECK = 57465
const MISC_CHECK = 57466
const URL = 57467
const PATH = 57468
const DIGEST = 57469
const STATUS_CODE = 57470
const CONNECT_TIMEOUT = 57471
const CONNECT_PORT = 57472
const CONNECT_IP = 57473
const BINDTO = 57474
const BIND_PORT = 57475
const RETRY = 57476
const HELO_NAME = 57477
const MISC_PATH = 57478
const MISC_TIMEOUT = 57479
const WARMUP = 57480
const MISC_DYNAMIC = 57481
const NB_GET_RETRY = 57482
const DELAY_BEFORE_RETRY = 57483
const VIRTUALHOST = 57484
const ALPHA = 57485
const OMEGA = 57486
const QUORUM = 57487
const HYSTERESIS = 57488
const QUORUM_UP = 57489
const QUORUM_DOWN = 57490

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"NUMBER",
	"ID",
	"STRING",
	"EMAIL",
	"IPV4",
	"IPV6",
	"IP_CIDR",
	"IPADDR_RANGE",
	"HEX32",
	"PATHSTR",
	"LB",
	"RB",
	"GLOBALDEFS",
	"NOTIFICATION_EMAIL",
	"NOTIFICATION_EMAIL_FROM",
	"SMTP_SERVER",
	"SMTP_CONNECT_TIMEOUT",
	"ROUTER_ID",
	"LVS_ID",
	"VRRP_MCAST_GROUP4",
	"VRRP_MCAST_GROUP6",
	"VRRP_GARP_MASTER_DELAY",
	"VRRP_GARP_MASTER_REPEAT",
	"VRRP_GARP_MASTER_REFRESH",
	"VRRP_GARP_MASTER_REFRESH_REPEAT",
	"VRRP_VERSION",
	"STATIC_IPADDRESS",
	"STATIC_ROUTES",
	"STATIC_RULES",
	"VRRP_SYNC_GROUP",
	"GROUP",
	"VRRP_INSTANCE",
	"USE_VMAC",
	"VERSION",
	"VMAC_XMIT_BASE",
	"NATIVE_IPV6",
	"INTERFACE",
	"MCAST_SRC_IP",
	"UNICAST_SRC_IP",
	"UNICAST_PEER",
	"LVS_SYNC_DAEMON_INTERFACE",
	"VIRTUAL_ROUTER_ID",
	"NOPREEMPT",
	"PRIORITY",
	"ADVERT_INT",
	"VIRTUAL_IPADDRESS",
	"VIRTUAL_IPADDRESS_EXCLUDED",
	"VIRTUAL_ROUTES",
	"STATE",
	"MASTER",
	"BACKUP",
	"GARP_MASTER_DELAY",
	"SMTP_ALERT",
	"AUTHENTICATION",
	"AUTH_TYPE",
	"AUTH_PASS",
	"PASS",
	"AH",
	"LABEL",
	"DEV",
	"SCOPE",
	"SITE",
	"LINK",
	"HOST",
	"NOWHERE",
	"GLOBAL",
	"BRD",
	"SRC",
	"FROM",
	"TO",
	"VIA",
	"GW",
	"OR",
	"TABLE",
	"METRIC",
	"TRACK_INTERFACE",
	"TRACK_SCRIPT",
	"DONT_TRACK_PRIMARY",
	"NOTIFY_MASTER",
	"NOTIFY_BACKUP",
	"NOTIFY_FAULT",
	"NOTIFY_STOP",
	"NOTIFY",
	"BLACKHOLE",
	"VRRP_SCRIPT",
	"SCRIPT",
	"INTERVAL",
	"TIMEOUT",
	"WEIGHT",
	"FALL",
	"RISE",
	"VIRTUAL_SERVER_GROUP",
	"VIRTUAL_SERVER",
	"DELAY_LOOP",
	"LB_ALGO",
	"LB_KIND",
	"LVS_SCHED",
	"LVS_METHOD",
	"RR",
	"WRR",
	"LC",
	"WLC",
	"LBLC",
	"SH",
	"DH",
	"NAT",
	"DR",
	"TUN",
	"PERSISTENCE_TIMEOUT",
	"PROTOCOL",
	"TCP",
	"UDP",
	"SORRY_SERVER",
	"REAL_SERVER",
	"FWMARK",
	"INHIBIT_ON_FAILURE",
	"TCP_CHECK",
	"HTTP_GET",
	"SSL_GET",
	"SMTP_CHECK",
	"MISC_CHECK",
	"URL",
	"PATH",
	"DIGEST",
	"STATUS_CODE",
	"CONNECT_TIMEOUT",
	"CONNECT_PORT",
	"CONNECT_IP",
	"BINDTO",
	"BIND_PORT",
	"RETRY",
	"HELO_NAME",
	"MISC_PATH",
	"MISC_TIMEOUT",
	"WARMUP",
	"MISC_DYNAMIC",
	"NB_GET_RETRY",
	"DELAY_BEFORE_RETRY",
	"VIRTUALHOST",
	"ALPHA",
	"OMEGA",
	"QUORUM",
	"HYSTERESIS",
	"QUORUM_UP",
	"QUORUM_DOWN",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.go.y:360

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 19,
	14, 125,
	-2, 277,
	-1, 200,
	15, 117,
	-2, 118,
}

const yyNprod = 290
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 525

var yyAct = [...]int{

	398, 456, 451, 63, 410, 396, 387, 380, 361, 326,
	322, 206, 318, 305, 311, 303, 30, 180, 301, 181,
	204, 199, 298, 33, 270, 182, 262, 54, 172, 142,
	389, 135, 59, 357, 390, 391, 412, 413, 414, 415,
	34, 35, 64, 394, 216, 392, 393, 383, 382, 95,
	60, 215, 384, 453, 454, 455, 385, 54, 214, 386,
	105, 34, 35, 213, 202, 115, 116, 278, 279, 118,
	119, 120, 399, 271, 272, 273, 74, 127, 34, 35,
	131, 133, 107, 324, 363, 434, 60, 409, 464, 201,
	76, 77, 211, 114, 69, 68, 72, 458, 408, 405,
	404, 406, 407, 61, 32, 62, 65, 66, 67, 70,
	71, 364, 365, 366, 367, 368, 369, 349, 350, 73,
	320, 321, 445, 409, 108, 109, 110, 111, 112, 183,
	184, 185, 186, 187, 408, 405, 404, 406, 407, 401,
	403, 432, 137, 400, 188, 189, 402, 52, 190, 191,
	238, 239, 130, 227, 228, 126, 263, 264, 265, 266,
	267, 268, 269, 431, 37, 423, 416, 379, 218, 325,
	138, 203, 371, 224, 192, 193, 194, 195, 196, 197,
	198, 174, 175, 176, 177, 178, 179, 355, 31, 352,
	347, 345, 341, 336, 280, 281, 34, 35, 55, 259,
	334, 104, 251, 87, 201, 34, 35, 64, 60, 211,
	332, 292, 275, 274, 287, 260, 258, 250, 139, 140,
	141, 223, 288, 242, 217, 129, 293, 34, 35, 310,
	12, 243, 425, 302, 89, 34, 35, 310, 113, 309,
	309, 103, 88, 378, 13, 14, 15, 16, 86, 17,
	313, 57, 58, 377, 376, 460, 447, 317, 56, 69,
	68, 72, 461, 448, 100, 375, 374, 359, 61, 358,
	62, 65, 66, 67, 70, 71, 244, 241, 237, 236,
	235, 306, 307, 229, 73, 33, 219, 85, 81, 308,
	314, 80, 79, 78, 25, 24, 23, 315, 330, 22,
	462, 101, 18, 34, 35, 134, 302, 90, 309, 20,
	19, 463, 340, 34, 35, 132, 309, 34, 35, 344,
	337, 335, 333, 438, 328, 211, 313, 342, 34, 35,
	128, 343, 348, 351, 353, 346, 339, 356, 39, 40,
	45, 46, 47, 48, 49, 50, 41, 42, 43, 44,
	51, 207, 338, 208, 209, 34, 35, 212, 450, 210,
	34, 35, 117, 300, 96, 331, 34, 35, 370, 34,
	35, 372, 232, 123, 231, 124, 286, 285, 282, 252,
	249, 248, 247, 395, 246, 245, 230, 449, 226, 417,
	222, 221, 220, 122, 121, 424, 106, 99, 98, 83,
	36, 28, 27, 433, 26, 444, 443, 442, 439, 440,
	441, 437, 436, 435, 430, 429, 446, 144, 145, 146,
	147, 148, 149, 150, 151, 152, 153, 154, 155, 156,
	157, 158, 159, 160, 428, 457, 161, 162, 163, 427,
	426, 422, 421, 420, 419, 418, 373, 360, 354, 329,
	297, 296, 295, 294, 291, 459, 290, 289, 457, 465,
	164, 165, 166, 167, 168, 169, 170, 171, 284, 283,
	276, 261, 257, 256, 255, 254, 253, 240, 234, 233,
	225, 125, 102, 97, 94, 93, 92, 91, 84, 82,
	1, 53, 205, 21, 312, 304, 411, 397, 452, 388,
	381, 362, 277, 29, 200, 173, 327, 323, 319, 316,
	143, 299, 136, 75, 38, 11, 10, 9, 8, 7,
	6, 5, 4, 3, 2,
}
var yyPact = [...]int{

	214, -1000, 214, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 285, 282, 281, 280, 398, 396, 395, 70,
	394, -1000, 321, 188, 197, 18, 279, 278, 277, 274,
	-1000, 485, 393, 484, -1000, -1000, 273, 233, 321, 228,
	300, 483, 482, 481, 480, 358, 479, 392, 391, 256,
	292, 478, 226, 188, -1000, -1000, 309, 390, 59, 223,
	197, 309, 352, -1000, -1000, 309, 309, 309, 388, 387,
	369, 477, 59, 320, 210, 18, 305, 295, 136, 381,
	92, 32, -1000, -1000, -1000, 53, -1000, -1000, 347, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -14, -19, -26, -33, 209, 136, -1000, 272, 386,
	385, 384, 206, 381, -1000, 476, -1000, -1000, 382, 309,
	309, 269, 380, 368, -1000, 475, 474, 266, 265, 264,
	97, 473, -1000, 263, 217, 262, -1000, 379, 378, 376,
	375, 374, 202, 92, 373, 472, 471, 470, 469, 468,
	201, 32, 197, 467, 54, -36, 54, -36, 466, -47,
	309, 309, 372, -1000, -1000, 465, 464, 371, 370, 199,
	53, 453, 452, 450, 196, 347, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 449, 448, 447, 446, -1000, -1000, 357,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 309,
	-1000, -1000, -1000, -1000, -1000, 219, 227, 32, -1000, -1000,
	-1000, 62, -1000, 77, 318, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	445, 361, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 195, 357,
	-1000, 185, 309, 178, 219, -1000, 346, 330, 309, -1000,
	-1000, 177, 227, -1000, 325, 309, 176, 32, 175, 62,
	57, 347, 174, 77, -1000, 444, 172, 318, -59, -1000,
	255, 253, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 443, -8, -8,
	-1000, 157, -8, 442, -1000, 252, 251, 240, 239, 229,
	152, -1000, -1000, -1000, -82, -95, -95, 5, -100, -1000,
	151, -82, 441, 440, 439, 438, 437, 150, -95, 218,
	436, 435, 430, 411, 410, 148, 126, 5, -1000, 71,
	409, 408, 407, 317, 309, 405, 309, 403, 402, 401,
	107, -100, 250, 383, 354, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -73, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -31, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 82, -73, 249, 288, 307, 73, -31, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 490, 524, 523, 522, 521, 520, 519, 518, 517,
	516, 515, 164, 514, 20, 3, 147, 32, 76, 513,
	31, 512, 22, 511, 29, 510, 18, 15, 14, 509,
	12, 10, 9, 508, 11, 507, 506, 28, 505, 19,
	17, 25, 21, 504, 503, 16, 26, 24, 502, 8,
	501, 7, 6, 5, 4, 500, 499, 2, 498, 497,
	0, 1, 496, 495, 494, 13, 492, 491, 82,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 12, 12, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 4, 5, 6, 18, 18, 7, 20, 20,
	21, 21, 21, 21, 21, 21, 22, 22, 23, 8,
	24, 24, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 30, 30, 33, 33, 33, 33,
	31, 31, 35, 35, 26, 26, 32, 32, 36, 36,
	36, 9, 37, 37, 38, 38, 38, 38, 38, 38,
	38, 29, 29, 39, 39, 11, 42, 42, 43, 43,
	43, 43, 10, 40, 40, 44, 44, 44, 44, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 49, 49, 50,
	50, 50, 50, 50, 50, 50, 50, 51, 51, 55,
	55, 55, 55, 55, 55, 52, 52, 56, 56, 56,
	56, 56, 56, 56, 57, 57, 58, 58, 58, 58,
	58, 53, 53, 59, 59, 59, 59, 59, 59, 59,
	61, 61, 60, 60, 60, 60, 60, 60, 60, 54,
	54, 62, 62, 62, 62, 62, 62, 46, 46, 46,
	46, 46, 46, 46, 46, 47, 47, 47, 47, 48,
	48, 48, 27, 27, 28, 28, 63, 63, 63, 63,
	64, 64, 64, 14, 14, 66, 16, 16, 67, 67,
	67, 67, 67, 67, 17, 17, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 19, 19, 19, 19, 19, 68, 68,
	68, 68, 68, 68, 65, 65, 65, 15, 15, 15,
	45, 45, 45, 34, 34, 34, 34, 34, 34, 34,
}
var yyR2 = [...]int{

	0, 2, 1, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 4, 2, 1, 0, 4, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 4, 4, 4, 2, 1, 5, 2, 1,
	0, 1, 4, 2, 2, 2, 2, 1, 1, 5,
	2, 1, 0, 1, 2, 1, 1, 2, 2, 2,
	4, 2, 2, 2, 1, 2, 2, 4, 4, 4,
	2, 2, 2, 1, 4, 2, 4, 4, 1, 2,
	2, 2, 2, 2, 2, 1, 0, 2, 2, 2,
	2, 1, 1, 2, 2, 1, 2, 1, 0, 1,
	3, 5, 2, 1, 0, 2, 2, 2, 2, 2,
	2, 2, 1, 2, 1, 5, 2, 1, 0, 2,
	2, 2, 5, 2, 1, 0, 1, 2, 2, 0,
	2, 2, 2, 2, 2, 2, 2, 3, 6, 6,
	2, 1, 1, 2, 2, 2, 2, 2, 1, 0,
	2, 1, 4, 4, 4, 4, 4, 2, 1, 0,
	2, 2, 2, 2, 2, 2, 1, 0, 4, 2,
	2, 2, 2, 2, 2, 1, 0, 2, 2, 2,
	2, 2, 1, 0, 1, 4, 2, 2, 2, 2,
	2, 1, 0, 2, 2, 2, 2, 2, 2, 2,
	1, 0, 2, 2, 2, 2, 1, 0, 1, 1,
	1, 1, 1, 1, 1, 0, 1, 1, 1, 0,
	1, 1, 2, 1, 2, 1, 1, 2, 2, 2,
	1, 2, 2, 2, 1, 1, 2, 1, 0, 1,
	1, 2, 2, 2, 2, 1, 0, 2, 2, 2,
	1, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 0, 4, 4, 4, 4, 0, 1,
	1, 1, 1, 1, 0, 1, 1, 0, 1, 1,
	0, 1, 2, 0, 1, 1, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, 16, 30, 31, 32, 33, 35, 88, 96,
	95, -1, 14, 14, 14, 14, 6, 6, 6, -44,
	-45, 118, 34, -15, 8, 9, 6, -12, -13, 17,
	18, 25, 26, 27, 28, 19, 20, 21, 22, 23,
	24, 29, -16, -67, -15, 10, 70, 63, 64, -17,
	-41, 71, 73, -15, 10, 74, 75, 76, 63, 62,
	77, 78, 64, 87, -18, -19, 72, 73, 14, 14,
	14, 14, 4, 6, 4, 14, 15, -12, 14, 6,
	7, 4, 4, 4, 4, -15, 6, 4, 6, 6,
	8, 9, 4, 15, -16, -15, 6, -68, 65, 66,
	67, 68, 69, 15, -17, -15, -15, 10, -15, -15,
	-15, 6, 6, 4, 6, 4, -68, -15, 10, 15,
	-18, -15, 10, -15, 10, -20, -21, 6, 34, 82,
	83, 84, -24, -25, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 55, 56, 57, 79, 80, 81, 82, 83, 84,
	85, 86, -37, -38, 89, 90, 91, 92, 93, 94,
	-40, -39, -41, 97, 98, 99, 100, 101, 112, 113,
	116, 117, 142, 143, 144, 145, 146, 147, 148, -42,
	-43, -15, 11, 118, -14, -66, -34, 4, 6, 7,
	12, -15, 10, 77, 77, 77, 77, 15, -20, 14,
	6, 6, 6, 15, -24, 4, 6, -15, -15, 14,
	6, 6, 4, 4, 4, 14, 14, 14, 53, 54,
	4, 14, 6, 14, 14, 6, 6, 6, 6, 6,
	15, -37, 6, 4, 4, 4, 4, 4, 15, -40,
	-17, 4, -46, 102, 103, 104, 105, 106, 107, 108,
	-47, 109, 110, 111, -46, -47, 4, -48, 114, 115,
	-15, -15, 6, 4, 4, 6, 6, 15, -42, 4,
	4, 4, 15, -14, 4, 4, 4, 4, -22, -23,
	6, -26, -15, -27, -63, -65, 62, 63, 70, -15,
	10, -28, -64, -65, 63, 70, -29, -39, -30, -33,
	58, 59, -31, -35, 6, 92, -32, -36, 6, 4,
	-45, 4, 15, -22, 15, -26, 15, -27, 6, 6,
	-15, 15, -28, 6, -15, 15, -40, 15, -30, 60,
	61, -34, 15, -31, 4, 15, -32, 92, 14, 14,
	4, -49, -50, 92, 119, 120, 121, 122, 123, 124,
	-49, 15, -49, 4, 14, 14, 14, 14, 14, 15,
	-51, -55, 130, 129, 134, 138, 141, -52, -56, 125,
	129, 130, 140, 141, 138, -52, -53, -59, -60, 67,
	138, 134, 141, 135, 131, 130, 132, 133, 129, 118,
	-54, -62, 136, 137, 138, 139, 15, -51, 4, 4,
	4, 4, 4, 15, -52, 14, 4, 4, 4, 4,
	4, 15, 15, -53, 14, 4, 4, 4, 6, -15,
	4, -15, 4, 4, 4, 15, -54, 6, 13, 4,
	4, -57, -58, 126, 127, 128, -61, -60, 15, -57,
	6, 13, 12, 4, 15, -61,
}
var yyDef = [...]int{

	3, -2, 2, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 0, 0, 0, 0, 0, 0, 0, -2,
	0, 1, 16, 238, 246, 263, 0, 0, 0, 0,
	126, 0, 0, 281, 278, 279, 0, 0, 15, 0,
	0, 0, 0, 0, 0, 277, 0, 0, 0, 0,
	0, 0, 0, 237, 239, 240, 277, 0, 268, 0,
	245, 277, 277, 250, 251, 277, 277, 277, 0, 0,
	0, 0, 268, 277, 0, 36, 277, 277, 40, 52,
	104, 129, 127, 128, 282, 118, 13, 14, 277, 18,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 32, 236, 241, 242, 243, 269, 270,
	271, 272, 273, 33, 244, 247, 248, 249, 252, 253,
	254, 255, 256, 257, 258, 259, 260, 261, 262, 34,
	35, 0, 0, 0, 0, 0, 39, 41, 0, 0,
	0, 0, 0, 51, 53, 0, 55, 56, 0, 277,
	277, 0, 0, 0, 64, 0, 0, 0, 0, 0,
	0, 0, 73, 0, 0, 0, 78, 0, 0, 0,
	0, 0, 0, 103, 0, 0, 0, 0, 0, 0,
	0, 124, 114, 0, 207, 215, 207, 215, 0, 219,
	277, 277, 0, 141, 142, 0, 0, 0, 0, 0,
	-2, 0, 0, 0, 0, 234, 235, 284, 285, 286,
	287, 288, 289, 0, 0, 0, 0, 37, 38, 0,
	43, 44, 45, 49, 50, 54, 57, 58, 59, 277,
	61, 62, 63, 65, 66, 274, 274, 129, 70, 71,
	72, 86, 75, 0, 98, 79, 80, 81, 82, 83,
	101, 102, 105, 106, 107, 108, 109, 110, 122, 123,
	113, 130, 131, 208, 209, 210, 211, 212, 213, 214,
	132, 216, 217, 218, 133, 134, 135, 136, 220, 221,
	0, 277, 140, 143, 144, 145, 146, 115, 116, 119,
	120, 121, 17, 233, 264, 266, 265, 267, 0, 47,
	48, 0, 95, 0, 223, 226, 0, 0, 277, 275,
	276, 0, 225, 230, 0, 277, 0, 112, 0, 85,
	0, 277, 0, 91, 92, 0, 0, 97, 99, 137,
	0, 0, 42, 46, 60, 94, 67, 222, 227, 228,
	229, 68, 224, 231, 232, 69, 111, 74, 84, 87,
	88, 89, 76, 90, 93, 77, 96, 0, 149, 149,
	100, 0, 148, 0, 151, 0, 0, 0, 0, 0,
	0, 138, 147, 150, 159, 167, 167, 183, 201, 139,
	0, 158, 0, 0, 0, 0, 0, 0, 166, 0,
	0, 0, 0, 0, 0, 0, 0, 182, 184, 0,
	0, 0, 0, 0, 277, 0, 277, 0, 0, 0,
	0, 200, 0, 0, 0, 206, 152, 157, 160, 161,
	162, 163, 164, 153, 165, 176, 169, 170, 171, 172,
	173, 154, 155, 181, 192, 186, 187, 188, 189, 193,
	194, 195, 196, 197, 198, 156, 199, 202, 203, 204,
	205, 0, 175, 0, 0, 0, 0, 191, 168, 174,
	177, 178, 179, 180, 185, 190,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
	82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121,
	122, 123, 124, 125, 126, 127, 128, 129, 130, 131,
	132, 133, 134, 135, 136, 137, 138, 139, 140, 141,
	142, 143, 144, 145, 146, 147, 148,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:29
		{
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:31
		{
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:32
		{
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:33
		{
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:34
		{
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:35
		{
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:36
		{
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:37
		{
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:38
		{
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:39
		{
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:40
		{
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:47
		{
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:48
		{
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:49
		{
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:50
		{
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:51
		{
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:52
		{
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:53
		{
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:54
		{
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:55
		{
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:56
		{
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:57
		{
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:58
		{
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:59
		{
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:60
		{
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:61
		{
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:75
		{
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:76
		{
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:77
		{
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:78
		{
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:79
		{
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:80
		{
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:82
		{
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:90
		{
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:91
		{
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:92
		{
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:93
		{
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:94
		{
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:95
		{
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:96
		{
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:97
		{
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:98
		{
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:99
		{
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:100
		{
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:101
		{
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:103
		{
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:104
		{
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:108
		{
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:109
		{
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:110
		{
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:111
		{
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:112
		{
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:113
		{
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:114
		{
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:115
		{
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:116
		{
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:117
		{
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:118
		{
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:119
		{
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:120
		{
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:121
		{
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:125
		{
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:126
		{
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:127
		{
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:128
		{
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:133
		{
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:139
		{
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:140
		{
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:141
		{
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:147
		{
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:148
		{
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:149
		{
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:150
		{
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:151
		{
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:152
		{
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:153
		{
		}
	case 118:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:163
		{
		}
	case 119:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:164
		{
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:165
		{
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:166
		{
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:174
		{
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:175
		{
		}
	case 129:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:177
		{
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:178
		{
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:179
		{
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:180
		{
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:181
		{
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:182
		{
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:183
		{
		}
	case 136:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:184
		{
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:185
		{
		}
	case 140:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:188
		{
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:191
		{
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:192
		{
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:193
		{
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:194
		{
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:196
		{
		}
	case 149:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:198
		{
		}
	case 150:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:199
		{
		}
	case 151:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:200
		{
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:201
		{
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:202
		{
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:203
		{
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:204
		{
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:205
		{
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:207
		{
		}
	case 159:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:209
		{
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:210
		{
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:211
		{
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:212
		{
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:213
		{
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:214
		{
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:216
		{
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:218
		{
		}
	case 168:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:219
		{
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:220
		{
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:221
		{
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:222
		{
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:223
		{
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:224
		{
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:226
		{
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:228
		{
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:229
		{
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:230
		{
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:231
		{
		}
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:232
		{
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:234
		{
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:236
		{
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:237
		{
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:244
		{
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:246
		{
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:247
		{
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:248
		{
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:249
		{
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:250
		{
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:251
		{
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:252
		{
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:254
		{
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:256
		{
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:257
		{
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:258
		{
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:259
		{
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:260
		{
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:261
		{
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:263
		{
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:264
		{
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:265
		{
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:266
		{
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:267
		{
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:268
		{
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:269
		{
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:270
		{
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:272
		{
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:273
		{
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:274
		{
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:275
		{
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:277
		{
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:278
		{
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:279
		{
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:281
		{
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:283
		{
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:285
		{
		}
	case 227:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:286
		{
		}
	case 228:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:287
		{
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:288
		{
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:290
		{
		}
	case 231:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:291
		{
		}
	case 232:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:292
		{
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:294
		{
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:296
		{
		}
	case 238:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:300
		{
		}
	case 246:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:309
		{
		}
	case 263:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:327
		{
		}
	case 268:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:333
		{
		}
	case 274:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:340
		{
		}
	case 277:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:344
		{
		}
	case 280:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:348
		{
		}
	case 283:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:352
		{
		}
	}
	goto yystack /* stack new state and value */
}
