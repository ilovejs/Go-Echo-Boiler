-- Note: No mysqldump or pg_dump equivalent for Microsoft SQL Server,
-- so generated tests must be supplemented by tables_schema.sql with CREATE TABLE ... queries

/*
USE [master]
DROP DATABASE test2;
create database test2;
use test2;
*/

create table roles
(
    id             int identity primary key,
    user_role_type nvarchar(50),
    created        datetime,
    updated        datetime default(getdate()),
)
go

create table users
(
    id               bigint identity primary key,

    user_role_id       int not null,
    creator_profile_id bigint not null,

    username         nvarchar(200) unique,
    password         nvarchar(256),
    email            nvarchar(256) not null,
    password_token   nvarchar(200),
    token_expires    datetime,
    deletion_reason  nvarchar(200),

    is_active        bit,
    is_deleted       bit,
    created          datetime,
    updated          datetime default(getdate()),
)
go

alter table users add constraint FK__logins__user_role_id foreign key (user_role_id) references roles (id);

go

create table profiles
(
    id                    bigint identity primary key,
    user_id               bigint not null,

    image_url			  nvarchar(200),
    mobile                nvarchar(50),
    company               nvarchar(200),
    first_name            nvarchar(200),
    last_name             nvarchar(200),

    is_active             bit,
    is_deleted            bit,
    created               datetime,
    updated               datetime default(getdate()),
)
go

--from user
alter table users add constraint FK__logins__creator_profile_id foreign key (creator_profile_id) references profiles (id);
alter table profiles add constraint FK__profiles__user_id foreign key(user_id) references users (id);

go

create table projects
(
    id                     bigint identity primary key,

    manager_profile_id     bigint not null,
    creator_profile_id     bigint not null,

    name                   nvarchar(200),
    total_item_breakdown   decimal(18,2),
    contractor_total_claim decimal(18,2),

    serial_no              nvarchar(50),
    details                nvarchar(200),
    total_contract_value   decimal(18,2),
    quantity_surveyor      nvarchar(200), --?
    notes                  nvarchar(200),

    is_deleted             bit,
    created                datetime,
    updated                datetime default(getdate())
)
go

alter table projects add constraint FK__projects__creator_profile_id foreign key(creator_profile_id) references profiles (id);
alter table projects add constraint FK__projects__manager_profile_id foreign key(manager_profile_id) references profiles (id);

/*
create table contractor_projects
(
    id         bigint identity primary key,
    user_id    bigint not null,

    project_id bigint not null references projects,
    trade_id   int    not null references basic_tradeitems,
    notes      nvarchar(max),

    is_deleted bit,
    created    datetime,
    updated    datetime
)
go
*/

create table basic_trades (
    id			bigint identity(1,1) primary key,

    name		nvarchar(200) not null,
    created		datetime null,
    updated	    datetime default(getdate()),
    is_deleted  bit      default(null),
)

create table trades
(
    id              bigint identity primary key,

    basic_trade_id  bigint      not null,
    profile_id      bigint      not null,
    project_id      bigint      not null,

    floor_level     nvarchar(20) not null,
    work_desc       nvarchar(200),
    item_breakdown  decimal(18,2),
    tempcheck      bit,

    created        datetime,
    updated        datetime default(getdate()),
    is_deleted     bit      default(null)
)
go

alter table trades add constraint FK__trades__basic_trade_id foreign key (basic_trade_id)   references basic_trades (id);
alter table trades add constraint FK__trades__profile_id     foreign key (profile_id) references profiles (id);
alter table trades add constraint FK__trades__project_id     foreign key (project_id) references projects (id)

create table claims
(
    id                 bigint identity primary key,
    trade_id           bigint not null,
    profile_id         bigint not null,
    basic_trade_id     bigint not null,

    total_amount       decimal(18, 2),
    claimed_amount     decimal(18, 2),
    previous_claimed   decimal(18, 2),
    amount_due         decimal(18, 2),
    cost_to_completed  decimal(18, 2),

    claim_number        nvarchar(200),
    claim_period        nvarchar(50),
    action_claim        bit,

    --AutoIncrement      bigint,
    old_claimed_amount decimal(18, 2),
    claim_percentage   decimal(18, 2),

    is_deleted         bit null,
    created            datetime,
    updated            datetime default(getdate()),
)
go

alter table claims add constraint FK__claims__trade_id foreign key (trade_id) references trades (id)
alter table claims add constraint FK__claims__profile_id foreign key (profile_id) references profiles (id)
alter table claims add constraint FK__claims__basic_trade_id foreign key (basic_trade_id) references basic_trades (id)

create table claim_histories
(
    id              bigint identity primary key,
    trade_id        bigint not null,
    claim_id        bigint not null,
    profile_id      bigint not null,

    previous_claimed  decimal(18, 2),
    created           datetime,
    updated           datetime default(getdate()),
    --AutoIncrement   bigint
)
go

alter table claim_histories add constraint FK__claim_histories_trade_id foreign key (trade_id) references trades (id);
alter table claim_histories add constraint FK__claim_histories_claim_id foreign key (claim_id) references claims (id);
alter table claim_histories add constraint FK__claim_histories_profile_id foreign key (profile_id) references profiles (id);
