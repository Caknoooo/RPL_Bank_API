CREATE DATABASE rekening_service;

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE OR REPLACE FUNCTION uuid_generate_v4()
RETURNS uuid
AS $$
BEGIN
  RETURN ('' || md5(random()::text || clock_timestamp()::text))::uuid;
END;
$$
LANGUAGE plpgsql;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Nasabah (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nama VARCHAR(100) NOT NULL,
    nomor_ktp VARCHAR(20),
    tempat_lahir VARCHAR(100) NOT NULL,
    tanggal_lahir VARCHAR(100) NOT NULL,
    alamat_asal TEXT NOT NULL,
    no_hp VARCHAR(30) NOT NULL,
    email VARCHAR(50) NOT NULL,
    jenis_kelamin VARCHAR(20) NOT NULL,
    pekerjaan VARCHAR(100) NOT NULL,
    alamat_pekerjaan TEXT,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE Rekening (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    jenis_akun VARCHAR(30) NOT NULL,
    mata_uang VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    nasabah_id UUID REFERENCES nasabah(id)
);

DROP DATABASE rekening_service;