create table if not exists products (
  id text not null,
  name text not null,
  media text not null,
  price number not null,
  bulk_price number not null,
  in_stock number not null
);
