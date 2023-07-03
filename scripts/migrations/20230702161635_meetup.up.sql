create table meetups(
  id text primary key,
  name varchar(250) not null,
  description text,
  user_id text references users (id) on delete cascade not null
);