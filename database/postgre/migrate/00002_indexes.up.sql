
-- email & phone
create unique index if not exists make_email_address_unique on emails (address);
create unique index if not exists make_phone_number_unique on phones (ccode, number);
create index if not exists phone_better_login on phones ((phones.ccode || phones.number));

-- documents
create index doc_uri on documents (uri);
create index doc_status on documents (status);
create index doc_project on documents (step, project_id);

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

