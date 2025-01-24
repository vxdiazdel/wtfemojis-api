CREATE TABLE emojis (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  value TEXT NOT NULL
);