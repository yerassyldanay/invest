
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');

insert into finances (id) values(1);
insert into costs (id) values(1);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

delete from invest.public.finances where True;
delete from costs where True;

