CREATE MIGRATION m1uv2zwvhjl7a4xwdjk2jsyffgkgpuummb6bd5ph7cmalckt4re4sq
    ONTO initial
{
  CREATE EXTENSION pgvector VERSION '0.7';
  CREATE EXTENSION ai VERSION '1.0';
  CREATE FUTURE simple_scoping;
  CREATE TYPE default::CharacterAttribute {
      CREATE REQUIRED PROPERTY name: std::str;
      CREATE REQUIRED PROPERTY value: std::int16 {
          CREATE CONSTRAINT std::max_value(10);
          CREATE CONSTRAINT std::min_value(-10);
      };
  };
  CREATE TYPE default::Character {
      CREATE MULTI LINK characteristics: default::CharacterAttribute {
          ON TARGET DELETE DELETE SOURCE;
      };
      CREATE MULTI LINK relationship: default::CharacterAttribute {
          ON TARGET DELETE DELETE SOURCE;
      };
      CREATE MULTI LINK skills: default::CharacterAttribute {
          ON TARGET DELETE DELETE SOURCE;
      };
      CREATE REQUIRED PROPERTY active: std::bool {
          SET default := false;
      };
      CREATE PROPERTY context: array<std::str>;
      CREATE PROPERTY description: std::str;
      CREATE REQUIRED PROPERTY main_character: std::bool {
          SET default := false;
      };
      CREATE REQUIRED PROPERTY name: std::str;
  };
  CREATE TYPE default::Game {
      CREATE MULTI LINK characters: default::Character {
          ON TARGET DELETE DELETE SOURCE;
      };
      CREATE PROPERTY characteristics: array<std::str>;
      CREATE PROPERTY description: std::str;
      CREATE REQUIRED PROPERTY is_running: std::bool {
          SET default := false;
      };
      CREATE REQUIRED PROPERTY is_template: std::bool {
          SET default := false;
      };
      CREATE PROPERTY objectives: std::str;
      CREATE PROPERTY relationship: array<std::str>;
      CREATE PROPERTY scenario: std::str;
      CREATE PROPERTY skills: array<std::str>;
      CREATE PROPERTY starting_message: std::str;
      CREATE CONSTRAINT std::expression ON ((NOT (__subject__.is_running) OR (__subject__.is_template OR (((((((__subject__.description = '') OR (__subject__.starting_message = '')) OR (__subject__.scenario = '')) OR (__subject__.objectives = '')) OR NOT (EXISTS (__subject__.skills))) OR NOT (EXISTS (__subject__.characteristics))) OR NOT (EXISTS (__subject__.relationship)))))) {
          SET errmessage := 'game data not fully satisfied, cannot set is_running to true';
      };
      CREATE PROPERTY create_time: std::datetime;
      CREATE PROPERTY last_activity_time: std::datetime;
      CREATE REQUIRED PROPERTY name: std::str;
      CREATE PROPERTY playthrough_end_time: std::datetime;
      CREATE PROPERTY playthrough_start_time: std::datetime;
  };
  CREATE TYPE default::History {
      CREATE REQUIRED LINK game: default::Game;
      CREATE REQUIRED PROPERTY choice: std::str;
      CREATE REQUIRED PROPERTY text: std::str;
      CREATE DEFERRED INDEX ext::ai::index(embedding_model := 'text-embedding-3-small') ON (((.text ++ ' ') ++ .choice));
      CREATE PROPERTY create_time: std::datetime;
      CREATE REQUIRED PROPERTY outcome: std::str;
  };
};
