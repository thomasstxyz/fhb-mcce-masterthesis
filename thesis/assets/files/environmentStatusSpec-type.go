// EnvironmentStatus defines the observed state of Environment
type EnvironmentStatus struct {
    // ObservedGeneration is the last observed generation of the Environment
    // object.
    // +optional
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`

    // Conditions is a list of the current conditions of the Environment.
    // +optional
    Conditions []metav1.Condition `json:"conditions,omitempty"`

    // ObservedCommitHash is the last observed commit hash of the Environment
    // object.
    // +optional
    ObservedCommitHash string `json:"observedCommitHash,omitempty"`
}
