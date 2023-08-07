package option

func (s *ServerRunOptions) Validate() []error {
	var errors []error

	errors = append(errors, s.KubernetesOption.Validate()...)

	return errors
}
