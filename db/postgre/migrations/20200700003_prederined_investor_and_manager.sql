
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

insert into users (password, fio, role_id, email_id, phone_id, verified) values
    ('$2a$11$87Lnp0qo3CJo2UdmVPJL9.bzTscq7gziBSCd6TrXrznslvisSZDcW', 'Глава / начальник инвестиционного департамента или финансового департамента', 1, 1, 1, true),
    ('$2a$11$6gqEIO0Oy9aBwcyhO.xbdOO/HMyyc0VyyteUuRDpd0LIl/2142ZlW', 'Менеджер (по умолчанию)', 2, 2, 2, true),
    ('$2a$11$DnmyFsSY54npwBOEOMAAaelSFrZBMqK0c/v4D0kSJNf83JIFAp4Cu', 'Инвестор (по умолчанию)', 3, 3, 3, true),
    ('$2a$11$5NXignmyT1RzEz7JHdMurONfnxasb09NdRK2TSUuSHjI2TKmbAXzS', 'Юрист (по умолчанию)', 4, 4, 4, true),
    ('$2a$11$nBr4Lt/aKraecTmkVN9SM.cWiHfjqfUNFMecm1mMUDtPQA8SShR9a', 'Финансист (по умолчанию)', 4, 5, 5, true);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
delete from users where True;

