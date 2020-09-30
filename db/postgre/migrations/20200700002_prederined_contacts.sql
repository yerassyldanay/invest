
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');

insert into emails (id, address, sent_code, sent_hash, deadline, verified) values
                (1, 'invest.dept.spk@inbox.ru', '', '', null, true),
                (2, 'manager.spk@inbox.ru', '', '', null, true),
                (3, 'investor.spk@inbox.ru', '', '', null, true),
                (4, 'financier.spk@inbox.ru', '', '', null, true),
                (5, 'lawyer.spk@inbox.ru', '', '', null, true);
--                 (2, 'finans.dept.spk@inbox.ru', '', '', null, true);

insert into phones (id, ccode, number, sent_code, verified) values
                (1, '+7', '7001002030', '', true),
                (2, '+7', '7001002031', '', true),
                (3, '+7', '7001002032', '', true),
                (4, '+7', '7001002033', '', true),
                (5, '+7', '7001002034', '', true),
                (6, '+7', '7001002035', '', true);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- delete from admins where 1=1;
delete from emails where True;
delete from phones where True;

