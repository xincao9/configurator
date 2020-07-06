insert into user (`username`, `password`, `role`) values ("admin", "21232f297a57a5a743894a0e4a801fc3", 3);

insert into env (`name`) values ("DEV");
insert into env (`name`) values ("TEST");
insert into env (`name`) values ("PROD");

insert into user_env (`user_id`, `env_id`) values (1, 1);
insert into user_env (`user_id`, `env_id`) values (1, 2);
insert into user_env (`user_id`, `env_id`) values (1, 3);
