package constant

const (
	ContextDb     = "contextdb"
	ContextClaims = "contextclaims"
	ContextError  = "contexterror"
)

const (
	APPLY_STATUS_PENDING int = 1
	APPLY_STATUS_PASS    int = 2
	APPLY_STATUS_FORBID  int = 3
)

const (
	USER_PERMISSION_TYPE_COLUMNS int = 1
	USER_PERMISSION_TYPE_MODEL   int = 2
)

const (
	TASK_API_METHOD_RUN     = "RunTask"
	TASK_API_METHOD_DEL     = "DeleteRows"
	TASK_API_METHOD_CANCEl  = "CancelTask"
	TASK_API_METHOD_RESTORE = "RestoreTask"
)

const (
	TASK_API_RUNTYPE_GETREQVALUE = "GetReqValue"
	TASK_API_RUNTYPE_RUNMODEL    = "RunModel"
)

const (
	DB_TYPE_MYSQL  string = "mysql"
	DB_TYPE_ORACLE string = "oracle"
	DB_TYPE_PG     string = "postgre"
)

const (
	GLOBAL_MODEL  string = "global"
	PORTION_MODEL string = "portion"
)
const (
	ENTER_IDENTITY_ALL   string = "all"
	ENTER_IDENTITY_BACK  string = "back"
	ENTER_IDENTITY_FRONT string = "front"
	ENTER_IDENTITY_APP   string = "app"
)

const (
	ENTRANCE_PLATFORM_TRITIUM  = "tritium"
	ENTRANCE_PLATFORM_EXCHANGE = "exchange"
	ENTRANCE_PLATFORM_ALL      = "all"
)
