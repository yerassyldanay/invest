
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

insert into emails (address, sent_code, deadline, verified) values
                ('invest.dept.spk@inbox.ru', '', null, true),
                ('manager.spk@inbox.ru', '', null, true),
                ('investor.spk@inbox.ru', '', null, true),
                ('financier.spk@inbox.ru', '', null, true),
                ('lawyer.spk@inbox.ru', '', null, true);
--                 (2, 'finans.dept.spk@inbox.ru', '', '', null, true);

insert into phones (ccode, number, sent_code, verified) values
                ('+7', '7001002030', '', true),
                ('+7', '7001002031', '', true),
                ('+7', '7001002032', '', true),
                ('+7', '7001002033', '', true),
                ('+7', '7001002034', '', true),
                ('+7', '7001002035', '', true);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- delete from admins where 1=1;
delete from emails where True;
delete from phones where True;

