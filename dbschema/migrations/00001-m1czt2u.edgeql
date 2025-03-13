CREATE MIGRATION m1czt2uersrgcz5qrmmhsxuybmwxt4bcpepxehw2e3bxykprkqu5ga
    ONTO initial
{
  CREATE TYPE default::Character {
      CREATE PROPERTY characteristics: array<tuple<name: std::str, value: std::int16>>;
      CREATE PROPERTY relationship: array<tuple<name: std::str, value: std::int16>>;
      CREATE PROPERTY stats: array<tuple<name: std::str, value: std::int16>>;
      CREATE REQUIRED PROPERTY active: std::bool;
      CREATE PROPERTY context: array<std::str>;
      CREATE PROPERTY description: std::str;
      CREATE REQUIRED PROPERTY main_character: std::bool;
      CREATE REQUIRED PROPERTY name: std::str;
  };
  CREATE FUTURE simple_scoping;
  CREATE TYPE default::Game {
      CREATE MULTI LINK characters: default::Character;
      CREATE PROPERTY background_message: std::str;
      CREATE PROPERTY description: std::str;
      CREATE PROPERTY initial_message: std::str;
      CREATE REQUIRED PROPERTY name: std::str;
      CREATE PROPERTY objectives: std::str;
      CREATE PROPERTY playthrough_end_time: std::datetime;
      CREATE PROPERTY playthrough_start_time: std::datetime;
      CREATE PROPERTY template: std::bool;
  };
};
