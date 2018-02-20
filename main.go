package main

import (
	"github.com/hunterBhough/gateway-api-tester/models"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"
	"errors"
)

var LIST_URL = "https://cwigapis.icfwebservices.com/api/organizations/v1/"

func main() {
	relatedLists, err := buildRelatedLists()

	if err != nil {
		fmt.Println(err)
	} else {
		err = runLists(relatedLists)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("API is passing all tests")
		}
	}

}

func runLists(relatedLists models.ListOfLists) error {
	for _, list := range relatedLists.Lists {
		switch list.RelatedListType {
		case "RCL":
			contactList, err := buildContactList(list.RSID)
			if err != nil {
				return err
			}
			err = compare(list.RSID, list.RelatedListType, list.ResourcesTitle, contactList.LISTCONFIGURATION.RELATEDLISTTYPE, contactList.LISTCONFIGURATION.RESOURCESTITLE)
			if err != nil {
				return err
			}
		case "ROL":
			orgList, err := buildOrgList(list.RSID)
			if err != nil {
				return err
			}
			err = compare(list.RSID, list.RelatedListType, list.ResourcesTitle, orgList.LISTCONFIGURATION.RELATEDLISTTYPE, orgList.LISTCONFIGURATION.RESOURCESTITLE)
			if err != nil {
				return err
			}
		default:
			err := errors.New("could not determine list type")
			return err
		}
	}
	return nil
}

func compare(id int, relatedListType string, relatedListTitle string, listType string, listTitle string) error {
	if relatedListType != listType {
		return errors.New("the master list type, " + relatedListType + ", is not equal to the list type, " + listType + " at item #" + strconv.Itoa(id))
	}
	if relatedListTitle != listTitle {
		return errors.New("the master list title, " + relatedListTitle + ", is not equal to the list title, " + listTitle + " at item #" + strconv.Itoa(id))
	}
	return nil
}

func buildOrgList(id int) (models.RolList, error) {
	u := LIST_URL + strconv.Itoa(id)
	var response models.RolList

	resp, e := http.Get(u)
	if e != nil {
		return response, e
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	e = json.Unmarshal(body, &response)
	if e != nil {
		return response, e
	}

	return response, nil
}

func buildContactList(id int) (models.RclList, error) {
	u := LIST_URL + strconv.Itoa(id)
	var response models.RclList

	resp, e := http.Get(u)
	if e != nil {
		return response, e
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	e = json.Unmarshal(body, &response)
	if e != nil {
		return response, e
	}

	return response, nil
}

func buildRelatedLists() (models.ListOfLists, error) {
	u := "https://cwigapis.icfwebservices.com/api/relatedLists/v1/"
	var response models.ListOfLists

	resp, e := http.Get(u)
	if e != nil {
		return response, e
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	e = json.Unmarshal(body, &response)
	if e != nil {
		return response, e
	}

	return response, nil
}
