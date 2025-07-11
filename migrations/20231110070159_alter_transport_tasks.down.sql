alter table public.transport_tasks
    add column trailer_id uuid references trailers (id);