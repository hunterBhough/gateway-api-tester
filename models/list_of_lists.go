package models

type ListOfLists struct {
	Lists []struct {
		DefaultGroupBy       int    `json:"DEFAULTGROUPBY"`
		DefaultOrderBy       string `json:"DEFAULTORDERBY"`
		IsDisplayAgencyTypes int    `json:"ISDISPLAYAGENCYTYPES"`
		IsDisplayContacts    int    `json:"ISDISPLAYCONTACTS"`
		IsDisplayPrograms    int    `json:"ISDISPLAYPROGRAMS"`
		IsNational           int    `json:"ISNATIONAL"`
		IsShowOnWebsite      int    `json:"ISSHOWONWEBSITE"`
		IsSortable           int    `json:"ISSORTABLE"`
		IsState              int    `json:"ISSTATE"`
		Mcode                string `json:"MCODE"`
		RelatedListType      string `json:"RELATEDLISTTYPE"`
		Resources            string `json:"RESOURCES"`
		ResourcesDescription string `json:"RESOURCES_DESCR"`
		ResourcesTitle       string `json:"RESOURCES_TITLE"`
		RolOrgLevel          string `json:"ROLORGLEVEL"`
		RSID                 int    `json:"RS_ID"`
		StdDisplayViews      int    `json:"STD_DISPLAY_VIEWS"`
	} `json:"LISTS"`
	Time string `json:"TIME"`
}
