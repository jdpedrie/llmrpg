CREATE MIGRATION m1hfnd6iwtaw3b2hk772njrbc7syn7ngbmfmrlcrkfklniiojddfvq
    ONTO m1vkunl5as6j7zczig2772bfznxlqtq6fevppuzydq65jmmufs3ina
{
  ALTER TYPE default::Game {
      CREATE PROPERTY create_time: std::datetime;
  };
};
