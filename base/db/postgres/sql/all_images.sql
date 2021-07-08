
create table all_images (
id uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
file_name text not null,
content_type text not null,
size INTEGER not NULL,
v text DEFAULT null,
uts timestamp default current_timestamp
);