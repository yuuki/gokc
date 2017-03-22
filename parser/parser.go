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
const FO = 57448
const OVF = 57449
const LBLC = 57450
const LBLCR = 57451
const SH = 57452
const DH = 57453
const SED = 57454
const NQ = 57455
const NAT = 57456
const DR = 57457
const TUN = 57458
const PERSISTENCE_TIMEOUT = 57459
const PROTOCOL = 57460
const TCP = 57461
const UDP = 57462
const SORRY_SERVER = 57463
const REAL_SERVER = 57464
const FWMARK = 57465
const INHIBIT_ON_FAILURE = 57466
const TCP_CHECK = 57467
const HTTP_GET = 57468
const SSL_GET = 57469
const SMTP_CHECK = 57470
const DNS_CHECK = 57471
const MISC_CHECK = 57472
const URL = 57473
const PATH = 57474
const DIGEST = 57475
const STATUS_CODE = 57476
const CONNECT_TIMEOUT = 57477
const CONNECT_PORT = 57478
const CONNECT_IP = 57479
const BINDTO = 57480
const BIND_PORT = 57481
const RETRY = 57482
const HELO_NAME = 57483
const TYPE = 57484
const NAME = 57485
const MISC_PATH = 57486
const MISC_TIMEOUT = 57487
const WARMUP = 57488
const MISC_DYNAMIC = 57489
const NB_GET_RETRY = 57490
const DELAY_BEFORE_RETRY = 57491
const VIRTUALHOST = 57492
const ALPHA = 57493
const OMEGA = 57494
const QUORUM = 57495
const HYSTERESIS = 57496
const QUORUM_UP = 57497
const QUORUM_DOWN = 57498

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

//line parser/parser.go.y:380

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 19,
	14, 126,
	-2, 296,
	-1, 201,
	15, 118,
	-2, 119,
}

const yyPrivate = 57344

const yyLast = 556

