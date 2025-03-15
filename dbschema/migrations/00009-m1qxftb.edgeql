CREATE MIGRATION m1qxftbtekmy2p2o5732l2enpacflijnx3oaubemcj4z3ybkcu4sta
    ONTO m1wlv5ly23x33u3zyt5pwq4l7s36ve5ht44qpzlzzinry2lwrgg33q
{
  ALTER TYPE default::Game {
      ALTER PROPERTY last_activity_time {
          RESET OPTIONALITY;
      };
  };
};
