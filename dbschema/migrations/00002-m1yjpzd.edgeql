CREATE MIGRATION m1yjpzdsvmnakdllamcwxltqzk5yol3iujemkh5x5om4bktzyhmwaa
    ONTO m1uv2zwvhjl7a4xwdjk2jsyffgkgpuummb6bd5ph7cmalckt4re4sq
{
  ALTER TYPE default::Character {
      ALTER LINK characteristics {
          ON SOURCE DELETE DELETE TARGET;
          RESET ON TARGET DELETE;
      };
      ALTER LINK relationship {
          ON SOURCE DELETE DELETE TARGET;
          RESET ON TARGET DELETE;
      };
      ALTER LINK skills {
          ON SOURCE DELETE DELETE TARGET;
          RESET ON TARGET DELETE;
      };
  };
  ALTER TYPE default::Game {
      ALTER LINK characters {
          ON SOURCE DELETE DELETE TARGET;
          RESET ON TARGET DELETE;
      };
  };
};
