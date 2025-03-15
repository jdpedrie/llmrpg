CREATE MIGRATION m1d6krtbwmj65kvkx7cflelnr7swxmgbr4koelmgdqyz7lsdxmnrja
    ONTO m1qrrsrjop3dz73vwm2gw67mbpkcbyafu77pvvfkgzsb4spcbbhb4q
{
  ALTER TYPE default::Game {
      CREATE LINK history := (.<game[IS default::History]);
  };
};
