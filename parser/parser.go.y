%{
package parser
%}

%union {
  string  string
  strings []string
  token   Token
  block   Block
  blocks_any  []BlockAny
  block_args  BlockArgs
  subblock_args SubBlockArgs
  stmt_any    StmtAny
  stmts_any   []StmtAny
  stmt       Stmt
  stmt_multi StmtMulti
  stmt_value StmtValue
  values  []Value
  vip_addr VIPAddr 
};

%token<token> NUMBER ID STRING EMAIL IPV4 IPV6 IP_CIDR IPADDR_RANGE HEX32 PATHSTR LB RB GLOBALDEFS NOTIFICATION_EMAIL NOTIFICATION_EMAIL_FROM SMTP_SERVER SMTP_CONNECT_TIMEOUT ROUTER_ID LVS_ID VRRP_MCAST_GROUP4 VRRP_MCAST_GROUP6 VRRP_GARP_MASTER_DELAY VRRP_GARP_MASTER_REPEAT VRRP_GARP_MASTER_REFRESH VRRP_GARP_MASTER_REFRESH_REPEAT VRRP_VERSION STATIC_IPADDRESS STATIC_ROUTES STATIC_RULES VRRP_SYNC_GROUP GROUP VRRP_INSTANCE USE_VMAC VERSION VMAC_XMIT_BASE NATIVE_IPV6 INTERFACE MCAST_SRC_IP UNICAST_SRC_IP UNICAST_PEER LVS_SYNC_DAEMON_INTERFACE VIRTUAL_ROUTER_ID NOPREEMPT PREEMPT_DELAY PRIORITY ADVERT_INT VIRTUAL_IPADDRESS VIRTUAL_IPADDRESS_EXCLUDED VIRTUAL_ROUTES STATE MASTER BACKUP GARP_MASTER_DELAY SMTP_ALERT AUTHENTICATION AUTH_TYPE AUTH_PASS PASS AH LABEL DEV SCOPE SITE LINK HOST NOWHERE GLOBAL BRD SRC FROM TO VIA GW OR TABLE METRIC TRACK_INTERFACE TRACK_SCRIPT DONT_TRACK_PRIMARY NOTIFY_MASTER NOTIFY_BACKUP NOTIFY_FAULT NOTIFY_STOP NOTIFY BLACKHOLE VRRP_SCRIPT SCRIPT INTERVAL TIMEOUT WEIGHT FALL RISE VIRTUAL_SERVER_GROUP VIRTUAL_SERVER DELAY_LOOP LB_ALGO LB_KIND LVS_SCHED LVS_METHOD RR WRR LC WLC FO OVF LBLC LBLCR SH DH SED NQ NAT DR TUN PERSISTENCE_TIMEOUT PROTOCOL TCP UDP SORRY_SERVER REAL_SERVER FWMARK INHIBIT_ON_FAILURE TCP_CHECK HTTP_GET SSL_GET SMTP_CHECK DNS_CHECK MISC_CHECK URL PATH DIGEST STATUS_CODE CONNECT_TIMEOUT CONNECT_PORT CONNECT_IP BINDTO BIND_PORT RETRY HELO_NAME TYPE NAME MISC_PATH MISC_TIMEOUT WARMUP MISC_DYNAMIC NB_GET_RETRY DELAY_BEFORE_RETRY VIRTUALHOST ALPHA OMEGA QUORUM HYSTERESIS QUORUM_UP QUORUM_DOWN

%type<blocks_any> configuration main_blocks
%type<block> global vrrp_instance_block static_ipaddress_block static_routes_block static_rules_block  vrrp_sync_group_block vrrp_instance_block vrrp_script_block virtual_server_group_block 
%type<block_args> virtual_server_block 
%type<stmts_any> global_statements vrrp_instance_statements vrrp_sync_group_statements vrrp_script_statements address_options route_options rule_options virtual_server_group_statements virtual_server_statements real_server_statements
%type<stmt_any> vrrp_instance_statement vrrp_sync_group_statement vrrp_script_statement virtual_server_group_statement virtual_server_statement route_option rule_option real_server_statement
%type<values> vips
%type<vip_addr> vip
%type<strings> virtual_server_arg
%type<string> ipaddr_literal ip46 lb_algo lb_kind protocol

