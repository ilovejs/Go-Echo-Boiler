
CREATE PROC [dbo].[ApproveTradeItemClaim](@AutoId bigint)
AS
BEGIN
UPDATE TradeItemClaim
SET ActionClaim = 1
WHERE AutoIncrement =  @AutoId
    END



    go

CREATE PROC [dbo].[ContractorActiveList]
/*
exec ContractorList
*/
AS
BEGIN
SELECT L.ID AS UserLoginID,P.Id AS UserProfileID,L.UserName, P.FirstName,P.LastName,P.Email,P.MobileNo , L.IsActive,L.IsDeleted,P.Company
FROM UserLogin AS L
         JOIN UserProfile AS P ON P.UserLoginId = L.Id
WHERE L.CreatedBy IS NOT NULL AND L.IsDeleted = 0 AND  L.IsActive = 1
ORDER BY L.Id DESC
    END
    go


CREATE PROC [dbo].[ContractorList]
/*
exec ContractorList
*/
AS
BEGIN
SELECT L.ID AS UserLoginID,P.Id AS UserProfileID,L.UserName, P.FirstName,P.LastName,P.Email,P.MobileNo,L.IsActive,L.IsDeleted,P.Company,
    [dbo].[GetProjectCount](P.Id) as TotalProjectCount
FROM UserLogin AS L
    JOIN UserProfile AS P ON P.UserLoginId = L.Id
WHERE L.CreatedBy IS NOT NULL AND L.IsDeleted = 0
ORDER BY L.Id DESC
    END
    go



CREATE  PROC [dbo].[DeleteTempClaim]
AS
BEGIN
DELETE FROM TradeItemClaimHistory
WHERE TradeItemClaimId IN (SELECT Id FROM TradeItemClaim
                           WHERE ActionClaim IS NULL)

DELETE FROM TradeItemClaim
WHERE ActionClaim IS NULL
    END



    go

CREATE  PROC [dbo].[DeleteTempTradeItems](@ProjectId bigint,@TradeId int)
AS
/*
Admin
exec DeleteTempTradeItems 10069,3
*/
BEGIN
DELETE FROM TradeItem
WHERE  ProjectId = @ProjectId AND TradeId = @TradeId  AND TempCheck = 1
    END


    go


CREATE  PROC [dbo].[DeleteTradeItemByID](@Id bigint)
AS
BEGIN
DELETE FROM TradeItem
WHERE  Id = @Id
    END



    go


CREATE PROC [dbo].[GenerateProjectClaimReport](@ProjectId bigint)
--exec GenerateProjectClaimReport 20269
AS
BEGIN
SELECT TC.[AutoIncrement],TC.[ClaimNumber],P.Name,TI.Id,TI.TradeId,TI.[Level],TI.DescriptionOfWork,TI.ProjectId,TI.ItemBreakdown,
       TM.TradeName,ISNULL([dbo].[GetPreviousClaimById](TI.ID),0)PreviousClaim,

       ISNULL([dbo].[GetReportClaimAmount](TI.ID),0)ClaimAmount,
       CAST( CAST(year(getdate())  AS varchar) + '-' + CAST(TC.ClaimPeriod AS varchar) + '-' + CAST(1 AS varchar) as datetime) AS StartDate,
       DATEADD(s,-1,DATEADD(mm, DATEDIFF(m,0,CAST( CAST(year(getdate())  AS varchar) + '-' + CAST(TC.ClaimPeriod AS varchar) + '-' + CAST(1 AS varchar) as datetime))+1,0)) AS EndDate
FROM TradeItem AS TI
         JOIN TradeItemClaim AS TC ON TC.TradeItemId = TI.Id
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId and TC.[AutoIncrement] = (select MAX(AutoIncrement) FROM TradeItemClaim where TradeItemId =TI.Id)

    END



    go


CREATE PROC [dbo].[GenerateReport](@ProjectId bigint,@Month varchar(50),@ClaimNumber varchar(200))
--exec GenerateReport 20269
AS
BEGIN
SELECT TI.TradeId,TM.TradeName,SUM(TI.ItemBreakdown) AS ItemBreakdownAmount,P.Name,
       ISNULL([dbo].[GetClaimedAssessmentReport](TI.TradeId,@ProjectId,@Month,@ClaimNumber),0)ClaimedAmount,

       ISNULL([dbo].[GetReportPreviousClaim](TI.TradeId,@ProjectId,(TC.AutoIncrement-1)),0)PreviousClaim,

       @ClaimNumber AS ClaimNumber,
       CAST( CAST(year(getdate())  AS varchar) + '-' + CAST(@Month AS varchar) + '-' + CAST(1 AS varchar) as datetime) AS StartDate,
       DATEADD(s,-1,DATEADD(mm, DATEDIFF(m,0,CAST( CAST(year(getdate())  AS varchar) + '-' + CAST(@Month AS varchar) + '-' + CAST(1 AS varchar) as datetime))+1,0)) AS EndDate

FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN TradeItemClaim AS TC ON TC.TradeItemId = TI.Id
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId AND TC.ClaimPeriod = @Month AND TC.ClaimNumber =  @ClaimNumber  AND TI.TempCheck = 0 AND TI.TradeId NOT IN (33,34,35,36,37)
GROUP BY TI.TradeId,TM.TradeName,TI.ProjectId,P.Name,TC.ClaimNumber,TC.ClaimPeriod,TC.AutoIncrement
ORDER BY TI.TradeId
    END



    go



CREATE PROC [dbo].[GenerateVariationsReport](@ProjectId bigint,@Month varchar(50),@ClaimNumber varchar(200))
--exec GenerateVariationsReport 20272,'January','1'

AS

BEGIN
SELECT TI.TradeId AS TradeId,TM.TradeName AS TradeName,TI.[Level],TI.DescriptionOfWork,TI.ItemBreakdown AS ItemBreakdown,P.Name,
       ISNULL([dbo].[GetClaimedAssessmentReport](TI.TradeId,@ProjectId,@Month,@ClaimNumber),0)ClaimAmount,
       ISNULL([dbo].[GetReportPreviousClaim](TI.TradeId,@ProjectId,(TC.AutoIncrement-1)),0)PreviousClaim,
       @ClaimNumber AS ClaimNumber,
       CAST( CAST(year(getdate())  AS varchar) + '-' + CAST(@Month AS varchar) + '-' + CAST(1 AS varchar) as datetime) AS StartDate,
       DATEADD(s,-1,DATEADD(mm, DATEDIFF(m,0,CAST( CAST(year(getdate())  AS varchar) + '-' + CAST(@Month AS varchar) + '-' + CAST(1 AS varchar) as datetime))+1,0)) AS EndDate
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN TradeItemClaim AS TC ON TC.TradeItemId = TI.Id
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId  AND TC.ClaimPeriod = @Month AND TC.ClaimNumber =  @ClaimNumber AND TI.TempCheck = 0 AND TI.TradeId IN (33,34,35,36,37)
GROUP BY TI.TradeId,TM.TradeName,TI.ProjectId,P.Name,TC.ClaimNumber,TC.ClaimPeriod,TI.[Level],TI.DescriptionOfWork,TI.ItemBreakdown,TC.AutoIncrement
ORDER BY TI.TradeId
    END



    go


CREATE function [dbo].[GetActionClaim](@TradeItemId bigint)
    returns bit
as
begin
    Declare @Status bit
select top 1   @Status=ActionClaim from  TradeItemClaim
where TradeItemId =@TradeItemId
    return @Status
    end
    go

CREATE function [dbo].[GetActiveTradeItemCount](@ProjectId bigint)
    returns int
as
begin
    Declare @ItemCnt int
