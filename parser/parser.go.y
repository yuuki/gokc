%{
package parser
%}

%union {
  integer   int
  symbol    string
};

%token <integer> NUMBER
%token <symbol>	 ID STRING EMAIL IPADDR IP_CIDR IPADDR_RANGE HEX32 PATHSTR
%token           LB RB
%token           GLOBALDEFS
%token           NOTIFICATION_EMAIL NOTIFICATION_EMAIL_FROM SMTP_SERVER SMTP_CONNECT_TIMEOUT ROUTER_ID LVS_ID
%token           STATIC_IPADDRESS
%token           STATIC_ROUTES
%token           STATIC_RULES
%token           VRRP_SYNC_GROUP GROUP
%token           VRRP_INSTANCE
%token           INTERFACE MCAST_SRC_IP UNICAST_SRC_IP UNICAST_PEER LVS_SYNC_DAEMON_INTERFACE VIRTUAL_ROUTER_ID NOPREEMPT PRIORITY ADVERT_INT VIRTUAL_IPADDRESS VIRTUAL_IPADDRESS_EXCLUDED VIRTUAL_ROUTES STATE MASTER BACKUP GARP_MASTER_DELAY SMTP_ALERT AUTHENTICATION AUTH_TYPE AUTH_PASS PASS AH LABEL DEV SCOPE SITE LINK HOST NOWHERE GLOBAL BRD SRC FROM TO VIA GW OR TABLE METRIC TRACK_INTERFACE TRACK_SCRIPT NOTIFY_MASTER NOTIFY_BACKUP NOTIFY_FAULT NOTIFY_STOP NOTIFY BLACKHOLE
%token           VRRP_SCRIPT
%token           SCRIPT INTERVAL FALL RISE
%token           VIRTUAL_SERVER_GROUP
%token           VIRTUAL_SERVER
%token           DELAY_LOOP LB_ALGO LB_KIND LVS_SCHED LVS_METHOD RR WRR LC WLC LBLC SH DH NAT DR TUN PERSISTENCE_TIMEOUT PROTOCOL TCP UDP SORRY_SERVER REAL_SERVER FWMARK WEIGHT INHIBIT_ON_FAILURE TCP_CHECK HTTP_GET SSL_GET SMTP_CHECK MISC_CHECK URL PATH DIGEST STATUS_CODE CONNECT_TIMEOUT CONNECT_PORT CONNECT_IP BINDTO BIND_PORT RETRY HELO_NAME MISC_PATH MISC_TIMEOUT WARMUP MISC_DYNAMIC NB_GET_RETRY DELAY_BEFORE_RETRY VIRTUALHOST ALPHA OMEGA QUORUM HYSTERESIS QUORUM_UP QUORUM_DOWN,


%%
configuration:  main_statements configuration | main_statements { }

main_statements:  { }
| global { }
| static_ipaddress_block { }
| static_routes_block { }
| static_rules_block { }
| vrrp_sync_group_block { }
| vrrp_instance_block { }
| vrrp_script_block { }
| virtual_server_block { }
| virtual_server_group_block { }

global:	GLOBALDEFS LB global_statements RB

global_statements:	global_statement global_statements | global_statement

global_statement:
| NOTIFICATION_EMAIL LB mail_statements RB  { }
| NOTIFICATION_EMAIL STRING  { }
| NOTIFICATION_EMAIL_FROM EMAIL { }
| SMTP_SERVER IPADDR  { }
| SMTP_SERVER STRING  { }
| SMTP_CONNECT_TIMEOUT NUMBER { }
| ROUTER_ID STRING { }
| LVS_ID STRING { }

static_ipaddress_block: STATIC_IPADDRESS LB address_options RB

static_routes_block: STATIC_ROUTES LB route_options RB

static_rules_block: STATIC_RULES LB rule_options RB

rule_options: rule_option rule_options | rule_option

vrrp_sync_group_block: VRRP_SYNC_GROUP STRING LB vrrp_sync_group_statements RB

vrrp_sync_group_statements: vrrp_sync_group_statement vrrp_sync_group_statements | vrrp_sync_group_statement

vrrp_sync_group_statement: { }
| STRING { }
| GROUP LB vrrp_group_ids RB { }
| NOTIFY_MASTER	STRING { }
| NOTIFY_BACKUP	STRING { }
| NOTIFY_FAULT	STRING { }

vrrp_group_ids: vrrp_group_id vrrp_group_ids | vrrp_group_id { }

vrrp_group_id: STRING

vrrp_instance_block: VRRP_INSTANCE STRING LB vrrp_instance_statements RB

vrrp_instance_statements: vrrp_instance_statement vrrp_instance_statements | vrrp_instance_statement

vrrp_instance_statement: { }
| INTERFACE STRING { }
| MCAST_SRC_IP IPADDR { }
| UNICAST_SRC_IP IPADDR { }
| UNICAST_PEER LB ipaddr_list RB { }
| LVS_SYNC_DAEMON_INTERFACE STRING { }
| VIRTUAL_ROUTER_ID STRING { }
| VIRTUAL_ROUTER_ID NUMBER { }
| NOPREEMPT
| PRIORITY NUMBER { }
| ADVERT_INT NUMBER { }
| VIRTUAL_IPADDRESS LB vips RB
| VIRTUAL_IPADDRESS_EXCLUDED LB vips_ex RB
| VIRTUAL_ROUTES LB virtual_routes_statements RB
| STATE MASTER { }
| STATE BACKUP { }
| GARP_MASTER_DELAY NUMBER { }
| SMTP_ALERT { }
| AUTHENTICATION LB authentication_statements RB { }
| TRACK_INTERFACE STRING { }
| TRACK_INTERFACE LB interfaces RB { }
| TRACK_SCRIPT LB track_script_statements RB { }
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