%%

configuration: main_blocks configuration
  {
    $$ = append($1, $2...)
    yylex.(*Lexer).result = $$
  }
  | main_blocks
  { 
    $$ = $1
    yylex.(*Lexer).result = $$
  }

main_blocks:
    global { $$ = []BlockAny{$1} }
  | static_ipaddress_block { $$ = []BlockAny{$1} }
  | static_routes_block { $$ = []BlockAny{$1} }
  | static_rules_block { $$ = []BlockAny{$1} }
  | vrrp_sync_group_block { $$ = []BlockAny{$1} }
  | vrrp_instance_block { $$ = []BlockAny{$1} }
  | vrrp_script_block { $$ = []BlockAny{$1} }
  | virtual_server_block { $$ = []BlockAny{$1} }
  | virtual_server_group_block { $$ = []BlockAny{$1} }

global:
	GLOBALDEFS LB global_statements RB
  {
    $$ = Block{name: $1.lit, stmts: $3}
  }

global_statements:	global_statement global_statements { } | global_statement { }

global_statement:
| NOTIFICATION_EMAIL LB mail_statements RB  { }
| NOTIFICATION_EMAIL STRING  { }
| NOTIFICATION_EMAIL_FROM EMAIL { }
| VRRP_GARP_MASTER_DELAY NUMBER { }
| VRRP_GARP_MASTER_REPEAT NUMBER { }
| VRRP_GARP_MASTER_REFRESH NUMBER { }
| VRRP_GARP_MASTER_REFRESH_REPEAT NUMBER { }
| SMTP_SERVER ip46  { }
| SMTP_SERVER STRING  { }
| SMTP_CONNECT_TIMEOUT NUMBER { }
| ROUTER_ID STRING { }
| LVS_ID STRING { }
| VRRP_MCAST_GROUP4 IPV4 { }
| VRRP_MCAST_GROUP6 IPV6 { }
| VRRP_VERSION NUMBER { }

static_ipaddress_block:
  STATIC_IPADDRESS LB address_options RB
  {
    $$ = Block{name: $1.lit, stmts: $3}
  }

static_routes_block:
  STATIC_ROUTES LB route_options RB
  {
    $$ = Block{name: $1.lit, stmts: $3}
  }

static_rules_block:
  STATIC_RULES LB rule_options RB
  {
    $$ = Block{name: $1.lit, stmts: $3}
  }

rule_options:
 rule_option rule_options { } | rule_option { }

vrrp_sync_group_block:
  VRRP_SYNC_GROUP STRING LB vrrp_sync_group_statements RB
  {
    $$ = Block{name: $1.lit, stmts: $4}
  }

vrrp_sync_group_statements:
  vrrp_sync_group_statement vrrp_sync_group_statements
  {
    $$ = append([]StmtAny{$1}, $2...)
  }
  | vrrp_sync_group_statement { $$ = []StmtAny{$1} }

vrrp_sync_group_statement: { }
| STRING { }
| GROUP LB vrrp_group_ids RB { }
| NOTIFY_MASTER	STRING { }
| NOTIFY_BACKUP	STRING { }
| NOTIFY_FAULT	STRING { }
| NOTIFY	STRING { }

vrrp_group_ids: vrrp_group_id vrrp_group_ids | vrrp_group_id { }

vrrp_group_id: STRING

vrrp_instance_block :
  VRRP_INSTANCE STRING LB vrrp_instance_statements RB
  {
    $$ = Block{name: $1.lit, stmts: $4}
  }

