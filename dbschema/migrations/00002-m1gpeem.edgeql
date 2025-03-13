CREATE MIGRATION m1gpeemmmvxf2bhqzlkoobribjt4kmnjkjbbgcglnll2suesw3sifa
    ONTO m1czt2uersrgcz5qrmmhsxuybmwxt4bcpepxehw2e3bxykprkqu5ga
{
  ALTER TYPE default::Game {
      ALTER PROPERTY background_message {
          RENAME TO scenario;
      };
  };
  ALTER TYPE default::Game {
      ALTER PROPERTY initial_message {
          RENAME TO starting_message;
      };
  };
  ALTER TYPE default::Game {
      CREATE PROPERTY is_template: std::bool;
  };
  ALTER TYPE default::Game {
      CREATE PROPERTY is_running: std::bool;
      DROP PROPERTY template;
  };
};
