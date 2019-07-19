/*
-- Note: No mysqldump or pg_dump equivalent for Microsoft SQL Server,
-- so generated tests must be supplemented by tables_schema.sql with CREATE TABLE ... queries
-- Don't run in test environment

USE [master]
DROP DATABASE test2;
create database test2;
ALTER AUTHORIZATION ON DATABASE::test2 TO tester;
use test2;
*/

create table roles
(
    id             int identity(1,1) primary key,
    user_role_type varchar(50),
    created        datetime,
    updated        datetime default(getdate()),
)
go

create table users
(
    id               int identity(1,1) primary key,

    user_role_id       int not null,

    username         varchar(200) unique,
    password         varchar(256),
    email            varchar(256) not null,
    password_token   varchar(200),
    token_expires    datetime,
    deletion_reason  varchar(200),

    is_active        bit,
    is_deleted       bit,
    created          datetime,
    updated          datetime default(getdate()),
)
go

create table profiles
(
    id                    int identity(1,1) primary key,
    user_id               int not null,

    image_url			  varchar(200),
    mobile                varchar(50),
    company               varchar(200),
    first_name            varchar(200),
    last_name             varchar(200),

    is_active             bit,
    is_deleted            bit,
    created               datetime,
    updated               datetime default(getdate()),
)
go

go

create table projects
(
    id                     int identity(1,1) primary key,

    manager_profile_id     int not null,
    creator_profile_id     int not null,

    name                   varchar(200),
    total_item_breakdown   float,
    contractor_total_claim float,

    serial_no              varchar(50),
    details                varchar(200),
    total_contract_value   float,
    quantity_surveyor      varchar(200), --?
    notes                  varchar(200),

    is_deleted             bit,
    created                datetime,
    updated                datetime default(getdate())
)
go


create table basic_trades
(
  id			int identity(1,1) primary key,
  name		varchar(200) not null,
  created		datetime null,
  updated	    datetime default(getdate()),
  is_deleted  bit      default(null),
)

create table trades
(
    id              int identity(1,1) primary key,

    basic_trade_id  int      not null,
    profile_id      int      not null,
    project_id      int      not null,

    floor_level     varchar(20) not null,
    work_desc       varchar(200),
    item_breakdown  float,
    tempcheck      bit,

    created        datetime,
    updated        datetime default(getdate()),
    is_deleted     bit      default(null)
)
go

create table claims
(
    id                 int identity(1,1) primary key,
    trade_id           int not null,
    profile_id         int not null,
    basic_trade_id     int not null,

    total_amount       float,
    claimed_amount     float,
    previous_claimed   float,
    amount_due         float,
    cost_to_completed  float,

    claim_number        varchar(200),
    claim_period        varchar(50),
    action_claim        bit,

    old_claimed_amount float,
    claim_percentage   float,

    is_deleted         bit null,
    created            datetime,
    updated            datetime default(getdate()),
)
go

create table claim_histories
(
    id              int identity(1,1) primary key,
    trade_id        int not null,
    claim_id        int not null,
    profile_id      int not null,

    previous_claimed  float,
    created           datetime,
    updated           datetime default(getdate()),
)
go

/* Do not run in test environment
alter table users add constraint FK__users__user_role_id foreign key (user_role_id) references roles (id);

alter table profiles add constraint FK__profiles__user_id foreign key(user_id) references users (id);
alter table projects add constraint FK__projects__creator_profile_id foreign key(creator_profile_id) references profiles (id);
alter table projects add constraint FK__projects__manager_profile_id foreign key(manager_profile_id) references profiles (id);

alter table trades add constraint FK__trades__basic_trade_id foreign key (basic_trade_id)   references basic_trades (id);
alter table trades add constraint FK__trades__profile_id     foreign key (profile_id) references profiles (id);
alter table trades add constraint FK__trades__project_id     foreign key (project_id) references projects (id)

alter table claims add constraint FK__claims__trade_id foreign key (trade_id) references trades (id)
alter table claims add constraint FK__claims__profile_id foreign key (profile_id) references profiles (id)
alter table claims add constraint FK__claims__basic_trade_id foreign key (basic_trade_id) references basic_trades (id)

alter table claim_histories add constraint FK__claim_histories_trade_id foreign key (trade_id) references trades (id);
alter table claim_histories add constraint FK__claim_histories_claim_id foreign key (claim_id) references claims (id);
alter table claim_histories add constraint FK__claim_histories_profile_id foreign key (profile_id) references profiles (id);

create table contractor_projects
(
    id         int identity(1,1)(1,1) primary key,
    user_id    int not null,

    project_id int not null references projects,
    trade_id   int    not null references basic_tradeitems,
    notes      varchar(max),

    is_deleted bit,
    created    datetime,
    updated    datetime
)
go
*/