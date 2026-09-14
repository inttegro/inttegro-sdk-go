package payout

// SettingsLookupScheduleAgingSpec describes when balance transactions become
// eligible for payout in the settings read model.
type SettingsLookupScheduleAgingSpec struct {
	// Abide describes how the aging period is applied.
	Abide string `json:"abide"`

	// Label is the human-readable aging rule label.
	Label string `json:"label"`

	// TPlus is the required transaction age, such as "168h".
	TPlus string `json:"t_plus"`
}

// SettingsLookupSchedule is the active schedule returned by Service.Settings.
type SettingsLookupSchedule struct {
	// AgingSpec describes when balance transactions become eligible.
	AgingSpec SettingsLookupScheduleAgingSpec `json:"aging_spec"`

	// Description explains the schedule behavior.
	Description string `json:"description"`

	// Interval is the payout frequency.
	Interval string `json:"interval"`

	// Name is the schedule's display name.
	Name string `json:"name"`

	// ScheduleOn describes when automatic payouts run.
	ScheduleOn string `json:"schedule_on"`

	// Type identifies the schedule mode.
	Type string `json:"type"`
}

// SettingsMutationScheduleSpec describes the aging rule returned after a
// payout settings mutation.
type SettingsMutationScheduleSpec struct {
	Abide string `json:"abide"`
	ID    string `json:"id"`
	Label string `json:"label"`
	TPlus string `json:"t_plus"`
}

// SettingsMutationSchedule is the updated schedule returned after a payout
// settings mutation.
type SettingsMutationSchedule struct {
	Description string                       `json:"description"`
	ID          string                       `json:"id"`
	Interval    string                       `json:"interval"`
	Name        string                       `json:"name"`
	ScheduleOn  string                       `json:"schedule_on"`
	Spec        SettingsMutationScheduleSpec `json:"spec"`
	Type        string                       `json:"type"`
}
