package command

import "github.com/vishvananda/netlink"

type LinkShowResult struct {
	*netlink.LinkAttrs
	Type         string
	HardwareAddr string
}

func LinkShow() error {
	results, err := GetLink()
	if err != nil {
		return err
	}
	toJson(results)
	return err
}

func GetLink() ([]LinkShowResult, error) {
	linkShowResults := []LinkShowResult{}
	links, err := netlink.LinkList()
	if err != nil {
		return nil, err
	}
	for _, link := range links {
		linkAttrs := link.Attrs()
		linkShowResult := LinkShowResult{linkAttrs, link.Type(), linkAttrs.HardwareAddr.String()}
		linkShowResults = append(linkShowResults, linkShowResult)
	}
	return linkShowResults, nil
}
