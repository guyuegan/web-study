drop table userinfo;
create table if not exists userinfo (
    uid serial not null ,
    username character varying(100) not null ,
    department character varying(500) not null ,
    created date,
    constraint userinfo_pkey primary key (uid)
)
with (oids=false);

drop table userdetail;
create table if not exists userdetail
(
    uid integer,
    intro character varying(100),
    profile character varying(100)
)
without oids