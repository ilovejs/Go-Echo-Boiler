
INSERT INTO test2.dbo.roles (user_role_type, created, updated) VALUES ('Supervisor', '2019-07-22 02:07:26.163', '2019-07-22 02:05:35.503');
INSERT INTO test2.dbo.roles (user_role_type, created, updated) VALUES ('Banker', '2019-07-22 02:05:54.097', '2019-07-22 02:05:56.887');
INSERT INTO test2.dbo.roles (user_role_type, created, updated) VALUES ('Quantity Surveyor', '2019-07-22 02:06:20.123', '2019-07-22 02:06:24.027');
INSERT INTO test2.dbo.roles (user_role_type, created, updated) VALUES ('System', '2019-07-22 02:07:09.897', '2019-07-22 02:07:13.397');
INSERT INTO test2.dbo.roles (user_role_type, created, updated) VALUES ('Contractor', '2019-07-22 02:07:22.733', '2019-07-22 02:07:26.117');
INSERT INTO test2.dbo.roles (user_role_type, created, updated) VALUES ('Reviewer', '2019-07-22 02:07:56.697', '2019-07-22 02:07:59.827');

SET IDENTITY_INSERT test2.dbo.users ON

INSERT INTO test2.dbo.users (id, user_role_id, username, password, email, password_token, token_expires, deletion_reason, is_active, is_deleted, created, updated)
    VALUES (13, 1, 'dummy', '$2a$10$tpz0joUBZ1qj51Pg7AOW4.2gQiwAbzEJ1BNik6.EJDJP/Kcn1TeqK', 'dummy@gmail.com', null, null, null, 1, 0, null, '2019-07-24 10:56:11.627');
INSERT INTO test2.dbo.users (id, user_role_id, username, password, email, password_token, token_expires, deletion_reason, is_active, is_deleted, created, updated)
    VALUES (14, 2, 'jason', '$2a$10$R.0FXTu8lu2NUJdMCzj9Iu941Or99UItwevg06fnL7oxnpyd40s/W', 'jason@gmail.com', null, null, null, 1, 0, null, '2019-08-08 17:38:24.517');

SET IDENTITY_INSERT test2.dbo.users OFF

INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Preliminaries', 1, 0, null, '2019-08-12 10:49:54.090');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Margin', 1, 0, null, '2019-08-12 10:49:54.093');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Groundworks', 1, 0, null, '2019-08-12 10:49:54.093');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Piling / Site Retention', 1, 0, null, '2019-08-12 10:49:54.093');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Concrete', 1, 0, null, '2019-08-12 10:49:54.097');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Formwork', 1, 0, null, '2019-08-12 10:49:54.097');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Reinforcement', 1, 0, null, '2019-08-12 10:49:54.097');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Post Tensioning', 1, 0, null, '2019-08-12 10:49:54.097');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Precast', 1, 0, null, '2019-08-12 10:49:54.097');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Masonry', 1, 0, null, '2019-08-12 10:49:54.097');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Roofing', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Façade', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Metal Work', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Carpentry', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Plastering & Insulation', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Floor and Wall Tiling', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Doors and Door hardware', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Joinery', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Timber Flooring', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Carpet', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Painting', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Final Cleaning', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Special Equipment', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Mechanical Services', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Electrical Services', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Fire Services', 1, 0, null, '2019-08-12 10:49:54.100');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Hydraulic Services', 1, 0, null, '2019-08-12 10:49:54.103');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Appliances', 1, 0, null, '2019-08-12 10:49:54.103');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Vertical Transportation', 1, 0, null, '2019-08-12 10:49:54.103');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('External Works', 1, 0, null, '2019-08-12 10:49:54.103');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('External Services', 1, 0, null, '2019-08-12 10:49:54.103');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Provisional Sums', 1, 0, null, '2019-08-12 10:49:54.103');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Approved Variations', 1, 0, null, '2019-08-12 10:49:54.103');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Pending Variations', 1, 0, null, '2019-08-12 10:49:54.103');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Forecast Variations', 1, 0, null, '2019-08-12 10:49:54.103');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Approved Provisional Sum Adjustment', 1, 0, null, '2019-08-12 10:49:54.103');
INSERT INTO test2.dbo.trade_categories (name, is_active, is_deleted, created, updated) VALUES ('Pending Provisoinal Sum Adjustment', 1, 0, null, '2019-08-12 10:49:54.103');


INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'GreenTownInf', 0, 0, '', '', 0, '', null, 1, 1, null, '2019-07-24 17:18:21.807');
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'GreenTown3', null, null, '', '', 300, '', null, 1, 1, null, '2019-07-24 11:14:17.290');
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'Green Doom 2', null, null, 'SYD-GD-12', '1 Orbe St, Auburn, NSW', null, null, null, 1, 1, null, '2019-08-09 15:18:53.460');
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'Green Doom 21', null, null, 'SYD-GD-12', '1 Orbe St, Auburn, NSW', 1000000, 'Michael Jason', null, 1, 0, null, '2019-08-09 15:23:41.540');
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'Green Doom 21', null, null, 'SYD-GD-12', '1 Orbe St, Auburn, NSW', 1000000, 'Michael Jason', null, 1, 0, '2019-08-09 15:36:11.040', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'Green Doom 21', null, null, 'SYD-GD-12', '1 Orbe St, Auburn, NSW', 1000000, 'Michael Jason', null, 1, 1, '2019-08-09 15:36:44.413', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'GreenTownInf', 0, 0, '', '', 0, '', '', 1, 1, '2019-08-09 15:39:32.983', '2019-08-09 15:42:38.230');
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'Green Doom 21', null, null, 'SYD-GD-12', '1 Orbe St, Auburn, NSW', 1000000, 'Michael Jason', 'notes asdfasfdasdfasfasfdsa
asdfasdfas', 1, 1, '2019-08-09 15:43:20.363', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'Green Doom 21', null, null, 'SYD-GD-12', '1 Orbe St, Auburn, NSW', 1000000, 'Michael Jason', 'notes asdfasfdasdfasfasfdsa
asdfasdfas', 1, 1, '2019-08-09 16:33:39.703', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'Green Doom 21', null, null, 'SYD-GD-12', '1 Orbe St, Auburn, NSW', 1000000, 'Michael Jason', 'notes asdfasfdasdfasfasfdsa
asdfasdfas', 1, 1, '2019-08-09 16:34:50.800', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'sadfsadf', null, null, '', '', 0, '', '', 1, 1, '2019-08-09 18:13:41.157', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'Tower One', null, null, '', '', 0, '', '', 1, 1, '2019-08-14 10:45:04.920', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'asdfsadf', null, null, 'AAAA', '1 get st', 10000, 'michael', '', 1, 0, '2019-08-14 11:04:52.890', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'ss', null, null, 'zhuang', '1 get st', 10000, 'michael', 'sadfsadf', 1, 0, '2019-08-14 12:25:04.600', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'Capochino', null, null, 'C12312313', '1 Sunset Road', 2000000, 'Michael', 'foo bar foo bar', 1, 0, '2019-08-14 12:34:29.343', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'asdf', null, null, 'asdfasdf', 'sadfasdf', 111111, 'asdfasdf', '', 1, 0, '2019-08-14 12:35:51.733', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'asdfsadf', null, null, 'sadfsadf', '1 get st', 1111, 'sadf', '', 1, 0, '2019-08-14 12:39:30.310', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'asdfasdf', null, null, 'asdfasdf', 'asdfsadf', 10000, 'asdfasdf', '', 1, 0, '2019-08-14 12:40:52.610', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, 'asdf', null, null, 'asdfasf', 'sadfsadf', 1111, 'sdfasdf', '', 1, 0, '2019-08-14 12:41:33.690', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, '22', null, null, '22', '22', 222, '22', '', 1, 0, '2019-08-14 12:43:57.383', null);
INSERT INTO test2.dbo.projects (manager_id, creator_id, name, total_trade_value, contractor_total_claim, serial_no, address, total_contract_value, quantity_surveyor, notes, is_active, is_deleted, created, updated) VALUES (13, 13, '33', null, null, '33', '33', 333, '33', '', 1, 0, '2019-08-14 12:45:38.530', null);


INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 3, 'floor 3 lower', 'sub a.b', 1234, null, 1, 1, '2019-08-16 17:49:26.830', '2019-08-16 17:49:26.833');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (14, 13, 15, '21', 'f1', 13.1, null, 1, 0, '2019-08-16 18:41:08.993', '2019-08-16 18:41:08.997');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (14, 13, 15, '22', 'f1', 14.1, null, 1, 0, '2019-08-16 18:41:36.617', '2019-08-16 18:41:36.620');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (15, 13, 15, '21', '133 ff', 33, null, 1, 0, '2019-08-16 18:46:04.757', '2019-08-16 18:46:04.763');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (15, 13, 15, '22', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:12.837', '2019-08-16 18:46:12.843');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 15, '22', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:20.117', '2019-08-16 18:46:20.117');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 15, '213', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:28.047', '2019-08-16 18:46:28.050');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 15, '213', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:29.407', '2019-08-16 18:46:29.410');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 15, '213', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:29.917', '2019-08-16 18:46:29.917');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 15, '213', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:30.423', '2019-08-16 18:46:30.423');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 15, '213', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:30.830', '2019-08-16 18:46:30.830');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 15, '213', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:31.220', '2019-08-16 18:46:31.220');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 15, '213', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:33.263', '2019-08-16 18:46:33.263');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 15, '213', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:34.003', '2019-08-16 18:46:34.000');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 15, '213', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:34.443', '2019-08-16 18:46:34.440');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 15, '213', '133 ffs', 33, null, 1, 0, '2019-08-16 18:46:35.080', '2019-08-16 18:46:35.077');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 3, 'floor 3 lower', 'sub g', 1234, null, 1, 0, '2019-08-16 19:07:52.973', '2019-08-16 19:07:52.980');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 3, 'floor 3 lower', 'sub z', 1234, null, 1, 0, '2019-08-16 19:07:57.013', '2019-08-16 19:07:57.010');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 3, 'floor 3 lower', 'sub s', 1234, null, 1, 0, '2019-08-16 19:08:00.030', '2019-08-16 19:08:00.030');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 3, 'floor 3 lower', 'sub s1', 1234, null, 1, 0, '2019-08-16 19:08:04.617', '2019-08-16 19:08:04.623');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 3, 'floor 3 lower', 'sub s21', 1234, null, 1, 0, '2019-08-16 19:08:07.753', '2019-08-16 19:08:07.757');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 3, 'floor 3 lower', 'sub s231', 1234, null, 1, 0, '2019-08-16 19:08:09.950', '2019-08-16 19:08:09.950');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (15, 13, 15, '21', 'f1fsadfasdfsafd', 332, null, 1, 0, '2019-08-16 19:37:31.693', '2019-08-16 19:37:31.700');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (15, 13, 15, '33', 'f1fsadfasdfsafd', 332, null, 1, 0, '2019-08-16 19:37:36.000', '2019-08-16 19:37:36.000');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (15, 13, 24, '12', 'subtitle 1', 1000, null, 1, 0, '2019-08-19 09:44:40.163', '2019-08-19 09:44:40.190');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (15, 13, 24, '12', 'subtitle 1', 1000, null, 1, 0, '2019-08-19 09:44:42.993', '2019-08-19 09:44:42.993');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (15, 13, 24, '12', 'subtitle 1', 1000, null, 1, 0, '2019-08-19 09:44:43.983', '2019-08-19 09:44:43.980');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 24, '12', 'subtitle 1', 1000, null, 1, 0, '2019-08-19 09:44:47.450', '2019-08-19 09:44:47.447');
INSERT INTO test2.dbo.trades (trade_category_id, creator_id, project_id, level, subtitle, value, temp, editable, is_deleted, created, updated) VALUES (13, 13, 24, '12', 'subtitle 1.2', 1000, null, 1, 0, '2019-08-19 09:45:08.510', '2019-08-19 09:45:08.513');

