// PromotionSpec defines the desired state of Promotion
type PromotionSpec struct {
    // The source environment to promote from.
    // +required
    SourceEnvironmentRef *corev1.LocalObjectReference `json:"sourceEnvironmentRef"`

    // The target environment to promote to.
    // +required
    TargetEnvironmentRef *corev1.LocalObjectReference `json:"targetEnvironmentRef"`

    // Copy defines a list of copy operations to perform.
    // +required
    Copy []CopyOperation `json:"copy"`

    // Strategy defines the strategy to use when promoting.
    // +required
    // +kubebuilder:validation:Enum=pull-request
    Strategy string `json:"strategy"`
}

// CopyOperation defines a file/directory copy operation.
type CopyOperation struct {
    // Name is the name you want to give this copy operation.
    // E.g. "Application Version"
    // +required
    Name string `json:"name"`

    // The source path to copy from.
    // +required
    Source string `json:"source"`

    // The target path to copy to.
    // +required
    Target string `json:"target"`
}
