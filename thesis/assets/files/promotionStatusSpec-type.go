// PromotionStatus defines the observed state of Promotion
type PromotionStatus struct {
    // ObservedGeneration is the last observed generation of the Promotion
    // object.
    // +optional
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`

    // Conditions is a list of the current conditions of the Promotion.
    // +optional
    Conditions []metav1.Condition `json:"conditions,omitempty"`

    // LastPullRequestURL is the URL of the pull request created by the promotion.
    // +optional
    LastPullRequestURL string `json:"lastPullRequestUrl,omitempty"`

    // LastPullRequestNumber is the number of the pull request created by the promotion.
    // +optional
    LastPullRequestNumber int `json:"lastPullRequestNumber,omitempty"`
}
