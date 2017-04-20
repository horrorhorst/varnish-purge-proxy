package providers


type IPProvider struct {
	Hosts []string;
}


func (a* IPProvider) GetPrivateIPs() []string {
	return a.Hosts;
}

func (a *IPProvider) Auth() error {
	return nil
}