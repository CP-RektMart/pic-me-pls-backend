package report

// @Summary     Get report by id
// @Description Get a report of a user by id
// @Tags		customer
// @Router      /api/v1/reports [GET]
// @Security    ApiKeyAuth
// @Param       body  body  dto.CreateReportRequest  true  "Report details"
// @Success     200
// @Failure     400   {object}  dto.HttpError
// @Failure     401   {object}  dto.HttpError
// @Failure     403   {object}  dto.HttpError
// @Failure     404   {object}  dto.HttpError
// @Failure     500   {object}  dto.HttpError
