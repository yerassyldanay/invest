
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');

insert into invest.public.categors (id, kaz, rus, eng) values
            (1, 'құрылыс', 'строительство', 'construction'),
            (2, 'құрылыс ЖТҮ', 'строительство МЖК', 'YRH'),
            (3, 'өндіріс', 'производство', 'production'),
            (4, 'ауыл шаруашылығы', 'сельское хозяйство', 'agriculture'),
            (5, 'қызметтер', 'услуги', 'services');
select setval('categors_id_seq', (select coalesce(max(id), 0) as id from invest.public.categors) + 1);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

delete from categors where true;
select e.* from users u join roles r on r.id = u.role_id join emails e on u.email_id = e.id where r.name = 'admin';