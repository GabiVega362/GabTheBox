CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE DOMAIN EMAIL AS VARCHAR(255) CHECK (VALUE ~ '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$');
CREATE DOMAIN PORT AS INTEGER CHECK (VALUE > 0 AND VALUE < 65536);
CREATE DOMAIN USERNAME AS VARCHAR(30) CHECK (VALUE ~ '^[a-zA-Z0-9_]+$');