// EnvironmentSpec defines the desired state of Environment
type EnvironmentSpec struct {
    // Path is the filesystem path to the environment directory
    // relative from the root of the source repository.
    // Defaults to the root of the repository.
    // +optional
    Path string `json:"path,omitempty"`

    // Source defines the source repository of the environment.
    // +required
    Source Source `json:"source"`

    // ApiTokenSecretRef refers to a secret containing the API token
    // needed for doing pull requests.
    // Its a generic secret with the key "token".
    // +optional
    ApiTokenSecretRef *corev1.LocalObjectReference `json:"apiTokenSecretRef,omitempty"`

    // GitProvider is the name of the git provider.
    // Required for pull request strategy.
    // +Kubebuilder:Validation:Enum=github
    // +optional
    GitProvider string `json:"gitProvider"`
}

// Source defines the source repository of the environment.
type Source struct {
    // URL is the URL of the source repository.
    // +required
    URL string `json:"url"`

    // Ref defines the git reference to use.
    // Defaults to the "master" branch.
    // +optional
    Reference *GitRepositoryRef `json:"ref,omitempty"`

    // SecretRef is the name of the secret containing the credentials
    // to access the source repository.
    // +optional
    SecretRef *corev1.LocalObjectReference `json:"secretRef,omitempty"`
}

// GitRepositoryRef specifies the Git reference to resolve and checkout.
type GitRepositoryRef struct {
    // Branch to check out, defaults to 'master' if no other field is defined.
    // +optional
    Branch string `json:"branch,omitempty"`
}
