-- Crear ENUM tipo de persona
CREATE TYPE person_type AS ENUM ('NATURAL', 'JURIDICA');
CREATE TYPE natural_person_type AS ENUM ('DNI', 'PASAPORTE','CARNET DE EXTRANJERIA');

-- Tabla person 
CREATE TABLE person (
    id UUID PRIMARY KEY NOT NULL,
    type person_type NOT NULL,
    status BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP NOT NULL
);

-- Tabla natural_person (datos personales naturales)
CREATE TABLE natural_person (
    id UUID PRIMARY KEY NOT NULL, -- Mismo ID que person
    document_type natural_person_type NOT NULL,
    document_number VARCHAR(50),
    first_name VARCHAR(100),
    last_name_paternal VARCHAR(100),
    last_name_maternal VARCHAR(100),
    status BOOLEAN DEFAULT TRUE, -- Soft delete controlado
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_natural_person FOREIGN KEY (id) REFERENCES person(id) ON DELETE CASCADE
);

-- Tabla juridic_person (datos de persona jurídica)
CREATE TABLE juridic_person (
    id UUID PRIMARY KEY NOT NULL, -- Mismo ID que person
    ruc VARCHAR(20),
    legal_name VARCHAR(255) NOT NULL,
    commercial_name VARCHAR(255),
    address VARCHAR(255),
    district VARCHAR(100),
    province VARCHAR(100),
    department VARCHAR(100),
    country VARCHAR(100),
    has_constituted BOOLEAN NOT NULL,
    status BOOLEAN DEFAULT TRUE, -- Soft delete controlado
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_juridic_person FOREIGN KEY (id) REFERENCES person(id) ON DELETE CASCADE
);
