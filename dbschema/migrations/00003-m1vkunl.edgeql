CREATE MIGRATION m1vkunl5as6j7zczig2772bfznxlqtq6fevppuzydq65jmmufs3ina
    ONTO m1gpeemmmvxf2bhqzlkoobribjt4kmnjkjbbgcglnll2suesw3sifa
{
  ALTER TYPE default::Game {
      CREATE PROPERTY last_activity_time: std::datetime;
  };
};
