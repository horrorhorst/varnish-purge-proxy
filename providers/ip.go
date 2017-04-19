package providers


type IPProvider struct {
	Hosts []string;
}


func (a* IPProvider) GetPrivateIPs() []string {
	return a.Hosts;
}
