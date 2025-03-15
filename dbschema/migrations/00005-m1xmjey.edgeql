CREATE MIGRATION m1xmjeyd2eil4i5yfngbfxxqtaes43t5lzczhdy7hhg6mg7memm6hq
    ONTO m1ddnv73qsleh3quiwdm2rtkcl55cqeswk42j4qv34do7xn5rvmm7q
{
  ALTER TYPE default::Game {
      ALTER PROPERTY create_time {
          CREATE REWRITE
              INSERT 
              USING (std::datetime_of_statement());
          SET REQUIRED USING (<std::datetime>{});
      };
      ALTER PROPERTY last_activity_time {
          CREATE REWRITE
              UPDATE 
              USING (std::datetime_of_statement());
          SET REQUIRED USING (<std::datetime>{});
      };
  };
  ALTER TYPE default::History {
      ALTER PROPERTY create_time {
          CREATE REWRITE
              INSERT 
              USING (std::datetime_of_statement());
      };
  };
};