vrrp_instance_statements :
  vrrp_instance_statement vrrp_instance_statements 
  {
    $$ = append([]StmtAny{$1}, $2...)
  }
  | vrrp_instance_statement { $$ = []StmtAny{$1} }

vrrp_instance_statement: { }
| USE_VMAC { }
| VERSION NUMBER { }
| VMAC_XMIT_BASE { }
| NATIVE_IPV6 { }
| INTERFACE STRING { }
| MCAST_SRC_IP ip46 { }
| UNICAST_SRC_IP ip46 { }
| UNICAST_PEER LB ipaddr_list RB { }
| LVS_SYNC_DAEMON_INTERFACE STRING { }
| VIRTUAL_ROUTER_ID STRING { }
| VIRTUAL_ROUTER_ID NUMBER { }
| NOPREEMPT { }
| PREEMPT_DELAY NUMBER { }
| PRIORITY NUMBER { }
| ADVERT_INT NUMBER { }
| VIRTUAL_IPADDRESS LB vips RB
  {
    $$ = StmtMulti{$1.lit, $3}
  }
| VIRTUAL_IPADDRESS_EXCLUDED LB vips RB
  {
    $$ = StmtMulti{$1.lit, $3}
  }
| VIRTUAL_ROUTES LB virtual_routes_statements RB { }
| STATE MASTER { }
| STATE BACKUP { }
| GARP_MASTER_DELAY NUMBER { }
| SMTP_ALERT { }
| AUTHENTICATION LB authentication_statements RB { }
| TRACK_INTERFACE STRING { }
| TRACK_INTERFACE LB interfaces RB { }
| TRACK_SCRIPT LB track_script_statements RB { }
| DONT_TRACK_PRIMARY { }
| NOTIFY_MASTER STRING { }
| NOTIFY_BACKUP STRING { }
| NOTIFY_FAULT STRING { }
| NOTIFY_STOP STRING { }
| NOTIFY STRING { }

authentication_statements: authentication_statement authentication_statements | authentication_statement

authentication_statement: { }
| AUTH_TYPE PASS { }
| AUTH_TYPE AH { }
| AUTH_PASS any_literal { }

interfaces: interface interfaces | interface

interface: STRING
| WEIGHT NUMBER { }

ipaddr_list: ip46 ipaddr_list | ip46

track_script_statements: track_script_statement track_script_statements | track_script_statement

track_script_statement: { }
| STRING { }
| STRING WEIGHT NUMBER { }

vrrp_script_block:
  VRRP_SCRIPT STRING LB vrrp_script_statements RB
  {
    $$ = Block{name: $1.lit, stmts: $4}
  }

vrrp_script_statements:
  vrrp_script_statement vrrp_script_statements
  {
    $$ = append([]StmtAny{$1}, $2...)
  }
  | vrrp_script_statement { $$ = []StmtAny{$1} }

vrrp_script_statement: { }
| SCRIPT STRING { }
| INTERVAL NUMBER { }
| TIMEOUT NUMBER { }
| WEIGHT NUMBER { }
| FALL NUMBER { }
| RISE NUMBER { }

virtual_routes_statements: virtual_server_statement virtual_server_statements | virtual_server_statement 

virtual_server_statement:
  route_option route_options
  {
    $$ = append([]StmtAny{$1}, $2...)
  }
  | route_option { $$ = []StmtAny{$1} }

virtual_server_group_block:
  VIRTUAL_SERVER_GROUP STRING LB virtual_server_group_statements RB
  {
    $$ = Block{name: $1.lit, stmts: $4}
  }

virtual_server_group_statements:
  virtual_server_group_statement virtual_server_group_statements
  {
    $$ = append([]StmtAny{$1}, $2...)
  }
  | virtual_server_group_statement { $$ = []StmtAny{$1} }

virtual_server_group_statement: { }
| ip46 NUMBER { }
| IPADDR_RANGE NUMBER { }
| FWMARK NUMBER { }

