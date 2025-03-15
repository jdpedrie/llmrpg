CREATE MIGRATION m1qrrsrjop3dz73vwm2gw67mbpkcbyafu77pvvfkgzsb4spcbbhb4q
    ONTO m1xmjeyd2eil4i5yfngbfxxqtaes43t5lzczhdy7hhg6mg7memm6hq
{
  ALTER TYPE default::History {
      ALTER PROPERTY create_time {
          SET REQUIRED USING (<std::datetime>{});
      };
  };
};
