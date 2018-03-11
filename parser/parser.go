//line parser/parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser/parser.go.y:2
import "fmt"

//line parser/parser.go.y:6
type yySymType struct {
	yys    int
	token  Token
	block  Block
	blocks []Block
	stmt   Stmt
	stmts  []Stmt
	value  Value
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

//line parser/parser.go.y:417

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 558

var yyAct = [...]int{

	409, 489, 484, 63, 432, 421, 407, 398, 391, 370,
	335, 331, 208, 327, 182, 314, 320, 312, 310, 30,
	307, 183, 206, 33, 201, 279, 174, 54, 266, 184,
	34, 35, 64, 400, 59, 410, 143, 401, 402, 287,
	288, 394, 393, 486, 487, 488, 395, 135, 405, 95,
	403, 404, 396, 366, 60, 397, 218, 54, 217, 216,
	105, 434, 435, 436, 437, 115, 116, 34, 35, 118,
	119, 120, 34, 35, 333, 204, 215, 127, 76, 77,
	131, 133, 497, 107, 491, 69, 68, 72, 52, 203,
	60, 420, 213, 32, 61, 114, 62, 65, 66, 67,
	70, 71, 478, 419, 416, 415, 417, 418, 412, 414,
	73, 74, 456, 428, 411, 358, 359, 413, 467, 454,
	185, 186, 187, 188, 189, 427, 424, 423, 425, 426,
	429, 372, 430, 431, 420, 280, 281, 282, 329, 330,
	190, 191, 104, 12, 192, 193, 419, 416, 415, 417,
	418, 242, 243, 453, 230, 231, 126, 13, 14, 15,
	16, 334, 17, 373, 374, 375, 376, 377, 378, 379,
	37, 137, 445, 194, 195, 196, 197, 198, 199, 200,
	447, 227, 438, 31, 220, 390, 381, 130, 205, 176,
	177, 178, 179, 180, 181, 364, 289, 290, 263, 138,
	361, 356, 255, 354, 350, 345, 203, 343, 341, 87,
	301, 213, 296, 262, 60, 284, 18, 283, 254, 264,
	34, 35, 55, 20, 19, 226, 246, 297, 219, 129,
	302, 34, 35, 64, 247, 389, 311, 113, 495, 103,
	34, 35, 319, 318, 318, 86, 100, 388, 139, 140,
	141, 387, 142, 34, 35, 319, 322, 108, 109, 110,
	111, 112, 386, 326, 267, 268, 269, 270, 271, 272,
	273, 274, 275, 276, 277, 278, 57, 58, 493, 385,
	384, 89, 368, 56, 480, 494, 69, 68, 72, 88,
	367, 481, 101, 248, 33, 61, 323, 62, 65, 66,
	67, 70, 71, 324, 245, 241, 240, 239, 315, 316,
	339, 73, 232, 221, 85, 311, 317, 318, 81, 80,
	79, 349, 78, 25, 24, 318, 23, 22, 353, 342,
	344, 346, 90, 96, 213, 34, 35, 322, 351, 477,
	476, 355, 357, 360, 362, 34, 35, 365, 39, 40,
	45, 46, 47, 48, 49, 50, 41, 42, 43, 44,
	51, 209, 460, 210, 211, 34, 35, 214, 496, 212,
	34, 35, 134, 34, 35, 132, 337, 235, 380, 234,
	483, 382, 34, 35, 128, 34, 35, 117, 340, 123,
	482, 124, 34, 35, 406, 352, 348, 347, 309, 295,
	294, 439, 291, 256, 253, 252, 251, 446, 250, 249,
	233, 229, 225, 224, 223, 455, 222, 122, 121, 461,
	106, 463, 99, 98, 83, 36, 28, 469, 468, 471,
	27, 26, 475, 474, 473, 472, 470, 466, 479, 465,
	464, 462, 459, 458, 457, 452, 451, 450, 449, 448,
	444, 443, 442, 441, 440, 383, 369, 490, 145, 146,
	147, 148, 149, 150, 151, 152, 153, 154, 155, 156,
	157, 158, 159, 160, 161, 162, 363, 338, 163, 164,
	165, 306, 305, 304, 303, 300, 299, 298, 492, 293,
	292, 490, 498, 285, 265, 261, 260, 259, 258, 257,
	244, 238, 166, 167, 168, 169, 170, 171, 172, 173,
	237, 236, 228, 125, 102, 97, 94, 93, 92, 91,
	84, 82, 1, 53, 207, 21, 433, 422, 408, 485,
	399, 392, 371, 286, 29, 202, 175, 336, 332, 328,
	325, 308, 136, 75, 38, 11, 10, 9, 7, 6,
	5, 4, 3, 321, 313, 8, 144, 2,
}
var yyPact = [...]int{

	127, -1000, 127, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 313, 312, 310, 309, 425, 424, 420, 59,
	419, -1000, 331, 212, 223, 5, 308, 306, 305, 304,
	-1000, 517, 418, 516, -1000, -1000, 300, 230, 331, 275,
	325, 515, 514, 513, 512, 327, 511, 417, 416, 238,
	283, 510, 224, 212, -1000, -1000, 337, 414, 191, 222,
	223, 337, 377, -1000, -1000, 337, 337, 337, 412, 411,
	385, 509, 191, 374, 214, 5, 365, 362, 165, 422,
	99, 22, -1000, -1000, -1000, 64, -1000, -1000, 357, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -2, -19, -20, -22, 213, 165, -1000, 299, 410,
	408, 407, 406, 210, 422, -1000, 508, -1000, -1000, 405,
	337, 337, 298, 404, 373, -1000, 507, 506, 497, 293,
	292, 291, 97, 496, -1000, 290, 220, 279, -1000, 403,
	402, 400, 399, 398, 203, 99, 397, 495, 494, 493,
	492, 491, 198, 22, 223, 490, 161, 20, 161, 20,
	489, -81, 337, 337, 396, -1000, -1000, 486, 485, 394,
	393, 197, 64, 483, 482, 481, 195, 357, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 480, 479, 478, 477, -1000,
	-1000, 392, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 337, -1000, -1000, -1000, -1000, -1000, -1000, 245,
	232, 22, -1000, -1000, -1000, 79, -1000, 68, 370, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 473,
	384, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 193, 392, -1000,
	192, 337, 190, 245, -1000, 391, 390, 337, -1000, -1000,
	189, 232, -1000, 389, 337, 188, 22, 186, 79, 54,
	357, 185, 68, -1000, 472, 180, 370, -40, -1000, 276,
	268, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 452, 38, 38, -1000,
	171, 38, 451, -1000, 266, 265, 248, 237, 233, 221,
	170, -1000, -1000, -1000, -95, -99, -99, -33, -11, -84,
	-1000, 167, -95, 450, 449, 448, 447, 446, 157, -99,
	166, 445, 444, 443, 442, 441, 138, 104, -33, -1000,
	98, 440, 439, 438, 356, 337, 437, 337, 436, 435,
	433, 103, -11, 337, 432, 337, 431, 430, 429, 428,
	334, 333, 87, -84, 278, 386, 376, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -90, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 10, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 69, -90, 272, 226, 364, 67,
	10, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 522, 557, 36, 556, 555, 17, 16, 554, 553,
	15, 3, 552, 551, 550, 549, 548, 547, 546, 545,
	170, 544, 22, 88, 34, 111, 543, 47, 542, 20,
	541, 18, 540, 13, 11, 10, 539, 12, 538, 537,
	26, 536, 21, 14, 29, 24, 535, 534, 19, 28,
	25, 533, 9, 532, 8, 7, 6, 5, 4, 531,
	530, 2, 529, 528, 0, 1, 527, 526, 524, 523,
	83,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 12, 20, 20, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 13, 14, 15, 25, 25, 16, 27, 27,
	28, 28, 28, 28, 28, 28, 28, 29, 29, 30,
	5, 3, 3, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4, 4, 4, 4, 33, 33, 36, 36,
	36, 36, 34, 34, 38, 38, 31, 31, 35, 35,
	39, 39, 39, 17, 40, 40, 41, 41, 41, 41,
	41, 41, 41, 32, 32, 42, 42, 19, 45, 45,
	46, 46, 46, 46, 18, 43, 43, 47, 47, 47,
	47, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 52,
	52, 53, 53, 53, 53, 53, 53, 53, 53, 53,
	54, 54, 59, 59, 59, 59, 59, 59, 55, 55,
	60, 60, 60, 60, 60, 60, 60, 61, 61, 62,
	62, 62, 62, 62, 56, 56, 63, 63, 63, 63,
	63, 63, 63, 65, 65, 64, 64, 64, 64, 64,
	64, 64, 57, 57, 66, 66, 66, 66, 66, 66,
	66, 66, 66, 66, 58, 58, 67, 67, 67, 67,
	67, 67, 49, 49, 49, 49, 49, 49, 49, 49,
	49, 49, 49, 49, 49, 50, 50, 50, 50, 51,
	51, 51, 6, 6, 7, 7, 8, 8, 8, 8,
	9, 9, 9, 22, 22, 68, 23, 23, 69, 69,
	69, 69, 69, 69, 24, 24, 44, 44, 44, 44,
	44, 44, 44, 44, 44, 44, 44, 44, 44, 44,
	44, 44, 44, 26, 26, 26, 26, 26, 70, 70,
	70, 70, 70, 70, 10, 10, 11, 11, 48, 48,
	48, 37, 37, 37, 37, 37, 37, 37,
}
var yyR2 = [...]int{

	0, 2, 1, 0, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 4, 2, 1, 0, 4, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 4, 4, 4, 2, 1, 5, 2, 1,
	0, 1, 4, 2, 2, 2, 2, 2, 1, 1,
	5, 2, 1, 0, 1, 2, 1, 1, 2, 2,
	2, 4, 2, 2, 2, 1, 2, 2, 2, 4,
	4, 4, 2, 2, 2, 1, 4, 2, 4, 4,
	1, 2, 2, 2, 2, 2, 2, 1, 0, 2,
	2, 2, 2, 1, 1, 2, 2, 1, 2, 1,
	0, 1, 3, 5, 2, 1, 0, 2, 2, 2,
	2, 2, 2, 2, 1, 2, 1, 5, 2, 1,
	0, 2, 2, 2, 5, 2, 1, 0, 1, 2,
	2, 0, 2, 2, 2, 2, 2, 2, 2, 3,
	6, 6, 2, 1, 1, 2, 2, 2, 2, 2,
	1, 0, 2, 1, 4, 4, 4, 4, 4, 4,
	2, 1, 0, 2, 2, 2, 2, 2, 2, 1,
	0, 4, 2, 2, 2, 2, 2, 2, 1, 0,
	2, 2, 2, 2, 2, 1, 0, 1, 4, 2,
	2, 2, 2, 2, 1, 0, 2, 2, 2, 2,
	2, 2, 2, 1, 0, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 1, 0, 2, 2, 2,
	2, 1, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 0, 1, 1, 1, 0,
	1, 1, 2, 1, 2, 1, 1, 2, 2, 2,
	1, 2, 2, 2, 1, 1, 2, 1, 0, 1,
	1, 2, 2, 2, 2, 1, 0, 2, 2, 2,
	1, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 0, 4, 4, 4, 4, 0, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 0, 1,
	2, 0, 1, 1, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -12, -13, -14, -15, -16, -5, -17,
	-18, -19, 16, 30, 31, 32, 33, 35, 89, 97,
	96, -1, 14, 14, 14, 14, 6, 6, 6, -47,
	-48, 124, 34, -11, 8, 9, 6, -20, -21, 17,
	18, 25, 26, 27, 28, 19, 20, 21, 22, 23,
	24, 29, -23, -69, -11, 10, 71, 64, 65, -24,
	-44, 72, 74, -11, 10, 75, 76, 77, 64, 63,
	78, 79, 65, 88, -25, -26, 73, 74, 14, 14,
	14, 14, 4, 6, 4, 14, 15, -20, 14, 6,
	7, 4, 4, 4, 4, -11, 6, 4, 6, 6,
	8, 9, 4, 15, -23, -11, 6, -70, 66, 67,
	68, 69, 70, 15, -24, -11, -11, 10, -11, -11,
	-11, 6, 6, 4, 6, 4, -70, -11, 10, 15,
	-25, -11, 10, -11, 10, -27, -28, 6, 34, 83,
	84, 85, 87, -3, -4, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 56, 57, 58, 80, 81, 82, 83,
	84, 85, 86, 87, -40, -41, 90, 91, 92, 93,
	94, 95, -43, -42, -44, 98, 99, 100, 101, 102,
	118, 119, 122, 123, 151, 152, 153, 154, 155, 156,
	157, -45, -46, -11, 11, 124, -22, -68, -37, 4,
	6, 7, 12, -11, 10, 78, 78, 78, 78, 15,
	-27, 14, 6, 6, 6, 6, 15, -3, 4, 6,
	-11, -11, 14, 6, 6, 4, 4, 4, 4, 14,
	14, 14, 54, 55, 4, 14, 6, 14, 14, 6,
	6, 6, 6, 6, 15, -40, 6, 4, 4, 4,
	4, 4, 15, -43, -24, 4, -49, 103, 104, 105,
	106, 107, 108, 109, 110, 111, 112, 113, 114, -50,
	115, 116, 117, -49, -50, 4, -51, 120, 121, -11,
	-11, 6, 4, 4, 6, 6, 15, -45, 4, 4,
	4, 15, -22, 4, 4, 4, 4, -29, -30, 6,
	-31, -11, -6, -8, -10, 63, 64, 71, -11, 10,
	-7, -9, -10, 64, 71, -32, -42, -33, -36, 59,
	60, -34, -38, 6, 93, -35, -39, 6, 4, -48,
	4, 15, -29, 15, -31, 15, -6, 6, 6, -11,
	15, -7, 6, -11, 15, -43, 15, -33, 61, 62,
	-37, 15, -34, 4, 15, -35, 93, 14, 14, 4,
	-52, -53, 93, 125, 126, 127, 128, 129, 130, 131,
	-52, 15, -52, 4, 14, 14, 14, 14, 14, 14,
	15, -54, -59, 137, 136, 141, 147, 150, -55, -60,
	132, 136, 137, 149, 150, 147, -55, -56, -63, -64,
	68, 147, 141, 150, 142, 138, 137, 139, 140, 136,
	124, -57, -66, 138, 137, 139, 140, 136, 124, 141,
	143, 144, -58, -67, 145, 146, 147, 148, 15, -54,
	4, 4, 4, 4, 4, 15, -55, 14, 4, 4,
	4, 4, 4, 15, 15, -56, 14, 4, 4, 4,
	6, -11, 4, -11, 4, 4, 4, 15, -57, -11,
	4, -11, 4, 4, 4, 4, 6, 6, 15, -58,
	6, 13, 4, 4, -61, -62, 133, 134, 135, -65,
	-64, 15, -61, 6, 13, 12, 4, 15, -65,
}
var yyDef = [...]int{

	3, -2, 2, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 0, 0, 0, 0, 0, 0, 0, 127,
	0, 1, 16, 258, 266, 283, 0, 0, 0, 0,
	128, 0, 0, 299, 296, 297, 0, 0, 15, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 257, 259, 260, 0, 0, 288, 0,
	265, 0, 0, 270, 271, 0, 0, 0, 0, 0,
	0, 0, 288, 0, 0, 36, 0, 0, 40, 53,
	106, 131, 129, 130, 300, 120, 13, 14, 301, 18,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 32, 256, 261, 262, 263, 289, 290,
	291, 292, 293, 33, 264, 267, 268, 269, 272, 273,
	274, 275, 276, 277, 278, 279, 280, 281, 282, 34,
	35, 0, 0, 0, 0, 0, 39, 41, 0, 0,
	0, 0, 0, 0, 52, 54, 0, 56, 57, 0,
	0, 0, 0, 0, 0, 65, 0, 0, 0, 0,
	0, 0, 0, 0, 75, 0, 0, 0, 80, 0,
	0, 0, 0, 0, 0, 105, 0, 0, 0, 0,
	0, 0, 0, 126, 116, 0, 222, 235, 222, 235,
	0, 239, 0, 0, 0, 143, 144, 0, 0, 0,
	0, 0, 119, 0, 0, 0, 0, 254, 255, 302,
	303, 304, 305, 306, 307, 0, 0, 0, 0, 37,
	38, 0, 43, 44, 45, 46, 50, 51, 55, 58,
	59, 60, 0, 62, 63, 64, 66, 67, 68, 0,
	0, 131, 72, 73, 74, 88, 77, 0, 100, 81,
	82, 83, 84, 85, 103, 104, 107, 108, 109, 110,
	111, 112, 124, 125, 115, 132, 133, 223, 224, 225,
	226, 227, 228, 229, 230, 231, 232, 233, 234, 134,
	236, 237, 238, 135, 136, 137, 138, 240, 241, 0,
	298, 142, 145, 146, 147, 148, 117, 118, 121, 122,
	123, 17, 253, 284, 286, 285, 287, 0, 48, 49,
	0, 97, 0, 243, 246, 0, 0, 0, 294, 295,
	0, 245, 250, 0, 0, 0, 114, 0, 87, 0,
	301, 0, 93, 94, 0, 0, 99, 101, 139, 0,
	0, 42, 47, 61, 96, 69, 242, 247, 248, 249,
	70, 244, 251, 252, 71, 113, 76, 86, 89, 90,
	91, 78, 92, 95, 79, 98, 0, 151, 151, 102,
	0, 150, 0, 153, 0, 0, 0, 0, 0, 0,
	0, 140, 149, 152, 162, 170, 170, 186, 204, 216,
	141, 0, 161, 0, 0, 0, 0, 0, 0, 169,
	0, 0, 0, 0, 0, 0, 0, 0, 185, 187,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 203, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 215, 0, 0, 0, 221, 154, 160,
	163, 164, 165, 166, 167, 155, 168, 179, 172, 173,
	174, 175, 176, 156, 157, 184, 195, 189, 190, 191,
	192, 196, 197, 198, 199, 200, 201, 158, 202, 205,
	206, 207, 208, 209, 210, 211, 212, 213, 159, 214,
	217, 218, 219, 220, 0, 178, 0, 0, 0, 0,
	194, 171, 177, 180, 181, 182, 183, 188, 193,
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
		//line parser/parser.go.y:28
		{
			yyVAL.blocks = append(yyDollar[1].blocks, yyDollar[2].blocks...)
			yylex.(*Lexer).result = yyVAL.blocks
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:33
		{
			yyVAL.blocks = yyDollar[1].blocks
			yylex.(*Lexer).result = yyVAL.blocks
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:38
		{
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:39
		{
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:40
		{
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:41
		{
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:42
		{
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:43
		{
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:44
		{
			yyVAL.blocks = []Block{yyDollar[1].block}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:45
		{
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:46
		{
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:47
		{
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:54
		{
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:55
		{
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:56
		{
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:57
		{
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:58
		{
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:59
		{
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:60
		{
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:61
		{
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:62
		{
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:63
		{
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:64
		{
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:65
		{
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:66
		{
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:67
		{
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:68
		{
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:82
		{
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:83
		{
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:84
		{
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:85
		{
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:86
		{
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:87
		{
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:88
		{
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:90
		{
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:96
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[4].stmts}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:102
		{
			yyVAL.stmts = append([]Stmt{yyDollar[1].stmt}, yyDollar[2].stmts...)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:106
		{
			yyVAL.stmts = []Stmt{yyDollar[1].stmt}
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:110
		{
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:111
		{
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:112
		{
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:113
		{
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:114
		{
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:115
		{
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:116
		{
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:117
		{
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:118
		{
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:119
		{
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:120
		{
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:121
		{
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:122
		{
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:123
		{
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:124
		{
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:125
		{
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:127
		{
			yyVAL.stmt = Stmt{name: yyDollar[1].token.lit, val: yyDollar[3].value}
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:131
		{
			yyVAL.stmt = Stmt{name: yyDollar[1].token.lit, val: yyDollar[3].value}
		}
	case 71:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:134
		{
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:135
		{
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:136
		{
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:137
		{
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:138
		{
		}
	case 76:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:139
		{
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:140
		{
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:141
		{
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:142
		{
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:143
		{
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:144
		{
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:145
		{
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:146
		{
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:147
		{
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:148
		{
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:152
		{
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:153
		{
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:154
		{
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:155
		{
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:160
		{
		}
	case 100:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:166
		{
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:167
		{
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:168
		{
		}
	case 106:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:174
		{
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:175
		{
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:176
		{
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:177
		{
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:178
		{
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:179
		{
		}
	case 112:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:180
		{
		}
	case 120:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:190
		{
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:191
		{
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:192
		{
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:193
		{
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:201
		{
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:202
		{
		}
	case 131:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:204
		{
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:205
		{
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:206
		{
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:207
		{
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:208
		{
		}
	case 136:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:209
		{
		}
	case 137:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:210
		{
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:211
		{
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:212
		{
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:215
		{
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:218
		{
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:219
		{
		}
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:220
		{
		}
	case 148:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:221
		{
		}
	case 150:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:223
		{
		}
	case 151:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:225
		{
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:226
		{
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:227
		{
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:228
		{
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:229
		{
		}
	case 156:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:230
		{
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:231
		{
		}
	case 158:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:232
		{
		}
	case 159:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:233
		{
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:235
		{
		}
	case 162:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:237
		{
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:238
		{
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:239
		{
		}
	case 165:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:240
		{
		}
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:241
		{
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:242
		{
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:244
		{
		}
	case 170:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:246
		{
		}
	case 171:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:247
		{
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:248
		{
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:249
		{
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:250
		{
		}
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:251
		{
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:252
		{
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:254
		{
		}
	case 179:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:256
		{
		}
	case 180:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:257
		{
		}
	case 181:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:258
		{
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:259
		{
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:260
		{
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:262
		{
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:264
		{
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:265
		{
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:272
		{
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:274
		{
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:275
		{
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:276
		{
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:277
		{
		}
	case 199:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:278
		{
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:279
		{
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:280
		{
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:282
		{
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:284
		{
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:285
		{
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:286
		{
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:287
		{
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:288
		{
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:289
		{
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:290
		{
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:291
		{
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:292
		{
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:293
		{
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:295
		{
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:297
		{
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:298
		{
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:299
		{
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:300
		{
		}
	case 220:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:301
		{
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:302
		{
		}
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:304
		{
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:305
		{
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:306
		{
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:307
		{
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:308
		{
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:309
		{
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:310
		{
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:311
		{
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:312
		{
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:313
		{
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:314
		{
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:315
		{
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:316
		{
		}
	case 235:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:318
		{
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:319
		{
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:320
		{
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:321
		{
		}
	case 239:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:323
		{
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:324
		{
		}
	case 241:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:325
		{
		}
	case 242:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:329
		{
			// $$ = append($1, $2...)
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:332
		{
			yyVAL.value = yyDollar[1].value
		}
	case 244:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:336
		{
			//  $$ = append($1, $2...)
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:339
		{
			yyVAL.value = yyDollar[1].value
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:342
		{
			yyVAL.value = yyDollar[1].value
		}
	case 247:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:343
		{
		}
	case 248:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:344
		{
		}
	case 249:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:345
		{
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:347
		{
		}
	case 251:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:348
		{
		}
	case 252:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:349
		{
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:351
		{
		}
	case 255:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:353
		{
		}
	case 258:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:357
		{
		}
	case 266:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:366
		{
		}
	case 283:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:384
		{
		}
	case 288:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:390
		{
		}
	case 294:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:398
		{
			yyVAL.value = yyDollar[1].value
		}
	case 295:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:399
		{
			yyVAL.value = yyDollar[1].token.lit
		}
	case 296:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:402
		{
			fmt.Printf("%+v\n", yyDollar[1].token)
			yyVAL.value = yyDollar[1].token.lit
		}
	case 297:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:403
		{
			yyVAL.value = yyDollar[1].token.lit
		}
	case 298:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:405
		{
		}
	case 301:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:409
		{
		}
	}
	goto yystack /* stack new state and value */
}
