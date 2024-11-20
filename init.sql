drop TABLE t_user if exist;
create table t_user
(
    id serial not null
        primary key
        unique,
    user_id    varchar(45),
    email      varchar(45),
    password   varchar(45),
    salt varchar(45),
    create_ts timestamp default current_timestamp,
    update_ts timestamp default current_timestamp
);

comment on column t_user.id is 'id，auto increment';
comment on column t_user.user_id is 'user id,unique';
comment on column t_user.email is 'user email';
comment on column t_user.password is 'user password';
comment on column t_user.salt is 'user salt';
comment on column t_user.create_ts is 'create timestamp';
comment on table t_user is '用户表';

create table t_wallet
(
    id        serial not null
        primary key
        unique,
    user_id   varchar(45),
    account   varchar(45),
    amount    DECIMAL(10, 9),
    create_ts TIMESTAMP,
    update_ts TIMESTAMP
);

comment on column t_wallet.id is 'id,auto increment';

comment
on column t_wallet.user_id is 'user id';

comment
on column t_wallet.account is 'wallet account';

comment
on column t_wallet.amount is 'amount';

comment
on column t_wallet.create_ts is 'create timestamp';

comment
on column t_wallet.update_ts is 'update timestamp';
create table t_transfer_log
(
    id  serial not null
        primary key
        unique,
    from_user_id    varchar(45),
    to_user_id      varchar(45),
    amount          decimal(10, 9),
    transfer_status varchar(1),
    transfer_type varchar(1),
    create_ts       timestamp
);

comment
on table t_transfer_log is 'transfer log';

comment
on column t_transfer_log.id is 'id,auto increment';

comment
on column t_transfer_log.from_user_id is 'from user id';

comment
on column t_transfer_log.to_user_id is 'to user id';

comment
on column t_transfer_log.amount is 'transfer amount';

comment
on column t_transfer_log.transfer_status is 'transfer status 1.success 0.fail';
comment on column t_transfer_log.transfer_type is '1. deposit 2.withdraw 3.send in 4.send out';
comment
on column t_transfer_log.create_ts is 'create timestamp';