select @ItemCnt=count(Q.Id) from TradeItem as  Q
                                     join TradeMaster as tm on tm.Id = Q.TradeId
where Q.ProjectId=@ProjectId and isnull(Q.isdeleted,0)=0 and tm.isdeleted = 0 and Q.TempCheck = 0
    return @ItemCnt
    end
    go

CREATE   PROC [dbo].[GetAdminEmail](@ProjectId bigint)
AS
--exec GetAdminEmail 10165
BEGIN
SELECT Email,isnull([dbo].[GetContractorName](@ProjectId),'')Name
FROM UserProfile AS P
         JOIN UserLogin AS L ON P.UserLoginId = L.id
         JOIN UserRole AS R ON R.Id = L.UserRoleId
WHERE R.UserRoleType = 'On-site Team'
    END



    go

CREATE   PROC [dbo].[GetAllProjectList]
AS
/*
Admin
exec GetAllProjectList
*/
BEGIN
SELECT P.Id,P.UserId,P.CreatedBy,P.Name,P.TotalItemBreakdown,P.ContractorTotalClaim,P.ProjectNumber,p.ProjectDetails,P.QuantitySurveyor,p.TotalContractAmount,P.SiteNote,
    [dbo].[GetActiveTradeItemCount](P.Id) as ItemCount,
    (UP.FirstName +' '+UP.LastName +'('+ UP.Email + ')') AS Contractor
FROM Project AS P
    JOIN UserProfile AS UP ON UP.UserLoginId = P.UserId
WHERE P.IsDeleted=0 AND [dbo].[GetActiveTradeItemCount](P.Id) > 0
ORDER BY  P.Id DESC
    END

    go

CREATE function [dbo].[GetAmountDue](@TradeItemId int)
    returns decimal(18,2)
as
begin
    Declare @AmountDue Decimal(18,2)
select top 1  @AmountDue= SUM(TC.AmountDue) from TradeItem as TT
                                                     join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.TradeId =@TradeItemId
    return @AmountDue
    end
    go


CREATE PROC [dbo].[GetApprovedClaimNumberList](@ProjectId bigint)
--exec GetApprovedClaimNumberList 20271
AS
BEGIN
SELECT TC.AutoIncrement AS Id,TC.ClaimPeriod,TC.[ClaimNumber]
FROM TradeItemClaim AS TC
         JOIN TradeItem AS TI ON TC.TradeItemId = TI.Id
WHERE  TI.ProjectId = @ProjectId AND TC.ActionClaim = 1
group by TC.AutoIncrement,TC.ClaimPeriod,TC.[ClaimNumber]
order by TC.AutoIncrement desc
    END
    go


create function [dbo].[GetApprovedClaimProjectCount](@ProjectId bigint)
    returns int
as
begin
    Declare @ClaimCnt int
select @ClaimCnt= COUNT(TC.Id) from TradeItem as TT
                                        join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.ProjectId=@ProjectId and TC.ActionClaim = 1
    return @ClaimCnt
    end


    go


CREATE   PROC [dbo].[GetApprovedClaimProjectList]
AS
/*
Admin
exec GetAllProjectList
*/
BEGIN
SELECT P.Id,P.Name FROM Project AS P
WHERE P.IsDeleted=0 AND [dbo].[GetActiveTradeItemCount](P.Id) > 0 AND [dbo].[GetApprovedClaimProjectCount](P.Id) > 0
ORDER BY  P.Id DESC
    END

    go


CREATE PROC [dbo].[GetAssementByLevel](@ProjectId bigint,@AutoId bigint,@level VARCHAR(200))
--exec GetAssementByLevel 20272,1,'00'
AS
BEGIN
SELECT TI.Id,TI.[Level],TI.DescriptionOfWork,TI.ProjectId,TI.ItemBreakdown,
       ISNULL([dbo].[GetPreviousApprovedClaimById](TI.ID,@AutoId-1),0)PreviousClaim,
       TC.ActionClaim AS ActionClaim,TC.Id AS TradeItemClaimId,
       isnull([dbo].[GetClaimHistoryID](TI.ID),0)ClaimHistoryId,
       TC.ClaimPercentage,
       TC.ClaimedAmount AS ApproveClaimed , TM.Id AS TradeID , TM.TradeName AS TradeName
FROM TradeItem AS TI
         JOIN TradeItemClaim AS TC ON TI.ID = TC.TradeItemId
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 AND TI.Level = @level  AND TC.AutoIncrement = @AutoId
    END



    go

CREATE PROC [dbo].[GetAssementList](@ProjectId bigint,@TradeId int)
AS
BEGIN
SELECT TI.Id,TI.TradeId,TI.[Level],TI.DescriptionOfWork,TI.ProjectId,TI.ItemBreakdown,TM.TradeName,
       ISNULL([dbo].[GetAssessmentTotalClaim](TI.ID),0)TotalClaim,
       ISNULL([dbo].[GetAssessmentPreviousClaim](TI.ID),0)PreviousClaim,
       ISNULL([dbo].[GetAssessmentClaimID](TI.ID),0)ClaimHistoryID,
       ISNULL([dbo].[GetActionClaim](TI.ID),null)ActionClaim
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 and Ti.TradeId= @TradeId
    END



    go


CREATE PROC [dbo].[GetAssementOfTrade](@ProjectId bigint,@TradeId int,@AutoId bigint)
--exec GetAssementOfTrade 20272,1,4
AS
BEGIN
SELECT TI.Id,TI.[Level],TI.DescriptionOfWork,TI.ProjectId,TI.ItemBreakdown,
       ISNULL([dbo].[GetPreviousApprovedClaimById](TI.ID,@AutoId-1),0)PreviousClaim,
       TC.ActionClaim AS ActionClaim,TC.Id AS TradeItemClaimId,
       isnull([dbo].[GetClaimHistoryID](TI.ID),0)ClaimHistoryId,
       TC.ClaimPercentage,
       TC.ClaimedAmount AS ApproveClaimed , TM.Id AS TradeID , TM.TradeName AS TradeName
FROM TradeItem AS TI
         JOIN TradeItemClaim AS TC ON TI.ID = TC.TradeItemId
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 AND TI.TradeId = @TradeId  AND TC.AutoIncrement = @AutoId
    END



    go

CREATE function [dbo].[GetAssessmentClaimID](@TradeItemId bigint)
    returns bigint
as
begin
    Declare @Id bigint
select top 1  @Id=Id from  TradeItemClaimHistory
where TradeItemId =@TradeItemId
order by CreatedOn desc
    return @Id
    end
    go


CREATE function [dbo].[GetAssessmentPreviousClaim](@TradeItemId bigint)
    returns decimal(18,2)
as
begin
    Declare @PClaimAmount Decimal(18,2)
select top 2  @PClaimAmount= PreviousClaimed from  TradeItemClaimHistory
where TradeItemId =@TradeItemId
order by CreatedOn desc
    return @PClaimAmount
    end
    go

CREATE function [dbo].[GetAssessmentTotalClaim](@TradeItemId bigint)
    returns decimal(18,2)
as
begin
    Declare @PClaimAmount Decimal(18,2)
select top 1  @PClaimAmount= PreviousClaimed from  TradeItemClaimHistory
where TradeItemId =@TradeItemId
order by CreatedOn desc
    return @PClaimAmount
    end
    go

CREATE   PROC [dbo].[GetAssignContractor](@ProjectId bigint)
AS
/*
exec GetAssignContractor 11
*/
BEGIN
SELECT AC.Id AS AssignContarctorID,AC.UserId,AC.ProjectId,AC.TradeId,AC.SiteNotes,
    [dbo].[GetTradeItemCount](AC.ProjectId) AS ItemCount,TM.TradeName,
    (UP.FirstName +' '+UP.LastName +'('+ UP.Email + ')') AS Contractor
