create table categors
(
    id  bigserial primary key,
    kaz text not null unique,
    rus text not null unique,
    eng text not null unique
);

create table comments
(
    id bigserial primary key,
    body text default ''::text,
    user_id    bigint,
    project_id bigint,
    status     text      not null,
    created    timestamp with time zone default now()
);

create table costs
(
    id  bigserial not null primary key,
    project_id  bigint unique,
    building_repair_investor      integer,
    building_repair_involved      integer,
    technology_equipment_investor integer,
    technology_equipment_involved integer,
    working_capital_investor      integer,
    working_capital_involved      integer,
    other_cost_investor           integer,
    other_cost_involved           integer
);

create table documents
(
    id            bigserial not null
        constraint documents_pkey
            primary key,
    kaz           text      not null,
    rus           text      not null,
    eng           text      not null,
    uri           text                     default ''::text,
    modified      timestamp with time zone default now(),
    created       timestamp with time zone default now(),
    status        text                     default 'new_one'::text,
    step          integer                  default 1,
    is_additional boolean                  default false,
    project_id    bigint,
    responsible   text
);

-- alter table documents
--     owner to spkuser;

create table emails
(
    id        bigserial   primary key,
    address   text unique,
    verified  boolean default false,
    sent_code text,
    deadline  timestamp with time zone
);

-- alter table emails
--     owner to spkuser;

create table finances
(
    id                     bigserial   primary key,
    project_id             bigint unique,
    total_income           integer,
    total_production       integer,
    production_cost        integer,
    operational_profit     integer,
    settlement_obligations integer,
    other_cost             integer,
    pure_profit            integer,
    taxes                  integer
);

create table forget_passwords
(
    email_address text,
    code          text,
    deadline      timestamp with time zone
);

create table gantas
(
    id               bigserial   primary key,
    is_additional    boolean                  default false,
    project_id       bigint,
    kaz              text                     default ''::text,
    rus              text                     default ''::text,
    eng              text                     default ''::text,
    start_date       timestamp with time zone default now(),
    duration_in_days bigint,
    deadline         timestamp with time zone,
    ganta_parent_id  bigint,
    step             integer                  default 1,
    status           text,
    is_done          boolean                  default false,
    responsible      text                     default 'spk'::text,
    is_doc_check     boolean                  default false,
    not_to_show      boolean                  default false
);

create table organizations
(
    id      bigserial primary key,
    lang    text,
    bin     text,
    name    text      not null,
    fio     text,
    regdate timestamp with time zone,
    address text      not null
);

create table permissions
(
    id          bigserial  primary key,
    name        text,
    description text
);

create table phones
(
    id        bigserial   primary key,
    ccode     text,
    number    text,
    sent_code text,
    verified  boolean
);

create table projects_users
(
    project_id bigint not null,
    user_id    bigint not null,
    created    timestamp with time zone default now(),
    constraint projects_users_pkey
        primary key (project_id, user_id)
);

create table projects_categors
(
    project_id bigint not null,
    categor_id bigint not null,
    constraint projects_categors_pkey
        primary key (project_id, categor_id)
);

create table projects
(
    id                  bigserial  primary key,
    name                text,
    description         text                     default ''::text,
    info                text,
    employee_count      integer,
    email               text                     default ''::text,
    phone_number        text                     default ''::text,
    organization_id     bigint,
    offered_by_id       bigint    not null,
    offered_by_position text      not null,
    status              text                     default 'pending_admin'::text,
    step                integer,
    land_plot_from      text                     default 'investor'::text,
    land_area           integer,
    land_address        text,
    is_manager_assigned boolean                  default false,
    created             timestamp with time zone default now(),
    deleted             timestamp with time zone
);

create table roles_permissions
(
    role_id       bigint not null,
    permission_id bigint not null,
    constraint roles_permissions_pkey
        primary key (role_id, permission_id)
);

create table roles
(
    id          bigserial  primary key,
    name        text,
    description text
);

create table smtp_servers
(
    id        bigserial  primary key,
    host      text,
    port      integer,
    username  text,
    password  text,
    last_used timestamp with time zone
);


create table smtp_headers
(
    id             bigserial  primary key,
    smtp_server_id bigint,
    key            text,
    value          text
);

create table users
(
    id              bigserial  primary key,
    password        text,
    fio             text,
    role_id         bigint,
    email_id        bigint,
    phone_id        bigint,
    verified        boolean                  default false,
    organization_id bigint                   default 0,
    blocked         boolean                  default false,
    created         timestamp with time zone default now()
);

create table notifications
(
    id           bigserial  primary key,
    from_address text,
    project_id   bigint                   default 0,
    html         text,
    plain        text,
    created      timestamp with time zone default now()
);

create table notification_instances
(
    to_address      text,
    notification_id bigint
);

