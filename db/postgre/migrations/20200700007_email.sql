
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');

insert into invest.public.emails (address, sent_code, deadline) values
            ('yerassyl.danay@mail.ru', '7777', now() + '100 years');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

delete from invest.public.emails where invest.public.emails.address = 'yerassyl.danay@mail.ru';
