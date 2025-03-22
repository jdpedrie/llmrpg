-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgvector";

-- Character attribute table for skills, characteristics, and relationships
CREATE TABLE character_attributes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    value SMALLINT NOT NULL,
    attribute_type TEXT NOT NULL, -- 'skill', 'characteristic', or 'relationship'
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT value_range CHECK (value >= -10 AND value <= 10)
);

-- Character table
CREATE TABLE characters (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    description TEXT,
    context TEXT[], -- array of strings for character context
    active BOOLEAN NOT NULL DEFAULT FALSE,
    main_character BOOLEAN NOT NULL DEFAULT FALSE,
    game_id UUID,  -- Foreign key to games table, added later with ALTER
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Games table
CREATE TABLE games (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    description TEXT,
    starting_message TEXT,
    scenario TEXT,
    objectives TEXT,
    skills TEXT[],
    characteristics TEXT[],
    relationship TEXT[],
    is_template BOOLEAN NOT NULL DEFAULT FALSE,
    is_running BOOLEAN NOT NULL DEFAULT FALSE,
    playthrough_start_time TIMESTAMPTZ,
    playthrough_end_time TIMESTAMPTZ,
    last_activity_time TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Add foreign key to characters table
ALTER TABLE characters
    ADD CONSTRAINT fk_characters_game
    FOREIGN KEY (game_id)
    REFERENCES games(id)
    ON DELETE CASCADE;

-- Inventory items table
CREATE TABLE inventory_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    game_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE
);

-- Character-attribute relationships
CREATE TABLE character_to_attributes (
    character_id UUID NOT NULL,
    attribute_id UUID NOT NULL,
    relationship_type TEXT NOT NULL, -- 'skill', 'characteristic', 'relationship'
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (character_id, attribute_id),
    FOREIGN KEY (character_id) REFERENCES characters(id) ON DELETE CASCADE,
    FOREIGN KEY (attribute_id) REFERENCES character_attributes(id) ON DELETE CASCADE
);

-- Game history table
CREATE TABLE history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    game_id UUID NOT NULL,
    text TEXT NOT NULL,
    choice TEXT NOT NULL,
    outcome TEXT NOT NULL,
    embedding vector(1536), -- Assuming text-embedding-3-small uses 1536 dimensions
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE
);

-- Game context table - stores unstructured context for semantic retrieval
CREATE TABLE game_contexts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    game_id UUID NOT NULL,
    content TEXT NOT NULL,
    embedding vector(1536), -- Vector embedding of the content for semantic search
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (game_id) REFERENCES games(id) ON DELETE CASCADE
);

-- Triggers for updated_at columns
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_games_modtime
    BEFORE UPDATE ON games
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

CREATE TRIGGER update_characters_modtime
    BEFORE UPDATE ON characters
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

CREATE TRIGGER update_character_attributes_modtime
    BEFORE UPDATE ON character_attributes
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

CREATE TRIGGER update_inventory_items_modtime
    BEFORE UPDATE ON inventory_items
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

CREATE TRIGGER update_game_contexts_modtime
    BEFORE UPDATE ON game_contexts
    FOR EACH ROW
    EXECUTE PROCEDURE update_modified_column();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS update_game_contexts_modtime ON game_contexts;
DROP TRIGGER IF EXISTS update_inventory_items_modtime ON inventory_items;
DROP TRIGGER IF EXISTS update_character_attributes_modtime ON character_attributes;
DROP TRIGGER IF EXISTS update_characters_modtime ON characters;
DROP TRIGGER IF EXISTS update_games_modtime ON games;
DROP FUNCTION IF EXISTS update_modified_column();

DROP TABLE IF EXISTS game_contexts;
DROP TABLE IF EXISTS history;
DROP TABLE IF EXISTS character_to_attributes;
DROP TABLE IF EXISTS inventory_items;
DROP TABLE IF EXISTS characters;
DROP TABLE IF EXISTS character_attributes;
DROP TABLE IF EXISTS games;

DROP EXTENSION IF EXISTS "pgvector";
DROP EXTENSION IF EXISTS "uuid-ossp";

-- +goose StatementEnd