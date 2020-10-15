
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');
insert into smtp_servers (host, port, username, password, last_used)
    values('smtp-relay.sendinblue.com', 587, 'yerassyl.danay@nu.edu.kz', 'pkhRjzw93cBFI6NE', now());

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
delete from smtp_servers where true;

