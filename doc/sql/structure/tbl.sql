create database test2 collate SQL_Latin1_General_CP1_CI_AS
go

create schema dbo
    go


create table roles
(
    id int identity
        primary key,
    user_role_type varchar(50) not null,
    created datetime not null,
    updated datetime default getdate()
)
    go

create table trade_categories
(
    id int identity
        primary key,
    name varchar(200) not null,
    is_active bit default 1 not null,
    is_deleted bit default 0 not null,
    created datetime constraint DF__basic_tra__creat__4AB81AF0 default getdate(),
    updated datetime
)
    go

create unique index basic_trades_name_uindex
    on trade_categories (name)
    go

create table users
(
    id int identity
        primary key,
    user_role_id int not null
        constraint users_roles_id_fk
        references roles,
    username varchar(200) not null,
    password varchar(256) not null,
    email varchar(256) not null,
    password_token varchar(200),
    token_expires datetime,
    deletion_reason varchar(200),
    is_active bit default 1 not null,
    is_deleted bit default 0 not null,
    created datetime,
    updated datetime default getdate()
)
    go

create table profiles
(
    id int identity
        primary key,
    user_id int not null
        constraint profiles_users_id_fk
        references users,
    image_url varchar(200),
    mobile varchar(50),
    company varchar(200),
    first_name varchar(200),
    last_name varchar(200),
    is_active bit default 1 not null,
    is_deleted bit default 0 not null,
    created datetime,
    updated datetime default getdate()
)
    go

create table projects
(
    id int identity
        primary key,
    manager_id int not null
        constraint fk_project_manager_id
        references users,
    creator_id int not null
        constraint fk_project_creator_id
        references users,
    name varchar(200),
    total_trade_value float,
    contractor_total_claim float,
    serial_no varchar(50),
    address varchar(200),
    total_contract_value float,
    quantity_surveyor varchar(200),
    notes varchar(200),
    is_active bit default 1 not null,
    is_deleted bit default 0 not null,
    created datetime constraint DF__projects__created__45F365D3 default getdate(),
    updated datetime
)
    go

create table trades
(
    id int identity
        primary key,
    trade_category_id int not null
        constraint trades_trade_category_id_fk
        references trade_categories,
    creator_id int not null
        constraint trades_users_id_fk
        references users,
    project_id int not null
        constraint trades_projects_id_fk
        references projects,
    level varchar(20) not null,
    subtitle varchar(200),
    value float,
    temp bit,
    editable bit default 1 not null,
    is_deleted bit default 0 not null,
    created datetime,
    updated datetime default getdate()
)
    go

create table claims
(
    id int identity
        primary key,
    trade_id int not null
        constraint claims_trades_id_fk
        references trades,
    user_id int not null
        constraint claims_users_id_fk
        references users,
    trade_category_id int not null
        constraint claims_trade_category_id_fk
        references trade_categories,
    total_amount float,
    claimed_amount float,
    previous_claimed float,
    amount_due float,
    cost_to_completed float,
    claim_no varchar(200),
    claim_period varchar(50),
    action_claim bit,
    old_claimed_amount float,
    claim_percentage float,
    is_active bit default 1 not null,
    is_deleted bit default 0 not null,
    created datetime,
    updated datetime default getdate()
)
    go

create table claim_histories
(
    id int identity
        primary key,
    trade_id int not null
        constraint claim_histories_trades_id_fk
        references trades,
    claim_id int not null
        constraint claim_histories_claims_id_fk
        references claims,
    profile_id int not null,
    previous_claimed float,
    is_active bit default 1 not null,
    is_deleted bit default 0 not null,
    created datetime,
    updated datetime default getdate()
)
    go

create unique index users_email_uindex
    on users (email)
    go


/*
-- Identity problem
SET IDENTITY_INSERT test2.dbo.trades ON
-- insertions ...ADD
SET IDENTITY_INSERT test2.dbo.trades OFF
GO

ALTER AUTHORIZATION ON test2.dbo.roles
TO onsitedbadmin
go
*/