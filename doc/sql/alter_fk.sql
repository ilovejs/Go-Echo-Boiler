use contractworksummary;

exec sp_rename 'dbo.AssignContarctor.ProjectId', project_id, 'COLUMN'
go

exec sp_rename 'dbo.AssignContarctor.UserId', user_id, 'COLUMN'
exec sp_rename 'dbo.AssignContarctor.TradeId', trade_id, 'COLUMN'
go

exec sp_rename 'dbo.AssignContarctor.UserId', user_id, 'COLUMN'
go

truncate table dbo.ErrorLog
go