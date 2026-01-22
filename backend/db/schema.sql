-- Create the table for storing Amiibo items
create table if not exists public.amiibos (
  id uuid not null default gen_random_uuid (),
  name text not null,
  image_url text null,
  series text not null default 'Super Smash Bros',
  release_date date null,
  created_at timestamp with time zone not null default now(),
  constraint amiibos_pkey primary key (id)
);

-- Enable Row Level Security (RLS)
alter table public.amiibos enable row level security;

-- Create a policy that allows anyone to read
create policy "Public read access"
  on public.amiibos
  for select
  using (true);

-- Create a policy that allows only authenticated users to insert/update/delete 
-- (Assuming the scripts run with a service role or user is authenticated)
create policy "Authenticated users can modify"
  on public.amiibos
  for all
  using (auth.role() = 'authenticated');