FROM AssignContarctor AS AC
    JOIN UserProfile AS UP ON UP.UserLoginId = AC.UserId
    JOIN TradeMaster AS TM ON TM.Id = AC.TradeId
WHERE  AC.ProjectId = @ProjectId
    END

    go

CREATE   PROC [dbo].[GetAssigndProject](@UserProfileId bigint)
AS
/*
exec GetAssigndProject 10014
*/
BEGIN
SELECT P.Id,P.UserId,P.CreatedBy,P.Name,P.TotalItemBreakdown,P.ContractorTotalClaim,P.ProjectNumber,
       P.ProjectDetails,P.QuantitySurveyor,p.TotalContractAmount,P.SiteNote,
       (UP.FirstName +' '+UP.LastName +'('+ UP.Email + ')') AS Contractor
FROM Project AS P
         JOIN UserProfile AS UP ON UP.UserLoginId = P.UserId
WHERE P.UserId= @UserProfileId AND P.IsDeleted = 0 AND [dbo].[GetActiveTradeItemCount](P.Id) > 0
    END

    go


CREATE function [dbo].[GetClaimAmount](@TradeItemId bigint)
    returns decimal(18,2)
as
begin
    Declare @PClaimAmount Decimal(18,2)
select  @PClaimAmount= ClaimedAmount  from  TradeItemClaim
where TradeItemId =@TradeItemId and ActionClaim = 0 and AutoIncrement = (select MAX(AutoIncrement) from TradeItemClaim where TradeItemId =@TradeItemId)
    return @PClaimAmount
    end
    go


CREATE PROC [dbo].[GetClaimAssement](@ProjectId bigint,@AutoId bigint)
--exec GetClaimAssement 20220,1
AS
BEGIN
SELECT TI.TradeId,SUM(TI.ItemBreakdown) AS TotalAmount,
       TM.TradeName,ISNULL([dbo].[GetPreviousApprovedClaim](TI.TradeId,@ProjectId,(@AutoId-1)),0)PreviousClaim,
       ISNULL([dbo].[GetClaimedAssessment](TI.TradeId,@AutoId,@ProjectId),0)ClaimedAmount
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 GROUP BY TI.TradeId,TM.TradeName,TI.ProjectId
    END



    go


CREATE function [dbo].[GetClaimDate](@TradeId bigint,@ProjectId bigint)
    returns datetime
as
begin
    Declare @date datetime
select top 1 @date= TC.UpdatedOn from TradeItem as TT
                                          join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.TradeId =@TradeId and TT.ProjectId = @ProjectId
order by TC.AutoIncrement




    return DATEADD(mm, DATEDIFF(mm, 0, @date), 0)
    end
    go


CREATE function [dbo].[GetClaimEndDate](@TradeId bigint,@ProjectId bigint)
    returns datetime
as
begin
    Declare @month varchar(50)
select top 1 @month= TC.ClaimPeriod from TradeItem as TT
                                             join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.TradeId =@TradeId and TT.ProjectId = @ProjectId and TC.ActionClaim = 1
order by TC.AutoIncrement desc
    return DATEADD(s,-1,DATEADD(mm, DATEDIFF(m,0,CAST( CAST(year(getdate())  AS varchar) + '-' + CAST(@month AS varchar) + '-' + CAST(1 AS varchar) as datetime))+1,0))
    end
    go


CREATE function [dbo].[GetClaimHistoryID](@TradeItemId bigint)
    returns bigint
as
begin
    Declare @Id bigint
select top 1  @Id=Id from  TradeItemClaimHistory
where TradeItemId =@TradeItemId
order by CreatedOn desc
    return @Id
    end
    go


CREATE function [dbo].[GetClaimProjectCount](@ProjectId bigint)
    returns int
as
begin
    Declare @ClaimCnt int
select @ClaimCnt= COUNT(TC.Id) from TradeItem as TT
                                        join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.ProjectId=@ProjectId and (TC.ActionClaim = 0 OR TC.ActionClaim = 1)
    return @ClaimCnt
    end
    go


CREATE   PROC [dbo].[GetClaimProjectList]
AS
/*
Admin
exec GetAllProjectList
*/
BEGIN
SELECT P.Id,P.Name FROM Project AS P
WHERE P.IsDeleted=0 AND [dbo].[GetActiveTradeItemCount](P.Id) > 0 AND [dbo].[GetClaimProjectCount](P.Id) > 0
ORDER BY  P.Id DESC
    END

    go


CREATE function [dbo].[GetClaimReportClaimNumber](@TradeItemId bigint,@ProjectId bigint)
    returns varchar(max)
as
begin
    Declare @ClaimNumber varchar(max)
select top 1 @ClaimNumber= ClaimNumber from TradeItemClaim
where TradeItemId =@TradeItemId and ActionClaim = 0
order by AutoIncrement desc
    return @ClaimNumber
    end
    go


CREATE function [dbo].[GetClaimStartDate](@TradeId bigint,@ProjectId bigint)
    returns datetime
as
begin
    Declare @month varchar(50)
select top 1 @month= TC.ClaimPeriod from TradeItem as TT
                                             join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.TradeId =@TradeId and TT.ProjectId = @ProjectId and TC.ActionClaim = 1
order by TC.AutoIncrement
    return CAST( CAST(year(getdate())  AS varchar) + '-' + CAST(@month AS varchar) + '-' + CAST(1 AS varchar) as datetime)
    end

    go

CREATE PROC [dbo].[GetClaimSummaryByMonth](@ProjectId bigint)
--exec GetClaimSummaryByMonth 101624
AS
BEGIN
SELECT SUM(TC.ClaimedAmount) AS Amount,TC.ActionClaim,
       TC.ClaimPeriod AS CreatedMonthName,TC.AutoIncrement
FROM TradeItemClaim AS TC
         JOIN TradeItem AS TI ON TI.id = TC.TradeItemId
WHERE  TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 AND (TC.ActionClaim = 0 OR TC.ActionClaim = 1)
GROUP BY TC.ClaimPeriod,TC.ActionClaim,TC.AutoIncrement
ORDER BY TC.AutoIncrement DESC
    END



    go

CREATE function [dbo].[GetClaimedAmount](@TradeItemId int)
    returns decimal(18,2)
as
begin
    Declare @CAmount Decimal(18,2)
select top 1  @CAmount= SUM(TC.ClaimedAmount) from TradeItem as TT join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.TradeId =@TradeItemId and TC.AutoIncrement = [dbo].[GetMaxAutoIncrement](TC.TradeItemId) and TC.ActionClaim is null
    return @CAmount
    end
    go


CREATE function [dbo].[GetClaimedAmountAssessment](@TradeId bigint,@Id bigint)
    returns decimal(18,2)
as
begin
    Declare @CAmount Decimal(18,2)
select  @CAmount= SUM(TC.ClaimedAmount) from TradeItem as TT join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.TradeId =@TradeId and TC.AutoIncrement = @Id
group by TT.TradeId
    return @CAmount
    end
    go


Create function [dbo].[GetClaimedAmountById](@TradeItemId bigint)
    returns decimal(18,2)
as
begin
    Declare @PClaimAmount Decimal(18,2)
select  @PClaimAmount= ClaimedAmount  from  TradeItemClaim
where TradeItemId =@TradeItemId
    return @PClaimAmount
    end
    go


create function [dbo].[GetClaimedAssessment](@TradeId bigint,@Id bigint,@ProjectId bigint)
    returns decimal(18,2)
