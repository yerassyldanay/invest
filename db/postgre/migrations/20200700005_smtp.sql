
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
insert into smtp_servers (host, port, username, password, last_used)
    values('smtp-relay.sendinblue.com', 587, 'yerassyl.danay@nu.edu.kz', 'pkhRjzw93cBFI6NE', now() + '-10 days'),
           ('smtp-relay.sendinblue.com', 587, 'yerassyl.danay.nu@gmail.com', 'JLZRatsrpn9BANDT', now() + '-10 days');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
delete from smtp_servers where true;

