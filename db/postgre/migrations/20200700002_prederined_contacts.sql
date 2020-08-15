
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');

insert into emails (id, address, sent_code, sent_hash, deadline, verified) values
                (1, 'yerassyl.danay@nu.edu.kz', '', '', null, true),
                (2, 'yerassyl.danay.nu@gmail.com', '', '', null, true),
                (3, 'yerassyl.danay@mail.ru', '', '', null, true);

insert into phones (id, ccode, number, sent_code, verified) values
                (1, '+7', '7001002030', '', true),
                (2, '+7', '7058686509', '', true),
                (3, '+7', '7058686509', '', true);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- delete from admins where 1=1;
delete from emails where True;
delete from phones where True;