as
begin
    Declare @CAmount Decimal(18,2)
select @CAmount= SUM(TC.ClaimedAmount) from TradeItem as TT
                                                join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.TradeId =@TradeId and TC.AutoIncrement = @Id and TT.ProjectId = @ProjectId
group by TT.TradeId
    return @CAmount
    end
    go


CREATE function [dbo].[GetClaimedAssessmentReport](@TradeId bigint,@ProjectId bigint,@Month varchar(50),@ClaimNumber varchar(200))
    returns decimal(18,2)
as
begin
    Declare @CAmount Decimal(18,2)
select @CAmount= SUM(TC.ClaimedAmount) from TradeItem as TT
                                                join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.TradeId =@TradeId  and TT.ProjectId = @ProjectId and TC.ActionClaim = 1 and TC.ClaimPeriod = @Month and TC.ClaimNumber = @ClaimNumber
group by TT.TradeId
    return @CAmount
    end
    go

CREATE   PROC [dbo].[GetClaimhistory](@TradeclaimId bigint)
AS
/*
exec GetClaimhistory 10168
*/
BEGIN
SELECT CM.Id,CM.TradeItemId,CM.TradeItemClaimId,CM.PreviousClaimed,CM.CreatedOn,CM.CreatedBy,UR.UserRoleType,
       Convert(varchar,CM.CreatedOn,103) As Date
FROM TradeItemClaimHistory AS CM
         JOIN UserProfile AS UP ON UP.Id = CM.CreatedBy
         JOIN UserLogin AS UL ON UL.Id = UP.UserLoginId
         JOIN UserRole AS UR ON UR.Id = UL.UserRoleId
    --JOIN TradeItemClaim AS TC ON TC.Id = CM.TradeItemClaimId
WHERE  CM.TradeItemId = @TradeclaimId --AND TC.ActionClaim = 1
ORDER BY CM.Id DESC
    END



    go

CREATE PROC [dbo].[GetContractorDetail](@LoginUserID BIGINT)
/*
exec GetContractorDetail 2
*/
AS
BEGIN
SELECT L.ID AS UserLoginID,P.Id AS UserProfileID,L.UserName, P.FirstName,P.LastName,P.Email,P.MobileNo,
       L.IsActive,L.IsDeleted,P.Company
FROM UserLogin AS L
         JOIN UserProfile AS P ON P.UserLoginId = L.Id
WHERE  L.ID = @LoginUserID
    END
    go

CREATE function [dbo].[GetContractorName](@ProjectId bigint)
    returns VARCHAR(MAX)
as
begin
    Declare @Name VARCHAR(MAX)
select  @Name= (U.FirstName+' '+U.LastName)  from UserProfile as U
                                                      join Project as P on P.UserId = U.Id
where P.Id = @ProjectId
    return @Name
    end
    go

CREATE   PROC [dbo].[GetContractorProjectList](@UserId bigint)
AS
BEGIN
SELECT P.Id,P.UserId,P.CreatedBy,P.Name,P.TotalItemBreakdown,P.ContractorTotalClaim,P.ProjectNumber,p.ProjectDetails,P.QuantitySurveyor,p.TotalContractAmount,P.SiteNote,
    [dbo].[GetActiveTradeItemCount](P.Id) AS TotalTradeItem,
    (UP.FirstName +' '+UP.LastName +'('+ UP.Email + ')') AS Contractor
FROM Project AS P
    LEFT JOIN (SELECT COUNT(Id) AS  TotalTradeItem,ProjectId FROM TradeItem WHERE IsDeleted=0 AND TempCheck=0 GROUP BY ProjectId  ) AS AC ON  P.Id = AC.ProjectId
    JOIN UserProfile AS UP ON UP.UserLoginId = P.UserId
WHERE isnull(P.IsDeleted,0)=0 and p.UserId=@UserId AND [dbo].[GetActiveTradeItemCount](P.Id) > 0
ORDER BY  P.Id DESC
    END

    go

CREATE function [dbo].[GetCostToCompleted](@TradeItemId int)
    returns decimal(18,2)
as
begin
    Declare @CTC Decimal(18,2)
select top 1  @CTC= SUM(TC.CostToCompleted) from TradeItem as TT join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.TradeId =@TradeItemId
    return @CTC
    end
    go

CREATE   PROC [dbo].[GetLastAutoIncrement](@ProjectId bigint)
AS
/*
*/
BEGIN
SELECT CASE WHEN TC.ActionClaim IS NULL THEN ISNULL(Max(TC.AutoIncrement),0) ELSE ISNULL(Max(TC.AutoIncrement),0)+1 END FROM TradeItemClaim AS TC
                                                                                                                                 JOIN TradeItem AS TI ON TI.Id = TC.TradeItemId
WHERE TI.ProjectId = @ProjectId
GROUP BY TC.ActionClaim
    END

    go


CREATE function [dbo].[GetLastClaimNumber](@TradeId bigint,@ProjectId bigint)
    returns varchar(max)
as
begin
    Declare @ClaimNumber varchar(max)
select top 1 @ClaimNumber= TC.ClaimNumber from TradeItem as TT
                                                   join TradeItemClaim as TC on TT.Id=TC.TradeItemId
where TT.TradeId =@TradeId and TT.ProjectId = @ProjectId and TC.ActionClaim = 1
order by TC.AutoIncrement desc
    return @ClaimNumber
    end
    go


create function [dbo].[GetMaxAutoIncrement](@TradeItemId int)
    returns bigint
as
begin
    Declare @Maxvalue bigint
select @Maxvalue =MAX(AutoIncrement) FROM TradeItemClaim
where TradeItemId =@TradeItemId
    return @Maxvalue
    end
    go


Create function [dbo].[GetPreviousApprovedClaim](@TradeId int,@ProjectId bigint,@AutoId bigint)
    returns decimal(18,2)
as
begin
    Declare @PClaimAmount Decimal(18,2)
select top 1  @PClaimAmount= SUM(TH.PreviousClaimed) from  TradeItemClaimHistory AS TH
                                                               join TradeItemClaim AS TC ON TC.Id = TH.TradeItemClaimId
                                                               join TradeItem AS TI ON TI.Id = TC.TradeItemId
where TI.ProjectId=@ProjectId AND TC.ActionClaim = 1 AND TI.TradeId =@TradeId AND TH.AutoIncrement = @AutoId
group by TH.AutoIncrement
order by TH.AutoIncrement desc
    return @PClaimAmount
    end

    go


create function [dbo].[GetPreviousApprovedClaimById](@TradeItemId bigint,@AutoId bigint)
    returns decimal(18,2)
as
begin
    Declare @PClaimAmount Decimal(18,2)
select top 1  @PClaimAmount= SUM(TH.PreviousClaimed) from  TradeItemClaimHistory AS TH
                                                               join TradeItemClaim AS TC ON TC.Id = TH.TradeItemClaimId
where TH.TradeItemId =@TradeItemId AND TC.ActionClaim = 1  AND TH.AutoIncrement = @AutoId
group by TH.AutoIncrement
order by TH.AutoIncrement desc
    return @PClaimAmount
    end
    go

CREATE function [dbo].[GetPreviousClaim](@TradeId int,@ProjectId bigint)
    returns decimal(18,2)
as
begin
    Declare @PClaimAmount Decimal(18,2)
select top 1  @PClaimAmount= SUM(TH.PreviousClaimed) from  TradeItemClaimHistory AS TH
                                                               join TradeItemClaim AS TC ON TC.Id = TH.TradeItemClaimId
                                                               join TradeItem AS TI ON TI.Id = TC.TradeItemId