ipaddr_list: IPADDR ipaddr_list | IPADDR

track_script_statements: track_script_statement track_script_statements | track_script_statement

track_script_statement: { }
| STRING { }
| STRING WEIGHT NUMBER { }

vrrp_script_block: VRRP_SCRIPT STRING LB vrrp_script_statements RB

vrrp_script_statements: vrrp_script_statement vrrp_script_statements | vrrp_script_statement

vrrp_script_statement: { }
| SCRIPT STRING { }
| INTERVAL NUMBER { }
| WEIGHT NUMBER { }
| FALL NUMBER { }
| RISE NUMBER { }

virtual_routes_statements: virtual_server_statement virtual_server_statements | virtual_server_statement 

virtual_server_statement: route_option route_options | route_option

virtual_server_group_block: VIRTUAL_SERVER_GROUP STRING LB virtual_server_group_statements RB

virtual_server_group_statements: virtual_server_group_statement virtual_server_group_statements | virtual_server_group_statement

virtual_server_group_statement: { }
| IPADDR NUMBER { }
| IPADDR_RANGE NUMBER { }
| FWMARK NUMBER { }

virtual_server_block: VIRTUAL_SERVER virtual_server_arg LB virtual_server_statements RB

virtual_server_statements: virtual_server_statement virtual_server_statements | virtual_server_statement

virtual_server_arg:
| ipport
| FWMARK NUMBER { }
| GROUP STRING { }

virtual_server_statement: { }
| DELAY_LOOP NUMBER { }
| LB_ALGO lb_algo { }
| LB_KIND lb_kind { }
| LVS_SCHED lb_algo { }
| LVS_METHOD lb_kind { }
| PERSISTENCE_TIMEOUT NUMBER { }
| PROTOCOL protocol { }
| SORRY_SERVER IPADDR NUMBER { }
| REAL_SERVER IPADDR ipport LB real_server_statements RB
| REAL_SERVER IPADDR NUMBER LB real_server_statements RB
| VIRTUALHOST STRING { }
| ALPHA
| OMEGA
| QUORUM NUMBER { }
| HYSTERESIS NUMBER { }
| QUORUM_UP STRING { }
| QUORUM_DOWN STRING { }

real_server_statements: real_server_statement real_server_statements | real_server_statement { }

real_server_statement: { }
| WEIGHT NUMBER { }
| INHIBIT_ON_FAILURE { }
| TCP_CHECK LB tcp_check_statements RB { }
| HTTP_GET LB http_get_statements RB { }
| SSL_GET LB http_get_statements RB { }
| SMTP_CHECK LB smtp_check_statements RB { }
| MISC_CHECK LB misc_check_statements RB { }

tcp_check_statements: tcp_check_statement tcp_check_statements | tcp_check_statement { }

tcp_check_statement: { }
| CONNECT_PORT NUMBER { }
| CONNECT_TIMEOUT NUMBER { }

http_get_statements: http_get_statement http_get_statements | http_get_statement { }

http_get_statement: { }
| URL LB url_statements RB { }
| CONNECT_TIMEOUT NUMBER { }
| CONNECT_PORT NUMBER { }
| NB_GET_RETRY NUMBER { }
| DELAY_BEFORE_RETRY NUMBER { }

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
| CONNECT_IP IPADDR { }
| CONNECT_PORT NUMBER { }
| BINDTO IPADDR { }
| BIND_PORT NUMBER { }
| CONNECT_TIMEOUT NUMBER { }
| FWMARK NUMBER { }

misc_check_statements: misc_check_statement misc_check_statements | misc_check_statement { }

misc_check_statement: { }
| MISC_PATH STRING { }
| MISC_PATH PATHSTR { }
| MISC_TIMEOUT NUMBER { }
| WARMUP NUMBER { }
| MISC_DYNAMIC { }

lb_algo: { }
| RR   { }
| WRR  { }
| LC   { }
| WLC  { }
| LBLC { }
| SH   { }
| DH   { }

lb_kind: { }
| NAT	{ }	
| DR  { }
| TUN	{ }

protocol: { }
| TCP { }
| UDP { }

vips: vip vips | vip { }

vips_ex: vip_ex vips_ex | vip_ex { }

vip: ipaddr_literal { }
| LABEL STRING { }
| DEV STRING { }
| BRD IPADDR { }

vip_ex: ipaddr_literal { }
| DEV STRING { }
| BRD IPADDR { }

mail_statements: mail_statement mail_statements |  mail_statement { }

mail_statement:	any_literal	{ }

address_options: address_option address_options | address_option

address_option: { }
| IPADDR
| IP_CIDR
| BRD IPADDR
| DEV STRING
| SCOPE scope_val

route_options: route_option route_options | route_option

route_option: { }
| SRC IPADDR
| TO IPADDR
| TO IP_CIDR
| IPADDR
| IP_CIDR
| VIA IPADDR
| GW IPADDR
| OR IPADDR
| DEV STRING
| LABEL STRING
| TABLE NUMBER
| TABLE STRING
| METRIC NUMBER
| SCOPE scope_val
| BLACKHOLE IPADDR
| BLACKHOLE IP_CIDR

rule_option: { }
| FROM IPADDR TABLE NUMBER
| TO IPADDR TABLE NUMBER
| FROM IP_CIDR TABLE NUMBER
| TO IP_CIDR TABLE NUMBER

scope_val: { }
| SITE
| LINK
| HOST
| NOWHERE
| GLOBAL

ipaddr_literal: { }
| IPADDR
| IP_CIDR

ipport: { }
| IPADDR
| IPADDR NUMBER

any_literal: { }
| NUMBER
| STRING
| EMAIL
| HEX32
| IPADDR
| IP_CIDR

%%
