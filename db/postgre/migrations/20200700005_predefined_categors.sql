
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');

select setval('categors_id_seq', (select coalesce(max(id), 0) as id from invest.public.categors) + 1);
insert into invest.public.categors (name)
    values ('спорт'), ('медицина'), ('арт'), ('развлечение');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

delete from categors where true;
