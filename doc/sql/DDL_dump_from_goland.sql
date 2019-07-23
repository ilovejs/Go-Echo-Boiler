create table basic_trades
(
    id         int identity
        primary key,
    name       varchar(200)       not null,
    is_active  bit      default 1 not null,
    is_deleted bit      default 0 not null,
    created    datetime,
    updated    datetime default getdate()
)
go

create unique index basic_trades_name_uindex
    on basic_trades (name)
go

create table roles
(
    id             int identity
        primary key,
    user_role_type varchar(50) not null,
    created        datetime    not null,
    updated        datetime default getdate()
)
go

create table users
(
    id              int identity
        primary key,
    user_role_id    int                not null
        constraint users_roles_id_fk
            references roles,
    username        varchar(200)       not null,
    password        varchar(256)       not null,
    email           varchar(256)       not null,
    password_token  varchar(200),
    token_expires   datetime,
    deletion_reason varchar(200),
    is_active       bit      default 1 not null,
    is_deleted      bit      default 0 not null,
    created         datetime,
    updated         datetime default getdate()
)
go

create table profiles
(
    id         int identity
        primary key,
    user_id    int                not null
        constraint profiles_users_id_fk
            references users,
    image_url  varchar(200),
    mobile     varchar(50),
    company    varchar(200),
    first_name varchar(200),
    last_name  varchar(200),
    is_active  bit      default 1 not null,
    is_deleted bit      default 0 not null,
    created    datetime,
    updated    datetime default getdate()
)
go

create table projects
(
    id                     int identity
        primary key,
    manager_id             int                not null
        constraint fk_project_manager_id
            references users,
    creator_id             int                not null
        constraint fk_project_creator_id
            references users,
    name                   varchar(200),
    total_item_breakdown   float,
    contractor_total_claim float,
    serial_no              varchar(50),
    details                varchar(200),
    total_contract_value   float,
    quantity_surveyor      varchar(200),
    notes                  varchar(200),
    is_active              bit      default 1 not null,
    is_deleted             bit      default 0 not null,
    created                datetime,
    updated                datetime default getdate()
)
go

create table trades
(
    id             int identity
        primary key,
    basic_trade_id int                not null
        constraint trades_basic_trades_id_fk
            references basic_trades,
    surveyor_id    int                not null
        constraint trades_users_id_fk
            references users,
    project_id     int                not null
        constraint trades_projects_id_fk
            references projects,
    floor_level    varchar(20)        not null,
    work_desc      varchar(200),
    item_breakdown float,
    tempcheck      bit,
    is_active      bit      default 1 not null,
    is_deleted     bit      default 0 not null,
    created        datetime,
    updated        datetime default getdate()
)
go

create table claims
(
    id                 int identity
        primary key,
    trade_id           int                not null
        constraint claims_trades_id_fk
            references trades,
    user_id            int                not null
        constraint claims_users_id_fk
            references users,
    basic_trade_id     int                not null
        constraint claims_basic_trades_id_fk
            references basic_trades,
    total_amount       float,
    claimed_amount     float,
    previous_claimed   float,
    amount_due         float,
    cost_to_completed  float,
    claim_number       varchar(200),
    claim_period       varchar(50),
    action_claim       bit,
    old_claimed_amount float,
    claim_percentage   float,
    is_active          bit      default 1 not null,
    is_deleted         bit      default 0 not null,
    created            datetime,
    updated            datetime default getdate()
)
go

create table claim_histories
(
    id               int identity
        primary key,
    trade_id         int                not null
        constraint claim_histories_trades_id_fk
            references trades,
    claim_id         int                not null
        constraint claim_histories_claims_id_fk
            references claims,
    profile_id       int                not null,
    previous_claimed float,
    is_active        bit      default 1 not null,
    is_deleted       bit      default 0 not null,
    created          datetime,
    updated          datetime default getdate()
)
go

create unique index users_email_uindex
    on users (email)
go

