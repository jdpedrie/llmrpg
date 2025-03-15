CREATE MIGRATION m1ddnv73qsleh3quiwdm2rtkcl55cqeswk42j4qv34do7xn5rvmm7q
    ONTO m165f76wyzjqgcirbfqu663qzzcu6vkvreysewszptgoh7zd4gmunq
{
  ALTER TYPE default::Game {
      ALTER CONSTRAINT std::expression ON ((NOT (__subject__.is_running) OR (NOT (__subject__.is_template) AND (((((((__subject__.description != '') AND (__subject__.starting_message != '')) AND (__subject__.scenario != '')) AND (__subject__.objectives != '')) AND EXISTS (__subject__.skills)) AND EXISTS (__subject__.characteristics)) AND EXISTS (__subject__.relationship))))) {
          SET errmessage := 'game is a template or data not fully satisfied; cannot set is_running to true';
      };
  };
};
