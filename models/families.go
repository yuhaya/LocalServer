package models

type Families struct {
	Id         uint64 `orm:"fk;auto"`
	Name	   string `orm:"size(50)"`
	Guid       string `orm:"unique;size(50)"`
	FirstGuid  string `orm:"size(50)"`
	SchoolGuid string `orm:"size(50)"`
}


type FamiliyModel struct{

}

func (this *FamiliyModel) GetListByFamlilyName(name string) []*Families{
	
	var family_slice []*Families

	o := orm.NewOrm()
	var r RawSeter
	var lists []orm.ParamsList
	r = o.Raw("SELECT f.*,u.guid as user_guid,u.phone as user_guid,u.realname as user_realname, FROM ittr_families AS f LEFT JOIN ittr_users as u ON f.first_guid = u.guid WHERE f.name LIKE '%?%'", name)
	num, err = r.Values(&lists)
	
}