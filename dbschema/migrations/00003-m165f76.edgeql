CREATE MIGRATION m165f76wyzjqgcirbfqu663qzzcu6vkvreysewszptgoh7zd4gmunq
    ONTO m1yjpzdsvmnakdllamcwxltqzk5yol3iujemkh5x5om4bktzyhmwaa
{
  ALTER TYPE default::Game {
      DROP CONSTRAINT std::expression ON ((NOT (__subject__.is_running) OR (__subject__.is_template OR (((((((__subject__.description = '') OR (__subject__.starting_message = '')) OR (__subject__.scenario = '')) OR (__subject__.objectives = '')) OR NOT (EXISTS (__subject__.skills))) OR NOT (EXISTS (__subject__.characteristics))) OR NOT (EXISTS (__subject__.relationship))))));
  };
  ALTER TYPE default::Game {
      CREATE CONSTRAINT std::expression ON ((NOT (__subject__.is_running) OR (NOT (__subject__.is_template) AND (((((((__subject__.description != '') AND (__subject__.starting_message != '')) AND (__subject__.scenario != '')) AND (__subject__.objectives != '')) AND EXISTS (__subject__.skills)) AND EXISTS (__subject__.characteristics)) AND EXISTS (__subject__.relationship))))) {
          SET errmessage := 'game data not fully satisfied, cannot set is_running to true';
      };
  };
};