where TI.ProjectId=@ProjectId AND TC.ActionClaim = 1 AND TI.TradeId =@TradeId
group by TH.AutoIncrement
order by TH.AutoIncrement desc
    return @PClaimAmount
    end

    go

CREATE function [dbo].[GetPreviousClaimById](@TradeItemId bigint)
    returns decimal(18,2)
as
begin
    Declare @PClaimAmount Decimal(18,2)
select top 1  @PClaimAmount= SUM(TH.PreviousClaimed) from  TradeItemClaimHistory AS TH
                                                               join TradeItemClaim AS TC ON TC.Id = TH.TradeItemClaimId
where TH.TradeItemId =@TradeItemId AND TC.ActionClaim = 1
group by TH.AutoIncrement
order by TH.AutoIncrement desc
    return @PClaimAmount
    end
    go


CREATE PROC [dbo].[GetProjectAmountOfTrade](@ProjectId bigint)
AS
/*
exec GetProjectAmountOfTrade 20272
*/
BEGIN
SELECT ISNULL(SUM(TI.ItemBreakdown),0) AS ItemBreakdown, P.TotalContractAmount	FROM TradeItem AS TI
                                                                                           JOIN Project AS P ON P.Id = @ProjectId
WHERE  TI.ProjectId = @ProjectId AND TI.IsDeleted = 0 AND TI.TempCheck = 0
GROUP BY P.TotalContractAmount

    END



    go

CREATE function [dbo].[GetProjectCount](@UserProfileId bigint)
    returns int
as
begin
    Declare @ProjectCnt int
select @ProjectCnt=count(Q.Id) from Project as  Q
where Q.UserId=@UserProfileId and isnull(Q.isdeleted,0)=0 and [dbo].[GetActiveTradeItemCount](Q.Id) > 0
    return @ProjectCnt
    end
    go

CREATE   PROC [dbo].[GetProjectDetailById](@ProjectId bigint)
AS
/*
exec GetAllProjectList
*/
BEGIN
SELECT P.Id,P.UserId,P.CreatedBy,P.Name,P.TotalItemBreakdown,P.ContractorTotalClaim,P.ProjectNumber,p.ProjectDetails,P.QuantitySurveyor,p.TotalContractAmount,P.SiteNote,
       (UP.FirstName +' '+UP.LastName +'('+ UP.Email + ')') AS Contractor
FROM Project AS P
         JOIN UserProfile AS UP ON UP.UserLoginId = P.UserId
WHERE P.Id= @ProjectId
    END

    go


CREATE function [dbo].[GetReportClaimAmount](@TradeItemId bigint)
    returns decimal(18,2)
as
begin
    Declare @PClaimAmount Decimal(18,2)
select top 1  @PClaimAmount= ClaimedAmount  from  TradeItemClaim
where TradeItemId =@TradeItemId and ActionClaim = 0
order by AutoIncrement desc
    return @PClaimAmount
    end

    go


CREATE function [dbo].[GetReportPreviousClaim](@TradeId int,@ProjectId bigint, @AutoIncrement bigint)
    returns decimal(18,2)
as
begin
    Declare @PClaimAmount Decimal(18,2)
select top 1  @PClaimAmount= SUM(TH.PreviousClaimed) from  TradeItemClaimHistory AS TH
                                                               join TradeItemClaim AS TC ON TC.Id = TH.TradeItemClaimId
                                                               join TradeItem AS TI ON TI.Id = TC.TradeItemId
where TI.ProjectId=@ProjectId AND TC.ActionClaim = 1 AND TI.TradeId =@TradeId and TH.AutoIncrement =@AutoIncrement
group by TH.AutoIncrement
order by TH.AutoIncrement desc

    return @PClaimAmount
    end
    go


CREATE function [dbo].[GetTempClaimAmount](@TradeItemId bigint)
    returns decimal(18,2)
as
begin
    Declare @PClaimAmount Decimal(18,2)
select  @PClaimAmount= ClaimedAmount  from  TradeItemClaim
where TradeItemId =@TradeItemId and ActionClaim is null --and AutoIncrement = MAX(AutoIncrement)
    return @PClaimAmount
    end
    go


create function [dbo].[GetTotalAmount](@ProjectId bigint)
    returns decimal(18,2)
as
begin
    Declare @TotalAmount Decimal(18,2)
select @TotalAmount= SUM(ItemBreakDown) from TradeItem
where ProjectId =@ProjectId AND TempCheck = 0 group by ProjectId
    return @TotalAmount
    end
    go

CREATE   PROC [dbo].[GetTradeItemByID](@Id bigint)
AS
/*
exec GetTradeList 10157
*/
BEGIN
SELECT TI.Id,TI.TradeId,TI.ProjectId,TI.[Level],TI.DescriptionOfWork,TI.ItemBreakdown,TM.TradeName
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
WHERE  TI.Id = @Id AND TI.IsDeleted = 0 AND TI.TempCheck = 0
    END



    go


CREATE PROC [dbo].[GetTradeItemByProjectId](@ProjectId bigint)
AS
/*
exec GetTradeList 10200
*/
BEGIN
SELECT TI.Id
FROM TradeItem AS TI
WHERE  TI.ProjectId = @ProjectId  AND TI.IsDeleted = 0 AND TI.TempCheck = 0
ORDER BY TI.Id
    END



    go


CREATE PROC [dbo].[GetTradeItemByTradeID](@ProjectId bigint,@TradeId int)

AS
BEGIN
SELECT TI.Id,TI.TradeId,TI.[Level],TI.DescriptionOfWork,TI.ProjectId,TI.ItemBreakdown,
       TM.TradeName,ISNULL([dbo].[GetPreviousClaimById](TI.ID),0)PreviousClaim,
       ISNULL([dbo].[GetClaimedAmountById](TI.ID),0)ClaimedAmount
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 and Ti.TradeId= @TradeId
    END



    go

CREATE function [dbo].[GetTradeItemCount](@ProjectId bigint)
    returns int
as
begin
    Declare @ItemCnt int
select @ItemCnt=count(Id) from TradeItem as  Q
where Q.ProjectId=@ProjectId and isnull(isdeleted,0)=0
    return @ItemCnt
    end
    go

CREATE PROC [dbo].[GetTradeItemLevelList](@ProjectId bigint)
AS
/*
exec GetTradeItemLevelList 20272
*/
BEGIN
SELECT DISTINCT TI.ProjectId AS Id, TI.[Level]  AS Level
FROM TradeItem AS TI
WHERE  TI.ProjectId = @ProjectId
ORDER BY TI.[Level]
    END



    go


CREATE PROC [dbo].[GetTradeItemList](@ProjectId bigint,@TradeId int)
--exec GetTradeItemList 20274,5
AS
BEGIN
SELECT TI.Id,TI.TradeId,TI.[Level],TI.DescriptionOfWork,TI.ProjectId,TI.ItemBreakdown,
       TM.TradeName,ISNULL([dbo].[GetPreviousClaimById](TI.ID),0)PreviousClaim,
       ISNULL([dbo].[GetClaimAmount](TI.ID),0)ClaimAmount,
       ISNULL([dbo].[GetClaimHistoryID](TI.ID),0)ClaimHistoryId
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 and Ti.TradeId= @TradeId
    END



    go


CREATE PROC [dbo].[GetTradeItemListByLevel](@ProjectId bigint,@level VARCHAR(200))
--exec GetTradeItemListByLevel
AS
BEGIN
SELECT TI.Id,TI.TradeId,TI.[Level],TI.DescriptionOfWork,TI.ProjectId,TI.ItemBreakdown,
       TM.TradeName,ISNULL([dbo].[GetPreviousClaimById](TI.ID),0)PreviousClaim,
       ISNULL([dbo].[GetClaimAmount](TI.ID),0)ClaimAmount,
       ISNULL([dbo].[GetClaimHistoryID](TI.ID),0)ClaimHistoryId
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 and TI.Level= @level
    END



    go



