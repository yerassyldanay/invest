
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

insert into roles(name, description) values
                                                ('admin', 'админ | admin'),
                                                ('manager', 'менеджер | manager'),
                                                ('investor', 'инвестор | investor'),
                                                ('expert', 'эксперт | expert');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- delete from admins where 1=1;
delete from roles where True;
delete from roles_permissions where True;