var yyAct = [...]int{

	407, 487, 482, 63, 430, 419, 405, 396, 389, 368,
	333, 329, 207, 325, 181, 312, 318, 310, 308, 30,
	305, 182, 205, 33, 200, 277, 173, 54, 264, 183,
	34, 35, 64, 59, 143, 398, 364, 392, 391, 399,
	400, 217, 393, 432, 433, 434, 435, 135, 394, 95,
	403, 395, 401, 402, 60, 331, 408, 54, 285, 286,
	105, 484, 485, 486, 216, 115, 116, 215, 214, 118,
	119, 120, 34, 35, 454, 203, 495, 127, 76, 77,
	131, 133, 489, 74, 69, 68, 72, 476, 107, 202,
	60, 426, 212, 61, 114, 62, 65, 66, 67, 70,
	71, 418, 465, 425, 422, 421, 423, 424, 427, 73,
	428, 429, 418, 417, 414, 413, 415, 416, 452, 184,
	185, 186, 187, 188, 417, 414, 413, 415, 416, 410,
	412, 370, 34, 35, 451, 409, 356, 357, 411, 189,
	190, 332, 37, 191, 192, 278, 279, 280, 175, 176,
	177, 178, 179, 180, 229, 230, 327, 328, 32, 130,
	137, 126, 443, 371, 372, 373, 374, 375, 376, 377,
	240, 241, 193, 194, 195, 196, 197, 198, 199, 226,
	52, 87, 436, 388, 219, 379, 362, 204, 138, 34,
	35, 317, 359, 354, 352, 287, 288, 261, 34, 35,
	317, 253, 348, 445, 12, 202, 343, 341, 339, 299,
	212, 294, 260, 60, 282, 252, 281, 262, 13, 14,
	15, 16, 244, 17, 225, 218, 295, 493, 129, 300,
	245, 34, 35, 64, 104, 309, 139, 140, 141, 113,
	142, 316, 316, 313, 314, 103, 86, 31, 491, 387,
	386, 315, 478, 321, 320, 492, 34, 35, 55, 479,
	322, 324, 265, 266, 267, 268, 269, 270, 271, 272,
	273, 274, 275, 276, 89, 385, 18, 108, 109, 110,
	111, 112, 88, 20, 19, 69, 68, 72, 384, 383,
	382, 366, 33, 365, 61, 246, 62, 65, 66, 67,
	70, 71, 243, 239, 238, 237, 231, 220, 337, 85,
	73, 57, 58, 309, 81, 316, 80, 79, 56, 347,
	78, 25, 24, 316, 23, 22, 351, 340, 342, 344,
	101, 96, 212, 34, 35, 320, 349, 475, 100, 353,
	355, 358, 360, 34, 35, 363, 39, 40, 45, 46,
	47, 48, 49, 50, 41, 42, 43, 44, 51, 208,
	90, 209, 210, 34, 35, 213, 494, 211, 34, 35,
	134, 34, 35, 132, 474, 234, 378, 233, 481, 380,
	34, 35, 128, 34, 35, 117, 338, 123, 480, 124,
	34, 35, 404, 458, 335, 350, 346, 345, 307, 437,
	293, 292, 289, 254, 251, 444, 250, 249, 248, 247,
	232, 228, 224, 453, 223, 222, 221, 459, 122, 461,
	121, 106, 99, 98, 83, 467, 466, 469, 36, 28,
	27, 26, 473, 472, 471, 470, 477, 468, 464, 463,
	462, 460, 457, 456, 455, 450, 449, 448, 447, 446,
	442, 441, 440, 439, 438, 488, 145, 146, 147, 148,
	149, 150, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 160, 161, 381, 367, 162, 163, 164, 361, 336,
	304, 303, 302, 301, 298, 297, 490, 296, 291, 488,
	496, 290, 283, 263, 259, 258, 257, 256, 255, 165,
	166, 167, 168, 169, 170, 171, 172, 242, 236, 235,
	227, 125, 102, 97, 94, 93, 92, 91, 84, 82,
	1, 53, 206, 21, 319, 311, 431, 420, 406, 483,
	397, 390, 369, 284, 29, 201, 174, 334, 330, 326,
	323, 144, 306, 136, 75, 38, 11, 10, 9, 8,
	7, 6, 5, 4, 3, 2,
}
var yyPact = [...]int{

	188, -1000, 188, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 311, 310, 308, 307, 425, 424, 423, 124,
	422, -1000, 329, 248, 223, 6, 306, 303, 302, 300,
	-1000, 515, 418, 514, -1000, -1000, 295, 231, 329, 268,
	353, 513, 512, 511, 510, 325, 509, 417, 416, 330,
	321, 508, 230, 248, -1000, -1000, 335, 415, 212, 224,
	223, 335, 375, -1000, -1000, 335, 335, 335, 414, 412,
	383, 507, 212, 372, 213, 6, 363, 360, 154, 420,
	59, 22, -1000, -1000, -1000, 64, -1000, -1000, 355, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -9, -10, -13, -36, 210, 154, -1000, 293, 410,
	409, 408, 406, 209, 420, -1000, 506, -1000, -1000, 405,
	335, 335, 292, 404, 371, -1000, 505, 504, 291, 290,
	289, 117, 503, -1000, 288, 216, 281, -1000, 403, 402,
	401, 400, 398, 200, 59, 397, 494, 493, 492, 491,
	490, 197, 22, 223, 489, 160, 31, 160, 31, 488,
	-61, 335, 335, 396, -1000, -1000, 487, 484, 395, 394,
	196, 64, 483, 481, 480, 194, 355, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 479, 478, 477, 476, -1000, -1000,
	392, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 335, -1000, -1000, -1000, -1000, -1000, 181, 190, 22,
	-1000, -1000, -1000, 98, -1000, 49, 388, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 475, 382, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 193, 392, -1000, 192, 335,
	191, 181, -1000, 391, 390, 335, -1000, -1000, 187, 190,
	-1000, 389, 335, 179, 22, 178, 98, 76, 355, 177,
	49, -1000, 474, 171, 388, -56, -1000, 279, 277, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 470, 39, 39, -1000, 170, 39,
	469, -1000, 276, 275, 274, 261, 236, 235, 168, -1000,
	-1000, -1000, -98, -96, -96, -11, -32, -101, -1000, 167,
	-98, 450, 449, 448, 447, 446, 147, -96, 189, 445,
	444, 443, 442, 441, 119, 103, -11, -1000, 60, 440,
	439, 438, 387, 335, 437, 335, 436, 435, 434, 87,
	-32, 335, 433, 335, 431, 430, 429, 428, 368, 331,
	72, -101, 246, 384, 374, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -71, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -22, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 67, -71, 242, 215, 362, 61, -22, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 520, 555, 554, 553, 552, 551, 550, 549, 548,
	547, 546, 142, 545, 22, 3, 180, 33, 83, 544,
	47, 543, 20, 542, 34, 541, 18, 17, 16, 540,
	13, 11, 10, 539, 12, 538, 537, 26, 536, 21,
	14, 29, 24, 535, 534, 19, 28, 25, 533, 9,
	532, 8, 7, 6, 5, 4, 531, 530, 2, 529,
	528, 0, 1, 527, 526, 525, 524, 15, 522, 521,
	88,
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
	50, 50, 50, 50, 50, 50, 50, 50, 50, 51,
	51, 56, 56, 56, 56, 56, 56, 52, 52, 57,
	57, 57, 57, 57, 57, 57, 58, 58, 59, 59,
	59, 59, 59, 53, 53, 60, 60, 60, 60, 60,
	60, 60, 62, 62, 61, 61, 61, 61, 61, 61,
	61, 54, 54, 63, 63, 63, 63, 63, 63, 63,
	63, 63, 63, 55, 55, 64, 64, 64, 64, 64,
	64, 46, 46, 46, 46, 46, 46, 46, 46, 46,
	46, 46, 46, 46, 47, 47, 47, 47, 48, 48,
	48, 27, 27, 28, 28, 65, 65, 65, 65, 66,
	66, 66, 14, 14, 68, 16, 16, 69, 69, 69,
	69, 69, 69, 17, 17, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 19, 19, 19, 19, 19, 70, 70, 70,
	70, 70, 70, 67, 67, 67, 15, 15, 15, 45,
	45, 45, 34, 34, 34, 34, 34, 34, 34,
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
	0, 2, 1, 4, 4, 4, 4, 4, 4, 2,
	1, 0, 2, 2, 2, 2, 2, 2, 1, 0,
	4, 2, 2, 2, 2, 2, 2, 1, 0, 2,
	2, 2, 2, 2, 1, 0, 1, 4, 2, 2,
	2, 2, 2, 1, 0, 2, 2, 2, 2, 2,
	2, 2, 1, 0, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 1, 0, 2, 2, 2, 2,
	1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 0, 1, 1, 1, 0, 1,
	1, 2, 1, 2, 1, 1, 2, 2, 2, 1,
	2, 2, 2, 1, 1, 2, 1, 0, 1, 1,
	2, 2, 2, 2, 1, 0, 2, 2, 2, 1,
	1, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 0, 4, 4, 4, 4, 0, 1, 1,
	1, 1, 1, 0, 1, 1, 0, 1, 1, 0,
	1, 2, 0, 1, 1, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, 16, 30, 31, 32, 33, 35, 88, 96,
	95, -1, 14, 14, 14, 14, 6, 6, 6, -44,
	-45, 123, 34, -15, 8, 9, 6, -12, -13, 17,
	18, 25, 26, 27, 28, 19, 20, 21, 22, 23,
	24, 29, -16, -69, -15, 10, 70, 63, 64, -17,
	-41, 71, 73, -15, 10, 74, 75, 76, 63, 62,
	77, 78, 64, 87, -18, -19, 72, 73, 14, 14,
	14, 14, 4, 6, 4, 14, 15, -12, 14, 6,
	7, 4, 4, 4, 4, -15, 6, 4, 6, 6,
	8, 9, 4, 15, -16, -15, 6, -70, 65, 66,
	67, 68, 69, 15, -17, -15, -15, 10, -15, -15,
	-15, 6, 6, 4, 6, 4, -70, -15, 10, 15,
	-18, -15, 10, -15, 10, -20, -21, 6, 34, 82,
	83, 84, 86, -24, -25, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 55, 56, 57, 79, 80, 81, 82, 83,
	84, 85, 86, -37, -38, 89, 90, 91, 92, 93,
	94, -40, -39, -41, 97, 98, 99, 100, 101, 117,
	118, 121, 122, 150, 151, 152, 153, 154, 155, 156,
	-42, -43, -15, 11, 123, -14, -68, -34, 4, 6,
	7, 12, -15, 10, 77, 77, 77, 77, 15, -20,
	14, 6, 6, 6, 6, 15, -24, 4, 6, -15,
	-15, 14, 6, 6, 4, 4, 4, 14, 14, 14,
	53, 54, 4, 14, 6, 14, 14, 6, 6, 6,
	6, 6, 15, -37, 6, 4, 4, 4, 4, 4,
	15, -40, -17, 4, -46, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, -47, 114, 115,
	116, -46, -47, 4, -48, 119, 120, -15, -15, 6,
	4, 4, 6, 6, 15, -42, 4, 4, 4, 15,
	-14, 4, 4, 4, 4, -22, -23, 6, -26, -15,
	-27, -65, -67, 62, 63, 70, -15, 10, -28, -66,
	-67, 63, 70, -29, -39, -30, -33, 58, 59, -31,
	-35, 6, 92, -32, -36, 6, 4, -45, 4, 15,
	-22, 15, -26, 15, -27, 6, 6, -15, 15, -28,
	6, -15, 15, -40, 15, -30, 60, 61, -34, 15,
	-31, 4, 15, -32, 92, 14, 14, 4, -49, -50,
	92, 124, 125, 126, 127, 128, 129, 130, -49, 15,
	-49, 4, 14, 14, 14, 14, 14, 14, 15, -51,
	-56, 136, 135, 140, 146, 149, -52, -57, 131, 135,
	136, 148, 149, 146, -52, -53, -60, -61, 67, 146,
	140, 149, 141, 137, 136, 138, 139, 135, 123, -54,
	-63, 137, 136, 138, 139, 135, 123, 140, 142, 143,
	-55, -64, 144, 145, 146, 147, 15, -51, 4, 4,
	4, 4, 4, 15, -52, 14, 4, 4, 4, 4,
	4, 15, 15, -53, 14, 4, 4, 4, 6, -15,
	4, -15, 4, 4, 4, 15, -54, -15, 4, -15,
	4, 4, 4, 4, 6, 6, 15, -55, 6, 13,
	4, 4, -58, -59, 132, 133, 134, -62, -61, 15,
	-58, 6, 13, 12, 4, 15, -62,
}
var yyDef = [...]int{

	3, -2, 2, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 0, 0, 0, 0, 0, 0, 0, -2,
	0, 1, 16, 257, 265, 282, 0, 0, 0, 0,
	127, 0, 0, 300, 297, 298, 0, 0, 15, 0,
	0, 0, 0, 0, 0, 296, 0, 0, 0, 0,
	0, 0, 0, 256, 258, 259, 296, 0, 287, 0,
	264, 296, 296, 269, 270, 296, 296, 296, 0, 0,
	0, 0, 287, 296, 0, 36, 296, 296, 40, 53,
	105, 130, 128, 129, 301, 119, 13, 14, 296, 18,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 32, 255, 260, 261, 262, 288, 289,
	290, 291, 292, 33, 263, 266, 267, 268, 271, 272,
	273, 274, 275, 276, 277, 278, 279, 280, 281, 34,
	35, 0, 0, 0, 0, 0, 39, 41, 0, 0,
	0, 0, 0, 0, 52, 54, 0, 56, 57, 0,
	296, 296, 0, 0, 0, 65, 0, 0, 0, 0,
	0, 0, 0, 74, 0, 0, 0, 79, 0, 0,
	0, 0, 0, 0, 104, 0, 0, 0, 0, 0,
	0, 0, 125, 115, 0, 221, 234, 221, 234, 0,
	238, 296, 296, 0, 142, 143, 0, 0, 0, 0,
	0, -2, 0, 0, 0, 0, 253, 254, 303, 304,
	305, 306, 307, 308, 0, 0, 0, 0, 37, 38,
	0, 43, 44, 45, 46, 50, 51, 55, 58, 59,
	60, 296, 62, 63, 64, 66, 67, 293, 293, 130,
	71, 72, 73, 87, 76, 0, 99, 80, 81, 82,
	83, 84, 102, 103, 106, 107, 108, 109, 110, 111,
	123, 124, 114, 131, 132, 222, 223, 224, 225, 226,
	227, 228, 229, 230, 231, 232, 233, 133, 235, 236,
	237, 134, 135, 136, 137, 239, 240, 0, 296, 141,
	144, 145, 146, 147, 116, 117, 120, 121, 122, 17,
	252, 283, 285, 284, 286, 0, 48, 49, 0, 96,
	0, 242, 245, 0, 0, 296, 294, 295, 0, 244,
	249, 0, 296, 0, 113, 0, 86, 0, 296, 0,
	92, 93, 0, 0, 98, 100, 138, 0, 0, 42,
	47, 61, 95, 68, 241, 246, 247, 248, 69, 243,
	250, 251, 70, 112, 75, 85, 88, 89, 90, 77,
	91, 94, 78, 97, 0, 150, 150, 101, 0, 149,
	0, 152, 0, 0, 0, 0, 0, 0, 0, 139,
	148, 151, 161, 169, 169, 185, 203, 215, 140, 0,
	160, 0, 0, 0, 0, 0, 0, 168, 0, 0,
	0, 0, 0, 0, 0, 0, 184, 186, 0, 0,
	0, 0, 0, 296, 0, 296, 0, 0, 0, 0,
	202, 296, 0, 296, 0, 0, 0, 0, 0, 0,
	0, 214, 0, 0, 0, 220, 153, 159, 162, 163,
	164, 165, 166, 154, 167, 178, 171, 172, 173, 174,
	175, 155, 156, 183, 194, 188, 189, 190, 191, 195,
	196, 197, 198, 199, 200, 157, 201, 204, 205, 206,
	207, 208, 209, 210, 211, 212, 158, 213, 216, 217,
	218, 219, 0, 177, 0, 0, 0, 0, 193, 170,
	176, 179, 180, 181, 182, 187, 192,
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
	152, 153, 154, 155, 156,
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
	case 158:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:207
		{
		}
	case 160:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:209
		{
		}
	case 161:
		yyDollar = yyS[yypt-0 : yypt+1]
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
	case 166:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:216
		{
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:218
		{
		}
	case 169:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:220
		{
		}
	case 170:
		yyDollar = yyS[yypt-4 : yypt+1]
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
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:226
		{
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:228
		{
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
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
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:234
		{
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:236
		{
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:238
		{
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:239
		{
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:246
		{
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
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
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:254
		{
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:256
		{
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:262
		{
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:263
		{
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:264
		{
		}
	case 210:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:265
		{
		}
	case 211:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:266
		{
		}
	case 212:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:267
		{
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:269
		{
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:271
		{
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:272
		{
		}
	case 217:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:273
		{
		}
	case 218:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:274
		{
		}
	case 219:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:275
		{
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:276
		{
		}
	case 221:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:278
		{
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:279
		{
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:280
		{
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:281
		{
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:282
		{
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:283
		{
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:284
		{
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:285
		{
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:286
		{
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:287
		{
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:288
		{
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:289
		{
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:290
		{
		}
	case 234:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:292
		{
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:293
		{
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:294
		{
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:295
		{
		}
	case 238:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:297
		{
		}
	case 239:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:298
		{
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:299
		{
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:301
		{
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:303
		{
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:305
		{
		}
	case 246:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:306
		{
		}
	case 247:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:307
		{
		}
	case 248:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:308
		{
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:310
		{
		}
	case 250:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:311
		{
		}
	case 251:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:312
		{
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:314
		{
		}
	case 254:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:316
		{
		}
	case 257:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:320
		{
		}
	case 265:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:329
		{
		}
	case 282:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:347
		{
		}
	case 287:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:353
		{
		}
	case 293:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:360
		{
		}
	case 296:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:364
		{
		}
	case 299:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:368
		{
		}
	case 302:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:372
		{
		}
	}
	goto yystack /* stack new state and value */
}
