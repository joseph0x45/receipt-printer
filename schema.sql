create table if not exists products (
  id text not null,
  name text not null,
  media text not null,
  price integer not null,
  bulk_price integer not null,
  in_stock integer not null
);

create table if not exists sessions(
  id text not null,
  active boolean not null default true
);
