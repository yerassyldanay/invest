
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- insert into admins (username, password, fio) values('admin', '$2a$11$dQ4HxX824pacAoqhXDQ/0em0ebug8gN6AETJU3HhbHZB1KPB0N5FW', 'default admin');

insert into roles(id, name, description) values
                                                (1, 'admin', 'глава / начальник инвестиционного отделения или финансового отделения'),
                                                (2, 'manager', 'менеджер'),
                                                (3, 'investor', 'инвестор'),
                                                (4, 'lawyer', 'юрист'),
                                                (5, 'financier', 'финансист');

insert into roles_permissions(role_id, permission_id) values(1, 1), (1, 2), (1, 7),
                                                            (2, 3), (2, 5), (2, 7),
                                                            (3, 3), (3, 4), (3, 7),
                                                            (4, 3), (4, 5), (4, 7),
                                                            (5, 3), (5, 5), (5, 7);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- delete from admins where 1=1;
delete from roles where True;
delete from roles_permissions where True;

