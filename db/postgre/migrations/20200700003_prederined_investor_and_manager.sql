
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');

insert into users (id, username, password, fio, role_id, email_id, phone_id, verified) values
    (1, 'admin', '$2a$11$5NXignmyT1RzEz7JHdMurONfnxasb09NdRK2TSUuSHjI2TKmbAXzS', 'Default Admin', 1, 1, 1, true),
    (2, 'manager', '$2a$11$6gqEIO0Oy9aBwcyhO.xbdOO/HMyyc0VyyteUuRDpd0LIl/2142ZlW', 'Default Manager', 2, 2, 2, true),
    (3, 'investor', '$2a$11$DnmyFsSY54npwBOEOMAAaelSFrZBMqK0c/v4D0kSJNf83JIFAp4Cu', 'Default Investor', 3, 3, 3, true);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
delete from users where 1=1;

