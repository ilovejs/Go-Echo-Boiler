-- Note: No mysqldump or pg_dump equivalent for Microsoft SQL Server,
-- so generated tests must be supplemented by tables_schema.sql with CREATE TABLE ... queries


create schema dbo
go

create table basic_tradeitems
(
    id        int identity
        primary key,
    TradeName varchar(200),
    CreatedOn datetime,
    IsDeleted bit
)
go

create table roles
(
    id             int identity
        primary key,
    user_role_type varchar(20),
    created        datetime
)
go

create table logins
(
    id               bigint identity
        primary key,
    username         nvarchar(200),
    password         nvarchar(256),
    user_role_id     int
        references roles,
    is_active        bit,
    is_deleted       bit,
    deletion_reason  nvarchar(200),
    created          datetime,
    updated          datetime,
    creator_login_id bigint
        references logins
)
go

create table profiles
(
    id                    bigint identity
        primary key,
    login_id              bigint
        references logins,
    mobile                varchar(20),
    first_name            nvarchar(200),
    last_name             nvarchar(200),
    email                 nvarchar(256) not null,
    is_active             bit,
    is_deleted            bit,
    created               datetime,
    updated               datetime,
    forget_password_token varchar(200),
    token_expires         datetime,
    company               nvarchar(200)
)
go

create table projects
(
    id                     bigint identity
        primary key,
    user_id                bigint not null
        references profiles,
    name                   nvarchar(200),
    total_item_breakdown   decimal(18, 2),
    contractor_total_claim decimal(18, 2),
    is_deleted             bit,
    created                datetime,
    updated                datetime,
    login_id               bigint not null
        references profiles,
    serial_no              nvarchar(50),
    details                nvarchar(200),
    total_contract_value   decimal(18, 2),
    quantity_surveyor      nvarchar(200),
    notes                  nvarchar(200)
)
go

create table contractor_projects
(
    id         bigint identity
        primary key,
    user_id    bigint not null
        references profiles,
    project_id bigint not null
        references projects,
    trade_id   int    not null
        references basic_tradeitems,
    notes      nvarchar(max),
    is_deleted bit,
    created    datetime,
    updated    datetime
)
go

create table trade_items
(
    id             bigint identity
        constraint PK__TradeIte__3214EC078EB45EF1
            primary key,
    trade_id       int         not null
        constraint FK__TradeItem__Trade__2A4B4B5E
            references basic_tradeitems,
    floor_level    varchar(20) not null,
    work_desc      nvarchar(200),
    item_breakdown decimal(18, 2),
    is_deleted     bit,
    created        datetime,
    updated        datetime,
    login_id       bigint      not null
        constraint FK__TradeItem__Creat__2B3F6F97
            references profiles,
    project_id     bigint      not null
        constraint FK__TradeItem__Proje__5EBF139D
            references projects,
    tempcheck      bit
)
go

create table claims
(
    id                 bigint identity
        constraint PK__TradeIte__3214EC0775241EB8
            primary key,
    trade_item_id      bigint not null
        constraint FK__TradeItem__bigin__70DDC3D8
            references trade_items,
    total_amount       decimal(18, 2),
    ClaimedAmount      decimal(18, 2),
    PreviousClaimed    decimal(18, 2),
    AmountDue          decimal(18, 2),
    CostToCompleted    decimal(18, 2),
    IsDeleted          bit,
    CreatedOn          datetime,
    UpdatedOn          datetime,
    login_id           bigint not null
        constraint FK__TradeItem__Creat__71D1E811
            references profiles,
    ClaimNumber        nvarchar(200),
    ClaimPeriod        varchar(50),
    action_claim       bit,
    AutoIncrement      bigint,
    old_claimed_amount decimal(18, 2),
    claim_percentage   decimal(18, 2),
    trade_id           int    not null
        constraint FK_TradeItemClaim_TradeMaster
            references basic_tradeitems
)
go

create table claim_histories
(
    id              bigint identity
        primary key,
    trade_item_id   bigint not null
        references trade_items,
    claim_id        bigint not null
        constraint FK__TradeItem__Trade__3F115E1A
            references claims,
    PreviousClaimed decimal(18, 2),
    CreatedOn       datetime,
    login_id        bigint not null
        references profiles,
    AutoIncrement   bigint
)
go

