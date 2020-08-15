
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');

insert into users (id, username, password, fio, position, role_id, email_id, phone_id, verified) values
    (10, 'admin', '$2a$11$5NXignmyT1RzEz7JHdMurONfnxasb09NdRK2TSUuSHjI2TKmbAXzS', 'Default Admin', 'adminstrator', 8, 9, 10, true);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
delete from users where 1=1;