virtual_server_block:
  VIRTUAL_SERVER virtual_server_arg LB virtual_server_statements RB
  {
    $$ = BlockArgs{Name: $1.lit, Args: $2, Stmts: $4}
  }

virtual_server_statements:
  virtual_server_statement virtual_server_statements
  {
    $$ = append([]StmtAny{$1}, $2...)
  }
  | virtual_server_statement { $$ = []StmtAny{$1} }

virtual_server_arg:
    ip46 NUMBER { $$ = []string{$1, $2.lit} }
  | FWMARK NUMBER { $$ = []string{$1.lit, $2.lit} }
  | GROUP STRING { $$ = []string{$1.lit, $2.lit} }

virtual_server_statement:
  DELAY_LOOP NUMBER { $$ = Stmt{$1.lit, $2.lit} }
| LB_ALGO lb_algo { $$ = Stmt{$1.lit, $2} }
| LB_KIND lb_kind { $$ = Stmt{$1.lit, $2} }
| LVS_SCHED lb_algo { $$ = Stmt{$1.lit, $2} }
| LVS_METHOD lb_kind { $$ = Stmt{$1.lit, $2} }
| PERSISTENCE_TIMEOUT NUMBER { $$ = Stmt{$1.lit, $2.lit} }
| PROTOCOL protocol { $$ = Stmt{$1.lit, $2} }
| SORRY_SERVER ip46 NUMBER { $$ = StmtMulti{$1.lit, []Value{$2, $3.lit}} }
| REAL_SERVER ip46 NUMBER LB real_server_statements RB
{ 
  $$ = SubBlockArgs{Name: $1.lit, Args: []string{$2, $3.lit}, Stmts: $5}
}
| VIRTUALHOST STRING { $$ = Stmt{$1.lit, $2.lit} }
| ALPHA { $$ = $1.lit }
| OMEGA { $$ = $1.lit }
| QUORUM NUMBER { $$ = Stmt{$1.lit, $2.lit} }
| HYSTERESIS NUMBER { $$ = Stmt{$1.lit, $2.lit} }
| QUORUM_UP STRING { $$ = Stmt{$1.lit, $2.lit} }
| QUORUM_DOWN STRING { $$ = Stmt{$1.lit, $2.lit} }

real_server_statements: real_server_statement real_server_statements { } | real_server_statement { }

real_server_statement: { }
| WEIGHT NUMBER { }
| INHIBIT_ON_FAILURE { }
| TCP_CHECK LB tcp_check_statements RB { }
| HTTP_GET LB http_get_statements RB { }
| SSL_GET LB http_get_statements RB { }
| SMTP_CHECK LB smtp_check_statements RB { }
| DNS_CHECK LB dns_check_statements RB { }
| MISC_CHECK LB misc_check_statements RB { }

tcp_check_statements: tcp_check_statement tcp_check_statements | tcp_check_statement { }

tcp_check_statement: { }
| CONNECT_PORT NUMBER { }
| CONNECT_TIMEOUT NUMBER { }
| RETRY NUMBER { }
| WARMUP NUMBER { }
| DELAY_BEFORE_RETRY NUMBER { }

http_get_statements: http_get_statement http_get_statements | http_get_statement { }

http_get_statement: { }
| URL LB url_statements RB { }
| CONNECT_TIMEOUT NUMBER { }
| CONNECT_PORT NUMBER { }
| NB_GET_RETRY NUMBER { }
| DELAY_BEFORE_RETRY NUMBER { }
| WARMUP NUMBER { }

url_statements: url_statement url_statements | url_statement { }

url_statement: { }
| PATH STRING { }
| PATH PATHSTR { }
| DIGEST HEX32 { }
| STATUS_CODE NUMBER { }

smtp_check_statements: smtp_check_statement smtp_check_statements | smtp_check_statement { }

