
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');

insert into roles(id, name, description) values
                                                (1, 'admin', 'админ | admin'),
                                                (2, 'manager', 'менеджер | manager'),
                                                (3, 'investor', 'инвестор | investor'),
                                                (4, 'expert', 'эксперт | expert');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- delete from admins where 1=1;
delete from roles where True;
delete from roles_permissions where True;

