CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create schema if not exists hackerevolution;

create table if not exists hackerevolution.organizations
(
    id_org uuid         not null default uuid_generate_v4(),
    name   varchar(255) not null,
    primary key (id_org)
);

create table if not exists hackerevolution.type_group
(
    id_type serial       not null,
    type    varchar(255) not null,
    primary key (id_type)
);

create table if not exists hackerevolution.groups
(
    id_group   uuid         not null default uuid_generate_v4(),
    name_group varchar(255) not null,
    year       smallint,
    id_type    int,
    primary key (id_group),
    foreign key (id_type) references hackerevolution.type_group (id_type)
);

create table if not exists hackerevolution.org_group
(
    id_org_group uuid not null default uuid_generate_v4(),
    id_org       uuid not null,
    id_group     uuid not null,
    primary key (id_org_group),
    foreign key (id_org) references hackerevolution.organizations (id_org),
    foreign key (id_group) references hackerevolution.groups (id_group)
);

create table if not exists hackerevolution.role
(
    id_role   serial       not null,
    type_role varchar(255) not null,
    primary key (id_role)
);

create table if not exists hackerevolution.users
(
    id_user      uuid         not null default uuid_generate_v4(),
    fio          varchar(255) not null,
    login        varchar(255) not null,
    password     varchar(255) not null,
    id_org       uuid,
    id_org_group uuid,
    epoints      bigint,
    exp          bigint,
    id_role      int          not null,
    primary key (id_user),
    foreign key (id_org) references hackerevolution.organizations (id_org),
    foreign key (id_org_group) references hackerevolution.org_group (id_org_group),
    foreign key (id_role) references hackerevolution.role (id_role)
);

create table if not exists hackerevolution.features
(
    id_feature    uuid         not null default uuid_generate_v4(),
    name_feature  varchar(255) not null,
    price_feature int                   default 10,
    primary key (id_feature)
);

create table if not exists hackerevolution.feature_user
(
    id_feature_user uuid not null default uuid_generate_v4(),
    id_feature      uuid not null,
    id_user         uuid not null,
    is_on           bool not null default false

);

create table if not exists hackerevolution.files
(
    id_file   uuid    not null default uuid_generate_v4(),
    file_path varchar not null,
    primary key (id_file)
);

create table if not exists hackerevolution.courses
(
    id_course             uuid not null default uuid_generate_v4(),
    id_photo_file         uuid,
    type_course boolean default true,
    name_of_course        varchar(255),
    description_of_course varchar,
    primary key (id_course)
);



create table if not exists hackerevolution.course_roles
(
    id_course_roles uuid    not null default uuid_generate_v4(),
    type_role       varchar not null,
    primary key (id_course_roles)
);

create table if not exists hackerevolution.user_courses
(
    id      uuid not null default uuid_generate_v4(),
    id_c    uuid not null,
    id_c_r  uuid not null,
    pinned  boolean       default false,
    id_user uuid not null,
    primary key (id),
    foreign key (id_c) references hackerevolution.courses (id_course),
    foreign key (id_c_r) references hackerevolution.course_roles (id_course_roles),
    foreign key (id_user) references hackerevolution.users (id_user)
);

create table if not exists hackerevolution.user_note_course
(
    id_note uuid not null default uuid_generate_v4(),
    id_user uuid not null,
    id_course uuid not null,
    note varchar null,
    primary key (id_note),
    foreign key (id_user) references hackerevolution.users (id_user),
    foreign key (id_course) references hackerevolution.courses (id_course)
);

create table if not exists hackerevolution.tasks
(
    id_task  uuid not null default uuid_generate_v4(),
    text     varchar(255),
    max_mark int,
    primary key (id_task)
);

create table if not exists hackerevolution.posts
(
    id_post    uuid                        not null default uuid_generate_v4(),
    id_course  uuid                        not null,
    id_task    uuid                                 default null,
    text       varchar,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC+3'),
    updated_at TIMESTAMP WITHOUT TIME ZONE NULL,
    primary key (id_post),
    foreign key (id_course) references hackerevolution.courses (id_course),
    foreign key (id_task) references hackerevolution.tasks (id_task)
);

create table if not exists hackerevolution.settings
(
    id_setting uuid not null default uuid_generate_v4(),
    type       varchar(255),
    primary key (id_setting)
);

create table if not exists hackerevolution.course_settings
(
    id_course_settings uuid not null default uuid_generate_v4(),
    id_course          uuid not null,
    id_setting         uuid not null,
    primary key (id_setting),
    foreign key (id_course) references hackerevolution.courses (id_course),
    foreign key (id_setting) references hackerevolution.settings (id_setting)
);

create table if not exists hackerevolution.posts_files
(
    id_post_file uuid not null default uuid_generate_v4(),
    id_p         uuid not null,
    id_f         uuid not null,
    id_u         uuid not null,
    primary key (id_post_file),
    foreign key (id_p) references hackerevolution.posts (id_post),
    foreign key (id_f) references hackerevolution.files (id_file),
    foreign key (id_u) references hackerevolution.users (id_user)
);

create table if not exists hackerevolution.task_user
(
    id_t_u uuid not null default uuid_generate_v4(),
    id_t   uuid not null,
    id_u   uuid not null,
    primary key (id_t_u),
    foreign key (id_t) references hackerevolution.tasks (id_task),
    foreign key (id_u) references hackerevolution.users (id_user)
);

create table if not exists hackerevolution.task_user_file
(
    id_t_u_f uuid not null default uuid_generate_v4(),
    id_t_u   uuid not null,
    id_f     uuid not null,
    primary key (id_t_u_f),
    foreign key (id_t_u) references hackerevolution.task_user (id_t_u),
    foreign key (id_f) references hackerevolution.files (id_file)
);

Create table if not exists hackerevolution.tokens
(
    id_user     uuid,
    token       varchar,
    update_time timestamp,
    primary key (id_user)
);
