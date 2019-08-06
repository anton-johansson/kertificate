drop table "User";
create table "User"
(
    "userId"        serial primary key  not null
,   "username"      varchar(30)         not null
,   "firstName"     varchar(100)        not null
,   "lastName"      varchar(100)        not null
,   "emailAddress"  varchar(100)        not null
,   "active"        boolean             not null    default true
,   "loggedInAt"    timestamptz         not null    default now()
);

create type "KeySize" as enum
(
    '1024'
,   '2048'
,   '4096'
);

create type "Digest" as enum
(
    'SHA-1'
);

create table "CertificateTemplate"
(
    "templateId"            serial primary key  not null
,   "name"                  varchar(30)         not null
,   "countryCode"           varchar(100)        not null
,   "state"                 varchar(100)        not null
,   "locality"              varchar(100)        not null
,   "organizationName"      varchar(100)        not null
,   "organizationalUnit"    varchar(100)        not null
,   "emailAddress"          varchar(100)        not null
,   "validFor"              int                 not null
,   "keySize"               "KeySize"           not null
,   "digest"                "Digest"            not null
);

insert into "CertificateTemplate"
(
    "name"
,   "countryCode"
,   "state"
,   "locality"
,   "organizationName"
,   "organizationalUnit"
,   "emailAddress"
,   "validFor"
,   "keySize"
,   "digest"
)
values
(
    'Basic'
,   'SE'
,   'Västra Götalands län'
,   'Borås'
,   'Anton AB'
,   'Operations'
,   'antoon.johansson@gmail.com'
,   365
,   '4096'
,   'SHA-1'
);

select  *
from    "CertificateTemplate";

drop table if exists "PrivateKey";
create table "PrivateKey"
(
    "keyId"     serial primary key  not null
,   "content"   varchar(1000)       not null
);

drop table if exists "PublicKey";
create table "PublicKey"
(
    "keyId"     serial primary key  not null
,   "content"   varchar(1000)       not null
);

drop table if exists "CertificateData";
create table "CertificateData"
(
    "certificateDataId"     serial primary key  not null
,   "privateKeyData"        bytea               not null
,   "certificateData"       bytea               not null
,   "expiresAt"             timestamptz         not null
)
select * from "CertificateData";

drop table if exists "CommonAuthority";
create table "CommonAuthority"
(
    "commonAuthorityId"     serial primary key  not null
,   "name"                  varchar(30)         not null
,   "certificateDataId"     int                 not null    references "CertificateData"("certificateDataId")
,   "createdBy"             int                 not null    references "User"("userId")
);
select * from "CommonAuthority";

drop table if exists "ConsumerType";
create table "ConsumerType"
(
    "typeId"        serial primary key  not null
,   "name"          varchar(30)         not null
);

drop table if exists "Consumer";
create table "Consumer"
(
    "consumerId"    serial primary key  not null
,   "name"          varchar(30)         not null
,   "consumerType"  int                 references "ConsumerType"("typeId")
,   "emailAddress"  varchar(100)        not null
,   "comments"      varchar             not null
);

drop table if exists "Certificate";
create table "Certificate"
(
    "certificateId"         serial primary key  not null
,   "name"                  varchar(30)         not null
,   "certificateDataId"     int                 not null    references "CertificateData"("certificateDataId")
,   "commonAuthorityId"     int                 not null    references "CommonAuthority"("commonAuthorityId")
,   "createdBy"             int                 not null    references "User"("userId")
);
select * from "Certificate";


update "User" set "active" = true where "userId" = 4;
select * from "ConsumerType";
select * from "User";
delete from "User" where "userId" = 7;
