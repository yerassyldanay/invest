
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');

insert into invest.public.forget_passwords (email_address, code, deadline)
    values ('invest.dept.spk@inbox.ru', '7777', now() + '100 years');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

delete from invest.public.forget_passwords where invest.public.forget_passwords.email_address = 'invest.dept.spk@inbox.ru';