CREATE PROC [dbo].[GetTradeItemOfProjectByLevel](@ProjectId bigint,@level VARCHAR(200))
--exec GetTradeItemOfProjectByLevel 10162
AS
BEGIN
SELECT TI.Id,TI.TradeId,TI.[Level],TI.DescriptionOfWork,TI.ProjectId,TI.ItemBreakdown,
       TM.TradeName,ISNULL([dbo].[GetPreviousClaimById](TI.ID),0)PreviousClaim,
       ISNULL([dbo].[GetTempClaimAmount](TI.ID),0)ClaimAmount,
       ISNULL([dbo].[GetClaimHistoryID](TI.ID),0)ClaimHistoryId
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 and TI.Level = @level
    END



    go


CREATE PROC [dbo].[GetTradeItemOfProjectByTradeID](@ProjectId bigint,@TradeId int)
--exec GetTradeItemOfProjectByTradeID 10162
AS
BEGIN
SELECT TI.Id,TI.TradeId,TI.[Level],TI.DescriptionOfWork,TI.ProjectId,TI.ItemBreakdown,
       TM.TradeName,ISNULL([dbo].[GetPreviousClaimById](TI.ID),0)PreviousClaim,
       ISNULL([dbo].[GetTempClaimAmount](TI.ID),0)ClaimAmount,
       ISNULL([dbo].[GetClaimHistoryID](TI.ID),0)ClaimHistoryId
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 and Ti.TradeId= @TradeId
    END



    go

CREATE   PROC [dbo].[GetTradeList](@ProjectId bigint)
AS
/*
exec GetTradeList 10200
*/
BEGIN
SELECT TI.Id,TI.TradeId,TI.ProjectId,TI.[Level],TI.DescriptionOfWork,TI.ItemBreakdown,TM.TradeName
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
WHERE  TI.ProjectId = @ProjectId AND TI.IsDeleted = 0 AND TI.TempCheck = 0 AND TM.IsDeleted = 0
ORDER BY TI.Id
    END



    go

CREATE PROC [dbo].[GetTradeListById](@TradeId int,@ProjectId bigint)
AS
/*
exec GetTradeList 10200
*/
BEGIN
SELECT TI.Id,TI.TradeId,TI.ProjectId,TI.[Level],TI.DescriptionOfWork,TI.ItemBreakdown,TM.TradeName
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
WHERE  TI.ProjectId = @ProjectId AND TI.TradeId =@TradeId AND TI.IsDeleted = 0 AND TI.TempCheck = 0 AND TM.IsDeleted = 0
ORDER BY TI.Id
    END



    go

CREATE PROC [dbo].[GetTradeMasterList]
/*
exec GetTradeMasterList
*/
AS
BEGIN
SELECT Id , TradeName FROM TradeMaster
WHERE IsDeleted = 0
    END
    go


CREATE PROC [dbo].[GetTradeOfProjectReport](@ProjectId bigint)
--exec GetTradeOfProjectReport 20269
AS
BEGIN
SELECT distinct TI.TradeId,TM.TradeName
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
WHERE  TI.ProjectId = @ProjectId
    END



    go



CREATE PROC [dbo].[GetTradeOfVariationsReport](@ProjectId bigint)
--exec GetTradeOfVariationsReport 20272
AS
BEGIN
SELECT distinct TI.TradeId,TM.TradeName
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
WHERE  TI.ProjectId = @ProjectId AND TM.Id IN (33,34,35,36,37)
    END



    go

CREATE PROC [dbo].[GetTradeSummary](@ProjectId bigint)
AS
/*
exec GetTradeSummary 202
*/
BEGIN
SELECT TI.TradeId,TM.TradeName ,ISNULL(SUM(TI.ItemBreakdown),0) AS ItemBreakdown
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
WHERE  TI.ProjectId = @ProjectId AND TI.IsDeleted = 0 AND TI.TempCheck = 0 AND TM.IsDeleted = 0
GROUP BY TI.TradeId,TM.TradeName
ORDER BY TI.TradeId
    END



    go

CREATE   PROC [dbo].[GetTradeUseCount](@TradeId int)
AS
/*
select * from GetTradeUseCount
*/
BEGIN
SELECT COUNT(TI.Id) AS TradeCount
FROM TradeItem AS TI
WHERE TI.TradeId = @TradeId AND TI.IsDeleted = 0
    END



    go

CREATE  PROC [dbo].[ProjectFinalize](@ProjectId bigint)
AS
/*
 exec ProjectFinalize
*/
BEGIN
UPDATE TradeItem
SET TempCheck = 0
WHERE  ProjectId = @ProjectId

SELECT TradeId,ProjectId,SUM(ItemBreakDown) AS Total,
    [dbo].[GetTotalAmount](@ProjectId)TotalItemBreakDown
FROM TradeItem
WHERE ProjectId = @ProjectId AND TempCheck = 0 group by TradeId ,ProjectId
    END



    go

CREATE  PROC [dbo].[ProjectTradeDetailClaim](@ProjectId bigint)
AS
/*
exec ProjectTradeDetailClaim 10162
*/
BEGIN
SELECT TI.Id,TI.TradeId,TM.TradeName,TI.ProjectId,SUM(TI.ItemBreakDown) AS Total,
       TM.TradeName,isnull([dbo].[GetPreviousClaim](TI.Id,@ProjectId),0)PreviousClaim
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
WHERE TI.ProjectId = @ProjectId AND TI.TempCheck = 0 AND TM.IsDeleted = 0 group by TI.TradeId ,TI.ProjectId,TM.TradeName,TI.Id
    END



    go


CREATE PROC [dbo].[TradeGenerateReport](@ProjectId bigint,@TradeId int,@Month varchar(50),@ClaimNumber varchar(200))
--exec TradeGenerateReport 20262,33

AS
DECLARE @RCount int
BEGIN
SELECT @RCount=COUNT(TI.TradeId)
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 AND TI.TradeId = @TradeId--IN (33,34,35,36,37)
GROUP BY TI.TradeId,TM.TradeName,TI.ProjectId,P.Name
ORDER BY TI.TradeId

    IF @RCount > 0
BEGIN
SELECT TI.TradeId AS TradeId,TM.TradeName AS TradeName,ISNULL(SUM(TI.ItemBreakdown),0) AS ItemBreakdownAmount,
       ISNULL([dbo].[GetClaimedAssessmentReport](TI.TradeId,@ProjectId,@Month,@ClaimNumber),0)ClaimedAmount,
       ISNULL([dbo].[GetReportPreviousClaim](TI.TradeId,@ProjectId,(TC.AutoIncrement-1)),0)PreviousClaim
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN TradeItemClaim AS TC ON TC.TradeItemId = TI.Id
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId  AND TC.ClaimPeriod = @Month AND TC.ClaimNumber =  @ClaimNumber AND TI.TempCheck = 0 AND TI.TradeId = @TradeId --IN (33,34,35,36,37)
GROUP BY TI.TradeId,TM.TradeName,TI.ProjectId,P.Name,TC.ClaimNumber,TC.ClaimPeriod,TC.AutoIncrement
ORDER BY TI.TradeId
    END
    ELSE
BEGIn
SELECT TM.Id AS TradeId,TM.TradeName AS TradeName,0.00 AS ItemBreakdownAmount,
       0.00 AS ClaimedAmount,
       0.00 AS PreviousClaim
