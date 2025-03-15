CREATE MIGRATION m1wlv5ly23x33u3zyt5pwq4l7s36ve5ht44qpzlzzinry2lwrgg33q
    ONTO m1d6krtbwmj65kvkx7cflelnr7swxmgbr4koelmgdqyz7lsdxmnrja
{
  ALTER TYPE default::Game {
      DROP LINK history;
  };
};
