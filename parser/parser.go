//line parser/parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser/parser.go.y:2
//line parser/parser.go.y:5
type yySymType struct {
	yys           int
	string        string
	strings       []string
	token         Token
	block         Block
	blocks_any    []BlockAny
	block_args    BlockArgs
	subblock_args SubBlockArgs
	stmt_any      StmtAny
	stmts_any     []StmtAny
	stmt          Stmt
	stmt_multi    StmtMulti
	stmt_value    StmtValue
	values        []Value
	vip_addr      VIPAddr
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

//line parser/parser.go.y:472

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 538

var yyAct = [...]int{

	397, 477, 472, 62, 420, 409, 395, 386, 379, 360,
	330, 326, 207, 322, 181, 311, 309, 306, 182, 205,
	200, 265, 278, 30, 173, 142, 134, 53, 33, 34,
	63, 357, 388, 183, 286, 287, 389, 390, 217, 216,
	58, 422, 423, 424, 425, 328, 215, 393, 94, 391,
	392, 214, 382, 381, 444, 398, 53, 383, 59, 104,
	474, 475, 476, 384, 114, 115, 385, 485, 117, 118,
	119, 33, 34, 479, 203, 466, 126, 75, 76, 130,
	132, 324, 325, 68, 67, 71, 73, 455, 202, 442,
	416, 212, 60, 59, 61, 64, 65, 66, 69, 70,
	113, 408, 415, 412, 411, 413, 414, 417, 72, 418,
	419, 408, 362, 407, 404, 403, 405, 406, 184, 185,
	186, 187, 188, 407, 404, 403, 405, 406, 400, 402,
	33, 34, 329, 441, 399, 136, 106, 401, 189, 190,
	349, 350, 191, 192, 363, 364, 365, 366, 367, 368,
	369, 241, 242, 229, 230, 433, 32, 279, 280, 281,
	36, 129, 219, 137, 12, 426, 370, 355, 51, 226,
	352, 193, 194, 195, 196, 197, 198, 199, 13, 14,
	15, 16, 347, 17, 33, 34, 318, 204, 175, 176,
	177, 178, 179, 180, 345, 288, 289, 262, 86, 254,
	107, 108, 109, 110, 111, 202, 344, 339, 125, 282,
	212, 283, 138, 139, 140, 337, 141, 59, 33, 34,
	63, 103, 296, 335, 263, 208, 301, 209, 210, 33,
	34, 213, 300, 211, 295, 310, 435, 18, 261, 314,
	315, 253, 317, 317, 20, 19, 31, 316, 33, 34,
	54, 225, 218, 128, 112, 319, 102, 85, 378, 321,
	266, 267, 268, 269, 270, 271, 272, 273, 274, 275,
	276, 277, 245, 68, 67, 71, 377, 376, 375, 374,
	246, 373, 60, 88, 61, 64, 65, 66, 69, 70,
	481, 87, 468, 358, 247, 244, 240, 482, 72, 469,
	100, 239, 238, 231, 56, 57, 220, 84, 80, 79,
	78, 55, 77, 25, 310, 24, 317, 23, 22, 483,
	343, 33, 34, 133, 99, 336, 89, 338, 340, 212,
	33, 34, 131, 33, 34, 127, 346, 348, 351, 353,
	33, 34, 356, 38, 39, 44, 45, 46, 47, 48,
	49, 40, 41, 42, 43, 50, 33, 34, 116, 95,
	465, 33, 34, 234, 122, 233, 123, 464, 448, 332,
	342, 371, 341, 308, 294, 293, 290, 255, 252, 251,
	250, 249, 248, 394, 232, 228, 224, 223, 222, 427,
	221, 121, 120, 105, 98, 434, 97, 83, 35, 28,
	27, 26, 484, 443, 471, 470, 463, 449, 462, 451,
	461, 460, 458, 454, 52, 457, 456, 459, 453, 452,
	450, 447, 446, 445, 440, 439, 467, 438, 437, 436,
	432, 431, 430, 429, 428, 372, 359, 354, 334, 333,
	305, 304, 303, 302, 299, 478, 144, 145, 146, 147,
	148, 149, 150, 151, 152, 153, 154, 155, 156, 157,
	158, 159, 160, 161, 298, 297, 162, 163, 164, 292,
	291, 284, 264, 260, 259, 258, 480, 257, 256, 478,
	486, 243, 237, 236, 235, 227, 124, 101, 96, 93,
	165, 166, 167, 168, 169, 170, 171, 172, 92, 91,
	90, 82, 81, 1, 206, 421, 21, 410, 396, 473,
	387, 380, 331, 327, 323, 320, 307, 37, 285, 313,
	29, 312, 361, 74, 201, 174, 135, 143, 10, 11,
	9, 7, 6, 5, 4, 8, 3, 2,
}
var yyPact = [...]int{

	148, -1000, 148, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 304, 303, 301, 299, 395, 394, 393, 122,
	392, -1000, 326, 240, 210, 4, 298, 296, 295, 294,
	498, 497, 391, -1000, -1000, 293, 242, 326, 277, 319,
	496, 495, 494, 485, 353, 484, 390, 388, 316, 291,
	483, 241, 240, -1000, -1000, 332, 387, 134, 239, 210,
	332, 348, -1000, -1000, 332, 332, 332, 386, 385, 360,
	482, 134, 325, 238, 4, 322, 313, 129, 410, 98,
	20, -1000, -1000, -1000, 63, -1000, -1000, 221, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-27, -32, -39, -40, 237, 129, -1000, 292, 384, 382,
	381, 380, 236, 410, -1000, 481, -1000, -1000, 379, 332,
	332, 289, 378, 359, -1000, 480, 479, 478, 288, 287,
	282, 97, 477, -1000, 281, 266, 280, -1000, 376, 375,
	374, 373, 372, 226, 98, 371, 474, 473, 471, 470,
	469, 223, 20, 210, 468, 157, 42, 157, 42, 467,
	-86, 332, 332, 370, -1000, -1000, 466, 465, 369, 368,
	219, 63, 461, 460, 440, 217, 221, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 439, 438, 437, 436, -1000, -1000,
	367, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 332, -1000, -1000, -1000, -1000, -1000, -1000, 176, 176,
	20, -1000, -1000, -1000, 22, -1000, 39, 363, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 435, 434,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 208, 367, -1000, 200,
	332, 192, 176, -1000, 366, 364, 332, -1000, -1000, 191,
	179, 20, 167, 22, 79, 221, 155, 39, -1000, 433,
	152, 363, -62, -1000, 279, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 432, 19, -1000,
	151, 19, 431, -1000, 267, 265, 264, 263, 262, 244,
	-1000, -1000, -1000, -84, -100, -100, -13, -34, -104, 150,
	-84, 430, 429, 428, 427, 426, 140, -100, 222, 425,
	424, 423, 421, 420, 118, 74, -13, -1000, 40, 419,
	418, 417, 362, 332, 416, 332, 415, 414, 409, 72,
	-34, 332, 408, 332, 407, 406, 404, 402, 361, 354,
	60, -104, 286, 401, 400, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -73, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -23, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 58, -73, 284, 307, 398, 52, -23, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 503, 537, 536, 535, 534, 533, 532, 531, 530,
	529, 528, 160, 25, 26, 24, 168, 40, 86, 20,
	14, 9, 527, 526, 525, 524, 18, 33, 523, 522,
	15, 521, 520, 519, 3, 21, 22, 518, 517, 19,
	17, 516, 16, 515, 13, 11, 10, 514, 12, 513,
	512, 8, 7, 6, 5, 4, 511, 510, 2, 509,
	508, 0, 1, 507, 505, 504, 414, 136, 414,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 3, 12, 12, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 5, 6, 7, 18, 18, 8, 14, 14, 23,
	23, 23, 23, 23, 23, 23, 40, 40, 41, 4,
	13, 13, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 22, 22, 22, 44, 44, 47, 47, 47,
	47, 45, 45, 49, 49, 42, 42, 46, 46, 50,
	50, 50, 9, 15, 15, 24, 24, 24, 24, 24,
	24, 24, 43, 43, 26, 26, 10, 19, 19, 25,
	25, 25, 25, 11, 20, 20, 32, 32, 32, 26,
	26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, 26, 21, 21, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 51, 51, 56, 56,
	56, 56, 56, 56, 52, 52, 57, 57, 57, 57,
	57, 57, 57, 58, 58, 59, 59, 59, 59, 59,
	53, 53, 60, 60, 60, 60, 60, 60, 60, 62,
	62, 61, 61, 61, 61, 61, 61, 61, 54, 54,
	63, 63, 63, 63, 63, 63, 63, 63, 63, 63,
	55, 55, 64, 64, 64, 64, 64, 64, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	36, 36, 36, 37, 37, 30, 30, 31, 31, 31,
	31, 39, 39, 65, 16, 16, 66, 66, 66, 66,
	66, 66, 17, 17, 27, 27, 27, 27, 27, 27,
	27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
	27, 28, 28, 28, 28, 67, 67, 67, 67, 67,
	67, 33, 33, 34, 34, 68, 68, 68, 48, 48,
	48, 48, 48, 48, 48,
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
	2, 2, 2, 2, 2, 2, 3, 6, 2, 1,
	1, 2, 2, 2, 2, 2, 1, 0, 2, 1,
	4, 4, 4, 4, 4, 4, 2, 1, 0, 2,
	2, 2, 2, 2, 2, 1, 0, 4, 2, 2,
	2, 2, 2, 2, 1, 0, 2, 2, 2, 2,
	2, 1, 0, 1, 4, 2, 2, 2, 2, 2,
	1, 0, 2, 2, 2, 2, 2, 2, 2, 1,
	0, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 1, 0, 2, 2, 2, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 1, 1, 2, 2,
	2, 2, 1, 1, 2, 1, 0, 1, 1, 2,
	2, 2, 2, 1, 0, 2, 2, 2, 1, 1,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 4, 4, 4, 4, 0, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 0, 1, 2, 0, 1,
	1, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -2, -3, -5, -6, -7, -8, -4, -9,
	-11, -10, 16, 30, 31, 32, 33, 35, 89, 97,
	96, -1, 14, 14, 14, 14, 6, 6, 6, -32,
	-34, 124, 34, 8, 9, 6, -12, -38, 17, 18,
	25, 26, 27, 28, 19, 20, 21, 22, 23, 24,
	29, -16, -66, -34, 10, 71, 64, 65, -17, -27,
	72, 74, -34, 10, 75, 76, 77, 64, 63, 78,
	79, 65, 88, -18, -28, 73, 74, 14, 14, 14,
	14, 4, 4, 6, 14, 15, -12, 14, 6, 7,
	4, 4, 4, 4, -34, 6, 4, 6, 6, 8,
	9, 4, 15, -16, -34, 6, -67, 66, 67, 68,
	69, 70, 15, -17, -34, -34, 10, -34, -34, -34,
	6, 6, 4, 6, 4, -67, -34, 10, 15, -18,
	-34, 10, -34, 10, -14, -23, 6, 34, 83, 84,
	85, 87, -13, -22, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 56, 57, 58, 80, 81, 82, 83, 84,
	85, 86, 87, -15, -24, 90, 91, 92, 93, 94,
	95, -20, -26, -27, 98, 99, 100, 101, 102, 118,
	119, 122, 123, 151, 152, 153, 154, 155, 156, 157,
	-19, -25, -34, 11, 124, -39, -65, -48, 4, 6,
	7, 12, -34, 10, 78, 78, 78, 78, 15, -14,
	14, 6, 6, 6, 6, 15, -13, 4, 6, -34,
	-34, 14, 6, 6, 4, 4, 4, 4, 14, 14,
	14, 54, 55, 4, 14, 6, 14, 14, 6, 6,
	6, 6, 6, 15, -15, 6, 4, 4, 4, 4,
	4, 15, -20, -17, 4, -35, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114, -36, 115,
	116, 117, -35, -36, 4, -37, 120, 121, -34, -34,
	6, 4, 4, 6, 6, 15, -19, 4, 4, 4,
	15, -39, 4, 4, 4, 4, -40, -41, 6, -42,
	-34, -30, -31, -33, 63, 64, 71, -34, 10, -30,
	-43, -26, -44, -47, 59, 60, -45, -49, 6, 93,
	-46, -50, 6, 4, 4, 15, -40, 15, -42, 15,
	-30, 6, 6, -34, 15, 15, -20, 15, -44, 61,
	62, -48, 15, -45, 4, 15, -46, 93, 14, 4,
	-21, -29, 93, 125, 126, 127, 128, 129, 130, 131,
	15, -21, 4, 14, 14, 14, 14, 14, 14, -51,
	-56, 137, 136, 141, 147, 150, -52, -57, 132, 136,
	137, 149, 150, 147, -52, -53, -60, -61, 68, 147,
	141, 150, 142, 138, 137, 139, 140, 136, 124, -54,
	-63, 138, 137, 139, 140, 136, 124, 141, 143, 144,
	-55, -64, 145, 146, 147, 148, 15, -51, 4, 4,
	4, 4, 4, 15, -52, 14, 4, 4, 4, 4,
	4, 15, 15, -53, 14, 4, 4, 4, 6, -34,
	4, -34, 4, 4, 4, 15, -54, -34, 4, -34,
	4, 4, 4, 4, 6, 6, 15, -55, 6, 13,
	4, 4, -58, -59, 133, 134, 135, -62, -61, 15,
	-58, 6, 13, 12, 4, 15, -62,
}
var yyDef = [...]int{

	0, -2, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 1, 15, 246, 254, 0, 0, 0, 0, 0,
	0, 0, 0, 283, 284, 0, 0, 14, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 245, 247, 248, 0, 0, 275, 0, 253,
	0, 0, 258, 259, 0, 0, 0, 0, 0, 0,
	0, 275, 0, 0, 35, 0, 0, 39, 52, 105,
	254, 126, 127, 128, 119, 12, 13, 288, 17, 18,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	29, 30, 31, 244, 249, 250, 251, 276, 277, 278,
	279, 280, 32, 252, 255, 256, 257, 260, 261, 262,
	263, 264, 265, 266, 267, 268, 269, 270, 33, 34,
	0, 0, 0, 0, 0, 38, 40, 0, 0, 0,
	0, 0, 0, 51, 53, 0, 55, 56, 0, 0,
	0, 0, 0, 0, 64, 0, 0, 0, 0, 0,
	0, 0, 0, 74, 0, 0, 0, 79, 0, 0,
	0, 0, 0, 0, 104, 0, 0, 0, 0, 0,
	0, 0, 125, 115, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 139, 140, 0, 0, 0, 0,
	0, 118, 0, 0, 0, 0, 242, 243, 289, 290,
	291, 292, 293, 294, 0, 0, 0, 0, 36, 37,
	0, 42, 43, 44, 45, 49, 50, 54, 57, 58,
	59, 0, 61, 62, 63, 65, 66, 67, 0, 0,
	254, 71, 72, 73, 87, 76, 0, 99, 80, 81,
	82, 83, 84, 102, 103, 106, 107, 108, 109, 110,
	111, 123, 124, 114, 129, 130, 218, 219, 220, 221,
	222, 223, 224, 225, 226, 227, 228, 229, 131, 230,
	231, 232, 132, 133, 134, 135, 233, 234, 0, 0,
	138, 141, 142, 143, 144, 116, 117, 120, 121, 122,
	16, 241, 271, 273, 272, 274, 0, 47, 48, 0,
	96, 0, 236, 237, 0, 0, 0, 281, 282, 0,
	0, 113, 0, 86, 0, 288, 0, 92, 93, 0,
	0, 98, 100, 136, 0, 41, 46, 60, 95, 68,
	235, 238, 239, 240, 69, 70, 112, 75, 85, 88,
	89, 90, 77, 91, 94, 78, 97, 0, 147, 101,
	0, 146, 0, 149, 0, 0, 0, 0, 0, 0,
	137, 145, 148, 158, 166, 166, 182, 200, 212, 0,
	157, 0, 0, 0, 0, 0, 0, 165, 0, 0,
	0, 0, 0, 0, 0, 0, 181, 183, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	199, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 211, 0, 0, 0, 217, 150, 156, 159, 160,
	161, 162, 163, 151, 164, 175, 168, 169, 170, 171,
	172, 152, 153, 180, 191, 185, 186, 187, 188, 192,
	193, 194, 195, 196, 197, 154, 198, 201, 202, 203,
	204, 205, 206, 207, 208, 209, 155, 210, 213, 214,
	215, 216, 0, 174, 0, 0, 0, 0, 190, 167,
	173, 176, 177, 178, 179, 184, 189,
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
		//line parser/parser.go.y:37
		{
			yyVAL.blocks_any = append(yyDollar[1].blocks_any, yyDollar[2].blocks_any...)
			yylex.(*Lexer).result = yyVAL.blocks_any
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:42
		{
			yyVAL.blocks_any = yyDollar[1].blocks_any
			yylex.(*Lexer).result = yyVAL.blocks_any
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:48
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:49
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:50
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:51
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:52
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:53
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:54
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:55
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block_args}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:56
		{
			yyVAL.blocks_any = []BlockAny{yyDollar[1].block}
		}
	case 12:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:60
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[3].stmts_any}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:64
		{
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:64
		{
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:67
		{
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:68
		{
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:69
		{
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:70
		{
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:71
		{
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:72
		{
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:73
		{
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:74
		{
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:75
		{
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:76
		{
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:77
		{
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:78
		{
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:79
		{
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:80
		{
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:81
		{
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:85
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[3].stmts_any}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:91
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[3].stmts_any}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:97
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[3].stmts_any}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:102
		{
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:102
		{
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:106
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[4].stmts_any}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:112
		{
			yyVAL.stmts_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:115
		{
			yyVAL.stmts_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 39:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:117
		{
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:118
		{
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:119
		{
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:120
		{
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:121
		{
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:122
		{
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:123
		{
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:125
		{
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:131
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[4].stmts_any}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:137
		{
			yyVAL.stmts_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:140
		{
			yyVAL.stmts_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 52:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:142
		{
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:143
		{
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:144
		{
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:145
		{
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:146
		{
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:147
		{
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:148
		{
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:149
		{
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:150
		{
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:151
		{
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:152
		{
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:153
		{
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:154
		{
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:155
		{
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:156
		{
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:157
		{
		}
	case 68:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:159
		{
			yyVAL.stmt_any = StmtMulti{yyDollar[1].token.lit, yyDollar[3].values}
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:163
		{
			yyVAL.stmt_any = StmtMulti{yyDollar[1].token.lit, yyDollar[3].values}
		}
	case 70:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:166
		{
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:167
		{
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:168
		{
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:169
		{
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:170
		{
		}
	case 75:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:171
		{
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:172
		{
		}
	case 77:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:173
		{
		}
	case 78:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:174
		{
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:175
		{
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:176
		{
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:177
		{
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:178
		{
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:179
		{
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:180
		{
		}
	case 87:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:184
		{
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:185
		{
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:186
		{
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:187
		{
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:192
		{
		}
	case 99:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:198
		{
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:199
		{
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:200
		{
		}
	case 102:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:204
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[4].stmts_any}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:210
		{
			yyVAL.stmts_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:213
		{
			yyVAL.stmts_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 105:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:215
		{
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:216
		{
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:217
		{
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:218
		{
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:219
		{
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:220
		{
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:221
		{
		}
	case 114:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:227
		{
			yyVAL.stmt_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:230
		{
			yyVAL.stmt_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 116:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:234
		{
			yyVAL.block = Block{name: yyDollar[1].token.lit, stmts: yyDollar[4].stmts_any}
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:240
		{
			yyVAL.stmts_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:243
		{
			yyVAL.stmts_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 119:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:245
		{
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:246
		{
		}
	case 121:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:247
		{
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:248
		{
		}
	case 123:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser/parser.go.y:252
		{
			yyVAL.block_args = BlockArgs{Name: yyDollar[1].token.lit, Args: yyDollar[2].strings, Stmts: yyDollar[4].stmts_any}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:258
		{
			yyVAL.stmts_any = append([]StmtAny{yyDollar[1].stmt_any}, yyDollar[2].stmts_any...)
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:261
		{
			yyVAL.stmts_any = []StmtAny{yyDollar[1].stmt_any}
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:264
		{
			yyVAL.strings = []string{yyDollar[1].string, yyDollar[2].token.lit}
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:265
		{
			yyVAL.strings = []string{yyDollar[1].token.lit, yyDollar[2].token.lit}
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:266
		{
			yyVAL.strings = []string{yyDollar[1].token.lit, yyDollar[2].token.lit}
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:269
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].token.lit}
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:270
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].string}
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:271
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].string}
		}
	case 132:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:272
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].string}
		}
	case 133:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:273
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].string}
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:274
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].token.lit}
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:275
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].string}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser/parser.go.y:276
		{
			yyVAL.stmt_any = StmtMulti{yyDollar[1].token.lit, []Value{yyDollar[2].string, yyDollar[3].token.lit}}
		}
	case 137:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser/parser.go.y:278
		{
			yyVAL.stmt_any = SubBlockArgs{Name: yyDollar[1].token.lit, Args: []string{yyDollar[2].string, yyDollar[3].token.lit}, Stmts: yyDollar[5].stmts_any}
		}
	case 138:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:281
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].token.lit}
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:282
		{
			yyVAL.stmt_any = yyDollar[1].token.lit
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:283
		{
			yyVAL.stmt_any = yyDollar[1].token.lit
		}
	case 141:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:284
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].token.lit}
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:285
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].token.lit}
		}
	case 143:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:286
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].token.lit}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:287
		{
			yyVAL.stmt_any = Stmt{yyDollar[1].token.lit, yyDollar[2].token.lit}
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:289
		{
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:289
		{
		}
	case 147:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:291
		{
		}
	case 148:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:292
		{
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:293
		{
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:294
		{
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:295
		{
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:296
		{
		}
	case 153:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:297
		{
		}
	case 154:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:298
		{
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:299
		{
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:301
		{
		}
	case 158:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:303
		{
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:304
		{
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:305
		{
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:306
		{
		}
	case 162:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:307
		{
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:308
		{
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:310
		{
		}
	case 166:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:312
		{
		}
	case 167:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:313
		{
		}
	case 168:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:314
		{
		}
	case 169:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:315
		{
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:316
		{
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:317
		{
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:318
		{
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:320
		{
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:322
		{
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:323
		{
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:324
		{
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:325
		{
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:326
		{
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:328
		{
		}
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:330
		{
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:331
		{
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:338
		{
		}
	case 191:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:340
		{
		}
	case 192:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:341
		{
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:342
		{
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:343
		{
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:344
		{
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:345
		{
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:346
		{
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:348
		{
		}
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:350
		{
		}
	case 201:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:351
		{
		}
	case 202:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:352
		{
		}
	case 203:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:353
		{
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:354
		{
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:355
		{
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:356
		{
		}
	case 207:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:357
		{
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:358
		{
		}
	case 209:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:359
		{
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:361
		{
		}
	case 212:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:363
		{
		}
	case 213:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:364
		{
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:365
		{
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:366
		{
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:367
		{
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:368
		{
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:371
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:372
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:373
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:374
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:375
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:376
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:377
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:378
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:379
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:380
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:381
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:382
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:385
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:386
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:387
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:390
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:391
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 235:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:395
		{
			yyVAL.values = append([]Value{yyDollar[1].vip_addr}, yyDollar[2].values...)
		}
	case 236:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:398
		{
			yyVAL.values = []Value{yyDollar[1].vip_addr}
		}
	case 237:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:401
		{
			yyVAL.vip_addr = VIPAddr{Addr: yyDollar[1].string}
		}
	case 238:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:402
		{
		}
	case 239:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:403
		{
		}
	case 240:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:404
		{
		}
	case 242:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:406
		{
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:408
		{
		}
	case 244:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:410
		{
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:410
		{
		}
	case 246:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:412
		{
		}
	case 252:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:419
		{
		}
	case 253:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:419
		{
		}
	case 254:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:421
		{
		}
	case 255:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:422
		{
		}
	case 256:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:423
		{
		}
	case 257:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:424
		{
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:425
		{
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:426
		{
		}
	case 260:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:427
		{
		}
	case 261:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:428
		{
		}
	case 262:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:429
		{
		}
	case 263:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:430
		{
		}
	case 264:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:431
		{
		}
	case 265:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:432
		{
		}
	case 266:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:433
		{
		}
	case 267:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:434
		{
		}
	case 268:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:435
		{
		}
	case 269:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:436
		{
		}
	case 270:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser/parser.go.y:437
		{
		}
	case 271:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:440
		{
		}
	case 272:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:441
		{
		}
	case 273:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:442
		{
		}
	case 274:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser/parser.go.y:443
		{
		}
	case 275:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:445
		{
		}
	case 281:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:453
		{
			yyVAL.string = yyDollar[1].string
		}
	case 282:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:454
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 283:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:457
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 284:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser/parser.go.y:458
		{
			yyVAL.string = yyDollar[1].token.lit
		}
	case 285:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:460
		{
		}
	case 288:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser/parser.go.y:464
		{
		}
	}
	goto yystack /* stack new state and value */
}