FROM TradeMaster AS TM
WHERE TM.Id =@TradeId --IN (33,34,35,36,37)
ORDER BY TM.Id
    END
    END



    go


CREATE PROC [dbo].[TradeItemsDetail](@ProjectId bigint)
--exec TradeItemsDetail 20215
AS
BEGIN
SELECT TradeId,SUM(TI.ItemBreakdown) AS TotalAmount,TM.TradeName,
       ISNULL([dbo].[GetPreviousClaim](TI.TradeId,@ProjectId),0)PreviousClaim,
       ISNULL([dbo].[GetCostToCompleted](TI.TradeId),0)CostToCompleted,
       ISNULL([dbo].[GetAmountDue](TI.TradeId),0)AmountDue,
       ISNULL([dbo].[GetClaimedAmount](TI.TradeId),0)ClaimedAmount
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
         JOIN Project AS P On P.id = @ProjectId
WHERE  TI.ProjectId = @ProjectId  AND TI.TempCheck = 0 GROUP BY TI.TradeId,TM.TradeName,TI.ProjectId
    END



    go

CREATE  PROC [dbo].[TradeItemsListParament](@ProjectId bigint,@TradeId int)
AS
/*
exec TradeItemsListTemp 11,44
*/
BEGIN
SELECT TI.Id,TI.TradeId,TI.ProjectId,TI.[Level],TI.DescriptionOfWork,TI.ItemBreakdown,TM.TradeName
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
WHERE  TI.ProjectId = @ProjectId AND TI.TradeId = @TradeId  AND TI.TempCheck = 0
    END



    go

CREATE  PROC [dbo].[TradeItemsListTemp](@ProjectId bigint,@TradeId int)
AS
/*
exec TradeItemsListTemp 11,44

*/
BEGIN
SELECT TI.Id,TI.TradeId,TI.ProjectId,TI.[Level],TI.DescriptionOfWork,TI.ItemBreakdown,TM.TradeName
FROM TradeItem AS TI
         JOIN TradeMaster AS TM ON TM.Id = TI.TradeId
WHERE  TI.ProjectId = @ProjectId AND TI.TradeId = @TradeId
    END



    go


CREATE PROC [dbo].[TradeNotAvailableGenerateReport](@ProjectId bigint)
--exec TradeGenerateReport 20262
AS
BEGIN
SELECT TM.Id,TM.TradeName,0 AS ItemBreakdownAmount,
       0 AS ClaimedAmount,
       0 AS PreviousClaim
FROM TradeMaster AS TM

WHERE   TM.Id IN (33,34,35,36,37)

ORDER BY TM.Id
    END



    go


CREATE PROC [dbo].[UpdateTradeItemClaim](@TradeItemId bigint,@Month varchar(50),@Number varchar(200))
AS
BEGIN
UPDATE TradeItemClaim
SET ActionClaim = 0,ClaimNumber = @Number,ClaimPeriod = @Month
WHERE TradeItemId = @TradeItemId AND AutoIncrement = [dbo].[GetMaxAutoIncrement](@TradeItemId)
    END



    go


CREATE FUNCTION dbo.fn_diagramobjects()
    RETURNS int
WITH EXECUTE AS N'dbo'
    AS
BEGIN
    declare @id_upgraddiagrams		int
		declare @id_sysdiagrams			int
		declare @id_helpdiagrams		int
		declare @id_helpdiagramdefinition	int
		declare @id_creatediagram	int
		declare @id_renamediagram	int
		declare @id_alterdiagram 	int
		declare @id_dropdiagram		int
		declare @InstalledObjects	int

select @InstalledObjects = 0

select 	@id_upgraddiagrams = object_id(N'dbo.sp_upgraddiagrams'),
          @id_sysdiagrams = object_id(N'dbo.sysdiagrams'),
          @id_helpdiagrams = object_id(N'dbo.sp_helpdiagrams'),
          @id_helpdiagramdefinition = object_id(N'dbo.sp_helpdiagramdefinition'),
          @id_creatediagram = object_id(N'dbo.sp_creatediagram'),
          @id_renamediagram = object_id(N'dbo.sp_renamediagram'),
          @id_alterdiagram = object_id(N'dbo.sp_alterdiagram'),
          @id_dropdiagram = object_id(N'dbo.sp_dropdiagram')

    if @id_upgraddiagrams is not null
select @InstalledObjects = @InstalledObjects + 1
    if @id_sysdiagrams is not null
select @InstalledObjects = @InstalledObjects + 2
    if @id_helpdiagrams is not null
select @InstalledObjects = @InstalledObjects + 4
    if @id_helpdiagramdefinition is not null
select @InstalledObjects = @InstalledObjects + 8
    if @id_creatediagram is not null
select @InstalledObjects = @InstalledObjects + 16
    if @id_renamediagram is not null
select @InstalledObjects = @InstalledObjects + 32
    if @id_alterdiagram  is not null
select @InstalledObjects = @InstalledObjects + 64
    if @id_dropdiagram is not null
select @InstalledObjects = @InstalledObjects + 128

    return @InstalledObjects
    END
    go


CREATE PROCEDURE dbo.sp_alterdiagram
    (
    @diagramname 	sysname,
		@owner_id	int	= null,
		@version 	int,
		@definition 	varbinary(max)
	)
	WITH EXECUTE AS 'dbo'
	AS
BEGIN
set nocount on

    declare @theId 			int
		declare @retval 		int
		declare @IsDbo 			int

		declare @UIDFound 		int
		declare @DiagId			int
		declare @ShouldChangeUID	int

		if(@diagramname is null)
begin
    RAISERROR ('Invalid ARG', 16, 1)
              return -1
		end

execute as caller;
select @theId = DATABASE_PRINCIPAL_ID();
select @IsDbo = IS_MEMBER(N'db_owner');
if(@owner_id is null)
select @owner_id = @theId;
revert;

select @ShouldChangeUID = 0
select @DiagId = diagram_id, @UIDFound = principal_id from dbo.sysdiagrams where principal_id = @owner_id and name = @diagramname

    if(@DiagId IS NULL or (@IsDbo = 0 and @theId <> @UIDFound))
begin
    RAISERROR ('Diagram does not exist or you do not have permission.', 16, 1);
return -3
		end

		if(@IsDbo <> 0)
begin
    if(@UIDFound is null or USER_NAME(@UIDFound) is null) -- invalid principal_id
begin
select @ShouldChangeUID = 1 ;
end
		end

		-- update dds data
update dbo.sysdiagrams set definition = @definition where diagram_id = @DiagId ;

-- change owner
if(@ShouldChangeUID = 1)
update dbo.sysdiagrams set principal_id = @theId where diagram_id = @DiagId ;

-- update dds version
if(@version is not null)
update dbo.sysdiagrams set version = @version where diagram_id = @DiagId ;

return 0
	END
go


CREATE PROCEDURE dbo.sp_creatediagram
    (
    @diagramname 	sysname,
		@owner_id		int	= null,
		@version 		int,
		@definition 	varbinary(max)
	)
	WITH EXECUTE AS 'dbo'
	AS
BEGIN
set nocount on

    declare @theId int
		declare @retval int
		declare @IsDbo	int
		declare @userName sysname
		if(@version is null or @diagramname is null)
begin
    RAISERROR (N'E_INVALIDARG', 16, 1);
return -1
		end

execute as caller;
select @theId = DATABASE_PRINCIPAL_ID();
select @IsDbo = IS_MEMBER(N'db_owner');
revert;

if @owner_id is null
begin
select @owner_id = @theId;
end
		else
begin
    if @theId <> @owner_id
begin
    if @IsDbo = 0
begin
    RAISERROR (N'E_INVALIDARG', 16, 1);
