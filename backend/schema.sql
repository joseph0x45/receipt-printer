create table if not exists products (
  id text not null primary key,
  name text not null,
  price integer not null,
  bulk_price integer not null,
  is_available boolean not null
);
