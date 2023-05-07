create table mall_db.user
(
    id              bigint unsigned auto_increment
        primary key,
    created_at      datetime      null,
    updated_at      datetime      null,
    deleted_at      datetime      null,
    user_name       varchar(256)  null,
    email           varchar(256)  null,
    password_digest varchar(256)  null,
    nick_name       varchar(256)  null,
    status          varchar(256)  null,
    avatar          varchar(1000) null,
    money           varchar(256)  null,
    constraint user_name
        unique (user_name)
);

create table mall_db.product
(
    id             bigint unsigned auto_increment
        primary key,
    created_at     datetime             null,
    updated_at     datetime             null,
    deleted_at     datetime             null,
    name           varchar(255)         null,
    category_id    bigint unsigned      not null,
    title          varchar(256)         null,
    info           varchar(1000)        null,
    img_path       varchar(256)         null,
    price          varchar(256)         null,
    discount_price varchar(256)         null,
    on_sale        tinyint(1) default 0 null,
    num            bigint               null,
    boss_id        bigint unsigned      null,
    boss_name      varchar(256)         null,
    boss_avatar    varchar(256)         null
);


create table mall_db.address
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    user_id    bigint unsigned not null,
    name       varchar(20)     not null,
    phone      varchar(11)     not null,
    address    varchar(50)     not null
);

create table mall_db.admin
(
    id              bigint unsigned auto_increment
        primary key,
    created_at      datetime      null,
    updated_at      datetime      null,
    deleted_at      datetime      null,
    user_name       varchar(256)  null,
    password_digest varchar(256)  null,
    avatar          varchar(1000) null
);

create table mall_db.carousel
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    img_path   varchar(256)    null,
    product_id bigint unsigned not null
);

create table mall_db.cart
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    user_id    bigint unsigned null,
    product_id bigint unsigned not null,
    boss_id    bigint unsigned null,
    num        bigint unsigned null,
    max_num    bigint unsigned null,
    `check`    tinyint(1)      null
);

create table mall_db.category
(
    id            bigint unsigned auto_increment
        primary key,
    created_at    datetime     null,
    updated_at    datetime     null,
    deleted_at    datetime     null,
    category_name varchar(256) null
);

create table mall_db.favorite
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    user_id    bigint unsigned not null,
    product_id bigint unsigned not null,
    boss_id    bigint unsigned not null,
    constraint fk_favorite_boss
        foreign key (boss_id) references mall_db.user (id),
    constraint fk_favorite_product
        foreign key (product_id) references mall_db.product (id),
    constraint fk_favorite_user
        foreign key (user_id) references mall_db.user (id)
);

create table mall_db.notice
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime null,
    updated_at datetime null,
    deleted_at datetime null,
    text       text     null
);

create table mall_db.`order`
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    user_id    bigint unsigned not null,
    product_id bigint unsigned not null,
    boss_id    bigint unsigned not null,
    address_id bigint unsigned not null,
    num        bigint          null,
    order_num  bigint unsigned null,
    type       bigint unsigned null,
    money      double          null
);

create index idx_product_name
    on mall_db.product (name);

create table mall_db.product_img
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime        null,
    updated_at datetime        null,
    deleted_at datetime        null,
    product_id bigint unsigned not null,
    img_path   varchar(256)    null
);

create table mall_db.skill_product
(
    id          bigint unsigned auto_increment
        primary key,
    product_id  bigint unsigned not null,
    boss_id     bigint unsigned not null,
    title       varchar(256)    null,
    money       double          null,
    num         bigint          not null,
    custom_id   bigint unsigned null,
    custom_name varchar(256)    null
);
