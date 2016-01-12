package objects

type Catalog struct {
	Services []Service `yaml:"services"`
}

type Service struct {
	Plans         []Plan            `yaml:"plans",json:"plans"`
	Id            string            `yaml:"id",json:"id"`
	Name          string            `yaml:"name",json:"name"`
	Description   string            `yaml:"description",json:"description"`
	Bindable      bool              `yaml:"bindable",json:"bindable"`
	Tags          []string          `yaml:"tags",json:"tags"`
	Metadata      map[string]string `yaml:"metadata",json:"metadata"`
	Requires      []string          `yaml:"requires",json:"requires"`
	PlanUpdatable bool              `yaml:"plan_updatable",json:"plan_updatable"`
}

type Plan struct {
	Id              string            `yaml:"id",json:"id"`
	Name            string            `yaml:"name",json:"name"`
	Description     string            `yaml:"description",json:"description"`
	Metadata        map[string]string `yaml:"metadata",json:"metadata"`
	Free            bool              `yaml:"free",json:"free"`
	DashboardClient DashboardClient   `yaml:"dashboard_client",json:"dashboard_client"`
}

type DashboardClient struct {
	Id          string `yaml:"id",json:"id"`
	Secret      string `yaml:"secret",json:"secret"`
	RedirectUrl string `yaml:"redirect_url",json:"redirect_url"`
}
