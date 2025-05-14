package exec_svc

type ServiceGroup struct {
	ScriptService
	ScriptLogService
	ScriptLibraryService
	TransferService
	TaskManageService
	ManageLogService
	TaskLogService
	CronService
	CronLogService
}
