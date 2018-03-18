//line parser/parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser/parser.go.y:2
//line parser/parser.go.y:5
type yySymType struct {
	yys        int
	string     string
	strings    []string
	token      Token
	block      Block
	blocks_any []BlockAny
	block_args BlockArgs
	stmt_any   StmtAny
	stmts_any  []StmtAny
	stmt       Stmt
	stmt_multi StmtMulti
	stmt_value StmtValue
	values     []Value
	vip_addr   VIPAddr
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
const PREEMPT_DELAY = 57389
const PRIORITY = 57390
const ADVERT_INT = 57391
const VIRTUAL_IPADDRESS = 57392
const VIRTUAL_IPADDRESS_EXCLUDED = 57393
const VIRTUAL_ROUTES = 57394
const STATE = 57395
const MASTER = 57396
const BACKUP = 57397
const GARP_MASTER_DELAY = 57398
const SMTP_ALERT = 57399
const AUTHENTICATION = 57400
const AUTH_TYPE = 57401
const AUTH_PASS = 57402
const PASS = 57403
const AH = 57404
const LABEL = 57405
const DEV = 57406
const SCOPE = 57407
const SITE = 57408
const LINK = 57409
const HOST = 57410
const NOWHERE = 57411
const GLOBAL = 57412
const BRD = 57413
const SRC = 57414
const FROM = 57415
const TO = 57416
const VIA = 57417
const GW = 57418
const OR = 57419
const TABLE = 57420
const METRIC = 57421
const TRACK_INTERFACE = 57422
const TRACK_SCRIPT = 57423
const DONT_TRACK_PRIMARY = 57424
const NOTIFY_MASTER = 57425
const NOTIFY_BACKUP = 57426
const NOTIFY_FAULT = 57427
const NOTIFY_STOP = 57428
const NOTIFY = 57429
const BLACKHOLE = 57430
const VRRP_SCRIPT = 57431
const SCRIPT = 57432
const INTERVAL = 57433
const TIMEOUT = 57434
const WEIGHT = 57435
const FALL = 57436
const RISE = 57437
const VIRTUAL_SERVER_GROUP = 57438
const VIRTUAL_SERVER = 57439
const DELAY_LOOP = 57440
const LB_ALGO = 57441
const LB_KIND = 57442
const LVS_SCHED = 57443
const LVS_METHOD = 57444
const RR = 57445
const WRR = 57446
const LC = 57447
const WLC = 57448
const FO = 57449
const OVF = 57450
const LBLC = 57451
const LBLCR = 57452
const SH = 57453
const DH = 57454
const SED = 57455
const NQ = 57456
const NAT = 57457
const DR = 57458
const TUN = 57459
const PERSISTENCE_TIMEOUT = 57460
const PROTOCOL = 57461
const TCP = 57462
const UDP = 57463
const SORRY_SERVER = 57464
const REAL_SERVER = 57465
const FWMARK = 57466
const INHIBIT_ON_FAILURE = 57467
const TCP_CHECK = 57468
const HTTP_GET = 57469
const SSL_GET = 57470
const SMTP_CHECK = 57471
const DNS_CHECK = 57472
const MISC_CHECK = 57473
const URL = 57474
const PATH = 57475
const DIGEST = 57476
const STATUS_CODE = 57477
const CONNECT_TIMEOUT = 57478
const CONNECT_PORT = 57479
const CONNECT_IP = 57480
const BINDTO = 57481
const BIND_PORT = 57482
const RETRY = 57483
const HELO_NAME = 57484
const TYPE = 57485
const NAME = 57486
const MISC_PATH = 57487
const MISC_TIMEOUT = 57488
const WARMUP = 57489
const MISC_DYNAMIC = 57490
const NB_GET_RETRY = 57491
const DELAY_BEFORE_RETRY = 57492
const VIRTUALHOST = 57493
const ALPHA = 57494
const OMEGA = 57495
const QUORUM = 57496
const HYSTERESIS = 57497
const QUORUM_UP = 57498
const QUORUM_DOWN = 57499

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
	"PREEMPT_DELAY",
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
	"FO",
	"OVF",
	"LBLC",
	"LBLCR",
	"SH",
	"DH",
	"SED",
	"NQ",
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
	"DNS_CHECK",
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
	"TYPE",
	"NAME",
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

//line parser/parser.go.y:480

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 558

var yyAct = [...]int{

	410, 490, 485, 62, 433, 422, 408, 399, 392, 371,
	334, 330, 207, 326, 181, 313, 319, 311, 309, 182,
	306, 205, 200, 30, 278, 173, 265, 53, 33, 34,
	63, 366, 401, 183, 142, 217, 402, 403, 487, 488,
	489, 58, 435, 436, 437, 438, 134, 406, 94, 404,
	405, 216, 395, 394, 332, 411, 53, 396, 59, 104,
	286, 287, 215, 397, 114, 115, 398, 214, 117, 118,
	119, 33, 34, 498, 203, 492, 126, 75, 76, 130,
	132, 73, 479, 68, 67, 71, 106, 51, 202, 457,
	429, 212, 60, 59, 61, 64, 65, 66, 69, 70,
	421, 113, 428, 425, 424, 426, 427, 430, 72, 431,
	432, 421, 420, 417, 416, 418, 419, 36, 184, 185,
	186, 187, 188, 420, 417, 416, 418, 419, 413, 415,
	373, 33, 34, 468, 412, 358, 359, 414, 189, 190,
	103, 333, 191, 192, 279, 280, 281, 175, 176, 177,
	178, 179, 180, 229, 230, 86, 129, 32, 125, 455,
	136, 454, 374, 375, 376, 377, 378, 379, 380, 328,
	329, 193, 194, 195, 196, 197, 198, 199, 226, 241,
	242, 446, 219, 439, 391, 382, 364, 204, 137, 33,
	34, 318, 361, 356, 354, 288, 289, 262, 350, 345,
	254, 343, 341, 245, 12, 202, 300, 295, 261, 253,
	212, 246, 448, 283, 282, 225, 218, 59, 13, 14,
	15, 16, 128, 17, 296, 263, 112, 494, 301, 102,
	85, 33, 34, 63, 495, 310, 390, 138, 139, 140,
	88, 141, 317, 317, 314, 315, 481, 31, 87, 33,
	34, 54, 316, 482, 99, 321, 389, 388, 387, 386,
	325, 266, 267, 268, 269, 270, 271, 272, 273, 274,
	275, 276, 277, 385, 368, 367, 247, 18, 107, 108,
	109, 110, 111, 244, 20, 19, 68, 67, 71, 33,
	34, 318, 240, 340, 239, 60, 238, 61, 64, 65,
	66, 69, 70, 231, 220, 56, 57, 84, 80, 79,
	78, 72, 55, 77, 310, 25, 317, 24, 23, 22,
	349, 496, 100, 95, 317, 33, 34, 353, 342, 344,
	346, 33, 34, 212, 339, 89, 321, 351, 33, 34,
	355, 357, 360, 362, 478, 322, 365, 33, 34, 133,
	477, 461, 323, 38, 39, 44, 45, 46, 47, 48,
	49, 40, 41, 42, 43, 50, 208, 336, 209, 210,
	33, 34, 213, 497, 211, 33, 34, 131, 381, 33,
	34, 127, 383, 33, 34, 116, 234, 122, 233, 123,
	352, 348, 347, 308, 294, 407, 293, 290, 255, 252,
	251, 250, 440, 249, 248, 232, 228, 224, 447, 223,
	222, 221, 121, 120, 105, 98, 456, 97, 83, 35,
	462, 28, 464, 27, 26, 484, 483, 52, 470, 469,
	472, 476, 475, 474, 473, 471, 467, 466, 465, 480,
	463, 460, 459, 458, 453, 452, 451, 450, 449, 445,
	444, 443, 442, 441, 384, 370, 369, 363, 491, 144,
	145, 146, 147, 148, 149, 150, 151, 152, 153, 154,
	155, 156, 157, 158, 159, 160, 161, 337, 305, 162,
	163, 164, 304, 303, 302, 299, 298, 297, 292, 493,
	291, 284, 491, 499, 264, 260, 259, 258, 257, 256,
	243, 237, 236, 165, 166, 167, 168, 169, 170, 171,
	172, 235, 227, 124, 101, 96, 93, 92, 91, 90,
	82, 81, 1, 206, 434, 21, 423, 409, 486, 400,
	393, 372, 338, 285, 335, 331, 327, 324, 307, 37,
	29, 320, 312, 74, 201, 174, 135, 143, 10, 11,
	9, 7, 6, 5, 4, 8, 3, 2,
}
var yyPact = [...]int{

	188, -1000, 188, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 305, 304, 303, 301, 418, 417, 415, 123,
	413, -1000, 336, 241, 223, 4, 299, 296, 295, 294,
	517, 516, 412, -1000, -1000, 293, 215, 336, 234, 328,
	515, 514, 513, 512, 317, 511, 411, 409, 246, 313,
	510, 214, 241, -1000, -1000, 323, 408, 212, 211, 223,
	323, 375, -1000, -1000, 323, 323, 323, 407, 406, 383,
	509, 212, 371, 207, 4, 367, 339, 154, 423, 57,
	20, -1000, -1000, -1000, 63, -1000, -1000, 362, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-11, -16, -27, -43, 201, 154, -1000, 290, 405, 404,
	403, 401, 200, 423, -1000, 508, -1000, -1000, 400, 323,
	323, 289, 399, 382, -1000, 507, 498, 497, 282, 280,
	278, 125, 496, -1000, 269, 197, 262, -1000, 398, 397,
	395, 394, 393, 194, 57, 392, 495, 494, 493, 492,
	491, 193, 20, 223, 490, 158, 29, 158, 29, 487,
	-60, 323, 323, 391, -1000, -1000, 486, 484, 390, 388,
	192, 63, 483, 482, 481, 191, 362, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 480, 479, 478, 474, -1000, -1000,
	387, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 323, -1000, -1000, -1000, -1000, -1000, -1000, 181, 281,
	20, -1000, -1000, -1000, 110, -1000, 48, 361, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 473, 330,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 187, 387, -1000, 186,
	323, 184, 181, -1000, 386, 385, 323, -1000, -1000, 183,
	281, -1000, 384, 323, 179, 20, 178, 110, 74, 362,
	177, 48, -1000, 453, 171, 361, -62, -1000, 261, 260,
	452, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 451, 37, 37, -1000,
	-1000, 170, 37, 450, -1000, 259, 245, 244, 243, 242,
	222, 169, -1000, -1000, -1000, -84, -100, -100, -13, -34,
	-103, -1000, 168, -84, 449, 448, 447, 446, 445, 166,
	-100, 198, 444, 443, 442, 441, 440, 146, 144, -13,
	-1000, 75, 439, 438, 437, 345, 323, 436, 323, 434,
	433, 432, 118, -34, 323, 431, 323, 430, 429, 428,
	427, 344, 338, 67, -103, 240, 422, 421, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -95, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -24, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 60, -95, 221, 309, 369,
	58, -24, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 522, 557, 556, 555, 554, 553, 552, 551, 550,
	549, 548, 117, 34, 46, 25, 87, 41, 81, 22,
	14, 547, 546, 545, 544, 19, 33, 543, 17, 16,
	542, 541, 540, 15, 3, 539, 21, 20, 538, 18,
	537, 13, 11, 10, 536, 12, 535, 534, 26, 24,
	533, 532, 9, 531, 8, 7, 6, 5, 4, 530,
	529, 2, 528, 527, 0, 1, 526, 524, 523, 427,
	86,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 3, 12, 12, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	35, 5, 6, 7, 18, 18, 8, 14, 14, 22,
	22, 22, 22, 22, 22, 22, 37, 37, 38, 4,
	13, 13, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 41, 41, 44, 44, 44,
	44, 42, 42, 46, 46, 39, 39, 43, 43, 47,
	47, 47, 9, 15, 15, 23, 23, 23, 23, 23,
	23, 23, 40, 40, 25, 25, 10, 19, 19, 24,
	24, 24, 24, 11, 20, 20, 32, 32, 32, 25,
	25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 25, 25, 52, 52, 53, 53,
	53, 53, 53, 53, 53, 53, 53, 54, 54, 59,
	59, 59, 59, 59, 59, 55, 55, 60, 60, 60,
	60, 60, 60, 60, 61, 61, 62, 62, 62, 62,
	62, 56, 56, 63, 63, 63, 63, 63, 63, 63,
	65, 65, 64, 64, 64, 64, 64, 64, 64, 57,
	57, 66, 66, 66, 66, 66, 66, 66, 66, 66,
	66, 58, 58, 67, 67, 67, 67, 67, 67, 48,
	48, 48, 48, 48, 48, 48, 48, 48, 48, 48,
	48, 48, 49, 49, 49, 49, 50, 50, 50, 28,
	28, 29, 29, 30, 30, 30, 30, 31, 31, 31,
	36, 36, 68, 16, 16, 69, 69, 69, 69, 69,
	69, 17, 17, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	27, 27, 27, 27, 70, 70, 70, 70, 70, 70,
	33, 33, 34, 34, 51, 51, 51, 45, 45, 45,
	45, 45, 45, 45,
}
var yyR2 = [...]int{

	0, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 4, 2, 1, 0, 4, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 4, 4, 4, 2, 1, 5, 2, 1, 0,
	1, 4, 2, 2, 2, 2, 2, 1, 1, 5,
	2, 1, 0, 1, 2, 1, 1, 2, 2, 2,
	4, 2, 2, 2, 1, 2, 2, 2, 4, 4,
	4, 2, 2, 2, 1, 4, 2, 4, 4, 1,
	2, 2, 2, 2, 2, 2, 1, 0, 2, 2,
	2, 2, 1, 1, 2, 2, 1, 2, 1, 0,
	1, 3, 5, 2, 1, 0, 2, 2, 2, 2,
	2, 2, 2, 1, 2, 1, 5, 2, 1, 0,
	2, 2, 2, 5, 2, 1, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 3, 6, 6, 2,
	1, 1, 2, 2, 2, 2, 2, 1, 0, 2,
	1, 4, 4, 4, 4, 4, 4, 2, 1, 0,
	2, 2, 2, 2, 2, 2, 1, 0, 4, 2,
	2, 2, 2, 2, 2, 1, 0, 2, 2, 2,
	2, 2, 1, 0, 1, 4, 2, 2, 2, 2,
	2, 1, 0, 2, 2, 2, 2, 2, 2, 2,
	1, 0, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 1, 0, 2, 2, 2, 2, 1, 0,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 0, 1, 1, 1, 0, 1, 1, 2,
	1, 2, 1, 1, 2, 2, 2, 1, 2, 2,
	2, 1, 1, 2, 1, 0, 1, 1, 2, 2,
	2, 2, 1, 0, 2, 2, 2, 1, 1, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	4, 4, 4, 4, 0, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 0, 1, 2, 0, 1, 1,
	1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -5, -6, -7, -8, -4, -9,
	-11, -10, 16, 30, 31, 32, 33, 35, 89, 97,
	96, -1, 14, 14, 14, 14, 6, 6, 6, -32,
	-34, 124, 34, 8, 9, 6, -12, -35, 17, 18,
	25, 26, 27, 28, 19, 20, 21, 22, 23, 24,
	29, -16, -69, -34, 10, 71, 64, 65, -17, -26,
	72, 74, -34, 10, 75, 76, 77, 64, 63, 78,
	79, 65, 88, -18, -27, 73, 74, 14, 14, 14,
	14, 4, 4, 6, 14, 15, -12, 14, 6, 7,
	4, 4, 4, 4, -34, 6, 4, 6, 6, 8,
	9, 4, 15, -16, -34, 6, -70, 66, 67, 68,
	69, 70, 15, -17, -34, -34, 10, -34, -34, -34,
	6, 6, 4, 6, 4, -70, -34, 10, 15, -18,
	-34, 10, -34, 10, -14, -22, 6, 34, 83, 84,
	85, 87, -13, -21, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 56, 57, 58, 80, 81, 82, 83, 84,
	85, 86, 87, -15, -23, 90, 91, 92, 93, 94,
	95, -20, -25, -26, 98, 99, 100, 101, 102, 118,
	119, 122, 123, 151, 152, 153, 154, 155, 156, 157,
	-19, -24, -34, 11, 124, -36, -68, -45, 4, 6,
	7, 12, -34, 10, 78, 78, 78, 78, 15, -14,
	14, 6, 6, 6, 6, 15, -13, 4, 6, -34,
	-34, 14, 6, 6, 4, 4, 4, 4, 14, 14,
	14, 54, 55, 4, 14, 6, 14, 14, 6, 6,
	6, 6, 6, 15, -15, 6, 4, 4, 4, 4,
	4, 15, -20, -17, 4, -48, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114, -49, 115,
	116, 117, -48, -49, 4, -50, 120, 121, -34, -34,
	6, 4, 4, 6, 6, 15, -19, 4, 4, 4,
	15, -36, 4, 4, 4, 4, -37, -38, 6, -39,
	-34, -28, -30, -33, 63, 64, 71, -34, 10, -29,
	-31, -33, 64, 71, -40, -25, -41, -44, 59, 60,
	-42, -46, 6, 93, -43, -47, 6, 4, -51, 4,
	-34, 15, -37, 15, -39, 15, -28, 6, 6, -34,
	15, -29, 6, -34, 15, -20, 15, -41, 61, 62,
	-45, 15, -42, 4, 15, -43, 93, 14, 14, 4,
	4, -52, -53, 93, 125, 126, 127, 128, 129, 130,
	131, -52, 15, -52, 4, 14, 14, 14, 14, 14,
	14, 15, -54, -59, 137, 136, 141, 147, 150, -55,
	-60, 132, 136, 137, 149, 150, 147, -55, -56, -63,
	-64, 68, 147, 141, 150, 142, 138, 137, 139, 140,
	136, 124, -57, -66, 138, 137, 139, 140, 136, 124,
	141, 143, 144, -58, -67, 145, 146, 147, 148, 15,
	-54, 4, 4, 4, 4, 4, 15, -55, 14, 4,
	4, 4, 4, 4, 15, 15, -56, 14, 4, 4,
	4, 6, -34, 4, -34, 4, 4, 4, 15, -57,
	-34, 4, -34, 4, 4, 4, 4, 6, 6, 15,
	-58, 6, 13, 4, 4, -61, -62, 133, 134, 135,
	-65, -64, 15, -61, 6, 13, 12, 4, 15, -65,
}
var yyDef = [...]int{

	0, -2, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 1, 15, 255, 263, 0, 0, 0, 0, 0,
	0, 0, 0, 292, 293, 0, 0, 14, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 254, 256, 257, 0, 0, 284, 0, 262,
	0, 0, 267, 268, 0, 0, 0, 0, 0, 0,
	0, 284, 0, 0, 35, 0, 0, 39, 52, 105,
	263, 126, 127, 128, 119, 12, 13, 297, 17, 18,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 253, 258, 259, 260, 285, 286, 287,
	288, 289, 32, 261, 264, 265, 266, 269, 270, 271,
	272, 273, 274, 275, 276, 277, 278, 279, 33, 34,
	0, 0, 0, 0, 0, 38, 40, 0, 0, 0,
	0, 0, 0, 51, 53, 0, 55, 56, 0, 0,
	0, 0, 0, 0, 64, 0, 0, 0, 0, 0,
	0, 0, 0, 74, 0, 0, 0, 79, 0, 0,
	0, 0, 0, 0, 104, 0, 0, 0, 0, 0,
	0, 0, 125, 115, 0, 219, 232, 219, 232, 0,
	236, 0, 0, 0, 140, 141, 0, 0, 0, 0,
	0, 118, 0, 0, 0, 0, 251, 252, 298, 299,
	300, 301, 302, 303, 0, 0, 0, 0, 36, 37,
	0, 42, 43, 44, 45, 49, 50, 54, 57, 58,
	59, 0, 61, 62, 63, 65, 66, 67, 0, 0,
	263, 71, 72, 73, 87, 76, 0, 99, 80, 81,
	82, 83, 84, 102, 103, 106, 107, 108, 109, 110,
	111, 123, 124, 114, 129, 130, 220, 221, 222, 223,
	224, 225, 226, 227, 228, 229, 230, 231, 131, 233,
	234, 235, 132, 133, 134, 135, 237, 238, 0, 294,
	139, 142, 143, 144, 145, 116, 117, 120, 121, 122,
	16, 250, 280, 282, 281, 283, 0, 47, 48, 0,
	96, 0, 240, 243, 0, 0, 0, 290, 291, 0,
	242, 247, 0, 0, 0, 113, 0, 86, 0, 297,
	0, 92, 93, 0, 0, 98, 100, 136, 0, 0,
	295, 41, 46, 60, 95, 68, 239, 244, 245, 246,
	69, 241, 248, 249, 70, 112, 75, 85, 88, 89,
	90, 77, 91, 94, 78, 97, 0, 148, 148, 296,
	101, 0, 147, 0, 150, 0, 0, 0, 0, 0,
	0, 0, 137, 146, 149, 159, 167, 167, 183, 201,
	213, 138, 0, 158, 0, 0, 0, 0, 0, 0,
	166, 0, 0, 0, 0, 0, 0, 0, 0, 182,
	184, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 200, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 212, 0, 0, 0, 218, 151,
	157, 160, 161, 162, 163, 164, 152, 165, 176, 169,
	170, 171, 172, 173, 153, 154, 181, 192, 186, 187,
	188, 189, 193, 194, 195, 196, 197, 198, 155, 199,
	202, 203, 204, 205, 206, 207, 208, 209, 210, 156,
	211, 214, 215, 216, 217, 0, 175, 0, 0, 0,
	0, 191, 168, 174, 177, 178, 179, 180, 185, 190,
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
	142, 143, 144, 145, 146, 147, 148, 149, 150, 151,
	152, 153, 154, 155, 156, 157,
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

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:36
		{
			yyVAL.blocks_any = append(yyDollar[1].blocks_any, yyDollar[2].blocks_any...)
			yylex.(*Lexer).result = yyVAL.blocks_any
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:41
		{
			yyVAL.blocks_any = yyDollar[1].blocks_any
			yylex.(*Lexer).result = yyVAL.blocks_any
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:47
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:48
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:49
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:50
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:51
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:52
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:53
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:54
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block_args}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:55
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 12:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:59
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[3].stmts_any}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:63
		{
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:63
		{
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:66
		{
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:67
		{
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:68
		{
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:69
		{
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:70
		{
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:71
		{
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:72
		{
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:73
		{
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:74
		{
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:75
		{
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:76
		{
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:77
		{
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:78
		{
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:79
		{
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:80
		{
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:84
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[3].stmts_any}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:90
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[3].stmts_any}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:96
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[3].stmts_any}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:101
		{
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:101
		{
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:105
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[4].stmts_any}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:111
		{
			yyVAL.stmts_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:114
		{
			yyVAL.stmts_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:116
		{
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:117
		{
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:118
		{
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:119
		{
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:120
		{
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:121
		{
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:122
		{
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:124
		{
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:130
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[4].stmts_any}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:136
		{
			yyVAL.stmts_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:139
		{
			yyVAL.stmts_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:141
		{
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:142
		{
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:143
		{
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:144
		{
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:145
		{
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:146
		{
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:147
		{
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:148
		{
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:149
		{
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:150
		{
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:151
		{
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:152
		{
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:153
		{
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:154
		{
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:155
		{
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:156
		{
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:158
		{
			yyVAL.stmt_any = StmtMulti{yyDollar[1].token.lit: yyDollar[3].values}
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:162
		{
			yyVAL.stmt_any = StmtMulti{yyDollar[1].token.lit: yyDollar[3].values}
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:165
		{
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:166
		{
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:167
		{
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:168
		{
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:169
		{
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:170
		{
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:171
		{
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:172
		{
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:173
		{
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:174
		{
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:175
		{
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:176
		{
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:177
		{
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:178
		{
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:179
		{
		}
	case 87:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:183
		{
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:184
		{
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:185
		{
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:186
		{
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:191
		{
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:197
		{
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:198
		{
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:199
		{
		}
	case 102:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:203
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[4].stmts_any}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:209
		{
			yyVAL.stmts_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:212
		{
			yyVAL.stmts_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:214
		{
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:215
		{
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:216
		{
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:217
		{
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:218
		{
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:219
		{
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:220
		{
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:226
		{
			yyVAL.stmt_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:229
		{
			yyVAL.stmt_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:233
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[4].stmts_any}
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:239
		{
			yyVAL.stmts_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:242
		{
			yyVAL.stmts_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 119:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:244
		{
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:245
		{
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:246
		{
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:247
		{
		}
	case 123:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:251
		{
			yyVAL.block_args = BlockArgs{Name: yyDollar[1].token.lit, Args: yyDollar[2].strings, Stmts: yyDollar[4].stmts_any}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:257
		{
			yyVAL.stmts_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:260
		{
			yyVAL.stmts_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:263
		{
			yyVAL.strings = []string{yyDollar[1].string, yyDollar[2].token.lit}
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:264
		{
			yyVAL.strings = []string{yyDollar[1].token.lit, yyDollar[2].token.lit}
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:265
		{
			yyVAL.strings = []string{yyDollar[1].token.lit, yyDollar[2].token.lit}
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:268
		{
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:269
		{
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:270
		{
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:271
		{
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:272
		{
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:273
		{
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:274
		{
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:275
		{
		}
	case 137:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:276
		{
		}
	case 138:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:277
		{
		}
	case 139:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:278
		{
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:279
		{
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:280
		{
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:281
		{
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:282
		{
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:283
		{
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:284
		{
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:286
		{
		}
	case 148:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:288
		{
		}
	case 149:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:289
		{
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:290
		{
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:291
		{
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:292
		{
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:293
		{
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:294
		{
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:295
		{
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:296
		{
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:298
		{
		}
	case 159:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:300
		{
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:301
		{
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:302
		{
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:303
		{
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:304
		{
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:305
		{
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:307
		{
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:309
		{
		}
	case 168:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:310
		{
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:311
		{
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:312
		{
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:313
		{
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:314
		{
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:315
		{
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:317
		{
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:319
		{
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:320
		{
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:321
		{
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:322
		{
		}
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:323
		{
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:325
		{
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:327
		{
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:328
		{
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:335
		{
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:337
		{
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:338
		{
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:339
		{
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:340
		{
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:341
		{
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:342
		{
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:343
		{
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:345
		{
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:347
		{
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:348
		{
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:349
		{
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:350
		{
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:351
		{
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:352
		{
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:353
		{
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:354
		{
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:355
		{
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:356
		{
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:358
		{
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:360
		{
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:361
		{
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:362
		{
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:363
		{
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:364
		{
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:365
		{
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:367
		{
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:368
		{
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:369
		{
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:370
		{
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:371
		{
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:372
		{
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:373
		{
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:374
		{
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:375
		{
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:376
		{
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:377
		{
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:378
		{
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:379
		{
		}
	case 232:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:381
		{
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:382
		{
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:383
		{
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:384
		{
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:386
		{
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:387
		{
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:388
		{
		}
	case 239:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:392
		{
			yyVAL.values = append([]Value{yyDollar[1].vip_addr}, yyDollar[2].values...)
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:395
		{
			yyVAL.values = []Value{yyDollar[1].vip_addr}
		}
	case 241:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:399
		{
			yyVAL.values = append([]Value{yyDollar[1].vip_addr}, yyDollar[2].values...)
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:402
		{
			yyVAL.values = []Value{yyDollar[1].vip_addr}
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:405
		{
			yyVAL.vip_addr = VIPAddr{Addr: yyDollar[1].string}
		}
	case 244:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:406
		{
		}
	case 245:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:407
		{
		}
	case 246:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:408
		{
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:410
		{
		}
	case 248:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:411
		{
		}
	case 249:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:412
		{
		}
	case 251:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:414
		{
		}
	case 252:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:416
		{
		}
	case 253:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:418
		{
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:418
		{
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:420
		{
		}
	case 261:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:427
		{
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:427
		{
		}
	case 263:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:429
		{
		}
	case 264:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:430
		{
		}
	case 265:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:431
		{
		}
	case 266:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:432
		{
		}
	case 267:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:433
		{
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:434
		{
		}
	case 269:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:435
		{
		}
	case 270:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:436
		{
		}
	case 271:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:437
		{
		}
	case 272:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:438
		{
		}
	case 273:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:439
		{
		}
	case 274:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:440
		{
		}
	case 275:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:441
		{
		}
	case 276:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:442
		{
		}
	case 277:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:443
		{
		}
	case 278:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:444
		{
		}
	case 279:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:445
		{
		}
	case 280:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:448
		{
		}
	case 281:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:449
		{
		}
	case 282:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:450
		{
		}
	case 283:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:451
		{
		}
	case 284:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:453
		{
		}
	case 290:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:461
		{
			yyVAL.string = yyDollar[1].string
		}
	case 291:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:462
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 292:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:465
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 293:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:466
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 294:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:468
		{
		}
	case 297:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:472
		{
		}
	}
	goto yystack /* stack new state and value */
}
