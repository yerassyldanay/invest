
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- email & phone
create unique index if not exists make_email_address_unique on spkdb.public.emails (address);
create unique index if not exists make_phone_number_unique on spkdb.public.phones (ccode, number);
create index if not exists phone_better_login on phones ((spkdb.public.phones.ccode || spkdb.public.phones.number));

-- documents
create index doc_uri on spkdb.public.documents (uri);
create index doc_status on spkdb.public.documents (status);
create index doc_project on spkdb.public.documents (step, project_id);

-- gantt table
create index gantt_project_id on gantas (project_id);
create index gantt_step on gantas (step);

-- notifications
create index ni_address on notification_instances (to_address);
create index ni_project_id on notifications (project_id);
create index ni_created on notifications (created);

-- project
create unique index if not exists make_project_unique on projects (name);
create index if not exists project_status on projects (status);
create index if not exists project_step on projects (step asc);
create index if not exists project_status_and_step on projects(status, step asc);
create index if not exists project_created on projects(created);

-- smtp
create index if not exists smtp_sorted on smtp_servers (last_used);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

drop index make_email_address_unique;
drop index make_phone_number_unique;
drop index phone_better_login;

drop index doc_uri, doc_status, doc_project;
drop index gantt_project_id;
drop index gantt_step;

drop index ni_address, ni_created, ni_project_id;

drop index project_status, project_step;

drop index smtp_sorted;

