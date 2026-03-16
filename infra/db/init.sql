-- Schema inicial do projeto (executa somente no 1º start do volume)

CREATE TABLE IF NOT EXISTS tblservices (
  id             SERIAL PRIMARY KEY,
  name           TEXT NOT NULL,
  price_cent     INTEGER NOT NULL CHECK (price_cent >= 0),
  time_minutes   INTEGER NOT NULL CHECK (time_minutes >= 0),
  is_maintenance BOOLEAN NOT NULL DEFAULT FALSE
);

-- Seed opcional (remova se não quiser dados padrão)
INSERT INTO tblservices (name, price_cent, time_minutes, is_maintenance)
VALUES
  ('Corte', 3000, 30, false),
  ('Barba', 2500, 20, false),
  ('Manutenção', 5000, 60, true)
ON CONFLICT DO NOTHING;

