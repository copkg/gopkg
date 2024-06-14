package schema

type Customer struct {
	CustNo              string `json:"cust_no" db:"cust_no"`
	CustIdCard          string `json:"cust_id_card" db:"cust_id_card"`
	CustName            string `json:"cust_name" db:"cust_name"`
	CustGender          int64  `json:"cust_gender" db:"cust_gender"`
	CustBirthday        string `json:"cust_birthday" db:"cust_birthday"`
	CustMarital         string `json:"cust_marital" db:"cust_marital"`
	CustBlood           string `json:"cust_blood" db:"cust_blood"`
	CustNativePlace     string `json:"cust_native_place" db:"cust_native_place"`
	CustNationality     string `json:"cust_nationality" db:"cust_nationality"`
	CustNation          string `json:"cust_nation" db:"cust_nation"`
	CustAddress         string `json:"cust_address" db:"cust_address"`
	CustProvince        string `json:"cust_province" db:"cust_province"`
	CustCity            string `json:"cust_city" db:"cust_city"`
	CustDistrict        string `json:"cust_district" db:"cust_district"`
	CustOccupation      string `json:"cust_occpation" db:"cust_occpation"`
	CustEdu             string `json:"cust_edu" db:"cust_edu"`
	CustFirstRecordDate string `json:"cust_first_record_date" db:"cust_first_record_date"`
	CustLastRecordDate  string `json:"cust_last_record_date" db:"cust_last_record_date"`
	CustVisitNum        int64  `json:"cust_visit_num" db:"cust_visit_num"`
	CustInNum           int64  `json:"cust_in_num" db:"cust_in_num"`
	CustAllergy         string `json:"cust_allergy" db:"cust_allergy"`
	CustNcd             string `json:"cust_ncd" db:"cust_ncd"`
	CustWeight          string `json:"cust_weight" db:"cust_weight"`
	CustHeight          string `json:"cust_height" db:"cust_height"`
	CustMobile          string `json:"cust_mobile" db:"cust_mobile"`
}

type CustomerListRequest struct {
	Page                   int64    `json:"page"`
	Size                   int64    `json:"size"` // 每页大小
	CustNo                 []string `json:"cust_no"`
	CustName               string   `json:"cust_name"`
	CustIdCard             string   `json:"cust_id_card"`
	CustMobile             string   `json:"cust_mobile"`
	Gender                 []int64  `json:"gender"`
	Age                    []string `json:"age"`
	Marital                []string `json:"marital"`
	Ncd                    []string `json:"ncd"`
	Tags                   []string `json:"tags"`
	CustProvince           []string `json:"cust_province"`
	CustCity               []string `json:"cust_city"`
	CustDistrict           []string `json:"cust_district"`
	CustVisitNumMin        int64    `json:"cust_visit_num_min"`
	CustVisitNumMax        int64    `json:"cust_visit_num_max"`
	CustInNumMin           int64    `json:"cust_in_num_min"`
	CustInNumMax           int64    `json:"cust_in_num_max"`
	CustBlood              []string `json:"cust_blood"`
	CustAddress            string   `json:"cust_address"`
	CustNativePlace        []string `json:"cust_native_place"`
	CustOccpation          []string `json:"cust_occpation"`
	CustEdu                []string `json:"cust_edu"`
	CustFirstRecordDateMin string   `json:"cust_first_record_date_min"`
	CustFirstRecordDateMax string   `json:"cust_first_record_date_max"`
	CustLastRecordDateMin  string   `json:"cust_last_record_date_min"`
	CustLastRecordDateMax  string   `json:"cust_last_record_date_max"`
	TID                    int64    `json:"tid"`
}

type CustomerListResponse struct {
	Customers []*Customer `json:"customers"`
	*Error
	Total int64 `json:"total"`
}

type CustomerMobileResponse struct {
	Mobiles []string `json:"mobiles,omitempty"`
	*Error
}
