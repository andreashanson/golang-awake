-- migrate:up
CREATE TABLE influencers (
    id serial PRIMARY KEY,
    name text NOT NULL,
    lastname text NOT NULL,
    email text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);

-- migrate:down
drop table influencers