smtp_check_statement: { }
| host_statement { }
| HOST LB host_statements RB
| WARMUP NUMBER
| RETRY NUMBER
| DELAY_BEFORE_RETRY NUMBER
| HELO_NAME STRING

host_statements: host_statement host_statements | host_statement { }

host_statement: { }
| CONNECT_IP ip46 { }
| CONNECT_PORT NUMBER { }
| BINDTO ip46 { }
| BIND_PORT NUMBER { }
| CONNECT_TIMEOUT NUMBER { }
| FWMARK NUMBER { }

dns_check_statements: dns_check_statement dns_check_statements | dns_check_statement { }

dns_check_statement: { }
| CONNECT_IP ip46 { }
| CONNECT_PORT NUMBER { }
| BINDTO ip46 { }
| BIND_PORT NUMBER { }
| CONNECT_TIMEOUT NUMBER { }
| FWMARK NUMBER { }
| RETRY NUMBER { }
| TYPE STRING { }
| NAME STRING { }

misc_check_statements: misc_check_statement misc_check_statements | misc_check_statement { }

misc_check_statement: { }
| MISC_PATH STRING { }
| MISC_PATH PATHSTR { }
| MISC_TIMEOUT NUMBER { }
| WARMUP NUMBER { }
| MISC_DYNAMIC { }

lb_algo:
    RR   { $$ = $1.lit }
  | WRR  { $$ = $1.lit }
  | LC   { $$ = $1.lit }
  | WLC  { $$ = $1.lit }
  | FO   { $$ = $1.lit }
  | OVF  { $$ = $1.lit }
  | LBLC { $$ = $1.lit }
  | LBLCR { $$ = $1.lit }
  | SH   { $$ = $1.lit }
  | DH   { $$ = $1.lit }
  | SED   { $$ = $1.lit }
  | NQ    { $$ = $1.lit }

lb_kind:
    NAT	{ $$ = $1.lit }
  | DR  { $$ = $1.lit }
  | TUN	{ $$ = $1.lit }

protocol:
    TCP { $$ = $1.lit }
  | UDP { $$ = $1.lit }

vips:
  vip vips
  {
    $$ = append([]Value{$1}, $2...)
  }
  | vip { $$ = []Value{$1} }

vip: 
  ipaddr_literal { $$ = VIPAddr{Addr: $1} }
| LABEL STRING { }
| DEV STRING { }
| BRD ip46 { }

mail_statements: mail_statement mail_statements |  mail_statement { }

mail_statement:	any_literal	{ }

address_options: address_option address_options { } | address_option { }

address_option: { }
| ip46
| IP_CIDR
| BRD ip46
| DEV STRING
| SCOPE scope_val

route_options: route_option route_options { } | route_option { }

route_option: { }
| SRC ip46 { }
| TO ip46 { }
| TO IP_CIDR { }
| ip46 { }
| IP_CIDR { }
| VIA ip46 { }
| GW ip46 { }
| OR ip46 { }
| DEV STRING { }
| LABEL STRING { }
| TABLE NUMBER { }
| TABLE STRING { }
| METRIC NUMBER { }
| SCOPE scope_val { }
| BLACKHOLE ip46 { }
| BLACKHOLE IP_CIDR { }

rule_option:
  FROM ip46 TABLE NUMBER { }
| TO ip46 TABLE NUMBER { }
| FROM IP_CIDR TABLE NUMBER { }
| TO IP_CIDR TABLE NUMBER { }

scope_val: { }
| SITE
| LINK
| HOST
| NOWHERE
| GLOBAL

ipaddr_literal:
  ip46 { $$ = $1 }
  | IP_CIDR { $$ = $1.lit }

ip46:
    IPV4 { $$ = $1.lit }
  | IPV6 { $$ = $1.lit }

ipport: { }
| ip46
| ip46 NUMBER

any_literal: { }
| NUMBER
| STRING
| EMAIL
| HEX32
| ip46
| IP_CIDR

%%
