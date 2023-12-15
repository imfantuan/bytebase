package iam

type Permission string

const (
	PermissionInstancesList     Permission = "bb.instances.list"
	PermissionInstancesGet      Permission = "bb.instances.get"
	PermissionInstancesCreate   Permission = "bb.instances.create"
	PermissionInstancesUpdate   Permission = "bb.instances.update"
	PermissionInstancesDelete   Permission = "bb.instances.delete"
	PermissionInstancesUndelete Permission = "bb.instances.undelete"
	PermissionInstancesSync     Permission = "bb.instances.sync"

	PermissionDatabasesList                Permission = "bb.databases.list"
	PermissionDatabasesGet                 Permission = "bb.databases.get"
	PermissionDatabasesUpdate              Permission = "bb.databases.update"
	PermissionDatabasesSync                Permission = "bb.databases.sync"
	PermissionDatabasesGetMetadata         Permission = "bb.databases.getMetadata"
	PermissionDatabasesUpdateMetadata      Permission = "bb.databases.updateMetadata"
	PermissionDatabasesGetSchema           Permission = "bb.databases.getSchema"
	PermissionDatabasesGetBackupSetting    Permission = "bb.databases.getBackupSetting"
	PermissionDatabasesUpdateBackupSetting Permission = "bb.databases.updateBackupSetting"
	PermissionBackupsList                  Permission = "bb.backups.list"
	PermissionBackupsCreate                Permission = "bb.backups.create"
	PermissionChangeHistoriesList          Permission = "bb.changeHistories.list"
	PermissionChangeHistoriesGet           Permission = "bb.changeHistories.get"
	PermissionDatabaseSecretsList          Permission = "bb.databaseSecrets.list"
	PermissionDatabaseSecretsUpdate        Permission = "bb.databaseSecrets.update"
	PermissionDatabaseSecretsDelete        Permission = "bb.databaseSecrets.delete"
	PermissionSlowQueriesList              Permission = "bb.slowQueries.list"

	PermissionEnvironmentsList     Permission = "bb.environments.list"
	PermissionEnvironmentsGet      Permission = "bb.environments.get"
	PermissionEnvironmentsCreate   Permission = "bb.environments.create"
	PermissionEnvironmentsUpdate   Permission = "bb.environments.update"
	PermissionEnvironmentsDelete   Permission = "bb.environments.delete"
	PermissionEnvironmentsUndelete Permission = "bb.environments.undelete"

	PermissionIssuesList          Permission = "bb.issues.list"
	PermissionIssuesGet           Permission = "bb.issues.get"
	PermissionIssuesCreate        Permission = "bb.issues.create"
	PermissionIssuesUpdate        Permission = "bb.issues.update"
	PermissionIssueCommentsCreate Permission = "bb.issueComments.create"
	PermissionIssueCommentsUpdate Permission = "bb.issueComments.update"

	// Project Service.
	PermissionProjectsList         Permission = "bb.projects.list"
	PermissionProjectsGet          Permission = "bb.projects.get"
	PermissionProjectsCreate       Permission = "bb.projects.create"
	PermissionProjectsUpdate       Permission = "bb.projects.update"
	PermissionProjectsDelete       Permission = "bb.projects.delete"
	PermissionProjectsUndelete     Permission = "bb.projects.undelete"
	PermissionProjectsGetIAMPolicy Permission = "bb.projects.getIamPolicy"
	PermissionProjectsSetIAMPolicy Permission = "bb.projects.setIamPolicy"

	PermissionRisksList   Permission = "bb.risks.list"
	PermissionRisksCreate Permission = "bb.risks.create"
	PermissionRisksUpdate Permission = "bb.risks.update"
	PermissionRisksDelete Permission = "bb.risks.delete"

	PermissionRolesList   Permission = "bb.roles.list"
	PermissionRolesCreate Permission = "bb.roles.create"
	PermissionRolesUpdate Permission = "bb.roles.update"
	PermissionRolesDelete Permission = "bb.roles.delete"

	PermissionChangelistsList   Permission = "bb.changelists.list"
	PermissionChangelistsGet    Permission = "bb.changelists.get"
	PermissionChangelistsUpdate Permission = "bb.changelists.update"
	PermissionChangelistsCreate Permission = "bb.changelists.create"
	PermissionChangelistsDelete Permission = "bb.changelists.delete"

	PermissionInstanceRolesList     Permission = "bb.instanceRoles.list"
	PermissionInstanceRolesGet      Permission = "bb.instanceRoles.get"
	PermissionInstanceRolesCreate   Permission = "bb.instanceRoles.create"
	PermissionInstanceRolesUpdate   Permission = "bb.instanceRoles.update"
	PermissionInstanceRolesDelete   Permission = "bb.instanceRoles.delete"
	PermissionInstanceRolesUndelete Permission = "bb.instanceRoles.undelete"

	PermissionExternalVersionControlsGet            Permission = "bb.externalVersionControls.get"
	PermissionExternalVersionControlsList           Permission = "bb.externalVersionControls.list"
	PermissionExternalVersionControlsCreate         Permission = "bb.externalVersionControls.create"
	PermissionExternalVersionControlsUpdate         Permission = "bb.externalVersionControls.update"
	PermissionExternalVersionControlsDelete         Permission = "bb.externalVersionControls.delete"
	PermissionExternalVersionControlsSearchProjects Permission = "bb.externalVersionControls.searchProjects"
	PermissionExternalVersionControlsListProjects   Permission = "bb.externalVersionControls.listProjects"
	PermissionPlansList                             Permission = "bb.plans.list"
	PermissionPlansGet                              Permission = "bb.plans.get"
	PermissionPlansCreate                           Permission = "bb.plans.create"
	PermissionPlansUpdate                           Permission = "bb.plans.update"
	PermissionRolloutsGet                           Permission = "bb.rollouts.get"
	PermissionRolloutsCreate                        Permission = "bb.rollouts.create"
	PermissionRolloutsPreview                       Permission = "bb.rollouts.preview"
	PermissionTaskRunsList                          Permission = "bb.taskRuns.list"
	PermissionPlanCheckRunsList                     Permission = "bb.planCheckRuns.list"
	PermissionPlanCheckRunsRun                      Permission = "bb.planCheckRuns.run"
	PermissionTasksRun                              Permission = "bb.tasks.run"
	PermissionTasksSkip                             Permission = "bb.tasks.skip"
	PermissionTaskRunsCancel                        Permission = "bb.taskRuns.cancel"

	PermissionSettingsList Permission = "bb.settings.list"
	PermissionSettingsGet  Permission = "bb.settings.get"
	PermissionSettingsSet  Permission = "bb.settings.set"

	PermissionPoliciesList   Permission = "bb.policies.list"
	PermissionPoliciesGet    Permission = "bb.policies.get"
	PermissionPoliciesCreate Permission = "bb.policies.create"
	PermissionPoliciesUpdate Permission = "bb.policies.update"
	PermissionPoliciesDelete Permission = "bb.policies.delete"
)
