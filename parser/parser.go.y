%{
package parser

const MAX_PORT_NUM int = 65535
%}

%union {
  integer   int
  symbol    string
};

%token <integer> NUMBER
%token <symbol>	 ID STRING EMAIL IPADDR IP_CIDR HEX32 PATHSTR
%token           LB RB
%token           GLOBALDEFS
%token           NOTIFICATION_EMAIL NOTIFICATION_EMAIL_FROM SMTP_SERVER SMTP_CONNECT_TIMEOUT ROUTER_ID LVS_ID
%token           VRRP_SYNC_GROUP GROUP
%token           VRRP_INSTANCE
%token           INTERFACE LVS_SYNC_DAEMON_INTERFACE VIRTUAL_ROUTER_ID NOPREEMPT PRIORITY ADVERT_INT VIRTUAL_IPADDRESS STATE MASTER BACKUP GARP_MASTER_DELAY SMTP_ALERT AUTHENTICATION AUTH_TYPE AUTH_PASS PASS AH LABEL DEV BRD TRACK_INTERFACE TRACK_SCRIPT NOTIFY_MASTER NOTIFY_BACKUP NOTIFY_FAULT NOTIFY_STOP NOTIFY
%token           VRRP_SCRIPT
%token           SCRIPT INTERVAL FALL RISE
%token           VIRTUAL_SERVER
%token           DELAY_LOOP LB_ALGO LB_KIND LVS_SCHED LVS_METHOD RR WRR LC WLC LBLC SH DH NAT DR TUN PERSISTENCE_TIMEOUT PROTOCOL TCP UDP SORRY_SERVER REAL_SERVER FWMARK WEIGHT HTTP_GET URL PATH DIGEST STATUS_CODE CONNECT_TIMEOUT NB_GET_RETRY DELAY_BEFORE_RETRY

%%
configuration:  main_statements configuration | main_statements { }

main_statements:  { }
| global { }
| vrrp_sync_group_block { }
| vrrp_instance_block { }
| vrrp_script_block { }
| virtual_server_block { }

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
| LVS_SYNC_DAEMON_INTERFACE STRING { }
| VIRTUAL_ROUTER_ID STRING { }
| VIRTUAL_ROUTER_ID NUMBER { }
| NOPREEMPT
| PRIORITY NUMBER { }
| ADVERT_INT NUMBER { }
| VIRTUAL_IPADDRESS LB vips RB
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

virtual_server_block: VIRTUAL_SERVER iporfw LB virtual_server_statements RB

virtual_server_statements: virtual_server_statement virtual_server_statements | virtual_server_statement

virtual_server_statement: { }
| DELAY_LOOP NUMBER { }
| LB_ALGO lb_algo { }
| LB_KIND lb_kind { }
| LVS_SCHED lb_algo { }
| LVS_METHOD lb_kind { }
| PERSISTENCE_TIMEOUT NUMBER { }
| PROTOCOL protocol { }
| SORRY_SERVER IPADDR NUMBER { }
| REAL_SERVER IPADDR ipport real_server_statements RB
| REAL_SERVER IPADDR NUMBER LB real_server_statements RB

real_server_statements: real_server_statement real_server_statements | real_server_statement { }

real_server_statement: { }
| WEIGHT NUMBER { }
| HTTP_GET LB http_get_statements RB { }

http_get_statements: http_get_statement http_get_statements | http_get_statement { }

http_get_statement: { }
| URL LB url_statements RB { }
| CONNECT_TIMEOUT NUMBER { }
| NB_GET_RETRY NUMBER { }
| DELAY_BEFORE_RETRY NUMBER { }

url_statements: url_statement url_statements | url_statement { }

url_statement: { }
| PATH STRING { }
| PATH PATHSTR { }
| DIGEST HEX32 { }
| STATUS_CODE NUMBER { }

ipport: { }
| IPADDR { }
| IPADDR NUMBER { }

iporfw: { }
| ipport { }
| FWMARK NUMBER { }

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

vip: ipaddr_literal { }
| LABEL STRING { }
| DEV STRING { }
| BRD IPADDR { }

mail_statements: mail_statement mail_statements |  mail_statement { }

mail_statement:	any_literal	{ }

ipaddr_literal: { }
| IPADDR
| IP_CIDR

any_literal: { }
| NUMBER
| STRING
| EMAIL
| HEX32
| IPADDR
| IP_CIDR

%%
