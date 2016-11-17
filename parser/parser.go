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

//line parser/parser.go.y:361

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 19,
	14, 126,
	-2, 278,
	-1, 201,
	15, 118,
	-2, 119,
}

const yyNprod = 291
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 527

var yyAct = [...]int{

	400, 458, 453, 63, 412, 398, 389, 382, 363, 328,
	324, 207, 320, 307, 313, 305, 30, 181, 303, 182,
	205, 200, 300, 33, 272, 183, 264, 54, 173, 135,
	391, 326, 59, 359, 392, 393, 143, 455, 456, 457,
	34, 35, 64, 396, 217, 394, 395, 385, 384, 95,
	60, 216, 386, 34, 35, 215, 387, 54, 214, 388,
	105, 414, 415, 416, 417, 115, 116, 280, 281, 118,
	119, 120, 401, 34, 35, 74, 203, 127, 466, 32,
	131, 133, 76, 77, 365, 460, 60, 411, 447, 202,
	434, 107, 212, 114, 69, 68, 72, 433, 410, 407,
	406, 408, 409, 61, 425, 62, 65, 66, 67, 70,
	71, 366, 367, 368, 369, 370, 371, 327, 418, 73,
	273, 274, 275, 411, 108, 109, 110, 111, 112, 184,
	185, 186, 187, 188, 410, 407, 406, 408, 409, 403,
	405, 351, 352, 402, 189, 190, 404, 381, 191, 192,
	52, 130, 436, 12, 229, 230, 265, 266, 267, 268,
	269, 270, 271, 31, 126, 373, 219, 13, 14, 15,
	16, 137, 17, 37, 193, 194, 195, 196, 197, 198,
	199, 226, 357, 204, 175, 176, 177, 178, 179, 180,
	354, 34, 35, 312, 349, 282, 283, 322, 323, 138,
	261, 240, 241, 253, 104, 202, 34, 35, 64, 60,
	212, 347, 87, 277, 276, 343, 262, 338, 336, 334,
	294, 289, 260, 290, 252, 18, 225, 295, 34, 35,
	55, 244, 20, 19, 218, 304, 34, 35, 312, 245,
	427, 311, 311, 464, 129, 308, 309, 139, 140, 141,
	113, 142, 315, 310, 103, 89, 86, 462, 449, 319,
	69, 68, 72, 88, 463, 450, 101, 380, 379, 61,
	378, 62, 65, 66, 67, 70, 71, 377, 376, 361,
	360, 246, 243, 57, 58, 73, 239, 33, 238, 237,
	56, 316, 231, 220, 85, 81, 80, 79, 317, 208,
	332, 209, 210, 34, 35, 213, 78, 211, 304, 90,
	311, 25, 24, 23, 342, 22, 100, 96, 311, 34,
	35, 346, 339, 337, 335, 34, 35, 212, 315, 344,
	34, 35, 134, 440, 350, 353, 355, 348, 330, 358,
	39, 40, 45, 46, 47, 48, 49, 50, 41, 42,
	43, 44, 51, 34, 35, 132, 34, 35, 128, 34,
	35, 117, 333, 234, 465, 233, 34, 35, 345, 123,
	372, 124, 341, 374, 340, 302, 288, 287, 284, 254,
	251, 250, 249, 248, 247, 397, 232, 228, 224, 452,
	223, 419, 222, 221, 122, 121, 106, 426, 99, 98,
	83, 36, 28, 27, 26, 435, 451, 446, 445, 444,
	441, 442, 443, 439, 438, 437, 432, 431, 448, 145,
	146, 147, 148, 149, 150, 151, 152, 153, 154, 155,
	156, 157, 158, 159, 160, 161, 430, 459, 162, 163,
	164, 429, 428, 424, 423, 422, 421, 420, 375, 362,
	356, 331, 299, 298, 297, 296, 293, 461, 292, 291,
	459, 467, 165, 166, 167, 168, 169, 170, 171, 172,
	286, 285, 278, 263, 259, 258, 257, 256, 255, 242,
	236, 235, 227, 125, 102, 97, 94, 93, 92, 91,
	84, 82, 1, 53, 206, 21, 314, 306, 413, 399,
	454, 390, 383, 364, 279, 29, 201, 174, 329, 325,
	321, 318, 144, 301, 136, 75, 38, 11, 10, 9,
	8, 7, 6, 5, 4, 3, 2,
}
var yyPact = [...]int{

	137, -1000, 137, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 301, 299, 298, 297, 398, 397, 396, 45,
	395, -1000, 323, 220, 198, 10, 292, 283, 282, 281,
	-1000, 487, 394, 486, -1000, -1000, 280, 241, 323, 249,
	302, 485, 484, 483, 482, 311, 481, 393, 392, 308,
	257, 480, 239, 220, -1000, -1000, 317, 390, 59, 235,
	198, 317, 351, -1000, -1000, 317, 317, 317, 389, 388,
	365, 479, 59, 348, 229, 10, 345, 322, 165, 383,
	95, 32, -1000, -1000, -1000, 65, -1000, -1000, 295, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -19, -22, -26, -33, 219, 165, -1000, 279, 387,
	386, 384, 382, 211, 383, -1000, 478, -1000, -1000, 381,
	317, 317, 278, 380, 359, -1000, 477, 476, 275, 274,
	272, 148, 475, -1000, 268, 225, 267, -1000, 378, 377,
	376, 375, 374, 209, 95, 373, 474, 473, 472, 471,
	470, 207, 32, 198, 469, 54, 11, 54, 11, 468,
	-47, 317, 317, 372, -1000, -1000, 467, 466, 371, 370,
	206, 65, 455, 454, 452, 205, 295, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 451, 450, 449, 448, -1000, -1000,
	369, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 317, -1000, -1000, -1000, -1000, -1000, 183, 228, 32,
	-1000, -1000, -1000, 139, -1000, 25, 332, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 447, 358, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	204, 369, -1000, 203, 317, 202, 183, -1000, 368, 366,
	317, -1000, -1000, 200, 228, -1000, 362, 317, 196, 32,
	179, 139, 81, 295, 175, 25, -1000, 446, 167, 332,
	-59, -1000, 266, 265, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 445,
	-8, -8, -1000, 150, -8, 444, -1000, 264, 263, 256,
	254, 253, 132, -1000, -1000, -1000, -82, -95, -95, 5,
	-75, -1000, 103, -82, 443, 442, 441, 440, 439, 89,
	-95, 226, 438, 437, 432, 413, 412, 82, 75, 5,
	-1000, 138, 411, 410, 409, 327, 317, 407, 317, 405,
	404, 403, 73, -75, 252, 402, 385, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -89, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -31, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 70, -89, 251, 231, 360, 63, -31,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 492, 526, 525, 524, 523, 522, 521, 520, 519,
	518, 517, 173, 516, 20, 3, 150, 32, 75, 515,
	29, 514, 22, 513, 36, 512, 18, 15, 14, 511,
	12, 10, 9, 510, 11, 509, 508, 28, 507, 19,
	17, 25, 21, 506, 505, 16, 26, 24, 504, 8,
	503, 7, 6, 5, 4, 502, 501, 2, 500, 499,
	0, 1, 498, 497, 496, 13, 494, 493, 91,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 12, 12, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 4, 5, 6, 18, 18, 7, 20, 20,
	21, 21, 21, 21, 21, 21, 21, 22, 22, 23,
	8, 24, 24, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 30, 30, 33, 33, 33,
	33, 31, 31, 35, 35, 26, 26, 32, 32, 36,
	36, 36, 9, 37, 37, 38, 38, 38, 38, 38,
	38, 38, 29, 29, 39, 39, 11, 42, 42, 43,
	43, 43, 43, 10, 40, 40, 44, 44, 44, 44,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 49, 49,
	50, 50, 50, 50, 50, 50, 50, 50, 51, 51,
	55, 55, 55, 55, 55, 55, 52, 52, 56, 56,
	56, 56, 56, 56, 56, 57, 57, 58, 58, 58,
	58, 58, 53, 53, 59, 59, 59, 59, 59, 59,
	59, 61, 61, 60, 60, 60, 60, 60, 60, 60,
	54, 54, 62, 62, 62, 62, 62, 62, 46, 46,
	46, 46, 46, 46, 46, 46, 47, 47, 47, 47,
	48, 48, 48, 27, 27, 28, 28, 63, 63, 63,
	63, 64, 64, 64, 14, 14, 66, 16, 16, 67,
	67, 67, 67, 67, 67, 17, 17, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 19, 19, 19, 19, 19, 68,
	68, 68, 68, 68, 68, 65, 65, 65, 15, 15,
	15, 45, 45, 45, 34, 34, 34, 34, 34, 34,
	34,
}
var yyR2 = [...]int{

	0, 2, 1, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 4, 2, 1, 0, 4, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 4, 4, 4, 2, 1, 5, 2, 1,
	0, 1, 4, 2, 2, 2, 2, 2, 1, 1,
	5, 2, 1, 0, 1, 2, 1, 1, 2, 2,
	2, 4, 2, 2, 2, 1, 2, 2, 4, 4,
	4, 2, 2, 2, 1, 4, 2, 4, 4, 1,
	2, 2, 2, 2, 2, 2, 1, 0, 2, 2,
	2, 2, 1, 1, 2, 2, 1, 2, 1, 0,
	1, 3, 5, 2, 1, 0, 2, 2, 2, 2,
	2, 2, 2, 1, 2, 1, 5, 2, 1, 0,
	2, 2, 2, 5, 2, 1, 0, 1, 2, 2,
	0, 2, 2, 2, 2, 2, 2, 2, 3, 6,
	6, 2, 1, 1, 2, 2, 2, 2, 2, 1,
	0, 2, 1, 4, 4, 4, 4, 4, 2, 1,
	0, 2, 2, 2, 2, 2, 2, 1, 0, 4,
	2, 2, 2, 2, 2, 2, 1, 0, 2, 2,
	2, 2, 2, 1, 0, 1, 4, 2, 2, 2,
	2, 2, 1, 0, 2, 2, 2, 2, 2, 2,
	2, 1, 0, 2, 2, 2, 2, 1, 0, 1,
	1, 1, 1, 1, 1, 1, 0, 1, 1, 1,
	0, 1, 1, 2, 1, 2, 1, 1, 2, 2,
	2, 1, 2, 2, 2, 1, 1, 2, 1, 0,
	1, 1, 2, 2, 2, 2, 1, 0, 2, 2,
	2, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 0, 4, 4, 4, 4, 0,
	1, 1, 1, 1, 1, 0, 1, 1, 0, 1,
	1, 0, 1, 2, 0, 1, 1, 1, 1, 1,
	1,
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
	83, 84, 86, -24, -25, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 55, 56, 57, 79, 80, 81, 82, 83,
	84, 85, 86, -37, -38, 89, 90, 91, 92, 93,
	94, -40, -39, -41, 97, 98, 99, 100, 101, 112,
	113, 116, 117, 142, 143, 144, 145, 146, 147, 148,
	-42, -43, -15, 11, 118, -14, -66, -34, 4, 6,
	7, 12, -15, 10, 77, 77, 77, 77, 15, -20,
	14, 6, 6, 6, 6, 15, -24, 4, 6, -15,
	-15, 14, 6, 6, 4, 4, 4, 14, 14, 14,
	53, 54, 4, 14, 6, 14, 14, 6, 6, 6,
	6, 6, 15, -37, 6, 4, 4, 4, 4, 4,
	15, -40, -17, 4, -46, 102, 103, 104, 105, 106,
	107, 108, -47, 109, 110, 111, -46, -47, 4, -48,
	114, 115, -15, -15, 6, 4, 4, 6, 6, 15,
	-42, 4, 4, 4, 15, -14, 4, 4, 4, 4,
	-22, -23, 6, -26, -15, -27, -63, -65, 62, 63,
	70, -15, 10, -28, -64, -65, 63, 70, -29, -39,
	-30, -33, 58, 59, -31, -35, 6, 92, -32, -36,
	6, 4, -45, 4, 15, -22, 15, -26, 15, -27,
	6, 6, -15, 15, -28, 6, -15, 15, -40, 15,
	-30, 60, 61, -34, 15, -31, 4, 15, -32, 92,
	14, 14, 4, -49, -50, 92, 119, 120, 121, 122,
	123, 124, -49, 15, -49, 4, 14, 14, 14, 14,
	14, 15, -51, -55, 130, 129, 134, 138, 141, -52,
	-56, 125, 129, 130, 140, 141, 138, -52, -53, -59,
	-60, 67, 138, 134, 141, 135, 131, 130, 132, 133,
	129, 118, -54, -62, 136, 137, 138, 139, 15, -51,
	4, 4, 4, 4, 4, 15, -52, 14, 4, 4,
	4, 4, 4, 15, 15, -53, 14, 4, 4, 4,
	6, -15, 4, -15, 4, 4, 4, 15, -54, 6,
	13, 4, 4, -57, -58, 126, 127, 128, -61, -60,
	15, -57, 6, 13, 12, 4, 15, -61,
}
var yyDef = [...]int{

	3, -2, 2, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 0, 0, 0, 0, 0, 0, 0, -2,
	0, 1, 16, 239, 247, 264, 0, 0, 0, 0,
	127, 0, 0, 282, 279, 280, 0, 0, 15, 0,
	0, 0, 0, 0, 0, 278, 0, 0, 0, 0,
	0, 0, 0, 238, 240, 241, 278, 0, 269, 0,
	246, 278, 278, 251, 252, 278, 278, 278, 0, 0,
	0, 0, 269, 278, 0, 36, 278, 278, 40, 53,
	105, 130, 128, 129, 283, 119, 13, 14, 278, 18,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 32, 237, 242, 243, 244, 270, 271,
	272, 273, 274, 33, 245, 248, 249, 250, 253, 254,
	255, 256, 257, 258, 259, 260, 261, 262, 263, 34,
	35, 0, 0, 0, 0, 0, 39, 41, 0, 0,
	0, 0, 0, 0, 52, 54, 0, 56, 57, 0,
	278, 278, 0, 0, 0, 65, 0, 0, 0, 0,
	0, 0, 0, 74, 0, 0, 0, 79, 0, 0,
	0, 0, 0, 0, 104, 0, 0, 0, 0, 0,
	0, 0, 125, 115, 0, 208, 216, 208, 216, 0,
	220, 278, 278, 0, 142, 143, 0, 0, 0, 0,
	0, -2, 0, 0, 0, 0, 235, 236, 285, 286,
	287, 288, 289, 290, 0, 0, 0, 0, 37, 38,
	0, 43, 44, 45, 46, 50, 51, 55, 58, 59,
	60, 278, 62, 63, 64, 66, 67, 275, 275, 130,
	71, 72, 73, 87, 76, 0, 99, 80, 81, 82,
	83, 84, 102, 103, 106, 107, 108, 109, 110, 111,
	123, 124, 114, 131, 132, 209, 210, 211, 212, 213,
	214, 215, 133, 217, 218, 219, 134, 135, 136, 137,
	221, 222, 0, 278, 141, 144, 145, 146, 147, 116,
	117, 120, 121, 122, 17, 234, 265, 267, 266, 268,
	0, 48, 49, 0, 96, 0, 224, 227, 0, 0,
	278, 276, 277, 0, 226, 231, 0, 278, 0, 113,
	0, 86, 0, 278, 0, 92, 93, 0, 0, 98,
	100, 138, 0, 0, 42, 47, 61, 95, 68, 223,
	228, 229, 230, 69, 225, 232, 233, 70, 112, 75,
	85, 88, 89, 90, 77, 91, 94, 78, 97, 0,
	150, 150, 101, 0, 149, 0, 152, 0, 0, 0,
	0, 0, 0, 139, 148, 151, 160, 168, 168, 184,
	202, 140, 0, 159, 0, 0, 0, 0, 0, 0,
	167, 0, 0, 0, 0, 0, 0, 0, 0, 183,
	185, 0, 0, 0, 0, 0, 278, 0, 278, 0,
	0, 0, 0, 201, 0, 0, 0, 207, 153, 158,
	161, 162, 163, 164, 165, 154, 166, 177, 170, 171,
	172, 173, 174, 155, 156, 182, 193, 187, 188, 189,
	190, 194, 195, 196, 197, 198, 199, 157, 200, 203,
	204, 205, 206, 0, 176, 0, 0, 0, 0, 192,
	169, 175, 178, 179, 180, 181, 186, 191,
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
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:81
		{
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:83
		{
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:91
		{
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:92
		{
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:93
		{
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:94
		{
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:98
		{
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
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
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:102
		{
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:104
		{
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:105
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:111
		{
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:112
		{
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:113
		{
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:114
		{
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:115
		{
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:116
		{
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
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
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:122
		{
		}
	case 87:
		yyDollar = yyS[yypt-0 : yypt+1]
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
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:129
		{
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:134
		{
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:140
		{
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:141
		{
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:142
		{
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
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
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:154
		{
		}
	case 119:
		yyDollar = yyS[yypt-0 : yypt+1]
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
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:167
		{
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:175
		{
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:176
		{
		}
	case 130:
		yyDollar = yyS[yypt-0 : yypt+1]
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:185
		{
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:186
		{
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:189
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
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:195
		{
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:197
		{
		}
	case 150:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:199
		{
		}
	case 151:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:200
		{
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
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
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:206
		{
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:208
		{
		}
	case 160:
		yyDollar = yyS[yypt-0 : yypt+1]
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
	case 165:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:215
		{
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:217
		{
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:219
		{
		}
	case 169:
		yyDollar = yyS[yypt-4 : yypt+1]
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
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:225
		{
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:227
		{
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
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
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:233
		{
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:235
		{
		}
	case 184:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:237
		{
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:238
		{
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:245
		{
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
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
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:253
		{
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:255
		{
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:261
		{
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:262
		{
		}
	case 208:
		yyDollar = yyS[yypt-0 : yypt+1]
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:271
		{
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
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
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:276
		{
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:278
		{
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:279
		{
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:280
		{
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:282
		{
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:284
		{
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:289
		{
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:291
		{
		}
	case 232:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:292
		{
		}
	case 233:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:293
		{
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:295
		{
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:297
		{
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:301
		{
		}
	case 247:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:310
		{
		}
	case 264:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:328
		{
		}
	case 269:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:334
		{
		}
	case 275:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:341
		{
		}
	case 278:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:345
		{
		}
	case 281:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:349
		{
		}
	case 284:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:353
		{
		}
	}
	goto yystack /* stack new state and value */
}
