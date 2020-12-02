
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
insert into spkdb.public.categors (kaz, rus, eng) values
            ('құрылыс', 'строительство', 'construction'),
            ('құрылыс ЖТҮ', 'строительство МЖК', 'YRH'),
            ('өндіріс', 'производство', 'production'),
            ('ауыл шаруашылығы', 'сельское хозяйство', 'agriculture'),
            ('қызметтер', 'услуги', 'services');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

delete from categors where true;
