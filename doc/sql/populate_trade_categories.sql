INSERT INTO test2.dbo.trade_categories (name) VALUES ('Preliminaries');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Margin');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Groundworks');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Piling / Site Retention');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Concrete');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Formwork');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Reinforcement');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Post Tensioning');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Precast');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Masonry');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Roofing');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Façade');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Metal Work');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Carpentry');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Plastering & Insulation');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Floor and Wall Tiling');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Doors and Door hardware');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Joinery');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Timber Flooring');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Carpet');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Painting');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Final Cleaning');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Special Equipment');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Mechanical Services');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Electrical Services');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Fire Services');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Hydraulic Services');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Appliances');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Vertical Transportation');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('External Works');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('External Services');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Provisional Sums');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Approved Variations');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Pending Variations');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Forecast Variations');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Approved Provisional Sum Adjustment');
INSERT INTO test2.dbo.trade_categories (name) VALUES ('Pending Provisoinal Sum Adjustment');

alter table trade_categories
    add constraint DF__basic_tra__creat__4AB81AF0 default getdate() for created
go

alter table trade_categories
    drop constraint DF__basic_tra__updat__4AB81AF0;
go