return -1
				end
select @theId = @owner_id
           end
    end
       -- next 2 line only for test, will be removed after define name unique
    if EXISTS(select diagram_id from dbo.sysdiagrams where principal_id = @theId and name = @diagramname)
begin
    RAISERROR ('The name is already used.', 16, 1);
return -2
		end

insert into dbo.sysdiagrams(name, principal_id , version, definition)
VALUES(@diagramname, @theId, @version, @definition) ;

select @retval = @@IDENTITY
    return @retval
    END
    go


CREATE PROCEDURE dbo.sp_dropdiagram
    (
    @diagramname 	sysname,
		@owner_id	int	= null
	)
	WITH EXECUTE AS 'dbo'
	AS
BEGIN
set nocount on
    declare @theId 			int
		declare @IsDbo 			int

		declare @UIDFound 		int
		declare @DiagId			int

		if(@diagramname is null)
begin
    RAISERROR ('Invalid value', 16, 1);
return -1
		end

EXECUTE AS CALLER;
select @theId = DATABASE_PRINCIPAL_ID();
select @IsDbo = IS_MEMBER(N'db_owner');
if(@owner_id is null)
select @owner_id = @theId;
REVERT;

select @DiagId = diagram_id, @UIDFound = principal_id from dbo.sysdiagrams where principal_id = @owner_id and name = @diagramname
    if(@DiagId IS NULL or (@IsDbo = 0 and @UIDFound <> @theId))
begin
    RAISERROR ('Diagram does not exist or you do not have permission.', 16, 1)
              return -3
		end

delete from dbo.sysdiagrams where diagram_id = @DiagId;

return 0;
END
go


CREATE PROCEDURE dbo.sp_helpdiagramdefinition
    (
    @diagramname 	sysname,
		@owner_id	int	= null
	)
	WITH EXECUTE AS N'dbo'
	AS
BEGIN
set nocount on

    declare @theId 		int
		declare @IsDbo 		int
		declare @DiagId		int
		declare @UIDFound	int

		if(@diagramname is null)
begin
    RAISERROR (N'E_INVALIDARG', 16, 1);
return -1
		end

execute as caller;
select @theId = DATABASE_PRINCIPAL_ID();
select @IsDbo = IS_MEMBER(N'db_owner');
if(@owner_id is null)
select @owner_id = @theId;
revert;

select @DiagId = diagram_id, @UIDFound = principal_id from dbo.sysdiagrams where principal_id = @owner_id and name = @diagramname;
if(@DiagId IS NULL or (@IsDbo = 0 and @UIDFound <> @theId ))
begin
    RAISERROR ('Diagram does not exist or you do not have permission.', 16, 1);
return -3
		end

select version, definition FROM dbo.sysdiagrams where diagram_id = @DiagId ;
return 0
	END
go


CREATE PROCEDURE dbo.sp_helpdiagrams
    (
    @diagramname sysname = NULL,
		@owner_id int = NULL
	)
	WITH EXECUTE AS N'dbo'
	AS
BEGIN
    DECLARE @user sysname
		DECLARE @dboLogin bit
EXECUTE AS CALLER;
SET @user = USER_NAME();
SET @dboLogin = CONVERT(bit,IS_MEMBER('db_owner'));
REVERT;
SELECT
    [Database] = DB_NAME(),
    [Name] = name,
    [ID] = diagram_id,
    [Owner] = USER_NAME(principal_id),
    [OwnerID] = principal_id
FROM
    sysdiagrams
WHERE
    (@dboLogin = 1 OR USER_NAME(principal_id) = @user) AND
    (@diagramname IS NULL OR name = @diagramname) AND
    (@owner_id IS NULL OR principal_id = @owner_id)
ORDER BY
    4, 5, 1
    END
    go


CREATE PROCEDURE dbo.sp_renamediagram
    (
    @diagramname 		sysname,
		@owner_id		int	= null,
		@new_diagramname	sysname

	)
	WITH EXECUTE AS 'dbo'
	AS
BEGIN
set nocount on
    declare @theId 			int
		declare @IsDbo 			int

		declare @UIDFound 		int
		declare @DiagId			int
		declare @DiagIdTarg		int
		declare @u_name			sysname
		if((@diagramname is null) or (@new_diagramname is null))
begin
    RAISERROR ('Invalid value', 16, 1);
return -1
		end

EXECUTE AS CALLER;
select @theId = DATABASE_PRINCIPAL_ID();
select @IsDbo = IS_MEMBER(N'db_owner');
if(@owner_id is null)
select @owner_id = @theId;
REVERT;

select @u_name = USER_NAME(@owner_id)

select @DiagId = diagram_id, @UIDFound = principal_id from dbo.sysdiagrams where principal_id = @owner_id and name = @diagramname
    if(@DiagId IS NULL or (@IsDbo = 0 and @UIDFound <> @theId))
begin
    RAISERROR ('Diagram does not exist or you do not have permission.', 16, 1)
              return -3
		end

		-- if((@u_name is not null) and (@new_diagramname = @diagramname))	-- nothing will change
		--	return 0;

		if(@u_name is null)
select @DiagIdTarg = diagram_id from dbo.sysdiagrams where principal_id = @theId and name = @new_diagramname
    else
select @DiagIdTarg = diagram_id from dbo.sysdiagrams where principal_id = @owner_id and name = @new_diagramname

    if((@DiagIdTarg is not null) and  @DiagId <> @DiagIdTarg)
begin
    RAISERROR ('The name is already used.', 16, 1);
return -2
		end

		if(@u_name is null)
update dbo.sysdiagrams set [name] = @new_diagramname, principal_id = @theId where diagram_id = @DiagId
    else
update dbo.sysdiagrams set [name] = @new_diagramname where diagram_id = @DiagId
    return 0
    END
    go


CREATE PROCEDURE dbo.sp_upgraddiagrams
    AS
BEGIN
    IF OBJECT_ID(N'dbo.sysdiagrams') IS NOT NULL
			return 0;

CREATE TABLE dbo.sysdiagrams
(
    name sysname NOT NULL,
    principal_id int NOT NULL,	-- we may change it to varbinary(85)
    diagram_id int PRIMARY KEY IDENTITY,
    version int,

    definition varbinary(max)
    CONSTRAINT UK_principal_name UNIQUE
    (
        principal_id,
        name
    )
);


/* Add this if we need to have some form of extended properties for diagrams */
/*
IF OBJECT_ID(N'dbo.sysdiagram_properties') IS NULL
BEGIN
    CREATE TABLE dbo.sysdiagram_properties
    (
        diagram_id int,
        name sysname,
        value varbinary(max) NOT NULL
    )
END
*/

IF OBJECT_ID(N'dbo.dtproperties') IS NOT NULL
begin
insert into dbo.sysdiagrams
(
    [name],
    [principal_id],
    [version],
[definition]
)
select
    convert(sysname, dgnm.[uvalue]),
    DATABASE_PRINCIPAL_ID(N'dbo'),			-- will change to the sid of sa
    0,							-- zero for old format, dgdef.[version],
    dgdef.[lvalue]
from dbo.[dtproperties] dgnm
    inner join dbo.[dtproperties] dggd on dggd.[property] = 'DtgSchemaGUID' and dggd.[objectid] = dgnm.[objectid]
    inner join dbo.[dtproperties] dgdef on dgdef.[property] = 'DtgSchemaDATA' and dgdef.[objectid] = dgnm.[objectid]

where dgnm.[property] = 'DtgSchemaNAME' and dggd.[uvalue] like N'_EA3E6268-D998-11CE-9454-00AA00A3F36E_'
    return 2;
end
		return 1;
END
go

