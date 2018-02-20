package models

type RclList struct {
	INITIALEXECUTIONTIME string `json:"INITIALEXECUTIONTIME"`
	LISTCONFIGURATION struct {
		CONTACTROLES          int    `json:"CONTACTROLES"`
		DEFAULTGROUPBY        int    `json:"DEFAULTGROUPBY"`
		DEFAULTORDERBY        string `json:"DEFAULTORDERBY"`
		ISDISPLAYAGENCYTYPES  int    `json:"ISDISPLAYAGENCYTYPES"`
		ISDISPLAYCONTACTS     int    `json:"ISDISPLAYCONTACTS"`
		ISDISPLAYPROGRAMS     int    `json:"ISDISPLAYPROGRAMS"`
		ISNATIONAL            int    `json:"ISNATIONAL"`
		ISSHOWONWEBSITE       int    `json:"ISSHOWONWEBSITE"`
		ISSORTABLE            int    `json:"ISSORTABLE"`
		ISSTATE               int    `json:"ISSTATE"`
		MCODE                 string `json:"MCODE"`
		RELATEDLISTTYPE       string `json:"RELATEDLISTTYPE"`
		RESOURCES             string `json:"RESOURCES"`
		RESOURCESDESCR        string `json:"RESOURCES_DESCR"`
		RESOURCESTITLE        string `json:"RESOURCES_TITLE"`
		ROLORGLEVEL           string `json:"ROLORGLEVEL"`
		ROLPAGESECTION        string `json:"ROLPAGESECTION"`
		RSID                  int    `json:"RS_ID"`
		STDDISPLAYDESCRIPTION string `json:"STDDISPLAYDESCRIPTION"`
		STDDISPLAYID          int    `json:"STDDISPLAYID"`
		STDDISPLAYMETHODNAME  string `json:"STDDISPLAYMETHODNAME"`
		STDDISPLAYNAME        string `json:"STDDISPLAYNAME"`
		STDDISPLAYVIEWS       int    `json:"STD_DISPLAY_VIEWS"`
	} `json:"LISTCONFIGURATION"`
	LISTITEMS []struct {
		CONADDRESS1 string `json:"CONADDRESS1"`
		CONADDRESS2 string `json:"CONADDRESS2"`
		CONADDRESS3 string `json:"CONADDRESS3"`
		CONADDRESS4 string `json:"CONADDRESS4"`
		CONCITY     string `json:"CONCITY"`
		CONEMAIL    string `json:"CONEMAIL"`
		CONSTATE    string `json:"CONSTATE"`
		CONTACTID   int    `json:"CONTACTID"`
		CONTACTROLES []struct {
			CONTACTROLE     string `json:"CONTACTROLE"`
			CONTACTROLEDESC string `json:"CONTACTROLEDESC"`
			CONTACTROLEID   int    `json:"CONTACTROLEID"`
			CONTACTROLEURL  string `json:"CONTACTROLEURL"`
			ISDELETED       int    `json:"ISDELETED"`
			ISNFCAD         int    `json:"ISNFCAD"`
		} `json:"CONTACTROLES"`
		CONZIP                                    string `json:"CONZIP"`
		DSPSIMPLEHTMLVIEW                         string `json:"DSPSIMPLEHTMLVIEW"`
		DSPSTANDARDHTMLVIEW                       string `json:"DSPSTANDARDHTMLVIEW"`
		FIRSTNAME                                 string `json:"FIRST_NAME"`
		FUNCTIONALTITLE                           string `json:"FUNCTIONALTITLE"`
		GETROLORGSIMPLECONTACTINFODISPLAY         string `json:"GETROLORGSIMPLECONTACTINFODISPLAY"`
		GETROLORGSIMPLEWITHUSERCONTACTINFODISPLAY string `json:"GETROLORGSIMPLEWITHUSERCONTACTINFODISPLAY"`
		LASTNAME                                  string `json:"LAST_NAME"`
		NADID                                     int    `json:"NADID"`
		ORG struct {
			ACRONYM                string  `json:"ACRONYM"`
			ADDRESS1               string  `json:"ADDRESS1"`
			ADDRESS2               string  `json:"ADDRESS2"`
			ADDRESS3               string  `json:"ADDRESS3"`
			ADDRESS4               string  `json:"ADDRESS4"`
			CITY                   string  `json:"CITY"`
			CONTACTINFOUPDATEDDATE string  `json:"CONTACTINFOUPDATEDDATE"`
			COUNTIESSERVED         string  `json:"COUNTIES_SERVED"`
			EMAIL                  string  `json:"EMAIL"`
			FBURL                  string  `json:"FBURL"`
			ISHAGUEACCREDITED      int     `json:"ISHAGUEACCREDITED"`
			LATITUDE               float64 `json:"LATITUDE"`
			LOCALCHAPURL           string  `json:"LOCALCHAPURL"`
			LONGITUDE              float64 `json:"LONGITUDE"`
			NADID                  int     `json:"NADID"`
			NOTES                  string  `json:"NOTES"`
			ORGNAME1               string  `json:"ORGNAME1"`
			ORGNAME2               string  `json:"ORGNAME2"`
			ORGNAME3               string  `json:"ORGNAME3"`
			STATE                  string  `json:"STATE"`
			STATEABBR              string  `json:"STATE_ABBR"`
			TWITTERURL             string  `json:"TWITTERURL"`
			WEBURL                 string  `json:"WEBURL"`
			WEBURLSPANISH          string  `json:"WEBURLSPANISH"`
			YOUTUBEURL             string  `json:"YOUTUBEURL"`
			ZIP                    string  `json:"ZIP"`
		} `json:"ORG"`
		PHONENUMBERS []struct {
			DISPLAYSORTORDER int    `json:"DISPLAYSORTORDER"`
			PHONE            string `json:"PHONE"`
			PHONETYPEID      int    `json:"PHONETYPEID"`
			TYPENUM          string `json:"TYPENUM"`
		} `json:"PHONENUMBERS"`
		PRIMARYCONTACT                  int    `json:"PRIMARY_CONTACT"`
		RETURNCOMDISPLAYLISTITEMSSIMPLE string `json:"RETURNCOMDISPLAYLISTITEMSSIMPLE"`
		STATE                           string `json:"STATE"`
		TITLE                           string `json:"TITLE"`
	} `json:"LISTITEMS"`
}
